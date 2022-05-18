/*
 * Twilio - Media
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// MediaV1MediaRecording struct for MediaV1MediaRecording
type MediaV1MediaRecording struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The bitrate of the media
	Bitrate *int `json:"bitrate,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The duration of the MediaRecording
	Duration *int `json:"duration,omitempty"`
	// The format of the MediaRecording
	Format *string `json:"format,omitempty"`
	// The URLs of related resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// The SID of the MediaProcessor
	ProcessorSid *string `json:"processor_sid,omitempty"`
	// The dimensions of the video image in pixels
	Resolution *string `json:"resolution,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The size of the recording
	Size *int `json:"size,omitempty"`
	// The SID of the resource that generated the original media
	SourceSid *string `json:"source_sid,omitempty"`
	// The status of the MediaRecording
	Status *string `json:"status,omitempty"`
	// The URL to which Twilio will send MediaRecording event updates
	StatusCallback *string `json:"status_callback,omitempty"`
	// The HTTP method Twilio should use to call the `status_callback` URL
	StatusCallbackMethod *string `json:"status_callback_method,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
