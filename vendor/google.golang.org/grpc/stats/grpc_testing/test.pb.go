// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

/*
Package grpc_testing is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	SimpleRequest
	SimpleResponse
*/
package grpc_testing

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Unary request.
type SimpleRequest struct {
	Id int32 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
}

func (m *SimpleRequest) Reset()                    { *m = SimpleRequest{} }
func (m *SimpleRequest) String() string            { return proto.CompactTextString(m) }
func (*SimpleRequest) ProtoMessage()               {}
func (*SimpleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Unary response, as configured by the request.
type SimpleResponse struct {
	Id int32 `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
}

func (m *SimpleResponse) Reset()                    { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string            { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()               {}
func (*SimpleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*SimpleRequest)(nil), "grpc.testing.SimpleRequest")
	proto.RegisterType((*SimpleResponse)(nil), "grpc.testing.SimpleResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TestService service

type TestServiceClient interface {
	// One request followed by one response.
	// The server returns the client id as-is.
	UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	// A sequence of requests with each request served by the server immediately.
	// As one request could lead to multiple responses, this interface
	// demonstrates the idea of full duplexing.
	FullDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_FullDuplexCallClient, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := grpc.Invoke(ctx, "/grpc.testing.TestService/UnaryCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) FullDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_FullDuplexCallClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[0], c.cc, "/grpc.testing.TestService/FullDuplexCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceFullDuplexCallClient{stream}
	return x, nil
}

type TestService_FullDuplexCallClient interface {
	Send(*SimpleRequest) error
	Recv() (*SimpleResponse, error)
	grpc.ClientStream
}

type testServiceFullDuplexCallClient struct {
	grpc.ClientStream
}

func (x *testServiceFullDuplexCallClient) Send(m *SimpleRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceFullDuplexCallClient) Recv() (*SimpleResponse, error) {
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TestService service

type TestServiceServer interface {
	// One request followed by one response.
	// The server returns the client id as-is.
	UnaryCall(context.Context, *SimpleRequest) (*SimpleResponse, error)
	// A sequence of requests with each request served by the server immediately.
	// As one request could lead to multiple responses, this interface
	// demonstrates the idea of full duplexing.
	FullDuplexCall(TestService_FullDuplexCallServer) error
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_UnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).UnaryCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.testing.TestService/UnaryCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).UnaryCall(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_FullDuplexCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).FullDuplexCall(&testServiceFullDuplexCallServer{stream})
}

type TestService_FullDuplexCallServer interface {
	Send(*SimpleResponse) error
	Recv() (*SimpleRequest, error)
	grpc.ServerStream
}

type testServiceFullDuplexCallServer struct {
	grpc.ServerStream
}

func (x *testServiceFullDuplexCallServer) Send(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceFullDuplexCallServer) Recv() (*SimpleRequest, error) {
	m := new(SimpleRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryCall",
			Handler:    _TestService_UnaryCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FullDuplexCall",
			Handler:       _TestService_FullDuplexCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "test.proto",
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x49, 0x2f, 0x2a, 0x48, 0xd6, 0x03, 0x09, 0x64,
	0xe6, 0xa5, 0x2b, 0xc9, 0x73, 0xf1, 0x06, 0x67, 0xe6, 0x16, 0xe4, 0xa4, 0x06, 0xa5, 0x16, 0x96,
	0xa6, 0x16, 0x97, 0x08, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x29, 0x30, 0x6a, 0xb0, 0x06,
	0x31, 0x65, 0xa6, 0x28, 0x29, 0x70, 0xf1, 0xc1, 0x14, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x42,
	0x55, 0x30, 0xc3, 0x54, 0x18, 0x2d, 0x63, 0xe4, 0xe2, 0x0e, 0x49, 0x2d, 0x2e, 0x09, 0x4e, 0x2d,
	0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x72, 0xe3, 0xe2, 0x0c, 0xcd, 0x4b, 0x2c, 0xaa, 0x74, 0x4e, 0xcc,
	0xc9, 0x11, 0x92, 0xd6, 0x43, 0xb6, 0x4e, 0x0f, 0xc5, 0x2e, 0x29, 0x19, 0xec, 0x92, 0x50, 0x7b,
	0xfc, 0xb9, 0xf8, 0xdc, 0x4a, 0x73, 0x72, 0x5c, 0x4a, 0x0b, 0x72, 0x52, 0x2b, 0x28, 0x34, 0x4c,
	0x83, 0xd1, 0x80, 0x31, 0x89, 0x0d, 0x1c, 0x00, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8d,
	0x82, 0x5b, 0xdd, 0x0e, 0x01, 0x00, 0x00,
}
