// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: porter/v1/gke.proto

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

type GKENodePoolType int32

const (
	GKENodePoolType_GKE_NODE_POOL_TYPE_UNSPECIFIED GKENodePoolType = 0
	GKENodePoolType_GKE_NODE_POOL_TYPE_SYSTEM      GKENodePoolType = 1
	GKENodePoolType_GKE_NODE_POOL_TYPE_MONITORING  GKENodePoolType = 2
	GKENodePoolType_GKE_NODE_POOL_TYPE_APPLICATION GKENodePoolType = 3
	GKENodePoolType_GKE_NODE_POOL_TYPE_CUSTOM      GKENodePoolType = 4
	GKENodePoolType_GKE_NODE_POOL_TYPE_USER        GKENodePoolType = 5
)

// Enum value maps for GKENodePoolType.
var (
	GKENodePoolType_name = map[int32]string{
		0: "GKE_NODE_POOL_TYPE_UNSPECIFIED",
		1: "GKE_NODE_POOL_TYPE_SYSTEM",
		2: "GKE_NODE_POOL_TYPE_MONITORING",
		3: "GKE_NODE_POOL_TYPE_APPLICATION",
		4: "GKE_NODE_POOL_TYPE_CUSTOM",
		5: "GKE_NODE_POOL_TYPE_USER",
	}
	GKENodePoolType_value = map[string]int32{
		"GKE_NODE_POOL_TYPE_UNSPECIFIED": 0,
		"GKE_NODE_POOL_TYPE_SYSTEM":      1,
		"GKE_NODE_POOL_TYPE_MONITORING":  2,
		"GKE_NODE_POOL_TYPE_APPLICATION": 3,
		"GKE_NODE_POOL_TYPE_CUSTOM":      4,
		"GKE_NODE_POOL_TYPE_USER":        5,
	}
)

func (x GKENodePoolType) Enum() *GKENodePoolType {
	p := new(GKENodePoolType)
	*p = x
	return p
}

func (x GKENodePoolType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GKENodePoolType) Descriptor() protoreflect.EnumDescriptor {
	return file_porter_v1_gke_proto_enumTypes[0].Descriptor()
}

func (GKENodePoolType) Type() protoreflect.EnumType {
	return &file_porter_v1_gke_proto_enumTypes[0]
}

