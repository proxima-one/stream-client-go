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
	config         *config.Config
	client         *ProximaClient
	lastState      model.State
	mutex          sync.Mutex
	preprocessFunc model.TransitionPreprocessingFunc
	bufferLoad     float32
}

type ProximaStreamObject struct {
	Transition *model.Transition
	Preprocess *model.TransitionPreprocessingResult
}

func NewStreamReader(config config.Config,
	lastState model.State,
	preprocess model.TransitionPreprocessingFunc) (*StreamReader, error) {

	res := &StreamReader{
		config:         &config,
		lastState:      lastState,
		preprocessFunc: preprocess,
	}
	res.client = NewProximaClient(res.config)
	err := res.client.Connect()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (reader *StreamReader) Close() error {
	return reader.client.Close()
}

func (reader *StreamReader) GetStreamID() string {
	return reader.config.StreamID
}

func (reader *StreamReader) FetchNextTransitions(ctx context.Context, count int) ([]*model.Transition, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	reader.mutex.Lock()
	transitions, err := reader.client.GetTransitionsAfter(ctx, model.StreamState{
		StreamID: reader.config.StreamID,
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
	//todo: transitions is not empty
	reader.lastState = transitions[len(transitions)-1].NewState
	reader.mutex.Unlock()
	return transitions, err
}

func (reader *StreamReader) GetRawStreamFromState(ctx context.Context,
	state model.State, buffer int) (<-chan *model.Transition, <-chan error, error) {

	if ctx == nil {
		ctx = context.Background()
	}
	stream, errc, err := reader.client.GetStream(ctx, model.StreamState{
		StreamID: reader.config.StreamID,
		State:    state,
	}, buffer)
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
		Preprocess: model.NewTransitionPreprocessingResult(transition, reader.preprocessFunc),
	}
}

func (reader *StreamReader) GetRawStream(ctx context.Context,
	buffer int) (<-chan *ProximaStreamObject, <-chan error, error) {

	if ctx == nil {
		ctx = context.Background()
	}
	stream, errc, err := reader.GetRawStreamFromState(ctx, reader.lastState, buffer)
	result := make(chan *ProximaStreamObject, buffer)
	go func() {
		defer close(result)
		for {
			select {
			case <-ctx.Done():
				return
			case msg, ok := <-stream:
				if !ok {
					return
				}
				reader.lastState = msg.NewState
				reader.bufferLoad = float32(len(result)) / float32(buffer)
				result <- reader.streamObjForTransition(msg)
			}
		}
	}()
	return result, errc, err
}

func (reader *StreamReader) GetStreamBufferLoad() float32 {
	return reader.bufferLoad
}

func (reader *StreamReader) GetBatchedStream(ctx context.Context,
	buffer int, count int) (<-chan *ProximaStreamObject, <-chan error, error) {

	if ctx == nil {
		ctx = context.Background()
	}
	result := make(chan *ProximaStreamObject, buffer)
	errc := make(chan error, 1)
	go func() {
		defer close(result)
		defer close(errc)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				messages, err := reader.FetchNextTransitions(ctx, count)
				if err != nil {
					log.Printf("Error while receiving stream, %v\n", err)
					errc <- err
					return
				}
				reader.bufferLoad = float32(len(result)) / float32(buffer)
				for _, msg := range messages {
					result <- reader.streamObjForTransition(msg)
				}
			}
		}
	}()
	return result, errc, nil
}
