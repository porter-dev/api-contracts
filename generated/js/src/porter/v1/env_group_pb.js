// @generated by protoc-gen-es v1.8.0
// @generated from file porter/v1/env_group.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * EnvGroupProviderType is the type of the EnvGroupProvider.
 *
 * @generated from enum porter.v1.EnumEnvGroupProviderType
 */
export const EnumEnvGroupProviderType = /*@__PURE__*/ proto3.makeEnum(
  "porter.v1.EnumEnvGroupProviderType",
  [
    {no: 0, name: "ENUM_ENV_GROUP_PROVIDER_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "ENUM_ENV_GROUP_PROVIDER_TYPE_PORTER", localName: "PORTER"},
    {no: 2, name: "ENUM_ENV_GROUP_PROVIDER_TYPE_DOPPLER", localName: "DOPPLER"},
    {no: 3, name: "ENUM_ENV_GROUP_PROVIDER_TYPE_DATASTORE", localName: "DATASTORE"},
  ],
);

