// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: porter/v1/porter_app.proto

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

// PorterApp is the top-level configuration for a Porter application, usually found in porter.yaml
type PorterApp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the name of the application
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Services is a map of service names to service configurations
	Services map[string]*Service `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Environ is a map of environment variable names to values
	Env map[string]string `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Build is the build settings for the application
	Build *Build `protobuf:"bytes,4,opt,name=build,proto3" json:"build,omitempty"`
	// Predeploy is a job service to run before deploying the application
	Predeploy *Service `protobuf:"bytes,5,opt,name=predeploy,proto3" json:"predeploy,omitempty"`
	// Image is the image to use for a given revision of the application
	Image *AppImage `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *PorterApp) Reset() {
	*x = PorterApp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_porter_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PorterApp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PorterApp) ProtoMessage() {}

func (x *PorterApp) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_porter_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PorterApp.ProtoReflect.Descriptor instead.
func (*PorterApp) Descriptor() ([]byte, []int) {
	return file_porter_v1_porter_app_proto_rawDescGZIP(), []int{0}
}

func (x *PorterApp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PorterApp) GetServices() map[string]*Service {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *PorterApp) GetEnv() map[string]string {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *PorterApp) GetBuild() *Build {
	if x != nil {
		return x.Build
	}
	return nil
}

func (x *PorterApp) GetPredeploy() *Service {
	if x != nil {
		return x.Predeploy
	}
	return nil
}

func (x *PorterApp) GetImage() *AppImage {
	if x != nil {
		return x.Image
	}
	return nil
}

// Build is the build settings for the application
type Build struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Context is the path to the build context
	Context string `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	// Method is the build method to use, being one of "pack", "docker", or "registry"
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	// Builder is the builder to use for the "pack" build method
	Builder string `protobuf:"bytes,3,opt,name=builder,proto3" json:"builder,omitempty"`
	// Buildpacks is a list of buildpacks to use for the "pack" build method
	Buildpacks []string `protobuf:"bytes,4,rep,name=buildpacks,proto3" json:"buildpacks,omitempty"`
	// Dockerfile is the path to the Dockerfile to use for the "docker" build method
	Dockerfile string `protobuf:"bytes,5,opt,name=dockerfile,proto3" json:"dockerfile,omitempty"`
}

func (x *Build) Reset() {
	*x = Build{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_porter_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Build) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Build) ProtoMessage() {}

func (x *Build) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_porter_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Build.ProtoReflect.Descriptor instead.
func (*Build) Descriptor() ([]byte, []int) {
	return file_porter_v1_porter_app_proto_rawDescGZIP(), []int{1}
}

func (x *Build) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

func (x *Build) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Build) GetBuilder() string {
	if x != nil {
		return x.Builder
	}
	return ""
}

func (x *Build) GetBuildpacks() []string {
	if x != nil {
		return x.Buildpacks
	}
	return nil
}

func (x *Build) GetDockerfile() string {
	if x != nil {
		return x.Dockerfile
	}
	return ""
}

// AppImage is the image to use for a given revision of the application
type AppImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Repository is the repository to use for the image
	Repository string `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	// Tag is the tag to use for the image
	Tag string `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *AppImage) Reset() {
	*x = AppImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_porter_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppImage) ProtoMessage() {}

func (x *AppImage) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_porter_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppImage.ProtoReflect.Descriptor instead.
func (*AppImage) Descriptor() ([]byte, []int) {
	return file_porter_v1_porter_app_proto_rawDescGZIP(), []int{2}
}

func (x *AppImage) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *AppImage) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

var File_porter_v1_porter_app_proto protoreflect.FileDescriptor

var file_porter_v1_porter_app_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9e, 0x03, 0x0a, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x12, 0x2f, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x41, 0x70, 0x70, 0x2e, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03,
	0x65, 0x6e, 0x76, 0x12, 0x26, 0x0a, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x52, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x30, 0x0a, 0x09, 0x70,
	0x72, 0x65, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x09, 0x70, 0x72, 0x65, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x12, 0x29, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x1a, 0x4f, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x36, 0x0a, 0x08, 0x45, 0x6e, 0x76,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x93, 0x01, 0x0a, 0x05, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x70, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x70, 0x61, 0x63, 0x6b, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x3c, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x74, 0x61, 0x67, 0x42, 0xa9, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x41,
	0x70, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2d, 0x64, 0x65, 0x76,
	0x2f, 0x61, 0x70, 0x69, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x50,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_porter_v1_porter_app_proto_rawDescOnce sync.Once
	file_porter_v1_porter_app_proto_rawDescData = file_porter_v1_porter_app_proto_rawDesc
)

func file_porter_v1_porter_app_proto_rawDescGZIP() []byte {
	file_porter_v1_porter_app_proto_rawDescOnce.Do(func() {
		file_porter_v1_porter_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_porter_v1_porter_app_proto_rawDescData)
	})
	return file_porter_v1_porter_app_proto_rawDescData
}

var file_porter_v1_porter_app_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_porter_v1_porter_app_proto_goTypes = []interface{}{
	(*PorterApp)(nil), // 0: porter.v1.PorterApp
	(*Build)(nil),     // 1: porter.v1.Build
	(*AppImage)(nil),  // 2: porter.v1.AppImage
	nil,               // 3: porter.v1.PorterApp.ServicesEntry
	nil,               // 4: porter.v1.PorterApp.EnvEntry
	(*Service)(nil),   // 5: porter.v1.Service
}
var file_porter_v1_porter_app_proto_depIdxs = []int32{
	3, // 0: porter.v1.PorterApp.services:type_name -> porter.v1.PorterApp.ServicesEntry
	4, // 1: porter.v1.PorterApp.env:type_name -> porter.v1.PorterApp.EnvEntry
	1, // 2: porter.v1.PorterApp.build:type_name -> porter.v1.Build
	5, // 3: porter.v1.PorterApp.predeploy:type_name -> porter.v1.Service
	2, // 4: porter.v1.PorterApp.image:type_name -> porter.v1.AppImage
	5, // 5: porter.v1.PorterApp.ServicesEntry.value:type_name -> porter.v1.Service
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_porter_v1_porter_app_proto_init() }
func file_porter_v1_porter_app_proto_init() {
	if File_porter_v1_porter_app_proto != nil {
		return
	}
	file_porter_v1_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_porter_v1_porter_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PorterApp); i {
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
		file_porter_v1_porter_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Build); i {
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
		file_porter_v1_porter_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppImage); i {
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
			RawDescriptor: file_porter_v1_porter_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_porter_v1_porter_app_proto_goTypes,
		DependencyIndexes: file_porter_v1_porter_app_proto_depIdxs,
		MessageInfos:      file_porter_v1_porter_app_proto_msgTypes,
	}.Build()
	File_porter_v1_porter_app_proto = out.File
	file_porter_v1_porter_app_proto_rawDesc = nil
	file_porter_v1_porter_app_proto_goTypes = nil
	file_porter_v1_porter_app_proto_depIdxs = nil
}
