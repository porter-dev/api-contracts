// @generated by protoc-gen-es v1.1.0
// @generated from file porter/v1/cluster_control_plane.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { AssumeRoleChainLink } from "./aws_assume_role_pb.js";
import type { Contract, ContractRevision } from "./contract_pb.js";

/**
 * @generated from message porter.v1.QuotaPreflightCheckRequest
 */
export declare class QuotaPreflightCheckRequest extends Message<QuotaPreflightCheckRequest> {
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

  /**
   * @generated from field: string region = 4;
   */
  region: string;

  constructor(data?: PartialMessage<QuotaPreflightCheckRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.QuotaPreflightCheckRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QuotaPreflightCheckRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QuotaPreflightCheckRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QuotaPreflightCheckRequest;

  static equals(a: QuotaPreflightCheckRequest | PlainMessage<QuotaPreflightCheckRequest> | undefined, b: QuotaPreflightCheckRequest | PlainMessage<QuotaPreflightCheckRequest> | undefined): boolean;
}

/**
 * @generated from message porter.v1.QuotaPreflightCheckResponse
 */
export declare class QuotaPreflightCheckResponse extends Message<QuotaPreflightCheckResponse> {
  constructor(data?: PartialMessage<QuotaPreflightCheckResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.QuotaPreflightCheckResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QuotaPreflightCheckResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QuotaPreflightCheckResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QuotaPreflightCheckResponse;

  static equals(a: QuotaPreflightCheckResponse | PlainMessage<QuotaPreflightCheckResponse> | undefined, b: QuotaPreflightCheckResponse | PlainMessage<QuotaPreflightCheckResponse> | undefined): boolean;
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

  /**
   * @generated from field: string target_arn = 6;
   */
  targetArn: string;

  /**
   * @generated from field: string external_id = 7;
   */
  externalId: string;

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
 * @generated from message porter.v1.SaveAzureCredentialsRequest
 */
export declare class SaveAzureCredentialsRequest extends Message<SaveAzureCredentialsRequest> {
  /**
   * @generated from field: int64 project_id = 1;
   */
  projectId: bigint;

  /**
   * @generated from field: string client_id = 2;
   */
  clientId: string;

  /**
   * @generated from field: string subscription_id = 3;
   */
  subscriptionId: string;

  /**
   * @generated from field: string tenant_id = 4;
   */
  tenantId: string;

  /**
   * @generated from field: bytes service_principal_secret = 5;
   */
  servicePrincipalSecret: Uint8Array;

  constructor(data?: PartialMessage<SaveAzureCredentialsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.SaveAzureCredentialsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SaveAzureCredentialsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SaveAzureCredentialsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SaveAzureCredentialsRequest;

  static equals(a: SaveAzureCredentialsRequest | PlainMessage<SaveAzureCredentialsRequest> | undefined, b: SaveAzureCredentialsRequest | PlainMessage<SaveAzureCredentialsRequest> | undefined): boolean;
}

/**
 * @generated from message porter.v1.SaveAzureCredentialsResponse
 */
export declare class SaveAzureCredentialsResponse extends Message<SaveAzureCredentialsResponse> {
  /**
   * @generated from field: int64 project_id = 1;
   */
  projectId: bigint;

  /**
   * @generated from field: string credentials_identifier = 2;
   */
  credentialsIdentifier: string;

  constructor(data?: PartialMessage<SaveAzureCredentialsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.SaveAzureCredentialsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SaveAzureCredentialsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SaveAzureCredentialsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SaveAzureCredentialsResponse;

  static equals(a: SaveAzureCredentialsResponse | PlainMessage<SaveAzureCredentialsResponse> | undefined, b: SaveAzureCredentialsResponse | PlainMessage<SaveAzureCredentialsResponse> | undefined): boolean;
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
 * @generated from message porter.v1.CertificateAuthorityDataRequest
 */
export declare class CertificateAuthorityDataRequest extends Message<CertificateAuthorityDataRequest> {
  /**
   * @generated from field: int64 project_id = 1;
   */
  projectId: bigint;

  /**
   * @generated from field: int64 cluster_id = 2;
   */
  clusterId: bigint;

  constructor(data?: PartialMessage<CertificateAuthorityDataRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.CertificateAuthorityDataRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CertificateAuthorityDataRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CertificateAuthorityDataRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CertificateAuthorityDataRequest;

  static equals(a: CertificateAuthorityDataRequest | PlainMessage<CertificateAuthorityDataRequest> | undefined, b: CertificateAuthorityDataRequest | PlainMessage<CertificateAuthorityDataRequest> | undefined): boolean;
}

/**
 * @generated from message porter.v1.CertificateAuthorityDataResponse
 */
export declare class CertificateAuthorityDataResponse extends Message<CertificateAuthorityDataResponse> {
  /**
   * @generated from field: string certificate_authority_data = 1;
   */
  certificateAuthorityData: string;

  constructor(data?: PartialMessage<CertificateAuthorityDataResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.CertificateAuthorityDataResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CertificateAuthorityDataResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CertificateAuthorityDataResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CertificateAuthorityDataResponse;

  static equals(a: CertificateAuthorityDataResponse | PlainMessage<CertificateAuthorityDataResponse> | undefined, b: CertificateAuthorityDataResponse | PlainMessage<CertificateAuthorityDataResponse> | undefined): boolean;
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

/**
 * @generated from message porter.v1.ECRTokenForRegistryRequest
 */
export declare class ECRTokenForRegistryRequest extends Message<ECRTokenForRegistryRequest> {
  /**
   * @generated from field: int64 project_id = 1;
   */
  projectId: bigint;

  /**
   * @generated from field: string region = 2;
   */
  region: string;

  /**
   * @generated from field: string aws_account_id = 3;
   */
  awsAccountId: string;

  constructor(data?: PartialMessage<ECRTokenForRegistryRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.ECRTokenForRegistryRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ECRTokenForRegistryRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ECRTokenForRegistryRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ECRTokenForRegistryRequest;

  static equals(a: ECRTokenForRegistryRequest | PlainMessage<ECRTokenForRegistryRequest> | undefined, b: ECRTokenForRegistryRequest | PlainMessage<ECRTokenForRegistryRequest> | undefined): boolean;
}

/**
 * @generated from message porter.v1.ECRTokenForRegistryResponse
 */
export declare class ECRTokenForRegistryResponse extends Message<ECRTokenForRegistryResponse> {
  /**
   * @generated from field: string token = 1;
   */
  token: string;

  /**
   * @generated from field: google.protobuf.Timestamp expiry = 2;
   */
  expiry?: Timestamp;

  constructor(data?: PartialMessage<ECRTokenForRegistryResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.ECRTokenForRegistryResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ECRTokenForRegistryResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ECRTokenForRegistryResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ECRTokenForRegistryResponse;

  static equals(a: ECRTokenForRegistryResponse | PlainMessage<ECRTokenForRegistryResponse> | undefined, b: ECRTokenForRegistryResponse | PlainMessage<ECRTokenForRegistryResponse> | undefined): boolean;
}

/**
 * @generated from message porter.v1.AssumeRoleCredentialsRequest
 */
export declare class AssumeRoleCredentialsRequest extends Message<AssumeRoleCredentialsRequest> {
  /**
   * @generated from field: string aws_account_id = 1;
   */
  awsAccountId: string;

  /**
   * @generated from field: int64 project_id = 2;
   */
  projectId: bigint;

  constructor(data?: PartialMessage<AssumeRoleCredentialsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.AssumeRoleCredentialsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AssumeRoleCredentialsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AssumeRoleCredentialsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AssumeRoleCredentialsRequest;

  static equals(a: AssumeRoleCredentialsRequest | PlainMessage<AssumeRoleCredentialsRequest> | undefined, b: AssumeRoleCredentialsRequest | PlainMessage<AssumeRoleCredentialsRequest> | undefined): boolean;
}

/**
 * @generated from message porter.v1.AssumeRoleCredentialsResponse
 */
export declare class AssumeRoleCredentialsResponse extends Message<AssumeRoleCredentialsResponse> {
  /**
   * @generated from field: string aws_access_id = 1;
   */
  awsAccessId: string;

  /**
   * @generated from field: string aws_secret_key = 2;
   */
  awsSecretKey: string;

  /**
   * @generated from field: string aws_session_token = 3;
   */
  awsSessionToken: string;

  constructor(data?: PartialMessage<AssumeRoleCredentialsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.AssumeRoleCredentialsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AssumeRoleCredentialsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AssumeRoleCredentialsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AssumeRoleCredentialsResponse;

  static equals(a: AssumeRoleCredentialsResponse | PlainMessage<AssumeRoleCredentialsResponse> | undefined, b: AssumeRoleCredentialsResponse | PlainMessage<AssumeRoleCredentialsResponse> | undefined): boolean;
}

/**
 * @generated from message porter.v1.ListRepositoriesForRegistryRequest
 */
export declare class ListRepositoriesForRegistryRequest extends Message<ListRepositoriesForRegistryRequest> {
  /**
   * @generated from field: int64 project_id = 1;
   */
  projectId: bigint;

  /**
   * @generated from field: string registry_uri = 2;
   */
  registryUri: string;

  constructor(data?: PartialMessage<ListRepositoriesForRegistryRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.ListRepositoriesForRegistryRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListRepositoriesForRegistryRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListRepositoriesForRegistryRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListRepositoriesForRegistryRequest;

  static equals(a: ListRepositoriesForRegistryRequest | PlainMessage<ListRepositoriesForRegistryRequest> | undefined, b: ListRepositoriesForRegistryRequest | PlainMessage<ListRepositoriesForRegistryRequest> | undefined): boolean;
}

/**
 * @generated from message porter.v1.ListRepositoriesForRegistryResponse
 */
export declare class ListRepositoriesForRegistryResponse extends Message<ListRepositoriesForRegistryResponse> {
  /**
   * @generated from field: repeated porter.v1.Repository repositories = 1;
   */
  repositories: Repository[];

  constructor(data?: PartialMessage<ListRepositoriesForRegistryResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.ListRepositoriesForRegistryResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListRepositoriesForRegistryResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListRepositoriesForRegistryResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListRepositoriesForRegistryResponse;

  static equals(a: ListRepositoriesForRegistryResponse | PlainMessage<ListRepositoriesForRegistryResponse> | undefined, b: ListRepositoriesForRegistryResponse | PlainMessage<ListRepositoriesForRegistryResponse> | undefined): boolean;
}

/**
 * @generated from message porter.v1.Repository
 */
export declare class Repository extends Message<Repository> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: string uri = 3;
   */
  uri: string;

  constructor(data?: PartialMessage<Repository>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Repository";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Repository;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Repository;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Repository;

  static equals(a: Repository | PlainMessage<Repository> | undefined, b: Repository | PlainMessage<Repository> | undefined): boolean;
}

/**
 * @generated from message porter.v1.ListImagesForRepositoryRequest
 */
export declare class ListImagesForRepositoryRequest extends Message<ListImagesForRepositoryRequest> {
  /**
   * @generated from field: int64 project_id = 1;
   */
  projectId: bigint;

  /**
   * @generated from field: string registry_uri = 2;
   */
  registryUri: string;

  /**
   * @generated from field: string repo_name = 3;
   */
  repoName: string;

  constructor(data?: PartialMessage<ListImagesForRepositoryRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.ListImagesForRepositoryRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListImagesForRepositoryRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListImagesForRepositoryRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListImagesForRepositoryRequest;

  static equals(a: ListImagesForRepositoryRequest | PlainMessage<ListImagesForRepositoryRequest> | undefined, b: ListImagesForRepositoryRequest | PlainMessage<ListImagesForRepositoryRequest> | undefined): boolean;
}

/**
 * @generated from message porter.v1.ListImagesForRepositoryResponse
 */
export declare class ListImagesForRepositoryResponse extends Message<ListImagesForRepositoryResponse> {
  /**
   * @generated from field: repeated porter.v1.Image images = 1;
   */
  images: Image[];

  constructor(data?: PartialMessage<ListImagesForRepositoryResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.ListImagesForRepositoryResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListImagesForRepositoryResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListImagesForRepositoryResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListImagesForRepositoryResponse;

  static equals(a: ListImagesForRepositoryResponse | PlainMessage<ListImagesForRepositoryResponse> | undefined, b: ListImagesForRepositoryResponse | PlainMessage<ListImagesForRepositoryResponse> | undefined): boolean;
}

/**
 * @generated from message porter.v1.Image
 */
export declare class Image extends Message<Image> {
  /**
   * @generated from field: string repository_name = 1;
   */
  repositoryName: string;

  /**
   * @generated from field: string tag = 2;
   */
  tag: string;

  constructor(data?: PartialMessage<Image>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Image";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Image;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Image;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Image;

  static equals(a: Image | PlainMessage<Image> | undefined, b: Image | PlainMessage<Image> | undefined): boolean;
}

/**
 * @generated from message porter.v1.TokenForRegistryRequest
 */
export declare class TokenForRegistryRequest extends Message<TokenForRegistryRequest> {
  /**
   * @generated from field: int64 project_id = 1;
   */
  projectId: bigint;

  /**
   * @generated from field: string registry_uri = 2;
   */
  registryUri: string;

  constructor(data?: PartialMessage<TokenForRegistryRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.TokenForRegistryRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TokenForRegistryRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TokenForRegistryRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TokenForRegistryRequest;

  static equals(a: TokenForRegistryRequest | PlainMessage<TokenForRegistryRequest> | undefined, b: TokenForRegistryRequest | PlainMessage<TokenForRegistryRequest> | undefined): boolean;
}

/**
 * @generated from message porter.v1.TokenForRegistryResponse
 */
export declare class TokenForRegistryResponse extends Message<TokenForRegistryResponse> {
  /**
   * @generated from field: string token = 1;
   */
  token: string;

  /**
   * @generated from field: google.protobuf.Timestamp expiry = 2;
   */
  expiry?: Timestamp;

  constructor(data?: PartialMessage<TokenForRegistryResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.TokenForRegistryResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TokenForRegistryResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TokenForRegistryResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TokenForRegistryResponse;

  static equals(a: TokenForRegistryResponse | PlainMessage<TokenForRegistryResponse> | undefined, b: TokenForRegistryResponse | PlainMessage<TokenForRegistryResponse> | undefined): boolean;
}

/**
 * @generated from message porter.v1.DockerConfigFileForRegistryRequest
 */
export declare class DockerConfigFileForRegistryRequest extends Message<DockerConfigFileForRegistryRequest> {
  /**
   * @generated from field: int64 project_id = 1;
   */
  projectId: bigint;

  /**
   * @generated from field: string registry_uri = 2;
   */
  registryUri: string;

  constructor(data?: PartialMessage<DockerConfigFileForRegistryRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.DockerConfigFileForRegistryRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DockerConfigFileForRegistryRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DockerConfigFileForRegistryRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DockerConfigFileForRegistryRequest;

  static equals(a: DockerConfigFileForRegistryRequest | PlainMessage<DockerConfigFileForRegistryRequest> | undefined, b: DockerConfigFileForRegistryRequest | PlainMessage<DockerConfigFileForRegistryRequest> | undefined): boolean;
}

/**
 * @generated from message porter.v1.DockerConfigFileForRegistryResponse
 */
export declare class DockerConfigFileForRegistryResponse extends Message<DockerConfigFileForRegistryResponse> {
  /**
   * @generated from field: bytes docker_config_file = 1;
   */
  dockerConfigFile: Uint8Array;

  constructor(data?: PartialMessage<DockerConfigFileForRegistryResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.DockerConfigFileForRegistryResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DockerConfigFileForRegistryResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DockerConfigFileForRegistryResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DockerConfigFileForRegistryResponse;

  static equals(a: DockerConfigFileForRegistryResponse | PlainMessage<DockerConfigFileForRegistryResponse> | undefined, b: DockerConfigFileForRegistryResponse | PlainMessage<DockerConfigFileForRegistryResponse> | undefined): boolean;
}

/**
 * @generated from message porter.v1.ReadContractRequest
 */
export declare class ReadContractRequest extends Message<ReadContractRequest> {
  /**
   * @generated from field: porter.v1.ContractRevision contract_revision = 1;
   */
  contractRevision?: ContractRevision;

  constructor(data?: PartialMessage<ReadContractRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.ReadContractRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ReadContractRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ReadContractRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ReadContractRequest;

  static equals(a: ReadContractRequest | PlainMessage<ReadContractRequest> | undefined, b: ReadContractRequest | PlainMessage<ReadContractRequest> | undefined): boolean;
}

/**
 * @generated from message porter.v1.ReadContractResponse
 */
export declare class ReadContractResponse extends Message<ReadContractResponse> {
  /**
   * @generated from field: string base64_contract = 1;
   */
  base64Contract: string;

  constructor(data?: PartialMessage<ReadContractResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.ReadContractResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ReadContractResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ReadContractResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ReadContractResponse;

  static equals(a: ReadContractResponse | PlainMessage<ReadContractResponse> | undefined, b: ReadContractResponse | PlainMessage<ReadContractResponse> | undefined): boolean;
}

