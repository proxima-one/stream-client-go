// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: stream_consumer/v1alpha1/stream_consumer.proto

package stream_consumerv1alpha1

import (
	v1 "github.com/proxima-one/streamdb-client-go/api/proto/gen/proto/go/model/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Direction int32

const (
	Direction_NEXT Direction = 0
	Direction_LAST Direction = 1
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0: "NEXT",
		1: "LAST",
	}
	Direction_value = map[string]int32{
		"NEXT": 0,
		"LAST": 1,
	}
)

func (x Direction) Enum() *Direction {
	p := new(Direction)
	*p = x
	return p
}

func (x Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_stream_consumer_v1alpha1_stream_consumer_proto_enumTypes[0].Descriptor()
}

func (Direction) Type() protoreflect.EnumType {
	return &file_stream_consumer_v1alpha1_stream_consumer_proto_enumTypes[0]
}

func (x Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Direction.Descriptor instead.
func (Direction) EnumDescriptor() ([]byte, []int) {
	return file_stream_consumer_v1alpha1_stream_consumer_proto_rawDescGZIP(), []int{0}
}

type FindStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamId string `protobuf:"bytes,1,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
}

func (x *FindStreamRequest) Reset() {
	*x = FindStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindStreamRequest) ProtoMessage() {}

func (x *FindStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindStreamRequest.ProtoReflect.Descriptor instead.
func (*FindStreamRequest) Descriptor() ([]byte, []int) {
	return file_stream_consumer_v1alpha1_stream_consumer_proto_rawDescGZIP(), []int{0}
}

func (x *FindStreamRequest) GetStreamId() string {
	if x != nil {
		return x.StreamId
	}
	return ""
}

type FindStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stream *v1.Stream `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty"`
}

func (x *FindStreamResponse) Reset() {
	*x = FindStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindStreamResponse) ProtoMessage() {}

func (x *FindStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindStreamResponse.ProtoReflect.Descriptor instead.
func (*FindStreamResponse) Descriptor() ([]byte, []int) {
	return file_stream_consumer_v1alpha1_stream_consumer_proto_rawDescGZIP(), []int{1}
}

func (x *FindStreamResponse) GetStream() *v1.Stream {
	if x != nil {
		return x.Stream
	}
	return nil
}

type FindOffsetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamId string `protobuf:"bytes,1,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
	// Types that are assignable to OffsetFilter:
	//	*FindOffsetRequest_Height
	//	*FindOffsetRequest_Timestamp
	OffsetFilter isFindOffsetRequest_OffsetFilter `protobuf_oneof:"offset_filter"`
}

func (x *FindOffsetRequest) Reset() {
	*x = FindOffsetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOffsetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOffsetRequest) ProtoMessage() {}

func (x *FindOffsetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOffsetRequest.ProtoReflect.Descriptor instead.
func (*FindOffsetRequest) Descriptor() ([]byte, []int) {
	return file_stream_consumer_v1alpha1_stream_consumer_proto_rawDescGZIP(), []int{2}
}

func (x *FindOffsetRequest) GetStreamId() string {
	if x != nil {
		return x.StreamId
	}
	return ""
}

func (m *FindOffsetRequest) GetOffsetFilter() isFindOffsetRequest_OffsetFilter {
	if m != nil {
		return m.OffsetFilter
	}
	return nil
}

func (x *FindOffsetRequest) GetHeight() int64 {
	if x, ok := x.GetOffsetFilter().(*FindOffsetRequest_Height); ok {
		return x.Height
	}
	return 0
}

func (x *FindOffsetRequest) GetTimestamp() *v1.Timestamp {
	if x, ok := x.GetOffsetFilter().(*FindOffsetRequest_Timestamp); ok {
		return x.Timestamp
	}
	return nil
}

type isFindOffsetRequest_OffsetFilter interface {
	isFindOffsetRequest_OffsetFilter()
}

type FindOffsetRequest_Height struct {
	Height int64 `protobuf:"varint,2,opt,name=height,proto3,oneof"`
}

type FindOffsetRequest_Timestamp struct {
	Timestamp *v1.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3,oneof"`
}

