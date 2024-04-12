// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: porter/v1/app_event_webhook.proto

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

// WebhookAppEventType is the app event type a webhook can be configured for
type WebhookAppEventType int32

const (
	// WEBHOOK_APP_EVENT_TYPE_UNSPECIFIED is the default value
	WebhookAppEventType_WEBHOOK_APP_EVENT_TYPE_UNSPECIFIED WebhookAppEventType = 0
	// WEBHOOK_APP_EVENT_TYPE_DEPLOY is when the webhook is configured for deploy events
	WebhookAppEventType_WEBHOOK_APP_EVENT_TYPE_DEPLOY WebhookAppEventType = 1
	// WEBHOOK_APP_EVENT_TYPE_BUILD is when the webhook is configured for build events
	WebhookAppEventType_WEBHOOK_APP_EVENT_TYPE_BUILD WebhookAppEventType = 2
	// WEBHOOK_APP_EVENT_TYPE_PREDEPLOYT is when the webhook is configured for predeploy events
	WebhookAppEventType_WEBHOOK_APP_EVENT_TYPE_PREDEPLOY WebhookAppEventType = 3
	// WEBHOOK_APP_EVENT_TYPE_INIT_DEPLOY is when the webhook is configured for initial deploy
	WebhookAppEventType_WEBHOOK_APP_EVENT_TYPE_INIT_DEPLOY WebhookAppEventType = 4
)

// Enum value maps for WebhookAppEventType.
var (
	WebhookAppEventType_name = map[int32]string{
		0: "WEBHOOK_APP_EVENT_TYPE_UNSPECIFIED",
		1: "WEBHOOK_APP_EVENT_TYPE_DEPLOY",
		2: "WEBHOOK_APP_EVENT_TYPE_BUILD",
		3: "WEBHOOK_APP_EVENT_TYPE_PREDEPLOY",
		4: "WEBHOOK_APP_EVENT_TYPE_INIT_DEPLOY",
	}
	WebhookAppEventType_value = map[string]int32{
		"WEBHOOK_APP_EVENT_TYPE_UNSPECIFIED": 0,
		"WEBHOOK_APP_EVENT_TYPE_DEPLOY":      1,
		"WEBHOOK_APP_EVENT_TYPE_BUILD":       2,
		"WEBHOOK_APP_EVENT_TYPE_PREDEPLOY":   3,
		"WEBHOOK_APP_EVENT_TYPE_INIT_DEPLOY": 4,
	}
)

func (x WebhookAppEventType) Enum() *WebhookAppEventType {
	p := new(WebhookAppEventType)
	*p = x
	return p
}

func (x WebhookAppEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WebhookAppEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_porter_v1_app_event_webhook_proto_enumTypes[0].Descriptor()
}

func (WebhookAppEventType) Type() protoreflect.EnumType {
	return &file_porter_v1_app_event_webhook_proto_enumTypes[0]
}

