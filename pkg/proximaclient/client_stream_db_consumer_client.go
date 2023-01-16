package proximaclient

import (
	"context"
	"crypto/tls"
	"fmt"
	streamConsumer "github.com/proxima-one/streamdb-client-go/v2/api/proto/gen/proto/go/stream_consumer/v1alpha1"
	"github.com/proxima-one/streamdb-client-go/v2/internal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"strings"
	"time"
)

type streamDbConsumerClient struct {
	client streamConsumer.StreamConsumerServiceClient
}

func newStreamDbConsumerClient(uri string) (*streamDbConsumerClient, error) {
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

	return &streamDbConsumerClient{client: streamConsumer.NewStreamConsumerServiceClient(conn)}, nil
}

func (c *streamDbConsumerClient) GetEvents(
	stream string,
	offset *Offset,
	count int,
	direction Direction) ([]StreamEvent, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // todo timeout?
	defer cancel()
	resp, err := c.client.GetStateTransitions(ctx, &streamConsumer.GetStateTransitionsRequest{
		StreamId:  stream,
		Offset:    modelOffsetToProto(offset),
		Count:     int32(count),
		Direction: modelDirectionToProto(direction),
	})
	if err != nil {
		return nil, fmt.Errorf("streamDbConsumerClient.getEvents: %s", err.Error())
	}
	return internal.MapArray(resp.StateTransitions, protoStateTransitionToStreamEvent), nil
}

func (c *streamDbConsumerClient) StreamEvents(
	ctx context.Context,
	streamId string,
	offset *Offset,
	eventsStream chan<- StreamEvent) (*Offset, error) {

	stream, err := c.client.StreamStateTransitions(ctx, &streamConsumer.StreamStateTransitionsRequest{
		StreamId: streamId,
		Offset:   modelOffsetToProto(offset),
	})
	if err != nil {
		return offset, err
	}
	lastOffset := *offset
	for ctx.Err() == nil {
		resp, err := stream.Recv()
		if err != nil {
			if ctx.Err() != nil { // ignore context cancel error as we're in an infinite loop
				return &lastOffset, nil
			}
			return &lastOffset, err
		}
		for _, stateTransition := range resp.StateTransition {
			event := protoStateTransitionToStreamEvent(stateTransition)
			lastOffset = event.Offset
			eventsStream <- event
		}
	}
	return &lastOffset, nil
}