func (*FindOffsetRequest_Height) isFindOffsetRequest_OffsetFilter() {}

func (*FindOffsetRequest_Timestamp) isFindOffsetRequest_OffsetFilter() {}

type FindOffsetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset *v1.Offset `protobuf:"bytes,1,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *FindOffsetResponse) Reset() {
	*x = FindOffsetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOffsetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOffsetResponse) ProtoMessage() {}

func (x *FindOffsetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOffsetResponse.ProtoReflect.Descriptor instead.
func (*FindOffsetResponse) Descriptor() ([]byte, []int) {
	return file_stream_consumer_v1alpha1_stream_consumer_proto_rawDescGZIP(), []int{3}
}

func (x *FindOffsetResponse) GetOffset() *v1.Offset {
	if x != nil {
		return x.Offset
	}
	return nil
}

type GetStateTransitionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamId  string     `protobuf:"bytes,1,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
	Offset    *v1.Offset `protobuf:"bytes,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Count     int32      `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Direction Direction  `protobuf:"varint,4,opt,name=direction,proto3,enum=stream_consumer.v1alpha1.Direction" json:"direction,omitempty"`
}

func (x *GetStateTransitionsRequest) Reset() {
	*x = GetStateTransitionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStateTransitionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStateTransitionsRequest) ProtoMessage() {}

func (x *GetStateTransitionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStateTransitionsRequest.ProtoReflect.Descriptor instead.
func (*GetStateTransitionsRequest) Descriptor() ([]byte, []int) {
	return file_stream_consumer_v1alpha1_stream_consumer_proto_rawDescGZIP(), []int{4}
}

func (x *GetStateTransitionsRequest) GetStreamId() string {
	if x != nil {
		return x.StreamId
	}
	return ""
}

func (x *GetStateTransitionsRequest) GetOffset() *v1.Offset {
	if x != nil {
		return x.Offset
	}
	return nil
}

func (x *GetStateTransitionsRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetStateTransitionsRequest) GetDirection() Direction {
	if x != nil {
		return x.Direction
	}
	return Direction_NEXT
}

type GetStateTransitionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StateTransitions []*v1.StateTransition `protobuf:"bytes,1,rep,name=state_transitions,json=stateTransitions,proto3" json:"state_transitions,omitempty"`
}

func (x *GetStateTransitionsResponse) Reset() {
	*x = GetStateTransitionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStateTransitionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStateTransitionsResponse) ProtoMessage() {}

func (x *GetStateTransitionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStateTransitionsResponse.ProtoReflect.Descriptor instead.
func (*GetStateTransitionsResponse) Descriptor() ([]byte, []int) {
	return file_stream_consumer_v1alpha1_stream_consumer_proto_rawDescGZIP(), []int{5}
}

func (x *GetStateTransitionsResponse) GetStateTransitions() []*v1.StateTransition {
	if x != nil {
		return x.StateTransitions
	}
	return nil
}

type StreamStateTransitionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamId string     `protobuf:"bytes,1,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
	Offset   *v1.Offset `protobuf:"bytes,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *StreamStateTransitionsRequest) Reset() {
	*x = StreamStateTransitionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamStateTransitionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamStateTransitionsRequest) ProtoMessage() {}

func (x *StreamStateTransitionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamStateTransitionsRequest.ProtoReflect.Descriptor instead.
func (*StreamStateTransitionsRequest) Descriptor() ([]byte, []int) {
	return file_stream_consumer_v1alpha1_stream_consumer_proto_rawDescGZIP(), []int{6}
}

func (x *StreamStateTransitionsRequest) GetStreamId() string {
	if x != nil {
		return x.StreamId
	}
	return ""
}

func (x *StreamStateTransitionsRequest) GetOffset() *v1.Offset {
	if x != nil {
		return x.Offset
	}
	return nil
}

type StreamStateTransitionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StateTransition []*v1.StateTransition `protobuf:"bytes,1,rep,name=state_transition,json=stateTransition,proto3" json:"state_transition,omitempty"`
}

