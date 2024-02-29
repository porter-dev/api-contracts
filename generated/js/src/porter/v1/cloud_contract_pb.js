// @generated by protoc-gen-es v1.4.2
// @generated from file porter/v1/cloud_contract.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { ManagedDatastore } from "./datastore_pb.js";

/**
 * CloudContract is a contract for all Porter-managed infrastructure within a project
 *
 * @generated from message porter.v1.CloudContract
 */
export const CloudContract = proto3.makeMessageType(
  "porter.v1.CloudContract",
  () => [
    { no: 1, name: "datastores", kind: "message", T: ManagedDatastore, repeated: true },
  ],
);

