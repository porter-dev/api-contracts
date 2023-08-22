// @generated by protoc-gen-es v1.3.0
// @generated from file porter/v1/cluster_control_plane.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3, Timestamp } from "@bufbuild/protobuf";
import { Contract, ContractRevision } from "./contract_pb.js";
import { PorterApp } from "./porter_app_pb.js";
import { AssumeRoleChainLink } from "./aws_assume_role_pb.js";

/**
 * @generated from enum porter.v1.EnumCLIAction
 */
export const EnumCLIAction = proto3.makeEnum(
  "porter.v1.EnumCLIAction",
  [
    {no: 0, name: "ENUM_CLI_ACTION_UNSPECIFIED"},
    {no: 1, name: "ENUM_CLI_ACTION_NONE"},
    {no: 2, name: "ENUM_CLI_ACTION_BUILD"},
    {no: 3, name: "ENUM_CLI_ACTION_TRACK_PREDEPLOY"},
  ],
);

/**
 * @generated from message porter.v1.QuotaPreflightCheckRequest
 */
export const QuotaPreflightCheckRequest = proto3.makeMessageType(
  "porter.v1.QuotaPreflightCheckRequest",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "target_arn", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "external_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "region", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.QuotaPreflightCheckResponse
 */
export const QuotaPreflightCheckResponse = proto3.makeMessageType(
  "porter.v1.QuotaPreflightCheckResponse",
  [],
);

/**
 * @generated from message porter.v1.CreateAssumeRoleChainRequest
 * @deprecated
 */
export const CreateAssumeRoleChainRequest = proto3.makeMessageType(
  "porter.v1.CreateAssumeRoleChainRequest",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "source_arn", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "target_access_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "target_secret_key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "target_session_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "target_arn", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "external_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.CreateAssumeRoleChainResponse
 * @deprecated
 */
export const CreateAssumeRoleChainResponse = proto3.makeMessageType(
  "porter.v1.CreateAssumeRoleChainResponse",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "target_arn", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.SaveAzureCredentialsRequest
 * @deprecated
 */
export const SaveAzureCredentialsRequest = proto3.makeMessageType(
  "porter.v1.SaveAzureCredentialsRequest",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "client_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "subscription_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "tenant_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "service_principal_secret", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ],
);

/**
 * @generated from message porter.v1.SaveAzureCredentialsResponse
 * @deprecated
 */
export const SaveAzureCredentialsResponse = proto3.makeMessageType(
  "porter.v1.SaveAzureCredentialsResponse",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "credentials_identifier", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.KubeConfigForClusterRequest
 */
export const KubeConfigForClusterRequest = proto3.makeMessageType(
  "porter.v1.KubeConfigForClusterRequest",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "cluster_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * @generated from message porter.v1.KubeConfigForClusterResponse
 */
export const KubeConfigForClusterResponse = proto3.makeMessageType(
  "porter.v1.KubeConfigForClusterResponse",
  () => [
    { no: 1, name: "kube_config", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.UpdateContractRequest
 */
export const UpdateContractRequest = proto3.makeMessageType(
  "porter.v1.UpdateContractRequest",
  () => [
    { no: 1, name: "contract", kind: "message", T: Contract },
  ],
);

/**
 * @generated from message porter.v1.UpdateContractResponse
 */
export const UpdateContractResponse = proto3.makeMessageType(
  "porter.v1.UpdateContractResponse",
  () => [
    { no: 1, name: "contract_revision", kind: "message", T: ContractRevision },
  ],
);

/**
 * @generated from message porter.v1.ClusterStatusRequest
 */
export const ClusterStatusRequest = proto3.makeMessageType(
  "porter.v1.ClusterStatusRequest",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "cluster_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * @generated from message porter.v1.ClusterStatusResponse
 */
export const ClusterStatusResponse = proto3.makeMessageType(
  "porter.v1.ClusterStatusResponse",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "cluster_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "phase", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "infrastructure_status", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 5, name: "control_plane_status", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * @generated from message porter.v1.DeleteClusterRequest
 */
export const DeleteClusterRequest = proto3.makeMessageType(
  "porter.v1.DeleteClusterRequest",
  () => [
    { no: 1, name: "contract_revision", kind: "message", T: ContractRevision },
  ],
);

/**
 * @generated from message porter.v1.DeleteClusterResponse
 */
export const DeleteClusterResponse = proto3.makeMessageType(
  "porter.v1.DeleteClusterResponse",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "cluster_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * @generated from message porter.v1.ListRepositoriesForRegistryRequest
 */
export const ListRepositoriesForRegistryRequest = proto3.makeMessageType(
  "porter.v1.ListRepositoriesForRegistryRequest",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "registry_uri", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.ListRepositoriesForRegistryResponse
 */
export const ListRepositoriesForRegistryResponse = proto3.makeMessageType(
  "porter.v1.ListRepositoriesForRegistryResponse",
  () => [
    { no: 1, name: "repositories", kind: "message", T: Repository, repeated: true },
  ],
);

/**
 * @generated from message porter.v1.Repository
 */
export const Repository = proto3.makeMessageType(
  "porter.v1.Repository",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "uri", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.ListImagesForRepositoryRequest
 */
export const ListImagesForRepositoryRequest = proto3.makeMessageType(
  "porter.v1.ListImagesForRepositoryRequest",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "registry_uri", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "repo_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.ListImagesForRepositoryResponse
 */
export const ListImagesForRepositoryResponse = proto3.makeMessageType(
  "porter.v1.ListImagesForRepositoryResponse",
  () => [
    { no: 1, name: "images", kind: "message", T: Image, repeated: true },
  ],
);

/**
 * @generated from message porter.v1.Image
 */
export const Image = proto3.makeMessageType(
  "porter.v1.Image",
  () => [
    { no: 1, name: "repository_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "tag", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.TokenForRegistryRequest
 */
export const TokenForRegistryRequest = proto3.makeMessageType(
  "porter.v1.TokenForRegistryRequest",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "registry_uri", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.TokenForRegistryResponse
 */
export const TokenForRegistryResponse = proto3.makeMessageType(
  "porter.v1.TokenForRegistryResponse",
  () => [
    { no: 1, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "expiry", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message porter.v1.ReadContractRequest
 */
export const ReadContractRequest = proto3.makeMessageType(
  "porter.v1.ReadContractRequest",
  () => [
    { no: 1, name: "contract_revision", kind: "message", T: ContractRevision },
  ],
);

/**
 * @generated from message porter.v1.ReadContractResponse
 */
export const ReadContractResponse = proto3.makeMessageType(
  "porter.v1.ReadContractResponse",
  () => [
    { no: 1, name: "base64_contract", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.ValidatePorterAppRequest
 */
export const ValidatePorterAppRequest = proto3.makeMessageType(
  "porter.v1.ValidatePorterAppRequest",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "deployment_target_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "commit_sha", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "app", kind: "message", T: PorterApp },
  ],
);

/**
 * @generated from message porter.v1.ValidatePorterAppResponse
 */
export const ValidatePorterAppResponse = proto3.makeMessageType(
  "porter.v1.ValidatePorterAppResponse",
  () => [
    { no: 1, name: "app", kind: "message", T: PorterApp },
  ],
);

/**
 * @generated from message porter.v1.ApplyPorterAppRequest
 */
export const ApplyPorterAppRequest = proto3.makeMessageType(
  "porter.v1.ApplyPorterAppRequest",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "deployment_target_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "app", kind: "message", T: PorterApp },
    { no: 4, name: "porter_app_revision_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.ApplyPorterAppResponse
 */
export const ApplyPorterAppResponse = proto3.makeMessageType(
  "porter.v1.ApplyPorterAppResponse",
  () => [
    { no: 1, name: "porter_app_revision_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "cli_action", kind: "enum", T: proto3.getEnumType(EnumCLIAction) },
  ],
);

/**
 * @generated from message porter.v1.CurrentAppRevisionRequest
 */
export const CurrentAppRevisionRequest = proto3.makeMessageType(
  "porter.v1.CurrentAppRevisionRequest",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "porter_app_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "deployment_target_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.CurrentAppRevisionResponse
 */
export const CurrentAppRevisionResponse = proto3.makeMessageType(
  "porter.v1.CurrentAppRevisionResponse",
  () => [
    { no: 1, name: "app", kind: "message", T: PorterApp },
  ],
);

/**
 * @generated from message porter.v1.EKSBearerTokenRequest
 * @deprecated
 */
export const EKSBearerTokenRequest = proto3.makeMessageType(
  "porter.v1.EKSBearerTokenRequest",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "cluster_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * @generated from message porter.v1.EKSBearerTokenResponse
 * @deprecated
 */
export const EKSBearerTokenResponse = proto3.makeMessageType(
  "porter.v1.EKSBearerTokenResponse",
  () => [
    { no: 1, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.CertificateAuthorityDataRequest
 * @deprecated
 */
export const CertificateAuthorityDataRequest = proto3.makeMessageType(
  "porter.v1.CertificateAuthorityDataRequest",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "cluster_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * @generated from message porter.v1.CertificateAuthorityDataResponse
 * @deprecated
 */
export const CertificateAuthorityDataResponse = proto3.makeMessageType(
  "porter.v1.CertificateAuthorityDataResponse",
  () => [
    { no: 1, name: "certificate_authority_data", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.AssumeRoleChainTargetsRequest
 * @deprecated
 */
export const AssumeRoleChainTargetsRequest = proto3.makeMessageType(
  "porter.v1.AssumeRoleChainTargetsRequest",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.AssumeRoleChainTargetsResponse
 * @deprecated
 */
export const AssumeRoleChainTargetsResponse = proto3.makeMessageType(
  "porter.v1.AssumeRoleChainTargetsResponse",
  () => [
    { no: 1, name: "chain_links", kind: "message", T: AssumeRoleChainLink, repeated: true },
  ],
);

/**
 * @generated from message porter.v1.ECRTokenForRegistryRequest
 * @deprecated
 */
export const ECRTokenForRegistryRequest = proto3.makeMessageType(
  "porter.v1.ECRTokenForRegistryRequest",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "region", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "aws_account_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.ECRTokenForRegistryResponse
 * @deprecated
 */
export const ECRTokenForRegistryResponse = proto3.makeMessageType(
  "porter.v1.ECRTokenForRegistryResponse",
  () => [
    { no: 1, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "expiry", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message porter.v1.AssumeRoleCredentialsRequest
 * @deprecated
 */
export const AssumeRoleCredentialsRequest = proto3.makeMessageType(
  "porter.v1.AssumeRoleCredentialsRequest",
  () => [
    { no: 1, name: "aws_account_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * @generated from message porter.v1.AssumeRoleCredentialsResponse
 * @deprecated
 */
export const AssumeRoleCredentialsResponse = proto3.makeMessageType(
  "porter.v1.AssumeRoleCredentialsResponse",
  () => [
    { no: 1, name: "aws_access_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "aws_secret_key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "aws_session_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.DockerConfigFileForRegistryRequest
 * @deprecated
 */
export const DockerConfigFileForRegistryRequest = proto3.makeMessageType(
  "porter.v1.DockerConfigFileForRegistryRequest",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "registry_uri", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.DockerConfigFileForRegistryResponse
 * @deprecated
 */
export const DockerConfigFileForRegistryResponse = proto3.makeMessageType(
  "porter.v1.DockerConfigFileForRegistryResponse",
  () => [
    { no: 1, name: "docker_config_file", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ],
);

