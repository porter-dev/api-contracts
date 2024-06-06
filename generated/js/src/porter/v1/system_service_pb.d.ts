// @generated by protoc-gen-es v1.10.0
// @generated from file porter/v1/system_service.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { InvolvedObjectType } from "./prometheus_alerts_pb.js";

/**
 * Status represents the health status of a service
 *
 * @generated from enum porter.v1.Status
 */
export declare enum Status {
  /**
   * STATUS_UNSPECIFIED is the default value
   *
   * @generated from enum value: STATUS_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * STATUS_HEALTHY is the value for a fully healthy service
   *
   * @generated from enum value: STATUS_HEALTHY = 1;
   */
  HEALTHY = 1,

  /**
   * STATUS_PARTIAL_FAILURE is the health status for a partially failed service
   * a service is in partial failure if only less than the max unavailable number of replicas declared on the service are failed
   *
   * @generated from enum value: STATUS_PARTIAL_FAILURE = 2;
   */
  PARTIAL_FAILURE = 2,

  /**
   * STATUS_FAILURE is the health status for a fully failed service
   *
   * @generated from enum value: STATUS_FAILURE = 3;
   */
  FAILURE = 3,
}

/**
 * ClusterHealthType represents types of health history we generate for the cluster
 *
 * @generated from enum porter.v1.ClusterHealthType
 */
export declare enum ClusterHealthType {
  /**
   * CLUSTER_HEALTH_TYPE_UNSPECIFIED is the default value
   *
   * @generated from enum value: CLUSTER_HEALTH_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * CLUSTER_HEALTH_TYPE_PINGABLE is history of whether we were able to ping the api-server
   *
   * @generated from enum value: CLUSTER_HEALTH_TYPE_CONNECTED = 1;
   */
  CONNECTED = 1,

  /**
   * CLUSTER_HEALTH_TYPE_CONNECTED is history of weather we were able to connect to the cluster
   *
   * @generated from enum value: CLUSTER_HEALTH_TYPE_PINGABLE = 2;
   */
  PINGABLE = 2,

  /**
   * CLUSTER_HEALTH_TYPE_METRICS_HEALTHY is whether the metrics services in the cluster were healthy
   *
   * @generated from enum value: CLUSTER_HEALTH_TYPE_METRICS_HEALTHY = 3;
   */
  METRICS_HEALTHY = 3,
}

/**
 * SystemService identifies a system service by its name, namespace and the type of k8s object it runs
 *
 * @generated from message porter.v1.SystemService
 */
export declare class SystemService extends Message<SystemService> {
  /**
   * name is the name of the service
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * namespace is the namespace where it is deployed
   *
   * @generated from field: string namespace = 2;
   */
  namespace: string;

  /**
   * involved_object_type is the k8s object type the service runs
   *
   * @generated from field: porter.v1.InvolvedObjectType involved_object_type = 3;
   */
  involvedObjectType: InvolvedObjectType;

  constructor(data?: PartialMessage<SystemService>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.SystemService";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SystemService;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SystemService;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SystemService;

  static equals(a: SystemService | PlainMessage<SystemService> | undefined, b: SystemService | PlainMessage<SystemService> | undefined): boolean;
}

/**
 * HealthStatus is a single status observed over a certain period of time
 *
 * @generated from message porter.v1.HealthStatus
 */
export declare class HealthStatus extends Message<HealthStatus> {
  /**
   * start_time holds the timestamp for when this status began
   *
   * @generated from field: google.protobuf.Timestamp start_time = 1;
   */
  startTime?: Timestamp;

  /**
   * end_time holds the timestamp for when this status ended if provided
   * if not provided, the status is considered ongoing
   *
   * @generated from field: optional google.protobuf.Timestamp end_time = 2;
   */
  endTime?: Timestamp;

  /**
   * status gives the status for a service
   *
   * @generated from field: porter.v1.Status status = 3;
   */
  status: Status;

  /**
   * description provides more information on a status
   *
   * @generated from field: string description = 4;
   */
  description: string;

