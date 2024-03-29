// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: solpipe.proto

package solpipe

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
	WorkerStatus_OnStatus_FullMethodName = "/solpipe.WorkerStatus/OnStatus"
)

// WorkerStatusClient is the client API for WorkerStatus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkerStatusClient interface {
	OnStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (WorkerStatus_OnStatusClient, error)
}

type workerStatusClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerStatusClient(cc grpc.ClientConnInterface) WorkerStatusClient {
	return &workerStatusClient{cc}
}

func (c *workerStatusClient) OnStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (WorkerStatus_OnStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &WorkerStatus_ServiceDesc.Streams[0], WorkerStatus_OnStatus_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &workerStatusOnStatusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WorkerStatus_OnStatusClient interface {
	Recv() (*CapacityResponse, error)
	grpc.ClientStream
}

type workerStatusOnStatusClient struct {
	grpc.ClientStream
}

func (x *workerStatusOnStatusClient) Recv() (*CapacityResponse, error) {
	m := new(CapacityResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WorkerStatusServer is the server API for WorkerStatus service.
// All implementations must embed UnimplementedWorkerStatusServer
// for forward compatibility
type WorkerStatusServer interface {
	OnStatus(*Empty, WorkerStatus_OnStatusServer) error
	mustEmbedUnimplementedWorkerStatusServer()
}

// UnimplementedWorkerStatusServer must be embedded to have forward compatible implementations.
type UnimplementedWorkerStatusServer struct {
}

func (UnimplementedWorkerStatusServer) OnStatus(*Empty, WorkerStatus_OnStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method OnStatus not implemented")
}
func (UnimplementedWorkerStatusServer) mustEmbedUnimplementedWorkerStatusServer() {}

// UnsafeWorkerStatusServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkerStatusServer will
// result in compilation errors.
type UnsafeWorkerStatusServer interface {
	mustEmbedUnimplementedWorkerStatusServer()
}

func RegisterWorkerStatusServer(s grpc.ServiceRegistrar, srv WorkerStatusServer) {
	s.RegisterService(&WorkerStatus_ServiceDesc, srv)
}

func _WorkerStatus_OnStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WorkerStatusServer).OnStatus(m, &workerStatusOnStatusServer{stream})
}

type WorkerStatus_OnStatusServer interface {
	Send(*CapacityResponse) error
	grpc.ServerStream
}

type workerStatusOnStatusServer struct {
	grpc.ServerStream
}

func (x *workerStatusOnStatusServer) Send(m *CapacityResponse) error {
	return x.ServerStream.SendMsg(m)
}

// WorkerStatus_ServiceDesc is the grpc.ServiceDesc for WorkerStatus service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkerStatus_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "solpipe.WorkerStatus",
	HandlerType: (*WorkerStatusServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OnStatus",
			Handler:       _WorkerStatus_OnStatus_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "solpipe.proto",
}
