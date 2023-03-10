// @generated by protoc-gen-es v1.1.0
// @generated from file porter/v1/gke.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message porter.v1.GKE
 */
export declare class GKE extends Message<GKE> {
  /**
   * @generated from field: string cluster_name = 1;
   */
  clusterName: string;

  /**
   * @generated from field: string cluster_version = 2;
   */
  clusterVersion: string;

  /**
   * @generated from field: string issuer_email = 3;
   */
  issuerEmail: string;

  /**
   * @generated from field: string cidr_range = 4;
   */
  cidrRange: string;

  constructor(data?: PartialMessage<GKE>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.GKE";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GKE;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GKE;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GKE;

  static equals(a: GKE | PlainMessage<GKE> | undefined, b: GKE | PlainMessage<GKE> | undefined): boolean;
}

