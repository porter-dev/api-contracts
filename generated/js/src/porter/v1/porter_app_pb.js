// @generated by protoc-gen-es v1.4.2
// @generated from file porter/v1/porter_app.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { Service } from "./service_pb.js";

/**
 * DeploymentTargetIdentifier is the object that identifies a deployment target. One of id or name must be provided, with id taking precedence.
 *
 * @generated from message porter.v1.DeploymentTargetIdentifier
 */
export const DeploymentTargetIdentifier = proto3.makeMessageType(
  "porter.v1.DeploymentTargetIdentifier",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.DeploymentTarget
 */
export const DeploymentTarget = proto3.makeMessageType(
  "porter.v1.DeploymentTarget",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "namespace", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "cluster_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 5, name: "is_preview", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 6, name: "is_default", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 7, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * PorterApp is the top-level configuration for a Porter application, usually found in porter.yaml
 *
 * @generated from message porter.v1.PorterApp
 */
export const PorterApp = proto3.makeMessageType(
  "porter.v1.PorterApp",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "services", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: Service} },
    { no: 3, name: "env", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 4, name: "build", kind: "message", T: Build },
    { no: 5, name: "predeploy", kind: "message", T: Service },
    { no: 6, name: "image", kind: "message", T: AppImage },
    { no: 7, name: "env_groups", kind: "message", T: EnvGroup, repeated: true },
    { no: 8, name: "helm_overrides", kind: "message", T: HelmOverrides },
    { no: 9, name: "service_list", kind: "message", T: Service, repeated: true },
    { no: 10, name: "efs_storage", kind: "message", T: EFS },
    { no: 11, name: "required_apps", kind: "message", T: RequiredApp, repeated: true },
  ],
);

/**
 * RequiredApp specifies another porter app that this app expects to be deployed alongside it
 * These are used for preview environments to pull in the latest deployed version of an app from a different deployment target
 *
 * @generated from message porter.v1.RequiredApp
 */
export const RequiredApp = proto3.makeMessageType(
  "porter.v1.RequiredApp",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "from_target", kind: "message", T: DeploymentTargetIdentifier },
  ],
);

/**
 * HelmOverrides is a wrapper over stringified json of raw Helm overrides to use for the application. HelmOverrides will only be removed if values is explicitly set to an empty string.
 *
 * @generated from message porter.v1.HelmOverrides
 */
export const HelmOverrides = proto3.makeMessageType(
  "porter.v1.HelmOverrides",
  () => [
    { no: 1, name: "b64_values", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * EnvGroup represents the metadata for an env group. We do not want to store the actual variables with the PorterApp.
 *
 * @generated from message porter.v1.EnvGroup
 */
export const EnvGroup = proto3.makeMessageType(
  "porter.v1.EnvGroup",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "version", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * EnvGroupVariables represents the variables for an env group.
 *
 * @generated from message porter.v1.EnvGroupVariables
 */
export const EnvGroupVariables = proto3.makeMessageType(
  "porter.v1.EnvGroupVariables",
  () => [
    { no: 1, name: "normal", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 2, name: "secret", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
  ],
);

/**
 * ServiceDeletions contains all explicit deletions from a service
 *
 * @generated from message porter.v1.ServiceDeletions
 */
export const ServiceDeletions = proto3.makeMessageType(
  "porter.v1.ServiceDeletions",
  () => [
    { no: 1, name: "domain_names", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 2, name: "ingress_annotations", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * Deletions contains all explicit deletions from a PorterApp
 *
 * @generated from message porter.v1.Deletions
 */
export const Deletions = proto3.makeMessageType(
  "porter.v1.Deletions",
  () => [
    { no: 1, name: "service_names", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 2, name: "env_group_names", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 3, name: "predeploy_names", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "env_variable_names", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 5, name: "service_domains", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: DomainNameList} },
    { no: 6, name: "service_deletions", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: ServiceDeletions} },
  ],
);

/**
 * @generated from message porter.v1.DomainNameList
 */
export const DomainNameList = proto3.makeMessageType(
  "porter.v1.DomainNameList",
  () => [
    { no: 1, name: "domain_names", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * Build is the build settings for the application
 *
 * @generated from message porter.v1.Build
 */
export const Build = proto3.makeMessageType(
  "porter.v1.Build",
  () => [
    { no: 1, name: "context", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "method", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "builder", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "buildpacks", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 5, name: "dockerfile", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "commit_sha", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * AppImage is the image to use for a given revision of the application
 *
 * @generated from message porter.v1.AppImage
 */
export const AppImage = proto3.makeMessageType(
  "porter.v1.AppImage",
  () => [
    { no: 1, name: "repository", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "tag", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * EFS is the values to configure EFS settings
 *
 * @generated from message porter.v1.EFS
 */
export const EFS = proto3.makeMessageType(
  "porter.v1.EFS",
  () => [
    { no: 1, name: "enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "file_system_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

