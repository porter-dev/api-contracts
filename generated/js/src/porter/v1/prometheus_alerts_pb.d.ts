// @generated by protoc-gen-es v1.9.0
// @generated from file porter/v1/prometheus_alerts.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * InvolvedObjectType is the kubernetes object type the notification targets
 * We currently alert for deployments, statefulsets and daemonsets
 *
 * @generated from enum porter.v1.InvolvedObjectType
 */
export declare enum InvolvedObjectType {
  /**
   * @generated from enum value: INVOLVED_OBJECT_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: INVOLVED_OBJECT_TYPE_DEPLOYMENT = 1;
   */
  DEPLOYMENT = 1,

  /**
   * @generated from enum value: INVOLVED_OBJECT_TYPE_STATEFULSET = 2;
   */
  STATEFULSET = 2,

  /**
   * @generated from enum value: INVOLVED_OBJECT_TYPE_DAEMONSET = 3;
   */
  DAEMONSET = 3,
}

/**
 * Alert represents a prometheus alert for one target object that is a daemonset, statefulset or deployment
 *
 * @generated from message porter.v1.Alert
 */
export declare class Alert extends Message<Alert> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: string namespace = 2;
   */
  namespace: string;

  /**
   * @generated from field: porter.v1.InvolvedObjectType type = 3;
   */
  type: InvolvedObjectType;

  /**
   * @generated from field: int32 desired_replicas = 4;
   */
  desiredReplicas: number;

  /**
   * @generated from field: int32 available_replicas = 5;
   */
  availableReplicas: number;

  /**
   * @generated from field: int32 max_unavailable = 6;
   */
  maxUnavailable: number;

  constructor(data?: PartialMessage<Alert>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Alert";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Alert;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Alert;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Alert;

  static equals(a: Alert | PlainMessage<Alert> | undefined, b: Alert | PlainMessage<Alert> | undefined): boolean;
}

