// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: porter/v1/system_service.proto

package porterv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Status represents the health status of a service
type Status int32

const (
	// STATUS_UNSPECIFIED is the default value
	Status_STATUS_UNSPECIFIED Status = 0
	// STATUS_HEALTHY is the value for a fully healthy service
	Status_STATUS_HEALTHY Status = 1
	// STATUS_PARTIAL_FAILURE is the health status for a partially failed service
	// a service is in partial failure if only less than the max unavailable number of replicas declared on the service are failed
	Status_STATUS_PARTIAL_FAILURE Status = 2
	// STATUS_FAILURE is the health status for a fully failed service
	Status_STATUS_FAILURE Status = 3
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_HEALTHY",
		2: "STATUS_PARTIAL_FAILURE",
		3: "STATUS_FAILURE",
	}
	Status_value = map[string]int32{
		"STATUS_UNSPECIFIED":     0,
		"STATUS_HEALTHY":         1,
		"STATUS_PARTIAL_FAILURE": 2,
		"STATUS_FAILURE":         3,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_porter_v1_system_service_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_porter_v1_system_service_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_porter_v1_system_service_proto_rawDescGZIP(), []int{0}
}

// ClusterHealthType represents types of health history we generate for the cluster
type ClusterHealthType int32

const (
	// CLUSTER_HEALTH_TYPE_UNSPECIFIED is the default value
	ClusterHealthType_CLUSTER_HEALTH_TYPE_UNSPECIFIED ClusterHealthType = 0
	// CLUSTER_HEALTH_TYPE_PINGABLE is history of whether we were able to ping the api-server
	ClusterHealthType_CLUSTER_HEALTH_TYPE_CONNECTED ClusterHealthType = 1
	// CLUSTER_HEALTH_TYPE_CONNECTED is history of weather we were able to connect to the cluster
	ClusterHealthType_CLUSTER_HEALTH_TYPE_PINGABLE ClusterHealthType = 2
	// CLUSTER_HEALTH_TYPE_METRICS_HEALTHY is whether the metrics services in the cluster were healthy
	ClusterHealthType_CLUSTER_HEALTH_TYPE_METRICS_HEALTHY ClusterHealthType = 3
)

// Enum value maps for ClusterHealthType.
var (
	ClusterHealthType_name = map[int32]string{
		0: "CLUSTER_HEALTH_TYPE_UNSPECIFIED",
		1: "CLUSTER_HEALTH_TYPE_CONNECTED",
		2: "CLUSTER_HEALTH_TYPE_PINGABLE",
		3: "CLUSTER_HEALTH_TYPE_METRICS_HEALTHY",
	}
	ClusterHealthType_value = map[string]int32{
		"CLUSTER_HEALTH_TYPE_UNSPECIFIED":     0,
		"CLUSTER_HEALTH_TYPE_CONNECTED":       1,
		"CLUSTER_HEALTH_TYPE_PINGABLE":        2,
		"CLUSTER_HEALTH_TYPE_METRICS_HEALTHY": 3,
	}
)

func (x ClusterHealthType) Enum() *ClusterHealthType {
	p := new(ClusterHealthType)
	*p = x
	return p
}

func (x ClusterHealthType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClusterHealthType) Descriptor() protoreflect.EnumDescriptor {
	return file_porter_v1_system_service_proto_enumTypes[1].Descriptor()
}

func (ClusterHealthType) Type() protoreflect.EnumType {
	return &file_porter_v1_system_service_proto_enumTypes[1]
}

func (x ClusterHealthType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClusterHealthType.Descriptor instead.
func (ClusterHealthType) EnumDescriptor() ([]byte, []int) {
	return file_porter_v1_system_service_proto_rawDescGZIP(), []int{1}
}

// SystemService identifies a system service by its name, namespace and the type of k8s object it runs
type SystemService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the name of the service
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// namespace is the namespace where it is deployed
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// involved_object_type is the k8s object type the service runs
	InvolvedObjectType InvolvedObjectType `protobuf:"varint,3,opt,name=involved_object_type,json=involvedObjectType,proto3,enum=porter.v1.InvolvedObjectType" json:"involved_object_type,omitempty"`
}

func (x *SystemService) Reset() {
	*x = SystemService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_system_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemService) ProtoMessage() {}

