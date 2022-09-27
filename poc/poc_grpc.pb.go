// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: poc.proto

package poc

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	SayHello(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*Response, error)
	SayHelloWithClientStream(ctx context.Context, opts ...grpc.CallOption) (Service_SayHelloWithClientStreamClient, error)
	SayHelloWithServerStream(ctx context.Context, in *MultipleRequest, opts ...grpc.CallOption) (Service_SayHelloWithServerStreamClient, error)
	SayHelloWithClientServerStream(ctx context.Context, opts ...grpc.CallOption) (Service_SayHelloWithClientServerStreamClient, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) SayHello(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/poc.Service/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SayHelloWithClientStream(ctx context.Context, opts ...grpc.CallOption) (Service_SayHelloWithClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[0], "/poc.Service/SayHelloWithClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceSayHelloWithClientStreamClient{stream}
	return x, nil
}

type Service_SayHelloWithClientStreamClient interface {
	Send(*SingleRequest) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type serviceSayHelloWithClientStreamClient struct {
	grpc.ClientStream
}

func (x *serviceSayHelloWithClientStreamClient) Send(m *SingleRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serviceSayHelloWithClientStreamClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) SayHelloWithServerStream(ctx context.Context, in *MultipleRequest, opts ...grpc.CallOption) (Service_SayHelloWithServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[1], "/poc.Service/SayHelloWithServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceSayHelloWithServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_SayHelloWithServerStreamClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type serviceSayHelloWithServerStreamClient struct {
	grpc.ClientStream
}

func (x *serviceSayHelloWithServerStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) SayHelloWithClientServerStream(ctx context.Context, opts ...grpc.CallOption) (Service_SayHelloWithClientServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[2], "/poc.Service/SayHelloWithClientServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceSayHelloWithClientServerStreamClient{stream}
	return x, nil
}

type Service_SayHelloWithClientServerStreamClient interface {
	Send(*SingleRequest) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type serviceSayHelloWithClientServerStreamClient struct {
	grpc.ClientStream
}

func (x *serviceSayHelloWithClientServerStreamClient) Send(m *SingleRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serviceSayHelloWithClientServerStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	SayHello(context.Context, *SingleRequest) (*Response, error)
	SayHelloWithClientStream(Service_SayHelloWithClientStreamServer) error
	SayHelloWithServerStream(*MultipleRequest, Service_SayHelloWithServerStreamServer) error
	SayHelloWithClientServerStream(Service_SayHelloWithClientServerStreamServer) error
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) SayHello(context.Context, *SingleRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedServiceServer) SayHelloWithClientStream(Service_SayHelloWithClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloWithClientStream not implemented")
}
func (UnimplementedServiceServer) SayHelloWithServerStream(*MultipleRequest, Service_SayHelloWithServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloWithServerStream not implemented")
}
func (UnimplementedServiceServer) SayHelloWithClientServerStream(Service_SayHelloWithClientServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloWithClientServerStream not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poc.Service/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).SayHello(ctx, req.(*SingleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_SayHelloWithClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).SayHelloWithClientStream(&serviceSayHelloWithClientStreamServer{stream})
}

type Service_SayHelloWithClientStreamServer interface {
	SendAndClose(*Response) error
	Recv() (*SingleRequest, error)
	grpc.ServerStream
}

type serviceSayHelloWithClientStreamServer struct {
	grpc.ServerStream
}

func (x *serviceSayHelloWithClientStreamServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serviceSayHelloWithClientStreamServer) Recv() (*SingleRequest, error) {
	m := new(SingleRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Service_SayHelloWithServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MultipleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).SayHelloWithServerStream(m, &serviceSayHelloWithServerStreamServer{stream})
}

type Service_SayHelloWithServerStreamServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type serviceSayHelloWithServerStreamServer struct {
	grpc.ServerStream
}

func (x *serviceSayHelloWithServerStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _Service_SayHelloWithClientServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).SayHelloWithClientServerStream(&serviceSayHelloWithClientServerStreamServer{stream})
}

type Service_SayHelloWithClientServerStreamServer interface {
	Send(*Response) error
	Recv() (*SingleRequest, error)
	grpc.ServerStream
}

type serviceSayHelloWithClientServerStreamServer struct {
	grpc.ServerStream
}

func (x *serviceSayHelloWithClientServerStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serviceSayHelloWithClientServerStreamServer) Recv() (*SingleRequest, error) {
	m := new(SingleRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "poc.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Service_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloWithClientStream",
			Handler:       _Service_SayHelloWithClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SayHelloWithServerStream",
			Handler:       _Service_SayHelloWithServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayHelloWithClientServerStream",
			Handler:       _Service_SayHelloWithClientServerStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "poc.proto",
}