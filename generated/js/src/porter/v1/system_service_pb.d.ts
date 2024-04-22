// @generated by protoc-gen-es v1.9.0
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
 * ServiceStatus is the status for a single system service at a certain time stamp
 *
 * @generated from message porter.v1.ServiceStatus
 */
export declare class ServiceStatus extends Message<ServiceStatus> {
  /**
   * timestamp_field holds the timestamp for this status
   *
   * @generated from field: google.protobuf.Timestamp timestamp_field = 1;
   */
  timestampField?: Timestamp;

  /**
   * status gives the status for a service
   *
   * @generated from field: porter.v1.Status status = 2;
   */
  status: Status;

  constructor(data?: PartialMessage<ServiceStatus>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.ServiceStatus";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ServiceStatus;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ServiceStatus;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ServiceStatus;

  static equals(a: ServiceStatus | PlainMessage<ServiceStatus> | undefined, b: ServiceStatus | PlainMessage<ServiceStatus> | undefined): boolean;
}

/**
 * SystemServiceStatusHistory is a system service's status timeseries
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
   * status_history is a list of statuses for the service sorted in a descending order (so most recent comes first)
   *
   * @generated from field: repeated porter.v1.ServiceStatus status_history = 2;
   */
  statusHistory: ServiceStatus[];

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
 * ClusterStatus is the status for a cluster at a certain timestamp
 *
 * @generated from message porter.v1.ClusterStatus
 */
export declare class ClusterStatus extends Message<ClusterStatus> {
  /**
   * timestamp_field holds the timestamp for this status
   *
   * @generated from field: google.protobuf.Timestamp timestamp_field = 1;
   */
  timestampField?: Timestamp;

  /**
   * responsive is true if the cluster sent all the expected heartbeats in the time range represented by the timestamp
   *
   * @generated from field: bool responsive = 2;
   */
  responsive: boolean;

  constructor(data?: PartialMessage<ClusterStatus>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.ClusterStatus";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ClusterStatus;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ClusterStatus;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ClusterStatus;

  static equals(a: ClusterStatus | PlainMessage<ClusterStatus> | undefined, b: ClusterStatus | PlainMessage<ClusterStatus> | undefined): boolean;
}

