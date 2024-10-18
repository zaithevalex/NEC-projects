// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: protobuf/desc.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DescriptionService_GetDescription_FullMethodName = "/main.DescriptionService/GetDescription"
)

// DescriptionServiceClient is the client API for DescriptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DescriptionServiceClient interface {
	GetDescription(ctx context.Context, in *DescriptionRequest, opts ...grpc.CallOption) (*DescriptionReply, error)
}

type descriptionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDescriptionServiceClient(cc grpc.ClientConnInterface) DescriptionServiceClient {
	return &descriptionServiceClient{cc}
}

func (c *descriptionServiceClient) GetDescription(ctx context.Context, in *DescriptionRequest, opts ...grpc.CallOption) (*DescriptionReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescriptionReply)
	err := c.cc.Invoke(ctx, DescriptionService_GetDescription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DescriptionServiceServer is the server API for DescriptionService service.
// All implementations must embed UnimplementedDescriptionServiceServer
// for forward compatibility.
type DescriptionServiceServer interface {
	GetDescription(context.Context, *DescriptionRequest) (*DescriptionReply, error)
	mustEmbedUnimplementedDescriptionServiceServer()
}

// UnimplementedDescriptionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDescriptionServiceServer struct{}

func (UnimplementedDescriptionServiceServer) GetDescription(context.Context, *DescriptionRequest) (*DescriptionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDescription not implemented")
}
func (UnimplementedDescriptionServiceServer) mustEmbedUnimplementedDescriptionServiceServer() {}
func (UnimplementedDescriptionServiceServer) testEmbeddedByValue()                            {}

// UnsafeDescriptionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DescriptionServiceServer will
// result in compilation errors.
type UnsafeDescriptionServiceServer interface {
	mustEmbedUnimplementedDescriptionServiceServer()
}

func RegisterDescriptionServiceServer(s grpc.ServiceRegistrar, srv DescriptionServiceServer) {
	// If the following call pancis, it indicates UnimplementedDescriptionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DescriptionService_ServiceDesc, srv)
}

func _DescriptionService_GetDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DescriptionServiceServer).GetDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DescriptionService_GetDescription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DescriptionServiceServer).GetDescription(ctx, req.(*DescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DescriptionService_ServiceDesc is the grpc.ServiceDesc for DescriptionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DescriptionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.DescriptionService",
	HandlerType: (*DescriptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDescription",
			Handler:    _DescriptionService_GetDescription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/desc.proto",
}
