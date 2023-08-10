// @generated by protoc-gen-es v1.3.0
// @generated from file porter/v1/service.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * Service type is used to categorize services, being one of web, worker, or job
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
   * Run is the command to start the service
   *
   * @generated from field: string run = 1;
   */
  run: string;

  /**
   * Instances is the number of instances, or replicas, to run
   *
   * @generated from field: int32 instances = 2;
   */
  instances: number;

  /**
   * Port is the port the service listens on
   *
   * @generated from field: int32 port = 3;
   */
  port: number;

  /**
   * CPU cores is the number of CPU cores to allocate to the service
   *
   * @generated from field: float cpu_cores = 4;
   */
  cpuCores: number;

  /**
   * RAM megabytes is the amount of memory to allocate to the service
   *
   * @generated from field: int32 ram_megabytes = 5;
   */
  ramMegabytes: number;

  /**
   * Config is the service-specific configuration
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
   * Type is the type of service
   *
   * @generated from field: porter.v1.ServiceType type = 9;
   */
  type: ServiceType;

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
   * Autoscaling is the autoscaling configuration
   *
   * @generated from field: porter.v1.Autoscaling autoscaling = 1;
   */
  autoscaling?: Autoscaling;

  /**
   * Domains is the list of custom domains for this service
   *
   * @generated from field: repeated porter.v1.Domain domains = 2;
   */
  domains: Domain[];

  /**
   * HealthCheck is the health check configuration, used for liveness and readiness probes
   *
   * @generated from field: porter.v1.HealthCheck health_check = 3;
   */
  healthCheck?: HealthCheck;

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
   * AllowConcurrent indicates whether or not runs of the job can be processed concurrently
   *
   * @generated from field: bool allow_concurrent = 1;
   */
  allowConcurrent: boolean;

  /**
   * Cron is the cron expression for the job
   *
   * @generated from field: string cron = 2;
   */
  cron: string;

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
   * MinInstances is the minimum number of instances to run
   *
   * @generated from field: int32 min_instances = 1;
   */
  minInstances: number;

  /**
   * MaxInstances is the maximum number of instances to run
   *
   * @generated from field: int32 max_instances = 2;
   */
  maxInstances: number;

  /**
   * CPUThresholdPercent is the CPU threshold percentage to trigger scaling
   *
   * @generated from field: int32 cpu_threshold_percent = 3;
   */
  cpuThresholdPercent: number;

  /**
   * MemoryThresholdPercent is the memory threshold percentage to trigger scaling
   *
   * @generated from field: int32 memory_threshold_percent = 4;
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
   * HTTPPath is the HTTP path to use for the health check
   *
   * @generated from field: string http_path = 1;
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

