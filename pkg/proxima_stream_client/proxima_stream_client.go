package proxima_stream_client

import (
	"github.com/proxima-one/streamdb-client-go/pkg/model"
	"github.com/proxima-one/streamdb-client-go/pkg/stream_registy"
)

type ProximaStreamClient struct {
	registry stream_registy.StreamRegistry
}

func NewProximaStreamClient(options Options) *ProximaStreamClient {
	return &ProximaStreamClient{
		registry: options.Registry,
	}
}

func (client *ProximaStreamClient) FetchEvents(stream string, offset model.Offset, count int, direction Direction) []model.StreamEvent {
	return nil
}

func (client *ProximaStreamClient) StreamEvents(stream string, offset model.Offset) <-chan model.StreamEvent {
	return nil
}
