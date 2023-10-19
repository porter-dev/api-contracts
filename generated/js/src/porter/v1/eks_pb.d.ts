// @generated by protoc-gen-es v1.3.3
// @generated from file porter/v1/eks.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum porter.v1.NodeGroupType
 */
export declare enum NodeGroupType {
  /**
   * @generated from enum value: NODE_GROUP_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: NODE_GROUP_TYPE_SYSTEM = 1;
   */
  SYSTEM = 1,

  /**
   * @generated from enum value: NODE_GROUP_TYPE_MONITORING = 2;
   */
  MONITORING = 2,

  /**
   * @generated from enum value: NODE_GROUP_TYPE_APPLICATION = 3;
   */
  APPLICATION = 3,

  /**
   * @generated from enum value: NODE_GROUP_TYPE_CUSTOM = 4;
   */
  CUSTOM = 4,
}

/**
 * @generated from enum porter.v1.LoadBalancerType
 */
export declare enum LoadBalancerType {
  /**
   * @generated from enum value: LOAD_BALANCER_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: LOAD_BALANCER_TYPE_NLB = 1;
   */
  NLB = 1,

  /**
   * @generated from enum value: LOAD_BALANCER_TYPE_ALB = 2;
   */
  ALB = 2,
}

/**
 * @generated from message porter.v1.EKS
 */
export declare class EKS extends Message<EKS> {
  /**
   * @generated from field: string cluster_name = 1;
   */
  clusterName: string;

  /**
   * @generated from field: string cluster_version = 2;
   */
  clusterVersion: string;

  /**
   * use network.vpc_cidr instead
   *
   * @generated from field: string cidr_range = 3 [deprecated = true];
   * @deprecated
   */
  cidrRange: string;

  /**
   * @generated from field: string region = 4;
   */
  region: string;

  /**
   * @generated from field: repeated porter.v1.EKSNodeGroup node_groups = 5;
   */
  nodeGroups: EKSNodeGroup[];

  /**
   * @generated from field: porter.v1.LoadBalancer load_balancer = 6;
   */
  loadBalancer?: LoadBalancer;

  /**
   * @generated from field: bool enable_guard_duty = 7;
   */
  enableGuardDuty: boolean;

  /**
   * @generated from field: porter.v1.EKSLogging logging = 8;
   */
  logging?: EKSLogging;

  /**
   * enable_kms_encryption triggers a KMS key creation and enables encryption on the EKS cluster with that key.  Once enabled, encryption can never be disabled.
   *
   * @generated from field: bool enable_kms_encryption = 9;
   */
  enableKmsEncryption: boolean;

  /**
   * network is the network configuration for the EKS cluster.
   * If both cidr_range and network.vpc_cidr are set, network.vpc_cidr will be used.
   *
   * @generated from field: porter.v1.AWSClusterNetwork network = 10;
   */
  network?: AWSClusterNetwork;

  constructor(data?: PartialMessage<EKS>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.EKS";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EKS;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EKS;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EKS;

  static equals(a: EKS | PlainMessage<EKS> | undefined, b: EKS | PlainMessage<EKS> | undefined): boolean;
}

/**
 * AWSClusterNetwork contains all information required to configure the AWS cluster's network
 *
 * @generated from message porter.v1.AWSClusterNetwork
 */
export declare class AWSClusterNetwork extends Message<AWSClusterNetwork> {
  /**
   * vpc_cidr is the range of the VPC network. This is used to specify the network that Porter will use.
   *
   * @generated from field: string vpc_cidr = 1;
   */
  vpcCidr: string;

  /**
   * service_cidr is the range of the network that services will be assigned IPs from, on the AWS vpc.
   *
   * @generated from field: string service_cidr = 2;
   */
  serviceCidr: string;

  constructor(data?: PartialMessage<AWSClusterNetwork>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.AWSClusterNetwork";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AWSClusterNetwork;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AWSClusterNetwork;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AWSClusterNetwork;

  static equals(a: AWSClusterNetwork | PlainMessage<AWSClusterNetwork> | undefined, b: AWSClusterNetwork | PlainMessage<AWSClusterNetwork> | undefined): boolean;
}

/**
 * EKSNodeGroup is the configuration for an EKS node group/auto scaling group
 *
 * @generated from message porter.v1.EKSNodeGroup
 */
export declare class EKSNodeGroup extends Message<EKSNodeGroup> {
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
   * node_group_type is used to specify the type of node group. This is used to specify what node groups are used by Porter, vs users.
   *
   * @generated from field: porter.v1.NodeGroupType node_group_type = 4;
   */
  nodeGroupType: NodeGroupType;

  /**
   * is_stateful is deprecated. It was initially used to specify if a nodegroup had state, and needed to be kept in a single AZ.
   *
   * @generated from field: bool is_stateful = 5 [deprecated = true];
   * @deprecated
   */
  isStateful: boolean;

  /**
   * additional_policies is a list of IAM policies to attach to the node group role, on top of the policies applied by Porter.
   *
   * @generated from field: repeated string additional_policies = 6;
   */
  additionalPolicies: string[];

