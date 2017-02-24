// Code generated by protoc-gen-go.
// source: auth.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	auth.proto

It has these top-level messages:
	LogInRequest
	CredentialsRequest
	AuthResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

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

// The request message containing the user's name.
type LogInRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LogInRequest) Reset()                    { *m = LogInRequest{} }
func (m *LogInRequest) String() string            { return proto.CompactTextString(m) }
func (*LogInRequest) ProtoMessage()               {}
func (*LogInRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LogInRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LogInRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CredentialsRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *CredentialsRequest) Reset()                    { *m = CredentialsRequest{} }
func (m *CredentialsRequest) String() string            { return proto.CompactTextString(m) }
func (*CredentialsRequest) ProtoMessage()               {}
func (*CredentialsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CredentialsRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// The response message containing the greetings
type AuthResponse struct {
	Token      string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	ValidUntil string `protobuf:"bytes,2,opt,name=validUntil" json:"validUntil,omitempty"`
}

func (m *AuthResponse) Reset()                    { *m = AuthResponse{} }
func (m *AuthResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()               {}
func (*AuthResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AuthResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthResponse) GetValidUntil() string {
	if m != nil {
		return m.ValidUntil
	}
	return ""
}

func init() {
	proto.RegisterType((*LogInRequest)(nil), "pb.LogInRequest")
	proto.RegisterType((*CredentialsRequest)(nil), "pb.CredentialsRequest")
	proto.RegisterType((*AuthResponse)(nil), "pb.AuthResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Authenticator service

type AuthenticatorClient interface {
	// Sends a greeting
	LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	Validate(ctx context.Context, in *CredentialsRequest, opts ...grpc.CallOption) (*AuthResponse, error)
}

type authenticatorClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticatorClient(cc *grpc.ClientConn) AuthenticatorClient {
	return &authenticatorClient{cc}
}

func (c *authenticatorClient) LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := grpc.Invoke(ctx, "/pb.Authenticator/LogIn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticatorClient) Validate(ctx context.Context, in *CredentialsRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := grpc.Invoke(ctx, "/pb.Authenticator/Validate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Authenticator service

type AuthenticatorServer interface {
	// Sends a greeting
	LogIn(context.Context, *LogInRequest) (*AuthResponse, error)
	Validate(context.Context, *CredentialsRequest) (*AuthResponse, error)
}

func RegisterAuthenticatorServer(s *grpc.Server, srv AuthenticatorServer) {
	s.RegisterService(&_Authenticator_serviceDesc, srv)
}

func _Authenticator_LogIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServer).LogIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Authenticator/LogIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServer).LogIn(ctx, req.(*LogInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenticator_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Authenticator/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServer).Validate(ctx, req.(*CredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authenticator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Authenticator",
	HandlerType: (*AuthenticatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogIn",
			Handler:    _Authenticator_LogIn_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _Authenticator_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4a, 0xf4, 0x40,
	0x10, 0x85, 0xc9, 0xc0, 0xfc, 0xcc, 0x14, 0xf9, 0x45, 0x9b, 0x41, 0x87, 0x20, 0x22, 0xbd, 0x92,
	0x2c, 0x12, 0xd4, 0xdd, 0xec, 0x06, 0x45, 0x10, 0x74, 0x13, 0xd0, 0x7d, 0xc7, 0x14, 0x99, 0xc6,
	0xd8, 0xd5, 0xa6, 0x2b, 0x71, 0xef, 0x15, 0x3c, 0x81, 0x67, 0xf2, 0x0a, 0x1e, 0x44, 0x3a, 0xc9,
	0x48, 0xc0, 0x59, 0xbe, 0xaa, 0x7a, 0xdf, 0xe3, 0x15, 0x80, 0x6a, 0x78, 0x93, 0xd8, 0x9a, 0x98,
	0xc4, 0xc4, 0xe6, 0xd1, 0x71, 0x49, 0x54, 0x56, 0x98, 0x2a, 0xab, 0x53, 0x65, 0x0c, 0xb1, 0x62,
	0x4d, 0xc6, 0xf5, 0x17, 0xf2, 0x06, 0xc2, 0x3b, 0x2a, 0x6f, 0x4d, 0x86, 0xaf, 0x0d, 0x3a, 0x16,
	0x11, 0xcc, 0x1a, 0x87, 0xb5, 0x51, 0x2f, 0xb8, 0x0c, 0x4e, 0x83, 0xb3, 0x79, 0xf6, 0xab, 0xfd,
	0xce, 0x2a, 0xe7, 0xde, 0xa8, 0x2e, 0x96, 0x93, 0x7e, 0xb7, 0xd5, 0x32, 0x06, 0x71, 0x55, 0x63,
	0x81, 0x86, 0xb5, 0xaa, 0xdc, 0x96, 0xb6, 0x80, 0x29, 0xd3, 0x33, 0x9a, 0x01, 0xd5, 0x0b, 0x79,
	0x0d, 0xe1, 0xba, 0xe1, 0x4d, 0x86, 0xce, 0x92, 0x71, 0xb8, 0xfb, 0x4a, 0x9c, 0x00, 0xb4, 0xaa,
	0xd2, 0xc5, 0x83, 0x61, 0x5d, 0x0d, 0x79, 0xa3, 0xc9, 0xc5, 0x67, 0x00, 0xff, 0x3d, 0xc6, 0x47,
	0x3e, 0x29, 0xa6, 0x5a, 0xac, 0x61, 0xda, 0x75, 0x11, 0xfb, 0x89, 0xcd, 0x93, 0x71, 0xad, 0xa8,
	0x9b, 0x8c, 0x43, 0xe5, 0xe2, 0xfd, 0xeb, 0xfb, 0x63, 0xb2, 0x27, 0xe7, 0x69, 0x7b, 0x9e, 0x56,
	0x54, 0x6a, 0xb3, 0x0a, 0x62, 0x71, 0x0f, 0xb3, 0x47, 0x1f, 0xa1, 0x18, 0xc5, 0xa1, 0xf7, 0xfc,
	0x2d, 0xb5, 0x83, 0x75, 0xd4, 0xb1, 0x0e, 0x64, 0xe8, 0x59, 0xed, 0xe0, 0x5f, 0x05, 0x71, 0xfe,
	0xaf, 0x7b, 0xf2, 0xe5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x49, 0xff, 0x76, 0xef, 0x94, 0x01,
	0x00, 0x00,
}
