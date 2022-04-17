package client

import (
	"errors"
	"github.com/proxima-one/pocs/stream-db-endpoint/config"
	"github.com/proxima-one/pocs/stream-db-endpoint/model"
	"log"
	"sync"
	//"google.golang.org/grpc/credentials"
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

func (reader *StreamReader) FetchNextTransitions(count int) ([]model.Transition, error) {
	reader.mutex.Lock()
	transitions, err := reader.client.GetTransitionsAfter(model.StreamState{
		StreamID: reader.config.StreamID,
		State:    reader.lastState,
	}, count)
	if err != nil {
		log.Println("Error fetching transitions:", err)
		return nil, err
	}
	if len(transitions) == 0 {
		reader.mutex.Unlock()
		return nil, errors.New("No transitions found")
	}
	reader.lastState = transitions[len(transitions)-1].NewState
	reader.mutex.Unlock()
	return transitions, err
}

func (reader *StreamReader) GetRawStreamFromState(state model.State, buffer int) (chan model.Transition, error) {
	stream, err := reader.client.GetStream(model.StreamState{
		StreamID: reader.config.StreamID,
		State:    state,
	}, buffer)
	return stream, err
}

func (reader *StreamReader) GetBatchedStream(buffer int, countPerRequest int) (chan model.Transition, error) {
	result := make(chan model.Transition, buffer)
	go func() {
		for {
			messages, err := reader.FetchNextTransitions(countPerRequest)
			if err != nil {
				result <- model.Transition{}
				close(result)
				log.Println("Error while receiving stream, %v", err)
				break

			}
			for _, msg := range messages {
				result <- msg
			}
		}
	}()
	return result, nil
}
