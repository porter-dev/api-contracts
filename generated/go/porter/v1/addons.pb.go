// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: porter/v1/addons.proto

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

type AddonType int32

const (
	// ADDON_TYPE_UNSPECIFIED is the default value
	AddonType_ADDON_TYPE_UNSPECIFIED AddonType = 0
	// ADDON_TYPE_POSTGRES is the postgres addon type
	AddonType_ADDON_TYPE_POSTGRES AddonType = 1
	// ADDON_TYPE_REDIS is the redis addon type
	AddonType_ADDON_TYPE_REDIS AddonType = 2
	// ADDON_TYPE_DATADOG is the datadog addon type
	AddonType_ADDON_TYPE_DATADOG AddonType = 3
)

// Enum value maps for AddonType.
var (
	AddonType_name = map[int32]string{
		0: "ADDON_TYPE_UNSPECIFIED",
		1: "ADDON_TYPE_POSTGRES",
		2: "ADDON_TYPE_REDIS",
		3: "ADDON_TYPE_DATADOG",
	}
	AddonType_value = map[string]int32{
		"ADDON_TYPE_UNSPECIFIED": 0,
		"ADDON_TYPE_POSTGRES":    1,
		"ADDON_TYPE_REDIS":       2,
		"ADDON_TYPE_DATADOG":     3,
	}
)

func (x AddonType) Enum() *AddonType {
	p := new(AddonType)
	*p = x
	return p
}

func (x AddonType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddonType) Descriptor() protoreflect.EnumDescriptor {
	return file_porter_v1_addons_proto_enumTypes[0].Descriptor()
}

func (AddonType) Type() protoreflect.EnumType {
	return &file_porter_v1_addons_proto_enumTypes[0]
}

func (x AddonType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddonType.Descriptor instead.
func (AddonType) EnumDescriptor() ([]byte, []int) {
	return file_porter_v1_addons_proto_rawDescGZIP(), []int{0}
}

// PrerequisiteAddon specifies an addon that must be installed before any apps can be installed
// the addon should be installed with the specified config
type PrerequisiteAddon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// commit_sha is the commit SHA of the addon
	CommitSha string `protobuf:"bytes,1,opt,name=commit_sha,json=commitSha,proto3" json:"commit_sha,omitempty"`
}

func (x *PrerequisiteAddon) Reset() {
	*x = PrerequisiteAddon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_addons_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrerequisiteAddon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrerequisiteAddon) ProtoMessage() {}

func (x *PrerequisiteAddon) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_addons_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrerequisiteAddon.ProtoReflect.Descriptor instead.
func (*PrerequisiteAddon) Descriptor() ([]byte, []int) {
	return file_porter_v1_addons_proto_rawDescGZIP(), []int{0}
}

func (x *PrerequisiteAddon) GetCommitSha() string {
	if x != nil {
		return x.CommitSha
	}
	return ""
}

// Addon is the configuration object for tooling or services that can be applied to the cluster alongside porter apps.
type Addon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the name of the addon
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// type is the type of type of addon
	Type AddonType `protobuf:"varint,2,opt,name=type,proto3,enum=porter.v1.AddonType" json:"type,omitempty"`
	// env_groups is a list of environment groups that can be applied to the addon
	EnvGroups []*EnvGroup `protobuf:"bytes,3,rep,name=env_groups,json=envGroups,proto3" json:"env_groups,omitempty"`
	// config is the addon-specific configuration
	//
	// Types that are assignable to Config:
	//
	//	*Addon_Postgres
	//	*Addon_Redis
	//	*Addon_Datadog
	Config isAddon_Config `protobuf_oneof:"config"`
}

func (x *Addon) Reset() {
	*x = Addon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_addons_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Addon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Addon) ProtoMessage() {}

func (x *Addon) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_addons_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Addon.ProtoReflect.Descriptor instead.
func (*Addon) Descriptor() ([]byte, []int) {
	return file_porter_v1_addons_proto_rawDescGZIP(), []int{1}
}

func (x *Addon) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Addon) GetType() AddonType {
	if x != nil {
		return x.Type
	}
	return AddonType_ADDON_TYPE_UNSPECIFIED
}

func (x *Addon) GetEnvGroups() []*EnvGroup {
	if x != nil {
		return x.EnvGroups
	}
	return nil
}

func (m *Addon) GetConfig() isAddon_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (x *Addon) GetPostgres() *Postgres {
	if x, ok := x.GetConfig().(*Addon_Postgres); ok {
		return x.Postgres
	}
	return nil
}

func (x *Addon) GetRedis() *Redis {
	if x, ok := x.GetConfig().(*Addon_Redis); ok {
		return x.Redis
	}
	return nil
}

