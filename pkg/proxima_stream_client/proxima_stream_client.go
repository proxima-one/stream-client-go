package proxima_stream_client

import (
	"context"
	"github.com/proxima-one/streamdb-client-go/pkg/model"
	"github.com/proxima-one/streamdb-client-go/pkg/stream_registy"
)

type ProximaStreamClient struct {
	registry     stream_registy.StreamRegistry
	clientsByUri map[string]*streamDbConsumerClient
}

func NewProximaStreamClient(options Options) *ProximaStreamClient {
	return &ProximaStreamClient{
		registry: options.Registry,
	}
}

func (client *ProximaStreamClient) FetchEvents(stream string, offset *model.Offset, count int, direction Direction) ([]model.StreamEvent, error) {
	endpoint, err := client.findEndpoint(stream, offset)
	if err != nil {
		return nil, err
	}
	streamClient, err := client.getStreamConsumerClient(endpoint.Uri)
	if err != nil {
		return nil, err
	}
	return streamClient.getEvents(stream, offset, count, direction)
}

func (client *ProximaStreamClient) StreamEvents(
	ctx context.Context,
	streamId string,
	offset *model.Offset,
	bufferSize int) <-chan model.StreamEvent {
	// todo: handle errors
	stream := make(chan model.StreamEvent, bufferSize)

	endpoint, _ := client.findEndpoint(streamId, offset)
	streamClient, _ := client.getStreamConsumerClient(endpoint.Uri)
	go streamClient.streamEvents(ctx, streamId, offset, stream)
	return stream
}

func (client *ProximaStreamClient) findEndpoint(stream string, offset *model.Offset) (res *model.StreamEndpoint, err error) {
	endpoints, err := client.registry.GetStreamEndpoints(stream, offset)
	if err != nil {
		return nil, err
	}
	if len(endpoints) == 0 {
		return
	}
	res = &endpoints[0]
	for _, endpoint := range endpoints {
		if endpoint.Stats.EndOffset == nil {
			res = &endpoint
			break
		}
		if res.Stats.EndOffset.Height < endpoint.Stats.EndOffset.Height {
			res = &endpoint
		}
	}
	return
}

func (client *ProximaStreamClient) getStreamConsumerClient(uri string) (*streamDbConsumerClient, error) {
	if client, ok := client.clientsByUri[uri]; ok {
		return client, nil
	}
	res, err := newStreamDbConsumerClient(uri)
	if err != nil {
		return nil, err
	}
	client.clientsByUri[uri] = res
	return res, nil
}