func (x *SystemService) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_system_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemService.ProtoReflect.Descriptor instead.
func (*SystemService) Descriptor() ([]byte, []int) {
	return file_porter_v1_system_service_proto_rawDescGZIP(), []int{0}
}

func (x *SystemService) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SystemService) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *SystemService) GetInvolvedObjectType() InvolvedObjectType {
	if x != nil {
		return x.InvolvedObjectType
	}
	return InvolvedObjectType_INVOLVED_OBJECT_TYPE_UNSPECIFIED
}

// HealthStatus is a single status observed over a certain period of time
type HealthStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// start_time holds the timestamp for when this status began
	StartTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// end_time holds the timestamp for when this status ended if provided
	// if not provided, the status is considered ongoing
	EndTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3,oneof" json:"end_time,omitempty"`
	// status gives the status for a service
	Status Status `protobuf:"varint,3,opt,name=status,proto3,enum=porter.v1.Status" json:"status,omitempty"`
	// description provides more information on a status
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *HealthStatus) Reset() {
	*x = HealthStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_system_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthStatus) ProtoMessage() {}

func (x *HealthStatus) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_system_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthStatus.ProtoReflect.Descriptor instead.
func (*HealthStatus) Descriptor() ([]byte, []int) {
	return file_porter_v1_system_service_proto_rawDescGZIP(), []int{1}
}

func (x *HealthStatus) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *HealthStatus) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *HealthStatus) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STATUS_UNSPECIFIED
}

func (x *HealthStatus) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// StatusPercentage is the percentage in a day that a certain status was observed
type StatusPercentage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     Status  `protobuf:"varint,1,opt,name=status,proto3,enum=porter.v1.Status" json:"status,omitempty"`
	Percentage float32 `protobuf:"fixed32,2,opt,name=percentage,proto3" json:"percentage,omitempty"`
}

func (x *StatusPercentage) Reset() {
	*x = StatusPercentage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_system_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusPercentage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusPercentage) ProtoMessage() {}

func (x *StatusPercentage) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_system_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusPercentage.ProtoReflect.Descriptor instead.
func (*StatusPercentage) Descriptor() ([]byte, []int) {
	return file_porter_v1_system_service_proto_rawDescGZIP(), []int{2}
}

func (x *StatusPercentage) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STATUS_UNSPECIFIED
}

func (x *StatusPercentage) GetPercentage() float32 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

// DailyHealthStatus represents health status over a day
type DailyHealthStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// status_percentages holds information on what percentage of the day different statuses were observed
	// there should be only one entry for the different Status types
	StatusPercentages []*StatusPercentage `protobuf:"bytes,1,rep,name=status_percentages,json=statusPercentages,proto3" json:"status_percentages,omitempty"`
	// health_statuses is the different statuses observed over a day
	// if there is a missing time gap, the service can be considered to have been healthy during that gap
	HealthStatuses []*HealthStatus `protobuf:"bytes,2,rep,name=health_statuses,json=healthStatuses,proto3" json:"health_statuses,omitempty"`
}

func (x *DailyHealthStatus) Reset() {
	*x = DailyHealthStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_system_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailyHealthStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyHealthStatus) ProtoMessage() {}

func (x *DailyHealthStatus) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_system_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyHealthStatus.ProtoReflect.Descriptor instead.
func (*DailyHealthStatus) Descriptor() ([]byte, []int) {
	return file_porter_v1_system_service_proto_rawDescGZIP(), []int{3}
}

func (x *DailyHealthStatus) GetStatusPercentages() []*StatusPercentage {
	if x != nil {
		return x.StatusPercentages
	}
	return nil
}

func (x *DailyHealthStatus) GetHealthStatuses() []*HealthStatus {
	if x != nil {
		return x.HealthStatuses
	}
	return nil
}