func (x *Addon) GetDatadog() *Datadog {
	if x, ok := x.GetConfig().(*Addon_Datadog); ok {
		return x.Datadog
	}
	return nil
}

type isAddon_Config interface {
	isAddon_Config()
}

type Addon_Postgres struct {
	// postgres is the configuration for the postgres addon
	Postgres *Postgres `protobuf:"bytes,4,opt,name=postgres,proto3,oneof"`
}

type Addon_Redis struct {
	// redis is the configuration for the redis addon
	Redis *Redis `protobuf:"bytes,5,opt,name=redis,proto3,oneof"`
}

type Addon_Datadog struct {
	// datadog is the configuration for the datadog addon
	Datadog *Datadog `protobuf:"bytes,6,opt,name=datadog,proto3,oneof"`
}

func (*Addon_Postgres) isAddon_Config() {}

func (*Addon_Redis) isAddon_Config() {}

func (*Addon_Datadog) isAddon_Config() {}

// Postgres is the configuration for postgres
type Postgres struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cpu_cores is the number of CPU cores to allocate to the database
	CpuCores float32 `protobuf:"fixed32,1,opt,name=cpu_cores,json=cpuCores,proto3" json:"cpu_cores,omitempty"`
	// ram_megabytes is the amount of memory to allocate to the database
	RamMegabytes int32 `protobuf:"varint,2,opt,name=ram_megabytes,json=ramMegabytes,proto3" json:"ram_megabytes,omitempty"`
	// storage_gigabytes is the amount of storage to allocate to the database
	StorageGigabytes int32 `protobuf:"varint,3,opt,name=storage_gigabytes,json=storageGigabytes,proto3" json:"storage_gigabytes,omitempty"`
	// master_username is the username of the database
	MasterUsername *string `protobuf:"bytes,4,opt,name=master_username,json=masterUsername,proto3,oneof" json:"master_username,omitempty"`
	// master_user_password_literal is the string value of the password; this is only used for creating the datastore password secret and is wiped when the contract is saved
	MasterUserPasswordLiteral *string `protobuf:"bytes,5,opt,name=master_user_password_literal,json=masterUserPasswordLiteral,proto3,oneof" json:"master_user_password_literal,omitempty"`
}

func (x *Postgres) Reset() {
	*x = Postgres{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_addons_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Postgres) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Postgres) ProtoMessage() {}

func (x *Postgres) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_addons_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Postgres.ProtoReflect.Descriptor instead.
func (*Postgres) Descriptor() ([]byte, []int) {
	return file_porter_v1_addons_proto_rawDescGZIP(), []int{2}
}

func (x *Postgres) GetCpuCores() float32 {
	if x != nil {
		return x.CpuCores
	}
	return 0
}

func (x *Postgres) GetRamMegabytes() int32 {
	if x != nil {
		return x.RamMegabytes
	}
	return 0
}

func (x *Postgres) GetStorageGigabytes() int32 {
	if x != nil {
		return x.StorageGigabytes
	}
	return 0
}

func (x *Postgres) GetMasterUsername() string {
	if x != nil && x.MasterUsername != nil {
		return *x.MasterUsername
	}
	return ""
}

func (x *Postgres) GetMasterUserPasswordLiteral() string {
	if x != nil && x.MasterUserPasswordLiteral != nil {
		return *x.MasterUserPasswordLiteral
	}
	return ""
}

// Redis is the configuration for redis
type Redis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cpu_cores is the number of CPU cores to allocate to redis
	CpuCores float32 `protobuf:"fixed32,1,opt,name=cpu_cores,json=cpuCores,proto3" json:"cpu_cores,omitempty"`
	// ram_megabytes is the amount of memory to allocate to redis
	RamMegabytes int32 `protobuf:"varint,2,opt,name=ram_megabytes,json=ramMegabytes,proto3" json:"ram_megabytes,omitempty"`
	// storage_gigabytes is the amount of storage to allocate to redis
	StorageGigabytes int32 `protobuf:"varint,3,opt,name=storage_gigabytes,json=storageGigabytes,proto3" json:"storage_gigabytes,omitempty"`
	// master_user_password_literal is the string value of the password; this is only used for creating the datastore password secret and is wiped when the contract is saved
	MasterUserPasswordLiteral *string `protobuf:"bytes,4,opt,name=master_user_password_literal,json=masterUserPasswordLiteral,proto3,oneof" json:"master_user_password_literal,omitempty"`
}

func (x *Redis) Reset() {
	*x = Redis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_addons_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Redis) ProtoMessage() {}

func (x *Redis) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_addons_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Redis.ProtoReflect.Descriptor instead.
func (*Redis) Descriptor() ([]byte, []int) {
	return file_porter_v1_addons_proto_rawDescGZIP(), []int{3}
}

