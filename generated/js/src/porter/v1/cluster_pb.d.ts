// @generated by protoc-gen-es v1.8.0
// @generated from file porter/v1/cluster.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { EKS } from "./eks_pb.js";
import type { GKE } from "./gke_pb.js";
import type { AKS } from "./aks_pb.js";

/**
 * @generated from enum porter.v1.EnumKubernetesKind
 */
export declare enum EnumKubernetesKind {
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

  /**
   * @generated from enum value: ENUM_KUBERNETES_KIND_AKS = 3;
   */
  AKS = 3,
}

/**
 * @generated from enum porter.v1.EnumCloudProvider
 */
export declare enum EnumCloudProvider {
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

  /**
   * @generated from enum value: ENUM_CLOUD_PROVIDER_AZURE = 3;
   */
  AZURE = 3,
}

/**
 * @generated from message porter.v1.Cluster
 */
export declare class Cluster extends Message<Cluster> {
  /**
   * project_id [REQUIRED] represents the Porter project that the cluster will be joined to. This is required for all cluster creations and updates
   *
   * @generated from field: int32 project_id = 1;
   */
  projectId: number;

  /**
   * cluster_id [OPTIONAL] represents the Porter cluster. This is required for update operations, but should be left blank when creating a cluster
   *
   * @generated from field: int32 cluster_id = 2;
   */
  clusterId: number;

  /**
   * kind [REQUIRED] is the different types of supported kubernetes distros such as EKS, GKE etc.
   *
   * @generated from field: porter.v1.EnumKubernetesKind kind = 3;
   */
  kind: EnumKubernetesKind;

  /**
   * cloud_provider [REQUIRED] represents the provider that we will provisioning the cluster in
   *
   * @generated from field: porter.v1.EnumCloudProvider cloud_provider = 4;
   */
  cloudProvider: EnumCloudProvider;

  /**
   * cloud_provider_credentials_id [REQUIRED] is the Porter credentials that will be used for provisioning a cluster.
   * These must be stored within Porter, prior to cluster creation. For AWS this refers to the last link in an assume role chain
   *
   * @generated from field: string cloud_provider_credentials_id = 5;
   */
  cloudProviderCredentialsId: string;

  /**
   * kind_values are the required values, depending on the selected cloud_provider and kind
   *
   * @generated from oneof porter.v1.Cluster.kind_values
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
  } | {
    /**
     * @generated from field: porter.v1.AKS aks_kind = 8;
     */
    value: AKS;
    case: "aksKind";
  } | { case: undefined; value?: undefined };

  /**
   * is_soc2_compliant force enables all the various soc2-related fields on a cluster
   * deprecated: set soc2 to be enabled under compliance profiles for the contract
   *
   * @generated from field: bool is_soc2_compliant = 9 [deprecated = true];
   * @deprecated
   */
  isSoc2Compliant: boolean;

  constructor(data?: PartialMessage<Cluster>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.Cluster";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Cluster;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Cluster;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Cluster;

  static equals(a: Cluster | PlainMessage<Cluster> | undefined, b: Cluster | PlainMessage<Cluster> | undefined): boolean;
}

