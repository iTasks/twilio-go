/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.3
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListEventTypeResponse struct for ListEventTypeResponse
type ListEventTypeResponse struct {
	Meta  ListSchemaVersionResponseMeta `json:"meta,omitempty"`
	Types []EventsV1EventType           `json:"types,omitempty"`
}
