// @generated by protoc-gen-es v1.8.0
// @generated from file porter/v1/app_event_webhook.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * WebhookAppEventType is the app event type a webhook can be configured for
 *
 * @generated from enum porter.v1.WebhookAppEventType
 */
export declare enum WebhookAppEventType {
  /**
   * WEBHOOK_APP_EVENT_TYPE_UNSPECIFIED is the default value
   *
   * @generated from enum value: WEBHOOK_APP_EVENT_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * WEBHOOK_APP_EVENT_TYPE_DEPLOY is when the webhook is configured for deploy events
   *
   * @generated from enum value: WEBHOOK_APP_EVENT_TYPE_DEPLOY = 1;
   */
  DEPLOY = 1,

  /**
   * WEBHOOK_APP_EVENT_TYPE_BUILD is when the webhook is configured for build events
   *
   * @generated from enum value: WEBHOOK_APP_EVENT_TYPE_BUILD = 2;
   */
  BUILD = 2,

  /**
   * WEBHOOK_APP_EVENT_TYPE_PREDEPLOYT is when the webhook is configured for predeploy events
   *
   * @generated from enum value: WEBHOOK_APP_EVENT_TYPE_PREDEPLOY = 3;
   */
  PREDEPLOY = 3,

  /**
   * WEBHOOK_APP_EVENT_TYPE_INIT_DEPLOY is when the webhook is configured for initial deploy
   *
   * @generated from enum value: WEBHOOK_APP_EVENT_TYPE_INIT_DEPLOY = 4;
   */
  INIT_DEPLOY = 4,
}

/**
 * WebhookAppEventStatus is the app evnet status a webhook can be configured for
 *
 * @generated from enum porter.v1.WebhookAppEventStatus
 */
export declare enum WebhookAppEventStatus {
  /**
   * WEBHOOK_APP_EVENT_STATUS_UNSPECIFIED is the default value
   *
   * @generated from enum value: WEBHOOK_APP_EVENT_STATUS_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * WEBHOOK_APP_EVENT_STATUS_SUCCESS configures webhook for success events
   *
   * @generated from enum value: WEBHOOK_APP_EVENT_STATUS_SUCCESS = 1;
   */
  SUCCESS = 1,

  /**
   * WEBHOOK_APP_EVENT_STATUS_FAILED configures whebhook for failed events
   *
   * @generated from enum value: WEBHOOK_APP_EVENT_STATUS_FAILED = 2;
   */
  FAILED = 2,

  /**
   * WEBHOOK_APP_EVENT_STATUS_PROGRESSING configures webhook for progressing events
   *
   * @generated from enum value: WEBHOOK_APP_EVENT_STATUS_PROGRESSING = 3;
   */
  PROGRESSING = 3,

  /**
   * WEBHOOK_APP_EVENT_STATUS_PROGRESSING configures webhook for canceled events
   *
   * @generated from enum value: WEBHOOK_APP_EVENT_STATUS_CANCELED = 4;
   */
  CANCELED = 4,
}

/**
 * AppEventWebhook is a webhook configured to have requests sent to on app events
 *
 * @generated from message porter.v1.AppEventWebhook
 */
export declare class AppEventWebhook extends Message<AppEventWebhook> {
  /**
   * id is a unique identifier for a webhook
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * webhook_url is the url of the webhook
   *
   * @generated from field: bytes webhook_url = 2;
   */
  webhookUrl: Uint8Array;

  /**
   * app_event_type is the type of the app event this webhook is being added for
   *
   * @generated from field: porter.v1.WebhookAppEventType app_event_type = 3;
   */
  appEventType: WebhookAppEventType;

  /**
   * app_event_status is the status of the app event this webhook is being added for
   *
   * @generated from field: porter.v1.WebhookAppEventStatus app_event_status = 4;
   */
  appEventStatus: WebhookAppEventStatus;

  constructor(data?: PartialMessage<AppEventWebhook>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "porter.v1.AppEventWebhook";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AppEventWebhook;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AppEventWebhook;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AppEventWebhook;

  static equals(a: AppEventWebhook | PlainMessage<AppEventWebhook> | undefined, b: AppEventWebhook | PlainMessage<AppEventWebhook> | undefined): boolean;
}

