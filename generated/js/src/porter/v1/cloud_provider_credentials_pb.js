// @generated by protoc-gen-es v1.3.0
// @generated from file porter/v1/cloud_provider_credentials.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { EnumCloudProvider } from "./cluster_pb.js";

/**
 * @generated from message porter.v1.AWSCredentials
 */
export const AWSCredentials = proto3.makeMessageType(
  "porter.v1.AWSCredentials",
  () => [
    { no: 1, name: "chain", kind: "message", T: AssumeRoleHop, repeated: true },
  ],
);

/**
 * AssumeRoleHop contains all information required for the source_arn to assume the target_arn
 *
 * @generated from message porter.v1.AssumeRoleHop
 */
export const AssumeRoleHop = proto3.makeMessageType(
  "porter.v1.AssumeRoleHop",
  () => [
    { no: 1, name: "source_arn", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "target_arn", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "external_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.AzureCredentials
 */
export const AzureCredentials = proto3.makeMessageType(
  "porter.v1.AzureCredentials",
  () => [
    { no: 1, name: "client_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "subscription_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "tenant_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "service_principal_secret", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ],
);

/**
 * @generated from message porter.v1.GCPCredentials
 */
export const GCPCredentials = proto3.makeMessageType(
  "porter.v1.GCPCredentials",
  () => [
    { no: 1, name: "service_account_json_base64", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.UpdateCloudProviderCredentialsRequest
 */
export const UpdateCloudProviderCredentialsRequest = proto3.makeMessageType(
  "porter.v1.UpdateCloudProviderCredentialsRequest",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "cloud_provider", kind: "enum", T: proto3.getEnumType(EnumCloudProvider) },
    { no: 3, name: "aws_credentials", kind: "message", T: AWSCredentials, oneof: "cloud_provider_credentials" },
    { no: 4, name: "azure_credentials", kind: "message", T: AzureCredentials, oneof: "cloud_provider_credentials" },
    { no: 5, name: "gcp_credentials", kind: "message", T: GCPCredentials, oneof: "cloud_provider_credentials" },
    { no: 6, name: "credentials_identifier", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.UpdateCloudProviderCredentialsResponse
 */
export const UpdateCloudProviderCredentialsResponse = proto3.makeMessageType(
  "porter.v1.UpdateCloudProviderCredentialsResponse",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "credentials_identifier", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

