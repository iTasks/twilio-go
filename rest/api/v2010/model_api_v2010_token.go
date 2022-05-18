/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010Token struct for ApiV2010Token
type ApiV2010Token struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// An array representing the ephemeral credentials
	IceServers *[]ApiV2010AccountTokenIceServers `json:"ice_servers,omitempty"`
	// The temporary password used for authenticating
	Password *string `json:"password,omitempty"`
	// The duration in seconds the credentials are valid
	Ttl *string `json:"ttl,omitempty"`
	// The temporary username that uniquely identifies a Token
	Username *string `json:"username,omitempty"`
}
