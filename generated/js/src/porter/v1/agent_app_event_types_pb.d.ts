// @generated by protoc-gen-es v1.8.0
// @generated from file porter/v1/agent_app_event_types.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

/**
 * @generated from enum porter.v1.AppEventType
 */
export declare enum AppEventType {
  /**
   * APP_EVENT_TYPE_UNSPECIFIED is the default value
   *
   * @generated from enum value: APP_EVENT_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * APP_EVENT_TYPE_INSUFFICIENT_RESOURCES is generated when pods can't be scheduled becase of a resource shortage
   * This is either CPU or Memory, but agent cannot resolve
   *
   * @generated from enum value: APP_EVENT_TYPE_INSUFFICIENT_RESOURCES = 1;
   */
  INSUFFICIENT_RESOURCES = 1,

  /**
   * APP_EVENT_TYPE_INSUFFICIENT_CPU is generated when pods can't be scheduled because there is insufficient CPU
   *
   * @generated from enum value: APP_EVENT_TYPE_INSUFFICIENT_CPU = 5;
   */
  INSUFFICIENT_CPU = 5,

  /**
   * APP_EVENT_TYPE_INSUFFICIENT_MEMORY is generated when pods can't be scheduled because there is insufficient memory
   *
   * @generated from enum value: APP_EVENT_TYPE_INSUFFICIENT_MEMORY = 10;
   */
  INSUFFICIENT_MEMORY = 10,

  /**
   * APP_EVENT_TYPE_STUCK_PENDING is generated when pods can't be scheduled for other reasons
   *
   * @generated from enum value: APP_EVENT_TYPE_STUCK_PENDING = 15;
   */
  STUCK_PENDING = 15,

  /**
   * APP_EVENT_TYPE_INVALID_IMAGE is generated when the image specified in containers is not valid
   *
   * @generated from enum value: APP_EVENT_TYPE_INVALID_IMAGE = 20;
   */
  INVALID_IMAGE = 20,

  /**
   * APP_EVENT_TYPE_INVALID_START_COMMAND is generated when a start command in a container is not valid
   *
   * @generated from enum value: APP_EVENT_TYPE_INVALID_START_COMMAND = 25;
   */
  INVALID_START_COMMAND = 25,

  /**
   * APP_EVENT_TYPE_OUT_OF_MEMORY is generated when a container OOMs
   *
   * @generated from enum value: APP_EVENT_TYPE_OUT_OF_MEMORY = 30;
   */
  OUT_OF_MEMORY = 30,

  /**
   * APP_EVENT_TYPE_NON_ZERO_EXIT_CODE is generated when a container terminates with another non-zero code
   *
   * @generated from enum value: APP_EVENT_TYPE_NON_ZERO_EXIT_CODE = 35;
   */
  NON_ZERO_EXIT_CODE = 35,

  /**
   * APP_EVENT_TYPE_FAILING_HEALTH_CHECK is generated when a container is restarted due to a failing health check
   *
   * @generated from enum value: APP_EVENT_TYPE_FAILING_HEALTH_CHECK = 40;
   */
  FAILING_HEALTH_CHECK = 40,
}

