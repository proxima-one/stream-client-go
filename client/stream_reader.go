package client

import (
	"context"
	"errors"
	"github.com/proxima-one/streamdb-client-go/config"
	"github.com/proxima-one/streamdb-client-go/model"
	"log"
	"sync"
	//"google.golang.org/grpc/credentials" todo: add ssl and auth
)

type StreamReader struct {
	config         config.Config
	client         *ProximaClient
	lastState      model.State
	mutex          sync.Mutex
	preprocessFunc TransitionPreprocessingFunc

	processedCount int
	workingChan    *chan *ProximaStreamObject

	runningStream *chan *ProximaStreamObject
	errStream     *chan error
	ctx           context.Context
}

func NewStreamReader(config config.Config, preprocess TransitionPreprocessingFunc) (*StreamReader, error) {
	res := &StreamReader{
		config:         config,
		lastState:      config.GetState(),
		preprocessFunc: preprocess,
	}
	res.client = NewProximaClient(res.config)
	err := res.client.Connect()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (reader *StreamReader) GetBufferLoadPercent() int {
	if reader.workingChan == nil {
		return 0
	}
	load := float32(len(*reader.workingChan)) / float32(reader.config.GetChannelSize()) * 100
	return int(load)
}

func (reader *StreamReader) Metrics() (int, int) {
	return reader.processedCount, reader.GetBufferLoadPercent()
}

func (reader *StreamReader) GetStreamID() string {
	return reader.config.GetStreamID()
}

func (reader *StreamReader) FetchNextTransitions(ctx context.Context, count int) ([]*model.Transition, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	reader.mutex.Lock()
	transitions, err := reader.client.GetTransitionsAfter(ctx, model.StreamState{
		StreamID: reader.config.GetStreamID(),
		State:    reader.lastState,
	}, count)
	if err != nil {
		log.Println("Error fetching transitions: ", err)
		return nil, err
	}
	if len(transitions) == 0 {
		reader.mutex.Unlock()
		return nil, errors.New("no transitions found")
	}
	reader.lastState = transitions[len(transitions)-1].NewState
	reader.processedCount += len(transitions)
	reader.mutex.Unlock()
	return transitions, err
}

func (reader *StreamReader) GetRawStreamFromState(ctx context.Context,
	state model.State) (<-chan *model.Transition, <-chan error, error) {

	if ctx == nil {
		ctx = context.Background()
	}
	stream, errc, err := reader.client.GetStream(ctx, model.StreamState{
		StreamID: reader.config.GetStreamID(),
		State:    state,
	})
	return stream, errc, err
}

func (reader *StreamReader) streamObjForTransition(transition *model.Transition) *ProximaStreamObject {
	if reader.preprocessFunc == nil {
		return &ProximaStreamObject{
			Transition: transition,
			Preprocess: nil,
		}
	}
	return &ProximaStreamObject{
		Transition: transition,
		Preprocess: NewTransitionPreprocessingResult(transition, reader.preprocessFunc),
	}
}

func (reader *StreamReader) saveWorkingChan(newChan *chan *ProximaStreamObject) {
	if reader.workingChan != nil {
		log.Println("Warning: previous stream is still working, cancel context to close it")
		panic("reader is already working") //todo: add log levels and close previous stream, cancel context
	}
	reader.workingChan = newChan
}

func (reader *StreamReader) StartGrpcStreamChannel(ctx context.Context) (<-chan *ProximaStreamObject, <-chan error, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	stream, errc, err := reader.GetRawStreamFromState(ctx, reader.lastState)
	buffer := reader.config.GetChannelSize()
	result := make(chan *ProximaStreamObject, buffer)
	errChan := make(chan error, 1)
	reader.saveWorkingChan(&result)
	go func() {
		defer close(result)
		for {
			select {
			case <-ctx.Done():
				reader.workingChan = nil
				return
			case msg, ok := <-stream:
				if !ok {
					return
				}
				result <- reader.streamObjForTransition(msg)
				reader.lastState = msg.NewState
				reader.processedCount++

			case err := <-errc:
				errChan <- err
			}
		}
	}()
	reader.ctx = ctx
	reader.runningStream = &result
	reader.errStream = &errChan
	return result, errc, err
}

func (reader *StreamReader) StartGrpcRpcChannel(ctx context.Context,
	count int) (<-chan *ProximaStreamObject, <-chan error, error) {

	if ctx == nil {
		ctx = context.Background()
	}
	buffer := reader.config.GetChannelSize()
	result := make(chan *ProximaStreamObject, buffer)
	errc := make(chan error, 1)
	reader.saveWorkingChan(&result)
	go func() {
		defer close(result)
		defer close(errc)
		for {
			select {
			case <-ctx.Done():
				reader.workingChan = nil
				return
			default:
				messages, err := reader.FetchNextTransitions(ctx, count)
				if err != nil {
					errc <- err
					return
				}
				for _, msg := range messages {
					result <- reader.streamObjForTransition(msg)
				}
			}
		}
	}()
	reader.ctx = ctx
	reader.runningStream = &result
	reader.errStream = &errc
	return result, errc, nil
}

func (reader *StreamReader) Start(ctx context.Context) error {
	_, _, err := reader.StartGrpcStreamChannel(ctx) //default implementation uses grpc stream
	return err
}

func (reader *StreamReader) ReadNext() (*ProximaStreamObject, error) {
	select {
	case <-reader.ctx.Done():
		return nil, reader.ctx.Err()

	case obj := <-*reader.runningStream:
		if obj.Preprocess.err != nil {
			return nil, obj.Preprocess.err
		}
		return obj, nil

	case err := <-*reader.errStream:
		return nil, err

	}
}

func (reader *StreamReader) Close() error {
	return reader.client.Close()
}
