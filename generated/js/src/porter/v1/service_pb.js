// @generated by protoc-gen-es v1.3.0
// @generated from file porter/v1/service.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum porter.v1.ServiceType
 */
export const ServiceType = proto3.makeEnum(
  "porter.v1.ServiceType",
  [
    {no: 0, name: "SERVICE_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "SERVICE_TYPE_WEB", localName: "WEB"},
    {no: 2, name: "SERVICE_TYPE_WORKER", localName: "WORKER"},
    {no: 3, name: "SERVICE_TYPE_JOB", localName: "JOB"},
  ],
);

/**
 * @generated from message porter.v1.Service
 */
export const Service = proto3.makeMessageType(
  "porter.v1.Service",
  () => [
    { no: 1, name: "run", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "web_config", kind: "message", T: WebServiceConfig, oneof: "config" },
    { no: 3, name: "worker_config", kind: "message", T: WorkerServiceConfig, oneof: "config" },
    { no: 4, name: "job_config", kind: "message", T: JobServiceConfig, oneof: "config" },
    { no: 5, name: "type", kind: "enum", T: proto3.getEnumType(ServiceType) },
  ],
);

/**
 * @generated from message porter.v1.WebServiceConfig
 */
export const WebServiceConfig = proto3.makeMessageType(
  "porter.v1.WebServiceConfig",
  () => [
    { no: 1, name: "replica_count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "resources", kind: "message", T: Resources },
    { no: 3, name: "container", kind: "message", T: Container },
    { no: 4, name: "autoscaling", kind: "message", T: Autoscaling },
    { no: 5, name: "ingress", kind: "message", T: Ingress },
    { no: 6, name: "health", kind: "message", T: Health },
    { no: 7, name: "cloud_sql", kind: "message", T: CloudSql },
  ],
);

/**
 * @generated from message porter.v1.WorkerServiceConfig
 */
export const WorkerServiceConfig = proto3.makeMessageType(
  "porter.v1.WorkerServiceConfig",
  () => [
    { no: 1, name: "replica_count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "resources", kind: "message", T: Resources },
    { no: 3, name: "container", kind: "message", T: Container },
    { no: 4, name: "autoscaling", kind: "message", T: Autoscaling },
    { no: 5, name: "cloud_sql", kind: "message", T: CloudSql },
  ],
);

/**
 * @generated from message porter.v1.JobServiceConfig
 */
export const JobServiceConfig = proto3.makeMessageType(
  "porter.v1.JobServiceConfig",
  () => [
    { no: 1, name: "allow_concurrent", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "resources", kind: "message", T: Resources },
    { no: 3, name: "container", kind: "message", T: Container },
    { no: 4, name: "schedule", kind: "message", T: Schedule },
    { no: 5, name: "paused", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 6, name: "cloud_sql", kind: "message", T: CloudSql },
  ],
);

/**
 * @generated from message porter.v1.RequestResources
 */
export const RequestResources = proto3.makeMessageType(
  "porter.v1.RequestResources",
  () => [
    { no: 1, name: "cpu", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "memory", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.Resources
 */
export const Resources = proto3.makeMessageType(
  "porter.v1.Resources",
  () => [
    { no: 1, name: "requests", kind: "message", T: RequestResources },
  ],
);

/**
 * @generated from message porter.v1.Schedule
 */
export const Schedule = proto3.makeMessageType(
  "porter.v1.Schedule",
  () => [
    { no: 1, name: "enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "value", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.Container
 */
export const Container = proto3.makeMessageType(
  "porter.v1.Container",
  () => [
    { no: 1, name: "command", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "port", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.Ingress
 */
export const Ingress = proto3.makeMessageType(
  "porter.v1.Ingress",
  () => [
    { no: 1, name: "enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "custom_domain", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "hosts", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "porter_hosts", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 5, name: "annotations", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
  ],
);

/**
 * @generated from message porter.v1.Autoscaling
 */
export const Autoscaling = proto3.makeMessageType(
  "porter.v1.Autoscaling",
  () => [
    { no: 1, name: "enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "min_replicas", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "max_replicas", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "target_cpu_utilization_threshold", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "target_memory_utilization_threshold", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message porter.v1.CloudSql
 */
export const CloudSql = proto3.makeMessageType(
  "porter.v1.CloudSql",
  () => [
    { no: 1, name: "enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "connection_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "db_port", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "service_account_json", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.LiveCheck
 */
export const LiveCheck = proto3.makeMessageType(
  "porter.v1.LiveCheck",
  () => [
    { no: 1, name: "enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "failure_threshold", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "path", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "period_seconds", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message porter.v1.ReadyCheck
 */
export const ReadyCheck = proto3.makeMessageType(
  "porter.v1.ReadyCheck",
  () => [
    { no: 1, name: "enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "failure_threshold", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "path", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "initial_delay_seconds", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message porter.v1.Health
 */
export const Health = proto3.makeMessageType(
  "porter.v1.Health",
  () => [
    { no: 1, name: "startup_probe", kind: "message", T: LiveCheck },
    { no: 2, name: "readiness_probe", kind: "message", T: ReadyCheck },
    { no: 3, name: "liveness_probe", kind: "message", T: LiveCheck },
  ],
);

