// @generated by protoc-gen-es v1.8.0
// @generated from file porter/v1/app_event_webhook.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * WebhookAppEventType is the app event type a webhook can be configured for
 *
 * @generated from enum porter.v1.WebhookAppEventType
 */
export const WebhookAppEventType = /*@__PURE__*/ proto3.makeEnum(
  "porter.v1.WebhookAppEventType",
  [
    {no: 0, name: "WEBHOOK_APP_EVENT_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "WEBHOOK_APP_EVENT_TYPE_DEPLOY", localName: "DEPLOY"},
    {no: 2, name: "WEBHOOK_APP_EVENT_TYPE_BUILD", localName: "BUILD"},
    {no: 3, name: "WEBHOOK_APP_EVENT_TYPE_PREDEPLOY", localName: "PREDEPLOY"},
  ],
);

/**
 * WebhookAppEventStatus is the app evnet status a webhook can be configured for
 *
 * @generated from enum porter.v1.WebhookAppEventStatus
 */
export const WebhookAppEventStatus = /*@__PURE__*/ proto3.makeEnum(
  "porter.v1.WebhookAppEventStatus",
  [
    {no: 0, name: "WEBHOOK_APP_EVENT_STATUS_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "WEBHOOK_APP_EVENT_STATUS_SUCCESS", localName: "SUCCESS"},
    {no: 2, name: "WEBHOOK_APP_EVENT_STATUS_FAILED", localName: "FAILED"},
    {no: 3, name: "WEBHOOK_APP_EVENT_STATUS_CANCELED", localName: "CANCELED"},
  ],
);

/**
 * AppEventWebhook is a webhook configured to have requests sent to on app events
 *
 * @generated from message porter.v1.AppEventWebhook
 */
export const AppEventWebhook = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.AppEventWebhook",
  () => [
    { no: 1, name: "webhook_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "app_event_type", kind: "enum", T: proto3.getEnumType(WebhookAppEventType) },
    { no: 3, name: "app_event_status", kind: "enum", T: proto3.getEnumType(WebhookAppEventStatus) },
    { no: 4, name: "payload_encryption_key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