func (x GKENodePoolType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GKENodePoolType.Descriptor instead.
func (GKENodePoolType) EnumDescriptor() ([]byte, []int) {
	return file_porter_v1_gke_proto_rawDescGZIP(), []int{0}
}

type GKE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName    string         `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	ClusterVersion string         `protobuf:"bytes,2,opt,name=cluster_version,json=clusterVersion,proto3" json:"cluster_version,omitempty"`
	Network        *GKENetwork    `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
	Region         string         `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	NodePools      []*GKENodePool `protobuf:"bytes,5,rep,name=node_pools,json=nodePools,proto3" json:"node_pools,omitempty"`
}

func (x *GKE) Reset() {
	*x = GKE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_gke_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GKE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GKE) ProtoMessage() {}

func (x *GKE) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_gke_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GKE.ProtoReflect.Descriptor instead.
func (*GKE) Descriptor() ([]byte, []int) {
	return file_porter_v1_gke_proto_rawDescGZIP(), []int{0}
}

func (x *GKE) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *GKE) GetClusterVersion() string {
	if x != nil {
		return x.ClusterVersion
	}
	return ""
}

func (x *GKE) GetNetwork() *GKENetwork {
	if x != nil {
		return x.Network
	}
	return nil
}

func (x *GKE) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *GKE) GetNodePools() []*GKENodePool {
	if x != nil {
		return x.NodePools
	}
	return nil
}

type GKENodePool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceType string `protobuf:"bytes,1,opt,name=instance_type,json=instanceType,proto3" json:"instance_type,omitempty"`
	MinInstances uint32 `protobuf:"varint,2,opt,name=min_instances,json=minInstances,proto3" json:"min_instances,omitempty"`
	MaxInstances uint32 `protobuf:"varint,3,opt,name=max_instances,json=maxInstances,proto3" json:"max_instances,omitempty"`
	// node_pool_type is used to specify the type of node pool. This is used to specify what node pools are used by Porter, vs users.
	NodePoolType GKENodePoolType `protobuf:"varint,4,opt,name=node_pool_type,json=nodePoolType,proto3,enum=porter.v1.GKENodePoolType" json:"node_pool_type,omitempty"`
	// additional_taints is a list of NoSchedule taints to apply to the node pool.
	// These will be applied on top of the default porter.run/workload-kind taints.
	AdditionalTaints []string `protobuf:"bytes,5,rep,name=additional_taints,json=additionalTaints,proto3" json:"additional_taints,omitempty"`
	// is_spot_instance is used to specify whether the node pool should use spot instances.
	IsSpotInstance bool `protobuf:"varint,6,opt,name=is_spot_instance,json=isSpotInstance,proto3" json:"is_spot_instance,omitempty"`
	// node_pool_id is the id of the node pool. This uniquely identifies GKENodePoolType=User and is generated on creation if not provided.
	NodePoolId string `protobuf:"bytes,7,opt,name=node_pool_id,json=nodePoolId,proto3" json:"node_pool_id,omitempty"`
	// node_pool_name is the vanity name of the node pool. This is required for GKENodePoolType=User and can be changed by the user.
	NodePoolName string `protobuf:"bytes,8,opt,name=node_pool_name,json=nodePoolName,proto3" json:"node_pool_name,omitempty"`
}

func (x *GKENodePool) Reset() {
	*x = GKENodePool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_gke_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GKENodePool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GKENodePool) ProtoMessage() {}

func (x *GKENodePool) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_gke_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GKENodePool.ProtoReflect.Descriptor instead.
func (*GKENodePool) Descriptor() ([]byte, []int) {
	return file_porter_v1_gke_proto_rawDescGZIP(), []int{1}
}

func (x *GKENodePool) GetInstanceType() string {
	if x != nil {
		return x.InstanceType
	}
	return ""
}

func (x *GKENodePool) GetMinInstances() uint32 {
	if x != nil {
		return x.MinInstances
	}
	return 0
}

func (x *GKENodePool) GetMaxInstances() uint32 {
	if x != nil {
		return x.MaxInstances
	}
	return 0
}

func (x *GKENodePool) GetNodePoolType() GKENodePoolType {
	if x != nil {
		return x.NodePoolType
	}
	return GKENodePoolType_GKE_NODE_POOL_TYPE_UNSPECIFIED
}

func (x *GKENodePool) GetAdditionalTaints() []string {
	if x != nil {
		return x.AdditionalTaints
	}
	return nil
}

func (x *GKENodePool) GetIsSpotInstance() bool {
	if x != nil {
		return x.IsSpotInstance
	}
	return false
}

func (x *GKENodePool) GetNodePoolId() string {
	if x != nil {
		return x.NodePoolId
	}
	return ""
}

func (x *GKENodePool) GetNodePoolName() string {
	if x != nil {
		return x.NodePoolName
	}
	return ""
}

// GKENetwork contains all information required to configure the GKE cluster's network
type GKENetwork struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cidr_range is the range of the network. This is used to specify the network that Porter will use.
	CidrRange string `protobuf:"bytes,1,opt,name=cidr_range,json=cidrRange,proto3" json:"cidr_range,omitempty"`
	// control_plane_cidr is a range reserved by GKE for control plane functions such as running a global load balancer.
	ControlPlaneCidr string `protobuf:"bytes,2,opt,name=control_plane_cidr,json=controlPlaneCidr,proto3" json:"control_plane_cidr,omitempty"`
	// pod_cidr is the range of the network that pods will be assigned IPs from, on the GCP subnet.
	PodCidr string `protobuf:"bytes,3,opt,name=pod_cidr,json=podCidr,proto3" json:"pod_cidr,omitempty"`
	// service_cidr is the range of the network that services will be assigned IPs from, on the GCP subnet.
	ServiceCidr string `protobuf:"bytes,4,opt,name=service_cidr,json=serviceCidr,proto3" json:"service_cidr,omitempty"`
}

func (x *GKENetwork) Reset() {
	*x = GKENetwork{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_gke_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GKENetwork) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GKENetwork) ProtoMessage() {}

func (x *GKENetwork) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_gke_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GKENetwork.ProtoReflect.Descriptor instead.
func (*GKENetwork) Descriptor() ([]byte, []int) {
	return file_porter_v1_gke_proto_rawDescGZIP(), []int{2}
}

func (x *GKENetwork) GetCidrRange() string {
	if x != nil {
		return x.CidrRange
	}
	return ""
}

func (x *GKENetwork) GetControlPlaneCidr() string {
	if x != nil {
		return x.ControlPlaneCidr
	}
	return ""
}

func (x *GKENetwork) GetPodCidr() string {
	if x != nil {
		return x.PodCidr
	}
	return ""
}

func (x *GKENetwork) GetServiceCidr() string {
	if x != nil {
		return x.ServiceCidr
	}
	return ""
}

// GKEPreflightValues is cidr ranges needed for PreflightChecks
type GKEPreflightValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// //Network contains all information required to configure the GKE cluster's network
	Network *GKENetwork `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
}

func (x *GKEPreflightValues) Reset() {
	*x = GKEPreflightValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_gke_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GKEPreflightValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GKEPreflightValues) ProtoMessage() {}

func (x *GKEPreflightValues) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_gke_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GKEPreflightValues.ProtoReflect.Descriptor instead.
func (*GKEPreflightValues) Descriptor() ([]byte, []int) {
	return file_porter_v1_gke_proto_rawDescGZIP(), []int{3}
}

func (x *GKEPreflightValues) GetNetwork() *GKENetwork {
	if x != nil {
		return x.Network
	}
	return nil
}

