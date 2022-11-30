package proxima_stream_client

import (
	"context"
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/proxima-one/streamdb-client-go/pkg/proxima_stream_client/internal"
	"github.com/proxima-one/streamdb-client-go/pkg/stream_model"
	"github.com/proxima-one/streamdb-client-go/pkg/stream_registy"
	"time"
)

type ProximaStreamClient struct {
	registry              stream_registy.StreamRegistry
	clientsByUri          map[string]*internal.StreamDbConsumerClient
	endpointByOffsetCache *cache.Cache
}

func NewProximaStreamClient(options Options) *ProximaStreamClient {
	return &ProximaStreamClient{
		registry:              options.Registry,
		clientsByUri:          make(map[string]*internal.StreamDbConsumerClient),
		endpointByOffsetCache: cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (client *ProximaStreamClient) FetchEvents(stream string, offset *stream_model.Offset, count int, direction stream_model.Direction) ([]stream_model.StreamEvent, error) {
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

func (client *ProximaStreamClient) StreamEvents(
	ctx context.Context,
	streamId string,
	offset *stream_model.Offset,
	bufferSize int) <-chan stream_model.StreamEvent {

	stream := make(chan stream_model.StreamEvent, bufferSize)
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

func (client *ProximaStreamClient) findEndpoint(stream string, offset *stream_model.Offset) (res *stream_model.StreamEndpoint, err error) {
	if cachedEndpoint, ok := client.endpointByOffsetCache.Get(stream + offset.String()); ok {
		return cachedEndpoint.(*stream_model.StreamEndpoint), nil
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

func (client *ProximaStreamClient) getStreamConsumerClient(uri string) (*internal.StreamDbConsumerClient, error) {
	if client, ok := client.clientsByUri[uri]; ok {
		return client, nil
	}
	res, err := internal.NewStreamDbConsumerClient(uri)
	if err != nil {
		return nil, err
	}
	client.clientsByUri[uri] = res
	return res, nil
}
