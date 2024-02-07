// @generated by protoc-gen-es v1.4.2
// @generated from file porter/v1/notification.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * EnumNotificationStatus describes the status of a notification
 *
 * @generated from enum porter.v1.EnumNotificationStatus
 */
export declare enum EnumNotificationStatus {
  /**
   * @generated from enum value: ENUM_NOTIFICATION_STATUS_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: ENUM_NOTIFICATION_STATUS_PROGRESSING = 1;
   */
  PROGRESSING = 1,

  /**
   * @generated from enum value: ENUM_NOTIFICATION_STATUS_SUCCESSFUL = 2;
   */
  SUCCESSFUL = 2,

  /**
   * @generated from enum value: ENUM_NOTIFICATION_STATUS_FAILED = 3;
   */
  FAILED = 3,
}

/**
 * EnumNotificationEventType describes the type of notification event
 *
 * @generated from enum porter.v1.EnumNotificationEventType
 */
export declare enum EnumNotificationEventType {
  /**
   * @generated from enum value: ENUM_NOTIFICATION_EVENT_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: ENUM_NOTIFICATION_EVENT_TYPE_BUILD = 1;
   */
  BUILD = 1,

  /**
   * @generated from enum value: ENUM_NOTIFICATION_EVENT_TYPE_PREDEPLOY = 2;
   */
  PREDEPLOY = 2,

  /**
   * @generated from enum value: ENUM_NOTIFICATION_EVENT_TYPE_DEPLOY = 3;
   */
  DEPLOY = 3,
}

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
export declare class NotificationConfig extends Message<NotificationConfig> {
  /**
   * app_instance_ids is a list of app instance ids
   *
   * @generated from field: repeated string app_instance_ids = 1;
   */
  appInstanceIds: string[];

  /**
   * cluster_ids is a list of cluster ids
   *
   * @generated from field: repeated int64 cluster_ids = 2;
   */
  clusterIds: bigint[];

  /**
   * deployment_target_ids is a list of deployment target ids
   *
   * @generated from field: repeated string deployment_target_ids = 3;
   */
  deploymentTargetIds: string[];

  /**
   * statuses is a list of statuses
   *
   * @generated from field: repeated porter.v1.EnumNotificationStatus statuses = 4;
   */
  statuses: EnumNotificationStatus[];

  /**
   * event_types is a list of event types
   *
   * @generated from field: repeated porter.v1.EnumNotificationEventType event_types = 5;
   */
  eventTypes: EnumNotificationEventType[];

  /**
   * slack_config is the slack-specific configuration for notifications
   *
   * @generated from field: porter.v1.SlackConfig slack_config = 6;
   */
  slackConfig?: SlackConfig;

  constructor(data?: PartialMessage<NotificationConfig>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.NotificationConfig";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NotificationConfig;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NotificationConfig;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NotificationConfig;

  static equals(a: NotificationConfig | PlainMessage<NotificationConfig> | undefined, b: NotificationConfig | PlainMessage<NotificationConfig> | undefined): boolean;
}

/**
 * SlackConfig describes the slack-specific configuration for notifications
 *
 * @generated from message porter.v1.SlackConfig
 */
export declare class SlackConfig extends Message<SlackConfig> {
  /**
   * mentions is a list of slack handles to mention on failure
   *
   * @generated from field: repeated string mentions = 1;
   */
  mentions: string[];

  constructor(data?: PartialMessage<SlackConfig>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.SlackConfig";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SlackConfig;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SlackConfig;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SlackConfig;

  static equals(a: SlackConfig | PlainMessage<SlackConfig> | undefined, b: SlackConfig | PlainMessage<SlackConfig> | undefined): boolean;
}