  constructor(data?: PartialMessage<HealthStatus>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.HealthStatus";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): HealthStatus;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): HealthStatus;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): HealthStatus;

  static equals(a: HealthStatus | PlainMessage<HealthStatus> | undefined, b: HealthStatus | PlainMessage<HealthStatus> | undefined): boolean;
}

/**
 * StatusPercentage is the percentage in a day that a certain status was observed
 *
 * @generated from message porter.v1.StatusPercentage
 */
export declare class StatusPercentage extends Message<StatusPercentage> {
  /**
   * @generated from field: porter.v1.Status status = 1;
   */
  status: Status;

  /**
   * @generated from field: float percentage = 2;
   */
  percentage: number;

  constructor(data?: PartialMessage<StatusPercentage>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.StatusPercentage";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StatusPercentage;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StatusPercentage;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StatusPercentage;

  static equals(a: StatusPercentage | PlainMessage<StatusPercentage> | undefined, b: StatusPercentage | PlainMessage<StatusPercentage> | undefined): boolean;
}

/**
 * DailyHealthStatus represents health status over a day
 *
 * @generated from message porter.v1.DailyHealthStatus
 */
export declare class DailyHealthStatus extends Message<DailyHealthStatus> {
  /**
   * status_percentages holds information on what percentage of the day different statuses were observed
   * there should be only one entry for the different Status types
   *
   * @generated from field: repeated porter.v1.StatusPercentage status_percentages = 1;
   */
  statusPercentages: StatusPercentage[];

  /**
   * health_statuses is the different statuses observed over a day
   * if there is a missing time gap, the service can be considered to have been healthy during that gap
   *
   * @generated from field: repeated porter.v1.HealthStatus health_statuses = 2;
   */
  healthStatuses: HealthStatus[];

  constructor(data?: PartialMessage<DailyHealthStatus>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.DailyHealthStatus";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DailyHealthStatus;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DailyHealthStatus;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DailyHealthStatus;

  static equals(a: DailyHealthStatus | PlainMessage<DailyHealthStatus> | undefined, b: DailyHealthStatus | PlainMessage<DailyHealthStatus> | undefined): boolean;
}

/**
 * SystemServiceStatusHistory holds the health status history for a particular system service over multiple days
 *
 * @generated from message porter.v1.SystemServiceStatusHistory
 */
export declare class SystemServiceStatusHistory extends Message<SystemServiceStatusHistory> {
  /**
   * system_service identifies the service
   *
   * @generated from field: porter.v1.SystemService system_service = 1;
   */
  systemService?: SystemService;

  /**
   * daily_status_history is daily health status for a service over multiple days
   * the keys are the number of days before today the DailyHealthStatus is from
   *
   * @generated from field: map<int32, porter.v1.DailyHealthStatus> daily_status_history = 2;
   */
  dailyStatusHistory: { [key: number]: DailyHealthStatus };

  constructor(data?: PartialMessage<SystemServiceStatusHistory>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.SystemServiceStatusHistory";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SystemServiceStatusHistory;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SystemServiceStatusHistory;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SystemServiceStatusHistory;

  static equals(a: SystemServiceStatusHistory | PlainMessage<SystemServiceStatusHistory> | undefined, b: SystemServiceStatusHistory | PlainMessage<SystemServiceStatusHistory> | undefined): boolean;
}

/**
 * ClusterStatusHistory holds health status history of a certain type for a cluster
 *
 * @generated from message porter.v1.ClusterStatusHistory
 */
export declare class ClusterStatusHistory extends Message<ClusterStatusHistory> {
  /**
   * cluster_health_type is the type this history is generated for
   *
   * @generated from field: porter.v1.ClusterHealthType cluster_health_type = 1;
   */
  clusterHealthType: ClusterHealthType;

  /**
   * daily_status_history is daily health status for the cluster over multiple days for the cluster_health_type
   * the keys are the number of days before today the DailyHealthStatus is from
   *
   * @generated from field: map<int32, porter.v1.DailyHealthStatus> daily_status_history = 2;
   */
  dailyStatusHistory: { [key: number]: DailyHealthStatus };

  constructor(data?: PartialMessage<ClusterStatusHistory>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.ClusterStatusHistory";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ClusterStatusHistory;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ClusterStatusHistory;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ClusterStatusHistory;

  static equals(a: ClusterStatusHistory | PlainMessage<ClusterStatusHistory> | undefined, b: ClusterStatusHistory | PlainMessage<ClusterStatusHistory> | undefined): boolean;
}

