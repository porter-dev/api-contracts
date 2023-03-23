// @generated by protoc-gen-es v1.1.0
// @generated from file porter/v1/cluster_control_plane.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { AssumeRoleChainLink } from "./aws_assume_role_pb.js";
import type { Contract, ContractRevision } from "./contract_pb.js";

/**
 * @generated from message porter.v1.RolePreflightCheckRequest
 */
export declare class RolePreflightCheckRequest extends Message<RolePreflightCheckRequest> {
  /**
   * @generated from field: int64 project_id = 1;
   */
  projectId: bigint;

  /**
   * @generated from field: string target_arn = 2;
   */
  targetArn: string;

  /**
   * @generated from field: string external_id = 3;
   */
  externalId: string;

  constructor(data?: PartialMessage<RolePreflightCheckRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.RolePreflightCheckRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RolePreflightCheckRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RolePreflightCheckRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RolePreflightCheckRequest;

  static equals(a: RolePreflightCheckRequest | PlainMessage<RolePreflightCheckRequest> | undefined, b: RolePreflightCheckRequest | PlainMessage<RolePreflightCheckRequest> | undefined): boolean;
}

/**
 * @generated from message porter.v1.RolePreflightCheckResponse
 */
export declare class RolePreflightCheckResponse extends Message<RolePreflightCheckResponse> {
  /**
   * @generated from field: int64 project_id = 1;
   */
  projectId: bigint;

  /**
   * @generated from field: string role_name = 2;
   */
  roleName: string;

  constructor(data?: PartialMessage<RolePreflightCheckResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.RolePreflightCheckResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RolePreflightCheckResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RolePreflightCheckResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RolePreflightCheckResponse;

  static equals(a: RolePreflightCheckResponse | PlainMessage<RolePreflightCheckResponse> | undefined, b: RolePreflightCheckResponse | PlainMessage<RolePreflightCheckResponse> | undefined): boolean;
}

/**
 * @generated from message porter.v1.CreateAssumeRoleChainRequest
 */
export declare class CreateAssumeRoleChainRequest extends Message<CreateAssumeRoleChainRequest> {
  /**
   * @generated from field: int64 project_id = 1;
   */
  projectId: bigint;

  /**
   * @generated from field: string source_arn = 2;
   */
  sourceArn: string;

  /**
   * @generated from field: string target_access_id = 3;
   */
  targetAccessId: string;

  /**
   * @generated from field: string target_secret_key = 4;
   */
  targetSecretKey: string;

  /**
   * @generated from field: string target_session_token = 5;
   */
  targetSessionToken: string;

  constructor(data?: PartialMessage<CreateAssumeRoleChainRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.CreateAssumeRoleChainRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateAssumeRoleChainRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateAssumeRoleChainRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateAssumeRoleChainRequest;

  static equals(a: CreateAssumeRoleChainRequest | PlainMessage<CreateAssumeRoleChainRequest> | undefined, b: CreateAssumeRoleChainRequest | PlainMessage<CreateAssumeRoleChainRequest> | undefined): boolean;
}

/**
 * @generated from message porter.v1.CreateAssumeRoleChainResponse
 */
export declare class CreateAssumeRoleChainResponse extends Message<CreateAssumeRoleChainResponse> {
  /**
   * @generated from field: int64 project_id = 1;
   */
  projectId: bigint;

  /**
   * @generated from field: string target_arn = 2;
   */
  targetArn: string;

  constructor(data?: PartialMessage<CreateAssumeRoleChainResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.CreateAssumeRoleChainResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateAssumeRoleChainResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateAssumeRoleChainResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateAssumeRoleChainResponse;

  static equals(a: CreateAssumeRoleChainResponse | PlainMessage<CreateAssumeRoleChainResponse> | undefined, b: CreateAssumeRoleChainResponse | PlainMessage<CreateAssumeRoleChainResponse> | undefined): boolean;
}

/**
 * @generated from message porter.v1.EKSBearerTokenRequest
 */
export declare class EKSBearerTokenRequest extends Message<EKSBearerTokenRequest> {
  /**
   * @generated from field: int64 project_id = 1;
   */
  projectId: bigint;

  /**
   * @generated from field: int64 cluster_id = 2;
   */
  clusterId: bigint;

