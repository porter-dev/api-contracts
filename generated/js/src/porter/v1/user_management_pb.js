// @generated by protoc-gen-es v1.2.0
// @generated from file porter/v1/user_management.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message porter.v1.PorterAgentConnectionRequest
 */
export const PorterAgentConnectionRequest = proto3.makeMessageType(
  "porter.v1.PorterAgentConnectionRequest",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "cluster_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * @generated from message porter.v1.PorterAgentConnectionResponse
 */
export const PorterAgentConnectionResponse = proto3.makeMessageType(
  "porter.v1.PorterAgentConnectionResponse",
  () => [
    { no: 1, name: "server_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "port", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

