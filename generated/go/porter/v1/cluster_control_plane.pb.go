// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: porter/v1/cluster_control_plane.proto

package porterv1

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

var File_porter_v1_cluster_control_plane_proto protoreflect.FileDescriptor

var file_porter_v1_cluster_control_plane_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x1a, 0x13, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x77,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x77, 0x0a, 0x13, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x12, 0x60,
	0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x25, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x75, 0x6d, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1e,
	0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x75, 0x6d,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x00,
	0x42, 0xa3, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x42, 0x18, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x15, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x50, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_porter_v1_cluster_control_plane_proto_goTypes = []interface{}{
	(*CreateAssumeRoleChainInput)(nil), // 0: porter.v1.CreateAssumeRoleChainInput
	(*AssumeRoleChainLink)(nil),        // 1: porter.v1.AssumeRoleChainLink
}
var file_porter_v1_cluster_control_plane_proto_depIdxs = []int32{
	0, // 0: porter.v1.ClusterControlPlane.CreateAssumeRoleChain:input_type -> porter.v1.CreateAssumeRoleChainInput
	1, // 1: porter.v1.ClusterControlPlane.CreateAssumeRoleChain:output_type -> porter.v1.AssumeRoleChainLink
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_porter_v1_cluster_control_plane_proto_init() }
func file_porter_v1_cluster_control_plane_proto_init() {
	if File_porter_v1_cluster_control_plane_proto != nil {
		return
	}
	file_porter_v1_aws_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_porter_v1_cluster_control_plane_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_porter_v1_cluster_control_plane_proto_goTypes,
		DependencyIndexes: file_porter_v1_cluster_control_plane_proto_depIdxs,
	}.Build()
	File_porter_v1_cluster_control_plane_proto = out.File
	file_porter_v1_cluster_control_plane_proto_rawDesc = nil
	file_porter_v1_cluster_control_plane_proto_goTypes = nil
	file_porter_v1_cluster_control_plane_proto_depIdxs = nil
}
