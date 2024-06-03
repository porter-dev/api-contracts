// @generated by protoc-gen-es v1.10.0
// @generated from file porter/v1/addons.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { EnvGroup, HelmOverrides } from "./porter_app_pb.js";
import type { Domain } from "./service_pb.js";

/**
 * @generated from enum porter.v1.AddonType
 */
export declare enum AddonType {
  /**
   * ADDON_TYPE_UNSPECIFIED is the default value
   *
   * @generated from enum value: ADDON_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * ADDON_TYPE_POSTGRES is the postgres addon type
   *
   * @generated from enum value: ADDON_TYPE_POSTGRES = 1;
   */
  POSTGRES = 1,

  /**
   * ADDON_TYPE_REDIS is the redis addon type
   *
   * @generated from enum value: ADDON_TYPE_REDIS = 2;
   */
  REDIS = 2,

  /**
   * ADDON_TYPE_DATADOG is the datadog addon type
   *
   * @generated from enum value: ADDON_TYPE_DATADOG = 3;
   */
  DATADOG = 3,

  /**
   * ADDON_TYPE_MEZMO is the mezmo addon type
   *
   * @generated from enum value: ADDON_TYPE_MEZMO = 4;
   */
  MEZMO = 4,

  /**
   * ADDON_TYPE_METABASE is the metabase addon type
   *
   * @generated from enum value: ADDON_TYPE_METABASE = 5;
   */
  METABASE = 5,

  /**
   * ADDON_TYPE_NEWRELIC is the newrelic addon type
   *
   * @generated from enum value: ADDON_TYPE_NEWRELIC = 6;
   */
  NEWRELIC = 6,

  /**
   * ADDON_TYPE_TAILSCALE is the tailscale addon type
   *
   * @generated from enum value: ADDON_TYPE_TAILSCALE = 7;
   */
  TAILSCALE = 7,

  /**
   * ADDON_TYPE_QUIVR is the quivr addon type
   *
   * @generated from enum value: ADDON_TYPE_QUIVR = 8;
   */
  QUIVR = 8,

  /**
   * ADDON_TYPE_DEEPGRAM is the deepgram addon type
   *
   * @generated from enum value: ADDON_TYPE_DEEPGRAM = 9;
   */
  DEEPGRAM = 9,
}

/**
 * AddonStatus specifies the status of an addon installation
 * this is currently used to track the installation status for complex addons like deepgram and other LLM addons
 *
 * @generated from enum porter.v1.AddonStatus
 */
export declare enum AddonStatus {
  /**
   * ADDON_STATUS_UNSPECIFIED is the default value
   *
   * @generated from enum value: ADDON_STATUS_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * ADDON_STATUS_QUOTA_PENDING is the status when an addon installation is waiting for a quota upgrade
   *
   * @generated from enum value: ADDON_STATUS_QUOTA_PENDING = 1;
   */
  QUOTA_PENDING = 1,

  /**
   * ADDON_STATUS_QUOTA_FAILED is the status when an addon can't be installed due to a failed quota upgrade
   *
   * @generated from enum value: ADDON_STATUS_QUOTA_FAILED = 2;
   */
  QUOTA_FAILED = 2,

  /**
   * ADDON_STATUS_INFRA_PROVISIONING is the status when an addon installation is waiting for infrastructure reprovisioning
   *
   * @generated from enum value: ADDON_STATUS_INFRA_PROVISIONING = 3;
   */
  INFRA_PROVISIONING = 3,

  /**
   * ADDON_STATUS_INFRA_PROVISIONING_FAILED is the status when an addon can't be installed due to a failed infrastructure reprovisioning
   *
   * @generated from enum value: ADDON_STATUS_INFRA_PROVISIONING_FAILED = 4;
   */
  INFRA_PROVISIONING_FAILED = 4,

  /**
   * ADDON_STATUS_DEPLOYING is the status when an addon is being deployed
   *
   * @generated from enum value: ADDON_STATUS_DEPLOYING = 5;
   */
  DEPLOYING = 5,
}

/**
 * PrerequisiteAddon specifies an addon that must be installed before any apps can be installed
 * the addon should be installed with the specified config
 *
 * @generated from message porter.v1.PrerequisiteAddon
 */
