// Code generated by protoc-gen-go.
// source: auth.proto
// DO NOT EDIT!

/*
Package grpc_rest is a generated protocol buffer package.

It is generated from these files:
	auth.proto

It has these top-level messages:
	LogInRequest
	CredentialsRequest
	AuthResponse
*/
package grpc_rest

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
	proto.RegisterType((*LogInRequest)(nil), "grpc_rest.LogInRequest")
	proto.RegisterType((*CredentialsRequest)(nil), "grpc_rest.CredentialsRequest")
	proto.RegisterType((*AuthResponse)(nil), "grpc_rest.AuthResponse")
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
	err := grpc.Invoke(ctx, "/grpc_rest.Authenticator/LogIn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticatorClient) Validate(ctx context.Context, in *CredentialsRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := grpc.Invoke(ctx, "/grpc_rest.Authenticator/Validate", in, out, c.cc, opts...)
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
	ServiceName: "grpc_rest.Authenticator",
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
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0x2f, 0x2a, 0x48, 0x8e, 0x2f, 0x4a, 0x2d,
	0x2e, 0x51, 0x72, 0xe3, 0xe2, 0xf1, 0xc9, 0x4f, 0xf7, 0xcc, 0x0b, 0x4a, 0x2d, 0x2c, 0x05, 0xf2,
	0x85, 0xa4, 0xb8, 0x38, 0x4a, 0x8b, 0x53, 0x8b, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0xe0, 0x7c, 0x90, 0x5c, 0x41, 0x62, 0x71, 0x71, 0x79, 0x7e, 0x51, 0x8a, 0x04,
	0x13, 0x44, 0x0e, 0xc6, 0x57, 0xd2, 0xe2, 0x12, 0x72, 0x2e, 0x4a, 0x4d, 0x49, 0xcd, 0x2b, 0xc9,
	0x4c, 0xcc, 0x29, 0x86, 0x99, 0x26, 0xc2, 0xc5, 0x5a, 0x92, 0x9f, 0x9d, 0x9a, 0x07, 0x35, 0x0a,
	0xc2, 0x51, 0x72, 0xe1, 0xe2, 0x71, 0x04, 0x3a, 0x26, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38,
	0x15, 0xbb, 0x2a, 0x21, 0x39, 0x2e, 0xae, 0xb2, 0xc4, 0x9c, 0xcc, 0x94, 0x50, 0xa0, 0x91, 0x39,
	0x50, 0xfb, 0x90, 0x44, 0x8c, 0x26, 0x31, 0x72, 0xf1, 0x82, 0x8c, 0x01, 0x59, 0x99, 0x9c, 0x58,
	0x92, 0x5f, 0x24, 0x64, 0xcd, 0xc5, 0x0a, 0xf6, 0x8b, 0x90, 0xb8, 0x1e, 0xdc, 0x83, 0x7a, 0xc8,
	0xbe, 0x93, 0x42, 0x96, 0x40, 0x76, 0x82, 0x12, 0x83, 0x90, 0x0b, 0x17, 0x47, 0x18, 0xc8, 0xf0,
	0xc4, 0x92, 0x54, 0x21, 0x59, 0x24, 0x65, 0x98, 0xbe, 0xc2, 0x63, 0x4a, 0x12, 0x1b, 0x38, 0x80,
	0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x76, 0x01, 0xd0, 0x6e, 0x01, 0x00, 0x00,
}
