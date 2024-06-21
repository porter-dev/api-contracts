// @generated by protoc-gen-es v1.10.0
// @generated from file porter/v1/storage.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { EnumCloudProvider } from "./cluster_pb.js";
import { ConnectedClusters } from "./datastore_pb.js";

/**
 * @generated from enum porter.v1.EnumStorageKind
 */
export const EnumStorageKind = /*@__PURE__*/ proto3.makeEnum(
  "porter.v1.EnumStorageKind",
  [
    {no: 0, name: "ENUM_STORAGE_KIND_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "ENUM_STORAGE_KIND_SHARED_FILE_SYSTEM", localName: "SHARED_FILE_SYSTEM"},
    {no: 2, name: "ENUM_STORAGE_KIND_OBJECT_STORAGE", localName: "OBJECT_STORAGE"},
  ],
);

/**
 * @generated from enum porter.v1.EnumSharedFileSystemEngine
 */
export const EnumSharedFileSystemEngine = /*@__PURE__*/ proto3.makeEnum(
  "porter.v1.EnumSharedFileSystemEngine",
  [
    {no: 0, name: "ENUM_SHARED_FILE_SYSTEM_ENGINE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "ENUM_SHARED_FILE_SYSTEM_ENGINE_AWS_EFS", localName: "AWS_EFS"},
    {no: 2, name: "ENUM_SHARED_FILE_SYSTEM_ENGINE_AZURE_FILES", localName: "AZURE_FILES"},
    {no: 3, name: "ENUM_SHARED_FILE_SYSTEM_ENGINE_GCP_FILESTORE", localName: "GCP_FILESTORE"},
  ],
);

/**
 * @generated from enum porter.v1.EnumObjectStorageEngine
 */
export const EnumObjectStorageEngine = /*@__PURE__*/ proto3.makeEnum(
  "porter.v1.EnumObjectStorageEngine",
  [
    {no: 0, name: "ENUM_OBJECT_STORAGE_ENGINE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "ENUM_OBJECT_STORAGE_ENGINE_AWS_S3", localName: "AWS_S3"},
    {no: 2, name: "ENUM_OBJECT_STORAGE_ENGINE_AZURE_BLOB_STORAGE", localName: "AZURE_BLOB_STORAGE"},
    {no: 3, name: "ENUM_OBJECT_STORAGE_ENGINE_GCP_FILESTORE", localName: "GCP_FILESTORE"},
  ],
);

/**
 * @generated from message porter.v1.Storage
 */
export const Storage = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.Storage",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "cloud_provider", kind: "enum", T: proto3.getEnumType(EnumCloudProvider) },
    { no: 3, name: "cloud_provider_credential_identifier", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "region", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "kind", kind: "enum", T: proto3.getEnumType(EnumStorageKind) },
    { no: 14, name: "shared_file_system_kind", kind: "message", T: SharedFileSystem, oneof: "kind_values" },
    { no: 15, name: "object_storage_kind", kind: "message", T: ObjectStorage, oneof: "kind_values" },
    { no: 8, name: "connected_clusters", kind: "message", T: ConnectedClusters },
  ],
);

/**
 * @generated from message porter.v1.SharedFileSystem
 */
export const SharedFileSystem = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.SharedFileSystem",
  () => [
    { no: 1, name: "engine", kind: "enum", T: proto3.getEnumType(EnumSharedFileSystemEngine) },
  ],
);

/**
 * @generated from message porter.v1.ObjectStorage
 */
export const ObjectStorage = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.ObjectStorage",
  () => [
    { no: 1, name: "engine", kind: "enum", T: proto3.getEnumType(EnumObjectStorageEngine) },
  ],
);

