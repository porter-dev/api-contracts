// @generated by protoc-gen-es v1.1.0 with parameter "target=ts,import_extension=none"
// @generated from file porter/v1/aws_assume_role.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * @generated from message porter.v1.AssumeRoleChainLink
 */
export class AssumeRoleChainLink extends Message<AssumeRoleChainLink> {
  /**
   * @generated from field: int64 project_id = 1;
   */
  projectId = protoInt64.zero;

  /**
   * @generated from field: string source_arn = 2;
   */
  sourceArn = "";

  /**
   * @generated from field: string target_arn = 3;
   */
  targetArn = "";

  /**
   * @generated from field: string external_id = 4;
   */
  externalId = "";

  constructor(data?: PartialMessage<AssumeRoleChainLink>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "porter.v1.AssumeRoleChainLink";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "source_arn", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "target_arn", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "external_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AssumeRoleChainLink {
    return new AssumeRoleChainLink().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AssumeRoleChainLink {
    return new AssumeRoleChainLink().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AssumeRoleChainLink {
    return new AssumeRoleChainLink().fromJsonString(jsonString, options);
  }

  static equals(a: AssumeRoleChainLink | PlainMessage<AssumeRoleChainLink> | undefined, b: AssumeRoleChainLink | PlainMessage<AssumeRoleChainLink> | undefined): boolean {
    return proto3.util.equals(AssumeRoleChainLink, a, b);
  }
}

