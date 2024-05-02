// @generated by protoc-gen-es v1.9.0
// @generated from file porter/v1/deployment_target.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * DeploymentTargetIdentifier is the object that identifies a deployment target. One of id or name must be provided, with id taking precedence.
 *
 * @generated from message porter.v1.DeploymentTargetIdentifier
 */
export const DeploymentTargetIdentifier = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.DeploymentTargetIdentifier",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.DeploymentTarget
 */
export const DeploymentTarget = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.DeploymentTarget",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "namespace", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "cluster_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 5, name: "is_preview", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 6, name: "is_default", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 7, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "metadata", kind: "message", T: DeploymentTargetMeta },
  ],
);

/**
 * @generated from message porter.v1.DeploymentTargetMeta
 */
export const DeploymentTargetMeta = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.DeploymentTargetMeta",
  () => [
    { no: 1, name: "pull_request", kind: "message", T: PullRequest },
  ],
);

/**
 * @generated from message porter.v1.PullRequest
 */
export const PullRequest = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.PullRequest",
  () => [
    { no: 1, name: "repository", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "number", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "head_ref", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

