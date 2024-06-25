// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: porter/v1/deployment_target.proto

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

// DeploymentTargetIdentifier is the object that identifies a deployment target. One of id or name must be provided, with id taking precedence.
type DeploymentTargetIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the id of the deployment target
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// name is the name of the deployment target
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeploymentTargetIdentifier) Reset() {
	*x = DeploymentTargetIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_deployment_target_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentTargetIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentTargetIdentifier) ProtoMessage() {}

func (x *DeploymentTargetIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_deployment_target_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentTargetIdentifier.ProtoReflect.Descriptor instead.
func (*DeploymentTargetIdentifier) Descriptor() ([]byte, []int) {
	return file_porter_v1_deployment_target_proto_rawDescGZIP(), []int{0}
}

func (x *DeploymentTargetIdentifier) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeploymentTargetIdentifier) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeploymentTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId int64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// name is the vanity name for the deployment target
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// namespace is the namespace that the deployment target points to
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// cluster_id is the id of the cluster that the deployment target points to
	ClusterId int64 `protobuf:"varint,4,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// is_preview indicates whether this is a preview deployment target or not
	IsPreview bool `protobuf:"varint,5,opt,name=is_preview,json=isPreview,proto3" json:"is_preview,omitempty"`
	// is_default indicates whether this is the default deployment target for the cluster
	IsDefault bool `protobuf:"varint,6,opt,name=is_default,json=isDefault,proto3" json:"is_default,omitempty"`
	// id is the id of the deployment target
	Id string `protobuf:"bytes,7,opt,name=id,proto3" json:"id,omitempty"`
	// metadata is the metadata for the deployment target, if any
	Metadata *DeploymentTargetMeta `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *DeploymentTarget) Reset() {
	*x = DeploymentTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_deployment_target_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentTarget) ProtoMessage() {}

func (x *DeploymentTarget) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_deployment_target_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentTarget.ProtoReflect.Descriptor instead.
func (*DeploymentTarget) Descriptor() ([]byte, []int) {
	return file_porter_v1_deployment_target_proto_rawDescGZIP(), []int{1}
}

func (x *DeploymentTarget) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *DeploymentTarget) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeploymentTarget) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DeploymentTarget) GetClusterId() int64 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

func (x *DeploymentTarget) GetIsPreview() bool {
	if x != nil {
		return x.IsPreview
	}
	return false
}

func (x *DeploymentTarget) GetIsDefault() bool {
	if x != nil {
		return x.IsDefault
	}
	return false
}

func (x *DeploymentTarget) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeploymentTarget) GetMetadata() *DeploymentTargetMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type DeploymentTargetMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pull_request is information about the pull request that triggered the deployment, if applicable
	PullRequest *PullRequest `protobuf:"bytes,1,opt,name=pull_request,json=pullRequest,proto3" json:"pull_request,omitempty"`
}

func (x *DeploymentTargetMeta) Reset() {
	*x = DeploymentTargetMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_deployment_target_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentTargetMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentTargetMeta) ProtoMessage() {}

func (x *DeploymentTargetMeta) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_deployment_target_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentTargetMeta.ProtoReflect.Descriptor instead.
func (*DeploymentTargetMeta) Descriptor() ([]byte, []int) {
	return file_porter_v1_deployment_target_proto_rawDescGZIP(), []int{2}
}

func (x *DeploymentTargetMeta) GetPullRequest() *PullRequest {
	if x != nil {
		return x.PullRequest
	}
	return nil
}

type PullRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// repository is the repository that the pull request is in
	Repository string `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	// number is the number of the pull request
	Number int64 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	// head_ref is the head ref of the pull request
	HeadRef string `protobuf:"bytes,3,opt,name=head_ref,json=headRef,proto3" json:"head_ref,omitempty"`
}

func (x *PullRequest) Reset() {
	*x = PullRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_deployment_target_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullRequest) ProtoMessage() {}

func (x *PullRequest) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_deployment_target_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullRequest.ProtoReflect.Descriptor instead.
func (*PullRequest) Descriptor() ([]byte, []int) {
	return file_porter_v1_deployment_target_proto_rawDescGZIP(), []int{3}
}

func (x *PullRequest) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *PullRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *PullRequest) GetHeadRef() string {
	if x != nil {
		return x.HeadRef
	}
	return ""
}

var File_porter_v1_deployment_target_proto protoreflect.FileDescriptor

var file_porter_v1_deployment_target_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x40,
	0x0a, 0x1a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x8d, 0x02, 0x0a, 0x10, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x51, 0x0a, 0x14, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x0c, 0x70, 0x75, 0x6c, 0x6c,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x70, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x60, 0x0a, 0x0b, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x65,
	0x61, 0x64, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x65,
	0x61, 0x64, 0x52, 0x65, 0x66, 0x42, 0xb0, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f,
	0x67, 0x6f, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x50, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x50, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_porter_v1_deployment_target_proto_rawDescOnce sync.Once
	file_porter_v1_deployment_target_proto_rawDescData = file_porter_v1_deployment_target_proto_rawDesc
)

func file_porter_v1_deployment_target_proto_rawDescGZIP() []byte {
	file_porter_v1_deployment_target_proto_rawDescOnce.Do(func() {
		file_porter_v1_deployment_target_proto_rawDescData = protoimpl.X.CompressGZIP(file_porter_v1_deployment_target_proto_rawDescData)
	})
	return file_porter_v1_deployment_target_proto_rawDescData
}

var file_porter_v1_deployment_target_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_porter_v1_deployment_target_proto_goTypes = []any{
	(*DeploymentTargetIdentifier)(nil), // 0: porter.v1.DeploymentTargetIdentifier
	(*DeploymentTarget)(nil),           // 1: porter.v1.DeploymentTarget
	(*DeploymentTargetMeta)(nil),       // 2: porter.v1.DeploymentTargetMeta
	(*PullRequest)(nil),                // 3: porter.v1.PullRequest
}
var file_porter_v1_deployment_target_proto_depIdxs = []int32{
	2, // 0: porter.v1.DeploymentTarget.metadata:type_name -> porter.v1.DeploymentTargetMeta
	3, // 1: porter.v1.DeploymentTargetMeta.pull_request:type_name -> porter.v1.PullRequest
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_porter_v1_deployment_target_proto_init() }
func file_porter_v1_deployment_target_proto_init() {
	if File_porter_v1_deployment_target_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_porter_v1_deployment_target_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DeploymentTargetIdentifier); i {
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
		file_porter_v1_deployment_target_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DeploymentTarget); i {
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
		file_porter_v1_deployment_target_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DeploymentTargetMeta); i {
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
		file_porter_v1_deployment_target_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PullRequest); i {
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
			RawDescriptor: file_porter_v1_deployment_target_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_porter_v1_deployment_target_proto_goTypes,
		DependencyIndexes: file_porter_v1_deployment_target_proto_depIdxs,
		MessageInfos:      file_porter_v1_deployment_target_proto_msgTypes,
	}.Build()
	File_porter_v1_deployment_target_proto = out.File
	file_porter_v1_deployment_target_proto_rawDesc = nil
	file_porter_v1_deployment_target_proto_goTypes = nil
	file_porter_v1_deployment_target_proto_depIdxs = nil
}
