// @generated by protoc-gen-es v1.8.0
// @generated from file porter/v1/addons.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { EnvGroup } from "./porter_app_pb.js";
import { Domain } from "./service_pb.js";

/**
 * @generated from enum porter.v1.AddonType
 */
export const AddonType = /*@__PURE__*/ proto3.makeEnum(
  "porter.v1.AddonType",
  [
    {no: 0, name: "ADDON_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "ADDON_TYPE_POSTGRES", localName: "POSTGRES"},
    {no: 2, name: "ADDON_TYPE_REDIS", localName: "REDIS"},
    {no: 3, name: "ADDON_TYPE_DATADOG", localName: "DATADOG"},
    {no: 4, name: "ADDON_TYPE_MEZMO", localName: "MEZMO"},
    {no: 5, name: "ADDON_TYPE_METABASE", localName: "METABASE"},
    {no: 6, name: "ADDON_TYPE_NEWRELIC", localName: "NEWRELIC"},
    {no: 7, name: "ADDON_TYPE_TAILSCALE", localName: "TAILSCALE"},
  ],
);

/**
 * PrerequisiteAddon specifies an addon that must be installed before any apps can be installed
 * the addon should be installed with the specified config
 *
 * @generated from message porter.v1.PrerequisiteAddon
 */
export const PrerequisiteAddon = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.PrerequisiteAddon",
  () => [
    { no: 1, name: "commit_sha", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * Addon is the configuration object for tooling or services that can be applied to the cluster alongside porter apps.
 *
 * @generated from message porter.v1.Addon
 */
export const Addon = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.Addon",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "type", kind: "enum", T: proto3.getEnumType(AddonType) },
    { no: 3, name: "env_groups", kind: "message", T: EnvGroup, repeated: true },
    { no: 4, name: "postgres", kind: "message", T: Postgres, oneof: "config" },
    { no: 5, name: "redis", kind: "message", T: Redis, oneof: "config" },
    { no: 6, name: "datadog", kind: "message", T: Datadog, oneof: "config" },
    { no: 7, name: "mezmo", kind: "message", T: Mezmo, oneof: "config" },
    { no: 8, name: "metabase", kind: "message", T: Metabase, oneof: "config" },
    { no: 9, name: "newrelic", kind: "message", T: Newrelic, oneof: "config" },
    { no: 10, name: "tailscale", kind: "message", T: Tailscale, oneof: "config" },
  ],
);

/**
 * Postgres is the configuration for postgres
 *
 * @generated from message porter.v1.Postgres
 */
export const Postgres = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.Postgres",
  () => [
    { no: 1, name: "cpu_cores", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 2, name: "ram_megabytes", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "storage_gigabytes", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "master_username", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 5, name: "master_user_password_literal", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ],
);

/**
 * Redis is the configuration for redis
 *
 * @generated from message porter.v1.Redis
 */
export const Redis = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.Redis",
  () => [
    { no: 1, name: "cpu_cores", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 2, name: "ram_megabytes", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "storage_gigabytes", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "master_user_password_literal", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ],
);

/**
 * Datadog is the configuration for Datadog
 *
 * @generated from message porter.v1.Datadog
 */
export const Datadog = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.Datadog",
  () => [
    { no: 1, name: "site", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 2, name: "api_key", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 3, name: "logging_enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 4, name: "dogstatsd_enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 5, name: "apm_enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 6, name: "cpu_cores", kind: "scalar", T: 2 /* ScalarType.FLOAT */, opt: true },
    { no: 7, name: "ram_megabytes", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
  ],
);

/**
 * Mezmo is the configuration for Mezmo
 *
 * @generated from message porter.v1.Mezmo
 */
export const Mezmo = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.Mezmo",
  () => [
    { no: 1, name: "ingestion_key", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ],
);

/**
 * MetabaseDatastore is the configuration for a datastore linked with Metabase
 *
 * @generated from message porter.v1.MetabaseDatastore
 */
export const MetabaseDatastore = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.MetabaseDatastore",
  () => [
    { no: 1, name: "host", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "port", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "database_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "master_username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "master_user_password_literal", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * Metabase is the configuration for Metabase
 *
 * @generated from message porter.v1.Metabase
 */
export const Metabase = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.Metabase",
  () => [
    { no: 1, name: "domains", kind: "message", T: Domain, repeated: true },
    { no: 2, name: "cpu_cores", kind: "scalar", T: 2 /* ScalarType.FLOAT */, opt: true },
    { no: 3, name: "ram_megabytes", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 4, name: "datastore", kind: "message", T: MetabaseDatastore, opt: true },
  ],
);

/**
 * Newrelic is the configuration for Newrelic
 *
 * @generated from message porter.v1.Newrelic
 */
export const Newrelic = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.Newrelic",
  () => [
    { no: 1, name: "cluster_name", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 2, name: "license_key", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 3, name: "insights_key", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 4, name: "personal_api_key", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 5, name: "account_id", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 6, name: "logging_enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 7, name: "metrics_adapter_enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 8, name: "prometheus_enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 9, name: "pixie_enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ],
);

/**
 * Tailscale is the configuration for Tailscale
 *
 * @generated from message porter.v1.Tailscale
 */
export const Tailscale = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.Tailscale",
  () => [
    { no: 1, name: "auth_key", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 2, name: "subnet_routes", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

