// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: porter/v1/aks.proto

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

type NodePoolType int32

const (
	NodePoolType_NODE_POOL_TYPE_UNSPECIFIED NodePoolType = 0
	NodePoolType_NODE_POOL_TYPE_SYSTEM      NodePoolType = 1
	NodePoolType_NODE_POOL_TYPE_MONITORING  NodePoolType = 2
	NodePoolType_NODE_POOL_TYPE_APPLICATION NodePoolType = 3
	NodePoolType_NODE_POOL_TYPE_CUSTOM      NodePoolType = 4
)

// Enum value maps for NodePoolType.
var (
	NodePoolType_name = map[int32]string{
		0: "NODE_POOL_TYPE_UNSPECIFIED",
		1: "NODE_POOL_TYPE_SYSTEM",
		2: "NODE_POOL_TYPE_MONITORING",
		3: "NODE_POOL_TYPE_APPLICATION",
		4: "NODE_POOL_TYPE_CUSTOM",
	}
	NodePoolType_value = map[string]int32{
		"NODE_POOL_TYPE_UNSPECIFIED": 0,
		"NODE_POOL_TYPE_SYSTEM":      1,
		"NODE_POOL_TYPE_MONITORING":  2,
		"NODE_POOL_TYPE_APPLICATION": 3,
		"NODE_POOL_TYPE_CUSTOM":      4,
	}
)

func (x NodePoolType) Enum() *NodePoolType {
	p := new(NodePoolType)
	*p = x
	return p
}

func (x NodePoolType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodePoolType) Descriptor() protoreflect.EnumDescriptor {
	return file_porter_v1_aks_proto_enumTypes[0].Descriptor()
}

func (NodePoolType) Type() protoreflect.EnumType {
	return &file_porter_v1_aks_proto_enumTypes[0]
}

func (x NodePoolType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodePoolType.Descriptor instead.
func (NodePoolType) EnumDescriptor() ([]byte, []int) {
	return file_porter_v1_aks_proto_rawDescGZIP(), []int{0}
}

type AKS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName    string         `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	ClusterVersion string         `protobuf:"bytes,2,opt,name=cluster_version,json=clusterVersion,proto3" json:"cluster_version,omitempty"`
	CidrRange      string         `protobuf:"bytes,3,opt,name=cidr_range,json=cidrRange,proto3" json:"cidr_range,omitempty"`
	Location       string         `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	NodePools      []*AKSNodePool `protobuf:"bytes,5,rep,name=node_pools,json=nodePools,proto3" json:"node_pools,omitempty"`
}

func (x *AKS) Reset() {
	*x = AKS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_aks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AKS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AKS) ProtoMessage() {}

func (x *AKS) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_aks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AKS.ProtoReflect.Descriptor instead.
func (*AKS) Descriptor() ([]byte, []int) {
	return file_porter_v1_aks_proto_rawDescGZIP(), []int{0}
}

func (x *AKS) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *AKS) GetClusterVersion() string {
	if x != nil {
		return x.ClusterVersion
	}
	return ""
}

func (x *AKS) GetCidrRange() string {
	if x != nil {
		return x.CidrRange
	}
	return ""
}

func (x *AKS) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *AKS) GetNodePools() []*AKSNodePool {
	if x != nil {
		return x.NodePools
	}
	return nil
}

type AKSNodePool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceType string       `protobuf:"bytes,1,opt,name=instance_type,json=instanceType,proto3" json:"instance_type,omitempty"`
	MinInstances uint32       `protobuf:"varint,2,opt,name=min_instances,json=minInstances,proto3" json:"min_instances,omitempty"`
	MaxInstances uint32       `protobuf:"varint,3,opt,name=max_instances,json=maxInstances,proto3" json:"max_instances,omitempty"`
	Mode         string       `protobuf:"bytes,4,opt,name=mode,proto3" json:"mode,omitempty"`
	NodePoolType NodePoolType `protobuf:"varint,5,opt,name=node_pool_type,json=nodePoolType,proto3,enum=porter.v1.NodePoolType" json:"node_pool_type,omitempty"`
}

func (x *AKSNodePool) Reset() {
	*x = AKSNodePool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_aks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AKSNodePool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AKSNodePool) ProtoMessage() {}

