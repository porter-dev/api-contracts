// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: porter/v1/auth_management.proto

package porterv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/porter-dev/api-contracts/generated/go/porter/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// AuthManagementServiceName is the fully-qualified name of the AuthManagementService service.
	AuthManagementServiceName = "porter.v1.AuthManagementService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AuthManagementServiceAPITokenProcedure is the fully-qualified name of the AuthManagementService's
	// APIToken RPC.
	AuthManagementServiceAPITokenProcedure = "/porter.v1.AuthManagementService/APIToken"
)

// AuthManagementServiceClient is a client for the porter.v1.AuthManagementService service.
type AuthManagementServiceClient interface {
	// APIToken gets a Porter token for programmatic access to the Porter API
	APIToken(context.Context, *connect.Request[v1.APITokenRequest]) (*connect.Response[v1.APITokenResponse], error)
}

// NewAuthManagementServiceClient constructs a client for the porter.v1.AuthManagementService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuthManagementServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AuthManagementServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &authManagementServiceClient{
		aPIToken: connect.NewClient[v1.APITokenRequest, v1.APITokenResponse](
			httpClient,
			baseURL+AuthManagementServiceAPITokenProcedure,
			opts...,
		),
	}
}

// authManagementServiceClient implements AuthManagementServiceClient.
type authManagementServiceClient struct {
	aPIToken *connect.Client[v1.APITokenRequest, v1.APITokenResponse]
}

// APIToken calls porter.v1.AuthManagementService.APIToken.
func (c *authManagementServiceClient) APIToken(ctx context.Context, req *connect.Request[v1.APITokenRequest]) (*connect.Response[v1.APITokenResponse], error) {
	return c.aPIToken.CallUnary(ctx, req)
}

// AuthManagementServiceHandler is an implementation of the porter.v1.AuthManagementService service.
type AuthManagementServiceHandler interface {
	// APIToken gets a Porter token for programmatic access to the Porter API
	APIToken(context.Context, *connect.Request[v1.APITokenRequest]) (*connect.Response[v1.APITokenResponse], error)
}

// NewAuthManagementServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuthManagementServiceHandler(svc AuthManagementServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	authManagementServiceAPITokenHandler := connect.NewUnaryHandler(
		AuthManagementServiceAPITokenProcedure,
		svc.APIToken,
		opts...,
	)
	return "/porter.v1.AuthManagementService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AuthManagementServiceAPITokenProcedure:
			authManagementServiceAPITokenHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAuthManagementServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAuthManagementServiceHandler struct{}

func (UnimplementedAuthManagementServiceHandler) APIToken(context.Context, *connect.Request[v1.APITokenRequest]) (*connect.Response[v1.APITokenResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("porter.v1.AuthManagementService.APIToken is not implemented"))
}
