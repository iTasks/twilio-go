/*
 * Twilio - Api
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

// Optional parameters for the method 'FetchCallNotification'
type FetchCallNotificationParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call Notification resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchCallNotificationParams) SetPathAccountSid(PathAccountSid string) *FetchCallNotificationParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

func (c *ApiService) FetchCallNotification(CallSid string, Sid string, params *FetchCallNotificationParams) (*ApiV2010AccountCallCallNotificationInstance, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Notifications/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountCallCallNotificationInstance{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListCallNotification'
type ListCallNotificationParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call Notification resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// Only read notifications of the specified log level. Can be:  `0` to read only ERROR notifications or `1` to read only WARNING notifications. By default, all notifications are read.
	Log *int `json:"Log,omitempty"`
	// Only show notifications for the specified date, formatted as `YYYY-MM-DD`. You can also specify an inequality, such as `<=YYYY-MM-DD` for messages logged at or before midnight on a date, or `>=YYYY-MM-DD` for messages logged at or after midnight on a date.
	MessageDate *string `json:"MessageDate,omitempty"`
	// Only show notifications for the specified date, formatted as `YYYY-MM-DD`. You can also specify an inequality, such as `<=YYYY-MM-DD` for messages logged at or before midnight on a date, or `>=YYYY-MM-DD` for messages logged at or after midnight on a date.
	MessageDateBefore *string `json:"MessageDate&lt;,omitempty"`
	// Only show notifications for the specified date, formatted as `YYYY-MM-DD`. You can also specify an inequality, such as `<=YYYY-MM-DD` for messages logged at or before midnight on a date, or `>=YYYY-MM-DD` for messages logged at or after midnight on a date.
	MessageDateAfter *string `json:"MessageDate&gt;,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListCallNotificationParams) SetPathAccountSid(PathAccountSid string) *ListCallNotificationParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListCallNotificationParams) SetLog(Log int) *ListCallNotificationParams {
	params.Log = &Log
	return params
}
func (params *ListCallNotificationParams) SetMessageDate(MessageDate string) *ListCallNotificationParams {
	params.MessageDate = &MessageDate
	return params
}
func (params *ListCallNotificationParams) SetMessageDateBefore(MessageDateBefore string) *ListCallNotificationParams {
	params.MessageDateBefore = &MessageDateBefore
	return params
}
func (params *ListCallNotificationParams) SetMessageDateAfter(MessageDateAfter string) *ListCallNotificationParams {
	params.MessageDateAfter = &MessageDateAfter
	return params
}
func (params *ListCallNotificationParams) SetPageSize(PageSize int) *ListCallNotificationParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a single page of CallNotification records from the API. Request is executed immediately.
func (c *ApiService) PageCallNotification(CallSid string, params *ListCallNotificationParams, pageToken string, pageNumber string) (*ListCallNotificationResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Notifications.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Log != nil {
		data.Set("Log", fmt.Sprint(*params.Log))
	}
	if params != nil && params.MessageDate != nil {
		data.Set("MessageDate", fmt.Sprint(*params.MessageDate))
	}
	if params != nil && params.MessageDateBefore != nil {
		data.Set("MessageDate<", fmt.Sprint(*params.MessageDateBefore))
	}
	if params != nil && params.MessageDateAfter != nil {
		data.Set("MessageDate>", fmt.Sprint(*params.MessageDateAfter))
	}
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

	ps := &ListCallNotificationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists CallNotification records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListCallNotification(CallSid string, params *ListCallNotificationParams, limit int) ([]ApiV2010AccountCallCallNotification, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageCallNotification(CallSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010AccountCallCallNotification

	for response != nil {
		records = append(records, response.Notifications...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListCallNotificationResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListCallNotificationResponse)
	}

	return records, err
}

// Streams CallNotification records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamCallNotification(CallSid string, params *ListCallNotificationParams, limit int) (chan ApiV2010AccountCallCallNotification, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageCallNotification(CallSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010AccountCallCallNotification, 1)

	go func() {
		for response != nil {
			for item := range response.Notifications {
				channel <- response.Notifications[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListCallNotificationResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListCallNotificationResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListCallNotificationResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCallNotificationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}