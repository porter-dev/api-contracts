// @generated by protoc-gen-es v1.4.2
// @generated from file porter/v1/eks.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum porter.v1.NodeGroupType
 */
export const NodeGroupType = proto3.makeEnum(
  "porter.v1.NodeGroupType",
  [
    {no: 0, name: "NODE_GROUP_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "NODE_GROUP_TYPE_SYSTEM", localName: "SYSTEM"},
    {no: 2, name: "NODE_GROUP_TYPE_MONITORING", localName: "MONITORING"},
    {no: 3, name: "NODE_GROUP_TYPE_APPLICATION", localName: "APPLICATION"},
    {no: 4, name: "NODE_GROUP_TYPE_CUSTOM", localName: "CUSTOM"},
  ],
);

/**
 * @generated from enum porter.v1.LoadBalancerType
 */
export const LoadBalancerType = proto3.makeEnum(
  "porter.v1.LoadBalancerType",
  [
    {no: 0, name: "LOAD_BALANCER_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "LOAD_BALANCER_TYPE_NLB", localName: "NLB"},
    {no: 2, name: "LOAD_BALANCER_TYPE_ALB", localName: "ALB"},
  ],
);

/**
 * @generated from message porter.v1.EKS
 */
export const EKS = proto3.makeMessageType(
  "porter.v1.EKS",
  () => [
    { no: 1, name: "cluster_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "cluster_version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "cidr_range", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "region", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "node_groups", kind: "message", T: EKSNodeGroup, repeated: true },
    { no: 6, name: "load_balancer", kind: "message", T: LoadBalancer },
    { no: 7, name: "enable_guard_duty", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 8, name: "logging", kind: "message", T: EKSLogging },
    { no: 9, name: "enable_kms_encryption", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 10, name: "network", kind: "message", T: AWSClusterNetwork },
    { no: 11, name: "enable_ecr_scanning", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * AWSClusterNetwork contains all information required to configure the AWS cluster's network
 *
 * @generated from message porter.v1.AWSClusterNetwork
 */
export const AWSClusterNetwork = proto3.makeMessageType(
  "porter.v1.AWSClusterNetwork",
  () => [
    { no: 1, name: "vpc_cidr", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "service_cidr", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * EKSNodeGroup is the configuration for an EKS node group/auto scaling group
 *
 * @generated from message porter.v1.EKSNodeGroup
 */
export const EKSNodeGroup = proto3.makeMessageType(
  "porter.v1.EKSNodeGroup",
  () => [
    { no: 1, name: "instance_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "min_instances", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 3, name: "max_instances", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 4, name: "node_group_type", kind: "enum", T: proto3.getEnumType(NodeGroupType) },
    { no: 5, name: "is_stateful", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 6, name: "additional_policies", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 7, name: "additional_taints", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message porter.v1.LoadBalancer
 */
export const LoadBalancer = proto3.makeMessageType(
  "porter.v1.LoadBalancer",
  () => [
    { no: 1, name: "load_balancer_type", kind: "enum", T: proto3.getEnumType(LoadBalancerType) },
    { no: 2, name: "wildcard_domain", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "allowlist_ip_ranges", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "enable_wafv2", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 5, name: "wafv2_arn", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "additional_certificate_arns", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 7, name: "tags", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 8, name: "enable_s3_access_logs", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * EKSLogging is the configuration for EKS control plane logging which is typically managed through the AWS Console. This will create a CloudWatch log group and log stream for the enabled logs.
 *
 * @generated from message porter.v1.EKSLogging
 */
export const EKSLogging = proto3.makeMessageType(
  "porter.v1.EKSLogging",
  () => [
    { no: 1, name: "enable_api_server_logs", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "enable_authenticator_logs", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "enable_scheduler_logs", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 4, name: "enable_audit_logs", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 5, name: "enable_controller_manager_logs", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 6, name: "enable_cloudwatch_logs_to_s3", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * EKSPreflightValues contains all needed values to perform EKS Preflight Checks
 *
 * @generated from message porter.v1.EKSPreflightValues
 */
export const EKSPreflightValues = proto3.makeMessageType(
  "porter.v1.EKSPreflightValues",
  () => [
    { no: 1, name: "region", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * AWSVpc contains all the properties representing an AWS VPC and any networking-related information
 *
 * @generated from message porter.v1.AWSVpc
 */
export const AWSVpc = proto3.makeMessageType(
  "porter.v1.AWSVpc",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "cidr_range", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "subnets", kind: "message", T: AWSSubnet, repeated: true },
  ],
);

/**
 * AWSSubnet contains all the properties representing a single subnet in an AWS VPC
 *
 * @generated from message porter.v1.AWSSubnet
 */
export const AWSSubnet = proto3.makeMessageType(
  "porter.v1.AWSSubnet",
  () => [
    { no: 1, name: "availability_zone", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "cidr_range", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

