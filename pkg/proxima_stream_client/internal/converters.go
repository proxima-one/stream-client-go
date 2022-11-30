package internal

import (
	pbModel "github.com/proxima-one/streamdb-client-go/api/proto/gen/proto/go/model/v1"
	pbConsumerModel "github.com/proxima-one/streamdb-client-go/api/proto/gen/proto/go/stream_consumer/v1alpha1"
	"github.com/proxima-one/streamdb-client-go/pkg/stream_model"
)

func ProtoStateTransitionToStreamEvent(transition *pbModel.StateTransition) stream_model.StreamEvent {
	return stream_model.StreamEvent{
		Payload:    transition.Event.Payload,
		Undo:       transition.Event.Undo,
		Offset:     ProtoOffsetToModel(transition.To),
		PrevOffset: ProtoOffsetToModel(transition.From),
		Timestamp:  ProtoTimestampToModel(transition.Event.Timestamp),
	}
}

func ProtoTimestampToModel(timestamp *pbModel.Timestamp) stream_model.Timestamp {
	return *stream_model.NewTimestamp(timestamp.EpochMs, timestamp.Parts)
}

func ProtoOffsetToModel(offset *pbModel.Offset) stream_model.Offset {
	return stream_model.Offset{
		OffsetId:  offset.Id,
		Height:    offset.Height,
		Timestamp: ProtoTimestampToModel(offset.Timestamp),
	}
}

func ModelOffsetToProto(offset *stream_model.Offset) *pbModel.Offset {
	return &pbModel.Offset{
		Id:        offset.OffsetId,
		Height:    offset.Height,
		Timestamp: ModelTimestampToProto(&offset.Timestamp),
	}
}

func ModelTimestampToProto(timestamp *stream_model.Timestamp) *pbModel.Timestamp {
	return &pbModel.Timestamp{
		EpochMs: timestamp.EpochMs,
		Parts:   timestamp.Parts,
	}
}

func ModelDirectionToProto(direction stream_model.Direction) pbConsumerModel.Direction {
	switch direction {
	case stream_model.DirectionLast:
		return pbConsumerModel.Direction_LAST
	case stream_model.DirectionNext:
		return pbConsumerModel.Direction_NEXT
	default:
		panic("modelDirectionToProto: unknown direction")
	}
}
