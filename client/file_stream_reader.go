package client

import (
	"context"
	"encoding/json"
	"github.com/proxima-one/streamdb-client-go/config"
	"github.com/proxima-one/streamdb-client-go/model"
	"io/ioutil"
	"time"
	//"google.golang.org/grpc/credentials" todo: add ssl and auth
)

// FileStreamReader copies StreamReader interface but reads messages from file with name of config.StreamID
// Example can be found in ../file_stream_reader_example folder
type FileStreamReader struct {
	config                 config.Config
	lastState              model.State
	preprocessFunc         TransitionPreprocessingFunc
	processedCount         int
	streamConnectionOption StreamConnectionOption
	allTransitions         []*model.Transition
	timeout                int // timeout between reads (for testing purposes)
}

func NewFileStreamReader(config config.Config, preprocess TransitionPreprocessingFunc) *FileStreamReader {
	return &FileStreamReader{
		config:         config,
		lastState:      config.GetState(),
		preprocessFunc: preprocess,
	}
}

func (reader *FileStreamReader) GetBufferLoadPercent() int { return 0 }

func (reader *FileStreamReader) Metrics() (int, int) {
	return reader.processedCount, reader.GetBufferLoadPercent()
}

func (reader *FileStreamReader) IsConnected() bool { return true }

func (reader *FileStreamReader) GetStreamID() string {
	return reader.config.GetStreamID()
}

func (reader *FileStreamReader) GetStreamChan() (<-chan *ProximaStreamObject, <-chan error) {
	return make(<-chan *ProximaStreamObject), make(<-chan error)
}

type JsonTransition struct {
	NewState model.State `json:"state"`
	Event    JsonEvent   `json:"event"`
}

type JsonEvent struct {
	Payload   json.RawMessage `json:"payload"`
	Timestamp time.Time       `json:"timestamp"`
	Undo      bool            `json:"undo"`
}

func (reader *FileStreamReader) Start(ctx context.Context, option *StreamConnectionOption) error {
	content, err := ioutil.ReadFile(reader.config.StreamID)
	if err != nil {
		panic(err.Error())
	}

	text := string(content)

	m := new(struct {
		Timeout int              `json:"timeout"`
		Data    []JsonTransition `json:"data"`
	})

	err = json.Unmarshal([]byte(text), &m)
	if err != nil {
		panic(err.Error())
	}

	transitions := m.Data
	reader.timeout = m.Timeout
	for _, transition := range transitions {
		raw, _ := transition.Event.Payload.MarshalJSON()
		trans := &model.Transition{
			NewState: transition.NewState,
			Event: model.Event{
				Payload:   &raw,
				Timestamp: transition.Event.Timestamp,
				Undo:      transition.Event.Undo,
			},
		}
		reader.allTransitions = append(reader.allTransitions, trans)
	}

	return nil
}

func (reader *FileStreamReader) Reconnect() error { return nil }

func (reader *FileStreamReader) Connect() error { return nil }

func (reader *FileStreamReader) Close() error { return nil }

func (reader *FileStreamReader) Restart() {}

func (reader *FileStreamReader) Stop() {}

func (reader *FileStreamReader) transitionToProximaStreamObject(transition *model.Transition) *ProximaStreamObject {
	return &ProximaStreamObject{
		Transition: transition,
		Preprocess: NewTransitionPreprocessingResult(transition, reader.preprocessFunc),
	}
}

func (reader *FileStreamReader) ReadNext() (*ProximaStreamObject, error) {
	time.Sleep(time.Duration(reader.timeout) * time.Millisecond)
	if reader.lastState.Id == "" {
		reader.lastState = reader.allTransitions[0].NewState
		return reader.transitionToProximaStreamObject(reader.allTransitions[0]), nil
	}
	for i, transition := range reader.allTransitions {
		if transition.NewState.Id == reader.lastState.Id && i+1 < len(reader.allTransitions) {
			reader.lastState = reader.allTransitions[i+1].NewState
			return reader.transitionToProximaStreamObject(reader.allTransitions[i+1]), nil
		}
	}
	println("FileReader: All messages done!")
	time.Sleep(10 * time.Second)
	return nil, nil
}
