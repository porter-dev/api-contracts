// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: porter/v1/datastore.proto

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

type EnumDatastoreKind int32

const (
	EnumDatastoreKind_ENUM_DATASTORE_KIND_UNSPECIFIED      EnumDatastoreKind = 0
	EnumDatastoreKind_ENUM_DATASTORE_KIND_AWS_RDS_POSTGRES EnumDatastoreKind = 1
)

// Enum value maps for EnumDatastoreKind.
var (
	EnumDatastoreKind_name = map[int32]string{
		0: "ENUM_DATASTORE_KIND_UNSPECIFIED",
		1: "ENUM_DATASTORE_KIND_AWS_RDS_POSTGRES",
	}
	EnumDatastoreKind_value = map[string]int32{
		"ENUM_DATASTORE_KIND_UNSPECIFIED":      0,
		"ENUM_DATASTORE_KIND_AWS_RDS_POSTGRES": 1,
	}
)

func (x EnumDatastoreKind) Enum() *EnumDatastoreKind {
	p := new(EnumDatastoreKind)
	*p = x
	return p
}

func (x EnumDatastoreKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumDatastoreKind) Descriptor() protoreflect.EnumDescriptor {
	return file_porter_v1_datastore_proto_enumTypes[0].Descriptor()
}

func (EnumDatastoreKind) Type() protoreflect.EnumType {
	return &file_porter_v1_datastore_proto_enumTypes[0]
}