export declare class PrerequisiteAddon extends Message<PrerequisiteAddon> {
  /**
   * commit_sha is the commit SHA of the addon
   *
   * @generated from field: string commit_sha = 1;
   */
  commitSha: string;

  constructor(data?: PartialMessage<PrerequisiteAddon>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.PrerequisiteAddon";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PrerequisiteAddon;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PrerequisiteAddon;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PrerequisiteAddon;

  static equals(a: PrerequisiteAddon | PlainMessage<PrerequisiteAddon> | undefined, b: PrerequisiteAddon | PlainMessage<PrerequisiteAddon> | undefined): boolean;
}

/**
 * Addon is the configuration object for tooling or services that can be applied to the cluster alongside porter apps.
 *
 * @generated from message porter.v1.Addon
 */
export declare class Addon extends Message<Addon> {
  /**
   * name is the name of the addon
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * type is the type of type of addon
   *
   * @generated from field: porter.v1.AddonType type = 2;
   */
  type: AddonType;

  /**
   * env_groups is a list of environment groups that can be applied to the addon
   *
   * @generated from field: repeated porter.v1.EnvGroup env_groups = 3;
   */
  envGroups: EnvGroup[];

  /**
   * config is the addon-specific configuration
   *
   * @generated from oneof porter.v1.Addon.config
   */
  config: {
    /**
     * postgres is the configuration for the postgres addon
     *
     * @generated from field: porter.v1.Postgres postgres = 4;
     */
    value: Postgres;
    case: "postgres";
  } | {
    /**
     * redis is the configuration for the redis addon
     *
     * @generated from field: porter.v1.Redis redis = 5;
     */
    value: Redis;
    case: "redis";
  } | {
    /**
     * datadog is the configuration for the datadog addon
     *
     * @generated from field: porter.v1.Datadog datadog = 6;
     */
    value: Datadog;
    case: "datadog";
  } | {
    /**
     * mezmo is the configuration for the mezmo addon
     *
     * @generated from field: porter.v1.Mezmo mezmo = 7;
     */
    value: Mezmo;
    case: "mezmo";
  } | {
    /**
     * metabase is the configuration for the metabase addon
     *
     * @generated from field: porter.v1.Metabase metabase = 8;
     */
    value: Metabase;
    case: "metabase";
  } | {
    /**
     * Newrelic is the configuration for the newrelic addon
     *
     * @generated from field: porter.v1.Newrelic newrelic = 9;
     */
    value: Newrelic;
    case: "newrelic";
  } | {
    /**
     * Tailscale is the configuration for the tailscale addon
     *
     * @generated from field: porter.v1.Tailscale tailscale = 10;
     */
    value: Tailscale;
    case: "tailscale";
  } | {
    /**
     * Quivr is the configuration for the quivr addon
     *
     * @generated from field: porter.v1.Quivr quivr = 11;
     */
    value: Quivr;
    case: "quivr";
  } | {
    /**
     * Deepgram is the configuration for the deepgram addon
     *
     * @generated from field: porter.v1.Deepgram deepgram = 12;
     */
    value: Deepgram;
    case: "deepgram";
  } | { case: undefined; value?: undefined };

  /**
   * helm_overrides is stringified json of raw Helm overrides to use for the addon. These will override any values generated by the contract.
   *
   * @generated from field: porter.v1.HelmOverrides helm_overrides = 13;
   */
  helmOverrides?: HelmOverrides;

  constructor(data?: PartialMessage<Addon>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Addon";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Addon;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Addon;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Addon;

  static equals(a: Addon | PlainMessage<Addon> | undefined, b: Addon | PlainMessage<Addon> | undefined): boolean;
}

/**
 * Postgres is the configuration for postgres
 *
 * @generated from message porter.v1.Postgres
 */
export declare class Postgres extends Message<Postgres> {
  /**
   * cpu_cores is the number of CPU cores to allocate to the database
   *
   * @generated from field: float cpu_cores = 1;
   */
  cpuCores: number;

  /**
   * ram_megabytes is the amount of memory to allocate to the database
   *
   * @generated from field: int32 ram_megabytes = 2;
   */
  ramMegabytes: number;

  /**
   * storage_gigabytes is the amount of storage to allocate to the database
   *
   * @generated from field: int32 storage_gigabytes = 3;
   */
  storageGigabytes: number;

