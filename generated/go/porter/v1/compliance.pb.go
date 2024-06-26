// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: porter/v1/compliance.proto

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

type EnumComplianceProfile int32

const (
	EnumComplianceProfile_ENUM_COMPLIANCE_PROFILE_UNSPECIFIED EnumComplianceProfile = 0
	// ENUM_COMPLIANCE_PROFILE_SOC2 represents the SOC2 compliance profile.
	EnumComplianceProfile_ENUM_COMPLIANCE_PROFILE_SOC2 EnumComplianceProfile = 1
	// ENUM_COMPLIANCE_PROFILE_HIPAA represents the HIPAA compliance profile.
	EnumComplianceProfile_ENUM_COMPLIANCE_PROFILE_HIPAA EnumComplianceProfile = 2
)

// Enum value maps for EnumComplianceProfile.
var (
	EnumComplianceProfile_name = map[int32]string{
		0: "ENUM_COMPLIANCE_PROFILE_UNSPECIFIED",
		1: "ENUM_COMPLIANCE_PROFILE_SOC2",
		2: "ENUM_COMPLIANCE_PROFILE_HIPAA",
	}
	EnumComplianceProfile_value = map[string]int32{
		"ENUM_COMPLIANCE_PROFILE_UNSPECIFIED": 0,
		"ENUM_COMPLIANCE_PROFILE_SOC2":        1,
		"ENUM_COMPLIANCE_PROFILE_HIPAA":       2,
	}
)

func (x EnumComplianceProfile) Enum() *EnumComplianceProfile {
	p := new(EnumComplianceProfile)
	*p = x
	return p
}

func (x EnumComplianceProfile) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumComplianceProfile) Descriptor() protoreflect.EnumDescriptor {
	return file_porter_v1_compliance_proto_enumTypes[0].Descriptor()
}

func (EnumComplianceProfile) Type() protoreflect.EnumType {
	return &file_porter_v1_compliance_proto_enumTypes[0]
}

func (x EnumComplianceProfile) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumComplianceProfile.Descriptor instead.
func (EnumComplianceProfile) EnumDescriptor() ([]byte, []int) {
	return file_porter_v1_compliance_proto_rawDescGZIP(), []int{0}
}

type EnumComplianceVendor int32

const (
	EnumComplianceVendor_ENUM_COMPLIANCE_VENDOR_UNSPECIFIED EnumComplianceVendor = 0
	// ENUM_COMPLIANCE_VENDOR_VANTA signifies the compliance vendor Vanta.
	EnumComplianceVendor_ENUM_COMPLIANCE_VENDOR_VANTA EnumComplianceVendor = 1
	// ENUM_COMPLIANCE_VENDOR_ONE_LEET signifies the compliance vendor One Leet.
	EnumComplianceVendor_ENUM_COMPLIANCE_VENDOR_ONE_LEET EnumComplianceVendor = 2
)

// Enum value maps for EnumComplianceVendor.
var (
	EnumComplianceVendor_name = map[int32]string{
		0: "ENUM_COMPLIANCE_VENDOR_UNSPECIFIED",
		1: "ENUM_COMPLIANCE_VENDOR_VANTA",
		2: "ENUM_COMPLIANCE_VENDOR_ONE_LEET",
	}
	EnumComplianceVendor_value = map[string]int32{
		"ENUM_COMPLIANCE_VENDOR_UNSPECIFIED": 0,
		"ENUM_COMPLIANCE_VENDOR_VANTA":       1,
		"ENUM_COMPLIANCE_VENDOR_ONE_LEET":    2,
	}
)

func (x EnumComplianceVendor) Enum() *EnumComplianceVendor {
	p := new(EnumComplianceVendor)
	*p = x
	return p
}

func (x EnumComplianceVendor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumComplianceVendor) Descriptor() protoreflect.EnumDescriptor {
	return file_porter_v1_compliance_proto_enumTypes[1].Descriptor()
}

func (EnumComplianceVendor) Type() protoreflect.EnumType {
	return &file_porter_v1_compliance_proto_enumTypes[1]
}

