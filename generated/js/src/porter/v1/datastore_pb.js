// @generated by protoc-gen-es v1.10.0
// @generated from file porter/v1/datastore.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { EnumCloudProvider } from "./cluster_pb.js";
import { Postgres, Redis } from "./addons_pb.js";

/**
 * @generated from enum porter.v1.EnumDatastoreKind
 */
export const EnumDatastoreKind = /*@__PURE__*/ proto3.makeEnum(
  "porter.v1.EnumDatastoreKind",
  [
    {no: 0, name: "ENUM_DATASTORE_KIND_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "ENUM_DATASTORE_KIND_AWS_RDS", localName: "AWS_RDS"},
    {no: 2, name: "ENUM_DATASTORE_KIND_AWS_ELASTICACHE", localName: "AWS_ELASTICACHE"},
    {no: 3, name: "ENUM_DATASTORE_KIND_MANAGED_POSTGRES", localName: "MANAGED_POSTGRES"},
    {no: 4, name: "ENUM_DATASTORE_KIND_MANAGED_REDIS", localName: "MANAGED_REDIS"},
    {no: 5, name: "ENUM_DATASTORE_KIND_NEON", localName: "NEON"},
    {no: 6, name: "ENUM_DATASTORE_KIND_UPSTASH", localName: "UPSTASH"},
  ],
);

/**
 * @generated from enum porter.v1.EnumAwsRdsEngine
 */
export const EnumAwsRdsEngine = /*@__PURE__*/ proto3.makeEnum(
  "porter.v1.EnumAwsRdsEngine",
  [
    {no: 0, name: "ENUM_AWS_RDS_ENGINE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "ENUM_AWS_RDS_ENGINE_POSTGRESQL", localName: "POSTGRESQL"},
    {no: 2, name: "ENUM_AWS_RDS_ENGINE_AURORA_POSTGRESQL", localName: "AURORA_POSTGRESQL"},
  ],
);

/**
 * @generated from enum porter.v1.EnumAwsElasticacheEngine
 */
export const EnumAwsElasticacheEngine = /*@__PURE__*/ proto3.makeEnum(
  "porter.v1.EnumAwsElasticacheEngine",
  [
    {no: 0, name: "ENUM_AWS_ELASTICACHE_ENGINE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "ENUM_AWS_ELASTICACHE_ENGINE_REDIS", localName: "REDIS"},
    {no: 2, name: "ENUM_AWS_ELASTICACHE_ENGINE_MEMCACHED", localName: "MEMCACHED"},
  ],
);

/**
 * ManagedDatastore is the specification for a Porter-managed datastore
 *
 * @generated from message porter.v1.ManagedDatastore
 */
export const ManagedDatastore = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.ManagedDatastore",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "cloud_provider", kind: "enum", T: proto3.getEnumType(EnumCloudProvider) },
    { no: 3, name: "cloud_provider_credential_identifier", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "region", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "kind", kind: "enum", T: proto3.getEnumType(EnumDatastoreKind) },
    { no: 7, name: "aws_rds_kind", kind: "message", T: AwsRds, oneof: "kind_values" },
    { no: 9, name: "aws_elasticache_kind", kind: "message", T: AwsElasticache, oneof: "kind_values" },
    { no: 10, name: "managed_postgres_kind", kind: "message", T: Postgres, oneof: "kind_values" },
    { no: 11, name: "managed_redis_kind", kind: "message", T: Redis, oneof: "kind_values" },
    { no: 12, name: "neon_kind", kind: "message", T: Neon, oneof: "kind_values" },
    { no: 13, name: "upstash_kind", kind: "message", T: Upstash, oneof: "kind_values" },
    { no: 8, name: "connected_clusters", kind: "message", T: ConnectedClusters },
    { no: 14, name: "public_networking", kind: "message", T: PublicNetworking },
  ],
);

/**
 * @generated from message porter.v1.ConnectedClusters
 */
export const ConnectedClusters = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.ConnectedClusters",
  () => [
    { no: 1, name: "connected_cluster_ids", kind: "scalar", T: 3 /* ScalarType.INT64 */, repeated: true },
  ],
);

/**
 * @generated from message porter.v1.PublicNetworking
 */
export const PublicNetworking = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.PublicNetworking",
  () => [
    { no: 1, name: "cidr_allowlist", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message porter.v1.DatastorePasswordSecretRef
 */
export const DatastorePasswordSecretRef = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.DatastorePasswordSecretRef",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "namespace", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.AwsRds
 */
export const AwsRds = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.AwsRds",
  () => [
    { no: 1, name: "database_name", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 2, name: "master_username", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 3, name: "master_user_password_literal", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 4, name: "master_user_password_secret_ref", kind: "message", T: DatastorePasswordSecretRef },
    { no: 5, name: "allocated_storage_gigabytes", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
    { no: 6, name: "instance_class", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 7, name: "engine", kind: "enum", T: proto3.getEnumType(EnumAwsRdsEngine) },
    { no: 8, name: "engine_version", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ],
);

/**
 * @generated from message porter.v1.AwsElasticache
 */
export const AwsElasticache = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.AwsElasticache",
  () => [
    { no: 1, name: "instance_class", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 2, name: "master_user_password_literal", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 3, name: "engine", kind: "enum", T: proto3.getEnumType(EnumAwsElasticacheEngine) },
    { no: 4, name: "engine_version", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ],
);

/**
 * @generated from message porter.v1.Neon
 */
export const Neon = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.Neon",
  [],
);

/**
 * @generated from message porter.v1.Upstash
 */
export const Upstash = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.Upstash",
  [],
);

/**
 * UpdateDatastorePayload is used to send messages via nats to update a datastore
 *
 * @generated from message porter.v1.UpdateDatastorePayload
 */
export const UpdateDatastorePayload = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.UpdateDatastorePayload",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "datastore_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * DatastoreCredential is used to connect to a datastore
 *
 * @generated from message porter.v1.DatastoreCredential
 */
export const DatastoreCredential = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.DatastoreCredential",
  () => [
    { no: 1, name: "host", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "database_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "port", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