var File_porter_v1_gke_proto protoreflect.FileDescriptor

var file_porter_v1_gke_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6b, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x22, 0xd1, 0x01, 0x0a, 0x03, 0x47, 0x4b, 0x45, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x4b, 0x45, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x07, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a,
	0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x4b,
	0x45, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x50,
	0x6f, 0x6f, 0x6c, 0x73, 0x22, 0xdd, 0x02, 0x0a, 0x0b, 0x47, 0x4b, 0x45, 0x4e, 0x6f, 0x64, 0x65,
	0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x69, 0x6e,
	0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0e, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x4b, 0x45, 0x4e, 0x6f, 0x64, 0x65, 0x50,
	0x6f, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x6f,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x10, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x54, 0x61, 0x69, 0x6e,
	0x74, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x73, 0x5f, 0x73, 0x70, 0x6f, 0x74, 0x5f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73,
	0x53, 0x70, 0x6f, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0c,
	0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x12, 0x24,
	0x0a, 0x0e, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x6f, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x97, 0x01, 0x0a, 0x0a, 0x47, 0x4b, 0x45, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x69, 0x64, 0x72, 0x5f, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x69, 0x64, 0x72, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x43, 0x69, 0x64, 0x72,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x64, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x64, 0x43, 0x69, 0x64, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x69, 0x64, 0x72, 0x22, 0x45,
	0x0a, 0x12, 0x47, 0x4b, 0x45, 0x50, 0x72, 0x65, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x4b, 0x45, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x07, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2a, 0xd7, 0x01, 0x0a, 0x0f, 0x47, 0x4b, 0x45, 0x4e, 0x6f, 0x64,
	0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x1e, 0x47, 0x4b, 0x45,
	0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x4f, 0x4f, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a,
	0x19, 0x47, 0x4b, 0x45, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x4f, 0x4f, 0x4c, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d,
	0x47, 0x4b, 0x45, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x4f, 0x4f, 0x4c, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4d, 0x4f, 0x4e, 0x49, 0x54, 0x4f, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12,
	0x22, 0x0a, 0x1e, 0x47, 0x4b, 0x45, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x4f, 0x4f, 0x4c,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x47, 0x4b, 0x45, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x5f,
	0x50, 0x4f, 0x4f, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d,
	0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x47, 0x4b, 0x45, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x50,
	0x4f, 0x4f, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x05, 0x42,
	0xa3, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x42, 0x08, 0x47, 0x6b, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2d, 0x64, 0x65, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x15, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_porter_v1_gke_proto_rawDescOnce sync.Once
	file_porter_v1_gke_proto_rawDescData = file_porter_v1_gke_proto_rawDesc
)

func file_porter_v1_gke_proto_rawDescGZIP() []byte {
	file_porter_v1_gke_proto_rawDescOnce.Do(func() {
		file_porter_v1_gke_proto_rawDescData = protoimpl.X.CompressGZIP(file_porter_v1_gke_proto_rawDescData)
	})
	return file_porter_v1_gke_proto_rawDescData
}

var file_porter_v1_gke_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_porter_v1_gke_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_porter_v1_gke_proto_goTypes = []any{
	(GKENodePoolType)(0),       // 0: porter.v1.GKENodePoolType
	(*GKE)(nil),                // 1: porter.v1.GKE
	(*GKENodePool)(nil),        // 2: porter.v1.GKENodePool
	(*GKENetwork)(nil),         // 3: porter.v1.GKENetwork
	(*GKEPreflightValues)(nil), // 4: porter.v1.GKEPreflightValues
}
var file_porter_v1_gke_proto_depIdxs = []int32{
	3, // 0: porter.v1.GKE.network:type_name -> porter.v1.GKENetwork
	2, // 1: porter.v1.GKE.node_pools:type_name -> porter.v1.GKENodePool
	0, // 2: porter.v1.GKENodePool.node_pool_type:type_name -> porter.v1.GKENodePoolType
	3, // 3: porter.v1.GKEPreflightValues.network:type_name -> porter.v1.GKENetwork
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_porter_v1_gke_proto_init() }
func file_porter_v1_gke_proto_init() {
	if File_porter_v1_gke_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_porter_v1_gke_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GKE); i {
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
		file_porter_v1_gke_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GKENodePool); i {
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
		file_porter_v1_gke_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GKENetwork); i {
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
		file_porter_v1_gke_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GKEPreflightValues); i {
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
			RawDescriptor: file_porter_v1_gke_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_porter_v1_gke_proto_goTypes,
		DependencyIndexes: file_porter_v1_gke_proto_depIdxs,
		EnumInfos:         file_porter_v1_gke_proto_enumTypes,
		MessageInfos:      file_porter_v1_gke_proto_msgTypes,
	}.Build()
	File_porter_v1_gke_proto = out.File
	file_porter_v1_gke_proto_rawDesc = nil
	file_porter_v1_gke_proto_goTypes = nil
	file_porter_v1_gke_proto_depIdxs = nil
}