  /**
   * master_username is the username of the database
   *
   * @generated from field: optional string master_username = 4;
   */
  masterUsername?: string;

  /**
   * master_user_password_literal is the string value of the password; this is only used for creating the datastore password secret and is wiped when the contract is saved
   *
   * @generated from field: optional string master_user_password_literal = 5;
   */
  masterUserPasswordLiteral?: string;

  constructor(data?: PartialMessage<Postgres>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Postgres";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Postgres;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Postgres;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Postgres;

  static equals(a: Postgres | PlainMessage<Postgres> | undefined, b: Postgres | PlainMessage<Postgres> | undefined): boolean;
}

/**
 * Redis is the configuration for redis
 *
 * @generated from message porter.v1.Redis
 */
export declare class Redis extends Message<Redis> {
  /**
   * cpu_cores is the number of CPU cores to allocate to redis
   *
   * @generated from field: float cpu_cores = 1;
   */
  cpuCores: number;

  /**
   * ram_megabytes is the amount of memory to allocate to redis
   *
   * @generated from field: int32 ram_megabytes = 2;
   */
  ramMegabytes: number;

  /**
   * storage_gigabytes is the amount of storage to allocate to redis
   *
   * @generated from field: int32 storage_gigabytes = 3;
   */
  storageGigabytes: number;

  /**
   * master_user_password_literal is the string value of the password; this is only used for creating the datastore password secret and is wiped when the contract is saved
   *
   * @generated from field: optional string master_user_password_literal = 4;
   */
  masterUserPasswordLiteral?: string;

  constructor(data?: PartialMessage<Redis>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Redis";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Redis;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Redis;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Redis;

  static equals(a: Redis | PlainMessage<Redis> | undefined, b: Redis | PlainMessage<Redis> | undefined): boolean;
}

/**
 * Datadog is the configuration for Datadog
 *
 * @generated from message porter.v1.Datadog
 */
export declare class Datadog extends Message<Datadog> {
  /**
   * site is the datadog site url
   *
   * @generated from field: optional string site = 1;
   */
  site?: string;

  /**
   * api_key is the api key
   *
   * @generated from field: optional string api_key = 2;
   */
  apiKey?: string;

  /**
   * logging_enabled determines whether all container logs go to datadog
   *
   * @generated from field: optional bool logging_enabled = 3;
   */
  loggingEnabled?: boolean;

  /**
   * dogstatsd_enabled determines whether dogstatsd is enabled
   *
   * @generated from field: optional bool dogstatsd_enabled = 4;
   */
  dogstatsdEnabled?: boolean;

  /**
   * apm_enabled determines whether apm is enabled
   *
   * @generated from field: optional bool apm_enabled = 5;
   */
  apmEnabled?: boolean;

  /**
   * cpu_cores is the number of CPU cores to allocate to datadog
   *
   * @generated from field: optional float cpu_cores = 6;
   */
  cpuCores?: number;

  /**
   * ram_megabytes is the amount of memory to allocate to datadog
   *
   * @generated from field: optional int32 ram_megabytes = 7;
   */
  ramMegabytes?: number;

  constructor(data?: PartialMessage<Datadog>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Datadog";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Datadog;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Datadog;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Datadog;

  static equals(a: Datadog | PlainMessage<Datadog> | undefined, b: Datadog | PlainMessage<Datadog> | undefined): boolean;
}

/**
 * Mezmo is the configuration for Mezmo
 *
 * @generated from message porter.v1.Mezmo
 */
export declare class Mezmo extends Message<Mezmo> {
  /**
   * ingestion_key is the mezmo ingestion key. This is stored as a secret on the client cluster
   *
   * @generated from field: optional string ingestion_key = 1;
   */
  ingestionKey?: string;

  constructor(data?: PartialMessage<Mezmo>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Mezmo";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Mezmo;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Mezmo;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Mezmo;

  static equals(a: Mezmo | PlainMessage<Mezmo> | undefined, b: Mezmo | PlainMessage<Mezmo> | undefined): boolean;
}

/**
 * MetabaseDatastore is the configuration for a datastore linked with Metabase
 *
 * @generated from message porter.v1.MetabaseDatastore
 */
export declare class MetabaseDatastore extends Message<MetabaseDatastore> {
  /**
   * @generated from field: string host = 1;
   */
  host: string;

