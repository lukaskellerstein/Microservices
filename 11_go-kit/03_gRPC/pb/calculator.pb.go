// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	calculator.proto

It has these top-level messages:
	PlusRequest
	PlusResponse
*/
package pb

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

type PlusRequest struct {
	A int32 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
}

func (m *PlusRequest) Reset()                    { *m = PlusRequest{} }
func (m *PlusRequest) String() string            { return proto.CompactTextString(m) }
func (*PlusRequest) ProtoMessage()               {}
func (*PlusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PlusRequest) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *PlusRequest) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type PlusResponse struct {
	Result int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *PlusResponse) Reset()                    { *m = PlusResponse{} }
func (m *PlusResponse) String() string            { return proto.CompactTextString(m) }
func (*PlusResponse) ProtoMessage()               {}
func (*PlusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PlusResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*PlusRequest)(nil), "pb.PlusRequest")
	proto.RegisterType((*PlusResponse)(nil), "pb.PlusResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Calculator service

type CalculatorClient interface {
	Plus(ctx context.Context, in *PlusRequest, opts ...grpc.CallOption) (*PlusResponse, error)
}

type calculatorClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorClient(cc *grpc.ClientConn) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Plus(ctx context.Context, in *PlusRequest, opts ...grpc.CallOption) (*PlusResponse, error) {
	out := new(PlusResponse)
	err := grpc.Invoke(ctx, "/pb.Calculator/Plus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Calculator service

type CalculatorServer interface {
	Plus(context.Context, *PlusRequest) (*PlusResponse, error)
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_Plus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Plus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Calculator/Plus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Plus(ctx, req.(*PlusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Plus",
			Handler:    _Calculator_Plus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator.proto",
}

func init() { proto.RegisterFile("calculator.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x4e, 0xcc, 0x49,
	0x2e, 0xcd, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48,
	0x52, 0xd2, 0xe4, 0xe2, 0x0e, 0xc8, 0x29, 0x2d, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11,
	0xe2, 0xe1, 0x62, 0x4c, 0x94, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62, 0x4c, 0x04, 0xf1, 0x92,
	0x24, 0x98, 0x20, 0xbc, 0x24, 0x25, 0x35, 0x2e, 0x1e, 0x88, 0xd2, 0xe2, 0x82, 0xfc, 0xbc, 0xe2,
	0x54, 0x21, 0x31, 0x2e, 0xb6, 0xa2, 0xd4, 0xe2, 0xd2, 0x9c, 0x12, 0xa8, 0x06, 0x28, 0xcf, 0xc8,
	0x92, 0x8b, 0xcb, 0x19, 0x6e, 0x95, 0x90, 0x36, 0x17, 0x0b, 0x48, 0x97, 0x10, 0xbf, 0x5e, 0x41,
	0x92, 0x1e, 0x92, 0x55, 0x52, 0x02, 0x08, 0x01, 0x88, 0x81, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x87,
	0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x68, 0x8f, 0x41, 0x05, 0xac, 0x00, 0x00, 0x00,
}