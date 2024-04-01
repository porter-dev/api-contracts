// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: porter/v1/agent_app_event_types.proto

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

type AppEventType int32

const (
	// APP_EVENT_TYPE_UNSPECIFIED is the default value
	AppEventType_APP_EVENT_TYPE_UNSPECIFIED AppEventType = 0
	// APP_EVENT_TYPE_INSUFFICIENT_RESOURCES is generated when pods can't be scheduled becase of a resource shortage
	// This is either CPU or Memory, but agent cannot resolve
	AppEventType_APP_EVENT_TYPE_INSUFFICIENT_RESOURCES AppEventType = 1
	// APP_EVENT_TYPE_MAX_RESOURCE_LIMIT_EXCEEDED is generated when pods can't be scheduled because the node group has reached maximum capacity
	AppEventType_APP_EVENT_TYPE_MAX_RESOURCE_LIMIT_EXCEEDED AppEventType = 3
	// APP_EVENT_TYPE_INSUFFICIENT_CPU is generated when pods can't be scheduled because there is insufficient CPU
	AppEventType_APP_EVENT_TYPE_INSUFFICIENT_CPU AppEventType = 5
	// APP_EVENT_TYPE_INSUFFICIENT_MEMORY is generated when pods can't be scheduled because there is insufficient memory
	AppEventType_APP_EVENT_TYPE_INSUFFICIENT_MEMORY AppEventType = 10
	// APP_EVENT_TYPE_STUCK_PENDING is generated when pods can't be scheduled for other reasons
	AppEventType_APP_EVENT_TYPE_STUCK_PENDING AppEventType = 15
	// APP_EVENT_TYPE_INVALID_IMAGE is generated when the image specified in containers is not valid
	AppEventType_APP_EVENT_TYPE_INVALID_IMAGE AppEventType = 20
	// APP_EVENT_TYPE_INVALID_START_COMMAND is generated when a start command in a container is not valid
	AppEventType_APP_EVENT_TYPE_INVALID_START_COMMAND AppEventType = 25
	// APP_EVENT_TYPE_OUT_OF_MEMORY is generated when a container OOMs
	AppEventType_APP_EVENT_TYPE_OUT_OF_MEMORY AppEventType = 30
	// APP_EVENT_TYPE_NON_ZERO_EXIT_CODE is generated when a container terminates with another non-zero code
	AppEventType_APP_EVENT_TYPE_NON_ZERO_EXIT_CODE AppEventType = 35
	// APP_EVENT_TYPE_FAILING_HEALTH_CHECK is generated when a container is restarted due to a failing health check
	AppEventType_APP_EVENT_TYPE_FAILING_HEALTH_CHECK AppEventType = 40
	// APP_EVENT_TYPE_JOB_TIMEOUT is generated when a job container runs for more than the timeout period set for it
	AppEventType_APP_EVENT_TYPE_JOB_TIMEOUT AppEventType = 45
)

// Enum value maps for AppEventType.
var (
	AppEventType_name = map[int32]string{
		0:  "APP_EVENT_TYPE_UNSPECIFIED",
		1:  "APP_EVENT_TYPE_INSUFFICIENT_RESOURCES",
		3:  "APP_EVENT_TYPE_MAX_RESOURCE_LIMIT_EXCEEDED",
		5:  "APP_EVENT_TYPE_INSUFFICIENT_CPU",
		10: "APP_EVENT_TYPE_INSUFFICIENT_MEMORY",
		15: "APP_EVENT_TYPE_STUCK_PENDING",
		20: "APP_EVENT_TYPE_INVALID_IMAGE",
		25: "APP_EVENT_TYPE_INVALID_START_COMMAND",
		30: "APP_EVENT_TYPE_OUT_OF_MEMORY",
		35: "APP_EVENT_TYPE_NON_ZERO_EXIT_CODE",
		40: "APP_EVENT_TYPE_FAILING_HEALTH_CHECK",
		45: "APP_EVENT_TYPE_JOB_TIMEOUT",
	}
	AppEventType_value = map[string]int32{
		"APP_EVENT_TYPE_UNSPECIFIED":                 0,
		"APP_EVENT_TYPE_INSUFFICIENT_RESOURCES":      1,
		"APP_EVENT_TYPE_MAX_RESOURCE_LIMIT_EXCEEDED": 3,
		"APP_EVENT_TYPE_INSUFFICIENT_CPU":            5,
		"APP_EVENT_TYPE_INSUFFICIENT_MEMORY":         10,
		"APP_EVENT_TYPE_STUCK_PENDING":               15,
		"APP_EVENT_TYPE_INVALID_IMAGE":               20,
		"APP_EVENT_TYPE_INVALID_START_COMMAND":       25,
		"APP_EVENT_TYPE_OUT_OF_MEMORY":               30,
		"APP_EVENT_TYPE_NON_ZERO_EXIT_CODE":          35,
		"APP_EVENT_TYPE_FAILING_HEALTH_CHECK":        40,
		"APP_EVENT_TYPE_JOB_TIMEOUT":                 45,
	}
)

