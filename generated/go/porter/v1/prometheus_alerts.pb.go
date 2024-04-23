// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: porter/v1/prometheus_alerts.proto

package porterv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// InvolvedObjectType is the kubernetes object type the notification targets
// We currently alert for deployments, statefulsets and daemonsets
type InvolvedObjectType int32

const (
	InvolvedObjectType_INVOLVED_OBJECT_TYPE_UNSPECIFIED InvolvedObjectType = 0
	InvolvedObjectType_INVOLVED_OBJECT_TYPE_DEPLOYMENT  InvolvedObjectType = 1
	InvolvedObjectType_INVOLVED_OBJECT_TYPE_STATEFULSET InvolvedObjectType = 2
	InvolvedObjectType_INVOLVED_OBJECT_TYPE_DAEMONSET   InvolvedObjectType = 3
)

// Enum value maps for InvolvedObjectType.
var (
	InvolvedObjectType_name = map[int32]string{
		0: "INVOLVED_OBJECT_TYPE_UNSPECIFIED",
		1: "INVOLVED_OBJECT_TYPE_DEPLOYMENT",
		2: "INVOLVED_OBJECT_TYPE_STATEFULSET",
		3: "INVOLVED_OBJECT_TYPE_DAEMONSET",
	}
	InvolvedObjectType_value = map[string]int32{
		"INVOLVED_OBJECT_TYPE_UNSPECIFIED": 0,
		"INVOLVED_OBJECT_TYPE_DEPLOYMENT":  1,
		"INVOLVED_OBJECT_TYPE_STATEFULSET": 2,
		"INVOLVED_OBJECT_TYPE_DAEMONSET":   3,
	}
)

func (x InvolvedObjectType) Enum() *InvolvedObjectType {
	p := new(InvolvedObjectType)
	*p = x
	return p
}

func (x InvolvedObjectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InvolvedObjectType) Descriptor() protoreflect.EnumDescriptor {
	return file_porter_v1_prometheus_alerts_proto_enumTypes[0].Descriptor()
}

func (InvolvedObjectType) Type() protoreflect.EnumType {
	return &file_porter_v1_prometheus_alerts_proto_enumTypes[0]
}

func (x InvolvedObjectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InvolvedObjectType.Descriptor instead.
func (InvolvedObjectType) EnumDescriptor() ([]byte, []int) {
	return file_porter_v1_prometheus_alerts_proto_rawDescGZIP(), []int{0}
}

// Alert represents a prometheus alert for one target object that is a daemonset, statefulset or deployment
type Alert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string             `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Type      InvolvedObjectType `protobuf:"varint,3,opt,name=type,proto3,enum=porter.v1.InvolvedObjectType" json:"type,omitempty"`
	// severity should be a string specifying how severe the alert is
	Severity string `protobuf:"bytes,4,opt,name=severity,proto3" json:"severity,omitempty"`
	// start_time holds the timestamp for when this status began
	StartTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// end_time holds the timestamp for when this status ended
	// if provided, this Alert is considered as a resolution alert
	EndTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=end_time,json=endTime,proto3,oneof" json:"end_time,omitempty"`
}

func (x *Alert) Reset() {
	*x = Alert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_prometheus_alerts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alert) ProtoMessage() {}

func (x *Alert) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_prometheus_alerts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alert.ProtoReflect.Descriptor instead.
func (*Alert) Descriptor() ([]byte, []int) {
	return file_porter_v1_prometheus_alerts_proto_rawDescGZIP(), []int{0}
}

func (x *Alert) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Alert) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Alert) GetType() InvolvedObjectType {
	if x != nil {
		return x.Type
	}
	return InvolvedObjectType_INVOLVED_OBJECT_TYPE_UNSPECIFIED
}

func (x *Alert) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

func (x *Alert) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Alert) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

var File_porter_v1_prometheus_alerts_proto protoreflect.FileDescriptor

var file_porter_v1_prometheus_alerts_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6d,
	0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x5f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8c, 0x02, 0x0a, 0x05, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x2a, 0xa9,
	0x01, 0x0a, 0x12, 0x49, 0x6e, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x20, 0x49, 0x4e, 0x56, 0x4f, 0x4c, 0x56, 0x45,
	0x44, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x23, 0x0a, 0x1f, 0x49,
	0x4e, 0x56, 0x4f, 0x4c, 0x56, 0x45, 0x44, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x01,
	0x12, 0x24, 0x0a, 0x20, 0x49, 0x4e, 0x56, 0x4f, 0x4c, 0x56, 0x45, 0x44, 0x5f, 0x4f, 0x42, 0x4a,
	0x45, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x46, 0x55,
	0x4c, 0x53, 0x45, 0x54, 0x10, 0x02, 0x12, 0x22, 0x0a, 0x1e, 0x49, 0x4e, 0x56, 0x4f, 0x4c, 0x56,
	0x45, 0x44, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44,
	0x41, 0x45, 0x4d, 0x4f, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x03, 0x42, 0xb0, 0x01, 0x0a, 0x0d, 0x63,
	0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x50, 0x72,
	0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x61, 0x70, 0x69,
	0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x3b, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58,
	0xaa, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x50,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x50, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_porter_v1_prometheus_alerts_proto_rawDescOnce sync.Once
	file_porter_v1_prometheus_alerts_proto_rawDescData = file_porter_v1_prometheus_alerts_proto_rawDesc
)

func file_porter_v1_prometheus_alerts_proto_rawDescGZIP() []byte {
	file_porter_v1_prometheus_alerts_proto_rawDescOnce.Do(func() {
		file_porter_v1_prometheus_alerts_proto_rawDescData = protoimpl.X.CompressGZIP(file_porter_v1_prometheus_alerts_proto_rawDescData)
	})
	return file_porter_v1_prometheus_alerts_proto_rawDescData
}

var file_porter_v1_prometheus_alerts_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_porter_v1_prometheus_alerts_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_porter_v1_prometheus_alerts_proto_goTypes = []interface{}{
	(InvolvedObjectType)(0),       // 0: porter.v1.InvolvedObjectType
	(*Alert)(nil),                 // 1: porter.v1.Alert
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_porter_v1_prometheus_alerts_proto_depIdxs = []int32{
	0, // 0: porter.v1.Alert.type:type_name -> porter.v1.InvolvedObjectType
	2, // 1: porter.v1.Alert.start_time:type_name -> google.protobuf.Timestamp
	2, // 2: porter.v1.Alert.end_time:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_porter_v1_prometheus_alerts_proto_init() }
func file_porter_v1_prometheus_alerts_proto_init() {
	if File_porter_v1_prometheus_alerts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_porter_v1_prometheus_alerts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alert); i {
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
	file_porter_v1_prometheus_alerts_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_porter_v1_prometheus_alerts_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_porter_v1_prometheus_alerts_proto_goTypes,
		DependencyIndexes: file_porter_v1_prometheus_alerts_proto_depIdxs,
		EnumInfos:         file_porter_v1_prometheus_alerts_proto_enumTypes,
		MessageInfos:      file_porter_v1_prometheus_alerts_proto_msgTypes,
	}.Build()
	File_porter_v1_prometheus_alerts_proto = out.File
	file_porter_v1_prometheus_alerts_proto_rawDesc = nil
	file_porter_v1_prometheus_alerts_proto_goTypes = nil
	file_porter_v1_prometheus_alerts_proto_depIdxs = nil
}