func (x EnumComplianceVendor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumComplianceVendor.Descriptor instead.
func (EnumComplianceVendor) EnumDescriptor() ([]byte, []int) {
	return file_porter_v1_compliance_proto_rawDescGZIP(), []int{1}
}

type EnumComplianceCheckStatus int32

const (
	EnumComplianceCheckStatus_ENUM_COMPLIANCE_CHECK_STATUS_UNSPECIFIED EnumComplianceCheckStatus = 0
	// ENUM_COMPLIANCE_CHECK_STATUS_PASSED signifies the compliance check passed.
	EnumComplianceCheckStatus_ENUM_COMPLIANCE_CHECK_STATUS_PASSED EnumComplianceCheckStatus = 1
	// ENUM_COMPLIANCE_CHECK_STATUS_FAILED signifies the compliance check failed.
	EnumComplianceCheckStatus_ENUM_COMPLIANCE_CHECK_STATUS_FAILED EnumComplianceCheckStatus = 2
	// ENUM_COMPLIANCE_CHECK_STATUS_NOT_APPLICABLE signifies the compliance check was not applicable.
	EnumComplianceCheckStatus_ENUM_COMPLIANCE_CHECK_STATUS_NOT_APPLICABLE EnumComplianceCheckStatus = 3
	// ENUM_COMPLIANCE_CHECK_STATUS_PENDING signifies the compliance check is pending (e.g. because the resource is updating).
	EnumComplianceCheckStatus_ENUM_COMPLIANCE_CHECK_STATUS_PENDING EnumComplianceCheckStatus = 4
)

// Enum value maps for EnumComplianceCheckStatus.
var (
	EnumComplianceCheckStatus_name = map[int32]string{
		0: "ENUM_COMPLIANCE_CHECK_STATUS_UNSPECIFIED",
		1: "ENUM_COMPLIANCE_CHECK_STATUS_PASSED",
		2: "ENUM_COMPLIANCE_CHECK_STATUS_FAILED",
		3: "ENUM_COMPLIANCE_CHECK_STATUS_NOT_APPLICABLE",
		4: "ENUM_COMPLIANCE_CHECK_STATUS_PENDING",
	}
	EnumComplianceCheckStatus_value = map[string]int32{
		"ENUM_COMPLIANCE_CHECK_STATUS_UNSPECIFIED":    0,
		"ENUM_COMPLIANCE_CHECK_STATUS_PASSED":         1,
		"ENUM_COMPLIANCE_CHECK_STATUS_FAILED":         2,
		"ENUM_COMPLIANCE_CHECK_STATUS_NOT_APPLICABLE": 3,
		"ENUM_COMPLIANCE_CHECK_STATUS_PENDING":        4,
	}
)

func (x EnumComplianceCheckStatus) Enum() *EnumComplianceCheckStatus {
	p := new(EnumComplianceCheckStatus)
	*p = x
	return p
}

func (x EnumComplianceCheckStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumComplianceCheckStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_porter_v1_compliance_proto_enumTypes[2].Descriptor()
}

func (EnumComplianceCheckStatus) Type() protoreflect.EnumType {
	return &file_porter_v1_compliance_proto_enumTypes[2]
}

func (x EnumComplianceCheckStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumComplianceCheckStatus.Descriptor instead.
func (EnumComplianceCheckStatus) EnumDescriptor() ([]byte, []int) {
	return file_porter_v1_compliance_proto_rawDescGZIP(), []int{2}
}

// ContractComplianceCheckGroup representes a porter internal concept that represents some infrastructure level configuration
// that is expected to be in place for a contract to be considered compliant.
type ContractComplianceCheckGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the name of the compliance check group.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// status is the current status of the compliance check group.
	//
	// Deprecated: Marked as deprecated in porter/v1/compliance.proto.
	Status EnumComplianceCheckStatus `protobuf:"varint,2,opt,name=status,proto3,enum=porter.v1.EnumComplianceCheckStatus" json:"status,omitempty"`
	// message is an optional message indicating why the compliance check group is in the current state.
	//
	// Deprecated: Marked as deprecated in porter/v1/compliance.proto.
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// resources_statuses is a list of resources and their statuses
	ResourceStatuses []*ResourceComplianceStatus `protobuf:"bytes,4,rep,name=resource_statuses,json=resourceStatuses,proto3" json:"resource_statuses,omitempty"`
}

func (x *ContractComplianceCheckGroup) Reset() {
	*x = ContractComplianceCheckGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_compliance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractComplianceCheckGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractComplianceCheckGroup) ProtoMessage() {}

func (x *ContractComplianceCheckGroup) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_compliance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractComplianceCheckGroup.ProtoReflect.Descriptor instead.
func (*ContractComplianceCheckGroup) Descriptor() ([]byte, []int) {
	return file_porter_v1_compliance_proto_rawDescGZIP(), []int{0}
}

func (x *ContractComplianceCheckGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Deprecated: Marked as deprecated in porter/v1/compliance.proto.
func (x *ContractComplianceCheckGroup) GetStatus() EnumComplianceCheckStatus {
	if x != nil {
		return x.Status
	}
	return EnumComplianceCheckStatus_ENUM_COMPLIANCE_CHECK_STATUS_UNSPECIFIED
}

// Deprecated: Marked as deprecated in porter/v1/compliance.proto.
func (x *ContractComplianceCheckGroup) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ContractComplianceCheckGroup) GetResourceStatuses() []*ResourceComplianceStatus {
	if x != nil {
		return x.ResourceStatuses
	}
	return nil
}

// ResourceComplianceCheck represents the status of a single resource for a particular compliance check
type ResourceComplianceStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the name of the resource.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// status is the current status of the resource.
	Status EnumComplianceCheckStatus `protobuf:"varint,2,opt,name=status,proto3,enum=porter.v1.EnumComplianceCheckStatus" json:"status,omitempty"`
	// message is an optional message indicating why the resource is in the current state.
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResourceComplianceStatus) Reset() {
	*x = ResourceComplianceStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_compliance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceComplianceStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceComplianceStatus) ProtoMessage() {}

