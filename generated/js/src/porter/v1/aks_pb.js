// @generated by protoc-gen-es v1.10.0
// @generated from file porter/v1/aks.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum porter.v1.AksSkuTier
 */
export const AksSkuTier = /*@__PURE__*/ proto3.makeEnum(
  "porter.v1.AksSkuTier",
  [
    {no: 0, name: "AKS_SKU_TIER_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "AKS_SKU_TIER_FREE", localName: "FREE"},
    {no: 2, name: "AKS_SKU_TIER_STANDARD", localName: "STANDARD"},
  ],
);

/**
 * @generated from enum porter.v1.NodePoolType
 */
export const NodePoolType = /*@__PURE__*/ proto3.makeEnum(
  "porter.v1.NodePoolType",
  [
    {no: 0, name: "NODE_POOL_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "NODE_POOL_TYPE_SYSTEM", localName: "SYSTEM"},
    {no: 2, name: "NODE_POOL_TYPE_MONITORING", localName: "MONITORING"},
    {no: 3, name: "NODE_POOL_TYPE_APPLICATION", localName: "APPLICATION"},
    {no: 4, name: "NODE_POOL_TYPE_CUSTOM", localName: "CUSTOM"},
    {no: 5, name: "NODE_POOL_TYPE_USER", localName: "USER"},
  ],
);

/**
 * @generated from message porter.v1.AKS
 */
export const AKS = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.AKS",
  () => [
    { no: 1, name: "cluster_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "cluster_version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "cidr_range", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "location", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "node_pools", kind: "message", T: AKSNodePool, repeated: true },
    { no: 6, name: "sku_tier", kind: "enum", T: proto3.getEnumType(AksSkuTier) },
    { no: 7, name: "enable_container_insights", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 8, name: "enable_rbac", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * @generated from message porter.v1.AKSNodePool
 */
export const AKSNodePool = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.AKSNodePool",
  () => [
    { no: 1, name: "instance_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "min_instances", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 3, name: "max_instances", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 4, name: "mode", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "node_pool_type", kind: "enum", T: proto3.getEnumType(NodePoolType) },
    { no: 6, name: "additional_taints", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 7, name: "node_pool_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "node_pool_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

