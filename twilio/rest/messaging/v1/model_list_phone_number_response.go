/*
 * Twilio - Messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListPhoneNumberResponse struct for ListPhoneNumberResponse
type ListPhoneNumberResponse struct {
	Meta         ListServiceResponseMeta         `json:"meta,omitempty"`
	PhoneNumbers []MessagingV1ServicePhoneNumber `json:"phone_numbers,omitempty"`
}
