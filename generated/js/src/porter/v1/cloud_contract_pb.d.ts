// @generated by protoc-gen-es v1.8.0
// @generated from file porter/v1/cloud_contract.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { ManagedDatastore } from "./datastore_pb.js";
import type { Addon } from "./addons_pb.js";

/**
 * CloudContract is a contract for all Porter-managed infrastructure within a project
 *
 * @generated from message porter.v1.CloudContract
 */
export declare class CloudContract extends Message<CloudContract> {
  /**
   * datastores is the list of datastores associated with the project
   *
   * @generated from field: repeated porter.v1.ManagedDatastore datastores = 1;
   */
  datastores: ManagedDatastore[];

  /**
   * addons is the list of addons associated with the project
   *
   * @generated from field: repeated porter.v1.CloudContractAddon addons = 2;
   */
  addons: CloudContractAddon[];

  constructor(data?: PartialMessage<CloudContract>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.CloudContract";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CloudContract;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CloudContract;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CloudContract;

  static equals(a: CloudContract | PlainMessage<CloudContract> | undefined, b: CloudContract | PlainMessage<CloudContract> | undefined): boolean;
}

/**
 * CloudContractAddon is a contract for an addon managed by the cloud contract
 *
 * @generated from message porter.v1.CloudContractAddon
 */
export declare class CloudContractAddon extends Message<CloudContractAddon> {
  /**
   * cluster_id is the id of the cluster that the addon is deployed on
   *
   * @generated from field: int32 cluster_id = 1;
   */
  clusterId: number;

  /**
   * addon is the addon configuration
   *
   * @generated from field: porter.v1.Addon addon = 2;
   */
  addon?: Addon;

  constructor(data?: PartialMessage<CloudContractAddon>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.CloudContractAddon";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CloudContractAddon;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CloudContractAddon;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CloudContractAddon;

  static equals(a: CloudContractAddon | PlainMessage<CloudContractAddon> | undefined, b: CloudContractAddon | PlainMessage<CloudContractAddon> | undefined): boolean;
}

/**
 * CloudContractRevision represents a cloud contract revision which should be reconciled
 *
 * @generated from message porter.v1.CloudContractRevision
 */
export declare class CloudContractRevision extends Message<CloudContractRevision> {
  /**
   * @generated from field: int32 project_id = 1;
   */
  projectId: number;

  /**
   * revision_id is the id of the cloud contract revision that this message applies to
   *
   * @generated from field: string revision_id = 2;
   */
  revisionId: string;

  constructor(data?: PartialMessage<CloudContractRevision>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.CloudContractRevision";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CloudContractRevision;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CloudContractRevision;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CloudContractRevision;

  static equals(a: CloudContractRevision | PlainMessage<CloudContractRevision> | undefined, b: CloudContractRevision | PlainMessage<CloudContractRevision> | undefined): boolean;
}

/**
 * CloudContractDeletionRevision represents a collection of resources that should be deleted
 *
 * @generated from message porter.v1.CloudContractDeletionRevision
 */
export declare class CloudContractDeletionRevision extends Message<CloudContractDeletionRevision> {
  /**
   * @generated from field: int32 project_id = 1;
   */
  projectId: number;

  /**
   * cloud_contract_deletions includes all resources that should be deleted
   *
   * @generated from field: porter.v1.CloudContract cloud_contract_deletions = 2;
   */
  cloudContractDeletions?: CloudContract;

  constructor(data?: PartialMessage<CloudContractDeletionRevision>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.CloudContractDeletionRevision";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CloudContractDeletionRevision;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CloudContractDeletionRevision;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CloudContractDeletionRevision;

  static equals(a: CloudContractDeletionRevision | PlainMessage<CloudContractDeletionRevision> | undefined, b: CloudContractDeletionRevision | PlainMessage<CloudContractDeletionRevision> | undefined): boolean;
}

/**
 * MachineType is a virtual machine type
 *
 * @generated from message porter.v1.MachineType
 */
export declare class MachineType extends Message<MachineType> {
  /**
   * name is the name of the machine type
   *
   * @generated from field: string name = 1;
   */
  name: string;

  constructor(data?: PartialMessage<MachineType>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.MachineType";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MachineType;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MachineType;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MachineType;

  static equals(a: MachineType | PlainMessage<MachineType> | undefined, b: MachineType | PlainMessage<MachineType> | undefined): boolean;
}

