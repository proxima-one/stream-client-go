package client

import (
	"context"
	"crypto/tls"
	"github.com/proxima-one/streamdb-client-go/config"
	pb "github.com/proxima-one/streamdb-client-go/gen/proto"
	"github.com/proxima-one/streamdb-client-go/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"time"
)

const defaultChannelSize = 10000
const defaultRpcChannelSize = 1000
const defaultRpcSleepTime = 200 * time.Millisecond

type ProximaClient struct {
	config config.Config
	conn   *grpc.ClientConn
	grpc   pb.MessagesServiceClient
}

func NewProximaClient(config config.Config) *ProximaClient {
	if config.User != "" || config.Password != "" || config.Token != "" {
		log.Println("[WARN] Authorization is not supported yet")
		log.Println("[WARN] User is not used in the client yet")
		log.Println("[WARN] Password is not used in the client yet")
		log.Println("[WARN] Token is not used in the client yet")
	}

	return &ProximaClient{
		config: config,
	}
}

func (client *ProximaClient) Connect() error {
	address := client.config.GetFullAddress()
	dialOption := grpc.WithInsecure()
	if client.config.GetPort() == 443 {
		dialOption = grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true}))
	}
	conn, err := grpc.Dial(address, dialOption)
	if err != nil {
		log.Fatalf("Error while making connection, %v", err)
		return err
	}
	client.conn = conn
	client.grpc = pb.NewMessagesServiceClient(conn)
	return nil
}

func (client *ProximaClient) Close() error {
	return client.conn.Close()
}

func streamMessageToModel(msg *pb.StreamMessage) *model.Transition {
	return &model.Transition{
		NewState: model.NewState(msg.Id),
		Event: model.Event{
			Undo:      msg.GetHeader().GetUndo(),
			Payload:   &msg.Payload,
			Timestamp: msg.GetTimestamp().AsTime(),
		},
	}
}

func (client *ProximaClient) GetTransitionsAfter(ctx context.Context,
	streamState model.StreamState,
	count int) ([]*model.Transition, error) {

	if ctx == nil {
		ctx = context.Background()
	}
	res, err := client.grpc.GetNextMessages(ctx, &pb.GetNextMessagesRequest{
		StreamId:      streamState.StreamID,
		LastMessageId: streamState.State.Id,
		Count:         int32(count),
	})
	if err != nil {
		return nil, err
	}
	transitions := make([]*model.Transition, len(res.Messages))
	for i, msg := range res.Messages {
		transitions[i] = streamMessageToModel(msg)
	}
	return transitions, nil
}

func (client *ProximaClient) GetStream(ctx context.Context,
	streamState model.StreamState) (<-chan *model.Transition, <-chan error, error) {

	if ctx == nil {
		ctx = context.Background()
	}
	streamClient, err := client.grpc.StreamMessages(ctx, &pb.StreamMessagesRequest{
		StreamId:      streamState.StreamID,
		LastMessageId: streamState.State.Id,
	})
	if err != nil {
		return nil, nil, err
	}
	result := make(chan *model.Transition, defaultChannelSize)
	errc := make(chan error, 1)
	go func() {
		defer close(result)
		defer close(errc)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				messages, err := streamClient.Recv()
				if err != nil {
					errc <- err
					return
				}
				for _, msg := range messages.Messages {
					result <- streamMessageToModel(msg)
				}
			}
		}
	}()

	return result, errc, nil
}

func (client *ProximaClient) GetStreamBasedOnRpc(ctx context.Context,
	streamState model.StreamState) (<-chan *model.Transition, <-chan error, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	result := make(chan *model.Transition, defaultChannelSize)
	errc := make(chan error, 1)
	go func() {
		defer close(result)
		defer close(errc)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				messages, err := client.GetTransitionsAfter(ctx, streamState, defaultRpcChannelSize)
				if err != nil {
					errc <- err
					return
				}
				if len(messages) == 0 {
					time.Sleep(defaultRpcSleepTime)
					continue
				}
				for _, msg := range messages {
					result <- msg
				}
				streamState.State = messages[len(messages)-1].NewState
			}
		}
	}()
	return result, errc, nil
}
