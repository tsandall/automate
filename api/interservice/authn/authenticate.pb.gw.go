// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: api/interservice/authn/authenticate.proto

/*
Package authn is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package authn

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_Authentication_Authenticate_0(ctx context.Context, marshaler runtime.Marshaler, client AuthenticationClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AuthenticateRequest
	var metadata runtime.ServerMetadata

	msg, err := client.Authenticate(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Authentication_Authenticate_0(ctx context.Context, marshaler runtime.Marshaler, server AuthenticationServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AuthenticateRequest
	var metadata runtime.ServerMetadata

	msg, err := server.Authenticate(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterAuthenticationHandlerServer registers the http handlers for service Authentication to "mux".
// UnaryRPC     :call AuthenticationServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
func RegisterAuthenticationHandlerServer(ctx context.Context, mux *runtime.ServeMux, server AuthenticationServer, opts []grpc.DialOption) error {

	mux.Handle("GET", pattern_Authentication_Authenticate_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Authentication_Authenticate_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Authentication_Authenticate_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterAuthenticationHandlerFromEndpoint is same as RegisterAuthenticationHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterAuthenticationHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterAuthenticationHandler(ctx, mux, conn)
}

// RegisterAuthenticationHandler registers the http handlers for service Authentication to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterAuthenticationHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterAuthenticationHandlerClient(ctx, mux, NewAuthenticationClient(conn))
}

// RegisterAuthenticationHandlerClient registers the http handlers for service Authentication
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "AuthenticationClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "AuthenticationClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "AuthenticationClient" to call the correct interceptors.
func RegisterAuthenticationHandlerClient(ctx context.Context, mux *runtime.ServeMux, client AuthenticationClient) error {

	mux.Handle("GET", pattern_Authentication_Authenticate_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Authentication_Authenticate_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Authentication_Authenticate_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Authentication_Authenticate_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"api", "authenticate"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_Authentication_Authenticate_0 = runtime.ForwardResponseMessage
)
