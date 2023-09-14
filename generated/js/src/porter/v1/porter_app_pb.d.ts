// @generated by protoc-gen-es v1.3.0
// @generated from file porter/v1/porter_app.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Service } from "./service_pb.js";

/**
 * PorterApp is the top-level configuration for a Porter application, usually found in porter.yaml
 *
 * @generated from message porter.v1.PorterApp
 */
export declare class PorterApp extends Message<PorterApp> {
  /**
   * name is the name of the application
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * services is a map of service names to service configurations
   *
   * @generated from field: map<string, porter.v1.Service> services = 2;
   */
  services: { [key: string]: Service };

  /**
   * env is deprecated in favor of env groups. It was a map of environment variable names to values
   *
   * @generated from field: map<string, string> env = 3 [deprecated = true];
   * @deprecated
   */
  env: { [key: string]: string };

  /**
   * build is the build settings for the application
   *
   * @generated from field: porter.v1.Build build = 4;
   */
  build?: Build;

  /**
   * predeploy is a job service to run before deploying the application
   *
   * @generated from field: porter.v1.Service predeploy = 5;
   */
  predeploy?: Service;

  /**
   * image is the image to use for a given revision of the application
   *
   * @generated from field: porter.v1.AppImage image = 6;
   */
  image?: AppImage;

  /**
   * env_groups is a map of environment variable group names to environment variable group configurations
   *
   * @generated from field: repeated porter.v1.EnvGroup env_groups = 7;
   */
  envGroups: EnvGroup[];

  constructor(data?: PartialMessage<PorterApp>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.PorterApp";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PorterApp;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PorterApp;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PorterApp;

  static equals(a: PorterApp | PlainMessage<PorterApp> | undefined, b: PorterApp | PlainMessage<PorterApp> | undefined): boolean;
}

/**
 * @generated from message porter.v1.EnvGroup
 */
export declare class EnvGroup extends Message<EnvGroup> {
  /**
   * name is the name of the environment variable group
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * version is the version of the environment variable group
   *
   * @generated from field: int64 version = 2;
   */
  version: bigint;

  constructor(data?: PartialMessage<EnvGroup>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.EnvGroup";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EnvGroup;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EnvGroup;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EnvGroup;

  static equals(a: EnvGroup | PlainMessage<EnvGroup> | undefined, b: EnvGroup | PlainMessage<EnvGroup> | undefined): boolean;
}

/**
 * @generated from message porter.v1.EnvGroupVariables
 */
export declare class EnvGroupVariables extends Message<EnvGroupVariables> {
  /**
   * normal is a map of non-sensitive variable names to values
   *
   * @generated from field: map<string, string> normal = 1;
   */
  normal: { [key: string]: string };

  /**
   * secret is a map of sensitive variable names to values
   *
   * @generated from field: map<string, string> secret = 2;
   */
  secret: { [key: string]: string };

  constructor(data?: PartialMessage<EnvGroupVariables>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.EnvGroupVariables";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EnvGroupVariables;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EnvGroupVariables;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EnvGroupVariables;

  static equals(a: EnvGroupVariables | PlainMessage<EnvGroupVariables> | undefined, b: EnvGroupVariables | PlainMessage<EnvGroupVariables> | undefined): boolean;
}

/**
 * Deletions contains all explicit deletions from a PorterApp
 *
 * @generated from message porter.v1.Deletions
 */
export declare class Deletions extends Message<Deletions> {
  /**
   * service_names is a list of service names to delete
   *
   * @generated from field: repeated string service_names = 1;
   */
  serviceNames: string[];

  /**
   * env_group_names is a list of environment variable group names to delete
   *
   * @generated from field: repeated string env_group_names = 2;
   */
  envGroupNames: string[];

  /**
   * env_variable_names is deprecated in favor of env_group_names. It was a list of environment variable names to delete
   *
   * @generated from field: repeated string env_variable_names = 4 [deprecated = true];
   * @deprecated
   */
  envVariableNames: string[];

  constructor(data?: PartialMessage<Deletions>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Deletions";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Deletions;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Deletions;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Deletions;

  static equals(a: Deletions | PlainMessage<Deletions> | undefined, b: Deletions | PlainMessage<Deletions> | undefined): boolean;
}

/**
 * Build is the build settings for the application
 *
 * @generated from message porter.v1.Build
 */
export declare class Build extends Message<Build> {
  /**
   * context is the path to the build context
   *
   * @generated from field: string context = 1;
   */
  context: string;

  /**
   * method is the build method to use, being one of "pack", "docker", or "registry"
   *
   * @generated from field: string method = 2;
   */
  method: string;

  /**
   * builder is the builder to use for the "pack" build method
   *
   * @generated from field: string builder = 3;
   */
  builder: string;

  /**
   * buildpacks is a list of buildpacks to use for the "pack" build method
   *
   * @generated from field: repeated string buildpacks = 4;
   */
  buildpacks: string[];

  /**
   * dockerfile is the path to the Dockerfile to use for the "docker" build method
   *
   * @generated from field: string dockerfile = 5;
   */
  dockerfile: string;

  /**
   * commit_sha is the commit SHA at which to build the application
   *
   * @generated from field: string commit_sha = 6;
   */
  commitSha: string;

  constructor(data?: PartialMessage<Build>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Build";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Build;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Build;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Build;

  static equals(a: Build | PlainMessage<Build> | undefined, b: Build | PlainMessage<Build> | undefined): boolean;
}

/**
 * AppImage is the image to use for a given revision of the application
 *
 * @generated from message porter.v1.AppImage
 */
export declare class AppImage extends Message<AppImage> {
  /**
   * repository is the repository to use for the image
   *
   * @generated from field: string repository = 1;
   */
  repository: string;

  /**
   * tag is the tag to use for the image
   *
   * @generated from field: string tag = 2;
   */
  tag: string;

  constructor(data?: PartialMessage<AppImage>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.AppImage";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AppImage;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AppImage;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AppImage;

  static equals(a: AppImage | PlainMessage<AppImage> | undefined, b: AppImage | PlainMessage<AppImage> | undefined): boolean;
}

