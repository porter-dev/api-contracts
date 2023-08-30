// @generated by protoc-gen-es v1.3.0
// @generated from file porter/v1/porter_app.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { Service } from "./service_pb.js";

/**
 * PorterApp is the top-level configuration for a Porter application, usually found in porter.yaml
 *
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
 * Deletions contains all explicit deletions from a PorterApp
 *
 * @generated from message porter.v1.Deletions
 */
export const Deletions = proto3.makeMessageType(
  "porter.v1.Deletions",
  () => [
    { no: 1, name: "service_names", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "env_variable_names", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * Build is the build settings for the application
 *
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
    { no: 6, name: "commit_sha", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * AppImage is the image to use for a given revision of the application
 *
 * @generated from message porter.v1.AppImage
 */
export const AppImage = proto3.makeMessageType(
  "porter.v1.AppImage",
  () => [
    { no: 1, name: "repository", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "tag", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