  /**
   * @generated from field: int64 port = 2;
   */
  port: bigint;

  /**
   * @generated from field: string database_name = 3;
   */
  databaseName: string;

  /**
   * @generated from field: string master_username = 4;
   */
  masterUsername: string;

  /**
   * @generated from field: string master_user_password_literal = 5;
   */
  masterUserPasswordLiteral: string;

  constructor(data?: PartialMessage<MetabaseDatastore>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.MetabaseDatastore";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MetabaseDatastore;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MetabaseDatastore;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MetabaseDatastore;

  static equals(a: MetabaseDatastore | PlainMessage<MetabaseDatastore> | undefined, b: MetabaseDatastore | PlainMessage<MetabaseDatastore> | undefined): boolean;
}

/**
 * Metabase is the configuration for Metabase
 *
 * @generated from message porter.v1.Metabase
 */
export declare class Metabase extends Message<Metabase> {
  /**
   * domains is the list of domains for this service
   *
   * @generated from field: repeated porter.v1.Domain domains = 1;
   */
  domains: Domain[];

  /**
   * cpu_cores is the number of CPU cores to allocate to metabase
   *
   * @generated from field: optional float cpu_cores = 2;
   */
  cpuCores?: number;

  /**
   * ram_megabytes is the amount of memory to allocate to metabase
   *
   * @generated from field: optional int32 ram_megabytes = 3;
   */
  ramMegabytes?: number;

  /**
   * datastore is the configuration of the datastore that metabase is connected to
   *
   * @generated from field: optional porter.v1.MetabaseDatastore datastore = 4;
   */
  datastore?: MetabaseDatastore;

  /**
   * ingress_enabled describes whether the metabase instance has ingress enabled
   *
   * @generated from field: optional bool ingress_enabled = 5;
   */
  ingressEnabled?: boolean;

  constructor(data?: PartialMessage<Metabase>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Metabase";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Metabase;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Metabase;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Metabase;

  static equals(a: Metabase | PlainMessage<Metabase> | undefined, b: Metabase | PlainMessage<Metabase> | undefined): boolean;
}

/**
 * Newrelic is the configuration for Newrelic
 *
 * @generated from message porter.v1.Newrelic
 */
export declare class Newrelic extends Message<Newrelic> {
  /**
   * cluster_name is the name of the cluster
   *
   * @generated from field: optional string cluster_name = 1;
   */
  clusterName?: string;

  /**
   * license_key is the license key
   *
   * @generated from field: optional string license_key = 2;
   */
  licenseKey?: string;

  /**
   * insights_key is the insights key
   *
   * @generated from field: optional string insights_key = 3;
   */
  insightsKey?: string;

  /**
   * personal_api_key is the personal api key
   *
   * @generated from field: optional string personal_api_key = 4;
   */
  personalApiKey?: string;

  /**
   * account_id is the account id
   *
   * @generated from field: optional string account_id = 5;
   */
  accountId?: string;

  /**
   * logging_enabled determines whether logging is enabled
   *
   * @generated from field: optional bool logging_enabled = 6;
   */
  loggingEnabled?: boolean;

  /**
   * metrics_adapter_enabled determines whether metrics adapter is enabled
   *
   * @generated from field: optional bool metrics_adapter_enabled = 7;
   */
  metricsAdapterEnabled?: boolean;

  /**
   * prometheus_enabled determines whether prometheus is enabled
   *
   * @generated from field: optional bool prometheus_enabled = 8;
   */
  prometheusEnabled?: boolean;

  /**
   * pixie_enabled determines whether pixie is enabled
   *
   * @generated from field: optional bool pixie_enabled = 9;
   */
  pixieEnabled?: boolean;

  /**
   * kube_events_enabled determines whether tracking for kube events is enabled
   *
   * @generated from field: optional bool kube_events_enabled = 10;
   */
  kubeEventsEnabled?: boolean;

