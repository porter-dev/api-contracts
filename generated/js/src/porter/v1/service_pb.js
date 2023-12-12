// @generated by protoc-gen-es v1.4.2
// @generated from file porter/v1/service.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * ServiceType is used to categorize services, being one of web, worker, or job
 *
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
 * Service is the top-level configuration for a service
 *
 * @generated from message porter.v1.Service
 */
export const Service = proto3.makeMessageType(
  "porter.v1.Service",
  () => [
    { no: 1, name: "run", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "instances", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "port", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "cpu_cores", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 5, name: "ram_megabytes", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 6, name: "web_config", kind: "message", T: WebServiceConfig, oneof: "config" },
    { no: 7, name: "worker_config", kind: "message", T: WorkerServiceConfig, oneof: "config" },
    { no: 8, name: "job_config", kind: "message", T: JobServiceConfig, oneof: "config" },
    { no: 9, name: "type", kind: "enum", T: proto3.getEnumType(ServiceType) },
    { no: 10, name: "absolute_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 11, name: "smart_optimization", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 12, name: "run_optional", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 13, name: "gpu_cores_nvidia", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 14, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 15, name: "instances_optional", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 16, name: "gpu_cores", kind: "scalar", T: 2 /* ScalarType.FLOAT */, opt: true },
    { no: 17, name: "gpu", kind: "message", T: GPU },
    { no: 18, name: "termination_grace_period_seconds", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
  ],
);

/**
 * WebServiceConfig is the configuration for a web service
 *
 * @generated from message porter.v1.WebServiceConfig
 */
export const WebServiceConfig = proto3.makeMessageType(
  "porter.v1.WebServiceConfig",
  () => [
    { no: 1, name: "autoscaling", kind: "message", T: Autoscaling },
    { no: 2, name: "domains", kind: "message", T: Domain, repeated: true },
    { no: 3, name: "health_check", kind: "message", T: HealthCheck },
    { no: 4, name: "private", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 5, name: "ingress_annotations", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 6, name: "disable_tls", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ],
);

/**
 * WorkerServiceConfig is the configuration for a worker service
 *
 * @generated from message porter.v1.WorkerServiceConfig
 */
export const WorkerServiceConfig = proto3.makeMessageType(
  "porter.v1.WorkerServiceConfig",
  () => [
    { no: 1, name: "autoscaling", kind: "message", T: Autoscaling },
  ],
);

/**
 * JobServiceConfig is the configuration for a job service
 *
 * @generated from message porter.v1.JobServiceConfig
 */
export const JobServiceConfig = proto3.makeMessageType(
  "porter.v1.JobServiceConfig",
  () => [
    { no: 1, name: "allow_concurrent", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "cron", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "suspend_cron", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 4, name: "allow_concurrent_optional", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 5, name: "timeout_seconds", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * Domain is the configuration for a custom domain for a web service
 *
 * @generated from message porter.v1.Domain
 */
export const Domain = proto3.makeMessageType(
  "porter.v1.Domain",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * Autoscaling is the autoscaling configuration
 *
 * @generated from message porter.v1.Autoscaling
 */
export const Autoscaling = proto3.makeMessageType(
  "porter.v1.Autoscaling",
  () => [
    { no: 1, name: "enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "min_instances", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "max_instances", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "cpu_threshold_percent", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "memory_threshold_percent", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * HealthCheck is the health check configuration
 *
 * @generated from message porter.v1.HealthCheck
 */
export const HealthCheck = proto3.makeMessageType(
  "porter.v1.HealthCheck",
  () => [
    { no: 1, name: "enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "http_path", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * GPU is the gpu configuration
 *
 * @generated from message porter.v1.GPU
 */
export const GPU = proto3.makeMessageType(
  "porter.v1.GPU",
  () => [
    { no: 1, name: "enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "gpu_cores_nvidia", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

