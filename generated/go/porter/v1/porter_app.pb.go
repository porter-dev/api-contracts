// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
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

	// name is the name of the application
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// services is a map of service names to service configurations
	Services map[string]*Service `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// env is deprecated in favor of env groups. It was a map of environment variable names to values
	//
	// Deprecated: Marked as deprecated in porter/v1/porter_app.proto.
	Env map[string]string `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// build is the build settings for the application
	Build *Build `protobuf:"bytes,4,opt,name=build,proto3" json:"build,omitempty"`
	// predeploy is a job service to run before deploying the application
	Predeploy *Service `protobuf:"bytes,5,opt,name=predeploy,proto3" json:"predeploy,omitempty"`
	// image is the image to use for a given revision of the application
	Image *AppImage `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
	// env_groups is a map of environment variable group names to environment variable group configurations
	EnvGroups []*EnvGroup `protobuf:"bytes,7,rep,name=env_groups,json=envGroups,proto3" json:"env_groups,omitempty"`
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

// Deprecated: Marked as deprecated in porter/v1/porter_app.proto.
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

func (x *PorterApp) GetEnvGroups() []*EnvGroup {
	if x != nil {
		return x.EnvGroups
	}
	return nil
}

// EnvGroup represents the metadata for an env group. We do not want to store the actual variables with the PorterApp.
type EnvGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the name of the environment variable group
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// version is the version of the environment variable group
	Version int64 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *EnvGroup) Reset() {
	*x = EnvGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_porter_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvGroup) ProtoMessage() {}

func (x *EnvGroup) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use EnvGroup.ProtoReflect.Descriptor instead.
func (*EnvGroup) Descriptor() ([]byte, []int) {
	return file_porter_v1_porter_app_proto_rawDescGZIP(), []int{1}
}

func (x *EnvGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EnvGroup) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

// EnvGroupVariables represents the variables for an env group.
type EnvGroupVariables struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// normal is a map of non-sensitive variable names to values
	Normal map[string]string `protobuf:"bytes,1,rep,name=normal,proto3" json:"normal,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// secret is a map of sensitive variable names to values
	Secret map[string]string `protobuf:"bytes,2,rep,name=secret,proto3" json:"secret,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EnvGroupVariables) Reset() {
	*x = EnvGroupVariables{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_porter_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvGroupVariables) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvGroupVariables) ProtoMessage() {}

func (x *EnvGroupVariables) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use EnvGroupVariables.ProtoReflect.Descriptor instead.
func (*EnvGroupVariables) Descriptor() ([]byte, []int) {
	return file_porter_v1_porter_app_proto_rawDescGZIP(), []int{2}
}

func (x *EnvGroupVariables) GetNormal() map[string]string {
	if x != nil {
		return x.Normal
	}
	return nil
}

func (x *EnvGroupVariables) GetSecret() map[string]string {
	if x != nil {
		return x.Secret
	}
	return nil
}

// Deletions contains all explicit deletions from a PorterApp
type Deletions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// service_names is a list of service names to delete
	ServiceNames []string `protobuf:"bytes,1,rep,name=service_names,json=serviceNames,proto3" json:"service_names,omitempty"`
	// env_group_names is a list of environment variable group names to delete
	EnvGroupNames []string `protobuf:"bytes,2,rep,name=env_group_names,json=envGroupNames,proto3" json:"env_group_names,omitempty"`
	// env_variable_names is deprecated in favor of env_group_names. It was a list of environment variable names to delete
	//
	// Deprecated: Marked as deprecated in porter/v1/porter_app.proto.
	EnvVariableNames []string `protobuf:"bytes,4,rep,name=env_variable_names,json=envVariableNames,proto3" json:"env_variable_names,omitempty"`
}

func (x *Deletions) Reset() {
	*x = Deletions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_porter_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deletions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deletions) ProtoMessage() {}

func (x *Deletions) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_porter_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deletions.ProtoReflect.Descriptor instead.
func (*Deletions) Descriptor() ([]byte, []int) {
	return file_porter_v1_porter_app_proto_rawDescGZIP(), []int{3}
}

func (x *Deletions) GetServiceNames() []string {
	if x != nil {
		return x.ServiceNames
	}
	return nil
}

func (x *Deletions) GetEnvGroupNames() []string {
	if x != nil {
		return x.EnvGroupNames
	}
	return nil
}

// Deprecated: Marked as deprecated in porter/v1/porter_app.proto.
func (x *Deletions) GetEnvVariableNames() []string {
	if x != nil {
		return x.EnvVariableNames
	}
	return nil
}

// Build is the build settings for the application
type Build struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// context is the path to the build context
	Context string `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	// method is the build method to use, being one of "pack", "docker", or "registry"
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	// builder is the builder to use for the "pack" build method
	Builder string `protobuf:"bytes,3,opt,name=builder,proto3" json:"builder,omitempty"`
	// buildpacks is a list of buildpacks to use for the "pack" build method
	Buildpacks []string `protobuf:"bytes,4,rep,name=buildpacks,proto3" json:"buildpacks,omitempty"`
	// dockerfile is the path to the Dockerfile to use for the "docker" build method
	Dockerfile string `protobuf:"bytes,5,opt,name=dockerfile,proto3" json:"dockerfile,omitempty"`
	// commit_sha is the commit SHA at which to build the application
	CommitSha string `protobuf:"bytes,6,opt,name=commit_sha,json=commitSha,proto3" json:"commit_sha,omitempty"`
}

