// @generated by protoc-gen-es v1.4.0
// @generated from file porter/v1/cloud_provider_credentials.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { EnumCloudProvider } from "./cluster_pb.js";

/**
 * @generated from message porter.v1.AWSCredentials
 */
export declare class AWSCredentials extends Message<AWSCredentials> {
  /**
   * @generated from field: repeated porter.v1.AssumeRoleHop chain = 1;
   */
  chain: AssumeRoleHop[];

  constructor(data?: PartialMessage<AWSCredentials>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.AWSCredentials";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AWSCredentials;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AWSCredentials;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AWSCredentials;

  static equals(a: AWSCredentials | PlainMessage<AWSCredentials> | undefined, b: AWSCredentials | PlainMessage<AWSCredentials> | undefined): boolean;
}

/**
 * AssumeRoleHop contains all information required for the source_arn to assume the target_arn
 *
 * @generated from message porter.v1.AssumeRoleHop
 */
export declare class AssumeRoleHop extends Message<AssumeRoleHop> {
  /**
   * @generated from field: string source_arn = 1;
   */
  sourceArn: string;

  /**
   * @generated from field: string target_arn = 2;
   */
  targetArn: string;

  /**
   * external_id is optional and is used to prevent privilege escalation
   *
   * @generated from field: string external_id = 3;
   */
  externalId: string;

  constructor(data?: PartialMessage<AssumeRoleHop>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.AssumeRoleHop";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AssumeRoleHop;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AssumeRoleHop;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AssumeRoleHop;

  static equals(a: AssumeRoleHop | PlainMessage<AssumeRoleHop> | undefined, b: AssumeRoleHop | PlainMessage<AssumeRoleHop> | undefined): boolean;
}

/**
 * @generated from message porter.v1.AzureCredentials
 */
export declare class AzureCredentials extends Message<AzureCredentials> {
  /**
   * @generated from field: string client_id = 1;
   */
  clientId: string;

  /**
   * @generated from field: string subscription_id = 2;
   */
  subscriptionId: string;

  /**
   * @generated from field: string tenant_id = 3;
   */
  tenantId: string;

  /**
   * @generated from field: bytes service_principal_secret = 4;
   */
  servicePrincipalSecret: Uint8Array;

  constructor(data?: PartialMessage<AzureCredentials>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.AzureCredentials";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AzureCredentials;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AzureCredentials;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AzureCredentials;

  static equals(a: AzureCredentials | PlainMessage<AzureCredentials> | undefined, b: AzureCredentials | PlainMessage<AzureCredentials> | undefined): boolean;
}

/**
 * @generated from message porter.v1.GCPCredentials
 */
export declare class GCPCredentials extends Message<GCPCredentials> {
  /**
   * service_account_json_base64 is the base64 encoded service account json which can be used to authenticate with GCP
   *
   * @generated from field: string service_account_json_base64 = 1;
   */
  serviceAccountJsonBase64: string;

  constructor(data?: PartialMessage<GCPCredentials>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.GCPCredentials";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GCPCredentials;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GCPCredentials;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GCPCredentials;

  static equals(a: GCPCredentials | PlainMessage<GCPCredentials> | undefined, b: GCPCredentials | PlainMessage<GCPCredentials> | undefined): boolean;
}

/**
 * @generated from message porter.v1.UpdateCloudProviderCredentialsRequest
 */
export declare class UpdateCloudProviderCredentialsRequest extends Message<UpdateCloudProviderCredentialsRequest> {
  /**
   * project_id [REQUIRED] is the project that we are updating the credentials for
   *
   * @generated from field: int64 project_id = 1;
   */
  projectId: bigint;

  /**
   * cloud_provider [REQUIRED] represents the provider that we will provisioning the cluster in
   *
   * @generated from field: porter.v1.EnumCloudProvider cloud_provider = 2;
   */
  cloudProvider: EnumCloudProvider;

  /**
   * cloud_provider_credentials are the credentials for the specified cloud provider
   *
   * @generated from oneof porter.v1.UpdateCloudProviderCredentialsRequest.cloud_provider_credentials
   */
  cloudProviderCredentials: {
    /**
     * @generated from field: porter.v1.AWSCredentials aws_credentials = 3;
     */
    value: AWSCredentials;
    case: "awsCredentials";
  } | {
    /**
     * @generated from field: porter.v1.AzureCredentials azure_credentials = 4;
     */
    value: AzureCredentials;
    case: "azureCredentials";
  } | {
    /**
     * @generated from field: porter.v1.GCPCredentials gcp_credentials = 5;
     */
    value: GCPCredentials;
    case: "gcpCredentials";
  } | { case: undefined; value?: undefined };

  /**
   * credentials_identifier is the identifier for the credentials that we are updating. This is required during update operations, otherwise a new credential will be created
   *
   * @generated from field: string credentials_identifier = 6;
   */
  credentialsIdentifier: string;

  constructor(data?: PartialMessage<UpdateCloudProviderCredentialsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.UpdateCloudProviderCredentialsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateCloudProviderCredentialsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateCloudProviderCredentialsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateCloudProviderCredentialsRequest;

  static equals(a: UpdateCloudProviderCredentialsRequest | PlainMessage<UpdateCloudProviderCredentialsRequest> | undefined, b: UpdateCloudProviderCredentialsRequest | PlainMessage<UpdateCloudProviderCredentialsRequest> | undefined): boolean;
}

/**
 * @generated from message porter.v1.UpdateCloudProviderCredentialsResponse
 */
export declare class UpdateCloudProviderCredentialsResponse extends Message<UpdateCloudProviderCredentialsResponse> {
  /**
   * @generated from field: int64 project_id = 1;
   */
  projectId: bigint;

  /**
   * @generated from field: string credentials_identifier = 2;
   */
  credentialsIdentifier: string;

  constructor(data?: PartialMessage<UpdateCloudProviderCredentialsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.UpdateCloudProviderCredentialsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateCloudProviderCredentialsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateCloudProviderCredentialsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateCloudProviderCredentialsResponse;

  static equals(a: UpdateCloudProviderCredentialsResponse | PlainMessage<UpdateCloudProviderCredentialsResponse> | undefined, b: UpdateCloudProviderCredentialsResponse | PlainMessage<UpdateCloudProviderCredentialsResponse> | undefined): boolean;
}

