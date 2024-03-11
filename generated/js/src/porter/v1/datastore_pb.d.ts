// @generated by protoc-gen-es v1.4.2
// @generated from file porter/v1/datastore.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { EnumCloudProvider } from "./cluster_pb.js";

/**
 * @generated from enum porter.v1.EnumDatastoreKind
 */
export declare enum EnumDatastoreKind {
  /**
   * @generated from enum value: ENUM_DATASTORE_KIND_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: ENUM_DATASTORE_KIND_AWS_RDS = 1;
   */
  AWS_RDS = 1,
}

/**
 * @generated from enum porter.v1.EnumAwsRdsEngine
 */
export declare enum EnumAwsRdsEngine {
  /**
   * @generated from enum value: ENUM_AWS_RDS_ENGINE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: ENUM_AWS_RDS_ENGINE_POSTGRESQL = 1;
   */
  POSTGRESQL = 1,

  /**
   * @generated from enum value: ENUM_AWS_RDS_ENGINE_AURORA_POSTGRESQL = 2;
   */
  AURORA_POSTGRESQL = 2,
}

/**
 * ManagedDatastore is the specification for a Porter-managed datastore
 *
 * @generated from message porter.v1.ManagedDatastore
 */
export declare class ManagedDatastore extends Message<ManagedDatastore> {
  /**
   * id represents the id of the datastore. This is required for update operations, but should be left blank when creating a datastore
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * cloud_provider represents the provider that the datastore is provisioned in
   *
   * @generated from field: porter.v1.EnumCloudProvider cloud_provider = 2;
   */
  cloudProvider: EnumCloudProvider;

  /**
   * cloud_provider_credential_identifier is the credential used to provision the datastore
   *
   * @generated from field: string cloud_provider_credential_identifier = 3;
   */
  cloudProviderCredentialIdentifier: string;

  /**
   * region is the region the datastore is provisioned in
   *
   * @generated from field: string region = 4;
   */
  region: string;

  /**
   * name is the name of the datastore
   *
   * @generated from field: string name = 5;
   */
  name: string;

  /**
   * kind represents the type of the datastore
   *
   * @generated from field: porter.v1.EnumDatastoreKind kind = 6;
   */
  kind: EnumDatastoreKind;

  /**
   * kind_values are the required values depending on kind
   *
   * @generated from oneof porter.v1.ManagedDatastore.kind_values
   */
  kindValues: {
    /**
     * @generated from field: porter.v1.AwsRds aws_rds_kind = 7;
     */
    value: AwsRds;
    case: "awsRdsKind";
  } | { case: undefined; value?: undefined };

  /**
   * peering describes the peering of this datastore
   *
   * @generated from field: porter.v1.Peering peering = 8;
   */
  peering?: Peering;

  constructor(data?: PartialMessage<ManagedDatastore>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.ManagedDatastore";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ManagedDatastore;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ManagedDatastore;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ManagedDatastore;

  static equals(a: ManagedDatastore | PlainMessage<ManagedDatastore> | undefined, b: ManagedDatastore | PlainMessage<ManagedDatastore> | undefined): boolean;
}

/**
 * @generated from message porter.v1.Peering
 */
export declare class Peering extends Message<Peering> {
  /**
   * peered_cluster_ids is a list of cluster ids that this datastore is peered to
   *
   * @generated from field: repeated int64 peered_cluster_ids = 1;
   */
  peeredClusterIds: bigint[];

  constructor(data?: PartialMessage<Peering>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Peering";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Peering;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Peering;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Peering;

  static equals(a: Peering | PlainMessage<Peering> | undefined, b: Peering | PlainMessage<Peering> | undefined): boolean;
}

/**
 * @generated from message porter.v1.DatastorePasswordSecretRef
 */
export declare class DatastorePasswordSecretRef extends Message<DatastorePasswordSecretRef> {
  /**
   * name is the name of the secret
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * namespace is the namespace of the secret
   *
   * @generated from field: string namespace = 2;
   */
  namespace: string;

  /**
   * key is the key of the password in the secret
   *
   * @generated from field: string key = 3;
   */
  key: string;

