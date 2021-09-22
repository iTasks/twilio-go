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

// ListMessageResponse struct for ListMessageResponse
type ListMessageResponse struct {
	Messages []IpMessagingV2Message     `json:"messages,omitempty"`
	Meta     ListCredentialResponseMeta `json:"meta,omitempty"`
}
