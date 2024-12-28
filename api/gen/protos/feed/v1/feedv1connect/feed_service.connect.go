// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: protos/feed/v1/feed_service.proto

package feedv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "feed/gen/protos/feed/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// FeedServiceName is the fully-qualified name of the FeedService service.
	FeedServiceName = "protos.feed.v1.FeedService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// FeedServiceSyncFeedProcedure is the fully-qualified name of the FeedService's SyncFeed RPC.
	FeedServiceSyncFeedProcedure = "/protos.feed.v1.FeedService/SyncFeed"
	// FeedServiceCreatePostProcedure is the fully-qualified name of the FeedService's CreatePost RPC.
	FeedServiceCreatePostProcedure = "/protos.feed.v1.FeedService/CreatePost"
	// FeedServiceFollowUserProcedure is the fully-qualified name of the FeedService's FollowUser RPC.
	FeedServiceFollowUserProcedure = "/protos.feed.v1.FeedService/FollowUser"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	feedServiceServiceDescriptor          = v1.File_protos_feed_v1_feed_service_proto.Services().ByName("FeedService")
	feedServiceSyncFeedMethodDescriptor   = feedServiceServiceDescriptor.Methods().ByName("SyncFeed")
	feedServiceCreatePostMethodDescriptor = feedServiceServiceDescriptor.Methods().ByName("CreatePost")
	feedServiceFollowUserMethodDescriptor = feedServiceServiceDescriptor.Methods().ByName("FollowUser")
)

// FeedServiceClient is a client for the protos.feed.v1.FeedService service.
type FeedServiceClient interface {
	SyncFeed(context.Context, *connect.Request[v1.SyncFeedRequest]) (*connect.Response[v1.SyncFeedResponse], error)
	CreatePost(context.Context, *connect.Request[v1.CreatePostRequest]) (*connect.Response[v1.CreatePostResponse], error)
	FollowUser(context.Context, *connect.Request[v1.FollowUserRequest]) (*connect.Response[v1.FollowUserResponse], error)
}

// NewFeedServiceClient constructs a client for the protos.feed.v1.FeedService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewFeedServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) FeedServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &feedServiceClient{
		syncFeed: connect.NewClient[v1.SyncFeedRequest, v1.SyncFeedResponse](
			httpClient,
			baseURL+FeedServiceSyncFeedProcedure,
			connect.WithSchema(feedServiceSyncFeedMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createPost: connect.NewClient[v1.CreatePostRequest, v1.CreatePostResponse](
			httpClient,
			baseURL+FeedServiceCreatePostProcedure,
			connect.WithSchema(feedServiceCreatePostMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		followUser: connect.NewClient[v1.FollowUserRequest, v1.FollowUserResponse](
			httpClient,
			baseURL+FeedServiceFollowUserProcedure,
			connect.WithSchema(feedServiceFollowUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// feedServiceClient implements FeedServiceClient.
type feedServiceClient struct {
	syncFeed   *connect.Client[v1.SyncFeedRequest, v1.SyncFeedResponse]
	createPost *connect.Client[v1.CreatePostRequest, v1.CreatePostResponse]
	followUser *connect.Client[v1.FollowUserRequest, v1.FollowUserResponse]
}

// SyncFeed calls protos.feed.v1.FeedService.SyncFeed.
func (c *feedServiceClient) SyncFeed(ctx context.Context, req *connect.Request[v1.SyncFeedRequest]) (*connect.Response[v1.SyncFeedResponse], error) {
	return c.syncFeed.CallUnary(ctx, req)
}

// CreatePost calls protos.feed.v1.FeedService.CreatePost.
func (c *feedServiceClient) CreatePost(ctx context.Context, req *connect.Request[v1.CreatePostRequest]) (*connect.Response[v1.CreatePostResponse], error) {
	return c.createPost.CallUnary(ctx, req)
}

// FollowUser calls protos.feed.v1.FeedService.FollowUser.
func (c *feedServiceClient) FollowUser(ctx context.Context, req *connect.Request[v1.FollowUserRequest]) (*connect.Response[v1.FollowUserResponse], error) {
	return c.followUser.CallUnary(ctx, req)
}

// FeedServiceHandler is an implementation of the protos.feed.v1.FeedService service.
type FeedServiceHandler interface {
	SyncFeed(context.Context, *connect.Request[v1.SyncFeedRequest]) (*connect.Response[v1.SyncFeedResponse], error)
	CreatePost(context.Context, *connect.Request[v1.CreatePostRequest]) (*connect.Response[v1.CreatePostResponse], error)
	FollowUser(context.Context, *connect.Request[v1.FollowUserRequest]) (*connect.Response[v1.FollowUserResponse], error)
}

// NewFeedServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewFeedServiceHandler(svc FeedServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	feedServiceSyncFeedHandler := connect.NewUnaryHandler(
		FeedServiceSyncFeedProcedure,
		svc.SyncFeed,
		connect.WithSchema(feedServiceSyncFeedMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	feedServiceCreatePostHandler := connect.NewUnaryHandler(
		FeedServiceCreatePostProcedure,
		svc.CreatePost,
		connect.WithSchema(feedServiceCreatePostMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	feedServiceFollowUserHandler := connect.NewUnaryHandler(
		FeedServiceFollowUserProcedure,
		svc.FollowUser,
		connect.WithSchema(feedServiceFollowUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/protos.feed.v1.FeedService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case FeedServiceSyncFeedProcedure:
			feedServiceSyncFeedHandler.ServeHTTP(w, r)
		case FeedServiceCreatePostProcedure:
			feedServiceCreatePostHandler.ServeHTTP(w, r)
		case FeedServiceFollowUserProcedure:
			feedServiceFollowUserHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedFeedServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedFeedServiceHandler struct{}

func (UnimplementedFeedServiceHandler) SyncFeed(context.Context, *connect.Request[v1.SyncFeedRequest]) (*connect.Response[v1.SyncFeedResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("protos.feed.v1.FeedService.SyncFeed is not implemented"))
}

func (UnimplementedFeedServiceHandler) CreatePost(context.Context, *connect.Request[v1.CreatePostRequest]) (*connect.Response[v1.CreatePostResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("protos.feed.v1.FeedService.CreatePost is not implemented"))
}

func (UnimplementedFeedServiceHandler) FollowUser(context.Context, *connect.Request[v1.FollowUserRequest]) (*connect.Response[v1.FollowUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("protos.feed.v1.FeedService.FollowUser is not implemented"))
}
