// @generated by protoc-gen-es v1.6.0
// @generated from file porter/v1/contract.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { Cluster } from "./cluster_pb.js";

/**
 * Contract represents a contract for managing clusters and compliance
 *
 * @generated from message porter.v1.Contract
 */
export const Contract = proto3.makeMessageType(
  "porter.v1.Contract",
  () => [
    { no: 1, name: "cluster", kind: "message", T: Cluster },
    { no: 2, name: "user", kind: "message", T: User },
    { no: 4, name: "compliance_profiles", kind: "message", T: ComplianceProfile },
  ],
);

/**
 * ContractRevision represents a contract revision which should be acted upon, depending on the stream
 * that is was passed through. This approach goes against microservice architectures, by expecting every consumer
 * to read the specific contract from the database
 *
 * @generated from message porter.v1.ContractRevision
 */
export const ContractRevision = proto3.makeMessageType(
  "porter.v1.ContractRevision",
  () => [
    { no: 1, name: "cluster_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "project_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "revision_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * User represents a user of Porter, who actioned the change to the contract
 *
 * @generated from message porter.v1.User
 */
export const User = proto3.makeMessageType(
  "porter.v1.User",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * ComplianceProfiles are the different compliance profiles that can be enforced on a contract
 *
 * @generated from message porter.v1.ComplianceProfile
 */
export const ComplianceProfile = proto3.makeMessageType(
  "porter.v1.ComplianceProfile",
  () => [
    { no: 1, name: "soc2", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "hipaa", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

