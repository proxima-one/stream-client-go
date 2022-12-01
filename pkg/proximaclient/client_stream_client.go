package proximaclient

import (
	"context"
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

type StreamClient struct {
	registry              StreamRegistry
	clientsByUri          map[string]*streamDbConsumerClient
	endpointByOffsetCache *cache.Cache
}

func NewStreamClient(options Options) *StreamClient {
	return &StreamClient{
		registry:              options.Registry,
		clientsByUri:          make(map[string]*streamDbConsumerClient),
		endpointByOffsetCache: cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (client *StreamClient) FetchEvents(stream string, offset *Offset, count int, direction Direction) ([]StreamEvent, error) {
	endpoint, err := client.findEndpoint(stream, offset)
	if err != nil {
		return nil, err
	}
	streamClient, err := client.getStreamConsumerClient(endpoint.Uri)
	if err != nil {
		return nil, err
	}
	events, err := streamClient.GetEvents(stream, offset, count, direction)
	if err != nil {
		return nil, err
	}
	client.endpointByOffsetCache.Set(stream+events[len(events)-1].Offset.String(), endpoint, 0) // with default expiration
	return events, nil
}

func (client *StreamClient) StreamEvents(
	ctx context.Context,
	streamId string,
	offset *Offset,
	bufferSize int) <-chan StreamEvent {

	stream := make(chan StreamEvent, bufferSize)
	lastOffset := offset
	go func() {
		for { // infinite retry
			endpoint, err := client.findEndpoint(streamId, lastOffset)
			if err != nil {
				client.endpointByOffsetCache.Delete(streamId + lastOffset.String()) // get new endpoint from registry next time
				continue
			}
			streamClient, err := client.getStreamConsumerClient(endpoint.Uri)
			if err != nil {
				delete(client.clientsByUri, endpoint.Uri) // recreate client next time
				continue
			}
			lastOffset, err = streamClient.StreamEvents(ctx, streamId, lastOffset, stream)
			if err != nil {
				continue
			}

			// streamEvents exited with nil error
			close(stream)
			break
		}
	}()
	return stream
}

func (client *StreamClient) findEndpoint(stream string, offset *Offset) (res *StreamEndpoint, err error) {
	if cachedEndpoint, ok := client.endpointByOffsetCache.Get(stream + offset.String()); ok {
		return cachedEndpoint.(*StreamEndpoint), nil
	}

	endpoints, err := client.registry.GetStreamEndpoints(stream, offset)
	if err != nil {
		return nil, err
	}
	if len(endpoints) == 0 {
		return nil, fmt.Errorf("no endpoints for %s - %s", stream, offset.String())
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
	client.endpointByOffsetCache.Set(stream+offset.String(), res, 0) // with default expiration
	return
}

func (client *StreamClient) getStreamConsumerClient(uri string) (*streamDbConsumerClient, error) {
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
