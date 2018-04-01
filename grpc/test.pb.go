// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/test.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	grpc/test.proto

It has these top-level messages:
	TestResponse
	TestRequest
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
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

type TestResponse struct {
	StrData string `protobuf:"bytes,1,opt,name=StrData" json:"StrData,omitempty"`
	IntData int32  `protobuf:"varint,2,opt,name=IntData" json:"IntData,omitempty"`
	ID      int32  `protobuf:"varint,3,opt,name=ID" json:"ID,omitempty"`
}

func (m *TestResponse) Reset()                    { *m = TestResponse{} }
func (m *TestResponse) String() string            { return proto.CompactTextString(m) }
func (*TestResponse) ProtoMessage()               {}
func (*TestResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TestResponse) GetStrData() string {
	if m != nil {
		return m.StrData
	}
	return ""
}

func (m *TestResponse) GetIntData() int32 {
	if m != nil {
		return m.IntData
	}
	return 0
}

func (m *TestResponse) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type TestRequest struct {
	ID int32 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *TestRequest) Reset()                    { *m = TestRequest{} }
func (m *TestRequest) String() string            { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()               {}
func (*TestRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TestRequest) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func init() {
	proto.RegisterType((*TestResponse)(nil), "TestResponse")
	proto.RegisterType((*TestRequest)(nil), "TestRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for TestDataProvider service

type TestDataProviderClient interface {
	GetTestData(ctx context.Context, in *TestRequest, opts ...grpc1.CallOption) (*TestResponse, error)
}

type testDataProviderClient struct {
	cc *grpc1.ClientConn
}

func NewTestDataProviderClient(cc *grpc1.ClientConn) TestDataProviderClient {
	return &testDataProviderClient{cc}
}

func (c *testDataProviderClient) GetTestData(ctx context.Context, in *TestRequest, opts ...grpc1.CallOption) (*TestResponse, error) {
	out := new(TestResponse)
	err := grpc1.Invoke(ctx, "/TestDataProvider/GetTestData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TestDataProvider service

type TestDataProviderServer interface {
	GetTestData(context.Context, *TestRequest) (*TestResponse, error)
}

func RegisterTestDataProviderServer(s *grpc1.Server, srv TestDataProviderServer) {
	s.RegisterService(&_TestDataProvider_serviceDesc, srv)
}

func _TestDataProvider_GetTestData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestDataProviderServer).GetTestData(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TestDataProvider/GetTestData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestDataProviderServer).GetTestData(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestDataProvider_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "TestDataProvider",
	HandlerType: (*TestDataProviderServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "GetTestData",
			Handler:    _TestDataProvider_GetTestData_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "grpc/test.proto",
}

func init() { proto.RegisterFile("grpc/test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xcf, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x49, 0x45, 0xc5, 0xec, 0xfa, 0x83, 0x9c, 0x8a, 0x20, 0x2c, 0x7b, 0x5a, 0x84, 0xa4,
	0xa0, 0x27, 0x11, 0x3c, 0x48, 0x41, 0x7a, 0x93, 0xe8, 0xc9, 0x5b, 0x1a, 0x87, 0x6c, 0xd8, 0x9a,
	0xc4, 0x64, 0xea, 0xdf, 0x2f, 0x89, 0x06, 0x7a, 0x7c, 0xef, 0x63, 0x1e, 0xf3, 0xd1, 0x4b, 0x13,
	0x83, 0xee, 0x10, 0x12, 0x8a, 0x10, 0x3d, 0xfa, 0xad, 0xa4, 0xeb, 0x77, 0x48, 0x28, 0x21, 0x05,
	0xef, 0x12, 0xb0, 0x96, 0x9e, 0xbe, 0x61, 0xec, 0x15, 0xaa, 0x96, 0x6c, 0xc8, 0xee, 0x4c, 0xd6,
	0x98, 0xc9, 0xe0, 0xb0, 0x90, 0x66, 0x43, 0x76, 0xc7, 0xb2, 0x46, 0x76, 0x41, 0x9b, 0xa1, 0x6f,
	0x8f, 0x4a, 0xd9, 0x0c, 0xfd, 0xf6, 0x86, 0xae, 0xfe, 0x36, 0xbf, 0x67, 0x48, 0xf8, 0x8f, 0x49,
	0xc5, 0x77, 0x4f, 0xf4, 0x2a, 0xe3, 0x7c, 0xfa, 0x1a, 0xfd, 0x8f, 0xfd, 0x84, 0xc8, 0x6e, 0xe9,
	0xea, 0x05, 0xb0, 0xd6, 0x6c, 0x2d, 0x16, 0x03, 0xd7, 0xe7, 0x62, 0xf9, 0xe2, 0xf3, 0xe3, 0xc7,
	0x83, 0xb1, 0xb8, 0x9f, 0x47, 0xa1, 0xfd, 0x57, 0x67, 0x8d, 0x8f, 0xfc, 0xa0, 0x62, 0x98, 0x0f,
	0x7b, 0xeb, 0xba, 0xec, 0xc7, 0xf5, 0x64, 0xc1, 0x21, 0x1f, 0xd5, 0xa4, 0x9c, 0xb6, 0xce, 0xf0,
	0x2c, 0x5c, 0xd0, 0x78, 0x52, 0xb4, 0xef, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xf5, 0x15,
	0x0a, 0x09, 0x01, 0x00, 0x00,
}
