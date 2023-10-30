// @generated by protoc-gen-es v1.4.1
// @generated from file porter/v1/service.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * ServiceType is used to categorize services, being one of web, worker, or job
 *
 * @generated from enum porter.v1.ServiceType
 */
export declare enum ServiceType {
  /**
   * @generated from enum value: SERVICE_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: SERVICE_TYPE_WEB = 1;
   */
  WEB = 1,

  /**
   * @generated from enum value: SERVICE_TYPE_WORKER = 2;
   */
  WORKER = 2,

  /**
   * @generated from enum value: SERVICE_TYPE_JOB = 3;
   */
  JOB = 3,
}

/**
 * Service is the top-level configuration for a service
 *
 * @generated from message porter.v1.Service
 */
export declare class Service extends Message<Service> {
  /**
   * run is the command to start the service. Deprecated: use run_optional instead, as this allows us to tell if the field was explicitly set or not.
   *
   * @generated from field: string run = 1 [deprecated = true];
   * @deprecated
   */
  run: string;

  /**
   * instances is the number of instances, or replicas, to run. Deprecated: use instances_optional instead.
   *
   * @generated from field: int32 instances = 2 [deprecated = true];
   * @deprecated
   */
  instances: number;

  /**
   * port is the port the service listens on
   *
   * @generated from field: int32 port = 3;
   */
  port: number;

  /**
   * cpu_cores is the number of CPU cores to allocate to the service
   *
   * @generated from field: float cpu_cores = 4;
   */
  cpuCores: number;

  /**
   * ram_megabytes is the amount of memory to allocate to the service
   *
   * @generated from field: int32 ram_megabytes = 5;
   */
  ramMegabytes: number;

  /**
   * config is the service-specific configuration
   *
   * @generated from oneof porter.v1.Service.config
   */
  config: {
    /**
     * @generated from field: porter.v1.WebServiceConfig web_config = 6;
     */
    value: WebServiceConfig;
    case: "webConfig";
  } | {
    /**
     * @generated from field: porter.v1.WorkerServiceConfig worker_config = 7;
     */
    value: WorkerServiceConfig;
    case: "workerConfig";
  } | {
    /**
     * @generated from field: porter.v1.JobServiceConfig job_config = 8;
     */
    value: JobServiceConfig;
    case: "jobConfig";
  } | { case: undefined; value?: undefined };

  /**
   * type is the type of service
   *
   * @generated from field: porter.v1.ServiceType type = 9;
   */
  type: ServiceType;

  /**
   * absolute_name is the name that should be used verbatim on the cluster. If not specified, a name is generated from the app, service and service type.
   *
   * @generated from field: string absolute_name = 10;
   */
  absoluteName: string;

  /**
   * smart_optimization is to toggle the smart optimization feature
   *
   * @generated from field: optional bool smart_optimization = 11;
   */
  smartOptimization?: boolean;

  /**
   * run_optional is the command to start the service
   *
   * @generated from field: optional string run_optional = 12;
   */
  runOptional?: string;

  /**
   * gpu_cores_nvidia is the number of GPU cores to allocate to the service. This is only used for nvidia-backed GPU's at present. If set to a non-zero value, this service will only be scheduled if a GPU node with the specified number of cores is available.
   *
   * @generated from field: float gpu_cores_nvidia = 13;
   */
  gpuCoresNvidia: number;

  /**
   * name is the user-defined name for the service
   *
   * @generated from field: string name = 14;
   */
  name: string;

  /**
   * instances_optional is the number of instances, or replicas, to run
   *
   * @generated from field: optional int32 instances_optional = 15;
   */
  instancesOptional?: number;

  constructor(data?: PartialMessage<Service>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Service";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Service;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Service;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Service;

  static equals(a: Service | PlainMessage<Service> | undefined, b: Service | PlainMessage<Service> | undefined): boolean;
}

/**
 * WebServiceConfig is the configuration for a web service
 *
 * @generated from message porter.v1.WebServiceConfig
 */
export declare class WebServiceConfig extends Message<WebServiceConfig> {
  /**
   * autoscaling is the autoscaling configuration
   *
   * @generated from field: porter.v1.Autoscaling autoscaling = 1;
   */
  autoscaling?: Autoscaling;

  /**
   * domains is the list of custom domains for this service
   *
   * @generated from field: repeated porter.v1.Domain domains = 2;
   */
  domains: Domain[];

  /**
   * health_check is the health check configuration, used for liveness and readiness probes
   *
   * @generated from field: porter.v1.HealthCheck health_check = 3;
   */
  healthCheck?: HealthCheck;

  /**
   * private indicates whether or not the web service should be private (not publicly accessible)
   *
   * @generated from field: optional bool private = 4;
   */
  private?: boolean;

  /**
   * ingress_annotations is a map of annotations to apply to the ingress resource
   *
   * @generated from field: map<string, string> ingress_annotations = 5;
   */
  ingressAnnotations: { [key: string]: string };