func (x *Redis) GetCpuCores() float32 {
	if x != nil {
		return x.CpuCores
	}
	return 0
}

func (x *Redis) GetRamMegabytes() int32 {
	if x != nil {
		return x.RamMegabytes
	}
	return 0
}

func (x *Redis) GetStorageGigabytes() int32 {
	if x != nil {
		return x.StorageGigabytes
	}
	return 0
}

func (x *Redis) GetMasterUserPasswordLiteral() string {
	if x != nil && x.MasterUserPasswordLiteral != nil {
		return *x.MasterUserPasswordLiteral
	}
	return ""
}

// Datadog is the configuration for Datadog
type Datadog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// site is the datadog site url
	Site *string `protobuf:"bytes,1,opt,name=site,proto3,oneof" json:"site,omitempty"`
	// api_key is the api key
	ApiKey *string `protobuf:"bytes,2,opt,name=api_key,json=apiKey,proto3,oneof" json:"api_key,omitempty"`
	// logging_enabled determines whether all container logs go to datadog
	LoggingEnabled *bool `protobuf:"varint,3,opt,name=logging_enabled,json=loggingEnabled,proto3,oneof" json:"logging_enabled,omitempty"`
	// dogstatsd_enabled determines whether dogstatsd is enabled
	DogstatsdEnabled *bool `protobuf:"varint,4,opt,name=dogstatsd_enabled,json=dogstatsdEnabled,proto3,oneof" json:"dogstatsd_enabled,omitempty"`
	// apm_enabled determines whether apm is enabled
	ApmEnabled *bool `protobuf:"varint,5,opt,name=apm_enabled,json=apmEnabled,proto3,oneof" json:"apm_enabled,omitempty"`
}

func (x *Datadog) Reset() {
	*x = Datadog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_addons_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Datadog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Datadog) ProtoMessage() {}

func (x *Datadog) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_addons_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Datadog.ProtoReflect.Descriptor instead.
func (*Datadog) Descriptor() ([]byte, []int) {
	return file_porter_v1_addons_proto_rawDescGZIP(), []int{4}
}

func (x *Datadog) GetSite() string {
	if x != nil && x.Site != nil {
		return *x.Site
	}
	return ""
}

func (x *Datadog) GetApiKey() string {
	if x != nil && x.ApiKey != nil {
		return *x.ApiKey
	}
	return ""
}

func (x *Datadog) GetLoggingEnabled() bool {
	if x != nil && x.LoggingEnabled != nil {
		return *x.LoggingEnabled
	}
	return false
}

func (x *Datadog) GetDogstatsdEnabled() bool {
	if x != nil && x.DogstatsdEnabled != nil {
		return *x.DogstatsdEnabled
	}
	return false
}

func (x *Datadog) GetApmEnabled() bool {
	if x != nil && x.ApmEnabled != nil {
		return *x.ApmEnabled
	}
	return false
}

var File_porter_v1_addons_proto protoreflect.FileDescriptor

