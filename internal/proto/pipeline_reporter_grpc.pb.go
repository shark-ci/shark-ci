// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: internal/proto/pipeline_reporter.proto

package proto

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
	PipelineReporter_PipelineStarted_FullMethodName   = "/PipelineReporter/PipelineStarted"
	PipelineReporter_PipelineFinnished_FullMethodName = "/PipelineReporter/PipelineFinnished"
)

// PipelineReporterClient is the client API for PipelineReporter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PipelineReporterClient interface {
	PipelineStarted(ctx context.Context, in *PipelineStartedRequest, opts ...grpc.CallOption) (*Empty, error)
	PipelineFinnished(ctx context.Context, in *PipelineFinnishedRequest, opts ...grpc.CallOption) (*Empty, error)
}

type pipelineReporterClient struct {
	cc grpc.ClientConnInterface
}

func NewPipelineReporterClient(cc grpc.ClientConnInterface) PipelineReporterClient {
	return &pipelineReporterClient{cc}
}

func (c *pipelineReporterClient) PipelineStarted(ctx context.Context, in *PipelineStartedRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, PipelineReporter_PipelineStarted_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineReporterClient) PipelineFinnished(ctx context.Context, in *PipelineFinnishedRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, PipelineReporter_PipelineFinnished_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PipelineReporterServer is the server API for PipelineReporter service.
// All implementations must embed UnimplementedPipelineReporterServer
// for forward compatibility
type PipelineReporterServer interface {
	PipelineStarted(context.Context, *PipelineStartedRequest) (*Empty, error)
	PipelineFinnished(context.Context, *PipelineFinnishedRequest) (*Empty, error)
	mustEmbedUnimplementedPipelineReporterServer()
}

// UnimplementedPipelineReporterServer must be embedded to have forward compatible implementations.
type UnimplementedPipelineReporterServer struct {
}

func (UnimplementedPipelineReporterServer) PipelineStarted(context.Context, *PipelineStartedRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PipelineStarted not implemented")
}
func (UnimplementedPipelineReporterServer) PipelineFinnished(context.Context, *PipelineFinnishedRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PipelineFinnished not implemented")
}
func (UnimplementedPipelineReporterServer) mustEmbedUnimplementedPipelineReporterServer() {}

// UnsafePipelineReporterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PipelineReporterServer will
// result in compilation errors.
type UnsafePipelineReporterServer interface {
	mustEmbedUnimplementedPipelineReporterServer()
}

func RegisterPipelineReporterServer(s grpc.ServiceRegistrar, srv PipelineReporterServer) {
	s.RegisterService(&PipelineReporter_ServiceDesc, srv)
}

func _PipelineReporter_PipelineStarted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PipelineStartedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineReporterServer).PipelineStarted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PipelineReporter_PipelineStarted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineReporterServer).PipelineStarted(ctx, req.(*PipelineStartedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineReporter_PipelineFinnished_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PipelineFinnishedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineReporterServer).PipelineFinnished(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PipelineReporter_PipelineFinnished_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineReporterServer).PipelineFinnished(ctx, req.(*PipelineFinnishedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PipelineReporter_ServiceDesc is the grpc.ServiceDesc for PipelineReporter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PipelineReporter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PipelineReporter",
	HandlerType: (*PipelineReporterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PipelineStarted",
			Handler:    _PipelineReporter_PipelineStarted_Handler,
		},
		{
			MethodName: "PipelineFinnished",
			Handler:    _PipelineReporter_PipelineFinnished_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/pipeline_reporter.proto",
}