  constructor(data?: PartialMessage<EKSNodeGroup>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.EKSNodeGroup";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EKSNodeGroup;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EKSNodeGroup;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EKSNodeGroup;

  static equals(a: EKSNodeGroup | PlainMessage<EKSNodeGroup> | undefined, b: EKSNodeGroup | PlainMessage<EKSNodeGroup> | undefined): boolean;
}

/**
 * @generated from message porter.v1.LoadBalancer
 */
export declare class LoadBalancer extends Message<LoadBalancer> {
  /**
   * load_balancer_type is the type of load balancer to deploy. If unspecified, this will be assumed to be NLB for AWS
   *
   * @generated from field: porter.v1.LoadBalancerType load_balancer_type = 1;
   */
  loadBalancerType: LoadBalancerType;

  /**
   * wildcard_domain is used in ALB loadbalancer deployment when allowlisting IPs. This has no effect when load_balancer_type is NLB
   *
   * @generated from field: string wildcard_domain = 2;
   */
  wildcardDomain: string;

  /**
   * allowlist_ip_ranges are comma separated CIDRS, which are the only ranges who will be granted access to ALB ingress resources. This has no effect when load_balancer_type is NLB
   *
   * @generated from field: string allowlist_ip_ranges = 3;
   */
  allowlistIpRanges: string;

  /**
   * enable_wafv2 enables WAFv2 on the ALB. This has no effect when load_balancer_type is NLB
   *
   * @generated from field: bool enable_wafv2 = 4;
   */
  enableWafv2: boolean;

  /**
   * wafv2_arn is the ARN of the WAFv2 ACL to attach to the ALB. This has no effect when load_balancer_type is NLB, or if enable_wafv2 is false
   *
   * @generated from field: string wafv2_arn = 5;
   */
  wafv2Arn: string;

  /**
   * additional_certificate_arns is a list of ACM certificate ARNs to attach to the ALB. This has no effect when load_balancer_type is NLB. These will be added in order, before the ACM certificate created by Porter.
   *
   * @generated from field: repeated string additional_certificate_arns = 6;
   */
  additionalCertificateArns: string[];

  /**
   * tags is a map of AWS tags to apply to the ALB. This has no effect when load_balancer_type is NLB.
   *
   * @generated from field: map<string, string> tags = 7;
   */
  tags: { [key: string]: string };

  /**
   * enable_s3_access_logs enables S3 access logs on the ALB. This has no effect when load_balancer_type is NLB.
   *
   * @generated from field: bool enable_s3_access_logs = 8;
   */
  enableS3AccessLogs: boolean;

  constructor(data?: PartialMessage<LoadBalancer>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.LoadBalancer";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LoadBalancer;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LoadBalancer;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LoadBalancer;

  static equals(a: LoadBalancer | PlainMessage<LoadBalancer> | undefined, b: LoadBalancer | PlainMessage<LoadBalancer> | undefined): boolean;
}

/**
 * EKSLogging is the configuration for EKS control plane logging which is typically managed through the AWS Console. This will create a CloudWatch log group and log stream for the enabled logs.
 *
 * @generated from message porter.v1.EKSLogging
 */
export declare class EKSLogging extends Message<EKSLogging> {
  /**
   * @generated from field: bool enable_api_server_logs = 1;
   */
  enableApiServerLogs: boolean;

  /**
   * @generated from field: bool enable_authenticator_logs = 2;
   */
  enableAuthenticatorLogs: boolean;

  /**
   * @generated from field: bool enable_scheduler_logs = 3;
   */
  enableSchedulerLogs: boolean;

  /**
   * @generated from field: bool enable_audit_logs = 4;
   */
  enableAuditLogs: boolean;

  /**
   * @generated from field: bool enable_controller_manager_logs = 5;
   */
  enableControllerManagerLogs: boolean;

  constructor(data?: PartialMessage<EKSLogging>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.EKSLogging";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EKSLogging;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EKSLogging;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EKSLogging;

  static equals(a: EKSLogging | PlainMessage<EKSLogging> | undefined, b: EKSLogging | PlainMessage<EKSLogging> | undefined): boolean;
}

/**
 * EKSPreflightValues contains all needed values to perform EKS Preflight Checks
 *
 * @generated from message porter.v1.EKSPreflightValues
 */
export declare class EKSPreflightValues extends Message<EKSPreflightValues> {
  /**
   * region the region to perform the preflight checks in
   *
   * @generated from field: string region = 1;
   */
  region: string;

  constructor(data?: PartialMessage<EKSPreflightValues>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.EKSPreflightValues";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EKSPreflightValues;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EKSPreflightValues;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EKSPreflightValues;

  static equals(a: EKSPreflightValues | PlainMessage<EKSPreflightValues> | undefined, b: EKSPreflightValues | PlainMessage<EKSPreflightValues> | undefined): boolean;
}

/**
 * AWSVpc contains all the properties representing an AWS VPC
 *
 * @generated from message porter.v1.AWSVpc
 */
export declare class AWSVpc extends Message<AWSVpc> {
  /**
   * id is a unique vpc identifier for an AWS VPC and can be used to reference the VPC in api calls
   *
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<AWSVpc>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.AWSVpc";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AWSVpc;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AWSVpc;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AWSVpc;

  static equals(a: AWSVpc | PlainMessage<AWSVpc> | undefined, b: AWSVpc | PlainMessage<AWSVpc> | undefined): boolean;
}