  constructor(data?: PartialMessage<DatastorePasswordSecretRef>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.DatastorePasswordSecretRef";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DatastorePasswordSecretRef;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DatastorePasswordSecretRef;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DatastorePasswordSecretRef;

  static equals(a: DatastorePasswordSecretRef | PlainMessage<DatastorePasswordSecretRef> | undefined, b: DatastorePasswordSecretRef | PlainMessage<DatastorePasswordSecretRef> | undefined): boolean;
}

/**
 * @generated from message porter.v1.AwsRds
 */
export declare class AwsRds extends Message<AwsRds> {
  /**
   * database_name is the name of the rds database
   *
   * @generated from field: optional string database_name = 1;
   */
  databaseName?: string;

  /**
   * master_username is the username of the database
   *
   * @generated from field: optional string master_username = 2;
   */
  masterUsername?: string;

  /**
   * master_user_password_literal is the string value of the password; this is only used for creating the database password secret and is wiped when the contract is saved
   *
   * @generated from field: optional string master_user_password_literal = 3;
   */
  masterUserPasswordLiteral?: string;

  /**
   * master_user_password_secret_ref is the reference to the secret that stores the password of the database
   *
   * @generated from field: porter.v1.DatastorePasswordSecretRef master_user_password_secret_ref = 4;
   */
  masterUserPasswordSecretRef?: DatastorePasswordSecretRef;

  /**
   * allocated_storage_gigabytes is the allocated storage of the database in gigabytes
   *
   * @generated from field: optional int64 allocated_storage_gigabytes = 5;
   */
  allocatedStorageGigabytes?: bigint;

  /**
   * instance_class is the instance class of the database
   *
   * @generated from field: optional string instance_class = 6;
   */
  instanceClass?: string;

  /**
   * engine is the engine of the database
   *
   * @generated from field: porter.v1.EnumAwsRdsEngine engine = 7;
   */
  engine: EnumAwsRdsEngine;

  /**
   * engine_version is the version of the engine
   *
   * @generated from field: optional string engine_version = 8;
   */
  engineVersion?: string;

  constructor(data?: PartialMessage<AwsRds>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.AwsRds";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AwsRds;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AwsRds;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AwsRds;

  static equals(a: AwsRds | PlainMessage<AwsRds> | undefined, b: AwsRds | PlainMessage<AwsRds> | undefined): boolean;
}

/**
 * UpdateDatastorePayload is used to send messages via nats to update a datastore
 *
 * @generated from message porter.v1.UpdateDatastorePayload
 */
export declare class UpdateDatastorePayload extends Message<UpdateDatastorePayload> {
  /**
   * project_id is the id of the project that the datastore belongs to
   *
   * @generated from field: int64 project_id = 1;
   */
  projectId: bigint;

  /**
   * datastore_id is the id of the datastore
   *
   * @generated from field: string datastore_id = 2;
   */
  datastoreId: string;

  constructor(data?: PartialMessage<UpdateDatastorePayload>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.UpdateDatastorePayload";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateDatastorePayload;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateDatastorePayload;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateDatastorePayload;

  static equals(a: UpdateDatastorePayload | PlainMessage<UpdateDatastorePayload> | undefined, b: UpdateDatastorePayload | PlainMessage<UpdateDatastorePayload> | undefined): boolean;
}

/**
 * DatastoreCredential is used to connect to a datastore
 *
 * @generated from message porter.v1.DatastoreCredential
 */
export declare class DatastoreCredential extends Message<DatastoreCredential> {
  /**
   * host is the datastore host
   *
   * @generated from field: string host = 1;
   */
  host: string;

  /**
   * database_name is the name of the database
   *
   * @generated from field: string database_name = 2;
   */
  databaseName: string;

  /**
   * username is the username required to access the datastore
   *
   * @generated from field: string username = 3;
   */
  username: string;

  /**
   * password is the password required to access the datastore
   *
   * @generated from field: string password = 4;
   */
  password: string;

  /**
   * port is the port to connect to
   *
   * @generated from field: int64 port = 5;
   */
  port: bigint;

  constructor(data?: PartialMessage<DatastoreCredential>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.DatastoreCredential";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DatastoreCredential;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DatastoreCredential;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DatastoreCredential;

  static equals(a: DatastoreCredential | PlainMessage<DatastoreCredential> | undefined, b: DatastoreCredential | PlainMessage<DatastoreCredential> | undefined): boolean;
}

