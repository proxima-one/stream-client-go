package proxima_stream_client

import (
	"github.com/proxima-one/streamdb-client-go/pkg/model"
	modelv1 "github.com/proxima-one/streamdb-client-go/proto/gen/proto/go/model/v1"
	stream_consumerv1alpha1 "github.com/proxima-one/streamdb-client-go/proto/gen/proto/go/stream_consumer/v1alpha1"
)

func protoStateTransitionToStreamEvent(transition *modelv1.StateTransition) model.StreamEvent {
	return model.StreamEvent{
		Payload:    transition.Event.Payload,
		Undo:       transition.Event.Undo,
		Offset:     protoOffsetToModel(transition.To),
		PrevOffset: protoOffsetToModel(transition.From),
		Timestamp:  protoTimestampToModel(transition.Event.Timestamp),
	}
}

func protoTimestampToModel(timestamp *modelv1.Timestamp) model.Timestamp {
	return *model.NewTimestamp(timestamp.EpochMs, timestamp.Parts)
}

func protoOffsetToModel(offset *modelv1.Offset) model.Offset {
	return model.Offset{
		OffsetId:  offset.Id,
		Height:    offset.Height,
		Timestamp: protoTimestampToModel(offset.Timestamp),
	}
}

func modelOffsetToProto(offset *model.Offset) *modelv1.Offset {
	return &modelv1.Offset{
		Id:        offset.OffsetId,
		Height:    offset.Height,
		Timestamp: modelTimestampToProto(&offset.Timestamp),
	}
}

func modelTimestampToProto(timestamp *model.Timestamp) *modelv1.Timestamp {
	return &modelv1.Timestamp{
		EpochMs: timestamp.EpochMs,
		Parts:   timestamp.Parts,
	}
}

func modelDirectionToProto(direction Direction) stream_consumerv1alpha1.Direction {
	switch direction {
	case DirectionLast:
		return stream_consumerv1alpha1.Direction_LAST
	case DirectionNext:
		return stream_consumerv1alpha1.Direction_NEXT
	default:
		panic("modelDirectionToProto: unknown direction")
	}
}
