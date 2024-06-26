// @generated by protoc-gen-es v1.10.0
// @generated from file porter/v1/errors.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * Error is a lightweight error message contract
 *
 * @generated from message porter.v1.Error
 */
export declare class Error extends Message<Error> {
  /**
   * code is an error code which can be programatically cased on
   *
   * @generated from field: string code = 1;
   */
  code: string;

  /**
   * message is a human readable error message, which should only be used in fallback cases. The caller should handle what is shown to the user
   *
   * @generated from field: string message = 2;
   */
  message: string;

  /**
   * metadata is a map of additional metadata about the error
   *
   * @generated from field: map<string, string> metadata = 3;
   */
  metadata: { [key: string]: string };

  constructor(data?: PartialMessage<Error>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Error";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Error;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Error;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Error;

  static equals(a: Error | PlainMessage<Error> | undefined, b: Error | PlainMessage<Error> | undefined): boolean;
}

