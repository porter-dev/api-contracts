// @generated by protoc-gen-es v1.1.1
// @generated from file porter/v1/gke.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message porter.v1.GKE
 */
export const GKE = proto3.makeMessageType(
  "porter.v1.GKE",
  () => [
    { no: 1, name: "cluster_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "cluster_version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "issuer_email", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "cidr_range", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

