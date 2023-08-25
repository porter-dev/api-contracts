// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: porter/v1/service.proto

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

// ServiceType is used to categorize services, being one of web, worker, or job
type ServiceType int32

const (
	ServiceType_SERVICE_TYPE_UNSPECIFIED ServiceType = 0
	ServiceType_SERVICE_TYPE_WEB         ServiceType = 1
	ServiceType_SERVICE_TYPE_WORKER      ServiceType = 2
	ServiceType_SERVICE_TYPE_JOB         ServiceType = 3
)

// Enum value maps for ServiceType.
var (
	ServiceType_name = map[int32]string{
		0: "SERVICE_TYPE_UNSPECIFIED",
		1: "SERVICE_TYPE_WEB",
		2: "SERVICE_TYPE_WORKER",
		3: "SERVICE_TYPE_JOB",
	}
	ServiceType_value = map[string]int32{
		"SERVICE_TYPE_UNSPECIFIED": 0,
		"SERVICE_TYPE_WEB":         1,
		"SERVICE_TYPE_WORKER":      2,
		"SERVICE_TYPE_JOB":         3,
	}
)

func (x ServiceType) Enum() *ServiceType {
	p := new(ServiceType)
	*p = x
	return p
}

func (x ServiceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceType) Descriptor() protoreflect.EnumDescriptor {
	return file_porter_v1_service_proto_enumTypes[0].Descriptor()
}

func (ServiceType) Type() protoreflect.EnumType {
	return &file_porter_v1_service_proto_enumTypes[0]
}

