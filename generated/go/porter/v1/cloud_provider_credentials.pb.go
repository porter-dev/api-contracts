// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: porter/v1/cloud_provider_credentials.proto

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

type AWSCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chain []*AssumeRoleHop `protobuf:"bytes,1,rep,name=chain,proto3" json:"chain,omitempty"`
}

func (x *AWSCredentials) Reset() {
	*x = AWSCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_cloud_provider_credentials_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSCredentials) ProtoMessage() {}

func (x *AWSCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_cloud_provider_credentials_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSCredentials.ProtoReflect.Descriptor instead.
func (*AWSCredentials) Descriptor() ([]byte, []int) {
	return file_porter_v1_cloud_provider_credentials_proto_rawDescGZIP(), []int{0}
}

func (x *AWSCredentials) GetChain() []*AssumeRoleHop {
	if x != nil {
		return x.Chain
	}
	return nil
}

// AssumeRoleHop contains all information required for the source_arn to assume the target_arn
type AssumeRoleHop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceArn string `protobuf:"bytes,1,opt,name=source_arn,json=sourceArn,proto3" json:"source_arn,omitempty"`
	TargetArn string `protobuf:"bytes,2,opt,name=target_arn,json=targetArn,proto3" json:"target_arn,omitempty"`
	// external_id is optional and is used to prevent privilege escalation
	ExternalId string `protobuf:"bytes,3,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
}

func (x *AssumeRoleHop) Reset() {
	*x = AssumeRoleHop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_cloud_provider_credentials_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssumeRoleHop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssumeRoleHop) ProtoMessage() {}

func (x *AssumeRoleHop) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_cloud_provider_credentials_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssumeRoleHop.ProtoReflect.Descriptor instead.
func (*AssumeRoleHop) Descriptor() ([]byte, []int) {
	return file_porter_v1_cloud_provider_credentials_proto_rawDescGZIP(), []int{1}
}

func (x *AssumeRoleHop) GetSourceArn() string {
	if x != nil {
		return x.SourceArn
	}
	return ""
}

func (x *AssumeRoleHop) GetTargetArn() string {
	if x != nil {
		return x.TargetArn
	}
	return ""
}

func (x *AssumeRoleHop) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

type AzureCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId               string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	SubscriptionId         string `protobuf:"bytes,2,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	TenantId               string `protobuf:"bytes,3,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	ServicePrincipalSecret []byte `protobuf:"bytes,4,opt,name=service_principal_secret,json=servicePrincipalSecret,proto3" json:"service_principal_secret,omitempty"`
}

func (x *AzureCredentials) Reset() {
	*x = AzureCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_cloud_provider_credentials_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AzureCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzureCredentials) ProtoMessage() {}

func (x *AzureCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_cloud_provider_credentials_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzureCredentials.ProtoReflect.Descriptor instead.
func (*AzureCredentials) Descriptor() ([]byte, []int) {
	return file_porter_v1_cloud_provider_credentials_proto_rawDescGZIP(), []int{2}
}

func (x *AzureCredentials) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *AzureCredentials) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

func (x *AzureCredentials) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *AzureCredentials) GetServicePrincipalSecret() []byte {
	if x != nil {
		return x.ServicePrincipalSecret
	}
	return nil
}

type GCPCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// service_account_json_base64 is the base64 encoded service account json which can be used to authenticate with GCP
	ServiceAccountJsonBase64 string `protobuf:"bytes,1,opt,name=service_account_json_base64,json=serviceAccountJsonBase64,proto3" json:"service_account_json_base64,omitempty"`
}

func (x *GCPCredentials) Reset() {
	*x = GCPCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_cloud_provider_credentials_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCPCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCPCredentials) ProtoMessage() {}

