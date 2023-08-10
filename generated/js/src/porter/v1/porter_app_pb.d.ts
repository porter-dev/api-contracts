// @generated by protoc-gen-es v1.2.0
// @generated from file porter/v1/porter_app.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Service } from "./service_pb.js";

/**
 * PorterApp is the top-level configuration for a Porter application, usually found in porter.yaml
 *
 * @generated from message porter.v1.PorterApp
 */
export declare class PorterApp extends Message<PorterApp> {
  /**
   * name is the name of the application
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * services is a map of service names to service configurations
   *
   * @generated from field: map<string, porter.v1.Service> services = 2;
   */
  services: { [key: string]: Service };

  /**
   * env is a map of environment variable names to values
   *
   * @generated from field: map<string, string> env = 3;
   */
  env: { [key: string]: string };

  /**
   * build is the build settings for the application
   *
   * @generated from field: porter.v1.Build build = 4;
   */
  build?: Build;

  /**
   * predeploy is a job service to run before deploying the application
   *
   * @generated from field: porter.v1.Service predeploy = 5;
   */
  predeploy?: Service;

  /**
   * image is the image to use for a given revision of the application
   *
   * @generated from field: porter.v1.AppImage image = 6;
   */
  image?: AppImage;

  constructor(data?: PartialMessage<PorterApp>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.PorterApp";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PorterApp;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PorterApp;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PorterApp;

  static equals(a: PorterApp | PlainMessage<PorterApp> | undefined, b: PorterApp | PlainMessage<PorterApp> | undefined): boolean;
}

/**
 * Build is the build settings for the application
 *
 * @generated from message porter.v1.Build
 */
export declare class Build extends Message<Build> {
  /**
   * context is the path to the build context
   *
   * @generated from field: string context = 1;
   */
  context: string;

  /**
   * method is the build method to use, being one of "pack", "docker", or "registry"
   *
   * @generated from field: string method = 2;
   */
  method: string;

  /**
   * builder is the builder to use for the "pack" build method
   *
   * @generated from field: string builder = 3;
   */
  builder: string;

  /**
   * buildpacks is a list of buildpacks to use for the "pack" build method
   *
   * @generated from field: repeated string buildpacks = 4;
   */
  buildpacks: string[];

  /**
   * dockerfile is the path to the Dockerfile to use for the "docker" build method
   *
   * @generated from field: string dockerfile = 5;
   */
  dockerfile: string;

  constructor(data?: PartialMessage<Build>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Build";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Build;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Build;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Build;

  static equals(a: Build | PlainMessage<Build> | undefined, b: Build | PlainMessage<Build> | undefined): boolean;
}

/**
 * AppImage is the image to use for a given revision of the application
 *
 * @generated from message porter.v1.AppImage
 */
export declare class AppImage extends Message<AppImage> {
  /**
   * repository is the repository to use for the image
   *
   * @generated from field: string repository = 1;
   */
  repository: string;

  /**
   * tag is the tag to use for the image
   *
   * @generated from field: string tag = 2;
   */
  tag: string;

  constructor(data?: PartialMessage<AppImage>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.AppImage";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AppImage;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AppImage;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AppImage;

  static equals(a: AppImage | PlainMessage<AppImage> | undefined, b: AppImage | PlainMessage<AppImage> | undefined): boolean;
}

