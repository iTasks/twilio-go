/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// IpMessagingV2ServiceUserUserChannel struct for IpMessagingV2ServiceUserUserChannel
type IpMessagingV2ServiceUserUserChannel struct {
	AccountSid               *string                 `json:"account_sid,omitempty"`
	ChannelSid               *string                 `json:"channel_sid,omitempty"`
	LastConsumedMessageIndex *int32                  `json:"last_consumed_message_index,omitempty"`
	Links                    *map[string]interface{} `json:"links,omitempty"`
	MemberSid                *string                 `json:"member_sid,omitempty"`
	NotificationLevel        *string                 `json:"notification_level,omitempty"`
	ServiceSid               *string                 `json:"service_sid,omitempty"`
	Status                   *string                 `json:"status,omitempty"`
	UnreadMessagesCount      *int32                  `json:"unread_messages_count,omitempty"`
	Url                      *string                 `json:"url,omitempty"`
	UserSid                  *string                 `json:"user_sid,omitempty"`
}
