// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: stream_stats/v1alpha1/stream_stats.proto

package stream_statsv1alpha1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	StreamStatsService_GetStreamStats_FullMethodName  = "/stream_stats.v1alpha1.StreamStatsService/GetStreamStats"
	StreamStatsService_GetStreamsStats_FullMethodName = "/stream_stats.v1alpha1.StreamStatsService/GetStreamsStats"
)

// StreamStatsServiceClient is the client API for StreamStatsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamStatsServiceClient interface {
	GetStreamStats(ctx context.Context, in *GetStreamStatsRequest, opts ...grpc.CallOption) (*GetStreamStatsResponse, error)
	GetStreamsStats(ctx context.Context, in *GetStreamsStatsRequest, opts ...grpc.CallOption) (*GetStreamsStatsResponse, error)
}

type streamStatsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamStatsServiceClient(cc grpc.ClientConnInterface) StreamStatsServiceClient {
	return &streamStatsServiceClient{cc}
}

func (c *streamStatsServiceClient) GetStreamStats(ctx context.Context, in *GetStreamStatsRequest, opts ...grpc.CallOption) (*GetStreamStatsResponse, error) {
	out := new(GetStreamStatsResponse)
	err := c.cc.Invoke(ctx, StreamStatsService_GetStreamStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamStatsServiceClient) GetStreamsStats(ctx context.Context, in *GetStreamsStatsRequest, opts ...grpc.CallOption) (*GetStreamsStatsResponse, error) {
	out := new(GetStreamsStatsResponse)
	err := c.cc.Invoke(ctx, StreamStatsService_GetStreamsStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamStatsServiceServer is the server API for StreamStatsService service.
// All implementations should embed UnimplementedStreamStatsServiceServer
// for forward compatibility
type StreamStatsServiceServer interface {
	GetStreamStats(context.Context, *GetStreamStatsRequest) (*GetStreamStatsResponse, error)
	GetStreamsStats(context.Context, *GetStreamsStatsRequest) (*GetStreamsStatsResponse, error)
}

// UnimplementedStreamStatsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedStreamStatsServiceServer struct {
}

func (UnimplementedStreamStatsServiceServer) GetStreamStats(context.Context, *GetStreamStatsRequest) (*GetStreamStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStreamStats not implemented")
}
func (UnimplementedStreamStatsServiceServer) GetStreamsStats(context.Context, *GetStreamsStatsRequest) (*GetStreamsStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStreamsStats not implemented")
}

// UnsafeStreamStatsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamStatsServiceServer will
// result in compilation errors.
type UnsafeStreamStatsServiceServer interface {
	mustEmbedUnimplementedStreamStatsServiceServer()
}

func RegisterStreamStatsServiceServer(s grpc.ServiceRegistrar, srv StreamStatsServiceServer) {
	s.RegisterService(&StreamStatsService_ServiceDesc, srv)
}

func _StreamStatsService_GetStreamStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStreamStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamStatsServiceServer).GetStreamStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamStatsService_GetStreamStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamStatsServiceServer).GetStreamStats(ctx, req.(*GetStreamStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamStatsService_GetStreamsStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStreamsStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamStatsServiceServer).GetStreamsStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamStatsService_GetStreamsStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamStatsServiceServer).GetStreamsStats(ctx, req.(*GetStreamsStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StreamStatsService_ServiceDesc is the grpc.ServiceDesc for StreamStatsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamStatsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stream_stats.v1alpha1.StreamStatsService",
	HandlerType: (*StreamStatsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStreamStats",
			Handler:    _StreamStatsService_GetStreamStats_Handler,
		},
		{
			MethodName: "GetStreamsStats",
			Handler:    _StreamStatsService_GetStreamsStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stream_stats/v1alpha1/stream_stats.proto",
}