func (x EnumDatastoreKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumDatastoreKind.Descriptor instead.
func (EnumDatastoreKind) EnumDescriptor() ([]byte, []int) {
	return file_porter_v1_datastore_proto_rawDescGZIP(), []int{0}
}

// PorterDatastore is the specification for a Porter-managed datastore
type PorterDatastore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id represents the id of the datastore. This is required for update operations, but should be left blank when creating a datastore
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// cloud_provider represents the provider that the datastore is provisioned in
	CloudProvider EnumCloudProvider `protobuf:"varint,2,opt,name=cloud_provider,json=cloudProvider,proto3,enum=porter.v1.EnumCloudProvider" json:"cloud_provider,omitempty"`
	// cloud_provider_credential_identifier is the credential used to provision the datastore
	CloudProviderCredentialIdentifier string `protobuf:"bytes,3,opt,name=cloud_provider_credential_identifier,json=cloudProviderCredentialIdentifier,proto3" json:"cloud_provider_credential_identifier,omitempty"`
	// region is the region the datastore is provisioned in
	Region string `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	// name is the name of the datastore
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// kind represents the type of the datastore
	Kind EnumDatastoreKind `protobuf:"varint,6,opt,name=kind,proto3,enum=porter.v1.EnumDatastoreKind" json:"kind,omitempty"`
	// kind_values are the required values depending on kind
	//
	// Types that are assignable to KindValues:
	//
	//	*PorterDatastore_AwsRdsPostgresKind
	KindValues isPorterDatastore_KindValues `protobuf_oneof:"kind_values"`
}

func (x *PorterDatastore) Reset() {
	*x = PorterDatastore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_datastore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PorterDatastore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PorterDatastore) ProtoMessage() {}

func (x *PorterDatastore) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_datastore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PorterDatastore.ProtoReflect.Descriptor instead.
func (*PorterDatastore) Descriptor() ([]byte, []int) {
	return file_porter_v1_datastore_proto_rawDescGZIP(), []int{0}
}

func (x *PorterDatastore) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PorterDatastore) GetCloudProvider() EnumCloudProvider {
	if x != nil {
		return x.CloudProvider
	}
	return EnumCloudProvider_ENUM_CLOUD_PROVIDER_UNSPECIFIED
}

func (x *PorterDatastore) GetCloudProviderCredentialIdentifier() string {
	if x != nil {
		return x.CloudProviderCredentialIdentifier
	}
	return ""
}

func (x *PorterDatastore) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *PorterDatastore) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PorterDatastore) GetKind() EnumDatastoreKind {
	if x != nil {
		return x.Kind
	}
	return EnumDatastoreKind_ENUM_DATASTORE_KIND_UNSPECIFIED
}

func (m *PorterDatastore) GetKindValues() isPorterDatastore_KindValues {
	if m != nil {
		return m.KindValues
	}
	return nil
}

func (x *PorterDatastore) GetAwsRdsPostgresKind() *AwsRdsPostgres {
	if x, ok := x.GetKindValues().(*PorterDatastore_AwsRdsPostgresKind); ok {
		return x.AwsRdsPostgresKind
	}
	return nil
}

type isPorterDatastore_KindValues interface {
	isPorterDatastore_KindValues()
}

type PorterDatastore_AwsRdsPostgresKind struct {
	AwsRdsPostgresKind *AwsRdsPostgres `protobuf:"bytes,7,opt,name=aws_rds_postgres_kind,json=awsRdsPostgresKind,proto3,oneof"`
}

func (*PorterDatastore_AwsRdsPostgresKind) isPorterDatastore_KindValues() {}

type DatastorePasswordSecretRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the name of the secret
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// namespace is the namespace of the secret
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// key is the key of the password in the secret
	Key string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DatastorePasswordSecretRef) Reset() {
	*x = DatastorePasswordSecretRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_datastore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatastorePasswordSecretRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatastorePasswordSecretRef) ProtoMessage() {}

func (x *DatastorePasswordSecretRef) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_datastore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatastorePasswordSecretRef.ProtoReflect.Descriptor instead.
func (*DatastorePasswordSecretRef) Descriptor() ([]byte, []int) {
	return file_porter_v1_datastore_proto_rawDescGZIP(), []int{1}
}

func (x *DatastorePasswordSecretRef) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DatastorePasswordSecretRef) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DatastorePasswordSecretRef) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type AwsRdsPostgres struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// database_name is the name of the rds database
	DatabaseName string `protobuf:"bytes,1,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	// master_username is the username of the database
	MasterUsername string `protobuf:"bytes,2,opt,name=master_username,json=masterUsername,proto3" json:"master_username,omitempty"`
	// master_user_password_literal is the string value of the password; this is only used for creating the database password secret and is wiped when the contract is saved
	MasterUserPasswordLiteral string `protobuf:"bytes,3,opt,name=master_user_password_literal,json=masterUserPasswordLiteral,proto3" json:"master_user_password_literal,omitempty"`
	// master_user_password_secret_ref is the reference to the secret that stores the password of the database
	MasterUserPasswordSecretRef *DatastorePasswordSecretRef `protobuf:"bytes,4,opt,name=master_user_password_secret_ref,json=masterUserPasswordSecretRef,proto3" json:"master_user_password_secret_ref,omitempty"`
	// allocated_storage_gigabytes is the allocated storage of the database in gigabytes
	AllocatedStorageGigabytes int64 `protobuf:"varint,5,opt,name=allocated_storage_gigabytes,json=allocatedStorageGigabytes,proto3" json:"allocated_storage_gigabytes,omitempty"`
	// instance_class is the instance class of the database
	InstanceClass string `protobuf:"bytes,6,opt,name=instance_class,json=instanceClass,proto3" json:"instance_class,omitempty"`
}

func (x *AwsRdsPostgres) Reset() {
	*x = AwsRdsPostgres{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_datastore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AwsRdsPostgres) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsRdsPostgres) ProtoMessage() {}

func (x *AwsRdsPostgres) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_datastore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsRdsPostgres.ProtoReflect.Descriptor instead.
func (*AwsRdsPostgres) Descriptor() ([]byte, []int) {
	return file_porter_v1_datastore_proto_rawDescGZIP(), []int{2}
}

func (x *AwsRdsPostgres) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

func (x *AwsRdsPostgres) GetMasterUsername() string {
	if x != nil {
		return x.MasterUsername
	}
	return ""
}

