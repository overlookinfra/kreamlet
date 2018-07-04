// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messaging/adminCreds.proto

/*
Package messaging is a generated protocol buffer package.

It is generated from these files:
	messaging/adminCreds.proto

It has these top-level messages:
	AdminCredsRequest
	AdminCredsResponse
*/
package messaging

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

type StatusCode int32

const (
	StatusCode_Unknown StatusCode = 0
	StatusCode_Ok      StatusCode = 1
	StatusCode_Failed  StatusCode = 2
)

var StatusCode_name = map[int32]string{
	0: "Unknown",
	1: "Ok",
	2: "Failed",
}
var StatusCode_value = map[string]int32{
	"Unknown": 0,
	"Ok":      1,
	"Failed":  2,
}

func (x StatusCode) String() string {
	return proto.EnumName(StatusCode_name, int32(x))
}
func (StatusCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AdminCredsRequest struct {
}

func (m *AdminCredsRequest) Reset()                    { *m = AdminCredsRequest{} }
func (m *AdminCredsRequest) String() string            { return proto.CompactTextString(m) }
func (*AdminCredsRequest) ProtoMessage()               {}
func (*AdminCredsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AdminCredsResponse struct {
	StatusCode StatusCode `protobuf:"varint,1,opt,name=StatusCode,enum=messaging.StatusCode" json:"StatusCode,omitempty"`
	Message    string     `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	Content    []byte     `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (m *AdminCredsResponse) Reset()                    { *m = AdminCredsResponse{} }
func (m *AdminCredsResponse) String() string            { return proto.CompactTextString(m) }
func (*AdminCredsResponse) ProtoMessage()               {}
func (*AdminCredsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AdminCredsResponse) GetStatusCode() StatusCode {
	if m != nil {
		return m.StatusCode
	}
	return StatusCode_Unknown
}

func (m *AdminCredsResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AdminCredsResponse) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*AdminCredsRequest)(nil), "messaging.AdminCredsRequest")
	proto.RegisterType((*AdminCredsResponse)(nil), "messaging.AdminCredsResponse")
	proto.RegisterEnum("messaging.StatusCode", StatusCode_name, StatusCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AdminCreds service

type AdminCredsClient interface {
	GetAdminCreds(ctx context.Context, in *AdminCredsRequest, opts ...grpc.CallOption) (*AdminCredsResponse, error)
}

type adminCredsClient struct {
	cc *grpc.ClientConn
}

func NewAdminCredsClient(cc *grpc.ClientConn) AdminCredsClient {
	return &adminCredsClient{cc}
}

func (c *adminCredsClient) GetAdminCreds(ctx context.Context, in *AdminCredsRequest, opts ...grpc.CallOption) (*AdminCredsResponse, error) {
	out := new(AdminCredsResponse)
	err := grpc.Invoke(ctx, "/messaging.AdminCreds/GetAdminCreds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AdminCreds service

type AdminCredsServer interface {
	GetAdminCreds(context.Context, *AdminCredsRequest) (*AdminCredsResponse, error)
}

func RegisterAdminCredsServer(s *grpc.Server, srv AdminCredsServer) {
	s.RegisterService(&_AdminCreds_serviceDesc, srv)
}

func _AdminCreds_GetAdminCreds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminCredsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminCredsServer).GetAdminCreds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging.AdminCreds/GetAdminCreds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminCredsServer).GetAdminCreds(ctx, req.(*AdminCredsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdminCreds_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messaging.AdminCreds",
	HandlerType: (*AdminCredsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdminCreds",
			Handler:    _AdminCreds_GetAdminCreds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messaging/adminCreds.proto",
}

func init() { proto.RegisterFile("messaging/adminCreds.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0xcf, 0xcc, 0x4b, 0xd7, 0x4f, 0x4c, 0xc9, 0xcd, 0xcc, 0x73, 0x2e, 0x4a, 0x4d, 0x29,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0xcb, 0x29, 0x09, 0x73, 0x09, 0x3a, 0xc2,
	0xa5, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94, 0xea, 0xb9, 0x84, 0x90, 0x05, 0x8b, 0x0b,
	0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x4c, 0xb9, 0xb8, 0x82, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0x9d, 0xf3,
	0x53, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0xf8, 0x8c, 0x44, 0xf5, 0xe0, 0x46, 0xe9, 0x21, 0x24,
	0x83, 0x90, 0x14, 0x0a, 0x49, 0x70, 0xb1, 0xfb, 0x82, 0xd5, 0xa4, 0x4a, 0x30, 0x29, 0x30, 0x6a,
	0x70, 0x06, 0xc1, 0xb8, 0x20, 0x19, 0xe7, 0xfc, 0xbc, 0x92, 0xd4, 0xbc, 0x12, 0x09, 0x66, 0x05,
	0x46, 0x0d, 0x9e, 0x20, 0x18, 0x57, 0x4b, 0x17, 0xd9, 0x2a, 0x21, 0x6e, 0x2e, 0xf6, 0xd0, 0xbc,
	0xec, 0xbc, 0xfc, 0xf2, 0x3c, 0x01, 0x06, 0x21, 0x36, 0x2e, 0x26, 0xff, 0x6c, 0x01, 0x46, 0x21,
	0x2e, 0x2e, 0x36, 0xb7, 0xc4, 0xcc, 0x9c, 0xd4, 0x14, 0x01, 0x26, 0xa3, 0x18, 0x2e, 0x2e, 0x84,
	0x7b, 0x85, 0xfc, 0xb8, 0x78, 0xdd, 0x53, 0x4b, 0x90, 0x04, 0x64, 0x90, 0x1c, 0x89, 0xe1, 0x59,
	0x29, 0x59, 0x1c, 0xb2, 0x10, 0x5f, 0x2b, 0x31, 0x24, 0xb1, 0x81, 0x03, 0xcd, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0xf9, 0x87, 0xf9, 0x80, 0x52, 0x01, 0x00, 0x00,
}