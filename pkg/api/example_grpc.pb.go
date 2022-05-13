// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: example.proto

package api

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

// ExampleServiceClient is the client API for ExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExampleServiceClient interface {
	CreateExample(ctx context.Context, in *CreateExampleRequest, opts ...grpc.CallOption) (*Example, error)
}

type exampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleServiceClient(cc grpc.ClientConnInterface) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) CreateExample(ctx context.Context, in *CreateExampleRequest, opts ...grpc.CallOption) (*Example, error) {
	out := new(Example)
	err := c.cc.Invoke(ctx, "/nokamoto.github.com.ExampleService/CreateExample", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServiceServer is the server API for ExampleService service.
// All implementations must embed UnimplementedExampleServiceServer
// for forward compatibility
type ExampleServiceServer interface {
	CreateExample(context.Context, *CreateExampleRequest) (*Example, error)
	mustEmbedUnimplementedExampleServiceServer()
}

// UnimplementedExampleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExampleServiceServer struct {
}

func (UnimplementedExampleServiceServer) CreateExample(context.Context, *CreateExampleRequest) (*Example, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExample not implemented")
}
func (UnimplementedExampleServiceServer) mustEmbedUnimplementedExampleServiceServer() {}

// UnsafeExampleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExampleServiceServer will
// result in compilation errors.
type UnsafeExampleServiceServer interface {
	mustEmbedUnimplementedExampleServiceServer()
}

func RegisterExampleServiceServer(s grpc.ServiceRegistrar, srv ExampleServiceServer) {
	s.RegisterService(&ExampleService_ServiceDesc, srv)
}

func _ExampleService_CreateExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).CreateExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nokamoto.github.com.ExampleService/CreateExample",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).CreateExample(ctx, req.(*CreateExampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExampleService_ServiceDesc is the grpc.ServiceDesc for ExampleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExampleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nokamoto.github.com.ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExample",
			Handler:    _ExampleService_CreateExample_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example.proto",
}
