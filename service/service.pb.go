// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	InterstallerRequest
	InterstallerReply
*/
package service

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

type InterstallerRequest struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *InterstallerRequest) Reset()                    { *m = InterstallerRequest{} }
func (m *InterstallerRequest) String() string            { return proto.CompactTextString(m) }
func (*InterstallerRequest) ProtoMessage()               {}
func (*InterstallerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InterstallerRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type InterstallerReply struct {
	Message          *string `protobuf:"bytes,1,req,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *InterstallerReply) Reset()                    { *m = InterstallerReply{} }
func (m *InterstallerReply) String() string            { return proto.CompactTextString(m) }
func (*InterstallerReply) ProtoMessage()               {}
func (*InterstallerReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *InterstallerReply) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*InterstallerRequest)(nil), "InterstallerRequest")
	proto.RegisterType((*InterstallerReply)(nil), "InterstallerReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for InterstallerCall service

type InterstallerCallClient interface {
	MakeInterstallerCall(ctx context.Context, in *InterstallerRequest, opts ...grpc.CallOption) (*InterstallerReply, error)
}

type interstallerCallClient struct {
	cc *grpc.ClientConn
}

func NewInterstallerCallClient(cc *grpc.ClientConn) InterstallerCallClient {
	return &interstallerCallClient{cc}
}

func (c *interstallerCallClient) MakeInterstallerCall(ctx context.Context, in *InterstallerRequest, opts ...grpc.CallOption) (*InterstallerReply, error) {
	out := new(InterstallerReply)
	err := grpc.Invoke(ctx, "/InterstallerCall/MakeInterstallerCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for InterstallerCall service

type InterstallerCallServer interface {
	MakeInterstallerCall(context.Context, *InterstallerRequest) (*InterstallerReply, error)
}

func RegisterInterstallerCallServer(s *grpc.Server, srv InterstallerCallServer) {
	s.RegisterService(&_InterstallerCall_serviceDesc, srv)
}

func _InterstallerCall_MakeInterstallerCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InterstallerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterstallerCallServer).MakeInterstallerCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InterstallerCall/MakeInterstallerCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterstallerCallServer).MakeInterstallerCall(ctx, req.(*InterstallerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InterstallerCall_serviceDesc = grpc.ServiceDesc{
	ServiceName: "InterstallerCall",
	HandlerType: (*InterstallerCallServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeInterstallerCall",
			Handler:    _InterstallerCall_MakeInterstallerCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe4, 0x12, 0xf6, 0xcc, 0x2b,
	0x49, 0x2d, 0x2a, 0x2e, 0x49, 0xcc, 0xc9, 0x49, 0x2d, 0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e,
	0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd2, 0xe0, 0x0c, 0x02,
	0xb3, 0x95, 0x74, 0xb9, 0x04, 0x51, 0x95, 0x16, 0xe4, 0x54, 0x0a, 0x49, 0x70, 0xb1, 0xe7, 0xa6,
	0x16, 0x17, 0x27, 0xa6, 0xc3, 0xd4, 0xc2, 0xb8, 0x46, 0x61, 0x5c, 0x02, 0xc8, 0xca, 0x9d, 0x13,
	0x73, 0x72, 0x84, 0x9c, 0xb8, 0x44, 0x7c, 0x13, 0xb3, 0x53, 0x31, 0xc4, 0x45, 0xf4, 0xb0, 0x38,
	0x42, 0x4a, 0x48, 0x0f, 0xc3, 0x3e, 0x25, 0x06, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x82,
	0xa7, 0xca, 0xc1, 0x00, 0x00, 0x00,
}
