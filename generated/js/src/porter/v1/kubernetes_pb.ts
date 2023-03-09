// @generated by protoc-gen-es v1.1.0 with parameter "target=ts"
// @generated from file porter/v1/kubernetes.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { EKS } from "./eks_pb.js";
import { GKE } from "./gke_pb.js";

/**
 * @generated from enum porter.v1.EnumKubernetesKind
 */
export enum EnumKubernetesKind {
  /**
   * @generated from enum value: ENUM_KUBERNETES_KIND_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: ENUM_KUBERNETES_KIND_EKS = 1;
   */
  EKS = 1,

  /**
   * @generated from enum value: ENUM_KUBERNETES_KIND_GKE = 2;
   */
  GKE = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(EnumKubernetesKind)
proto3.util.setEnumType(EnumKubernetesKind, "porter.v1.EnumKubernetesKind", [
  { no: 0, name: "ENUM_KUBERNETES_KIND_UNSPECIFIED" },
  { no: 1, name: "ENUM_KUBERNETES_KIND_EKS" },
  { no: 2, name: "ENUM_KUBERNETES_KIND_GKE" },
]);

/**
 * @generated from enum porter.v1.EnumCloudProvider
 */
export enum EnumCloudProvider {
  /**
   * @generated from enum value: ENUM_CLOUD_PROVIDER_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: ENUM_CLOUD_PROVIDER_AWS = 1;
   */
  AWS = 1,

  /**
   * @generated from enum value: ENUM_CLOUD_PROVIDER_GCP = 2;
   */
  GCP = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(EnumCloudProvider)
proto3.util.setEnumType(EnumCloudProvider, "porter.v1.EnumCloudProvider", [
  { no: 0, name: "ENUM_CLOUD_PROVIDER_UNSPECIFIED" },
  { no: 1, name: "ENUM_CLOUD_PROVIDER_AWS" },
  { no: 2, name: "ENUM_CLOUD_PROVIDER_GCP" },
]);

/**
 * @generated from message porter.v1.Kubernetes
 */
export class Kubernetes extends Message<Kubernetes> {
  /**
   * project_id [REQUIRED] represents the Porter project that the cluster will be joined to. This is required for all cluster creations and updates
   *
   * @generated from field: int32 project_id = 1;
   */
  projectId = 0;

  /**
   * cluster_id [OPTIONAL] represents the Porter cluster. This is required for update operations, but should be left blank when creating a cluster
   *
   * @generated from field: int32 cluster_id = 2;
   */
  clusterId = 0;

  /**
   * kind [REQUIRED] is the different types of supported kubernetes distros such as EKS, GKE etc.
   *
   * @generated from field: porter.v1.EnumKubernetesKind kind = 3;
   */
  kind = EnumKubernetesKind.UNSPECIFIED;

  /**
   * cloud_provider [REQUIRED] represents the provider that we will provisioning the cluster in
   *
   * @generated from field: porter.v1.EnumCloudProvider cloud_provider = 4;
   */
  cloudProvider = EnumCloudProvider.UNSPECIFIED;

  /**
   * cloud_provider_credentials_id [REQUIRED] is the Porter credentials that will be used for provisioning a cluster.
   * These must be stored within Porter, prior to cluster creation. For AWS this refers to the last link in an assume role chain
   *
   * @generated from field: string cloud_provider_credentials_id = 5;
   */
  cloudProviderCredentialsId = "";

  /**
   * kind_values are the required values, depending on the selected cloud_provider and kind
   *
   * @generated from oneof porter.v1.Kubernetes.kind_values
   */
  kindValues: {
    /**
     * @generated from field: porter.v1.EKS eks_kind = 6;
     */
    value: EKS;
    case: "eksKind";
  } | {
    /**
     * @generated from field: porter.v1.GKE gke_kind = 7;
     */
    value: GKE;
    case: "gkeKind";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<Kubernetes>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "porter.v1.Kubernetes";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "project_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "cluster_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "kind", kind: "enum", T: proto3.getEnumType(EnumKubernetesKind) },
    { no: 4, name: "cloud_provider", kind: "enum", T: proto3.getEnumType(EnumCloudProvider) },
    { no: 5, name: "cloud_provider_credentials_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "eks_kind", kind: "message", T: EKS, oneof: "kind_values" },
    { no: 7, name: "gke_kind", kind: "message", T: GKE, oneof: "kind_values" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Kubernetes {
    return new Kubernetes().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Kubernetes {
    return new Kubernetes().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Kubernetes {
    return new Kubernetes().fromJsonString(jsonString, options);
  }

  static equals(a: Kubernetes | PlainMessage<Kubernetes> | undefined, b: Kubernetes | PlainMessage<Kubernetes> | undefined): boolean {
    return proto3.util.equals(Kubernetes, a, b);
  }
}