// SystemServiceStatusHistory holds the health status history for a particular system service over multiple days
type SystemServiceStatusHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// system_service identifies the service
	SystemService *SystemService `protobuf:"bytes,1,opt,name=system_service,json=systemService,proto3" json:"system_service,omitempty"`
	// daily_status_history is daily health status for a service over multiple days
	// the keys are the number of days before today the DailyHealthStatus is from
	DailyStatusHistory map[int32]*DailyHealthStatus `protobuf:"bytes,2,rep,name=daily_status_history,json=dailyStatusHistory,proto3" json:"daily_status_history,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SystemServiceStatusHistory) Reset() {
	*x = SystemServiceStatusHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_system_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemServiceStatusHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemServiceStatusHistory) ProtoMessage() {}

func (x *SystemServiceStatusHistory) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_system_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemServiceStatusHistory.ProtoReflect.Descriptor instead.
func (*SystemServiceStatusHistory) Descriptor() ([]byte, []int) {
	return file_porter_v1_system_service_proto_rawDescGZIP(), []int{4}
}

func (x *SystemServiceStatusHistory) GetSystemService() *SystemService {
	if x != nil {
		return x.SystemService
	}
	return nil
}

func (x *SystemServiceStatusHistory) GetDailyStatusHistory() map[int32]*DailyHealthStatus {
	if x != nil {
		return x.DailyStatusHistory
	}
	return nil
}

// ClusterStatusHistory holds health status history of a certain type for a cluster
type ClusterStatusHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cluster_health_type is the type this history is generated for
	ClusterHealthType ClusterHealthType `protobuf:"varint,1,opt,name=cluster_health_type,json=clusterHealthType,proto3,enum=porter.v1.ClusterHealthType" json:"cluster_health_type,omitempty"`
	// daily_status_history is daily health status for the cluster over multiple days for the cluster_health_type
	// the keys are the number of days before today the DailyHealthStatus is from
	DailyStatusHistory map[int32]*DailyHealthStatus `protobuf:"bytes,2,rep,name=daily_status_history,json=dailyStatusHistory,proto3" json:"daily_status_history,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ClusterStatusHistory) Reset() {
	*x = ClusterStatusHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_system_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterStatusHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterStatusHistory) ProtoMessage() {}

func (x *ClusterStatusHistory) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_system_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterStatusHistory.ProtoReflect.Descriptor instead.
func (*ClusterStatusHistory) Descriptor() ([]byte, []int) {
	return file_porter_v1_system_service_proto_rawDescGZIP(), []int{5}
}

func (x *ClusterStatusHistory) GetClusterHealthType() ClusterHealthType {
	if x != nil {
		return x.ClusterHealthType
	}
	return ClusterHealthType_CLUSTER_HEALTH_TYPE_UNSPECIFIED
}

func (x *ClusterStatusHistory) GetDailyStatusHistory() map[int32]*DailyHealthStatus {
	if x != nil {
		return x.DailyStatusHistory
	}
	return nil
}

var File_porter_v1_system_service_proto protoreflect.FileDescriptor

var file_porter_v1_system_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65,
	0x75, 0x73, 0x5f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x92, 0x01, 0x0a, 0x0d, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x14, 0x69, 0x6e, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x5f,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x76, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x12, 0x69, 0x6e, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x22, 0xdf, 0x01, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x3a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x6e,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x5d, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x11, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4a, 0x0a, 0x12, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x52, 0x11, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0f, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0e, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x22, 0xb3, 0x02, 0x0a, 0x1a, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x3f, 0x0a, 0x0e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x0d, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a, 0x14, 0x64, 0x61, 0x69,
	0x6c, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x44,
	0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0x63, 0x0a, 0x17, 0x44, 0x61,
	0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0xb4, 0x02, 0x0a, 0x14, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x4c, 0x0a, 0x13, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x11, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x69, 0x0a, 0x14, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x64,
	0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x1a, 0x63, 0x0a, 0x17, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x64, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x59, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x49, 0x41, 0x4c, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x03, 0x2a, 0xa6, 0x01, 0x0a,
	0x11, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x23, 0x0a, 0x1f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x48, 0x45,
	0x41, 0x4c, 0x54, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x4c, 0x55, 0x53, 0x54,
	0x45, 0x52, 0x5f, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43,
	0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x4c,
	0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x50, 0x49, 0x4e, 0x47, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x27, 0x0a, 0x23,
	0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x5f, 0x48, 0x45, 0x41, 0x4c,
	0x54, 0x48, 0x59, 0x10, 0x03, 0x42, 0xad, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2d, 0x64, 0x65, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x15, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_porter_v1_system_service_proto_rawDescOnce sync.Once
	file_porter_v1_system_service_proto_rawDescData = file_porter_v1_system_service_proto_rawDesc
)

