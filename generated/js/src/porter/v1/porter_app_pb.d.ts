// @generated by protoc-gen-es v1.4.1
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
   * @generated from field: map<string, porter.v1.Service> services = 2 [deprecated = true];
   * @deprecated
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

  /**
   * helm_overrides is stringified json of raw Helm overrides to use for the application. These will override any values generated by the contract.
   *
   * @generated from field: porter.v1.HelmOverrides helm_overrides = 8;
   */
  helmOverrides?: HelmOverrides;

  /**
   * service_list is a list of service configurations
   *
   * @generated from field: repeated porter.v1.Service service_list = 9;
   */
  serviceList: Service[];

  /**
   * efs_storage shared storage across service
   *
   * @generated from field: porter.v1.EFS efs_storage = 10;
   */
  efsStorage?: EFS;

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
 * HelmOverrides is a wrapper over stringified json of raw Helm overrides to use for the application. HelmOverrides will only be removed if values is explicitly set to an empty string.
 *
 * @generated from message porter.v1.HelmOverrides
 */
export declare class HelmOverrides extends Message<HelmOverrides> {
  /**
   * Values is base64-encoded, stringified json of Helm values. We use a string because Typescript does not have a native bytes type.
   *
   * @generated from field: string b64_values = 1;
   */
  b64Values: string;

  constructor(data?: PartialMessage<HelmOverrides>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.HelmOverrides";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): HelmOverrides;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): HelmOverrides;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): HelmOverrides;

  static equals(a: HelmOverrides | PlainMessage<HelmOverrides> | undefined, b: HelmOverrides | PlainMessage<HelmOverrides> | undefined): boolean;
}

/**
 * EnvGroup represents the metadata for an env group. We do not want to store the actual variables with the PorterApp.
 *
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
 * EnvGroupVariables represents the variables for an env group.
 *
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
 * ServiceDeletions contains all explicit deletions from a service
 *
 * @generated from message porter.v1.ServiceDeletions
 */
export declare class ServiceDeletions extends Message<ServiceDeletions> {
  /**
   * domain_names is a list of domain names to delete
   *
   * @generated from field: repeated string domain_names = 1;
   */
  domainNames: string[];

  /**
   * ingress_annotations is a list of ingress annotation keys to delete
   *
   * @generated from field: repeated string ingress_annotations = 2;
   */
  ingressAnnotations: string[];

  constructor(data?: PartialMessage<ServiceDeletions>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.ServiceDeletions";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ServiceDeletions;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ServiceDeletions;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ServiceDeletions;

  static equals(a: ServiceDeletions | PlainMessage<ServiceDeletions> | undefined, b: ServiceDeletions | PlainMessage<ServiceDeletions> | undefined): boolean;
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
   * predeploy_names is a list of predeploy jobs to delete. Even though there is only one predeploy job, it is a list to be consistent with the service deletions.
   *
   * @generated from field: repeated string predeploy_names = 3;
   */
  predeployNames: string[];

  /**
   * env_variable_names is deprecated in favor of env_group_names. It was a list of environment variable names to delete
   *
   * @generated from field: repeated string env_variable_names = 4 [deprecated = true];
   * @deprecated
   */
  envVariableNames: string[];

  /**
   * service_domains is a map of service names to a list of domains to delete
   *
   * @generated from field: map<string, porter.v1.DomainNameList> service_domains = 5 [deprecated = true];
   * @deprecated
   */
  serviceDomains: { [key: string]: DomainNameList };

  /**
   * service_deletions is a map of service names to a list of deletions to perform on the service
   *
   * @generated from field: map<string, porter.v1.ServiceDeletions> service_deletions = 6;
   */
  serviceDeletions: { [key: string]: ServiceDeletions };

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
 * @generated from message porter.v1.DomainNameList
 */
export declare class DomainNameList extends Message<DomainNameList> {
  /**
   * domain_names is a list of domain names
   *
   * @generated from field: repeated string domain_names = 1;
   */
  domainNames: string[];

  constructor(data?: PartialMessage<DomainNameList>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.DomainNameList";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DomainNameList;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DomainNameList;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DomainNameList;

  static equals(a: DomainNameList | PlainMessage<DomainNameList> | undefined, b: DomainNameList | PlainMessage<DomainNameList> | undefined): boolean;
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

/**
 * EFS is the values to configure EFS settings
 *
 * @generated from message porter.v1.EFS
 */
export declare class EFS extends Message<EFS> {
  /**
   * enabled enabes shared storage across services
   *
   * @generated from field: bool enabled = 1;
   */
  enabled: boolean;

  constructor(data?: PartialMessage<EFS>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.EFS";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EFS;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EFS;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EFS;

  static equals(a: EFS | PlainMessage<EFS> | undefined, b: EFS | PlainMessage<EFS> | undefined): boolean;
}

