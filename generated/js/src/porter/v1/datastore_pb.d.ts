// @generated by protoc-gen-es v1.6.0
// @generated from file porter/v1/datastore.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { EnumCloudProvider } from "./cluster_pb.js";
import type { Postgres, Redis } from "./addons_pb.js";

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

  /**
   * @generated from enum value: ENUM_DATASTORE_KIND_AWS_ELASTICACHE = 2;
   */
  AWS_ELASTICACHE = 2,

  /**
   * ENUM_DATASTORE_KIND_MANAGED_POSTGRES refers to the postgres helm chart deployed on customer cluster
   *
   * @generated from enum value: ENUM_DATASTORE_KIND_MANAGED_POSTGRES = 3;
   */
  MANAGED_POSTGRES = 3,

  /**
   * ENUM_DATASTORE_KIND_MANAGED_REDIS refers to the redis helm chart deployed on customer cluster
   *
   * @generated from enum value: ENUM_DATASTORE_KIND_MANAGED_REDIS = 4;
   */
  MANAGED_REDIS = 4,

  /**
   * @generated from enum value: ENUM_DATASTORE_KIND_NEON = 5;
   */
  NEON = 5,

  /**
   * @generated from enum value: ENUM_DATASTORE_KIND_UPSTASH = 6;
   */
  UPSTASH = 6,
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
 * @generated from enum porter.v1.EnumAwsElasticacheEngine
 */
export declare enum EnumAwsElasticacheEngine {
  /**
   * @generated from enum value: ENUM_AWS_ELASTICACHE_ENGINE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: ENUM_AWS_ELASTICACHE_ENGINE_REDIS = 1;
   */
  REDIS = 1,

  /**
   * @generated from enum value: ENUM_AWS_ELASTICACHE_ENGINE_MEMCACHED = 2;
   */
  MEMCACHED = 2,
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
  } | {
    /**
     * @generated from field: porter.v1.AwsElasticache aws_elasticache_kind = 9;
     */
    value: AwsElasticache;
    case: "awsElasticacheKind";
  } | {
    /**
     * @generated from field: porter.v1.Postgres managed_postgres_kind = 10;
     */
    value: Postgres;
    case: "managedPostgresKind";
  } | {
    /**
     * @generated from field: porter.v1.Redis managed_redis_kind = 11;
     */
    value: Redis;
    case: "managedRedisKind";
  } | {
    /**
     * @generated from field: porter.v1.Neon neon_kind = 12;
     */
    value: Neon;
    case: "neonKind";
  } | {
    /**
     * @generated from field: porter.v1.Upstash upstash_kind = 13;
     */
    value: Upstash;
    case: "upstashKind";
  } | { case: undefined; value?: undefined };

  /**
   * connected_clusters describes the connection of this datastore to clusters
   *
   * @generated from field: porter.v1.ConnectedClusters connected_clusters = 8;
   */
  connectedClusters?: ConnectedClusters;

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
 * @generated from message porter.v1.ConnectedClusters
 */
export declare class ConnectedClusters extends Message<ConnectedClusters> {
  /**
   * connected_cluster_ids is a list of cluster ids that this datastore is connected to
   *
   * @generated from field: repeated int64 connected_cluster_ids = 1;
   */
  connectedClusterIds: bigint[];

  constructor(data?: PartialMessage<ConnectedClusters>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.ConnectedClusters";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ConnectedClusters;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ConnectedClusters;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ConnectedClusters;

  static equals(a: ConnectedClusters | PlainMessage<ConnectedClusters> | undefined, b: ConnectedClusters | PlainMessage<ConnectedClusters> | undefined): boolean;
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
 * @generated from message porter.v1.AwsElasticache
 */
export declare class AwsElasticache extends Message<AwsElasticache> {
  /**
   * instance_class is the instance class of the database
   *
   * @generated from field: optional string instance_class = 1;
   */
  instanceClass?: string;

  /**
   * master_user_password_literal is the string value of the password; this is only used for creating the database password secret and is wiped when the contract is saved
   *
   * @generated from field: optional string master_user_password_literal = 2;
   */
  masterUserPasswordLiteral?: string;

  /**
   * engine is the engine of the datastore
   *
   * @generated from field: porter.v1.EnumAwsElasticacheEngine engine = 3;
   */
  engine: EnumAwsElasticacheEngine;

  /**
   * engine_version is the version of the engine
   *
   * @generated from field: optional string engine_version = 4;
   */
  engineVersion?: string;

  constructor(data?: PartialMessage<AwsElasticache>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.AwsElasticache";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AwsElasticache;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AwsElasticache;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AwsElasticache;

  static equals(a: AwsElasticache | PlainMessage<AwsElasticache> | undefined, b: AwsElasticache | PlainMessage<AwsElasticache> | undefined): boolean;
}

/**
 * @generated from message porter.v1.Neon
 */
export declare class Neon extends Message<Neon> {
  constructor(data?: PartialMessage<Neon>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Neon";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Neon;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Neon;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Neon;

  static equals(a: Neon | PlainMessage<Neon> | undefined, b: Neon | PlainMessage<Neon> | undefined): boolean;
}

/**
 * @generated from message porter.v1.Upstash
 */
export declare class Upstash extends Message<Upstash> {
  constructor(data?: PartialMessage<Upstash>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Upstash";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Upstash;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Upstash;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Upstash;

  static equals(a: Upstash | PlainMessage<Upstash> | undefined, b: Upstash | PlainMessage<Upstash> | undefined): boolean;
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

