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

// ChatV2ServiceChannelChannelWebhook struct for ChatV2ServiceChannelChannelWebhook
type ChatV2ServiceChannelChannelWebhook struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Channel the Channel Webhook resource belongs to
	ChannelSid *string `json:"channel_sid,omitempty"`
	// The JSON string that describes the configuration object for the channel webhook
	Configuration *map[string]interface{} `json:"configuration,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The SID of the Service that the Channel Webhook resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The type of webhook
	Type *string `json:"type,omitempty"`
	// The absolute URL of the Channel Webhook resource
	Url *string `json:"url,omitempty"`
}
