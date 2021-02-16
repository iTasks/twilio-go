/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// TaskrouterV1Workspace struct for TaskrouterV1Workspace
type TaskrouterV1Workspace struct {
	AccountSid string `json:"AccountSid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	DefaultActivityName string `json:"DefaultActivityName,omitempty"`
	DefaultActivitySid string `json:"DefaultActivitySid,omitempty"`
	EventCallbackUrl string `json:"EventCallbackUrl,omitempty"`
	EventsFilter string `json:"EventsFilter,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	MultiTaskEnabled bool `json:"MultiTaskEnabled,omitempty"`
	PrioritizeQueueOrder QueueOrder `json:"PrioritizeQueueOrder,omitempty"`
	Sid string `json:"Sid,omitempty"`
	TimeoutActivityName string `json:"TimeoutActivityName,omitempty"`
	TimeoutActivitySid string `json:"TimeoutActivitySid,omitempty"`
	Url string `json:"Url,omitempty"`
}