// @generated by protoc-gen-es v1.9.0
// @generated from file porter/v1/auth_management.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message porter.v1.APITokenRequest
 */
export const APITokenRequest = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.APITokenRequest",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * @generated from message porter.v1.APITokenResponse
 */
export const APITokenResponse = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.APITokenResponse",
  () => [
    { no: 1, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

