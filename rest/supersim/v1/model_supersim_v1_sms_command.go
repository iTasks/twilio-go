/*
 * Twilio - Supersim
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

// SupersimV1SmsCommand struct for SupersimV1SmsCommand
type SupersimV1SmsCommand struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The direction of the SMS Command
	Direction *string `json:"direction,omitempty"`
	// The message body of the SMS Command sent to or from the SIM
	Payload *string `json:"payload,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The SID of the SIM that this SMS Command was sent to or from
	SimSid *string `json:"sim_sid,omitempty"`
	// The status of the SMS Command
	Status *string `json:"status,omitempty"`
	// The absolute URL of the SMS Command resource
	Url *string `json:"url,omitempty"`
}