func (x *StreamStateTransitionsResponse) Reset() {
	*x = StreamStateTransitionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamStateTransitionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamStateTransitionsResponse) ProtoMessage() {}

func (x *StreamStateTransitionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamStateTransitionsResponse.ProtoReflect.Descriptor instead.
func (*StreamStateTransitionsResponse) Descriptor() ([]byte, []int) {
	return file_stream_consumer_v1alpha1_stream_consumer_proto_rawDescGZIP(), []int{7}
}

func (x *StreamStateTransitionsResponse) GetStateTransition() []*v1.StateTransition {
	if x != nil {
		return x.StateTransition
	}
	return nil
}

var File_stream_consumer_v1alpha1_stream_consumer_proto protoreflect.FileDescriptor

var file_stream_consumer_v1alpha1_stream_consumer_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x14, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x30, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x49, 0x64, 0x22, 0x3e, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x22, 0x90, 0x01, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x33, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x0f, 0x0a, 0x0d, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x5f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x3e, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0xbc, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x12, 0x28, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x41, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x65, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x66, 0x0a, 0x1d, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x22, 0x66, 0x0a, 0x1e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x1f, 0x0a, 0x09, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x45, 0x58, 0x54,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x41, 0x53, 0x54, 0x10, 0x01, 0x32, 0x86, 0x04, 0x0a,
	0x15, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x0a, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x2b, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2c, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x69, 0x0a, 0x0a, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x2b, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x84, 0x01, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x8f, 0x01, 0x0a, 0x16, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x37,
	0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0xa3, 0x02, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x13, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x71, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x6d,
	0x61, 0x2d, 0x6f, 0x6e, 0x65, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x64, 0x62, 0x2d, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x17, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xca, 0x02, 0x17, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x23, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x18, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_stream_consumer_v1alpha1_stream_consumer_proto_rawDescOnce sync.Once
	file_stream_consumer_v1alpha1_stream_consumer_proto_rawDescData = file_stream_consumer_v1alpha1_stream_consumer_proto_rawDesc
)

func file_stream_consumer_v1alpha1_stream_consumer_proto_rawDescGZIP() []byte {
	file_stream_consumer_v1alpha1_stream_consumer_proto_rawDescOnce.Do(func() {
		file_stream_consumer_v1alpha1_stream_consumer_proto_rawDescData = protoimpl.X.CompressGZIP(file_stream_consumer_v1alpha1_stream_consumer_proto_rawDescData)
	})
	return file_stream_consumer_v1alpha1_stream_consumer_proto_rawDescData
}

