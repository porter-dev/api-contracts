// @generated by protoc-gen-es v1.2.0
// @generated from file porter/v1/aks.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum porter.v1.NodePoolType
 */
export declare enum NodePoolType {
  /**
   * @generated from enum value: NODE_POOL_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: NODE_POOL_TYPE_SYSTEM = 1;
   */
  SYSTEM = 1,

  /**
   * @generated from enum value: NODE_POOL_TYPE_MONITORING = 2;
   */
  MONITORING = 2,

  /**
   * @generated from enum value: NODE_POOL_TYPE_APPLICATION = 3;
   */
  APPLICATION = 3,

  /**
   * @generated from enum value: NODE_POOL_TYPE_CUSTOM = 4;
   */
  CUSTOM = 4,
}

/**
 * @generated from message porter.v1.AKS
 */
export declare class AKS extends Message<AKS> {
  /**
   * @generated from field: string cluster_name = 1;
   */
  clusterName: string;

  /**
   * @generated from field: string cluster_version = 2;
   */
  clusterVersion: string;

  /**
   * @generated from field: string cidr_range = 3;
   */
  cidrRange: string;

  /**
   * @generated from field: string location = 4;
   */
  location: string;

  /**
   * @generated from field: repeated porter.v1.AKSNodePool node_pools = 5;
   */
  nodePools: AKSNodePool[];

  constructor(data?: PartialMessage<AKS>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.AKS";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AKS;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AKS;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AKS;

  static equals(a: AKS | PlainMessage<AKS> | undefined, b: AKS | PlainMessage<AKS> | undefined): boolean;
}

/**
 * @generated from message porter.v1.AKSNodePool
 */
export declare class AKSNodePool extends Message<AKSNodePool> {
  /**
   * @generated from field: string instance_type = 1;
   */
  instanceType: string;

  /**
   * @generated from field: uint32 min_instances = 2;
   */
  minInstances: number;

  /**
   * @generated from field: uint32 max_instances = 3;
   */
  maxInstances: number;

  /**
   * @generated from field: string mode = 4;
   */
  mode: string;

  /**
   * @generated from field: porter.v1.NodePoolType node_pool_type = 5;
   */
  nodePoolType: NodePoolType;

  constructor(data?: PartialMessage<AKSNodePool>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.AKSNodePool";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AKSNodePool;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AKSNodePool;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AKSNodePool;

  static equals(a: AKSNodePool | PlainMessage<AKSNodePool> | undefined, b: AKSNodePool | PlainMessage<AKSNodePool> | undefined): boolean;
}
