// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: shared/proto/pipeline_reporter.proto

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

// PipelineReporterClient is the client API for PipelineReporter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PipelineReporterClient interface {
	PipelineStart(ctx context.Context, in *PipelineStartRequest, opts ...grpc.CallOption) (*Void, error)
	PipelineEnd(ctx context.Context, in *PipelineEndRequest, opts ...grpc.CallOption) (*Void, error)
}

type pipelineReporterClient struct {
	cc grpc.ClientConnInterface
}

func NewPipelineReporterClient(cc grpc.ClientConnInterface) PipelineReporterClient {
	return &pipelineReporterClient{cc}
}

func (c *pipelineReporterClient) PipelineStart(ctx context.Context, in *PipelineStartRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/PipelineReporter/PipelineStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineReporterClient) PipelineEnd(ctx context.Context, in *PipelineEndRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/PipelineReporter/PipelineEnd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PipelineReporterServer is the server API for PipelineReporter service.
// All implementations must embed UnimplementedPipelineReporterServer
// for forward compatibility
type PipelineReporterServer interface {
	PipelineStart(context.Context, *PipelineStartRequest) (*Void, error)
	PipelineEnd(context.Context, *PipelineEndRequest) (*Void, error)
	mustEmbedUnimplementedPipelineReporterServer()
}

// UnimplementedPipelineReporterServer must be embedded to have forward compatible implementations.
type UnimplementedPipelineReporterServer struct {
}

func (UnimplementedPipelineReporterServer) PipelineStart(context.Context, *PipelineStartRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PipelineStart not implemented")
}
func (UnimplementedPipelineReporterServer) PipelineEnd(context.Context, *PipelineEndRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PipelineEnd not implemented")
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

func _PipelineReporter_PipelineStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PipelineStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineReporterServer).PipelineStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PipelineReporter/PipelineStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineReporterServer).PipelineStart(ctx, req.(*PipelineStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineReporter_PipelineEnd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PipelineEndRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineReporterServer).PipelineEnd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PipelineReporter/PipelineEnd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineReporterServer).PipelineEnd(ctx, req.(*PipelineEndRequest))
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
			MethodName: "PipelineStart",
			Handler:    _PipelineReporter_PipelineStart_Handler,
		},
		{
			MethodName: "PipelineEnd",
			Handler:    _PipelineReporter_PipelineEnd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shared/proto/pipeline_reporter.proto",
}