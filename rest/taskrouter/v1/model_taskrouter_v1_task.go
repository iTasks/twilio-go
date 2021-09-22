/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.3
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// TaskrouterV1Task struct for TaskrouterV1Task
type TaskrouterV1Task struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// An object that contains the addon data for all installed addons
	Addons *string `json:"addons,omitempty"`
	// The number of seconds since the Task was created
	Age *int `json:"age,omitempty"`
	// The current status of the Task's assignment
	AssignmentStatus *string `json:"assignment_status,omitempty"`
	// The JSON string with custom attributes of the work
	Attributes *string `json:"attributes,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The URLs of related resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// Retrieve the list of all Tasks in the Workspace with the specified priority
	Priority *int `json:"priority,omitempty"`
	// The reason the Task was canceled or completed
	Reason *string `json:"reason,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The SID of the TaskChannel
	TaskChannelSid *string `json:"task_channel_sid,omitempty"`
	// The unique name of the TaskChannel
	TaskChannelUniqueName *string `json:"task_channel_unique_name,omitempty"`
	// The ISO 8601 date and time in GMT when the Task entered the TaskQueue.
	TaskQueueEnteredDate *time.Time `json:"task_queue_entered_date,omitempty"`
	// The friendly name of the TaskQueue
	TaskQueueFriendlyName *string `json:"task_queue_friendly_name,omitempty"`
	// The SID of the TaskQueue
	TaskQueueSid *string `json:"task_queue_sid,omitempty"`
	// The amount of time in seconds that the Task can live before being assigned
	Timeout *int `json:"timeout,omitempty"`
	// The absolute URL of the Task resource
	Url *string `json:"url,omitempty"`
	// The friendly name of the Workflow that is controlling the Task
	WorkflowFriendlyName *string `json:"workflow_friendly_name,omitempty"`
	// The SID of the Workflow that is controlling the Task
	WorkflowSid *string `json:"workflow_sid,omitempty"`
	// The SID of the Workspace that contains the Task
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
}
