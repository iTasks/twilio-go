/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// EventsV1Schema struct for EventsV1Schema
type EventsV1Schema struct {
	// Schema Identifier.
	Id *string `json:"id,omitempty"`
	// The date that the last schema version was created.
	LastCreated *time.Time `json:"last_created,omitempty"`
	// Last schema version.
	LastVersion *int32 `json:"last_version,omitempty"`
	// Nested resource URLs.
	Links *map[string]interface{} `json:"links,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
}
