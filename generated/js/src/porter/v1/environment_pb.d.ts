// @generated by protoc-gen-es v1.10.0
// @generated from file porter/v1/environment.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { PorterApp } from "./porter_app_pb.js";
import type { Addon } from "./addons_pb.js";

/**
 * @generated from message porter.v1.Environment
 */
export declare class Environment extends Message<Environment> {
  /**
   * name is the name of the environment
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * apps is a list of apps that should be included in the environment
   *
   * @generated from field: repeated porter.v1.PorterApp apps = 2;
   */
  apps: PorterApp[];

  /**
   * addons is a list of addons that should be included in the environment
   *
   * @generated from field: repeated porter.v1.Addon addons = 3;
   */
  addons: Addon[];

  constructor(data?: PartialMessage<Environment>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Environment";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Environment;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Environment;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Environment;

  static equals(a: Environment | PlainMessage<Environment> | undefined, b: Environment | PlainMessage<Environment> | undefined): boolean;
}

