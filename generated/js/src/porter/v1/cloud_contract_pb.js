// @generated by protoc-gen-es v1.8.0
// @generated from file porter/v1/cloud_contract.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { ManagedDatastore } from "./datastore_pb.js";
import { Addon } from "./addons_pb.js";

/**
 * CloudContract is a contract for all Porter-managed infrastructure within a project
 *
 * @generated from message porter.v1.CloudContract
 */
export const CloudContract = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.CloudContract",
  () => [
    { no: 1, name: "datastores", kind: "message", T: ManagedDatastore, repeated: true },
    { no: 2, name: "addons", kind: "message", T: CloudContractAddon, repeated: true },
  ],
);

/**
 * CloudContractAddon is a contract for an addon managed by the cloud contract
 *
 * @generated from message porter.v1.CloudContractAddon
 */
export const CloudContractAddon = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.CloudContractAddon",
  () => [
    { no: 1, name: "cluster_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "addon", kind: "message", T: Addon },
  ],
);

/**
 * CloudContractRevision represents a cloud contract revision which should be reconciled
 *
 * @generated from message porter.v1.CloudContractRevision
 */
export const CloudContractRevision = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.CloudContractRevision",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "revision_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * CloudContractDeletionRevision represents a collection of resources that should be deleted
 *
 * @generated from message porter.v1.CloudContractDeletionRevision
 */
export const CloudContractDeletionRevision = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.CloudContractDeletionRevision",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "cloud_contract_deletions", kind: "message", T: CloudContract },
  ],
);