func (x *GCPCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_cloud_provider_credentials_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCPCredentials.ProtoReflect.Descriptor instead.
func (*GCPCredentials) Descriptor() ([]byte, []int) {
	return file_porter_v1_cloud_provider_credentials_proto_rawDescGZIP(), []int{3}
}

func (x *GCPCredentials) GetServiceAccountJsonBase64() string {
	if x != nil {
		return x.ServiceAccountJsonBase64
	}
	return ""
}

type UpdateCloudProviderCredentialsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// project_id [REQUIRED] is the project that we are updating the credentials for
	ProjectId int64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// cloud_provider [REQUIRED] represents the provider that we will provisioning the cluster in
	CloudProvider EnumCloudProvider `protobuf:"varint,2,opt,name=cloud_provider,json=cloudProvider,proto3,enum=porter.v1.EnumCloudProvider" json:"cloud_provider,omitempty"`
	// cloud_provider_credentials are the credentials for the specified cloud provider
	//
	// Types that are assignable to CloudProviderCredentials:
	//
	//	*UpdateCloudProviderCredentialsRequest_AwsCredentials
	//	*UpdateCloudProviderCredentialsRequest_AzureCredentials
	//	*UpdateCloudProviderCredentialsRequest_GcpCredentials
	CloudProviderCredentials isUpdateCloudProviderCredentialsRequest_CloudProviderCredentials `protobuf_oneof:"cloud_provider_credentials"`
	// credentials_identifier is the identifier for the credentials that we are updating. This is required during update operations, otherwise a new credential will be created
	CredentialsIdentifier string `protobuf:"bytes,6,opt,name=credentials_identifier,json=credentialsIdentifier,proto3" json:"credentials_identifier,omitempty"`
}

func (x *UpdateCloudProviderCredentialsRequest) Reset() {
	*x = UpdateCloudProviderCredentialsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_cloud_provider_credentials_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCloudProviderCredentialsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCloudProviderCredentialsRequest) ProtoMessage() {}

func (x *UpdateCloudProviderCredentialsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_cloud_provider_credentials_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCloudProviderCredentialsRequest.ProtoReflect.Descriptor instead.
func (*UpdateCloudProviderCredentialsRequest) Descriptor() ([]byte, []int) {
	return file_porter_v1_cloud_provider_credentials_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCloudProviderCredentialsRequest) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *UpdateCloudProviderCredentialsRequest) GetCloudProvider() EnumCloudProvider {
	if x != nil {
		return x.CloudProvider
	}
	return EnumCloudProvider_ENUM_CLOUD_PROVIDER_UNSPECIFIED
}

func (m *UpdateCloudProviderCredentialsRequest) GetCloudProviderCredentials() isUpdateCloudProviderCredentialsRequest_CloudProviderCredentials {
	if m != nil {
		return m.CloudProviderCredentials
	}
	return nil
}

func (x *UpdateCloudProviderCredentialsRequest) GetAwsCredentials() *AWSCredentials {
	if x, ok := x.GetCloudProviderCredentials().(*UpdateCloudProviderCredentialsRequest_AwsCredentials); ok {
		return x.AwsCredentials
	}
	return nil
}

func (x *UpdateCloudProviderCredentialsRequest) GetAzureCredentials() *AzureCredentials {
	if x, ok := x.GetCloudProviderCredentials().(*UpdateCloudProviderCredentialsRequest_AzureCredentials); ok {
		return x.AzureCredentials
	}
	return nil
}

func (x *UpdateCloudProviderCredentialsRequest) GetGcpCredentials() *GCPCredentials {
	if x, ok := x.GetCloudProviderCredentials().(*UpdateCloudProviderCredentialsRequest_GcpCredentials); ok {
		return x.GcpCredentials
	}
	return nil
}

func (x *UpdateCloudProviderCredentialsRequest) GetCredentialsIdentifier() string {
	if x != nil {
		return x.CredentialsIdentifier
	}
	return ""
}

type isUpdateCloudProviderCredentialsRequest_CloudProviderCredentials interface {
	isUpdateCloudProviderCredentialsRequest_CloudProviderCredentials()
}

type UpdateCloudProviderCredentialsRequest_AwsCredentials struct {
	AwsCredentials *AWSCredentials `protobuf:"bytes,3,opt,name=aws_credentials,json=awsCredentials,proto3,oneof"`
}

