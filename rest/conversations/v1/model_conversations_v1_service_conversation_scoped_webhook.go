/*
 * Twilio - Conversations
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

// ConversationsV1ServiceConversationScopedWebhook struct for ConversationsV1ServiceConversationScopedWebhook
type ConversationsV1ServiceConversationScopedWebhook struct {
	// The unique ID of the Account responsible for this conversation.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Conversation Service that the resource is associated with.
	ChatServiceSid *string `json:"chat_service_sid,omitempty"`
	// The configuration of this webhook.
	Configuration *map[string]interface{} `json:"configuration,omitempty"`
	// The unique ID of the Conversation for this webhook.
	ConversationSid *string `json:"conversation_sid,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this resource was last updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// A 34 character string that uniquely identifies this resource.
	Sid *string `json:"sid,omitempty"`
	// The target of this webhook.
	Target *string `json:"target,omitempty"`
	// An absolute URL for this webhook.
	Url *string `json:"url,omitempty"`
}
