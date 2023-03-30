// @generated by protoc-gen-es v1.1.0
// @generated from file porter/v1/cluster_control_plane.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3, Timestamp } from "@bufbuild/protobuf";
import { AssumeRoleChainLink } from "./aws_assume_role_pb.js";
import { Contract, ContractRevision } from "./contract_pb.js";

/**
 * @generated from message porter.v1.RolePreflightCheckRequest
 */
export const RolePreflightCheckRequest = proto3.makeMessageType(
  "porter.v1.RolePreflightCheckRequest",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "target_arn", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "external_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.RolePreflightCheckResponse
 */
export const RolePreflightCheckResponse = proto3.makeMessageType(
  "porter.v1.RolePreflightCheckResponse",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "target_arn", kind: "scalar", T: 9 /* ScalarType.STRING */ },
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
 */
export const CreateAssumeRoleChainRequest = proto3.makeMessageType(
  "porter.v1.CreateAssumeRoleChainRequest",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "source_arn", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "target_access_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "target_secret_key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "target_session_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.CreateAssumeRoleChainResponse
 */
export const CreateAssumeRoleChainResponse = proto3.makeMessageType(
  "porter.v1.CreateAssumeRoleChainResponse",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "target_arn", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.EKSBearerTokenRequest
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
 */
export const EKSBearerTokenResponse = proto3.makeMessageType(
  "porter.v1.EKSBearerTokenResponse",
  () => [
    { no: 1, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.AssumeRoleChainTargetsRequest
 */
export const AssumeRoleChainTargetsRequest = proto3.makeMessageType(
  "porter.v1.AssumeRoleChainTargetsRequest",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.AssumeRoleChainTargetsResponse
 */
export const AssumeRoleChainTargetsResponse = proto3.makeMessageType(
  "porter.v1.AssumeRoleChainTargetsResponse",
  () => [
    { no: 1, name: "chain_links", kind: "message", T: AssumeRoleChainLink, repeated: true },
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
 * @generated from message porter.v1.ECRTokenForRegistryRequest
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
 */
export const AssumeRoleCredentialsResponse = proto3.makeMessageType(
  "porter.v1.AssumeRoleCredentialsResponse",
  () => [
    { no: 1, name: "aws_access_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "aws_secret_key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "aws_session_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

