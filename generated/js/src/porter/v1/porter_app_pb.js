// @generated by protoc-gen-es v1.3.0
// @generated from file porter/v1/porter_app.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { Service } from "./service_pb.js";

/**
 * @generated from message porter.v1.PorterApp
 */
export const PorterApp = proto3.makeMessageType(
  "porter.v1.PorterApp",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "services", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: Service} },
    { no: 3, name: "env", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 4, name: "build", kind: "message", T: Build },
    { no: 5, name: "predeploy", kind: "message", T: Service },
    { no: 6, name: "image", kind: "message", T: AppImage },
  ],
);

/**
 * @generated from message porter.v1.Build
 */
export const Build = proto3.makeMessageType(
  "porter.v1.Build",
  () => [
    { no: 1, name: "context", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "method", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "builder", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "buildpacks", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 5, name: "dockerfile", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message porter.v1.AppImage
 */
export const AppImage = proto3.makeMessageType(
  "porter.v1.AppImage",
  () => [
    { no: 1, name: "repository", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "tag", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