type UpdateCloudProviderCredentialsRequest_AzureCredentials struct {
	AzureCredentials *AzureCredentials `protobuf:"bytes,4,opt,name=azure_credentials,json=azureCredentials,proto3,oneof"`
}

type UpdateCloudProviderCredentialsRequest_GcpCredentials struct {
	GcpCredentials *GCPCredentials `protobuf:"bytes,5,opt,name=gcp_credentials,json=gcpCredentials,proto3,oneof"`
}

func (*UpdateCloudProviderCredentialsRequest_AwsCredentials) isUpdateCloudProviderCredentialsRequest_CloudProviderCredentials() {
}

func (*UpdateCloudProviderCredentialsRequest_AzureCredentials) isUpdateCloudProviderCredentialsRequest_CloudProviderCredentials() {
}

func (*UpdateCloudProviderCredentialsRequest_GcpCredentials) isUpdateCloudProviderCredentialsRequest_CloudProviderCredentials() {
}

type UpdateCloudProviderCredentialsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId             int64  `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	CredentialsIdentifier string `protobuf:"bytes,2,opt,name=credentials_identifier,json=credentialsIdentifier,proto3" json:"credentials_identifier,omitempty"`
}

func (x *UpdateCloudProviderCredentialsResponse) Reset() {
	*x = UpdateCloudProviderCredentialsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_cloud_provider_credentials_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCloudProviderCredentialsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCloudProviderCredentialsResponse) ProtoMessage() {}

func (x *UpdateCloudProviderCredentialsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_cloud_provider_credentials_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCloudProviderCredentialsResponse.ProtoReflect.Descriptor instead.
func (*UpdateCloudProviderCredentialsResponse) Descriptor() ([]byte, []int) {
	return file_porter_v1_cloud_provider_credentials_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCloudProviderCredentialsResponse) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *UpdateCloudProviderCredentialsResponse) GetCredentialsIdentifier() string {
	if x != nil {
		return x.CredentialsIdentifier
	}
	return ""
}

var File_porter_v1_cloud_provider_credentials_proto protoreflect.FileDescriptor

var file_porter_v1_cloud_provider_credentials_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x40, 0x0a, 0x0e, 0x41, 0x57, 0x53, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x12, 0x2e, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73,
	0x73, 0x75, 0x6d, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x48, 0x6f, 0x70, 0x52, 0x05, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x22, 0x6e, 0x0a, 0x0d, 0x41, 0x73, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x48, 0x6f, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x72,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41,
	0x72, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x72, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x72,
	0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x49, 0x64, 0x22, 0xaf, 0x01, 0x0a, 0x10, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x18, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x16, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x22, 0x4f, 0x0a, 0x0e, 0x47, 0x43, 0x50, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x3d, 0x0a, 0x1b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x62,
	0x61, 0x73, 0x65, 0x36, 0x34, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4a, 0x73, 0x6f, 0x6e, 0x42,
	0x61, 0x73, 0x65, 0x36, 0x34, 0x22, 0xb8, 0x03, 0x0a, 0x25, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x43,
	0x0a, 0x0e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x52, 0x0d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x0f, 0x61, 0x77, 0x73, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x57, 0x53, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x48, 0x00, 0x52, 0x0e, 0x61, 0x77, 0x73, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x4a, 0x0a, 0x11, 0x61, 0x7a, 0x75,
	0x72, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x48, 0x00, 0x52, 0x10, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x44, 0x0a, 0x0f, 0x67, 0x63, 0x70, 0x5f, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x43, 0x50, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x48, 0x00, 0x52, 0x0e, 0x67, 0x63, 0x70,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x35, 0x0a, 0x16, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x42, 0x1c, 0x0a, 0x1a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x22, 0x7e, 0x0a, 0x26, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x16, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x42, 0xb8, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x42, 0x1d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02,
	0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x50, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0a, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_porter_v1_cloud_provider_credentials_proto_rawDescOnce sync.Once
	file_porter_v1_cloud_provider_credentials_proto_rawDescData = file_porter_v1_cloud_provider_credentials_proto_rawDesc
)

func file_porter_v1_cloud_provider_credentials_proto_rawDescGZIP() []byte {
	file_porter_v1_cloud_provider_credentials_proto_rawDescOnce.Do(func() {
		file_porter_v1_cloud_provider_credentials_proto_rawDescData = protoimpl.X.CompressGZIP(file_porter_v1_cloud_provider_credentials_proto_rawDescData)
	})
	return file_porter_v1_cloud_provider_credentials_proto_rawDescData
}

var file_porter_v1_cloud_provider_credentials_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_porter_v1_cloud_provider_credentials_proto_goTypes = []interface{}{
	(*AWSCredentials)(nil),                         // 0: porter.v1.AWSCredentials
	(*AssumeRoleHop)(nil),                          // 1: porter.v1.AssumeRoleHop
	(*AzureCredentials)(nil),                       // 2: porter.v1.AzureCredentials
	(*GCPCredentials)(nil),                         // 3: porter.v1.GCPCredentials
	(*UpdateCloudProviderCredentialsRequest)(nil),  // 4: porter.v1.UpdateCloudProviderCredentialsRequest
	(*UpdateCloudProviderCredentialsResponse)(nil), // 5: porter.v1.UpdateCloudProviderCredentialsResponse
	(EnumCloudProvider)(0),                         // 6: porter.v1.EnumCloudProvider
}
var file_porter_v1_cloud_provider_credentials_proto_depIdxs = []int32{
	1, // 0: porter.v1.AWSCredentials.chain:type_name -> porter.v1.AssumeRoleHop
	6, // 1: porter.v1.UpdateCloudProviderCredentialsRequest.cloud_provider:type_name -> porter.v1.EnumCloudProvider
	0, // 2: porter.v1.UpdateCloudProviderCredentialsRequest.aws_credentials:type_name -> porter.v1.AWSCredentials
	2, // 3: porter.v1.UpdateCloudProviderCredentialsRequest.azure_credentials:type_name -> porter.v1.AzureCredentials
	3, // 4: porter.v1.UpdateCloudProviderCredentialsRequest.gcp_credentials:type_name -> porter.v1.GCPCredentials
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_porter_v1_cloud_provider_credentials_proto_init() }
func file_porter_v1_cloud_provider_credentials_proto_init() {
	if File_porter_v1_cloud_provider_credentials_proto != nil {
		return
	}
	file_porter_v1_cluster_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_porter_v1_cloud_provider_credentials_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AWSCredentials); i {
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
		file_porter_v1_cloud_provider_credentials_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssumeRoleHop); i {
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
		file_porter_v1_cloud_provider_credentials_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AzureCredentials); i {
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
		file_porter_v1_cloud_provider_credentials_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCPCredentials); i {
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
		file_porter_v1_cloud_provider_credentials_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCloudProviderCredentialsRequest); i {
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
		file_porter_v1_cloud_provider_credentials_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCloudProviderCredentialsResponse); i {
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
	file_porter_v1_cloud_provider_credentials_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*UpdateCloudProviderCredentialsRequest_AwsCredentials)(nil),
		(*UpdateCloudProviderCredentialsRequest_AzureCredentials)(nil),
		(*UpdateCloudProviderCredentialsRequest_GcpCredentials)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_porter_v1_cloud_provider_credentials_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_porter_v1_cloud_provider_credentials_proto_goTypes,
		DependencyIndexes: file_porter_v1_cloud_provider_credentials_proto_depIdxs,
		MessageInfos:      file_porter_v1_cloud_provider_credentials_proto_msgTypes,
	}.Build()
	File_porter_v1_cloud_provider_credentials_proto = out.File
	file_porter_v1_cloud_provider_credentials_proto_rawDesc = nil
	file_porter_v1_cloud_provider_credentials_proto_goTypes = nil
	file_porter_v1_cloud_provider_credentials_proto_depIdxs = nil
}
