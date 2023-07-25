// @generated by protoc-gen-es v1.2.0
// @generated from file porter/v1/aws_assume_role.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * AssumeRoleChainLink is used for hopping to a specific AWS role.
 * Deprecated - use AssumeRoleHop instead
 *
 * @generated from message porter.v1.AssumeRoleChainLink
 * @deprecated
 */
export declare class AssumeRoleChainLink extends Message<AssumeRoleChainLink> {
  /**
   * @generated from field: int64 project_id = 1;
   */
  projectId: bigint;

  /**
   * @generated from field: string source_arn = 2;
   */
  sourceArn: string;

  /**
   * @generated from field: string target_arn = 3;
   */
  targetArn: string;

  /**
   * @generated from field: string external_id = 4;
   */
  externalId: string;

  constructor(data?: PartialMessage<AssumeRoleChainLink>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.AssumeRoleChainLink";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AssumeRoleChainLink;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AssumeRoleChainLink;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AssumeRoleChainLink;

  static equals(a: AssumeRoleChainLink | PlainMessage<AssumeRoleChainLink> | undefined, b: AssumeRoleChainLink | PlainMessage<AssumeRoleChainLink> | undefined): boolean;
}