var file_stream_consumer_v1alpha1_stream_consumer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_stream_consumer_v1alpha1_stream_consumer_proto_goTypes = []interface{}{
	(Direction)(0),                         // 0: stream_consumer.v1alpha1.Direction
	(*FindStreamRequest)(nil),              // 1: stream_consumer.v1alpha1.FindStreamRequest
	(*FindStreamResponse)(nil),             // 2: stream_consumer.v1alpha1.FindStreamResponse
	(*FindOffsetRequest)(nil),              // 3: stream_consumer.v1alpha1.FindOffsetRequest
	(*FindOffsetResponse)(nil),             // 4: stream_consumer.v1alpha1.FindOffsetResponse
	(*GetStateTransitionsRequest)(nil),     // 5: stream_consumer.v1alpha1.GetStateTransitionsRequest
	(*GetStateTransitionsResponse)(nil),    // 6: stream_consumer.v1alpha1.GetStateTransitionsResponse
	(*StreamStateTransitionsRequest)(nil),  // 7: stream_consumer.v1alpha1.StreamStateTransitionsRequest
	(*StreamStateTransitionsResponse)(nil), // 8: stream_consumer.v1alpha1.StreamStateTransitionsResponse
	(*v1.Stream)(nil),                      // 9: model.v1.Stream
	(*v1.Timestamp)(nil),                   // 10: model.v1.Timestamp
	(*v1.Offset)(nil),                      // 11: model.v1.Offset
	(*v1.StateTransition)(nil),             // 12: model.v1.StateTransition
}
var file_stream_consumer_v1alpha1_stream_consumer_proto_depIdxs = []int32{
	9,  // 0: stream_consumer.v1alpha1.FindStreamResponse.stream:type_name -> model.v1.Stream
	10, // 1: stream_consumer.v1alpha1.FindOffsetRequest.timestamp:type_name -> model.v1.Timestamp
	11, // 2: stream_consumer.v1alpha1.FindOffsetResponse.offset:type_name -> model.v1.Offset
	11, // 3: stream_consumer.v1alpha1.GetStateTransitionsRequest.offset:type_name -> model.v1.Offset
	0,  // 4: stream_consumer.v1alpha1.GetStateTransitionsRequest.direction:type_name -> stream_consumer.v1alpha1.Direction
	12, // 5: stream_consumer.v1alpha1.GetStateTransitionsResponse.state_transitions:type_name -> model.v1.StateTransition
	11, // 6: stream_consumer.v1alpha1.StreamStateTransitionsRequest.offset:type_name -> model.v1.Offset
	12, // 7: stream_consumer.v1alpha1.StreamStateTransitionsResponse.state_transition:type_name -> model.v1.StateTransition
	1,  // 8: stream_consumer.v1alpha1.StreamConsumerService.FindStream:input_type -> stream_consumer.v1alpha1.FindStreamRequest
	3,  // 9: stream_consumer.v1alpha1.StreamConsumerService.FindOffset:input_type -> stream_consumer.v1alpha1.FindOffsetRequest
	5,  // 10: stream_consumer.v1alpha1.StreamConsumerService.GetStateTransitions:input_type -> stream_consumer.v1alpha1.GetStateTransitionsRequest
	7,  // 11: stream_consumer.v1alpha1.StreamConsumerService.StreamStateTransitions:input_type -> stream_consumer.v1alpha1.StreamStateTransitionsRequest
	2,  // 12: stream_consumer.v1alpha1.StreamConsumerService.FindStream:output_type -> stream_consumer.v1alpha1.FindStreamResponse
	4,  // 13: stream_consumer.v1alpha1.StreamConsumerService.FindOffset:output_type -> stream_consumer.v1alpha1.FindOffsetResponse
	6,  // 14: stream_consumer.v1alpha1.StreamConsumerService.GetStateTransitions:output_type -> stream_consumer.v1alpha1.GetStateTransitionsResponse
	8,  // 15: stream_consumer.v1alpha1.StreamConsumerService.StreamStateTransitions:output_type -> stream_consumer.v1alpha1.StreamStateTransitionsResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_stream_consumer_v1alpha1_stream_consumer_proto_init() }
func file_stream_consumer_v1alpha1_stream_consumer_proto_init() {
	if File_stream_consumer_v1alpha1_stream_consumer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindStreamRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindStreamResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOffsetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOffsetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStateTransitionsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStateTransitionsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamStateTransitionsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamStateTransitionsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*FindOffsetRequest_Height)(nil),
		(*FindOffsetRequest_Timestamp)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stream_consumer_v1alpha1_stream_consumer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stream_consumer_v1alpha1_stream_consumer_proto_goTypes,
		DependencyIndexes: file_stream_consumer_v1alpha1_stream_consumer_proto_depIdxs,
		EnumInfos:         file_stream_consumer_v1alpha1_stream_consumer_proto_enumTypes,
		MessageInfos:      file_stream_consumer_v1alpha1_stream_consumer_proto_msgTypes,
	}.Build()
	File_stream_consumer_v1alpha1_stream_consumer_proto = out.File
	file_stream_consumer_v1alpha1_stream_consumer_proto_rawDesc = nil
	file_stream_consumer_v1alpha1_stream_consumer_proto_goTypes = nil
	file_stream_consumer_v1alpha1_stream_consumer_proto_depIdxs = nil
}