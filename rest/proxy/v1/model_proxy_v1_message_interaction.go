/*
 * Twilio - Proxy
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

// ProxyV1MessageInteraction struct for ProxyV1MessageInteraction
type ProxyV1MessageInteraction struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// A JSON string that includes the message body sent to the participant
	Data *string `json:"data,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Always empty for Message Interactions
	InboundParticipantSid *string `json:"inbound_participant_sid,omitempty"`
	// Always empty for Message Interactions
	InboundResourceSid *string `json:"inbound_resource_sid,omitempty"`
	// Always empty for Message Interactions
	InboundResourceStatus *string `json:"inbound_resource_status,omitempty"`
	// Always empty for Message Interactions
	InboundResourceType *string `json:"inbound_resource_type,omitempty"`
	// Always empty for Message Interactions
	InboundResourceUrl *string `json:"inbound_resource_url,omitempty"`
	// The SID of the outbound Participant resource
	OutboundParticipantSid *string `json:"outbound_participant_sid,omitempty"`
	// The SID of the outbound Message resource
	OutboundResourceSid *string `json:"outbound_resource_sid,omitempty"`
	// The outbound resource status
	OutboundResourceStatus *string `json:"outbound_resource_status,omitempty"`
	// The outbound resource type
	OutboundResourceType *string `json:"outbound_resource_type,omitempty"`
	// The URL of the Twilio message resource
	OutboundResourceUrl *string `json:"outbound_resource_url,omitempty"`
	// The SID of the Participant resource
	ParticipantSid *string `json:"participant_sid,omitempty"`
	// The SID of the resource's parent Service
	ServiceSid *string `json:"service_sid,omitempty"`
	// The SID of the resource's parent Session
	SessionSid *string `json:"session_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The Type of Message Interaction
	Type *string `json:"type,omitempty"`
	// The absolute URL of the MessageInteraction resource
	Url *string `json:"url,omitempty"`
}
