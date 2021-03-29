/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// ChatV2Service struct for ChatV2Service
type ChatV2Service struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// DEPRECATED
	ConsumptionReportInterval *int32 `json:"consumption_report_interval,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The channel role assigned to a channel creator when they join a new channel
	DefaultChannelCreatorRoleSid *string `json:"default_channel_creator_role_sid,omitempty"`
	// The channel role assigned to users when they are added to a channel
	DefaultChannelRoleSid *string `json:"default_channel_role_sid,omitempty"`
	// The service role assigned to users when they are added to the service
	DefaultServiceRoleSid *string `json:"default_service_role_sid,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// An object that describes the limits of the service instance
	Limits *map[string]interface{} `json:"limits,omitempty"`
	// The absolute URLs of the Service's Channels, Roles, and Users
	Links *map[string]interface{} `json:"links,omitempty"`
	// The properties of the media that the service supports
	Media *map[string]interface{} `json:"media,omitempty"`
	// The notification configuration for the Service instance
	Notifications *map[string]interface{} `json:"notifications,omitempty"`
	// The number of times calls to the `post_webhook_url` will be retried
	PostWebhookRetryCount *int32 `json:"post_webhook_retry_count,omitempty"`
	// The URL for post-event webhooks
	PostWebhookUrl *string `json:"post_webhook_url,omitempty"`
	// Count of times webhook will be retried in case of timeout or 429/503/504 HTTP responses
	PreWebhookRetryCount *int32 `json:"pre_webhook_retry_count,omitempty"`
	// The webhook URL for pre-event webhooks
	PreWebhookUrl *string `json:"pre_webhook_url,omitempty"`
	// Whether the Reachability Indicator feature is enabled for this Service instance
	ReachabilityEnabled *bool `json:"reachability_enabled,omitempty"`
	// Whether the Message Consumption Horizon feature is enabled
	ReadStatusEnabled *bool `json:"read_status_enabled,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// How long in seconds to wait before assuming the user is no longer typing
	TypingIndicatorTimeout *int32 `json:"typing_indicator_timeout,omitempty"`
	// The absolute URL of the Service resource
	Url *string `json:"url,omitempty"`
	// The list of webhook events that are enabled for this Service instance
	WebhookFilters *[]string `json:"webhook_filters,omitempty"`
	// The HTTP method  to use for both PRE and POST webhooks
	WebhookMethod *string `json:"webhook_method,omitempty"`
}
