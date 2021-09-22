/*
 * Twilio - Messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.3
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListPhoneNumberResponse struct for ListPhoneNumberResponse
type ListPhoneNumberResponse struct {
	Meta         ListServiceResponseMeta  `json:"meta,omitempty"`
	PhoneNumbers []MessagingV1PhoneNumber `json:"phone_numbers,omitempty"`
}