func (x WebhookAppEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WebhookAppEventType.Descriptor instead.
func (WebhookAppEventType) EnumDescriptor() ([]byte, []int) {
	return file_porter_v1_app_event_webhook_proto_rawDescGZIP(), []int{0}
}

// WebhookAppEventStatus is the app evnet status a webhook can be configured for
type WebhookAppEventStatus int32

const (
	// WEBHOOK_APP_EVENT_STATUS_UNSPECIFIED is the default value
	WebhookAppEventStatus_WEBHOOK_APP_EVENT_STATUS_UNSPECIFIED WebhookAppEventStatus = 0
	// WEBHOOK_APP_EVENT_STATUS_SUCCESS configures webhook for success events
	WebhookAppEventStatus_WEBHOOK_APP_EVENT_STATUS_SUCCESS WebhookAppEventStatus = 1
	// WEBHOOK_APP_EVENT_STATUS_FAILED configures whebhook for failed events
	WebhookAppEventStatus_WEBHOOK_APP_EVENT_STATUS_FAILED WebhookAppEventStatus = 2
	// WEBHOOK_APP_EVENT_STATUS_PROGRESSING configures webhook for progressing events
	WebhookAppEventStatus_WEBHOOK_APP_EVENT_STATUS_PROGRESSING WebhookAppEventStatus = 3
	// WEBHOOK_APP_EVENT_STATUS_PROGRESSING configures webhook for canceled events
	WebhookAppEventStatus_WEBHOOK_APP_EVENT_STATUS_CANCELED WebhookAppEventStatus = 4
)

// Enum value maps for WebhookAppEventStatus.
var (
	WebhookAppEventStatus_name = map[int32]string{
		0: "WEBHOOK_APP_EVENT_STATUS_UNSPECIFIED",
		1: "WEBHOOK_APP_EVENT_STATUS_SUCCESS",
		2: "WEBHOOK_APP_EVENT_STATUS_FAILED",
		3: "WEBHOOK_APP_EVENT_STATUS_PROGRESSING",
		4: "WEBHOOK_APP_EVENT_STATUS_CANCELED",
	}
	WebhookAppEventStatus_value = map[string]int32{
		"WEBHOOK_APP_EVENT_STATUS_UNSPECIFIED": 0,
		"WEBHOOK_APP_EVENT_STATUS_SUCCESS":     1,
		"WEBHOOK_APP_EVENT_STATUS_FAILED":      2,
		"WEBHOOK_APP_EVENT_STATUS_PROGRESSING": 3,
		"WEBHOOK_APP_EVENT_STATUS_CANCELED":    4,
	}
)

func (x WebhookAppEventStatus) Enum() *WebhookAppEventStatus {
	p := new(WebhookAppEventStatus)
	*p = x
	return p
}

func (x WebhookAppEventStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WebhookAppEventStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_porter_v1_app_event_webhook_proto_enumTypes[1].Descriptor()
}

func (WebhookAppEventStatus) Type() protoreflect.EnumType {
	return &file_porter_v1_app_event_webhook_proto_enumTypes[1]
}

func (x WebhookAppEventStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WebhookAppEventStatus.Descriptor instead.
func (WebhookAppEventStatus) EnumDescriptor() ([]byte, []int) {
	return file_porter_v1_app_event_webhook_proto_rawDescGZIP(), []int{1}
}

// AppEventWebhook is a webhook configured to have requests sent to on app events
type AppEventWebhook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is a unique identifier for a webhook
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// webhook_url is the url of the webhook
	WebhookUrl []byte `protobuf:"bytes,2,opt,name=webhook_url,json=webhookUrl,proto3" json:"webhook_url,omitempty"`
	// app_event_type is the type of the app event this webhook is being added for
	AppEventType WebhookAppEventType `protobuf:"varint,3,opt,name=app_event_type,json=appEventType,proto3,enum=porter.v1.WebhookAppEventType" json:"app_event_type,omitempty"`
	// app_event_status is the status of the app event this webhook is being added for
	AppEventStatus WebhookAppEventStatus `protobuf:"varint,4,opt,name=app_event_status,json=appEventStatus,proto3,enum=porter.v1.WebhookAppEventStatus" json:"app_event_status,omitempty"`
}

func (x *AppEventWebhook) Reset() {
	*x = AppEventWebhook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_app_event_webhook_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppEventWebhook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppEventWebhook) ProtoMessage() {}

func (x *AppEventWebhook) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_app_event_webhook_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppEventWebhook.ProtoReflect.Descriptor instead.
func (*AppEventWebhook) Descriptor() ([]byte, []int) {
	return file_porter_v1_app_event_webhook_proto_rawDescGZIP(), []int{0}
}

func (x *AppEventWebhook) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AppEventWebhook) GetWebhookUrl() []byte {
	if x != nil {
		return x.WebhookUrl
	}
	return nil
}

func (x *AppEventWebhook) GetAppEventType() WebhookAppEventType {
	if x != nil {
		return x.AppEventType
	}
	return WebhookAppEventType_WEBHOOK_APP_EVENT_TYPE_UNSPECIFIED
}

func (x *AppEventWebhook) GetAppEventStatus() WebhookAppEventStatus {
	if x != nil {
		return x.AppEventStatus
	}
	return WebhookAppEventStatus_WEBHOOK_APP_EVENT_STATUS_UNSPECIFIED
}

var File_porter_v1_app_event_webhook_proto protoreflect.FileDescriptor