func (x *ResourceComplianceStatus) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_compliance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceComplianceStatus.ProtoReflect.Descriptor instead.
func (*ResourceComplianceStatus) Descriptor() ([]byte, []int) {
	return file_porter_v1_compliance_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceComplianceStatus) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *ResourceComplianceStatus) GetStatus() EnumComplianceCheckStatus {
	if x != nil {
		return x.Status
	}
	return EnumComplianceCheckStatus_ENUM_COMPLIANCE_CHECK_STATUS_UNSPECIFIED
}

func (x *ResourceComplianceStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// VendorComplianceCheckGroup represents a vendor provided compliance test, which porter deems to be passing based on the status of the corresponding check group.
type VendorComplianceCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// description is the vendor provided description of the compliance check.
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// check_group is the name of the compliance check group that this check belongs to.
	CheckGroup string `protobuf:"bytes,2,opt,name=check_group,json=checkGroup,proto3" json:"check_group,omitempty"`
	// status is the status of the compliance check.
	// Deprecated: if cloud account id is set in the request, the statuses will be returned in the resource_statuses field.
	//
	// Deprecated: Marked as deprecated in porter/v1/compliance.proto.
	Status EnumComplianceCheckStatus `protobuf:"varint,3,opt,name=status,proto3,enum=porter.v1.EnumComplianceCheckStatus" json:"status,omitempty"`
	// reason is an optional message indicating why the compliance check is in the current state.
	// Deprecated: if cloud account id is set in the request, the statuses will be returned in the resource_statuses field.
	//
	// Deprecated: Marked as deprecated in porter/v1/compliance.proto.
	Reason string `protobuf:"bytes,4,opt,name=reason,proto3" json:"reason,omitempty"`
	// vendor_check_id is the unique identifier for the compliance check in the vendor's system.
	VendorCheckId string `protobuf:"bytes,5,opt,name=vendor_check_id,json=vendorCheckId,proto3" json:"vendor_check_id,omitempty"`
	// resource_statuses is a list of resources and their statuses
	ResourceStatuses []*ResourceComplianceStatus `protobuf:"bytes,6,rep,name=resource_statuses,json=resourceStatuses,proto3" json:"resource_statuses,omitempty"`
}

func (x *VendorComplianceCheck) Reset() {
	*x = VendorComplianceCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_compliance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VendorComplianceCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VendorComplianceCheck) ProtoMessage() {}

func (x *VendorComplianceCheck) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_compliance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VendorComplianceCheck.ProtoReflect.Descriptor instead.
func (*VendorComplianceCheck) Descriptor() ([]byte, []int) {
	return file_porter_v1_compliance_proto_rawDescGZIP(), []int{2}
}

func (x *VendorComplianceCheck) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *VendorComplianceCheck) GetCheckGroup() string {
	if x != nil {
		return x.CheckGroup
	}
	return ""
}

// Deprecated: Marked as deprecated in porter/v1/compliance.proto.
func (x *VendorComplianceCheck) GetStatus() EnumComplianceCheckStatus {
	if x != nil {
		return x.Status
	}
	return EnumComplianceCheckStatus_ENUM_COMPLIANCE_CHECK_STATUS_UNSPECIFIED
}

// Deprecated: Marked as deprecated in porter/v1/compliance.proto.
func (x *VendorComplianceCheck) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *VendorComplianceCheck) GetVendorCheckId() string {
	if x != nil {
		return x.VendorCheckId
	}
	return ""
}

func (x *VendorComplianceCheck) GetResourceStatuses() []*ResourceComplianceStatus {
	if x != nil {
		return x.ResourceStatuses
	}
	return nil
}

var File_porter_v1_compliance_proto protoreflect.FileDescriptor

