// @generated by protoc-gen-es v1.4.2
// @generated from file porter/v1/env_group.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * EnvGroupProviderType is the type of the EnvGroupProvider.
 *
 * @generated from enum porter.v1.EnumEnvGroupProviderType
 */
export const EnumEnvGroupProviderType = proto3.makeEnum(
  "porter.v1.EnumEnvGroupProviderType",
  [
    {no: 0, name: "ENUM_ENV_GROUP_PROVIDER_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "ENUM_ENV_GROUP_PROVIDER_TYPE_PORTER", localName: "PORTER"},
    {no: 2, name: "ENUM_ENV_GROUP_PROVIDER_TYPE_DOPPLER", localName: "DOPPLER"},
    {no: 3, name: "ENUM_ENV_GROUP_PROVIDER_TYPE_DATASTORE", localName: "DATASTORE"},
  ],
);

/**
 * @generated from message porter.v1.UpdateAppsLinkedToEnvGroupPayload
 */
export const UpdateAppsLinkedToEnvGroupPayload = proto3.makeMessageType(
  "porter.v1.UpdateAppsLinkedToEnvGroupPayload",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "cluster_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "env_group_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "remove_env_group", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

