// @generated by protoc-gen-es v1.10.0
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
    {no: 4, name: "ENUM_ENV_GROUP_PROVIDER_TYPE_INFISICAL", localName: "INFISICAL"},
  ],
);

/**
 * EnumExternalEnvGroupOperatorType is the type of external env group provider.
 *
 * @generated from enum porter.v1.EnumExternalEnvGroupOperatorType
 */
export const EnumExternalEnvGroupOperatorType = /*@__PURE__*/ proto3.makeEnum(
  "porter.v1.EnumExternalEnvGroupOperatorType",
  [
    {no: 0, name: "ENUM_EXTERNAL_ENV_GROUP_OPERATOR_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "ENUM_EXTERNAL_ENV_GROUP_OPERATOR_TYPE_EXTERNAL_SECRETS", localName: "EXTERNAL_SECRETS"},
    {no: 2, name: "ENUM_EXTERNAL_ENV_GROUP_OPERATOR_TYPE_INFISICAL", localName: "INFISICAL"},
  ],
);

/**
 * @generated from message porter.v1.ExternalEnvGroupProviderEnabledStatus
 */
export const ExternalEnvGroupProviderEnabledStatus = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.ExternalEnvGroupProviderEnabledStatus",
  () => [
    { no: 1, name: "operator", kind: "enum", T: proto3.getEnumType(EnumExternalEnvGroupOperatorType) },
    { no: 2, name: "enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "reprovision_required", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 4, name: "k8s_upgrade_required", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * @generated from message porter.v1.InfisicalEnv
 */
export const InfisicalEnv = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.InfisicalEnv",
  () => [
    { no: 1, name: "environment_slug", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "environment_path", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

