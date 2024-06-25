// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: porter/v1/cloud_contract.proto

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

// CloudContract is a contract for all Porter-managed infrastructure within a project
type CloudContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// datastores is the list of datastores associated with the project
	Datastores []*ManagedDatastore `protobuf:"bytes,1,rep,name=datastores,proto3" json:"datastores,omitempty"`
	// addons is the list of addons associated with the project
	Addons []*CloudContractAddon `protobuf:"bytes,2,rep,name=addons,proto3" json:"addons,omitempty"`
	// compliance contains the cloud accounts in a project that have compliance enabled
	CloudAccounts []*CloudAccount `protobuf:"bytes,3,rep,name=cloud_accounts,json=cloudAccounts,proto3" json:"cloud_accounts,omitempty"`
}

func (x *CloudContract) Reset() {
	*x = CloudContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_cloud_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudContract) ProtoMessage() {}

func (x *CloudContract) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_cloud_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudContract.ProtoReflect.Descriptor instead.
func (*CloudContract) Descriptor() ([]byte, []int) {
	return file_porter_v1_cloud_contract_proto_rawDescGZIP(), []int{0}
}

func (x *CloudContract) GetDatastores() []*ManagedDatastore {
	if x != nil {
		return x.Datastores
	}
	return nil
}

func (x *CloudContract) GetAddons() []*CloudContractAddon {
	if x != nil {
		return x.Addons
	}
	return nil
}

func (x *CloudContract) GetCloudAccounts() []*CloudAccount {
	if x != nil {
		return x.CloudAccounts
	}
	return nil
}

// CloudAccount represents a single account in a cloud provider that is managed by Porter
type CloudAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the id of the cloud account
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// compliance_enabled is true if compliance is enabled for this cloud account
	ComplianceProfile *CloudAccountComplianceProfile `protobuf:"bytes,2,opt,name=compliance_profile,json=complianceProfile,proto3" json:"compliance_profile,omitempty"`
}

func (x *CloudAccount) Reset() {
	*x = CloudAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_cloud_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudAccount) ProtoMessage() {}

func (x *CloudAccount) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_cloud_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudAccount.ProtoReflect.Descriptor instead.
func (*CloudAccount) Descriptor() ([]byte, []int) {
	return file_porter_v1_cloud_contract_proto_rawDescGZIP(), []int{1}
}

func (x *CloudAccount) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CloudAccount) GetComplianceProfile() *CloudAccountComplianceProfile {
	if x != nil {
		return x.ComplianceProfile
	}
	return nil
}

// ComplianceProfiles are the different compliance profiles that can be enforced on a contract
type CloudAccountComplianceProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// soc2 indicates that the contract should be compliant with SOC2
	Soc2 bool `protobuf:"varint,1,opt,name=soc2,proto3" json:"soc2,omitempty"`
	// hipaa indicates that the contract should be compliant with HIPAA
	Hipaa bool `protobuf:"varint,2,opt,name=hipaa,proto3" json:"hipaa,omitempty"`
}

func (x *CloudAccountComplianceProfile) Reset() {
	*x = CloudAccountComplianceProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_cloud_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudAccountComplianceProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudAccountComplianceProfile) ProtoMessage() {}

func (x *CloudAccountComplianceProfile) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_cloud_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudAccountComplianceProfile.ProtoReflect.Descriptor instead.
func (*CloudAccountComplianceProfile) Descriptor() ([]byte, []int) {
	return file_porter_v1_cloud_contract_proto_rawDescGZIP(), []int{2}
}

func (x *CloudAccountComplianceProfile) GetSoc2() bool {
	if x != nil {
		return x.Soc2
	}
	return false
}

func (x *CloudAccountComplianceProfile) GetHipaa() bool {
	if x != nil {
		return x.Hipaa
	}
	return false
}

