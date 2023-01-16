package proximaclient

import (
	pbModel "github.com/proxima-one/streamdb-client-go/v2/api/proto/gen/proto/go/model/v1"
	pbConsumerModel "github.com/proxima-one/streamdb-client-go/v2/api/proto/gen/proto/go/stream_consumer/v1alpha1"
)

func protoStateTransitionToStreamEvent(transition *pbModel.StateTransition) StreamEvent {
	return StreamEvent{
		Payload:    transition.Event.Payload,
		Undo:       transition.Event.Undo,
		Offset:     protoOffsetToModel(transition.To),
		PrevOffset: protoOffsetToModel(transition.From),
		Timestamp:  protoTimestampToModel(transition.Event.Timestamp),
	}
}

func protoTimestampToModel(timestamp *pbModel.Timestamp) Timestamp {
	return *NewTimestamp(timestamp.EpochMs, timestamp.Parts)
}

func protoOffsetToModel(offset *pbModel.Offset) Offset {
	return Offset{
		OffsetId:  offset.Id,
		Height:    offset.Height,
		Timestamp: protoTimestampToModel(offset.Timestamp),
	}
}

func modelOffsetToProto(offset *Offset) *pbModel.Offset {
	return &pbModel.Offset{
		Id:        offset.OffsetId,
		Height:    offset.Height,
		Timestamp: modelTimestampToProto(&offset.Timestamp),
	}
}

func modelTimestampToProto(timestamp *Timestamp) *pbModel.Timestamp {
	return &pbModel.Timestamp{
		EpochMs: timestamp.EpochMs,
		Parts:   timestamp.Parts,
	}
}

func modelDirectionToProto(direction Direction) pbConsumerModel.Direction {
	switch direction {
	case DirectionLast:
		return pbConsumerModel.Direction_LAST
	case DirectionNext:
		return pbConsumerModel.Direction_NEXT
	default:
		panic("modelDirectionToProto: unknown direction")
	}
}
