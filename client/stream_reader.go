package client

import (
	"context"
	"errors"
	"github.com/proxima-one/pocs/stream-db-endpoint/config"
	"github.com/proxima-one/pocs/stream-db-endpoint/model"
	"log"
	"sync"
	//"google.golang.org/grpc/credentials" todo: add ssl and auth
)

type StreamReader struct {
	config    *config.Config
	client    *ProximaClient
	lastState model.State
	mutex     sync.Mutex
}

func NewStreamReader(config config.Config, lastState model.State) *StreamReader {
	return &StreamReader{
		config:    &config,
		lastState: lastState,
	}
}

func (reader *StreamReader) Connect() {
	reader.client = NewProximaClient(reader.config)
	reader.client.Connect()
}

func (reader *StreamReader) Disconnect() {
	reader.client.Disconnect()
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

func (reader *StreamReader) GetRawStreamFromState(ctx context.Context, state model.State, buffer int) (<-chan *model.Transition, <-chan error, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	stream, errc, err := reader.client.GetStream(ctx, model.StreamState{
		StreamID: reader.config.StreamID,
		State:    state,
	}, buffer)
	return stream, errc, err
}

func (reader *StreamReader) GetRawStream(ctx context.Context, buffer int) (<-chan *model.Transition, <-chan error, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	stream, errc, err := reader.GetRawStreamFromState(ctx, reader.lastState, buffer)
	result := make(chan *model.Transition, buffer)
	go func() {
		defer close(result)
		for {
			select {
			case <-ctx.Done():
				return
			case transition, ok := <-stream:
				if !ok {
					return
				}
				reader.lastState = transition.NewState
				result <- transition
			}
		}
	}()
	return result, errc, err
}

func (reader *StreamReader) GetBatchedStream(ctx context.Context, buffer int, count int) (<-chan *model.Transition, <-chan error, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	result := make(chan *model.Transition, buffer)
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
				for _, msg := range messages {
					result <- msg
				}
			}
		}
	}()
	return result, errc, nil
}
