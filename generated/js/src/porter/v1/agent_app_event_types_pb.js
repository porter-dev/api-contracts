// @generated by protoc-gen-es v1.6.0
// @generated from file porter/v1/agent_app_event_types.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum porter.v1.AppEventType
 */
export const AppEventType = proto3.makeEnum(
  "porter.v1.AppEventType",
  [
    {no: 0, name: "APP_EVENT_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "APP_EVENT_TYPE_INSUFFICIENT_RESOURCES", localName: "INSUFFICIENT_RESOURCES"},
    {no: 3, name: "APP_EVENT_TYPE_MAX_RESOURCE_LIMIT_EXCEEDED", localName: "MAX_RESOURCE_LIMIT_EXCEEDED"},
    {no: 5, name: "APP_EVENT_TYPE_INSUFFICIENT_CPU", localName: "INSUFFICIENT_CPU"},
    {no: 10, name: "APP_EVENT_TYPE_INSUFFICIENT_MEMORY", localName: "INSUFFICIENT_MEMORY"},
    {no: 15, name: "APP_EVENT_TYPE_STUCK_PENDING", localName: "STUCK_PENDING"},
    {no: 20, name: "APP_EVENT_TYPE_INVALID_IMAGE", localName: "INVALID_IMAGE"},
    {no: 25, name: "APP_EVENT_TYPE_INVALID_START_COMMAND", localName: "INVALID_START_COMMAND"},
    {no: 30, name: "APP_EVENT_TYPE_OUT_OF_MEMORY", localName: "OUT_OF_MEMORY"},
    {no: 35, name: "APP_EVENT_TYPE_NON_ZERO_EXIT_CODE", localName: "NON_ZERO_EXIT_CODE"},
    {no: 40, name: "APP_EVENT_TYPE_FAILING_HEALTH_CHECK", localName: "FAILING_HEALTH_CHECK"},
    {no: 45, name: "APP_EVENT_TYPE_JOB_TIMEOUT", localName: "JOB_TIMEOUT"},
  ],
);

