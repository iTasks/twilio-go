/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListLogResponse struct for ListLogResponse
type ListLogResponse struct {
	Logs []ServerlessV1ServiceEnvironmentLog `json:"logs,omitempty"`
	Meta ListServiceResponseMeta             `json:"meta,omitempty"`
}
