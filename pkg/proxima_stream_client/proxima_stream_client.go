package proxima_stream_client

import (
	"context"
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/proxima-one/streamdb-client-go/pkg/model"
	"github.com/proxima-one/streamdb-client-go/pkg/stream_registy"
	"time"
)

type ProximaStreamClient struct {
	registry              stream_registy.StreamRegistry
	clientsByUri          map[string]*streamDbConsumerClient
	offsetToEndpointCache *cache.Cache
}

func NewProximaStreamClient(options Options) *ProximaStreamClient {
	return &ProximaStreamClient{
		registry:              options.Registry,
		clientsByUri:          make(map[string]*streamDbConsumerClient),
		offsetToEndpointCache: cache.New(5*time.Minute, 10*time.Minute),
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
	events, err := streamClient.getEvents(stream, offset, count, direction)
	if err != nil {
		return nil, err
	}
	client.offsetToEndpointCache.Set(stream+events[len(events)-1].Offset.ToString(), endpoint, 0) // with default expiration
	return events, nil
}

func (client *ProximaStreamClient) StreamEvents(
	ctx context.Context,
	streamId string,
	offset *model.Offset,
	bufferSize int) <-chan model.StreamEvent {

	stream := make(chan model.StreamEvent, bufferSize)
	lastOffset := offset
	go func() {
		for { // infinite retry
			endpoint, err := client.findEndpoint(streamId, lastOffset)
			if err != nil {
				client.offsetToEndpointCache.Delete(streamId + lastOffset.ToString()) // get new endpoint from registry next time
				continue
			}
			streamClient, err := client.getStreamConsumerClient(endpoint.Uri)
			if err != nil {
				delete(client.clientsByUri, endpoint.Uri) // recreate client next time
				continue
			}
			lastOffset, err = streamClient.streamEvents(ctx, streamId, lastOffset, stream)
			if err != nil {
				continue
			}
		}
	}()
	return stream
}

func (client *ProximaStreamClient) findEndpoint(stream string, offset *model.Offset) (res *model.StreamEndpoint, err error) {
	if cachedEndpoint, ok := client.offsetToEndpointCache.Get(stream + offset.ToString()); ok {
		return cachedEndpoint.(*model.StreamEndpoint), nil
	}

	endpoints, err := client.registry.GetStreamEndpoints(stream, offset)
	if err != nil {
		return nil, err
	}
	if len(endpoints) == 0 {
		return nil, fmt.Errorf("no endpoints for %s - %s", stream, offset.ToString())
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
	client.offsetToEndpointCache.Set(stream+offset.ToString(), res, 0) // with default expiration
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
