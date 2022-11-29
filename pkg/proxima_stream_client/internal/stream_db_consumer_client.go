package internal

import (
	"context"
	"crypto/tls"
	"fmt"
	streamConsumer "github.com/proxima-one/streamdb-client-go/api/proto/gen/proto/go/stream_consumer/v1alpha1"
	"github.com/proxima-one/streamdb-client-go/internal"
	"github.com/proxima-one/streamdb-client-go/pkg/stream_model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"strings"
	"time"
)

type StreamDbConsumerClient struct {
	client streamConsumer.StreamConsumerServiceClient
}

func NewStreamDbConsumerClient(uri string) (*StreamDbConsumerClient, error) {
	isSecure := strings.Contains(uri, ":443")
	dialOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	if isSecure {
		dialOption = grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true}))
	}
	conn, err := grpc.Dial(
		uri,
		dialOption,
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                100 * time.Second,
			Timeout:             100 * time.Second,
			PermitWithoutStream: true,
		}),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(100*1024*1024)),
	)
	if err != nil {
		return nil, err
	}

	return &StreamDbConsumerClient{client: streamConsumer.NewStreamConsumerServiceClient(conn)}, nil
}

func (c *StreamDbConsumerClient) GetEvents(
	stream string,
	offset *stream_model.Offset,
	count int,
	direction stream_model.Direction) ([]stream_model.StreamEvent, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // todo timeout?
	defer cancel()
	resp, err := c.client.GetStateTransitions(ctx, &streamConsumer.GetStateTransitionsRequest{
		StreamId:  stream,
		Offset:    ModelOffsetToProto(offset),
		Count:     int32(count),
		Direction: ModelDirectionToProto(direction),
	})
	if err != nil {
		return nil, fmt.Errorf("StreamDbConsumerClient.getEvents: %s", err.Error())
	}
	return internal.MapArray(resp.StateTransitions, ProtoStateTransitionToStreamEvent), nil
}

func (c *StreamDbConsumerClient) StreamEvents(
	ctx context.Context,
	streamId string,
	offset *stream_model.Offset,
	eventsStream chan<- stream_model.StreamEvent) (*stream_model.Offset, error) {

	stream, err := c.client.StreamStateTransitions(ctx, &streamConsumer.StreamStateTransitionsRequest{
		StreamId: streamId,
		Offset:   ModelOffsetToProto(offset),
	})
	if err != nil {
		return offset, err
	}
	var lastOffset *stream_model.Offset
	for ctx.Err() == nil {
		resp, err := stream.Recv()
		if err != nil {
			if ctx.Err() != nil { // ignore context cancel error as we're in an infinite loop
				return lastOffset, nil
			}
			return lastOffset, err
		}
		for _, stateTransition := range resp.StateTransition {
			event := ProtoStateTransitionToStreamEvent(stateTransition)
			lastOffset = &event.Offset
			eventsStream <- event
		}
	}
	return lastOffset, nil
}
