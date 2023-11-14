// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: porter/v1/cluster.proto

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

type EnumKubernetesKind int32

const (
	EnumKubernetesKind_ENUM_KUBERNETES_KIND_UNSPECIFIED EnumKubernetesKind = 0
	EnumKubernetesKind_ENUM_KUBERNETES_KIND_EKS         EnumKubernetesKind = 1
	EnumKubernetesKind_ENUM_KUBERNETES_KIND_GKE         EnumKubernetesKind = 2
	EnumKubernetesKind_ENUM_KUBERNETES_KIND_AKS         EnumKubernetesKind = 3
)

// Enum value maps for EnumKubernetesKind.
var (
	EnumKubernetesKind_name = map[int32]string{
		0: "ENUM_KUBERNETES_KIND_UNSPECIFIED",
		1: "ENUM_KUBERNETES_KIND_EKS",
		2: "ENUM_KUBERNETES_KIND_GKE",
		3: "ENUM_KUBERNETES_KIND_AKS",
	}
	EnumKubernetesKind_value = map[string]int32{
		"ENUM_KUBERNETES_KIND_UNSPECIFIED": 0,
		"ENUM_KUBERNETES_KIND_EKS":         1,
		"ENUM_KUBERNETES_KIND_GKE":         2,
		"ENUM_KUBERNETES_KIND_AKS":         3,
	}
)

func (x EnumKubernetesKind) Enum() *EnumKubernetesKind {
	p := new(EnumKubernetesKind)
	*p = x
	return p
}

func (x EnumKubernetesKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumKubernetesKind) Descriptor() protoreflect.EnumDescriptor {
	return file_porter_v1_cluster_proto_enumTypes[0].Descriptor()
}

func (EnumKubernetesKind) Type() protoreflect.EnumType {
	return &file_porter_v1_cluster_proto_enumTypes[0]
}

func (x EnumKubernetesKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumKubernetesKind.Descriptor instead.
func (EnumKubernetesKind) EnumDescriptor() ([]byte, []int) {
	return file_porter_v1_cluster_proto_rawDescGZIP(), []int{0}
}

type EnumCloudProvider int32

const (
	EnumCloudProvider_ENUM_CLOUD_PROVIDER_UNSPECIFIED EnumCloudProvider = 0
	EnumCloudProvider_ENUM_CLOUD_PROVIDER_AWS         EnumCloudProvider = 1
	EnumCloudProvider_ENUM_CLOUD_PROVIDER_GCP         EnumCloudProvider = 2
	EnumCloudProvider_ENUM_CLOUD_PROVIDER_AZURE       EnumCloudProvider = 3
)

// Enum value maps for EnumCloudProvider.
var (
	EnumCloudProvider_name = map[int32]string{
		0: "ENUM_CLOUD_PROVIDER_UNSPECIFIED",
		1: "ENUM_CLOUD_PROVIDER_AWS",
		2: "ENUM_CLOUD_PROVIDER_GCP",
		3: "ENUM_CLOUD_PROVIDER_AZURE",
	}
	EnumCloudProvider_value = map[string]int32{
		"ENUM_CLOUD_PROVIDER_UNSPECIFIED": 0,
		"ENUM_CLOUD_PROVIDER_AWS":         1,
		"ENUM_CLOUD_PROVIDER_GCP":         2,
		"ENUM_CLOUD_PROVIDER_AZURE":       3,
	}
)

func (x EnumCloudProvider) Enum() *EnumCloudProvider {
	p := new(EnumCloudProvider)
	*p = x
	return p
}

func (x EnumCloudProvider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumCloudProvider) Descriptor() protoreflect.EnumDescriptor {
	return file_porter_v1_cluster_proto_enumTypes[1].Descriptor()
}

func (EnumCloudProvider) Type() protoreflect.EnumType {
	return &file_porter_v1_cluster_proto_enumTypes[1]
}

func (x EnumCloudProvider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumCloudProvider.Descriptor instead.
func (EnumCloudProvider) EnumDescriptor() ([]byte, []int) {
	return file_porter_v1_cluster_proto_rawDescGZIP(), []int{1}
}

type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// project_id [REQUIRED] represents the Porter project that the cluster will be joined to. This is required for all cluster creations and updates
	ProjectId int32 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// cluster_id [OPTIONAL] represents the Porter cluster. This is required for update operations, but should be left blank when creating a cluster
	ClusterId int32 `protobuf:"varint,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// kind [REQUIRED] is the different types of supported kubernetes distros such as EKS, GKE etc.
	Kind EnumKubernetesKind `protobuf:"varint,3,opt,name=kind,proto3,enum=porter.v1.EnumKubernetesKind" json:"kind,omitempty"`
	// cloud_provider [REQUIRED] represents the provider that we will provisioning the cluster in
	CloudProvider EnumCloudProvider `protobuf:"varint,4,opt,name=cloud_provider,json=cloudProvider,proto3,enum=porter.v1.EnumCloudProvider" json:"cloud_provider,omitempty"`
	// cloud_provider_credentials_id [REQUIRED] is the Porter credentials that will be used for provisioning a cluster.
	// These must be stored within Porter, prior to cluster creation. For AWS this refers to the last link in an assume role chain
	CloudProviderCredentialsId string `protobuf:"bytes,5,opt,name=cloud_provider_credentials_id,json=cloudProviderCredentialsId,proto3" json:"cloud_provider_credentials_id,omitempty"`
	// kind_values are the required values, depending on the selected cloud_provider and kind
	//
	// Types that are assignable to KindValues:
	//
	//	*Cluster_EksKind
	//	*Cluster_GkeKind
	//	*Cluster_AksKind
	KindValues isCluster_KindValues `protobuf_oneof:"kind_values"`
	// soc2_compliant force enables all the various soc2-related fields on a cluster
	Soc2Compliant bool `protobuf:"varint,9,opt,name=soc2_compliant,json=soc2Compliant,proto3" json:"soc2_compliant,omitempty"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_porter_v1_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *Cluster) GetProjectId() int32 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *Cluster) GetClusterId() int32 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