func (x *AwsRdsPostgres) GetMasterUserPasswordLiteral() string {
	if x != nil {
		return x.MasterUserPasswordLiteral
	}
	return ""
}

func (x *AwsRdsPostgres) GetMasterUserPasswordSecretRef() *DatastorePasswordSecretRef {
	if x != nil {
		return x.MasterUserPasswordSecretRef
	}
	return nil
}

func (x *AwsRdsPostgres) GetAllocatedStorageGigabytes() int64 {
	if x != nil {
		return x.AllocatedStorageGigabytes
	}
	return 0
}

func (x *AwsRdsPostgres) GetInstanceClass() string {
	if x != nil {
		return x.InstanceClass
	}
	return ""
}

// UpdateDatastorePayload is used to send messages via nats to update a datastore
type UpdateDatastorePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// project_id is the id of the project that the datastore belongs to
	ProjectId int64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// datastore_id is the id of the datastore
	DatastoreId string `protobuf:"bytes,2,opt,name=datastore_id,json=datastoreId,proto3" json:"datastore_id,omitempty"`
}

func (x *UpdateDatastorePayload) Reset() {
	*x = UpdateDatastorePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_datastore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDatastorePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDatastorePayload) ProtoMessage() {}

func (x *UpdateDatastorePayload) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_datastore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDatastorePayload.ProtoReflect.Descriptor instead.
func (*UpdateDatastorePayload) Descriptor() ([]byte, []int) {
	return file_porter_v1_datastore_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateDatastorePayload) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *UpdateDatastorePayload) GetDatastoreId() string {
	if x != nil {
		return x.DatastoreId
	}
	return ""
}

// DatastoreCredential is used to connect to a datastore
type DatastoreCredential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// host is the datastore host
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// database_name is the name of the database
	DatabaseName string `protobuf:"bytes,2,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	// username is the username required to access the datastore
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	// password is the password required to access the datastore
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	// port is the port to connect to
	Port int64 `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *DatastoreCredential) Reset() {
	*x = DatastoreCredential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_datastore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatastoreCredential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatastoreCredential) ProtoMessage() {}

func (x *DatastoreCredential) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_datastore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatastoreCredential.ProtoReflect.Descriptor instead.
func (*DatastoreCredential) Descriptor() ([]byte, []int) {
	return file_porter_v1_datastore_proto_rawDescGZIP(), []int{4}
}

func (x *DatastoreCredential) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *DatastoreCredential) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

func (x *DatastoreCredential) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DatastoreCredential) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *DatastoreCredential) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_porter_v1_datastore_proto protoreflect.FileDescriptor

var file_porter_v1_datastore_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf4, 0x02, 0x0a, 0x0f, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x43, 0x0a, 0x0e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x0d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x4f, 0x0a, 0x24, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x21, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4b, 0x69, 0x6e,
	0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x4e, 0x0a, 0x15, 0x61, 0x77, 0x73, 0x5f, 0x72,
	0x64, 0x73, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x77, 0x73, 0x52, 0x64, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65,
	0x73, 0x48, 0x00, 0x52, 0x12, 0x61, 0x77, 0x73, 0x52, 0x64, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x67,
	0x72, 0x65, 0x73, 0x4b, 0x69, 0x6e, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x6b, 0x69, 0x6e, 0x64, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x60, 0x0a, 0x1a, 0x44, 0x61, 0x74, 0x61, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x52, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xf3, 0x02, 0x0a, 0x0e, 0x41, 0x77, 0x73,
	0x52, 0x64, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x1c, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x5f, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x19, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x6b, 0x0a, 0x1f, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x66, 0x52, 0x1b, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x66, 0x12, 0x3e, 0x0a, 0x1b, 0x61, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x67, 0x69, 0x67,
	0x61, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x19, 0x61, 0x6c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x47, 0x69,
	0x67, 0x61, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x22, 0x5a,
	0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x22, 0x9a, 0x01, 0x0a, 0x13, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x2a, 0x62, 0x0a, 0x11, 0x45, 0x6e, 0x75, 0x6d, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x23, 0x0a, 0x1f,
	0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x4b,
	0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x28, 0x0a, 0x24, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x53, 0x54,
	0x4f, 0x52, 0x45, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x41, 0x57, 0x53, 0x5f, 0x52, 0x44, 0x53,
	0x5f, 0x50, 0x4f, 0x53, 0x54, 0x47, 0x52, 0x45, 0x53, 0x10, 0x01, 0x42, 0xa9, 0x01, 0x0a, 0x0d,
	0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
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
	file_porter_v1_datastore_proto_rawDescOnce sync.Once
	file_porter_v1_datastore_proto_rawDescData = file_porter_v1_datastore_proto_rawDesc
)

