// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: porter/v1/user_management.proto

package porterv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PorterAgentConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId int64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ClusterId int64 `protobuf:"varint,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
}

func (x *PorterAgentConnectionRequest) Reset() {
	*x = PorterAgentConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_user_management_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PorterAgentConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PorterAgentConnectionRequest) ProtoMessage() {}

func (x *PorterAgentConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_user_management_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PorterAgentConnectionRequest.ProtoReflect.Descriptor instead.
func (*PorterAgentConnectionRequest) Descriptor() ([]byte, []int) {
	return file_porter_v1_user_management_proto_rawDescGZIP(), []int{0}
}

func (x *PorterAgentConnectionRequest) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *PorterAgentConnectionRequest) GetClusterId() int64 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

type PorterAgentConnectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerUrl string `protobuf:"bytes,1,opt,name=server_url,json=serverUrl,proto3" json:"server_url,omitempty"`
	Port      int64  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Token     string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *PorterAgentConnectionResponse) Reset() {
	*x = PorterAgentConnectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_user_management_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PorterAgentConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PorterAgentConnectionResponse) ProtoMessage() {}

func (x *PorterAgentConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_user_management_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PorterAgentConnectionResponse.ProtoReflect.Descriptor instead.
func (*PorterAgentConnectionResponse) Descriptor() ([]byte, []int) {
	return file_porter_v1_user_management_proto_rawDescGZIP(), []int{1}
}

func (x *PorterAgentConnectionResponse) GetServerUrl() string {
	if x != nil {
		return x.ServerUrl
	}
	return ""
}

func (x *PorterAgentConnectionResponse) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *PorterAgentConnectionResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_porter_v1_user_management_proto protoreflect.FileDescriptor

var file_porter_v1_user_management_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a,
	0x1c, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x68, 0x0a, 0x1d, 0x50,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x7e, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x6c, 0x0a, 0x15, 0x50, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x27, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xae, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f,
	0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x15, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x50, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_porter_v1_user_management_proto_rawDescOnce sync.Once
	file_porter_v1_user_management_proto_rawDescData = file_porter_v1_user_management_proto_rawDesc
)

func file_porter_v1_user_management_proto_rawDescGZIP() []byte {
	file_porter_v1_user_management_proto_rawDescOnce.Do(func() {
		file_porter_v1_user_management_proto_rawDescData = protoimpl.X.CompressGZIP(file_porter_v1_user_management_proto_rawDescData)
	})
	return file_porter_v1_user_management_proto_rawDescData
}

var file_porter_v1_user_management_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_porter_v1_user_management_proto_goTypes = []interface{}{
	(*PorterAgentConnectionRequest)(nil),  // 0: porter.v1.PorterAgentConnectionRequest
	(*PorterAgentConnectionResponse)(nil), // 1: porter.v1.PorterAgentConnectionResponse
}
var file_porter_v1_user_management_proto_depIdxs = []int32{
	0, // 0: porter.v1.AuthManagement.PorterAgentConnection:input_type -> porter.v1.PorterAgentConnectionRequest
	1, // 1: porter.v1.AuthManagement.PorterAgentConnection:output_type -> porter.v1.PorterAgentConnectionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_porter_v1_user_management_proto_init() }
func file_porter_v1_user_management_proto_init() {
	if File_porter_v1_user_management_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_porter_v1_user_management_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PorterAgentConnectionRequest); i {
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
		file_porter_v1_user_management_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PorterAgentConnectionResponse); i {
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
			RawDescriptor: file_porter_v1_user_management_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_porter_v1_user_management_proto_goTypes,
		DependencyIndexes: file_porter_v1_user_management_proto_depIdxs,
		MessageInfos:      file_porter_v1_user_management_proto_msgTypes,
	}.Build()
	File_porter_v1_user_management_proto = out.File
	file_porter_v1_user_management_proto_rawDesc = nil
	file_porter_v1_user_management_proto_goTypes = nil
	file_porter_v1_user_management_proto_depIdxs = nil
}
