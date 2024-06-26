// @generated by protoc-gen-es v1.10.0
// @generated from file porter/v1/aks.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum porter.v1.AksSkuTier
 */
export declare enum AksSkuTier {
  /**
   * @generated from enum value: AKS_SKU_TIER_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: AKS_SKU_TIER_FREE = 1;
   */
  FREE = 1,

  /**
   * @generated from enum value: AKS_SKU_TIER_STANDARD = 2;
   */
  STANDARD = 2,
}

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

  /**
   * NODE_POOL_TYPE_USER indicates a user-managed node pool (can be created, edited, and deleted by users). Node pools of this type must specify a unique node_pool_name.
   *
   * @generated from enum value: NODE_POOL_TYPE_USER = 5;
   */
  USER = 5,
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

  /**
   * sku_tier is the SKU tier of the cluster (Free or Standard).  Standard should be used for all production workloads that require high availability.
   *
   * @generated from field: porter.v1.AksSkuTier sku_tier = 6;
   */
  skuTier: AksSkuTier;

  /**
   * enable_container_insights enables Azure Monitor for containers on the cluster.
   *
   * @generated from field: bool enable_container_insights = 7;
   */
  enableContainerInsights: boolean;

  /**
   * enable_rbac enables Azure RBAC on the cluster.
   *
   * @generated from field: bool enable_rbac = 8;
   */
  enableRbac: boolean;

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

  /**
   * additional_taints is a list of NoSchedule taints to apply to the node pool.
   * These will be applied on top of the default porter.run/workload-kind taints.
   *
   * @generated from field: repeated string additional_taints = 6;
   */
  additionalTaints: string[];

  /**
   * node_pool_id is the id of the node pool. This uniquely identifies NodePoolType=User and is generated on creation if not provided.
   *
   * @generated from field: string node_pool_id = 7;
   */
  nodePoolId: string;

  /**
   * node_pool_name is the vanity name of the node pool. This is required for NodePoolType=User and can be changed by the user.
   *
   * @generated from field: string node_pool_name = 8;
   */
  nodePoolName: string;

  constructor(data?: PartialMessage<AKSNodePool>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.AKSNodePool";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AKSNodePool;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AKSNodePool;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AKSNodePool;

  static equals(a: AKSNodePool | PlainMessage<AKSNodePool> | undefined, b: AKSNodePool | PlainMessage<AKSNodePool> | undefined): boolean;
}

