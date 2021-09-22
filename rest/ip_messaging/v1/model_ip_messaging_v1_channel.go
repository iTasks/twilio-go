/*
 * Twilio - Ip_messaging
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

// IpMessagingV1Channel struct for IpMessagingV1Channel
type IpMessagingV1Channel struct {
	AccountSid    *string                 `json:"account_sid,omitempty"`
	Attributes    *string                 `json:"attributes,omitempty"`
	CreatedBy     *string                 `json:"created_by,omitempty"`
	DateCreated   *time.Time              `json:"date_created,omitempty"`
	DateUpdated   *time.Time              `json:"date_updated,omitempty"`
	FriendlyName  *string                 `json:"friendly_name,omitempty"`
	Links         *map[string]interface{} `json:"links,omitempty"`
	MembersCount  *int                    `json:"members_count,omitempty"`
	MessagesCount *int                    `json:"messages_count,omitempty"`
	ServiceSid    *string                 `json:"service_sid,omitempty"`
	Sid           *string                 `json:"sid,omitempty"`
	Type          *string                 `json:"type,omitempty"`
	UniqueName    *string                 `json:"unique_name,omitempty"`
	Url           *string                 `json:"url,omitempty"`
}
