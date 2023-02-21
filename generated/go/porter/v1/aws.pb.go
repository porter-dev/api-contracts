// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: porter/v1/aws.proto

package porterv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AssumeRoleChainLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId  int64  `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	SourceArn  string `protobuf:"bytes,2,opt,name=source_arn,json=sourceArn,proto3" json:"source_arn,omitempty"`
	TargetArn  string `protobuf:"bytes,3,opt,name=target_arn,json=targetArn,proto3" json:"target_arn,omitempty"`
	ExternalId string `protobuf:"bytes,4,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
}

func (x *AssumeRoleChainLink) Reset() {
	*x = AssumeRoleChainLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_aws_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssumeRoleChainLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssumeRoleChainLink) ProtoMessage() {}

func (x *AssumeRoleChainLink) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_aws_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssumeRoleChainLink.ProtoReflect.Descriptor instead.
func (*AssumeRoleChainLink) Descriptor() ([]byte, []int) {
	return file_porter_v1_aws_proto_rawDescGZIP(), []int{0}
}

func (x *AssumeRoleChainLink) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *AssumeRoleChainLink) GetSourceArn() string {
	if x != nil {
		return x.SourceArn
	}
	return ""
}

func (x *AssumeRoleChainLink) GetTargetArn() string {
	if x != nil {
		return x.TargetArn
	}
	return ""
}

func (x *AssumeRoleChainLink) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

type CreateAssumeRoleChainInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId          int64  `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	SourceArn          string `protobuf:"bytes,2,opt,name=source_arn,json=sourceArn,proto3" json:"source_arn,omitempty"`
	TargetAccessId     string `protobuf:"bytes,3,opt,name=target_access_id,json=targetAccessId,proto3" json:"target_access_id,omitempty"`
	TargetSecretKey    string `protobuf:"bytes,4,opt,name=target_secret_key,json=targetSecretKey,proto3" json:"target_secret_key,omitempty"`
	TargetSessionToken string `protobuf:"bytes,5,opt,name=target_session_token,json=targetSessionToken,proto3" json:"target_session_token,omitempty"`
}

func (x *CreateAssumeRoleChainInput) Reset() {
	*x = CreateAssumeRoleChainInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_aws_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAssumeRoleChainInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAssumeRoleChainInput) ProtoMessage() {}

func (x *CreateAssumeRoleChainInput) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_aws_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAssumeRoleChainInput.ProtoReflect.Descriptor instead.
func (*CreateAssumeRoleChainInput) Descriptor() ([]byte, []int) {
	return file_porter_v1_aws_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAssumeRoleChainInput) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *CreateAssumeRoleChainInput) GetSourceArn() string {
	if x != nil {
		return x.SourceArn
	}
	return ""
}

func (x *CreateAssumeRoleChainInput) GetTargetAccessId() string {
	if x != nil {
		return x.TargetAccessId
	}
	return ""
}

func (x *CreateAssumeRoleChainInput) GetTargetSecretKey() string {
	if x != nil {
		return x.TargetSecretKey
	}
	return ""
}

func (x *CreateAssumeRoleChainInput) GetTargetSessionToken() string {
	if x != nil {
		return x.TargetSessionToken
	}
	return ""
}

var File_porter_v1_aws_proto protoreflect.FileDescriptor

var file_porter_v1_aws_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x77, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x22, 0x93, 0x01, 0x0a, 0x13, 0x41, 0x73, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x61, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x41, 0x72, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x61, 0x72, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x41, 0x72, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x22, 0xe2, 0x01, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x73, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61,
	0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x41, 0x72, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x2a, 0x0a,
	0x11, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x93, 0x01, 0x0a, 0x0d,
	0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x41,
	0x77, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2d, 0x64, 0x65, 0x76,
	0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x50,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_porter_v1_aws_proto_rawDescOnce sync.Once
	file_porter_v1_aws_proto_rawDescData = file_porter_v1_aws_proto_rawDesc
)

func file_porter_v1_aws_proto_rawDescGZIP() []byte {
	file_porter_v1_aws_proto_rawDescOnce.Do(func() {
		file_porter_v1_aws_proto_rawDescData = protoimpl.X.CompressGZIP(file_porter_v1_aws_proto_rawDescData)
	})
	return file_porter_v1_aws_proto_rawDescData
}

var file_porter_v1_aws_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_porter_v1_aws_proto_goTypes = []interface{}{
	(*AssumeRoleChainLink)(nil),        // 0: porter.v1.AssumeRoleChainLink
	(*CreateAssumeRoleChainInput)(nil), // 1: porter.v1.CreateAssumeRoleChainInput
}
var file_porter_v1_aws_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_porter_v1_aws_proto_init() }
func file_porter_v1_aws_proto_init() {
	if File_porter_v1_aws_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_porter_v1_aws_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssumeRoleChainLink); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_porter_v1_aws_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAssumeRoleChainInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_porter_v1_aws_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_porter_v1_aws_proto_goTypes,
		DependencyIndexes: file_porter_v1_aws_proto_depIdxs,
		MessageInfos:      file_porter_v1_aws_proto_msgTypes,
	}.Build()
	File_porter_v1_aws_proto = out.File
	file_porter_v1_aws_proto_rawDesc = nil
	file_porter_v1_aws_proto_goTypes = nil
	file_porter_v1_aws_proto_depIdxs = nil
}
