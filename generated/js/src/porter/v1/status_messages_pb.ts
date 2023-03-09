// @generated by protoc-gen-es v1.1.0 with parameter "target=ts"
// @generated from file porter/v1/status_messages.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message porter.v1.MessageSystemInfrastructure
 */
export class MessageSystemInfrastructure extends Message<MessageSystemInfrastructure> {
  /**
   * @generated from field: int32 cluster_id = 1;
   */
  clusterId = 0;

  /**
   * @generated from field: int32 project_id = 2;
   */
  projectId = 0;

  /**
   * @generated from field: bool is_error = 3;
   */
  isError = false;

  /**
   * @generated from field: bool is_retryable = 4;
   */
  isRetryable = false;

  /**
   * @generated from field: string message = 5;
   */
  message = "";

  constructor(data?: PartialMessage<MessageSystemInfrastructure>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "porter.v1.MessageSystemInfrastructure";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "cluster_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "project_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "is_error", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 4, name: "is_retryable", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 5, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MessageSystemInfrastructure {
    return new MessageSystemInfrastructure().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MessageSystemInfrastructure {
    return new MessageSystemInfrastructure().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MessageSystemInfrastructure {
    return new MessageSystemInfrastructure().fromJsonString(jsonString, options);
  }

  static equals(a: MessageSystemInfrastructure | PlainMessage<MessageSystemInfrastructure> | undefined, b: MessageSystemInfrastructure | PlainMessage<MessageSystemInfrastructure> | undefined): boolean {
    return proto3.util.equals(MessageSystemInfrastructure, a, b);
  }
}

