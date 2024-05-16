// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: porter/v1/dead_letter.proto

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

type DeadLetter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractRevision *ContractRevision `protobuf:"bytes,1,opt,name=contract_revision,json=contractRevision,proto3" json:"contract_revision,omitempty"`
	Error            *Error            `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeadLetter) Reset() {
	*x = DeadLetter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_dead_letter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeadLetter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeadLetter) ProtoMessage() {}

func (x *DeadLetter) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_dead_letter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeadLetter.ProtoReflect.Descriptor instead.
func (*DeadLetter) Descriptor() ([]byte, []int) {
	return file_porter_v1_dead_letter_proto_rawDescGZIP(), []int{0}
}

func (x *DeadLetter) GetContractRevision() *ContractRevision {
	if x != nil {
		return x.ContractRevision
	}
	return nil
}

func (x *DeadLetter) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_porter_v1_dead_letter_proto protoreflect.FileDescriptor

var file_porter_v1_dead_letter_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x61, 0x64,
	0x5f, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x0a, 0x44, 0x65,
	0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0xaa, 0x01, 0x0a, 0x0d, 0x63,
	0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x44, 0x65,
	0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67,
	0x6f, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x50, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x15, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x50, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_porter_v1_dead_letter_proto_rawDescOnce sync.Once
	file_porter_v1_dead_letter_proto_rawDescData = file_porter_v1_dead_letter_proto_rawDesc
)

func file_porter_v1_dead_letter_proto_rawDescGZIP() []byte {
	file_porter_v1_dead_letter_proto_rawDescOnce.Do(func() {
		file_porter_v1_dead_letter_proto_rawDescData = protoimpl.X.CompressGZIP(file_porter_v1_dead_letter_proto_rawDescData)
	})
	return file_porter_v1_dead_letter_proto_rawDescData
}

var file_porter_v1_dead_letter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_porter_v1_dead_letter_proto_goTypes = []interface{}{
	(*DeadLetter)(nil),       // 0: porter.v1.DeadLetter
	(*ContractRevision)(nil), // 1: porter.v1.ContractRevision
	(*Error)(nil),            // 2: porter.v1.Error
}
var file_porter_v1_dead_letter_proto_depIdxs = []int32{
	1, // 0: porter.v1.DeadLetter.contract_revision:type_name -> porter.v1.ContractRevision
	2, // 1: porter.v1.DeadLetter.error:type_name -> porter.v1.Error
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_porter_v1_dead_letter_proto_init() }
func file_porter_v1_dead_letter_proto_init() {
	if File_porter_v1_dead_letter_proto != nil {
		return
	}
	file_porter_v1_contract_proto_init()
	file_porter_v1_errors_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_porter_v1_dead_letter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeadLetter); i {
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
			RawDescriptor: file_porter_v1_dead_letter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_porter_v1_dead_letter_proto_goTypes,
		DependencyIndexes: file_porter_v1_dead_letter_proto_depIdxs,
		MessageInfos:      file_porter_v1_dead_letter_proto_msgTypes,
	}.Build()
	File_porter_v1_dead_letter_proto = out.File
	file_porter_v1_dead_letter_proto_rawDesc = nil
	file_porter_v1_dead_letter_proto_goTypes = nil
	file_porter_v1_dead_letter_proto_depIdxs = nil
}
