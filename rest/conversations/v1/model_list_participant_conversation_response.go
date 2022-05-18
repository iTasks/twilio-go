/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListParticipantConversationResponse struct for ListParticipantConversationResponse
type ListParticipantConversationResponse struct {
	Conversations []ConversationsV1ParticipantConversation `json:"conversations,omitempty"`
	Meta          ListConfigurationAddressResponseMeta     `json:"meta,omitempty"`
}