func (x *AKSNodePool) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_aks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AKSNodePool.ProtoReflect.Descriptor instead.
func (*AKSNodePool) Descriptor() ([]byte, []int) {
	return file_porter_v1_aks_proto_rawDescGZIP(), []int{1}
}

func (x *AKSNodePool) GetInstanceType() string {
	if x != nil {
		return x.InstanceType
	}
	return ""
}

func (x *AKSNodePool) GetMinInstances() uint32 {
	if x != nil {
		return x.MinInstances
	}
	return 0
}

func (x *AKSNodePool) GetMaxInstances() uint32 {
	if x != nil {
		return x.MaxInstances
	}
	return 0
}

func (x *AKSNodePool) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *AKSNodePool) GetNodePoolType() NodePoolType {
	if x != nil {
		return x.NodePoolType
	}
	return NodePoolType_NODE_POOL_TYPE_UNSPECIFIED
}

var File_porter_v1_aks_proto protoreflect.FileDescriptor

var file_porter_v1_aks_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6b, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x22, 0xc3, 0x01, 0x0a, 0x03, 0x41, 0x4b, 0x53, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x69, 0x64, 0x72, 0x5f, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x69, 0x64, 0x72, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x35, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x4b, 0x53, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x09, 0x6e, 0x6f, 0x64,
	0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x22, 0xcf, 0x01, 0x0a, 0x0b, 0x41, 0x4b, 0x53, 0x4e, 0x6f,
	0x64, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d,
	0x69, 0x6e, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x17, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x6e, 0x6f, 0x64, 0x65,
	0x50, 0x6f, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x2a, 0xa3, 0x01, 0x0a, 0x0c, 0x4e, 0x6f, 0x64,
	0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x4e, 0x4f, 0x44,
	0x45, 0x5f, 0x50, 0x4f, 0x4f, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x4e, 0x4f, 0x44,
	0x45, 0x5f, 0x50, 0x4f, 0x4f, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x54,
	0x45, 0x4d, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x4f, 0x4f,
	0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x4f, 0x4e, 0x49, 0x54, 0x4f, 0x52, 0x49, 0x4e,
	0x47, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x4f, 0x4f, 0x4c,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x4f, 0x4f, 0x4c,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x04, 0x42, 0xa3,
	0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x42, 0x08, 0x41, 0x6b, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2d,
	0x64, 0x65, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x15, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_porter_v1_aks_proto_rawDescOnce sync.Once
	file_porter_v1_aks_proto_rawDescData = file_porter_v1_aks_proto_rawDesc
)

func file_porter_v1_aks_proto_rawDescGZIP() []byte {
	file_porter_v1_aks_proto_rawDescOnce.Do(func() {
		file_porter_v1_aks_proto_rawDescData = protoimpl.X.CompressGZIP(file_porter_v1_aks_proto_rawDescData)
	})
	return file_porter_v1_aks_proto_rawDescData
}

var file_porter_v1_aks_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_porter_v1_aks_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_porter_v1_aks_proto_goTypes = []interface{}{
	(NodePoolType)(0),   // 0: porter.v1.NodePoolType
	(*AKS)(nil),         // 1: porter.v1.AKS
	(*AKSNodePool)(nil), // 2: porter.v1.AKSNodePool
}
var file_porter_v1_aks_proto_depIdxs = []int32{
	2, // 0: porter.v1.AKS.node_pools:type_name -> porter.v1.AKSNodePool
	0, // 1: porter.v1.AKSNodePool.node_pool_type:type_name -> porter.v1.NodePoolType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_porter_v1_aks_proto_init() }
func file_porter_v1_aks_proto_init() {
	if File_porter_v1_aks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_porter_v1_aks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AKS); i {
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
		file_porter_v1_aks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AKSNodePool); i {
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
			RawDescriptor: file_porter_v1_aks_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_porter_v1_aks_proto_goTypes,
		DependencyIndexes: file_porter_v1_aks_proto_depIdxs,
		EnumInfos:         file_porter_v1_aks_proto_enumTypes,
		MessageInfos:      file_porter_v1_aks_proto_msgTypes,
	}.Build()
	File_porter_v1_aks_proto = out.File
	file_porter_v1_aks_proto_rawDesc = nil
	file_porter_v1_aks_proto_goTypes = nil
	file_porter_v1_aks_proto_depIdxs = nil
}