  constructor(data?: PartialMessage<EKSBearerTokenRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.EKSBearerTokenRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EKSBearerTokenRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EKSBearerTokenRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EKSBearerTokenRequest;

  static equals(a: EKSBearerTokenRequest | PlainMessage<EKSBearerTokenRequest> | undefined, b: EKSBearerTokenRequest | PlainMessage<EKSBearerTokenRequest> | undefined): boolean;
}

/**
 * @generated from message porter.v1.EKSBearerTokenResponse
 */
export declare class EKSBearerTokenResponse extends Message<EKSBearerTokenResponse> {
  /**
   * @generated from field: string token = 1;
   */
  token: string;

  constructor(data?: PartialMessage<EKSBearerTokenResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.EKSBearerTokenResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EKSBearerTokenResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EKSBearerTokenResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EKSBearerTokenResponse;

  static equals(a: EKSBearerTokenResponse | PlainMessage<EKSBearerTokenResponse> | undefined, b: EKSBearerTokenResponse | PlainMessage<EKSBearerTokenResponse> | undefined): boolean;
}

/**
 * @generated from message porter.v1.AssumeRoleChainTargetsRequest
 */
export declare class AssumeRoleChainTargetsRequest extends Message<AssumeRoleChainTargetsRequest> {
  /**
   * @generated from field: string project_id = 1;
   */
  projectId: string;

  constructor(data?: PartialMessage<AssumeRoleChainTargetsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.AssumeRoleChainTargetsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AssumeRoleChainTargetsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AssumeRoleChainTargetsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AssumeRoleChainTargetsRequest;

  static equals(a: AssumeRoleChainTargetsRequest | PlainMessage<AssumeRoleChainTargetsRequest> | undefined, b: AssumeRoleChainTargetsRequest | PlainMessage<AssumeRoleChainTargetsRequest> | undefined): boolean;
}

/**
 * @generated from message porter.v1.AssumeRoleChainTargetsResponse
 */
export declare class AssumeRoleChainTargetsResponse extends Message<AssumeRoleChainTargetsResponse> {
  /**
   * @generated from field: repeated porter.v1.AssumeRoleChainLink chain_links = 1;
   */
  chainLinks: AssumeRoleChainLink[];

  constructor(data?: PartialMessage<AssumeRoleChainTargetsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.AssumeRoleChainTargetsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AssumeRoleChainTargetsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AssumeRoleChainTargetsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AssumeRoleChainTargetsResponse;

  static equals(a: AssumeRoleChainTargetsResponse | PlainMessage<AssumeRoleChainTargetsResponse> | undefined, b: AssumeRoleChainTargetsResponse | PlainMessage<AssumeRoleChainTargetsResponse> | undefined): boolean;
}

/**
 * @generated from message porter.v1.KubeConfigForClusterRequest
 */
export declare class KubeConfigForClusterRequest extends Message<KubeConfigForClusterRequest> {
  /**
   * @generated from field: int64 project_id = 1;
   */
  projectId: bigint;

  /**
   * @generated from field: int64 cluster_id = 2;
   */
  clusterId: bigint;

  constructor(data?: PartialMessage<KubeConfigForClusterRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.KubeConfigForClusterRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): KubeConfigForClusterRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): KubeConfigForClusterRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): KubeConfigForClusterRequest;

  static equals(a: KubeConfigForClusterRequest | PlainMessage<KubeConfigForClusterRequest> | undefined, b: KubeConfigForClusterRequest | PlainMessage<KubeConfigForClusterRequest> | undefined): boolean;
}

/**
 * @generated from message porter.v1.KubeConfigForClusterResponse
 */
export declare class KubeConfigForClusterResponse extends Message<KubeConfigForClusterResponse> {
  /**
   * @generated from field: string kube_config = 1;
   */
  kubeConfig: string;

  constructor(data?: PartialMessage<KubeConfigForClusterResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.KubeConfigForClusterResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): KubeConfigForClusterResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): KubeConfigForClusterResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): KubeConfigForClusterResponse;

  static equals(a: KubeConfigForClusterResponse | PlainMessage<KubeConfigForClusterResponse> | undefined, b: KubeConfigForClusterResponse | PlainMessage<KubeConfigForClusterResponse> | undefined): boolean;
}

/**
 * @generated from message porter.v1.UpdateContractRequest
 */
export declare class UpdateContractRequest extends Message<UpdateContractRequest> {
  /**
   * @generated from field: porter.v1.Contract contract = 1;
   */
  contract?: Contract;

