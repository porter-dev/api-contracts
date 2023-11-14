// @generated by protoc-gen-es v1.4.2
// @generated from file porter/v1/cluster.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { EKS } from "./eks_pb.js";
import { GKE } from "./gke_pb.js";
import { AKS } from "./aks_pb.js";

/**
 * @generated from enum porter.v1.EnumKubernetesKind
 */
export const EnumKubernetesKind = proto3.makeEnum(
  "porter.v1.EnumKubernetesKind",
  [
    {no: 0, name: "ENUM_KUBERNETES_KIND_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "ENUM_KUBERNETES_KIND_EKS", localName: "EKS"},
    {no: 2, name: "ENUM_KUBERNETES_KIND_GKE", localName: "GKE"},
    {no: 3, name: "ENUM_KUBERNETES_KIND_AKS", localName: "AKS"},
  ],
);

/**
 * @generated from enum porter.v1.EnumCloudProvider
 */
export const EnumCloudProvider = proto3.makeEnum(
  "porter.v1.EnumCloudProvider",
  [
    {no: 0, name: "ENUM_CLOUD_PROVIDER_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "ENUM_CLOUD_PROVIDER_AWS", localName: "AWS"},
    {no: 2, name: "ENUM_CLOUD_PROVIDER_GCP", localName: "GCP"},
    {no: 3, name: "ENUM_CLOUD_PROVIDER_AZURE", localName: "AZURE"},
  ],
);

/**
 * @generated from message porter.v1.Cluster
 */
export const Cluster = proto3.makeMessageType(
  "porter.v1.Cluster",
  () => [
    { no: 1, name: "project_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "cluster_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "kind", kind: "enum", T: proto3.getEnumType(EnumKubernetesKind) },
    { no: 4, name: "cloud_provider", kind: "enum", T: proto3.getEnumType(EnumCloudProvider) },
    { no: 5, name: "cloud_provider_credentials_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "eks_kind", kind: "message", T: EKS, oneof: "kind_values" },
    { no: 7, name: "gke_kind", kind: "message", T: GKE, oneof: "kind_values" },
    { no: 8, name: "aks_kind", kind: "message", T: AKS, oneof: "kind_values" },
    { no: 9, name: "is_soc2_compliant", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