func (x ServiceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceType.Descriptor instead.
func (ServiceType) EnumDescriptor() ([]byte, []int) {
	return file_porter_v1_service_proto_rawDescGZIP(), []int{0}
}

// Service is the top-level configuration for a service
type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// run is the command to start the service
	Run string `protobuf:"bytes,1,opt,name=run,proto3" json:"run,omitempty"`
	// instances is the number of instances, or replicas, to run
	Instances int32 `protobuf:"varint,2,opt,name=instances,proto3" json:"instances,omitempty"`
	// port is the port the service listens on
	Port int32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	// cpu_cores is the number of CPU cores to allocate to the service
	CpuCores float32 `protobuf:"fixed32,4,opt,name=cpu_cores,json=cpuCores,proto3" json:"cpu_cores,omitempty"`
	// ram_megabytes is the amount of memory to allocate to the service
	RamMegabytes int32 `protobuf:"varint,5,opt,name=ram_megabytes,json=ramMegabytes,proto3" json:"ram_megabytes,omitempty"`
	// config is the service-specific configuration
	//
	// Types that are assignable to Config:
	//
	//	*Service_WebConfig
	//	*Service_WorkerConfig
	//	*Service_JobConfig
	Config isService_Config `protobuf_oneof:"config"`
	// type is the type of service
	Type ServiceType `protobuf:"varint,9,opt,name=type,proto3,enum=porter.v1.ServiceType" json:"type,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_porter_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *Service) GetRun() string {
	if x != nil {
		return x.Run
	}
	return ""
}

func (x *Service) GetInstances() int32 {
	if x != nil {
		return x.Instances
	}
	return 0
}

func (x *Service) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Service) GetCpuCores() float32 {
	if x != nil {
		return x.CpuCores
	}
	return 0
}

func (x *Service) GetRamMegabytes() int32 {
	if x != nil {
		return x.RamMegabytes
	}
	return 0
}

func (m *Service) GetConfig() isService_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (x *Service) GetWebConfig() *WebServiceConfig {
	if x, ok := x.GetConfig().(*Service_WebConfig); ok {
		return x.WebConfig
	}
	return nil
}

func (x *Service) GetWorkerConfig() *WorkerServiceConfig {
	if x, ok := x.GetConfig().(*Service_WorkerConfig); ok {
		return x.WorkerConfig
	}
	return nil
}

func (x *Service) GetJobConfig() *JobServiceConfig {
	if x, ok := x.GetConfig().(*Service_JobConfig); ok {
		return x.JobConfig
	}
	return nil
}

func (x *Service) GetType() ServiceType {
	if x != nil {
		return x.Type
	}
	return ServiceType_SERVICE_TYPE_UNSPECIFIED
}

type isService_Config interface {
	isService_Config()
}

type Service_WebConfig struct {
	WebConfig *WebServiceConfig `protobuf:"bytes,6,opt,name=web_config,json=webConfig,proto3,oneof"`
}

type Service_WorkerConfig struct {
	WorkerConfig *WorkerServiceConfig `protobuf:"bytes,7,opt,name=worker_config,json=workerConfig,proto3,oneof"`
}

type Service_JobConfig struct {
	JobConfig *JobServiceConfig `protobuf:"bytes,8,opt,name=job_config,json=jobConfig,proto3,oneof"`
}

func (*Service_WebConfig) isService_Config() {}

func (*Service_WorkerConfig) isService_Config() {}

func (*Service_JobConfig) isService_Config() {}

// WebServiceConfig is the configuration for a web service
type WebServiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// autoscaling is the autoscaling configuration
	Autoscaling *Autoscaling `protobuf:"bytes,1,opt,name=autoscaling,proto3" json:"autoscaling,omitempty"`
	// domains is the list of custom domains for this service
	Domains []*Domain `protobuf:"bytes,2,rep,name=domains,proto3" json:"domains,omitempty"`
	// health_check is the health check configuration, used for liveness and readiness probes
	HealthCheck *HealthCheck `protobuf:"bytes,3,opt,name=health_check,json=healthCheck,proto3" json:"health_check,omitempty"`
	// private indicates whether or not the web service should be private (not publicly accessible)
	Private bool `protobuf:"varint,4,opt,name=private,proto3" json:"private,omitempty"`
}

func (x *WebServiceConfig) Reset() {
	*x = WebServiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebServiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebServiceConfig) ProtoMessage() {}

func (x *WebServiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebServiceConfig.ProtoReflect.Descriptor instead.
func (*WebServiceConfig) Descriptor() ([]byte, []int) {
	return file_porter_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *WebServiceConfig) GetAutoscaling() *Autoscaling {
	if x != nil {
		return x.Autoscaling
	}
	return nil
}

func (x *WebServiceConfig) GetDomains() []*Domain {
	if x != nil {
		return x.Domains
	}
	return nil
}

func (x *WebServiceConfig) GetHealthCheck() *HealthCheck {
	if x != nil {
		return x.HealthCheck
	}
	return nil
}

func (x *WebServiceConfig) GetPrivate() bool {
	if x != nil {
		return x.Private
	}
	return false
}

// WorkerServiceConfig is the configuration for a worker service
type WorkerServiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Autoscaling *Autoscaling `protobuf:"bytes,1,opt,name=autoscaling,proto3" json:"autoscaling,omitempty"`
}

func (x *WorkerServiceConfig) Reset() {
	*x = WorkerServiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerServiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerServiceConfig) ProtoMessage() {}

func (x *WorkerServiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerServiceConfig.ProtoReflect.Descriptor instead.
func (*WorkerServiceConfig) Descriptor() ([]byte, []int) {
	return file_porter_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *WorkerServiceConfig) GetAutoscaling() *Autoscaling {
	if x != nil {
		return x.Autoscaling
	}
	return nil
}

// JobServiceConfig is the configuration for a job service
type JobServiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// allow_concurrent indicates whether or not runs of the job can be processed concurrently
	AllowConcurrent bool `protobuf:"varint,1,opt,name=allow_concurrent,json=allowConcurrent,proto3" json:"allow_concurrent,omitempty"`
	// cron is the cron expression for the job
	Cron string `protobuf:"bytes,2,opt,name=cron,proto3" json:"cron,omitempty"`
}

func (x *JobServiceConfig) Reset() {
	*x = JobServiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobServiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobServiceConfig) ProtoMessage() {}

func (x *JobServiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobServiceConfig.ProtoReflect.Descriptor instead.
func (*JobServiceConfig) Descriptor() ([]byte, []int) {
	return file_porter_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *JobServiceConfig) GetAllowConcurrent() bool {
	if x != nil {
		return x.AllowConcurrent
	}
	return false
}

func (x *JobServiceConfig) GetCron() string {
	if x != nil {
		return x.Cron
	}
	return ""
}

// Domain is the configuration for a custom domain for a web service
type Domain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Domain) Reset() {
	*x = Domain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Domain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Domain) ProtoMessage() {}

func (x *Domain) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Domain.ProtoReflect.Descriptor instead.
func (*Domain) Descriptor() ([]byte, []int) {
	return file_porter_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *Domain) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Autoscaling is the autoscaling configuration
type Autoscaling struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// enabled explicitly enables or disables autoscaling
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// min_instances is the minimum number of instances to run
	MinInstances int32 `protobuf:"varint,2,opt,name=min_instances,json=minInstances,proto3" json:"min_instances,omitempty"`
	// max_instances is the maximum number of instances to run
	MaxInstances int32 `protobuf:"varint,3,opt,name=max_instances,json=maxInstances,proto3" json:"max_instances,omitempty"`
	// cpu_threshold_percent is the CPU threshold percentage to trigger scaling
	CpuThresholdPercent int32 `protobuf:"varint,4,opt,name=cpu_threshold_percent,json=cpuThresholdPercent,proto3" json:"cpu_threshold_percent,omitempty"`
	// memory_threshold_percent is the memory threshold percentage to trigger scaling
	MemoryThresholdPercent int32 `protobuf:"varint,5,opt,name=memory_threshold_percent,json=memoryThresholdPercent,proto3" json:"memory_threshold_percent,omitempty"`
}

func (x *Autoscaling) Reset() {
	*x = Autoscaling{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Autoscaling) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Autoscaling) ProtoMessage() {}

func (x *Autoscaling) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Autoscaling.ProtoReflect.Descriptor instead.
func (*Autoscaling) Descriptor() ([]byte, []int) {
	return file_porter_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *Autoscaling) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Autoscaling) GetMinInstances() int32 {
	if x != nil {
		return x.MinInstances
	}
	return 0
}

func (x *Autoscaling) GetMaxInstances() int32 {
	if x != nil {
		return x.MaxInstances
	}
	return 0
}

func (x *Autoscaling) GetCpuThresholdPercent() int32 {
	if x != nil {
		return x.CpuThresholdPercent
	}
	return 0
}

func (x *Autoscaling) GetMemoryThresholdPercent() int32 {
	if x != nil {
		return x.MemoryThresholdPercent
	}
	return 0
}

// HealthCheck is the health check configuration
type HealthCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// enabled explicitly enables or disables health checks
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// http_path is the HTTP path to use for the health check
	HttpPath string `protobuf:"bytes,2,opt,name=http_path,json=httpPath,proto3" json:"http_path,omitempty"`
}

func (x *HealthCheck) Reset() {
	*x = HealthCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_porter_v1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheck) ProtoMessage() {}

func (x *HealthCheck) ProtoReflect() protoreflect.Message {
	mi := &file_porter_v1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheck.ProtoReflect.Descriptor instead.
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return file_porter_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *HealthCheck) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *HealthCheck) GetHttpPath() string {
	if x != nil {
		return x.HttpPath
	}
	return ""
}

var File_porter_v1_service_proto protoreflect.FileDescriptor

var file_porter_v1_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x22, 0x88, 0x03, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x72, 0x75, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72,
	0x75, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x70, 0x75, 0x5f, 0x63, 0x6f, 0x72, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x63, 0x70, 0x75, 0x43, 0x6f, 0x72, 0x65,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x61, 0x6d, 0x5f, 0x6d, 0x65, 0x67, 0x61, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x61, 0x6d, 0x4d, 0x65, 0x67,
	0x61, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x0a, 0x77, 0x65, 0x62, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x09, 0x77, 0x65, 0x62, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x45, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0c, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3c, 0x0a, 0x0a, 0x6a,
	0x6f, 0x62, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x09,
	0x6a, 0x6f, 0x62, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0xce, 0x01, 0x0a, 0x10, 0x57, 0x65, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x38, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e,
	0x67, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x2b,
	0x0a, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x52, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x39, 0x0a, 0x0c, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x0b, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x22, 0x4f, 0x0a, 0x13, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x38, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x73,
	0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61,
	0x6c, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e,
	0x67, 0x22, 0x51, 0x0a, 0x10, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63,
	0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x72, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x72, 0x6f, 0x6e, 0x22, 0x1c, 0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0xdf, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69,
	0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x70, 0x75, 0x5f, 0x74, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x63, 0x70, 0x75, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x18, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x70,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x50, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x22, 0x44, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x68, 0x74, 0x74, 0x70, 0x50, 0x61, 0x74, 0x68, 0x2a, 0x70, 0x0a, 0x0b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x45, 0x52,
	0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x45, 0x52, 0x56, 0x49,
	0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x45, 0x42, 0x10, 0x01, 0x12, 0x17, 0x0a,
	0x13, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x4f,
	0x52, 0x4b, 0x45, 0x52, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4a, 0x4f, 0x42, 0x10, 0x03, 0x42, 0xa7, 0x01, 0x0a,
	0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43,
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
	file_porter_v1_service_proto_rawDescOnce sync.Once
	file_porter_v1_service_proto_rawDescData = file_porter_v1_service_proto_rawDesc
)

func file_porter_v1_service_proto_rawDescGZIP() []byte {
	file_porter_v1_service_proto_rawDescOnce.Do(func() {
		file_porter_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_porter_v1_service_proto_rawDescData)
	})
	return file_porter_v1_service_proto_rawDescData
}

var file_porter_v1_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_porter_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_porter_v1_service_proto_goTypes = []interface{}{
	(ServiceType)(0),            // 0: porter.v1.ServiceType
	(*Service)(nil),             // 1: porter.v1.Service
	(*WebServiceConfig)(nil),    // 2: porter.v1.WebServiceConfig
	(*WorkerServiceConfig)(nil), // 3: porter.v1.WorkerServiceConfig
	(*JobServiceConfig)(nil),    // 4: porter.v1.JobServiceConfig
	(*Domain)(nil),              // 5: porter.v1.Domain
	(*Autoscaling)(nil),         // 6: porter.v1.Autoscaling
	(*HealthCheck)(nil),         // 7: porter.v1.HealthCheck
}
var file_porter_v1_service_proto_depIdxs = []int32{
	2, // 0: porter.v1.Service.web_config:type_name -> porter.v1.WebServiceConfig
	3, // 1: porter.v1.Service.worker_config:type_name -> porter.v1.WorkerServiceConfig
	4, // 2: porter.v1.Service.job_config:type_name -> porter.v1.JobServiceConfig
	0, // 3: porter.v1.Service.type:type_name -> porter.v1.ServiceType
	6, // 4: porter.v1.WebServiceConfig.autoscaling:type_name -> porter.v1.Autoscaling
	5, // 5: porter.v1.WebServiceConfig.domains:type_name -> porter.v1.Domain
	7, // 6: porter.v1.WebServiceConfig.health_check:type_name -> porter.v1.HealthCheck
	6, // 7: porter.v1.WorkerServiceConfig.autoscaling:type_name -> porter.v1.Autoscaling
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_porter_v1_service_proto_init() }
func file_porter_v1_service_proto_init() {
	if File_porter_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_porter_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_porter_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebServiceConfig); i {
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
		file_porter_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerServiceConfig); i {
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
		file_porter_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobServiceConfig); i {
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
		file_porter_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Domain); i {
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
		file_porter_v1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Autoscaling); i {
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
		file_porter_v1_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheck); i {
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
	file_porter_v1_service_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Service_WebConfig)(nil),
		(*Service_WorkerConfig)(nil),
		(*Service_JobConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_porter_v1_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_porter_v1_service_proto_goTypes,
		DependencyIndexes: file_porter_v1_service_proto_depIdxs,
		EnumInfos:         file_porter_v1_service_proto_enumTypes,
		MessageInfos:      file_porter_v1_service_proto_msgTypes,
	}.Build()
	File_porter_v1_service_proto = out.File
	file_porter_v1_service_proto_rawDesc = nil
	file_porter_v1_service_proto_goTypes = nil
	file_porter_v1_service_proto_depIdxs = nil
}
