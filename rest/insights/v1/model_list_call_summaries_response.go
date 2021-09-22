/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.3
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListCallSummariesResponse struct for ListCallSummariesResponse
type ListCallSummariesResponse struct {
	CallSummaries []InsightsV1CallSummaries        `json:"call_summaries,omitempty"`
	Meta          ListVideoRoomSummaryResponseMeta `json:"meta,omitempty"`
}