var file_porter_v1_addons_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x32, 0x0a, 0x11, 0x50, 0x72, 0x65, 0x72, 0x65, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x65, 0x41,
	0x64, 0x64, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x73,
	0x68, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x53, 0x68, 0x61, 0x22, 0x90, 0x02, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x65,
	0x6e, 0x76, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x09, 0x65, 0x6e, 0x76, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12,
	0x31, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x48, 0x00, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72,
	0x65, 0x73, 0x12, 0x28, 0x0a, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x64, 0x69, 0x73, 0x48, 0x00, 0x52, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x12, 0x2e, 0x0a, 0x07,
	0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x64, 0x6f,
	0x67, 0x48, 0x00, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x42, 0x08, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xa2, 0x02, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x67,
	0x72, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x70, 0x75, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x63, 0x70, 0x75, 0x43, 0x6f, 0x72, 0x65, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x61, 0x6d, 0x5f, 0x6d, 0x65, 0x67, 0x61, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x61, 0x6d, 0x4d, 0x65, 0x67, 0x61,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x5f, 0x67, 0x69, 0x67, 0x61, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x47, 0x69, 0x67, 0x61, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x12, 0x2c, 0x0a, 0x0f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0e, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x44, 0x0a, 0x1c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x19, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x74, 0x65,
	0x72, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x1f, 0x0a, 0x1d, 0x5f, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x22, 0xdd, 0x01, 0x0a, 0x05,
	0x52, 0x65, 0x64, 0x69, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x70, 0x75, 0x5f, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x63, 0x70, 0x75, 0x43, 0x6f, 0x72,
	0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x61, 0x6d, 0x5f, 0x6d, 0x65, 0x67, 0x61, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x61, 0x6d, 0x4d, 0x65,
	0x67, 0x61, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x5f, 0x67, 0x69, 0x67, 0x61, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x47, 0x69, 0x67, 0x61, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x1c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x74,
	0x65, 0x72, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x19, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x1f, 0x0a, 0x1d, 0x5f, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x22, 0x95, 0x02, 0x0a, 0x07,
	0x44, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x12, 0x17, 0x0a, 0x04, 0x73, 0x69, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x73, 0x69, 0x74, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x1c, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x12, 0x2c,
	0x0a, 0x0f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x11,
	0x64, 0x6f, 0x67, 0x73, 0x74, 0x61, 0x74, 0x73, 0x64, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x03, 0x52, 0x10, 0x64, 0x6f, 0x67, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x24,
	0x0a, 0x0b, 0x61, 0x70, 0x6d, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x04, 0x52, 0x0a, 0x61, 0x70, 0x6d, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x69, 0x74, 0x65, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0x14, 0x0a,
	0x12, 0x5f, 0x64, 0x6f, 0x67, 0x73, 0x74, 0x61, 0x74, 0x73, 0x64, 0x5f, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x61, 0x70, 0x6d, 0x5f, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x2a, 0x6e, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x16, 0x41, 0x44, 0x44, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13,
	0x41, 0x44, 0x44, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4f, 0x53, 0x54, 0x47,
	0x52, 0x45, 0x53, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x44, 0x44, 0x4f, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x44, 0x49, 0x53, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x41,
	0x44, 0x44, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x44, 0x4f,
	0x47, 0x10, 0x03, 0x42, 0xa6, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x41, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2d,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x3b, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa,
	0x02, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x50, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_porter_v1_addons_proto_rawDescOnce sync.Once
	file_porter_v1_addons_proto_rawDescData = file_porter_v1_addons_proto_rawDesc
)

func file_porter_v1_addons_proto_rawDescGZIP() []byte {
	file_porter_v1_addons_proto_rawDescOnce.Do(func() {
		file_porter_v1_addons_proto_rawDescData = protoimpl.X.CompressGZIP(file_porter_v1_addons_proto_rawDescData)
	})
	return file_porter_v1_addons_proto_rawDescData
}

var file_porter_v1_addons_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_porter_v1_addons_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_porter_v1_addons_proto_goTypes = []interface{}{
	(AddonType)(0),            // 0: porter.v1.AddonType
	(*PrerequisiteAddon)(nil), // 1: porter.v1.PrerequisiteAddon
	(*Addon)(nil),             // 2: porter.v1.Addon
	(*Postgres)(nil),          // 3: porter.v1.Postgres
	(*Redis)(nil),             // 4: porter.v1.Redis
	(*Datadog)(nil),           // 5: porter.v1.Datadog
	(*EnvGroup)(nil),          // 6: porter.v1.EnvGroup
}
var file_porter_v1_addons_proto_depIdxs = []int32{
	0, // 0: porter.v1.Addon.type:type_name -> porter.v1.AddonType
	6, // 1: porter.v1.Addon.env_groups:type_name -> porter.v1.EnvGroup
	3, // 2: porter.v1.Addon.postgres:type_name -> porter.v1.Postgres
	4, // 3: porter.v1.Addon.redis:type_name -> porter.v1.Redis
	5, // 4: porter.v1.Addon.datadog:type_name -> porter.v1.Datadog
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_porter_v1_addons_proto_init() }
func file_porter_v1_addons_proto_init() {
	if File_porter_v1_addons_proto != nil {
		return
	}
	file_porter_v1_porter_app_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_porter_v1_addons_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrerequisiteAddon); i {
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
		file_porter_v1_addons_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Addon); i {
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
		file_porter_v1_addons_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Postgres); i {
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
		file_porter_v1_addons_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Redis); i {
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
		file_porter_v1_addons_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Datadog); i {
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
	file_porter_v1_addons_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Addon_Postgres)(nil),
		(*Addon_Redis)(nil),
		(*Addon_Datadog)(nil),
	}
	file_porter_v1_addons_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_porter_v1_addons_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_porter_v1_addons_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_porter_v1_addons_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_porter_v1_addons_proto_goTypes,
		DependencyIndexes: file_porter_v1_addons_proto_depIdxs,
		EnumInfos:         file_porter_v1_addons_proto_enumTypes,
		MessageInfos:      file_porter_v1_addons_proto_msgTypes,
	}.Build()
	File_porter_v1_addons_proto = out.File
	file_porter_v1_addons_proto_rawDesc = nil
	file_porter_v1_addons_proto_goTypes = nil
	file_porter_v1_addons_proto_depIdxs = nil
}
