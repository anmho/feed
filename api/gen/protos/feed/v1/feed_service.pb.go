// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: protos/feed/v1/feed_service.proto

package feedv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_protos_feed_v1_feed_service_proto protoreflect.FileDescriptor

var file_protos_feed_v1_feed_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x76, 0x31,
	0x2f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x65, 0x65, 0x64,
	0x2e, 0x76, 0x31, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x65, 0x65, 0x64,
	0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8c, 0x02, 0x0a, 0x0b, 0x46, 0x65, 0x65, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x08, 0x53, 0x79, 0x6e, 0x63, 0x46,
	0x65, 0x65, 0x64, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x65, 0x65,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x65,
	0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x66, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x55, 0x0a, 0x0a, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x21, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x76,
	0x31, 0x3b, 0x66, 0x65, 0x65, 0x64, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_protos_feed_v1_feed_service_proto_goTypes = []any{
	(*SyncFeedRequest)(nil),    // 0: protos.feed.v1.SyncFeedRequest
	(*CreatePostRequest)(nil),  // 1: protos.feed.v1.CreatePostRequest
	(*FollowUserRequest)(nil),  // 2: protos.feed.v1.FollowUserRequest
	(*SyncFeedResponse)(nil),   // 3: protos.feed.v1.SyncFeedResponse
	(*CreatePostResponse)(nil), // 4: protos.feed.v1.CreatePostResponse
	(*FollowUserResponse)(nil), // 5: protos.feed.v1.FollowUserResponse
}
var file_protos_feed_v1_feed_service_proto_depIdxs = []int32{
	0, // 0: protos.feed.v1.FeedService.SyncFeed:input_type -> protos.feed.v1.SyncFeedRequest
	1, // 1: protos.feed.v1.FeedService.CreatePost:input_type -> protos.feed.v1.CreatePostRequest
	2, // 2: protos.feed.v1.FeedService.FollowUser:input_type -> protos.feed.v1.FollowUserRequest
	3, // 3: protos.feed.v1.FeedService.SyncFeed:output_type -> protos.feed.v1.SyncFeedResponse
	4, // 4: protos.feed.v1.FeedService.CreatePost:output_type -> protos.feed.v1.CreatePostResponse
	5, // 5: protos.feed.v1.FeedService.FollowUser:output_type -> protos.feed.v1.FollowUserResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_feed_v1_feed_service_proto_init() }
func file_protos_feed_v1_feed_service_proto_init() {
	if File_protos_feed_v1_feed_service_proto != nil {
		return
	}
	file_protos_feed_v1_feed_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_feed_v1_feed_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_feed_v1_feed_service_proto_goTypes,
		DependencyIndexes: file_protos_feed_v1_feed_service_proto_depIdxs,
	}.Build()
	File_protos_feed_v1_feed_service_proto = out.File
	file_protos_feed_v1_feed_service_proto_rawDesc = nil
	file_protos_feed_v1_feed_service_proto_goTypes = nil
	file_protos_feed_v1_feed_service_proto_depIdxs = nil
}
