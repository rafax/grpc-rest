// Code generated by protoc-gen-grpc-gateway
// source: auth.proto
// DO NOT EDIT!

/*
Package pb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package pb

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gengo/grpc-gateway/runtime"
	"github.com/gengo/grpc-gateway/utilities"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
)

var _ codes.Code
var _ io.Reader
var _ = runtime.String
var _ = json.Marshal
var _ = utilities.NewDoubleArray

func request_Authenticator_LogIn_0(ctx context.Context, client AuthenticatorClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq LogInRequest
	var metadata runtime.ServerMetadata

	if err := json.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.LogIn(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_Authenticator_Validate_0(ctx context.Context, client AuthenticatorClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CredentialsRequest
	var metadata runtime.ServerMetadata

	if err := json.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Validate(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterAuthenticatorHandlerFromEndpoint is same as RegisterAuthenticatorHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterAuthenticatorHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterAuthenticatorHandler(ctx, mux, conn)
}

// RegisterAuthenticatorHandler registers the http handlers for service Authenticator to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterAuthenticatorHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := NewAuthenticatorClient(conn)

	mux.Handle("POST", pattern_Authenticator_LogIn_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		resp, md, err := request_Authenticator_LogIn_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_Authenticator_LogIn_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_Authenticator_Validate_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		resp, md, err := request_Authenticator_Validate_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_Authenticator_Validate_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Authenticator_LogIn_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "login"}, ""))

	pattern_Authenticator_Validate_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "validate"}, ""))
)

var (
	forward_Authenticator_LogIn_0 = runtime.ForwardResponseMessage

	forward_Authenticator_Validate_0 = runtime.ForwardResponseMessage
)
