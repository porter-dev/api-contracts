// @generated by protoc-gen-es v1.8.0
// @generated from file porter/v1/notification.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * EnumNotificationStatus describes the status of a notification
 *
 * @generated from enum porter.v1.EnumNotificationStatus
 */
export const EnumNotificationStatus = /*@__PURE__*/ proto3.makeEnum(
  "porter.v1.EnumNotificationStatus",
  [
    {no: 0, name: "ENUM_NOTIFICATION_STATUS_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "ENUM_NOTIFICATION_STATUS_PROGRESSING", localName: "PROGRESSING"},
    {no: 2, name: "ENUM_NOTIFICATION_STATUS_SUCCESSFUL", localName: "SUCCESSFUL"},
    {no: 3, name: "ENUM_NOTIFICATION_STATUS_FAILED", localName: "FAILED"},
  ],
);

/**
 * EnumNotificationEventType describes the type of notification event
 *
 * @generated from enum porter.v1.EnumNotificationEventType
 */
export const EnumNotificationEventType = /*@__PURE__*/ proto3.makeEnum(
  "porter.v1.EnumNotificationEventType",
  [
    {no: 0, name: "ENUM_NOTIFICATION_EVENT_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "ENUM_NOTIFICATION_EVENT_TYPE_BUILD", localName: "BUILD"},
    {no: 2, name: "ENUM_NOTIFICATION_EVENT_TYPE_PREDEPLOY", localName: "PREDEPLOY"},
    {no: 3, name: "ENUM_NOTIFICATION_EVENT_TYPE_DEPLOY", localName: "DEPLOY"},
    {no: 4, name: "ENUM_NOTIFICATION_EVENT_TYPE_ALERT", localName: "ALERT"},
  ],
);

/**
 * NotificationConfig describes the configuration available for notifications.
 * Empty lists are ignored.
 * Multiple lists are treated as an AND condition.
 * Multiple items in one list are treated as an OR condition.
 * For example, if statuses and cluster_ids are both provided, the notification
 * will only be sent if the event matches one of the statuses AND one of the cluster_ids.
 *
 * @generated from message porter.v1.NotificationConfig
 */
export const NotificationConfig = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.NotificationConfig",
  () => [
    { no: 1, name: "app_instance_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 2, name: "cluster_ids", kind: "scalar", T: 3 /* ScalarType.INT64 */, repeated: true },
    { no: 3, name: "deployment_target_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "statuses", kind: "enum", T: proto3.getEnumType(EnumNotificationStatus), repeated: true },
    { no: 5, name: "event_types", kind: "enum", T: proto3.getEnumType(EnumNotificationEventType), repeated: true },
    { no: 6, name: "slack_config", kind: "message", T: SlackConfig },
    { no: 7, name: "enabled_types", kind: "message", T: NotificationTypeEnabled, repeated: true },
    { no: 8, name: "enabled_statuses", kind: "message", T: NotificationStatusEnabled, repeated: true },
  ],
);

/**
 * NotificationTypeEnabled specifies whether notifications are enabled for a given type
 *
 * @generated from message porter.v1.NotificationTypeEnabled
 */
export const NotificationTypeEnabled = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.NotificationTypeEnabled",
  () => [
    { no: 1, name: "type", kind: "enum", T: proto3.getEnumType(EnumNotificationEventType) },
    { no: 2, name: "enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * NotificationStatusEnabled specifies whether notifications are enabled for a given status
 *
 * @generated from message porter.v1.NotificationStatusEnabled
 */
export const NotificationStatusEnabled = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.NotificationStatusEnabled",
  () => [
    { no: 1, name: "status", kind: "enum", T: proto3.getEnumType(EnumNotificationStatus) },
    { no: 2, name: "enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * SlackConfig describes the slack-specific configuration for notifications
 *
 * @generated from message porter.v1.SlackConfig
 */
export const SlackConfig = /*@__PURE__*/ proto3.makeMessageType(
  "porter.v1.SlackConfig",
  () => [
    { no: 1, name: "mentions", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