func file_porter_v1_system_service_proto_rawDescGZIP() []byte {
	file_porter_v1_system_service_proto_rawDescOnce.Do(func() {
		file_porter_v1_system_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_porter_v1_system_service_proto_rawDescData)
	})
	return file_porter_v1_system_service_proto_rawDescData
}

var file_porter_v1_system_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_porter_v1_system_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_porter_v1_system_service_proto_goTypes = []interface{}{
	(Status)(0),                        // 0: porter.v1.Status
	(ClusterHealthType)(0),             // 1: porter.v1.ClusterHealthType
	(*SystemService)(nil),              // 2: porter.v1.SystemService
	(*HealthStatus)(nil),               // 3: porter.v1.HealthStatus
	(*StatusPercentage)(nil),           // 4: porter.v1.StatusPercentage
	(*DailyHealthStatus)(nil),          // 5: porter.v1.DailyHealthStatus
	(*SystemServiceStatusHistory)(nil), // 6: porter.v1.SystemServiceStatusHistory
	(*ClusterStatusHistory)(nil),       // 7: porter.v1.ClusterStatusHistory
	nil,                                // 8: porter.v1.SystemServiceStatusHistory.DailyStatusHistoryEntry
	nil,                                // 9: porter.v1.ClusterStatusHistory.DailyStatusHistoryEntry
	(InvolvedObjectType)(0),            // 10: porter.v1.InvolvedObjectType
	(*timestamppb.Timestamp)(nil),      // 11: google.protobuf.Timestamp
}
var file_porter_v1_system_service_proto_depIdxs = []int32{
	10, // 0: porter.v1.SystemService.involved_object_type:type_name -> porter.v1.InvolvedObjectType
	11, // 1: porter.v1.HealthStatus.start_time:type_name -> google.protobuf.Timestamp
	11, // 2: porter.v1.HealthStatus.end_time:type_name -> google.protobuf.Timestamp
	0,  // 3: porter.v1.HealthStatus.status:type_name -> porter.v1.Status
	0,  // 4: porter.v1.StatusPercentage.status:type_name -> porter.v1.Status
	4,  // 5: porter.v1.DailyHealthStatus.status_percentages:type_name -> porter.v1.StatusPercentage
	3,  // 6: porter.v1.DailyHealthStatus.health_statuses:type_name -> porter.v1.HealthStatus
	2,  // 7: porter.v1.SystemServiceStatusHistory.system_service:type_name -> porter.v1.SystemService
	8,  // 8: porter.v1.SystemServiceStatusHistory.daily_status_history:type_name -> porter.v1.SystemServiceStatusHistory.DailyStatusHistoryEntry
	1,  // 9: porter.v1.ClusterStatusHistory.cluster_health_type:type_name -> porter.v1.ClusterHealthType
	9,  // 10: porter.v1.ClusterStatusHistory.daily_status_history:type_name -> porter.v1.ClusterStatusHistory.DailyStatusHistoryEntry
	5,  // 11: porter.v1.SystemServiceStatusHistory.DailyStatusHistoryEntry.value:type_name -> porter.v1.DailyHealthStatus
	5,  // 12: porter.v1.ClusterStatusHistory.DailyStatusHistoryEntry.value:type_name -> porter.v1.DailyHealthStatus
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_porter_v1_system_service_proto_init() }
func file_porter_v1_system_service_proto_init() {
	if File_porter_v1_system_service_proto != nil {
		return
	}
	file_porter_v1_prometheus_alerts_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_porter_v1_system_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemService); i {
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
		file_porter_v1_system_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthStatus); i {
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
		file_porter_v1_system_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusPercentage); i {
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
		file_porter_v1_system_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DailyHealthStatus); i {
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
		file_porter_v1_system_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemServiceStatusHistory); i {
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
		file_porter_v1_system_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterStatusHistory); i {
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
	file_porter_v1_system_service_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_porter_v1_system_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_porter_v1_system_service_proto_goTypes,
		DependencyIndexes: file_porter_v1_system_service_proto_depIdxs,
		EnumInfos:         file_porter_v1_system_service_proto_enumTypes,
		MessageInfos:      file_porter_v1_system_service_proto_msgTypes,
	}.Build()
	File_porter_v1_system_service_proto = out.File
	file_porter_v1_system_service_proto_rawDesc = nil
	file_porter_v1_system_service_proto_goTypes = nil
	file_porter_v1_system_service_proto_depIdxs = nil
}