var file_porter_v1_compliance_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0xe4, 0x01, 0x0a, 0x1c, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x42, 0x02, 0x18, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x50, 0x0a, 0x11,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x10, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x22, 0x97,
	0x01, 0x0a, 0x18, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x69, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x24, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x75,
	0x6d, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xb2, 0x02, 0x0a, 0x15, 0x56, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x02, 0x18, 0x01, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x50, 0x0a, 0x11, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x69, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x10, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x2a, 0x85, 0x01,
	0x0a, 0x15, 0x45, 0x6e, 0x75, 0x6d, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x23, 0x45, 0x4e, 0x55, 0x4d, 0x5f,
	0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49,
	0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x20, 0x0a, 0x1c, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x41,
	0x4e, 0x43, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x53, 0x4f, 0x43, 0x32,
	0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c,
	0x49, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x48, 0x49,
	0x50, 0x41, 0x41, 0x10, 0x02, 0x2a, 0x85, 0x01, 0x0a, 0x14, 0x45, 0x6e, 0x75, 0x6d, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x26,
	0x0a, 0x22, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x43,
	0x45, 0x5f, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x43,
	0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52,
	0x5f, 0x56, 0x41, 0x4e, 0x54, 0x41, 0x10, 0x01, 0x12, 0x23, 0x0a, 0x1f, 0x45, 0x4e, 0x55, 0x4d,
	0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x56, 0x45, 0x4e, 0x44,
	0x4f, 0x52, 0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x4c, 0x45, 0x45, 0x54, 0x10, 0x02, 0x2a, 0xf6, 0x01,
	0x0a, 0x19, 0x45, 0x6e, 0x75, 0x6d, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x28, 0x45,
	0x4e, 0x55, 0x4d, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x43,
	0x48, 0x45, 0x43, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x27, 0x0a, 0x23, 0x45, 0x4e, 0x55,
	0x4d, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x43, 0x48, 0x45,
	0x43, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x27, 0x0a, 0x23, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c,
	0x49, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x2f, 0x0a, 0x2b, 0x45,
	0x4e, 0x55, 0x4d, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x43,
	0x48, 0x45, 0x43, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x28, 0x0a, 0x24,
	0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x43, 0x45, 0x5f,
	0x43, 0x48, 0x45, 0x43, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e,
	0x44, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x42, 0xaa, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69,
	0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2d, 0x64,
	0x65, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x15, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_porter_v1_compliance_proto_rawDescOnce sync.Once
	file_porter_v1_compliance_proto_rawDescData = file_porter_v1_compliance_proto_rawDesc
)

func file_porter_v1_compliance_proto_rawDescGZIP() []byte {
	file_porter_v1_compliance_proto_rawDescOnce.Do(func() {
		file_porter_v1_compliance_proto_rawDescData = protoimpl.X.CompressGZIP(file_porter_v1_compliance_proto_rawDescData)
	})
	return file_porter_v1_compliance_proto_rawDescData
}

var file_porter_v1_compliance_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_porter_v1_compliance_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_porter_v1_compliance_proto_goTypes = []any{
	(EnumComplianceProfile)(0),           // 0: porter.v1.EnumComplianceProfile
	(EnumComplianceVendor)(0),            // 1: porter.v1.EnumComplianceVendor
	(EnumComplianceCheckStatus)(0),       // 2: porter.v1.EnumComplianceCheckStatus
	(*ContractComplianceCheckGroup)(nil), // 3: porter.v1.ContractComplianceCheckGroup
	(*ResourceComplianceStatus)(nil),     // 4: porter.v1.ResourceComplianceStatus
	(*VendorComplianceCheck)(nil),        // 5: porter.v1.VendorComplianceCheck
}
var file_porter_v1_compliance_proto_depIdxs = []int32{
	2, // 0: porter.v1.ContractComplianceCheckGroup.status:type_name -> porter.v1.EnumComplianceCheckStatus
	4, // 1: porter.v1.ContractComplianceCheckGroup.resource_statuses:type_name -> porter.v1.ResourceComplianceStatus
	2, // 2: porter.v1.ResourceComplianceStatus.status:type_name -> porter.v1.EnumComplianceCheckStatus
	2, // 3: porter.v1.VendorComplianceCheck.status:type_name -> porter.v1.EnumComplianceCheckStatus
	4, // 4: porter.v1.VendorComplianceCheck.resource_statuses:type_name -> porter.v1.ResourceComplianceStatus
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_porter_v1_compliance_proto_init() }
func file_porter_v1_compliance_proto_init() {
	if File_porter_v1_compliance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_porter_v1_compliance_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ContractComplianceCheckGroup); i {
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
		file_porter_v1_compliance_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ResourceComplianceStatus); i {
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
		file_porter_v1_compliance_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*VendorComplianceCheck); i {
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
			RawDescriptor: file_porter_v1_compliance_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_porter_v1_compliance_proto_goTypes,
		DependencyIndexes: file_porter_v1_compliance_proto_depIdxs,
		EnumInfos:         file_porter_v1_compliance_proto_enumTypes,
		MessageInfos:      file_porter_v1_compliance_proto_msgTypes,
	}.Build()
	File_porter_v1_compliance_proto = out.File
	file_porter_v1_compliance_proto_rawDesc = nil
	file_porter_v1_compliance_proto_goTypes = nil
	file_porter_v1_compliance_proto_depIdxs = nil
}