func (x *Cluster) GetKind() EnumKubernetesKind {
	if x != nil {
		return x.Kind
	}
	return EnumKubernetesKind_ENUM_KUBERNETES_KIND_UNSPECIFIED
}

func (x *Cluster) GetCloudProvider() EnumCloudProvider {
	if x != nil {
		return x.CloudProvider
	}
	return EnumCloudProvider_ENUM_CLOUD_PROVIDER_UNSPECIFIED
}

func (x *Cluster) GetCloudProviderCredentialsId() string {
	if x != nil {
		return x.CloudProviderCredentialsId
	}
	return ""
}

func (m *Cluster) GetKindValues() isCluster_KindValues {
	if m != nil {
		return m.KindValues
	}
	return nil
}

func (x *Cluster) GetEksKind() *EKS {
	if x, ok := x.GetKindValues().(*Cluster_EksKind); ok {
		return x.EksKind
	}
	return nil
}

func (x *Cluster) GetGkeKind() *GKE {
	if x, ok := x.GetKindValues().(*Cluster_GkeKind); ok {
		return x.GkeKind
	}
	return nil
}

func (x *Cluster) GetAksKind() *AKS {
	if x, ok := x.GetKindValues().(*Cluster_AksKind); ok {
		return x.AksKind
	}
	return nil
}

func (x *Cluster) GetSoc2Compliant() bool {
	if x != nil {
		return x.Soc2Compliant
	}
	return false
}

type isCluster_KindValues interface {
	isCluster_KindValues()
}

type Cluster_EksKind struct {
	EksKind *EKS `protobuf:"bytes,6,opt,name=eks_kind,json=eksKind,proto3,oneof"`
}

type Cluster_GkeKind struct {
	GkeKind *GKE `protobuf:"bytes,7,opt,name=gke_kind,json=gkeKind,proto3,oneof"`
}

type Cluster_AksKind struct {
	AksKind *AKS `protobuf:"bytes,8,opt,name=aks_kind,json=aksKind,proto3,oneof"`
}

func (*Cluster_EksKind) isCluster_KindValues() {}

func (*Cluster_GkeKind) isCluster_KindValues() {}

func (*Cluster_AksKind) isCluster_KindValues() {}

var File_porter_v1_cluster_proto protoreflect.FileDescriptor

