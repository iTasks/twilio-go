/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// InsightsV1Event struct for InsightsV1Event
type InsightsV1Event struct {
	AccountSid  *string      `json:"account_sid,omitempty"`
	CallSid     *string      `json:"call_sid,omitempty"`
	CarrierEdge *interface{} `json:"carrier_edge,omitempty"`
	ClientEdge  *interface{} `json:"client_edge,omitempty"`
	Edge        *string      `json:"edge,omitempty"`
	Group       *string      `json:"group,omitempty"`
	Level       *string      `json:"level,omitempty"`
	Name        *string      `json:"name,omitempty"`
	SdkEdge     *interface{} `json:"sdk_edge,omitempty"`
	SipEdge     *interface{} `json:"sip_edge,omitempty"`
	Timestamp   *string      `json:"timestamp,omitempty"`
}
