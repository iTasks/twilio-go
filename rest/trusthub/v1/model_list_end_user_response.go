/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.3
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListEndUserResponse struct for ListEndUserResponse
type ListEndUserResponse struct {
	Meta    ListCustomerProfileResponseMeta `json:"meta,omitempty"`
	Results []TrusthubV1EndUser             `json:"results,omitempty"`
}
