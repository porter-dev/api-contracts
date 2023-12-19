// @generated by protoc-gen-es v1.4.2
// @generated from file porter/v1/addons.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { EnvGroup } from "./porter_app_pb.js";

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
  } | { case: undefined; value?: undefined };

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

  constructor(data?: PartialMessage<Postgres>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Postgres";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Postgres;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Postgres;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Postgres;

  static equals(a: Postgres | PlainMessage<Postgres> | undefined, b: Postgres | PlainMessage<Postgres> | undefined): boolean;
}

