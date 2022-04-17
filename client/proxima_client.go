package client

import (
	"context"
	"crypto/tls"
	"github.com/proxima-one/pocs/stream-db-endpoint/config"
	pb "github.com/proxima-one/pocs/stream-db-endpoint/gen/proto"
	"github.com/proxima-one/pocs/stream-db-endpoint/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

type ProximaClient struct {
	config *config.Config
	conn   *grpc.ClientConn
	grpc   pb.MessagesServiceClient
}

func NewProximaClient(config *config.Config) *ProximaClient {
	return &ProximaClient{
		config: config,
	}
}

func (client *ProximaClient) Connect() error {
	address := client.config.GetFullAddress()
	dialOption := grpc.WithInsecure()
	if client.config.Port == 443 {
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

func (client *ProximaClient) Disconnect() {
	client.conn.Close()
}

func streamMessageToModel(msg *pb.StreamMessage) *model.Transition {
	return &model.Transition{
		NewState: model.NewState(msg.Id),
		Event: model.Event{
			Undo: msg.GetHeader().GetUndo(),
			//Payload:   msg.GetPayload(),
			Payload:   &msg.Payload, //copy payload or not?
			Timestamp: msg.GetTimestamp().AsTime(),
		},
	}
}

func (client *ProximaClient) GetTransitionsAfter(streamState model.StreamState, count int) ([]model.Transition, error) {
	res, err := client.grpc.GetNextMessages(context.Background(), &pb.GetNextMessagesRequest{
		StreamId:      streamState.StreamID,
		LastMessageId: streamState.State.Id,
		Count:         int32(count),
	})
	if err != nil {
		return nil, err
	}
	transitions := make([]model.Transition, len(res.Messages))
	for i, msg := range res.Messages {
		transitions[i] = *streamMessageToModel(msg)
	}
	return transitions, nil
}

func (client *ProximaClient) GetStream(streamState model.StreamState, bufferSize int) (chan model.Transition, error) {
	streamClient, err := client.grpc.StreamMessages(context.Background(), &pb.StreamMessagesRequest{
		StreamId:      streamState.StreamID,
		LastMessageId: streamState.State.Id,
	})
	if err != nil {
		log.Fatalf("Error while getting stream, %v", err)
		return nil, err
	}
	result := make(chan model.Transition, bufferSize)
	go func() {
		for {
			messages, err := streamClient.Recv()
			if err != nil {
				log.Printf("Error while reading stream, %v\n", err)
				result <- model.Transition{}
				close(result)
				break
			}
			for _, msg := range messages.Messages {
				result <- *streamMessageToModel(msg)
			}
		}
	}()

	return result, nil
}