func file_porter_v1_datastore_proto_rawDescGZIP() []byte {
	file_porter_v1_datastore_proto_rawDescOnce.Do(func() {
		file_porter_v1_datastore_proto_rawDescData = protoimpl.X.CompressGZIP(file_porter_v1_datastore_proto_rawDescData)
	})
	return file_porter_v1_datastore_proto_rawDescData
}

var file_porter_v1_datastore_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_porter_v1_datastore_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_porter_v1_datastore_proto_goTypes = []interface{}{
	(EnumDatastoreKind)(0),             // 0: porter.v1.EnumDatastoreKind
	(*PorterDatastore)(nil),            // 1: porter.v1.PorterDatastore
	(*DatastorePasswordSecretRef)(nil), // 2: porter.v1.DatastorePasswordSecretRef
	(*AwsRdsPostgres)(nil),             // 3: porter.v1.AwsRdsPostgres
	(*UpdateDatastorePayload)(nil),     // 4: porter.v1.UpdateDatastorePayload
	(*DatastoreCredential)(nil),        // 5: porter.v1.DatastoreCredential
	(EnumCloudProvider)(0),             // 6: porter.v1.EnumCloudProvider
}
var file_porter_v1_datastore_proto_depIdxs = []int32{
	6, // 0: porter.v1.PorterDatastore.cloud_provider:type_name -> porter.v1.EnumCloudProvider
	0, // 1: porter.v1.PorterDatastore.kind:type_name -> porter.v1.EnumDatastoreKind
	3, // 2: porter.v1.PorterDatastore.aws_rds_postgres_kind:type_name -> porter.v1.AwsRdsPostgres
	2, // 3: porter.v1.AwsRdsPostgres.master_user_password_secret_ref:type_name -> porter.v1.DatastorePasswordSecretRef
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_porter_v1_datastore_proto_init() }
func file_porter_v1_datastore_proto_init() {
	if File_porter_v1_datastore_proto != nil {
		return
	}
	file_porter_v1_cluster_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_porter_v1_datastore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PorterDatastore); i {
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
		file_porter_v1_datastore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatastorePasswordSecretRef); i {
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
		file_porter_v1_datastore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AwsRdsPostgres); i {
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
		file_porter_v1_datastore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDatastorePayload); i {
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
		file_porter_v1_datastore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatastoreCredential); i {
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
	file_porter_v1_datastore_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PorterDatastore_AwsRdsPostgresKind)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_porter_v1_datastore_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_porter_v1_datastore_proto_goTypes,
		DependencyIndexes: file_porter_v1_datastore_proto_depIdxs,
		EnumInfos:         file_porter_v1_datastore_proto_enumTypes,
		MessageInfos:      file_porter_v1_datastore_proto_msgTypes,
	}.Build()
	File_porter_v1_datastore_proto = out.File
	file_porter_v1_datastore_proto_rawDesc = nil
	file_porter_v1_datastore_proto_goTypes = nil
	file_porter_v1_datastore_proto_depIdxs = nil
}
