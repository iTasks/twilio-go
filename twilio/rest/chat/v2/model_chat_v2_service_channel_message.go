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

// ChatV2ServiceChannelMessage struct for ChatV2ServiceChannelMessage
type ChatV2ServiceChannelMessage struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The JSON string that stores application-specific data
	Attributes *string `json:"attributes,omitempty"`
	// The content of the message
	Body *string `json:"body,omitempty"`
	// The SID of the Channel the Message resource belongs to
	ChannelSid *string `json:"channel_sid,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The Identity of the message's author
	From *string `json:"from,omitempty"`
	// The index of the message within the Channel
	Index *int32 `json:"index,omitempty"`
	// The Identity of the User who last updated the Message
	LastUpdatedBy *string `json:"last_updated_by,omitempty"`
	// A Media object that describes the Message's media if attached; otherwise, null
	Media *map[string]interface{} `json:"media,omitempty"`
	// The SID of the Service that the resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The SID of the Channel that the message was sent to
	To *string `json:"to,omitempty"`
	// The Message type
	Type *string `json:"type,omitempty"`
	// The absolute URL of the Message resource
	Url *string `json:"url,omitempty"`
	// Whether the message has been edited since  it was created
	WasEdited *bool `json:"was_edited,omitempty"`
}
