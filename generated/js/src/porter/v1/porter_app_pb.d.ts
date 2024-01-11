// @generated by protoc-gen-es v1.4.2
// @generated from file porter/v1/porter_app.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Service } from "./service_pb.js";

/**
 * EnvVariableSource is the source type from which to pull the environment variable
 *
 * @generated from enum porter.v1.EnvVariableSource
 */
export declare enum EnvVariableSource {
  /**
   * @generated from enum value: ENV_VARIABLE_SOURCE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: ENV_VARIABLE_SOURCE_VALUE = 1;
   */
  VALUE = 1,

  /**
   * @generated from enum value: ENV_VARIABLE_SOURCE_FROM_APP = 2;
   */
  FROM_APP = 2,
}

/**
 * EnvValueFromApp is an enum of the possible values to use in the EnvVariableFromApp definition
 *
 * @generated from enum porter.v1.EnvValueFromApp
 */
export declare enum EnvValueFromApp {
  /**
   * @generated from enum value: ENV_VALUE_FROM_APP_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: ENV_VALUE_FROM_APP_PUBLIC_DOMAIN = 1;
   */
  PUBLIC_DOMAIN = 1,

  /**
   * @generated from enum value: ENV_VALUE_FROM_APP_INTERNAL_DOMAIN = 2;
   */
  INTERNAL_DOMAIN = 2,
}

/**
 * EnumAppRevisionStatus describes the status of an app revision
 *
 * @generated from enum porter.v1.EnumAppRevisionStatus
 */
export declare enum EnumAppRevisionStatus {
  /**
   * @generated from enum value: ENUM_APP_REVISION_STATUS_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: ENUM_APP_REVISION_STATUS_PROGRESSING = 1;
   */
  PROGRESSING = 1,

  /**
   * @generated from enum value: ENUM_APP_REVISION_STATUS_SUCCESSFUL = 2;
   */
  SUCCESSFUL = 2,

  /**
   * @generated from enum value: ENUM_APP_REVISION_STATUS_FAILED = 3;
   */
  FAILED = 3,
}

/**
 * DeploymentTargetIdentifier is the object that identifies a deployment target. One of id or name must be provided, with id taking precedence.
 *
 * @generated from message porter.v1.DeploymentTargetIdentifier
 */
export declare class DeploymentTargetIdentifier extends Message<DeploymentTargetIdentifier> {
  /**
   * id is the id of the deployment target
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * name is the name of the deployment target
   *
   * @generated from field: string name = 2;
   */
  name: string;

  constructor(data?: PartialMessage<DeploymentTargetIdentifier>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.DeploymentTargetIdentifier";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeploymentTargetIdentifier;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeploymentTargetIdentifier;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeploymentTargetIdentifier;

  static equals(a: DeploymentTargetIdentifier | PlainMessage<DeploymentTargetIdentifier> | undefined, b: DeploymentTargetIdentifier | PlainMessage<DeploymentTargetIdentifier> | undefined): boolean;
}

/**
 * @generated from message porter.v1.DeploymentTarget
 */
export declare class DeploymentTarget extends Message<DeploymentTarget> {
  /**
   * @generated from field: int64 project_id = 1;
   */
  projectId: bigint;

  /**
   * name is the vanity name for the deployment target
   *
   * @generated from field: string name = 2;
   */
  name: string;

  /**
   * namespace is the namespace that the deployment target points to
   *
   * @generated from field: string namespace = 3;
   */
  namespace: string;

  /**
   * cluster_id is the id of the cluster that the deployment target points to
   *
   * @generated from field: int64 cluster_id = 4;
   */
  clusterId: bigint;

  /**
   * is_preview indicates whether this is a preview deployment target or not
   *
   * @generated from field: bool is_preview = 5;
   */
  isPreview: boolean;

  /**
   * is_default indicates whether this is the default deployment target for the cluster
   *
   * @generated from field: bool is_default = 6;
   */
  isDefault: boolean;

  /**
   * id is the id of the deployment target
   *
   * @generated from field: string id = 7;
   */
  id: string;

  constructor(data?: PartialMessage<DeploymentTarget>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.DeploymentTarget";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeploymentTarget;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeploymentTarget;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeploymentTarget;

  static equals(a: DeploymentTarget | PlainMessage<DeploymentTarget> | undefined, b: DeploymentTarget | PlainMessage<DeploymentTarget> | undefined): boolean;
}

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
   * env is a list of environment variables to set on the application
   *
   * @generated from field: repeated porter.v1.EnvVariable env = 3;
   */
  env: EnvVariable[];

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

