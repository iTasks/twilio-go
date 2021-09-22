/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.3
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// ChatV1Message struct for ChatV1Message
type ChatV1Message struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The JSON string that stores application-specific data
	Attributes *string `json:"attributes,omitempty"`
	// The content of the message
	Body *string `json:"body,omitempty"`
	// The unique ID of the Channel the Message resource belongs to
	ChannelSid *string `json:"channel_sid,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The identity of the message's author
	From *string `json:"from,omitempty"`
	// The index of the message within the Channel
	Index *int `json:"index,omitempty"`
	// The SID of the Service that the resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The SID of the Channel that the message was sent to
	To *string `json:"to,omitempty"`
	// The absolute URL of the Message resource
	Url *string `json:"url,omitempty"`
	// Whether the message has been edited since  it was created
	WasEdited *bool `json:"was_edited,omitempty"`
}
