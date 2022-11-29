package internal

import (
	modelv1 "github.com/proxima-one/streamdb-client-go/api/proto/gen/proto/go/model/v1"
	stream_consumerv1alpha1 "github.com/proxima-one/streamdb-client-go/api/proto/gen/proto/go/stream_consumer/v1alpha1"
	"github.com/proxima-one/streamdb-client-go/pkg/stream_model"
)

func ProtoStateTransitionToStreamEvent(transition *modelv1.StateTransition) stream_model.StreamEvent {
	return stream_model.StreamEvent{
		Payload:    transition.Event.Payload,
		Undo:       transition.Event.Undo,
		Offset:     ProtoOffsetToModel(transition.To),
		PrevOffset: ProtoOffsetToModel(transition.From),
		Timestamp:  ProtoTimestampToModel(transition.Event.Timestamp),
	}
}

func ProtoTimestampToModel(timestamp *modelv1.Timestamp) stream_model.Timestamp {
	return *stream_model.NewTimestamp(timestamp.EpochMs, timestamp.Parts)
}

func ProtoOffsetToModel(offset *modelv1.Offset) stream_model.Offset {
	return stream_model.Offset{
		OffsetId:  offset.Id,
		Height:    offset.Height,
		Timestamp: ProtoTimestampToModel(offset.Timestamp),
	}
}

func ModelOffsetToProto(offset *stream_model.Offset) *modelv1.Offset {
	return &modelv1.Offset{
		Id:        offset.OffsetId,
		Height:    offset.Height,
		Timestamp: ModelTimestampToProto(&offset.Timestamp),
	}
}

func ModelTimestampToProto(timestamp *stream_model.Timestamp) *modelv1.Timestamp {
	return &modelv1.Timestamp{
		EpochMs: timestamp.EpochMs,
		Parts:   timestamp.Parts,
	}
}

func ModelDirectionToProto(direction stream_model.Direction) stream_consumerv1alpha1.Direction {
	switch direction {
	case stream_model.DirectionLast:
		return stream_consumerv1alpha1.Direction_LAST
	case stream_model.DirectionNext:
		return stream_consumerv1alpha1.Direction_NEXT
	default:
		panic("modelDirectionToProto: unknown direction")
	}
}
