// @generated by protoc-gen-es v1.4.2
// @generated from file porter/v1/aws_assume_role.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * AssumeRoleChainLink is used for hopping to a specific AWS role.
 * Deprecated - use AssumeRoleHop instead
 *
 * @generated from message porter.v1.AssumeRoleChainLink
 * @deprecated
 */
export const AssumeRoleChainLink = proto3.makeMessageType(
  "porter.v1.AssumeRoleChainLink",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "source_arn", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "target_arn", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "external_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

