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
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"

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
const _ = proto.ProtoPackageIsVersion1

// The request message containing the user's name.
type LogInRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LogInRequest) Reset()                    { *m = LogInRequest{} }
func (m *LogInRequest) String() string            { return proto.CompactTextString(m) }
func (*LogInRequest) ProtoMessage()               {}
func (*LogInRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CredentialsRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *CredentialsRequest) Reset()                    { *m = CredentialsRequest{} }
func (m *CredentialsRequest) String() string            { return proto.CompactTextString(m) }
func (*CredentialsRequest) ProtoMessage()               {}
func (*CredentialsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// The response message containing the greetings
type AuthResponse struct {
	Token      string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	ValidUntil string `protobuf:"bytes,2,opt,name=validUntil" json:"validUntil,omitempty"`
}

func (m *AuthResponse) Reset()                    { *m = AuthResponse{} }
func (m *AuthResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()               {}
func (*AuthResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*LogInRequest)(nil), "pb.LogInRequest")
	proto.RegisterType((*CredentialsRequest)(nil), "pb.CredentialsRequest")
	proto.RegisterType((*AuthResponse)(nil), "pb.AuthResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

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

func _Authenticator_LogIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(LogInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AuthenticatorServer).LogIn(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Authenticator_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AuthenticatorServer).Validate(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
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
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x4d, 0xa0, 0xd2, 0x0e, 0x55, 0x64, 0x29, 0x52, 0x82, 0x88, 0xe4, 0x24, 0x3d, 0x24,
	0xa8, 0xe0, 0xc1, 0x5b, 0x51, 0x04, 0xc1, 0x53, 0x40, 0xef, 0x13, 0x33, 0xc4, 0xc5, 0x75, 0x67,
	0xdd, 0xdd, 0xd4, 0xbb, 0x47, 0xaf, 0xfe, 0x34, 0xff, 0x82, 0x3f, 0xc4, 0x6c, 0xd2, 0x4a, 0xc0,
	0x1e, 0xdf, 0xcc, 0xec, 0xf7, 0xde, 0x5b, 0x00, 0x6c, 0xfc, 0x73, 0x66, 0x2c, 0x7b, 0x16, 0xb1,
	0x29, 0x93, 0xa3, 0x9a, 0xb9, 0x56, 0x94, 0xa3, 0x91, 0x39, 0x6a, 0xcd, 0x1e, 0xbd, 0x64, 0xed,
	0xfa, 0x8b, 0xf4, 0x16, 0xa6, 0xf7, 0x5c, 0xdf, 0xe9, 0x82, 0xde, 0x1a, 0x72, 0x5e, 0x24, 0x30,
	0x6e, 0x1c, 0x59, 0x8d, 0xaf, 0x34, 0x8f, 0x4e, 0xa2, 0xd3, 0x49, 0xf1, 0xa7, 0xc3, 0xce, 0xa0,
	0x73, 0xef, 0x6c, 0xab, 0x79, 0xdc, 0xef, 0x36, 0x3a, 0x5d, 0x80, 0xb8, 0xb6, 0x54, 0x91, 0xf6,
	0x12, 0x95, 0xdb, 0xd0, 0x66, 0x30, 0xf2, 0xfc, 0x42, 0x7a, 0x8d, 0xea, 0x45, 0x7a, 0x03, 0xd3,
	0x65, 0x9b, 0xb1, 0x20, 0x67, 0xda, 0x20, 0xb4, 0xfd, 0x4a, 0x1c, 0x03, 0xac, 0x50, 0xc9, 0xea,
	0xa1, 0x45, 0xaa, 0xb5, 0xdf, 0x60, 0x72, 0xfe, 0x19, 0xc1, 0x5e, 0xc0, 0x04, 0xcb, 0x27, 0xf4,
	0x6c, 0xc5, 0x12, 0x46, 0x5d, 0x17, 0x71, 0x90, 0x99, 0x32, 0x1b, 0xd6, 0x4a, 0xba, 0xc9, 0xd0,
	0x34, 0x9d, 0x7d, 0x7c, 0xff, 0x7c, 0xc5, 0xfb, 0xe9, 0x24, 0x5f, 0x9d, 0xe5, 0x8a, 0x6b, 0xa9,
	0xaf, 0xa2, 0x85, 0xb8, 0x84, 0xf1, 0x63, 0xb0, 0x40, 0x4f, 0xe2, 0x30, 0xbc, 0xf9, 0x5f, 0x6a,
	0x0b, 0x6b, 0xa7, 0xdc, 0xed, 0x7e, 0xf3, 0xe2, 0x37, 0x00, 0x00, 0xff, 0xff, 0xae, 0xf7, 0xa9,
	0x4c, 0x7d, 0x01, 0x00, 0x00,
}
