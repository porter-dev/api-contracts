// @generated by protoc-gen-es v1.10.0
// @generated from file porter/v1/dead_letter.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { ContractRevision } from "./contract_pb.js";
import { Error } from "./errors_pb.js";

/**
 * @generated from message porter.v1.DeadLetter
 */
export const DeadLetter = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.DeadLetter",
  () => [
    { no: 1, name: "contract_revision", kind: "message", T: ContractRevision },
    { no: 2, name: "error", kind: "message", T: Error },
  ],
);

