// @generated by protoc-gen-es v1.8.0
// @generated from file porter/v1/gke.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum porter.v1.GKENodePoolType
 */
export const GKENodePoolType = /*@__PURE__*/ proto3.makeEnum(
  "porter.v1.GKENodePoolType",
  [
    {no: 0, name: "GKE_NODE_POOL_TYPE_UNSPECIFIED"},
    {no: 1, name: "GKE_NODE_POOL_TYPE_SYSTEM"},
    {no: 2, name: "GKE_NODE_POOL_TYPE_MONITORING"},
    {no: 3, name: "GKE_NODE_POOL_TYPE_APPLICATION"},
    {no: 4, name: "GKE_NODE_POOL_TYPE_CUSTOM"},
  ],
);

/**
 * @generated from message porter.v1.GKE
 */
export const GKE = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.GKE",
  () => [
    { no: 1, name: "cluster_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "cluster_version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "network", kind: "message", T: GKENetwork },
    { no: 4, name: "region", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "node_pools", kind: "message", T: GKENodePool, repeated: true },
  ],
);

/**
 * @generated from message porter.v1.GKENodePool
 */
export const GKENodePool = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.GKENodePool",
  () => [
    { no: 1, name: "instance_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "min_instances", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 3, name: "max_instances", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 4, name: "node_pool_type", kind: "enum", T: proto3.getEnumType(GKENodePoolType) },
    { no: 5, name: "additional_taints", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 6, name: "is_spot_instance", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * GKENetwork contains all information required to configure the GKE cluster's network
 *
 * @generated from message porter.v1.GKENetwork
 */
export const GKENetwork = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.GKENetwork",
  () => [
    { no: 1, name: "cidr_range", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "control_plane_cidr", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "pod_cidr", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "service_cidr", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * GKEPreflightValues is cidr ranges needed for PreflightChecks
 *
 * @generated from message porter.v1.GKEPreflightValues
 */
export const GKEPreflightValues = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.GKEPreflightValues",
  () => [
    { no: 1, name: "network", kind: "message", T: GKENetwork },
  ],
);