var file_porter_v1_cluster_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6b, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x03, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x31, 0x0a,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x12, 0x43, 0x0a, 0x0e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x0d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x1d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1a, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x08, 0x65, 0x6b, 0x73, 0x5f,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x4b, 0x53, 0x48, 0x00, 0x52, 0x07, 0x65, 0x6b,
	0x73, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x2b, 0x0a, 0x08, 0x67, 0x6b, 0x65, 0x5f, 0x6b, 0x69, 0x6e,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x4b, 0x45, 0x48, 0x00, 0x52, 0x07, 0x67, 0x6b, 0x65, 0x4b, 0x69,
	0x6e, 0x64, 0x12, 0x2b, 0x0a, 0x08, 0x61, 0x6b, 0x73, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x4b, 0x53, 0x48, 0x00, 0x52, 0x07, 0x61, 0x6b, 0x73, 0x4b, 0x69, 0x6e, 0x64, 0x12,
	0x25, 0x0a, 0x0e, 0x73, 0x6f, 0x63, 0x32, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x6f, 0x63, 0x32, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x69, 0x61, 0x6e, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x6b, 0x69, 0x6e, 0x64, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x2a, 0x94, 0x01, 0x0a, 0x12, 0x45, 0x6e, 0x75, 0x6d, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x24, 0x0a, 0x20,
	0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x4b, 0x55, 0x42, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x45, 0x53, 0x5f,
	0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x4b, 0x55, 0x42, 0x45, 0x52,
	0x4e, 0x45, 0x54, 0x45, 0x53, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x45, 0x4b, 0x53, 0x10, 0x01,
	0x12, 0x1c, 0x0a, 0x18, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x4b, 0x55, 0x42, 0x45, 0x52, 0x4e, 0x45,
	0x54, 0x45, 0x53, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x47, 0x4b, 0x45, 0x10, 0x02, 0x12, 0x1c,
	0x0a, 0x18, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x4b, 0x55, 0x42, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x45,
	0x53, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x41, 0x4b, 0x53, 0x10, 0x03, 0x2a, 0x91, 0x01, 0x0a,
	0x11, 0x45, 0x6e, 0x75, 0x6d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x12, 0x23, 0x0a, 0x1f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x43, 0x4c, 0x4f, 0x55, 0x44,
	0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x4e, 0x55, 0x4d, 0x5f,
	0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x41,
	0x57, 0x53, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x43, 0x4c, 0x4f,
	0x55, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x47, 0x43, 0x50, 0x10,
	0x02, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f,
	0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x41, 0x5a, 0x55, 0x52, 0x45, 0x10, 0x03,
	0x42, 0xa7, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x42, 0x0c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x09,
	0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a,
	0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_porter_v1_cluster_proto_rawDescOnce sync.Once
	file_porter_v1_cluster_proto_rawDescData = file_porter_v1_cluster_proto_rawDesc
)

func file_porter_v1_cluster_proto_rawDescGZIP() []byte {
	file_porter_v1_cluster_proto_rawDescOnce.Do(func() {
		file_porter_v1_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_porter_v1_cluster_proto_rawDescData)
	})
	return file_porter_v1_cluster_proto_rawDescData
}

var file_porter_v1_cluster_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_porter_v1_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_porter_v1_cluster_proto_goTypes = []interface{}{
	(EnumKubernetesKind)(0), // 0: porter.v1.EnumKubernetesKind
	(EnumCloudProvider)(0),  // 1: porter.v1.EnumCloudProvider
	(*Cluster)(nil),         // 2: porter.v1.Cluster
	(*EKS)(nil),             // 3: porter.v1.EKS
	(*GKE)(nil),             // 4: porter.v1.GKE
	(*AKS)(nil),             // 5: porter.v1.AKS
}
var file_porter_v1_cluster_proto_depIdxs = []int32{
	0, // 0: porter.v1.Cluster.kind:type_name -> porter.v1.EnumKubernetesKind
	1, // 1: porter.v1.Cluster.cloud_provider:type_name -> porter.v1.EnumCloudProvider
	3, // 2: porter.v1.Cluster.eks_kind:type_name -> porter.v1.EKS
	4, // 3: porter.v1.Cluster.gke_kind:type_name -> porter.v1.GKE
	5, // 4: porter.v1.Cluster.aks_kind:type_name -> porter.v1.AKS
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_porter_v1_cluster_proto_init() }
func file_porter_v1_cluster_proto_init() {
	if File_porter_v1_cluster_proto != nil {
		return
	}
	file_porter_v1_aks_proto_init()
	file_porter_v1_eks_proto_init()
	file_porter_v1_gke_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_porter_v1_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
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
	file_porter_v1_cluster_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Cluster_EksKind)(nil),
		(*Cluster_GkeKind)(nil),
		(*Cluster_AksKind)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_porter_v1_cluster_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_porter_v1_cluster_proto_goTypes,
		DependencyIndexes: file_porter_v1_cluster_proto_depIdxs,
		EnumInfos:         file_porter_v1_cluster_proto_enumTypes,
		MessageInfos:      file_porter_v1_cluster_proto_msgTypes,
	}.Build()
	File_porter_v1_cluster_proto = out.File
	file_porter_v1_cluster_proto_rawDesc = nil
	file_porter_v1_cluster_proto_goTypes = nil
	file_porter_v1_cluster_proto_depIdxs = nil
}