var file_porter_v1_app_event_webhook_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0xd4,
	0x01, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x57, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x55, 0x72, 0x6c, 0x12, 0x44, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x41,
	0x70, 0x70, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x61, 0x70, 0x70,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4a, 0x0a, 0x10, 0x61, 0x70, 0x70,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x41, 0x70, 0x70, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0e, 0x61, 0x70, 0x70, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0xd0, 0x01, 0x0a, 0x13, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f,
	0x6b, 0x41, 0x70, 0x70, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a,
	0x22, 0x57, 0x45, 0x42, 0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x57, 0x45, 0x42, 0x48, 0x4f, 0x4f, 0x4b,
	0x5f, 0x41, 0x50, 0x50, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x57, 0x45, 0x42, 0x48,
	0x4f, 0x4f, 0x4b, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x57, 0x45,
	0x42, 0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x45, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x10, 0x03,
	0x12, 0x26, 0x0a, 0x22, 0x57, 0x45, 0x42, 0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x41, 0x50, 0x50, 0x5f,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x5f,
	0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x10, 0x04, 0x2a, 0xdd, 0x01, 0x0a, 0x15, 0x57, 0x65, 0x62,
	0x68, 0x6f, 0x6f, 0x6b, 0x41, 0x70, 0x70, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x28, 0x0a, 0x24, 0x57, 0x45, 0x42, 0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x41, 0x50,
	0x50, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x20,
	0x57, 0x45, 0x42, 0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x10, 0x01, 0x12, 0x23, 0x0a, 0x1f, 0x57, 0x45, 0x42, 0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x41, 0x50,
	0x50, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x28, 0x0a, 0x24, 0x57, 0x45, 0x42, 0x48, 0x4f,
	0x4f, 0x4b, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10,
	0x03, 0x12, 0x25, 0x0a, 0x21, 0x57, 0x45, 0x42, 0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x41, 0x50, 0x50,
	0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x41,
	0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x42, 0xaf, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d,
	0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x41, 0x70, 0x70, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x09,
	0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a,
	0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_porter_v1_app_event_webhook_proto_rawDescOnce sync.Once
	file_porter_v1_app_event_webhook_proto_rawDescData = file_porter_v1_app_event_webhook_proto_rawDesc
)

func file_porter_v1_app_event_webhook_proto_rawDescGZIP() []byte {
	file_porter_v1_app_event_webhook_proto_rawDescOnce.Do(func() {
		file_porter_v1_app_event_webhook_proto_rawDescData = protoimpl.X.CompressGZIP(file_porter_v1_app_event_webhook_proto_rawDescData)
	})
	return file_porter_v1_app_event_webhook_proto_rawDescData
}

var file_porter_v1_app_event_webhook_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_porter_v1_app_event_webhook_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_porter_v1_app_event_webhook_proto_goTypes = []interface{}{
	(WebhookAppEventType)(0),   // 0: porter.v1.WebhookAppEventType
	(WebhookAppEventStatus)(0), // 1: porter.v1.WebhookAppEventStatus
	(*AppEventWebhook)(nil),    // 2: porter.v1.AppEventWebhook
}
var file_porter_v1_app_event_webhook_proto_depIdxs = []int32{
	0, // 0: porter.v1.AppEventWebhook.app_event_type:type_name -> porter.v1.WebhookAppEventType
	1, // 1: porter.v1.AppEventWebhook.app_event_status:type_name -> porter.v1.WebhookAppEventStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_porter_v1_app_event_webhook_proto_init() }
func file_porter_v1_app_event_webhook_proto_init() {
	if File_porter_v1_app_event_webhook_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_porter_v1_app_event_webhook_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppEventWebhook); i {
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
			RawDescriptor: file_porter_v1_app_event_webhook_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_porter_v1_app_event_webhook_proto_goTypes,
		DependencyIndexes: file_porter_v1_app_event_webhook_proto_depIdxs,
		EnumInfos:         file_porter_v1_app_event_webhook_proto_enumTypes,
		MessageInfos:      file_porter_v1_app_event_webhook_proto_msgTypes,
	}.Build()
	File_porter_v1_app_event_webhook_proto = out.File
	file_porter_v1_app_event_webhook_proto_rawDesc = nil
	file_porter_v1_app_event_webhook_proto_goTypes = nil
	file_porter_v1_app_event_webhook_proto_depIdxs = nil
}
