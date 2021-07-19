/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateChannel'
type CreateChannelParams struct {
	// The chat channel's friendly name.
	ChatFriendlyName *string `json:"ChatFriendlyName,omitempty"`
	// The chat channel's unique name.
	ChatUniqueName *string `json:"ChatUniqueName,omitempty"`
	// The chat participant's friendly name.
	ChatUserFriendlyName *string `json:"ChatUserFriendlyName,omitempty"`
	// The SID of the Flex Flow.
	FlexFlowSid *string `json:"FlexFlowSid,omitempty"`
	// The `identity` value that uniquely identifies the new resource's chat User.
	Identity *string `json:"Identity,omitempty"`
	// Whether to create the channel as long-lived.
	LongLived *bool `json:"LongLived,omitempty"`
	// The pre-engagement data.
	PreEngagementData *string `json:"PreEngagementData,omitempty"`
	// The Target Contact Identity, for example the phone number of an SMS.
	Target *string `json:"Target,omitempty"`
	// The Task attributes to be added for the TaskRouter Task.
	TaskAttributes *string `json:"TaskAttributes,omitempty"`
	// The SID of the TaskRouter Task. Only valid when integration type is `task`. `null` for integration types `studio` & `external`
	TaskSid *string `json:"TaskSid,omitempty"`
}

func (params *CreateChannelParams) SetChatFriendlyName(ChatFriendlyName string) *CreateChannelParams {
	params.ChatFriendlyName = &ChatFriendlyName
	return params
}
func (params *CreateChannelParams) SetChatUniqueName(ChatUniqueName string) *CreateChannelParams {
	params.ChatUniqueName = &ChatUniqueName
	return params
}
func (params *CreateChannelParams) SetChatUserFriendlyName(ChatUserFriendlyName string) *CreateChannelParams {
	params.ChatUserFriendlyName = &ChatUserFriendlyName
	return params
}
func (params *CreateChannelParams) SetFlexFlowSid(FlexFlowSid string) *CreateChannelParams {
	params.FlexFlowSid = &FlexFlowSid
	return params
}
func (params *CreateChannelParams) SetIdentity(Identity string) *CreateChannelParams {
	params.Identity = &Identity
	return params
}
func (params *CreateChannelParams) SetLongLived(LongLived bool) *CreateChannelParams {
	params.LongLived = &LongLived
	return params
}
func (params *CreateChannelParams) SetPreEngagementData(PreEngagementData string) *CreateChannelParams {
	params.PreEngagementData = &PreEngagementData
	return params
}
func (params *CreateChannelParams) SetTarget(Target string) *CreateChannelParams {
	params.Target = &Target
	return params
}
func (params *CreateChannelParams) SetTaskAttributes(TaskAttributes string) *CreateChannelParams {
	params.TaskAttributes = &TaskAttributes
	return params
}
func (params *CreateChannelParams) SetTaskSid(TaskSid string) *CreateChannelParams {
	params.TaskSid = &TaskSid
	return params
}

func (c *ApiService) CreateChannel(params *CreateChannelParams) (*FlexV1Channel, error) {
	path := "/v1/Channels"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ChatFriendlyName != nil {
		data.Set("ChatFriendlyName", *params.ChatFriendlyName)
	}
	if params != nil && params.ChatUniqueName != nil {
		data.Set("ChatUniqueName", *params.ChatUniqueName)
	}
	if params != nil && params.ChatUserFriendlyName != nil {
		data.Set("ChatUserFriendlyName", *params.ChatUserFriendlyName)
	}
	if params != nil && params.FlexFlowSid != nil {
		data.Set("FlexFlowSid", *params.FlexFlowSid)
	}
	if params != nil && params.Identity != nil {
		data.Set("Identity", *params.Identity)
	}
	if params != nil && params.LongLived != nil {
		data.Set("LongLived", fmt.Sprint(*params.LongLived))
	}
	if params != nil && params.PreEngagementData != nil {
		data.Set("PreEngagementData", *params.PreEngagementData)
	}
	if params != nil && params.Target != nil {
		data.Set("Target", *params.Target)
	}
	if params != nil && params.TaskAttributes != nil {
		data.Set("TaskAttributes", *params.TaskAttributes)
	}
	if params != nil && params.TaskSid != nil {
		data.Set("TaskSid", *params.TaskSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1Channel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteChannel(Sid string) error {
	path := "/v1/Channels/{Sid}"
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

func (c *ApiService) FetchChannel(Sid string) (*FlexV1Channel, error) {
	path := "/v1/Channels/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1Channel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListChannel'
type ListChannelParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListChannelParams) SetPageSize(PageSize int) *ListChannelParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a single page of Channel records from the API. Request is executed immediately.
func (c *ApiService) PageChannel(params *ListChannelParams, pageToken string, pageNumber string) (*ListChannelResponse, error) {
	path := "/v1/Channels"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageToken != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListChannelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Channel records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListChannel(params *ListChannelParams, limit int) ([]FlexV1Channel, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageChannel(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []FlexV1Channel

	for response != nil {
		records = append(records, response.FlexChatChannels...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListChannelResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListChannelResponse)
	}

	return records, err
}

// Streams Channel records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamChannel(params *ListChannelParams, limit int) (chan FlexV1Channel, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageChannel(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan FlexV1Channel, 1)

	go func() {
		for response != nil {
			for item := range response.FlexChatChannels {
				channel <- response.FlexChatChannels[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListChannelResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListChannelResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListChannelResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListChannelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}