  constructor(data?: PartialMessage<WebServiceConfig>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.WebServiceConfig";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): WebServiceConfig;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): WebServiceConfig;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): WebServiceConfig;

  static equals(a: WebServiceConfig | PlainMessage<WebServiceConfig> | undefined, b: WebServiceConfig | PlainMessage<WebServiceConfig> | undefined): boolean;
}

/**
 * WorkerServiceConfig is the configuration for a worker service
 *
 * @generated from message porter.v1.WorkerServiceConfig
 */
export declare class WorkerServiceConfig extends Message<WorkerServiceConfig> {
  /**
   * @generated from field: porter.v1.Autoscaling autoscaling = 1;
   */
  autoscaling?: Autoscaling;

  constructor(data?: PartialMessage<WorkerServiceConfig>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.WorkerServiceConfig";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): WorkerServiceConfig;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): WorkerServiceConfig;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): WorkerServiceConfig;

  static equals(a: WorkerServiceConfig | PlainMessage<WorkerServiceConfig> | undefined, b: WorkerServiceConfig | PlainMessage<WorkerServiceConfig> | undefined): boolean;
}

/**
 * JobServiceConfig is the configuration for a job service
 *
 * @generated from message porter.v1.JobServiceConfig
 */
export declare class JobServiceConfig extends Message<JobServiceConfig> {
  /**
   * allow_concurrent indicates whether or not runs of the job can be processed concurrently.  Deprecated: use allow_concurrent_optional instead, as this allows us to tell if the field was set in porter.yaml or not.
   *
   * @generated from field: bool allow_concurrent = 1 [deprecated = true];
   * @deprecated
   */
  allowConcurrent: boolean;

  /**
   * cron is the cron expression for the job
   *
   * @generated from field: string cron = 2;
   */
  cron: string;

  /**
   * suspend_cron indicates whether or not the cron should be suspended
   *
   * @generated from field: optional bool suspend_cron = 3;
   */
  suspendCron?: boolean;

  /**
   * allow_concurrent_optional indicates whether or not runs of the job can be processed concurrently
   *
   * @generated from field: optional bool allow_concurrent_optional = 4;
   */
  allowConcurrentOptional?: boolean;

  /**
   * timeout_seconds is the number of seconds to wait before timing out the job
   *
   * @generated from field: int64 timeout_seconds = 5;
   */
  timeoutSeconds: bigint;

  constructor(data?: PartialMessage<JobServiceConfig>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.JobServiceConfig";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): JobServiceConfig;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): JobServiceConfig;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): JobServiceConfig;

  static equals(a: JobServiceConfig | PlainMessage<JobServiceConfig> | undefined, b: JobServiceConfig | PlainMessage<JobServiceConfig> | undefined): boolean;
}

/**
 * Domain is the configuration for a custom domain for a web service
 *
 * @generated from message porter.v1.Domain
 */
export declare class Domain extends Message<Domain> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  constructor(data?: PartialMessage<Domain>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Domain";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Domain;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Domain;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Domain;

  static equals(a: Domain | PlainMessage<Domain> | undefined, b: Domain | PlainMessage<Domain> | undefined): boolean;
}

/**
 * Autoscaling is the autoscaling configuration
 *
 * @generated from message porter.v1.Autoscaling
 */
export declare class Autoscaling extends Message<Autoscaling> {
  /**
   * enabled explicitly enables or disables autoscaling
   *
   * @generated from field: bool enabled = 1;
   */
  enabled: boolean;

  /**
   * min_instances is the minimum number of instances to run
   *
   * @generated from field: int32 min_instances = 2;
   */
  minInstances: number;

  /**
   * max_instances is the maximum number of instances to run
   *
   * @generated from field: int32 max_instances = 3;
   */
  maxInstances: number;

  /**
   * cpu_threshold_percent is the CPU threshold percentage to trigger scaling
   *
   * @generated from field: int32 cpu_threshold_percent = 4;
   */
  cpuThresholdPercent: number;

  /**
   * memory_threshold_percent is the memory threshold percentage to trigger scaling
   *
   * @generated from field: int32 memory_threshold_percent = 5;
   */
  memoryThresholdPercent: number;

  constructor(data?: PartialMessage<Autoscaling>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Autoscaling";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Autoscaling;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Autoscaling;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Autoscaling;

  static equals(a: Autoscaling | PlainMessage<Autoscaling> | undefined, b: Autoscaling | PlainMessage<Autoscaling> | undefined): boolean;
}

/**
 * HealthCheck is the health check configuration
 *
 * @generated from message porter.v1.HealthCheck
 */
export declare class HealthCheck extends Message<HealthCheck> {
  /**
   * enabled explicitly enables or disables health checks
   *
   * @generated from field: bool enabled = 1;
   */
  enabled: boolean;

  /**
   * http_path is the HTTP path to use for the health check
   *
   * @generated from field: string http_path = 2;
   */
  httpPath: string;

  constructor(data?: PartialMessage<HealthCheck>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.HealthCheck";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): HealthCheck;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): HealthCheck;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): HealthCheck;

  static equals(a: HealthCheck | PlainMessage<HealthCheck> | undefined, b: HealthCheck | PlainMessage<HealthCheck> | undefined): boolean;
}

