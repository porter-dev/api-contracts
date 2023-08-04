// @generated by protoc-gen-es v1.3.0
// @generated from file porter/v1/service.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
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
 * @generated from message porter.v1.Service
 */
export declare class Service extends Message<Service> {
  /**
   * @generated from field: string run = 1;
   */
  run: string;

  /**
   * @generated from oneof porter.v1.Service.config
   */
  config: {
    /**
     * @generated from field: porter.v1.WebServiceConfig web_config = 2;
     */
    value: WebServiceConfig;
    case: "webConfig";
  } | {
    /**
     * @generated from field: porter.v1.WorkerServiceConfig worker_config = 3;
     */
    value: WorkerServiceConfig;
    case: "workerConfig";
  } | {
    /**
     * @generated from field: porter.v1.JobServiceConfig job_config = 4;
     */
    value: JobServiceConfig;
    case: "jobConfig";
  } | { case: undefined; value?: undefined };

  /**
   * @generated from field: porter.v1.ServiceType type = 5;
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
 * @generated from message porter.v1.WebServiceConfig
 */
export declare class WebServiceConfig extends Message<WebServiceConfig> {
  /**
   * @generated from field: int32 replica_count = 1;
   */
  replicaCount: number;

  /**
   * @generated from field: porter.v1.Resources resources = 2;
   */
  resources?: Resources;

  /**
   * @generated from field: porter.v1.Container container = 3;
   */
  container?: Container;

  /**
   * @generated from field: porter.v1.Autoscaling autoscaling = 4;
   */
  autoscaling?: Autoscaling;

  /**
   * @generated from field: porter.v1.Ingress ingress = 5;
   */
  ingress?: Ingress;

  /**
   * @generated from field: porter.v1.Health health = 6;
   */
  health?: Health;

  /**
   * @generated from field: porter.v1.CloudSql cloud_sql = 7;
   */
  cloudSql?: CloudSql;

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
 * @generated from message porter.v1.WorkerServiceConfig
 */
export declare class WorkerServiceConfig extends Message<WorkerServiceConfig> {
  /**
   * @generated from field: int32 replica_count = 1;
   */
  replicaCount: number;

  /**
   * @generated from field: porter.v1.Resources resources = 2;
   */
  resources?: Resources;

  /**
   * @generated from field: porter.v1.Container container = 3;
   */
  container?: Container;

  /**
   * @generated from field: porter.v1.Autoscaling autoscaling = 4;
   */
  autoscaling?: Autoscaling;

  /**
   * @generated from field: porter.v1.CloudSql cloud_sql = 5;
   */
  cloudSql?: CloudSql;

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
 * @generated from message porter.v1.JobServiceConfig
 */
export declare class JobServiceConfig extends Message<JobServiceConfig> {
  /**
   * @generated from field: bool allow_concurrent = 1;
   */
  allowConcurrent: boolean;

  /**
   * @generated from field: porter.v1.Resources resources = 2;
   */
  resources?: Resources;

  /**
   * @generated from field: porter.v1.Container container = 3;
   */
  container?: Container;

  /**
   * @generated from field: porter.v1.Schedule schedule = 4;
   */
  schedule?: Schedule;

  /**
   * @generated from field: bool paused = 5;
   */
  paused: boolean;

  /**
   * @generated from field: porter.v1.CloudSql cloud_sql = 6;
   */
  cloudSql?: CloudSql;

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
 * @generated from message porter.v1.RequestResources
 */
export declare class RequestResources extends Message<RequestResources> {
  /**
   * @generated from field: string cpu = 1;
   */
  cpu: string;

  /**
   * @generated from field: string memory = 2;
   */
  memory: string;

  constructor(data?: PartialMessage<RequestResources>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.RequestResources";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RequestResources;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RequestResources;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RequestResources;

  static equals(a: RequestResources | PlainMessage<RequestResources> | undefined, b: RequestResources | PlainMessage<RequestResources> | undefined): boolean;
}

/**
 * @generated from message porter.v1.Resources
 */
export declare class Resources extends Message<Resources> {
  /**
   * @generated from field: porter.v1.RequestResources requests = 1;
   */
  requests?: RequestResources;

  constructor(data?: PartialMessage<Resources>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Resources";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Resources;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Resources;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Resources;

  static equals(a: Resources | PlainMessage<Resources> | undefined, b: Resources | PlainMessage<Resources> | undefined): boolean;
}

/**
 * @generated from message porter.v1.Schedule
 */
export declare class Schedule extends Message<Schedule> {
  /**
   * @generated from field: bool enabled = 1;
   */
  enabled: boolean;

  /**
   * @generated from field: string value = 2;
   */
  value: string;

  constructor(data?: PartialMessage<Schedule>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Schedule";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Schedule;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Schedule;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Schedule;

  static equals(a: Schedule | PlainMessage<Schedule> | undefined, b: Schedule | PlainMessage<Schedule> | undefined): boolean;
}

/**
 * @generated from message porter.v1.Container
 */
export declare class Container extends Message<Container> {
  /**
   * @generated from field: string command = 1;
   */
  command: string;

