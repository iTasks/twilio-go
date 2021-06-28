/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
)

// Optional parameters for the method 'CreateWebhook'
type CreateWebhookParams struct {
	// The list of space-separated events that this Webhook will subscribe to.
	Events *string `json:"Events,omitempty"`
	// An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the `sid` in the URL path to address the resource. This value must be unique and 64 characters or less in length.
	UniqueName *string `json:"UniqueName,omitempty"`
	// The method to be used when calling the webhook's URL.
	WebhookMethod *string `json:"WebhookMethod,omitempty"`
	// The URL associated with this Webhook.
	WebhookUrl *string `json:"WebhookUrl,omitempty"`
}

func (params *CreateWebhookParams) SetEvents(Events string) *CreateWebhookParams {
	params.Events = &Events
	return params
}
func (params *CreateWebhookParams) SetUniqueName(UniqueName string) *CreateWebhookParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *CreateWebhookParams) SetWebhookMethod(WebhookMethod string) *CreateWebhookParams {
	params.WebhookMethod = &WebhookMethod
	return params
}
func (params *CreateWebhookParams) SetWebhookUrl(WebhookUrl string) *CreateWebhookParams {
	params.WebhookUrl = &WebhookUrl
	return params
}

func (c *ApiService) CreateWebhook(AssistantSid string, params *CreateWebhookParams) (*AutopilotV1AssistantWebhook, error) {
	path := "/v1/Assistants/{AssistantSid}/Webhooks"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Events != nil {
		data.Set("Events", *params.Events)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.WebhookMethod != nil {
		data.Set("WebhookMethod", *params.WebhookMethod)
	}
	if params != nil && params.WebhookUrl != nil {
		data.Set("WebhookUrl", *params.WebhookUrl)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1AssistantWebhook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteWebhook(AssistantSid string, Sid string) error {
	path := "/v1/Assistants/{AssistantSid}/Webhooks/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ApiService) FetchWebhook(AssistantSid string, Sid string) (*AutopilotV1AssistantWebhook, error) {
	path := "/v1/Assistants/{AssistantSid}/Webhooks/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1AssistantWebhook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListWebhook'
type ListWebhookParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListWebhookParams) SetPageSize(PageSize int) *ListWebhookParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListWebhook(AssistantSid string, params *ListWebhookParams) (*ListWebhookResponse, error) {
	path := "/v1/Assistants/{AssistantSid}/Webhooks"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListWebhookResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateWebhook'
type UpdateWebhookParams struct {
	// The list of space-separated events that this Webhook will subscribe to.
	Events *string `json:"Events,omitempty"`
	// An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the `sid` in the URL path to address the resource. This value must be unique and 64 characters or less in length.
	UniqueName *string `json:"UniqueName,omitempty"`
	// The method to be used when calling the webhook's URL.
	WebhookMethod *string `json:"WebhookMethod,omitempty"`
	// The URL associated with this Webhook.
	WebhookUrl *string `json:"WebhookUrl,omitempty"`
}

func (params *UpdateWebhookParams) SetEvents(Events string) *UpdateWebhookParams {
	params.Events = &Events
	return params
}
func (params *UpdateWebhookParams) SetUniqueName(UniqueName string) *UpdateWebhookParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *UpdateWebhookParams) SetWebhookMethod(WebhookMethod string) *UpdateWebhookParams {
	params.WebhookMethod = &WebhookMethod
	return params
}
func (params *UpdateWebhookParams) SetWebhookUrl(WebhookUrl string) *UpdateWebhookParams {
	params.WebhookUrl = &WebhookUrl
	return params
}

func (c *ApiService) UpdateWebhook(AssistantSid string, Sid string, params *UpdateWebhookParams) (*AutopilotV1AssistantWebhook, error) {
	path := "/v1/Assistants/{AssistantSid}/Webhooks/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Events != nil {
		data.Set("Events", *params.Events)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.WebhookMethod != nil {
		data.Set("WebhookMethod", *params.WebhookMethod)
	}
	if params != nil && params.WebhookUrl != nil {
		data.Set("WebhookUrl", *params.WebhookUrl)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1AssistantWebhook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
