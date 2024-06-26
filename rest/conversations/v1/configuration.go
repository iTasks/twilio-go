/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Conversations
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/url"
)

// Fetch the global configuration of conversations on your account
func (c *ApiService) FetchConfiguration() (*ConversationsV1Configuration, error) {
	path := "/v1/Configuration"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1Configuration{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateConfiguration'
type UpdateConfigurationParams struct {
	// The SID of the default [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to use when creating a conversation.
	DefaultChatServiceSid *string `json:"DefaultChatServiceSid,omitempty"`
	// The SID of the default [Messaging Service](https://www.twilio.com/docs/messaging/api/service-resource) to use when creating a conversation.
	DefaultMessagingServiceSid *string `json:"DefaultMessagingServiceSid,omitempty"`
	// Default ISO8601 duration when conversation will be switched to `inactive` state. Minimum value for this timer is 1 minute.
	DefaultInactiveTimer *string `json:"DefaultInactiveTimer,omitempty"`
	// Default ISO8601 duration when conversation will be switched to `closed` state. Minimum value for this timer is 10 minutes.
	DefaultClosedTimer *string `json:"DefaultClosedTimer,omitempty"`
}

func (params *UpdateConfigurationParams) SetDefaultChatServiceSid(DefaultChatServiceSid string) *UpdateConfigurationParams {
	params.DefaultChatServiceSid = &DefaultChatServiceSid
	return params
}
func (params *UpdateConfigurationParams) SetDefaultMessagingServiceSid(DefaultMessagingServiceSid string) *UpdateConfigurationParams {
	params.DefaultMessagingServiceSid = &DefaultMessagingServiceSid
	return params
}
func (params *UpdateConfigurationParams) SetDefaultInactiveTimer(DefaultInactiveTimer string) *UpdateConfigurationParams {
	params.DefaultInactiveTimer = &DefaultInactiveTimer
	return params
}
func (params *UpdateConfigurationParams) SetDefaultClosedTimer(DefaultClosedTimer string) *UpdateConfigurationParams {
	params.DefaultClosedTimer = &DefaultClosedTimer
	return params
}

// Update the global configuration of conversations on your account
func (c *ApiService) UpdateConfiguration(params *UpdateConfigurationParams) (*ConversationsV1Configuration, error) {
	path := "/v1/Configuration"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.DefaultChatServiceSid != nil {
		data.Set("DefaultChatServiceSid", *params.DefaultChatServiceSid)
	}
	if params != nil && params.DefaultMessagingServiceSid != nil {
		data.Set("DefaultMessagingServiceSid", *params.DefaultMessagingServiceSid)
	}
	if params != nil && params.DefaultInactiveTimer != nil {
		data.Set("DefaultInactiveTimer", *params.DefaultInactiveTimer)
	}
	if params != nil && params.DefaultClosedTimer != nil {
		data.Set("DefaultClosedTimer", *params.DefaultClosedTimer)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1Configuration{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