  /**
   * required_apps are other porter apps that this app expects to be deployed alongside it
   *
   * @generated from field: repeated porter.v1.RequiredApp required_apps = 11;
   */
  requiredApps: RequiredApp[];

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
 * RequiredApp specifies another porter app that this app expects to be deployed alongside it
 * These are used for preview environments to pull in the latest deployed version of an app from a different deployment target
 *
 * @generated from message porter.v1.RequiredApp
 */
export declare class RequiredApp extends Message<RequiredApp> {
  /**
   * name is the name of the required app
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * from_target is the deployment target from which to pull the contract for the required app
   * If not provided, the contract will be pulled from the default deployment target
   *
   * @generated from field: porter.v1.DeploymentTargetIdentifier from_target = 2;
   */
  fromTarget?: DeploymentTargetIdentifier;

  constructor(data?: PartialMessage<RequiredApp>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.RequiredApp";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RequiredApp;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RequiredApp;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RequiredApp;

  static equals(a: RequiredApp | PlainMessage<RequiredApp> | undefined, b: RequiredApp | PlainMessage<RequiredApp> | undefined): boolean;
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

  /**
   * linked_applications is an optional list of applications linked to the env group
   *
   * @generated from field: repeated string linked_applications = 3;
   */
  linkedApplications: string[];

  /**
   * variables is an optional map of non-sensitive variable names to values
   *
   * @generated from field: map<string, string> variables = 4;
   */
  variables: { [key: string]: string };

  /**
   * secret is an optional map of sensitive variable names to values
   *
   * @generated from field: map<string, string> secret_variables = 5;
   */
  secretVariables: { [key: string]: string };

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

  /**
   * file_system_id is the AWS EFS File System ID to mount
   *
   * @generated from field: string file_system_id = 2;
   */
  fileSystemId: string;

  constructor(data?: PartialMessage<EFS>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.EFS";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EFS;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EFS;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EFS;

  static equals(a: EFS | PlainMessage<EFS> | undefined, b: EFS | PlainMessage<EFS> | undefined): boolean;
}

/**
 * @generated from message porter.v1.EnvVariable
 */
export declare class EnvVariable extends Message<EnvVariable> {
  /**
   * key is the name of the environment variable
   *
   * @generated from field: string key = 1;
   */
  key: string;

  /**
   * source specifies how the env variable should be set
   *
   * @generated from field: porter.v1.EnvVariableSource source = 2;
   */
  source: EnvVariableSource;

  /**
   * value is the value of the environment variable. This is only used if case is ENV_VARIABLE_SOURCE_VALUE
   *
   * @generated from field: string value = 3;
   */
  value: string;

  /**
   * definition is an object that defines how the environment variable should be set
   *
   * @generated from oneof porter.v1.EnvVariable.definition
   */
  definition: {
    /**
     * @generated from field: porter.v1.EnvVariableFromApp from_app = 4;
     */
    value: EnvVariableFromApp;
    case: "fromApp";
  } | { case: undefined; value?: undefined };

  constructor(data?: PartialMessage<EnvVariable>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.EnvVariable";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EnvVariable;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EnvVariable;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EnvVariable;

  static equals(a: EnvVariable | PlainMessage<EnvVariable> | undefined, b: EnvVariable | PlainMessage<EnvVariable> | undefined): boolean;
}

/**
 * @generated from message porter.v1.EnvVariableFromApp
 */
export declare class EnvVariableFromApp extends Message<EnvVariableFromApp> {
  /**
   * app_name is the name of the app
   *
   * @generated from field: string app_name = 1;
   */
  appName: string;

  /**
   * value is one of a discrete set of configured values on the app that can be used for the environment variable
   *
   * @generated from field: porter.v1.EnvValueFromApp value = 2;
   */
  value: EnvValueFromApp;

  /**
   * service_name is the name of the service where the value is defined
   *
   * @generated from field: string service_name = 3;
   */
  serviceName: string;

  constructor(data?: PartialMessage<EnvVariableFromApp>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.EnvVariableFromApp";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EnvVariableFromApp;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EnvVariableFromApp;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EnvVariableFromApp;

  static equals(a: EnvVariableFromApp | PlainMessage<EnvVariableFromApp> | undefined, b: EnvVariableFromApp | PlainMessage<EnvVariableFromApp> | undefined): boolean;
}

