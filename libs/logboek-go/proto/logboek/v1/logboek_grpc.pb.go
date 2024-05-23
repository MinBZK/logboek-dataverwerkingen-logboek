// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: proto/logboek/v1/logboek.proto

package v1

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
	LogboekService_Export_FullMethodName = "/logboek.v1.LogboekService/Export"
)

// LogboekServiceClient is the client API for LogboekService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogboekServiceClient interface {
	Export(ctx context.Context, in *ExportOperationsRequest, opts ...grpc.CallOption) (*ExportOperationsResponse, error)
}

type logboekServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogboekServiceClient(cc grpc.ClientConnInterface) LogboekServiceClient {
	return &logboekServiceClient{cc}
}

func (c *logboekServiceClient) Export(ctx context.Context, in *ExportOperationsRequest, opts ...grpc.CallOption) (*ExportOperationsResponse, error) {
	out := new(ExportOperationsResponse)
	err := c.cc.Invoke(ctx, LogboekService_Export_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogboekServiceServer is the server API for LogboekService service.
// All implementations must embed UnimplementedLogboekServiceServer
// for forward compatibility
type LogboekServiceServer interface {
	Export(context.Context, *ExportOperationsRequest) (*ExportOperationsResponse, error)
	mustEmbedUnimplementedLogboekServiceServer()
}

// UnimplementedLogboekServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLogboekServiceServer struct {
}

func (UnimplementedLogboekServiceServer) Export(context.Context, *ExportOperationsRequest) (*ExportOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Export not implemented")
}
func (UnimplementedLogboekServiceServer) mustEmbedUnimplementedLogboekServiceServer() {}

// UnsafeLogboekServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogboekServiceServer will
// result in compilation errors.
type UnsafeLogboekServiceServer interface {
	mustEmbedUnimplementedLogboekServiceServer()
}

func RegisterLogboekServiceServer(s grpc.ServiceRegistrar, srv LogboekServiceServer) {
	s.RegisterService(&LogboekService_ServiceDesc, srv)
}

func _LogboekService_Export_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogboekServiceServer).Export(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogboekService_Export_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogboekServiceServer).Export(ctx, req.(*ExportOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogboekService_ServiceDesc is the grpc.ServiceDesc for LogboekService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogboekService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logboek.v1.LogboekService",
	HandlerType: (*LogboekServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Export",
			Handler:    _LogboekService_Export_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/logboek/v1/logboek.proto",
}