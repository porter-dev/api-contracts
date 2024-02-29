// @generated by protoc-gen-es v1.4.2
// @generated from file porter/v1/datastore.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { EnumCloudProvider } from "./cluster_pb.js";

/**
 * @generated from enum porter.v1.EnumDatastoreKind
 */
export const EnumDatastoreKind = proto3.makeEnum(
  "porter.v1.EnumDatastoreKind",
  [
    {no: 0, name: "ENUM_DATASTORE_KIND_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "ENUM_DATASTORE_KIND_AWS_RDS_POSTGRES", localName: "AWS_RDS_POSTGRES"},
  ],
);

/**
 * PorterDatastore is the specification for a Porter-managed datastore
 *
 * @generated from message porter.v1.PorterDatastore
 */
export const PorterDatastore = proto3.makeMessageType(
  "porter.v1.PorterDatastore",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "cloud_provider", kind: "enum", T: proto3.getEnumType(EnumCloudProvider) },
    { no: 3, name: "cloud_provider_credential_identifier", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "region", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "kind", kind: "enum", T: proto3.getEnumType(EnumDatastoreKind) },
    { no: 7, name: "aws_rds_postgres_kind", kind: "message", T: AwsRdsPostgres, oneof: "kind_values" },
  ],
);

/**
 * @generated from message porter.v1.DatastorePasswordSecretRef
 */
export const DatastorePasswordSecretRef = proto3.makeMessageType(
  "porter.v1.DatastorePasswordSecretRef",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "namespace", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.AwsRdsPostgres
 */
export const AwsRdsPostgres = proto3.makeMessageType(
  "porter.v1.AwsRdsPostgres",
  () => [
    { no: 1, name: "database_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "master_username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "master_user_password_literal", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "master_user_password_secret_ref", kind: "message", T: DatastorePasswordSecretRef },
    { no: 5, name: "allocated_storage_gigabytes", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 6, name: "instance_class", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * UpdateDatastorePayload is used to send messages via nats to update a datastore
 *
 * @generated from message porter.v1.UpdateDatastorePayload
 */
export const UpdateDatastorePayload = proto3.makeMessageType(
  "porter.v1.UpdateDatastorePayload",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "datastore_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * DatastoreCredential is used to connect to a datastore
 *
 * @generated from message porter.v1.DatastoreCredential
 */
export const DatastoreCredential = proto3.makeMessageType(
  "porter.v1.DatastoreCredential",
  () => [
    { no: 1, name: "host", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "database_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "port", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