// CloudContractAddon is a contract for an addon managed by the cloud contract
type CloudContractAddon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cluster_id is the id of the cluster that the addon is deployed on
	ClusterId int32 `protobuf:"varint,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// addon is the addon configuration
	Addon *Addon `protobuf:"bytes,2,opt,name=addon,proto3" json:"addon,omitempty"`
}

func (x *CloudContractAddon) Reset() {
	*x = CloudContractAddon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_cloud_contract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudContractAddon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudContractAddon) ProtoMessage() {}

func (x *CloudContractAddon) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_cloud_contract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudContractAddon.ProtoReflect.Descriptor instead.
func (*CloudContractAddon) Descriptor() ([]byte, []int) {
	return file_porter_v1_cloud_contract_proto_rawDescGZIP(), []int{3}
}

func (x *CloudContractAddon) GetClusterId() int32 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

func (x *CloudContractAddon) GetAddon() *Addon {
	if x != nil {
		return x.Addon
	}
	return nil
}

// CloudContractRevision represents a cloud contract revision which should be reconciled
type CloudContractRevision struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId int32 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// revision_id is the id of the cloud contract revision that this message applies to
	RevisionId string `protobuf:"bytes,2,opt,name=revision_id,json=revisionId,proto3" json:"revision_id,omitempty"`
}

func (x *CloudContractRevision) Reset() {
	*x = CloudContractRevision{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_cloud_contract_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudContractRevision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudContractRevision) ProtoMessage() {}

func (x *CloudContractRevision) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_cloud_contract_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudContractRevision.ProtoReflect.Descriptor instead.
func (*CloudContractRevision) Descriptor() ([]byte, []int) {
	return file_porter_v1_cloud_contract_proto_rawDescGZIP(), []int{4}
}

func (x *CloudContractRevision) GetProjectId() int32 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *CloudContractRevision) GetRevisionId() string {
	if x != nil {
		return x.RevisionId
	}
	return ""
}

// CloudContractDeletionRevision represents a collection of resources that should be deleted
type CloudContractDeletionRevision struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId int32 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// cloud_contract_deletions includes all resources that should be deleted
	CloudContractDeletions *CloudContract `protobuf:"bytes,2,opt,name=cloud_contract_deletions,json=cloudContractDeletions,proto3" json:"cloud_contract_deletions,omitempty"`
}

func (x *CloudContractDeletionRevision) Reset() {
	*x = CloudContractDeletionRevision{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_cloud_contract_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudContractDeletionRevision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudContractDeletionRevision) ProtoMessage() {}

func (x *CloudContractDeletionRevision) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_cloud_contract_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudContractDeletionRevision.ProtoReflect.Descriptor instead.
func (*CloudContractDeletionRevision) Descriptor() ([]byte, []int) {
	return file_porter_v1_cloud_contract_proto_rawDescGZIP(), []int{5}
}

func (x *CloudContractDeletionRevision) GetProjectId() int32 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *CloudContractDeletionRevision) GetCloudContractDeletions() *CloudContract {
	if x != nil {
		return x.CloudContractDeletions
	}
	return nil
}

// MachineType is a virtual machine type
type MachineType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the name of the machine type
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *MachineType) Reset() {
	*x = MachineType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_cloud_contract_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineType) ProtoMessage() {}

func (x *MachineType) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_cloud_contract_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineType.ProtoReflect.Descriptor instead.
func (*MachineType) Descriptor() ([]byte, []int) {
	return file_porter_v1_cloud_contract_proto_rawDescGZIP(), []int{6}
}

func (x *MachineType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_porter_v1_cloud_contract_proto protoreflect.FileDescriptor

var file_porter_v1_cloud_contract_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x16, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3,
	0x01, 0x0a, 0x0d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x12, 0x3b, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x35, 0x0a,
	0x06, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x64,
	0x64, 0x6f, 0x6e, 0x73, 0x12, 0x3e, 0x0a, 0x0e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x22, 0x77, 0x0a, 0x0c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x57, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61,
	0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x11, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x49, 0x0a,
	0x1d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6f, 0x63, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x6f,
	0x63, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x69, 0x70, 0x61, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x68, 0x69, 0x70, 0x61, 0x61, 0x22, 0x5b, 0x0a, 0x12, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x6f, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a,
	0x05, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x6f, 0x6e, 0x52, 0x05,
	0x61, 0x64, 0x64, 0x6f, 0x6e, 0x22, 0x57, 0x0a, 0x15, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x92,
	0x01, 0x0a, 0x1d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x52, 0x0a, 0x18, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x16, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x21, 0x0a, 0x0b, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0xad, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f,
	0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x15, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x50, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_porter_v1_cloud_contract_proto_rawDescOnce sync.Once
	file_porter_v1_cloud_contract_proto_rawDescData = file_porter_v1_cloud_contract_proto_rawDesc
)

func file_porter_v1_cloud_contract_proto_rawDescGZIP() []byte {
	file_porter_v1_cloud_contract_proto_rawDescOnce.Do(func() {
		file_porter_v1_cloud_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_porter_v1_cloud_contract_proto_rawDescData)
	})
	return file_porter_v1_cloud_contract_proto_rawDescData
}

var file_porter_v1_cloud_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_porter_v1_cloud_contract_proto_goTypes = []any{
	(*CloudContract)(nil),                 // 0: porter.v1.CloudContract
	(*CloudAccount)(nil),                  // 1: porter.v1.CloudAccount
	(*CloudAccountComplianceProfile)(nil), // 2: porter.v1.CloudAccountComplianceProfile
	(*CloudContractAddon)(nil),            // 3: porter.v1.CloudContractAddon
	(*CloudContractRevision)(nil),         // 4: porter.v1.CloudContractRevision
	(*CloudContractDeletionRevision)(nil), // 5: porter.v1.CloudContractDeletionRevision
	(*MachineType)(nil),                   // 6: porter.v1.MachineType
	(*ManagedDatastore)(nil),              // 7: porter.v1.ManagedDatastore
	(*Addon)(nil),                         // 8: porter.v1.Addon
}
var file_porter_v1_cloud_contract_proto_depIdxs = []int32{
	7, // 0: porter.v1.CloudContract.datastores:type_name -> porter.v1.ManagedDatastore
	3, // 1: porter.v1.CloudContract.addons:type_name -> porter.v1.CloudContractAddon
	1, // 2: porter.v1.CloudContract.cloud_accounts:type_name -> porter.v1.CloudAccount
	2, // 3: porter.v1.CloudAccount.compliance_profile:type_name -> porter.v1.CloudAccountComplianceProfile
	8, // 4: porter.v1.CloudContractAddon.addon:type_name -> porter.v1.Addon
	0, // 5: porter.v1.CloudContractDeletionRevision.cloud_contract_deletions:type_name -> porter.v1.CloudContract
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_porter_v1_cloud_contract_proto_init() }
func file_porter_v1_cloud_contract_proto_init() {
	if File_porter_v1_cloud_contract_proto != nil {
		return
	}
	file_porter_v1_addons_proto_init()
	file_porter_v1_datastore_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_porter_v1_cloud_contract_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CloudContract); i {
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
		file_porter_v1_cloud_contract_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CloudAccount); i {
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
		file_porter_v1_cloud_contract_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CloudAccountComplianceProfile); i {
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
		file_porter_v1_cloud_contract_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CloudContractAddon); i {
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
		file_porter_v1_cloud_contract_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CloudContractRevision); i {
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
		file_porter_v1_cloud_contract_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CloudContractDeletionRevision); i {
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
		file_porter_v1_cloud_contract_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*MachineType); i {
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
			RawDescriptor: file_porter_v1_cloud_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_porter_v1_cloud_contract_proto_goTypes,
		DependencyIndexes: file_porter_v1_cloud_contract_proto_depIdxs,
		MessageInfos:      file_porter_v1_cloud_contract_proto_msgTypes,
	}.Build()
	File_porter_v1_cloud_contract_proto = out.File
	file_porter_v1_cloud_contract_proto_rawDesc = nil
	file_porter_v1_cloud_contract_proto_goTypes = nil
	file_porter_v1_cloud_contract_proto_depIdxs = nil
}
