// @generated by protoc-gen-es v1.4.2
// @generated from file porter/v1/contract.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Cluster } from "./cluster_pb.js";

/**
 * Contract represents a contract for managing clusters and compliance
 *
 * @generated from message porter.v1.Contract
 */
export declare class Contract extends Message<Contract> {
  /**
   * @generated from field: porter.v1.Cluster cluster = 1;
   */
  cluster?: Cluster;

  /**
   * @generated from field: porter.v1.User user = 2;
   */
  user?: User;

  /**
   * compliance_profiles is a list of compliance profiles that should be enforced on the contract
   *
   * @generated from field: porter.v1.ComplianceProfile compliance_profiles = 4;
   */
  complianceProfiles?: ComplianceProfile;

  constructor(data?: PartialMessage<Contract>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Contract";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Contract;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Contract;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Contract;

  static equals(a: Contract | PlainMessage<Contract> | undefined, b: Contract | PlainMessage<Contract> | undefined): boolean;
}

/**
 * ContractRevision represents a contract revision which should be acted upon, depending on the stream
 * that is was passed through. This approach goes against microservice architectures, by expecting every consumer
 * to read the specific contract from the database
 *
 * @generated from message porter.v1.ContractRevision
 */
export declare class ContractRevision extends Message<ContractRevision> {
  /**
   * @generated from field: int32 cluster_id = 1;
   */
  clusterId: number;

  /**
   * @generated from field: int32 project_id = 2;
   */
  projectId: number;

  /**
   * revision_id is the ID of the contract revision that this message applies to.
   * This field is a UUID represented as a string, for better compatibility.
   * Best practice is to parse this as a uuid upon receipt
   *
   * @generated from field: string revision_id = 3;
   */
  revisionId: string;

  constructor(data?: PartialMessage<ContractRevision>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.ContractRevision";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ContractRevision;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ContractRevision;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ContractRevision;

  static equals(a: ContractRevision | PlainMessage<ContractRevision> | undefined, b: ContractRevision | PlainMessage<ContractRevision> | undefined): boolean;
}

/**
 * User represents a user of Porter, who actioned the change to the contract
 *
 * @generated from message porter.v1.User
 */
export declare class User extends Message<User> {
  /**
   * @generated from field: int32 id = 1;
   */
  id: number;

  constructor(data?: PartialMessage<User>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.User";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): User;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): User;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): User;

  static equals(a: User | PlainMessage<User> | undefined, b: User | PlainMessage<User> | undefined): boolean;
}

/**
 * ComplianceProfiles are the different compliance profiles that can be enforced on a contract
 *
 * @generated from message porter.v1.ComplianceProfile
 */
export declare class ComplianceProfile extends Message<ComplianceProfile> {
  /**
   * soc2 indicates that the contract should be compliant with SOC2
   *
   * @generated from field: bool soc2 = 1;
   */
  soc2: boolean;

  /**
   * hipaa indicates that the contract should be compliant with HIPAA
   *
   * @generated from field: bool hipaa = 2;
   */
  hipaa: boolean;

  constructor(data?: PartialMessage<ComplianceProfile>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.ComplianceProfile";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ComplianceProfile;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ComplianceProfile;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ComplianceProfile;

  static equals(a: ComplianceProfile | PlainMessage<ComplianceProfile> | undefined, b: ComplianceProfile | PlainMessage<ComplianceProfile> | undefined): boolean;
}