  constructor(data?: PartialMessage<UpdateContractRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.UpdateContractRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateContractRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateContractRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateContractRequest;

  static equals(a: UpdateContractRequest | PlainMessage<UpdateContractRequest> | undefined, b: UpdateContractRequest | PlainMessage<UpdateContractRequest> | undefined): boolean;
}

/**
 * @generated from message porter.v1.UpdateContractResponse
 */
export declare class UpdateContractResponse extends Message<UpdateContractResponse> {
  /**
   * @generated from field: porter.v1.ContractRevision contract_revision = 1;
   */
  contractRevision?: ContractRevision;

  constructor(data?: PartialMessage<UpdateContractResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.UpdateContractResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateContractResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateContractResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateContractResponse;

  static equals(a: UpdateContractResponse | PlainMessage<UpdateContractResponse> | undefined, b: UpdateContractResponse | PlainMessage<UpdateContractResponse> | undefined): boolean;
}

/**
 * @generated from message porter.v1.ClusterStatusRequest
 */
export declare class ClusterStatusRequest extends Message<ClusterStatusRequest> {
  /**
   * @generated from field: int64 project_id = 1;
   */
  projectId: bigint;

  /**
   * @generated from field: int64 cluster_id = 2;
   */
  clusterId: bigint;

  constructor(data?: PartialMessage<ClusterStatusRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.ClusterStatusRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ClusterStatusRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ClusterStatusRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ClusterStatusRequest;

  static equals(a: ClusterStatusRequest | PlainMessage<ClusterStatusRequest> | undefined, b: ClusterStatusRequest | PlainMessage<ClusterStatusRequest> | undefined): boolean;
}

/**
 * @generated from message porter.v1.ClusterStatusResponse
 */
export declare class ClusterStatusResponse extends Message<ClusterStatusResponse> {
  /**
   * @generated from field: int64 project_id = 1;
   */
  projectId: bigint;

  /**
   * @generated from field: int64 cluster_id = 2;
   */
  clusterId: bigint;

  /**
   * @generated from field: string phase = 3;
   */
  phase: string;

  /**
   * @generated from field: bool infrastructure_status = 4;
   */
  infrastructureStatus: boolean;

  /**
   * @generated from field: bool control_plane_status = 5;
   */
  controlPlaneStatus: boolean;

  constructor(data?: PartialMessage<ClusterStatusResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.ClusterStatusResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ClusterStatusResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ClusterStatusResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ClusterStatusResponse;

  static equals(a: ClusterStatusResponse | PlainMessage<ClusterStatusResponse> | undefined, b: ClusterStatusResponse | PlainMessage<ClusterStatusResponse> | undefined): boolean;
}

/**
 * @generated from message porter.v1.DeleteClusterRequest
 */
export declare class DeleteClusterRequest extends Message<DeleteClusterRequest> {
  /**
   * @generated from field: porter.v1.ContractRevision contract_revision = 1;
   */
  contractRevision?: ContractRevision;

  constructor(data?: PartialMessage<DeleteClusterRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.DeleteClusterRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteClusterRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteClusterRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteClusterRequest;

  static equals(a: DeleteClusterRequest | PlainMessage<DeleteClusterRequest> | undefined, b: DeleteClusterRequest | PlainMessage<DeleteClusterRequest> | undefined): boolean;
}

/**
 * @generated from message porter.v1.DeleteClusterResponse
 */
export declare class DeleteClusterResponse extends Message<DeleteClusterResponse> {
  /**
   * @generated from field: int64 project_id = 1;
   */
  projectId: bigint;

  /**
   * @generated from field: int64 cluster_id = 2;
   */
  clusterId: bigint;

  constructor(data?: PartialMessage<DeleteClusterResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.DeleteClusterResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteClusterResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteClusterResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteClusterResponse;

  static equals(a: DeleteClusterResponse | PlainMessage<DeleteClusterResponse> | undefined, b: DeleteClusterResponse | PlainMessage<DeleteClusterResponse> | undefined): boolean;
}