func (x AppEventType) Enum() *AppEventType {
	p := new(AppEventType)
	*p = x
	return p
}

func (x AppEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_porter_v1_agent_app_event_types_proto_enumTypes[0].Descriptor()
}

func (AppEventType) Type() protoreflect.EnumType {
	return &file_porter_v1_agent_app_event_types_proto_enumTypes[0]
}

func (x AppEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppEventType.Descriptor instead.
func (AppEventType) EnumDescriptor() ([]byte, []int) {
	return file_porter_v1_agent_app_event_types_proto_rawDescGZIP(), []int{0}
}

var File_porter_v1_agent_app_event_types_proto protoreflect.FileDescriptor

var file_porter_v1_agent_app_event_types_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2a, 0xd6, 0x03, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x41, 0x50, 0x50, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x29, 0x0a, 0x25, 0x41, 0x50, 0x50, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x53, 0x55, 0x46, 0x46, 0x49, 0x43, 0x49, 0x45,
	0x4e, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x53, 0x10, 0x01, 0x12, 0x2e,
	0x0a, 0x2a, 0x41, 0x50, 0x50, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4d, 0x41, 0x58, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4c, 0x49,
	0x4d, 0x49, 0x54, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x03, 0x12, 0x23,
	0x0a, 0x1f, 0x41, 0x50, 0x50, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x49, 0x4e, 0x53, 0x55, 0x46, 0x46, 0x49, 0x43, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x50,
	0x55, 0x10, 0x05, 0x12, 0x26, 0x0a, 0x22, 0x41, 0x50, 0x50, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x53, 0x55, 0x46, 0x46, 0x49, 0x43, 0x49, 0x45,
	0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59, 0x10, 0x0a, 0x12, 0x20, 0x0a, 0x1c, 0x41,
	0x50, 0x50, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54,
	0x55, 0x43, 0x4b, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x0f, 0x12, 0x20, 0x0a,
	0x1c, 0x41, 0x50, 0x50, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x14, 0x12,
	0x28, 0x0a, 0x24, 0x41, 0x50, 0x50, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f,
	0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x10, 0x19, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x50, 0x50,
	0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x5f,
	0x4f, 0x46, 0x5f, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59, 0x10, 0x1e, 0x12, 0x25, 0x0a, 0x21, 0x41,
	0x50, 0x50, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f,
	0x4e, 0x5f, 0x5a, 0x45, 0x52, 0x4f, 0x5f, 0x45, 0x58, 0x49, 0x54, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x10, 0x23, 0x12, 0x27, 0x0a, 0x23, 0x41, 0x50, 0x50, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x48, 0x45, 0x41,
	0x4c, 0x54, 0x48, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x10, 0x28, 0x12, 0x1e, 0x0a, 0x1a, 0x41,
	0x50, 0x50, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4a, 0x4f,
	0x42, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x2d, 0x42, 0xb2, 0x01, 0x0a, 0x0d,
	0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x17, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x70, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2d, 0x64, 0x65, 0x76, 0x2f,
	0x61, 0x70, 0x69, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x50, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x50, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_porter_v1_agent_app_event_types_proto_rawDescOnce sync.Once
	file_porter_v1_agent_app_event_types_proto_rawDescData = file_porter_v1_agent_app_event_types_proto_rawDesc
)

func file_porter_v1_agent_app_event_types_proto_rawDescGZIP() []byte {
	file_porter_v1_agent_app_event_types_proto_rawDescOnce.Do(func() {
		file_porter_v1_agent_app_event_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_porter_v1_agent_app_event_types_proto_rawDescData)
	})
	return file_porter_v1_agent_app_event_types_proto_rawDescData
}

var file_porter_v1_agent_app_event_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_porter_v1_agent_app_event_types_proto_goTypes = []interface{}{
	(AppEventType)(0), // 0: porter.v1.AppEventType
}
var file_porter_v1_agent_app_event_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_porter_v1_agent_app_event_types_proto_init() }
func file_porter_v1_agent_app_event_types_proto_init() {
	if File_porter_v1_agent_app_event_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_porter_v1_agent_app_event_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_porter_v1_agent_app_event_types_proto_goTypes,
		DependencyIndexes: file_porter_v1_agent_app_event_types_proto_depIdxs,
		EnumInfos:         file_porter_v1_agent_app_event_types_proto_enumTypes,
	}.Build()
	File_porter_v1_agent_app_event_types_proto = out.File
	file_porter_v1_agent_app_event_types_proto_rawDesc = nil
	file_porter_v1_agent_app_event_types_proto_goTypes = nil
	file_porter_v1_agent_app_event_types_proto_depIdxs = nil
}
