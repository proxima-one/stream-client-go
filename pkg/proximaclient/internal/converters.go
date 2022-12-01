package internal

import (
	pbModel "github.com/proxima-one/streamdb-client-go/api/proto/gen/proto/go/model/v1"
	pbConsumerModel "github.com/proxima-one/streamdb-client-go/api/proto/gen/proto/go/stream_consumer/v1alpha1"
	"github.com/proxima-one/streamdb-client-go/pkg/proximaclient"
)

func ProtoStateTransitionToStreamEvent(transition *pbModel.StateTransition) proximaclient.StreamEvent {
	return proximaclient.StreamEvent{
		Payload:    transition.Event.Payload,
		Undo:       transition.Event.Undo,
		Offset:     ProtoOffsetToModel(transition.To),
		PrevOffset: ProtoOffsetToModel(transition.From),
		Timestamp:  ProtoTimestampToModel(transition.Event.Timestamp),
	}
}

func ProtoTimestampToModel(timestamp *pbModel.Timestamp) proximaclient.Timestamp {
	return *proximaclient.NewTimestamp(timestamp.EpochMs, timestamp.Parts)
}

func ProtoOffsetToModel(offset *pbModel.Offset) proximaclient.Offset {
	return proximaclient.Offset{
		OffsetId:  offset.Id,
		Height:    offset.Height,
		Timestamp: ProtoTimestampToModel(offset.Timestamp),
	}
}

func ModelOffsetToProto(offset *proximaclient.Offset) *pbModel.Offset {
	return &pbModel.Offset{
		Id:        offset.OffsetId,
		Height:    offset.Height,
		Timestamp: ModelTimestampToProto(&offset.Timestamp),
	}
}

func ModelTimestampToProto(timestamp *proximaclient.Timestamp) *pbModel.Timestamp {
	return &pbModel.Timestamp{
		EpochMs: timestamp.EpochMs,
		Parts:   timestamp.Parts,
	}
}

func ModelDirectionToProto(direction proximaclient.Direction) pbConsumerModel.Direction {
	switch direction {
	case proximaclient.DirectionLast:
		return pbConsumerModel.Direction_LAST
	case proximaclient.DirectionNext:
		return pbConsumerModel.Direction_NEXT
	default:
		panic("modelDirectionToProto: unknown direction")
	}
}