func (x *Build) Reset() {
	*x = Build{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_porter_app_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Build) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Build) ProtoMessage() {}

func (x *Build) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_porter_app_proto_msgTypes[4]
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
	return file_porter_v1_porter_app_proto_rawDescGZIP(), []int{4}
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

func (x *Build) GetCommitSha() string {
	if x != nil {
		return x.CommitSha
	}
	return ""
}

// AppImage is the image to use for a given revision of the application
type AppImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// repository is the repository to use for the image
	Repository string `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	// tag is the tag to use for the image
	Tag string `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *AppImage) Reset() {
	*x = AppImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_porter_app_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppImage) ProtoMessage() {}

func (x *AppImage) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_porter_app_proto_msgTypes[5]
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
	return file_porter_v1_porter_app_proto_rawDescGZIP(), []int{5}
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
	0x22, 0xd6, 0x03, 0x0a, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x12, 0x33, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x41, 0x70, 0x70, 0x2e, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x02,
	0x18, 0x01, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x26, 0x0a, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x12,
	0x30, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x09, 0x70, 0x72, 0x65, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x0a,
	0x65, 0x6e, 0x76, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x76,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x09, 0x65, 0x6e, 0x76, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x1a, 0x4f, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x36, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x38, 0x0a, 0x08, 0x45, 0x6e, 0x76,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x8d, 0x02, 0x0a, 0x11, 0x45, 0x6e, 0x76, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x06, 0x6e, 0x6f, 0x72,
	0x6d, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x12, 0x40, 0x0a, 0x06, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x1a, 0x39, 0x0a,
	0x0b, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x8a, 0x01, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x6e, 0x76, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0d, 0x65, 0x6e, 0x76, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x30,
	0x0a, 0x12, 0x65, 0x6e, 0x76, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x10,
	0x65, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x22, 0xb2, 0x01, 0x0a, 0x05, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x70,
	0x61, 0x63, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x70, 0x61, 0x63, 0x6b, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x5f, 0x73, 0x68, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x53, 0x68, 0x61, 0x22, 0x3c, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x74, 0x61, 0x67, 0x42, 0xa9, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x61,
	0x70, 0x69, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50,
	0x58, 0x58, 0xaa, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x50, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_porter_v1_porter_app_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_porter_v1_porter_app_proto_goTypes = []interface{}{
	(*PorterApp)(nil),         // 0: porter.v1.PorterApp
	(*EnvGroup)(nil),          // 1: porter.v1.EnvGroup
	(*EnvGroupVariables)(nil), // 2: porter.v1.EnvGroupVariables
	(*Deletions)(nil),         // 3: porter.v1.Deletions
	(*Build)(nil),             // 4: porter.v1.Build
	(*AppImage)(nil),          // 5: porter.v1.AppImage
	nil,                       // 6: porter.v1.PorterApp.ServicesEntry
	nil,                       // 7: porter.v1.PorterApp.EnvEntry
	nil,                       // 8: porter.v1.EnvGroupVariables.NormalEntry
	nil,                       // 9: porter.v1.EnvGroupVariables.SecretEntry
	(*Service)(nil),           // 10: porter.v1.Service
}
var file_porter_v1_porter_app_proto_depIdxs = []int32{
	6,  // 0: porter.v1.PorterApp.services:type_name -> porter.v1.PorterApp.ServicesEntry
	7,  // 1: porter.v1.PorterApp.env:type_name -> porter.v1.PorterApp.EnvEntry
	4,  // 2: porter.v1.PorterApp.build:type_name -> porter.v1.Build
	10, // 3: porter.v1.PorterApp.predeploy:type_name -> porter.v1.Service
	5,  // 4: porter.v1.PorterApp.image:type_name -> porter.v1.AppImage
	1,  // 5: porter.v1.PorterApp.env_groups:type_name -> porter.v1.EnvGroup
	8,  // 6: porter.v1.EnvGroupVariables.normal:type_name -> porter.v1.EnvGroupVariables.NormalEntry
	9,  // 7: porter.v1.EnvGroupVariables.secret:type_name -> porter.v1.EnvGroupVariables.SecretEntry
	10, // 8: porter.v1.PorterApp.ServicesEntry.value:type_name -> porter.v1.Service
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
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
			switch v := v.(*EnvGroup); i {
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
			switch v := v.(*EnvGroupVariables); i {
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
		file_porter_v1_porter_app_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deletions); i {
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
		file_porter_v1_porter_app_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_porter_v1_porter_app_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   10,
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
