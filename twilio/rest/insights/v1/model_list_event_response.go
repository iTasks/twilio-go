/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListEventResponse struct for ListEventResponse
type ListEventResponse struct {
	Events []InsightsV1CallEvent            `json:"events,omitempty"`
	Meta   ListVideoRoomSummaryResponseMeta `json:"meta,omitempty"`
}