  constructor(data?: PartialMessage<Newrelic>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Newrelic";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Newrelic;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Newrelic;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Newrelic;

  static equals(a: Newrelic | PlainMessage<Newrelic> | undefined, b: Newrelic | PlainMessage<Newrelic> | undefined): boolean;
}

/**
 * Tailscale is the configuration for Tailscale
 *
 * @generated from message porter.v1.Tailscale
 */
export declare class Tailscale extends Message<Tailscale> {
  /**
   * auth_key is the auth key for Tailscale
   *
   * @generated from field: optional string auth_key = 1;
   */
  authKey?: string;

  /**
   * subnet_routes are the subnet routes for Tailscale
   *
   * @generated from field: repeated string subnet_routes = 2;
   */
  subnetRoutes: string[];

  constructor(data?: PartialMessage<Tailscale>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Tailscale";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Tailscale;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Tailscale;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Tailscale;

  static equals(a: Tailscale | PlainMessage<Tailscale> | undefined, b: Tailscale | PlainMessage<Tailscale> | undefined): boolean;
}

/**
 * Quivr is the configuration for Quivr
 *
 * @generated from message porter.v1.Quivr
 */
export declare class Quivr extends Message<Quivr> {
  /**
   * domains is the list of domains for the service
   *
   * @generated from field: repeated porter.v1.Domain domains = 1;
   */
  domains: Domain[];

  /**
   * ingress_enabled describes whether the quivr instance has external ingress enabled
   *
   * @generated from field: optional bool ingress_enabled = 2;
   */
  ingressEnabled?: boolean;

  /**
   * @generated from field: optional string openai_api_key = 3;
   */
  openaiApiKey?: string;

  /**
   * @generated from field: optional string supabase_url = 4;
   */
  supabaseUrl?: string;

  /**
   * @generated from field: optional string supabase_service_key = 5;
   */
  supabaseServiceKey?: string;

  /**
   * @generated from field: optional string pg_database_url = 6;
   */
  pgDatabaseUrl?: string;

  /**
   * @generated from field: optional string jwt_secret_key = 7;
   */
  jwtSecretKey?: string;

  /**
   * @generated from field: optional string cohere_api_key = 8;
   */
  cohereApiKey?: string;

  /**
   * @generated from field: optional string anthropic_api_key = 9;
   */
  anthropicApiKey?: string;

  /**
   * @generated from field: optional string quivr_domain = 10;
   */
  quivrDomain?: string;

  /**
   * @generated from field: optional string brave_search_api_key = 11;
   */
  braveSearchApiKey?: string;

  constructor(data?: PartialMessage<Quivr>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Quivr";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Quivr;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Quivr;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Quivr;

  static equals(a: Quivr | PlainMessage<Quivr> | undefined, b: Quivr | PlainMessage<Quivr> | undefined): boolean;
}

/**
 * Deepgram is the configuration for Deepgram
 *
 * @generated from message porter.v1.Deepgram
 */
export declare class Deepgram extends Message<Deepgram> {
  /**
   * api_key is the deepgram API key
   *
   * @generated from field: optional string api_key = 1;
   */
  apiKey?: string;

  /**
   * ecr_username is the username to fetch the proprietary image
   *
   * @generated from field: optional string ecr_username = 2;
   */
  ecrUsername?: string;

  /**
   * ecr_password is the password to fetch the proprietary image
   *
   * @generated from field: optional string ecr_password = 3;
   */
  ecrPassword?: string;

  /**
   * ecr_email is email to fetch the proprietary image
   *
   * @generated from field: optional string ecr_email = 4;
   */
  ecrEmail?: string;

  /**
   * release_tag is the version of the deepgram application
   *
   * @generated from field: optional string release_tag = 5;
   */
  releaseTag?: string;

  /**
   * model_urls is the list of urls to download deepgram models from
   *
   * @generated from field: repeated string model_urls = 6;
   */
  modelUrls: string[];

  /**
   * manage_quota is a boolean that tells the backend if it should go ahead and modify quotas before installing the addon
   * if this field is not set to true, the endpoint will not be able to install the addon if the required quotas are not high enough
   *
   * @generated from field: optional bool manage_quota = 7;
   */
  manageQuota?: boolean;

  constructor(data?: PartialMessage<Deepgram>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Deepgram";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Deepgram;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Deepgram;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Deepgram;

  static equals(a: Deepgram | PlainMessage<Deepgram> | undefined, b: Deepgram | PlainMessage<Deepgram> | undefined): boolean;
}