  /**
   * @generated from field: string port = 2;
   */
  port: string;

  constructor(data?: PartialMessage<Container>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Container";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Container;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Container;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Container;

  static equals(a: Container | PlainMessage<Container> | undefined, b: Container | PlainMessage<Container> | undefined): boolean;
}

/**
 * @generated from message porter.v1.Ingress
 */
export declare class Ingress extends Message<Ingress> {
  /**
   * @generated from field: bool enabled = 1;
   */
  enabled: boolean;

  /**
   * @generated from field: bool custom_domain = 2;
   */
  customDomain: boolean;

  /**
   * @generated from field: repeated string hosts = 3;
   */
  hosts: string[];

  /**
   * @generated from field: repeated string porter_hosts = 4;
   */
  porterHosts: string[];

  /**
   * @generated from field: map<string, string> annotations = 5;
   */
  annotations: { [key: string]: string };

  constructor(data?: PartialMessage<Ingress>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Ingress";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Ingress;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Ingress;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Ingress;

  static equals(a: Ingress | PlainMessage<Ingress> | undefined, b: Ingress | PlainMessage<Ingress> | undefined): boolean;
}

/**
 * @generated from message porter.v1.Autoscaling
 */
export declare class Autoscaling extends Message<Autoscaling> {
  /**
   * @generated from field: bool enabled = 1;
   */
  enabled: boolean;

  /**
   * @generated from field: int32 min_replicas = 2;
   */
  minReplicas: number;

  /**
   * @generated from field: int32 max_replicas = 3;
   */
  maxReplicas: number;

  /**
   * @generated from field: int32 target_cpu_utilization_threshold = 4;
   */
  targetCpuUtilizationThreshold: number;

  /**
   * @generated from field: int32 target_memory_utilization_threshold = 5;
   */
  targetMemoryUtilizationThreshold: number;

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
 * @generated from message porter.v1.CloudSql
 */
export declare class CloudSql extends Message<CloudSql> {
  /**
   * @generated from field: bool enabled = 1;
   */
  enabled: boolean;

  /**
   * @generated from field: string connection_name = 2;
   */
  connectionName: string;

  /**
   * @generated from field: string db_port = 3;
   */
  dbPort: string;

  /**
   * @generated from field: string service_account_json = 4;
   */
  serviceAccountJson: string;

  constructor(data?: PartialMessage<CloudSql>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.CloudSql";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CloudSql;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CloudSql;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CloudSql;

  static equals(a: CloudSql | PlainMessage<CloudSql> | undefined, b: CloudSql | PlainMessage<CloudSql> | undefined): boolean;
}

/**
 * @generated from message porter.v1.LiveCheck
 */
export declare class LiveCheck extends Message<LiveCheck> {
  /**
   * @generated from field: bool enabled = 1;
   */
  enabled: boolean;

  /**
   * @generated from field: int32 failure_threshold = 2;
   */
  failureThreshold: number;

  /**
   * @generated from field: string path = 3;
   */
  path: string;

  /**
   * @generated from field: int32 period_seconds = 4;
   */
  periodSeconds: number;

  constructor(data?: PartialMessage<LiveCheck>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.LiveCheck";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LiveCheck;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LiveCheck;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LiveCheck;

  static equals(a: LiveCheck | PlainMessage<LiveCheck> | undefined, b: LiveCheck | PlainMessage<LiveCheck> | undefined): boolean;
}

/**
 * @generated from message porter.v1.ReadyCheck
 */
export declare class ReadyCheck extends Message<ReadyCheck> {
  /**
   * @generated from field: bool enabled = 1;
   */
  enabled: boolean;

  /**
   * @generated from field: int32 failure_threshold = 2;
   */
  failureThreshold: number;

  /**
   * @generated from field: string path = 3;
   */
  path: string;

  /**
   * @generated from field: int32 initial_delay_seconds = 4;
   */
  initialDelaySeconds: number;

  constructor(data?: PartialMessage<ReadyCheck>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.ReadyCheck";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ReadyCheck;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ReadyCheck;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ReadyCheck;

  static equals(a: ReadyCheck | PlainMessage<ReadyCheck> | undefined, b: ReadyCheck | PlainMessage<ReadyCheck> | undefined): boolean;
}

/**
 * @generated from message porter.v1.Health
 */
export declare class Health extends Message<Health> {
  /**
   * @generated from field: porter.v1.LiveCheck startup_probe = 1;
   */
  startupProbe?: LiveCheck;

  /**
   * @generated from field: porter.v1.ReadyCheck readiness_probe = 2;
   */
  readinessProbe?: ReadyCheck;

  /**
   * @generated from field: porter.v1.LiveCheck liveness_probe = 3;
   */
  livenessProbe?: LiveCheck;

  constructor(data?: PartialMessage<Health>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Health";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Health;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Health;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Health;

  static equals(a: Health | PlainMessage<Health> | undefined, b: Health | PlainMessage<Health> | undefined): boolean;
}

