// @generated by protoc-gen-es v1.1.0
// @generated from file porter/v1/errors.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * Error is a lightweight error message contract
 *
 * @generated from message porter.v1.Error
 */
export const Error = proto3.makeMessageType(
  "porter.v1.Error",
  () => [
    { no: 1, name: "code", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "metadata", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
  ],
);

