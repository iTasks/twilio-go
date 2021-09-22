/*
 * Twilio - Bulkexports
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.3
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

// Optional parameters for the method 'CreateExportCustomJob'
type CreateExportCustomJobParams struct {
	// The optional email to send the completion notification to. You can set both webhook, and email, or one or the other. If you set neither, the job will run but you will have to query to determine your job's status.
	Email *string `json:"Email,omitempty"`
	// The end day for the custom export specified as a string in the format of yyyy-mm-dd. End day is inclusive and must be 2 days earlier than the current UTC day.
	EndDay *string `json:"EndDay,omitempty"`
	// The friendly name specified when creating the job
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The start day for the custom export specified as a string in the format of yyyy-mm-dd
	StartDay *string `json:"StartDay,omitempty"`
	// This is the method used to call the webhook on completion of the job. If this is supplied, `WebhookUrl` must also be supplied.
	WebhookMethod *string `json:"WebhookMethod,omitempty"`
	// The optional webhook url called on completion of the job. If this is supplied, `WebhookMethod` must also be supplied. If you set neither webhook nor email, you will have to check your job's status manually.
	WebhookUrl *string `json:"WebhookUrl,omitempty"`
}

func (params *CreateExportCustomJobParams) SetEmail(Email string) *CreateExportCustomJobParams {
	params.Email = &Email
	return params
}
func (params *CreateExportCustomJobParams) SetEndDay(EndDay string) *CreateExportCustomJobParams {
	params.EndDay = &EndDay
	return params
}
func (params *CreateExportCustomJobParams) SetFriendlyName(FriendlyName string) *CreateExportCustomJobParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateExportCustomJobParams) SetStartDay(StartDay string) *CreateExportCustomJobParams {
	params.StartDay = &StartDay
	return params
}
func (params *CreateExportCustomJobParams) SetWebhookMethod(WebhookMethod string) *CreateExportCustomJobParams {
	params.WebhookMethod = &WebhookMethod
	return params
}
func (params *CreateExportCustomJobParams) SetWebhookUrl(WebhookUrl string) *CreateExportCustomJobParams {
	params.WebhookUrl = &WebhookUrl
	return params
}

func (c *ApiService) CreateExportCustomJob(ResourceType string, params *CreateExportCustomJobParams) (*BulkexportsV1ExportCustomJob, error) {
	path := "/v1/Exports/{ResourceType}/Jobs"
	path = strings.Replace(path, "{"+"ResourceType"+"}", ResourceType, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Email != nil {
		data.Set("Email", *params.Email)
	}
	if params != nil && params.EndDay != nil {
		data.Set("EndDay", *params.EndDay)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.StartDay != nil {
		data.Set("StartDay", *params.StartDay)
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

	ps := &BulkexportsV1ExportCustomJob{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteJob(JobSid string) error {
	path := "/v1/Exports/Jobs/{JobSid}"
	path = strings.Replace(path, "{"+"JobSid"+"}", JobSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ApiService) FetchJob(JobSid string) (*BulkexportsV1Job, error) {
	path := "/v1/Exports/Jobs/{JobSid}"
	path = strings.Replace(path, "{"+"JobSid"+"}", JobSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &BulkexportsV1Job{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListExportCustomJob'
type ListExportCustomJobParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListExportCustomJobParams) SetPageSize(PageSize int) *ListExportCustomJobParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListExportCustomJobParams) SetLimit(Limit int) *ListExportCustomJobParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ExportCustomJob records from the API. Request is executed immediately.
func (c *ApiService) PageExportCustomJob(ResourceType string, params *ListExportCustomJobParams, pageToken, pageNumber string) (*ListExportCustomJobResponse, error) {
	path := "/v1/Exports/{ResourceType}/Jobs"

	path = strings.Replace(path, "{"+"ResourceType"+"}", ResourceType, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListExportCustomJobResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ExportCustomJob records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListExportCustomJob(ResourceType string, params *ListExportCustomJobParams) ([]BulkexportsV1ExportCustomJob, error) {
	if params == nil {
		params = &ListExportCustomJobParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageExportCustomJob(ResourceType, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []BulkexportsV1ExportCustomJob

	for response != nil {
		records = append(records, response.Jobs...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListExportCustomJobResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListExportCustomJobResponse)
	}

	return records, err
}

// Streams ExportCustomJob records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamExportCustomJob(ResourceType string, params *ListExportCustomJobParams) (chan BulkexportsV1ExportCustomJob, error) {
	if params == nil {
		params = &ListExportCustomJobParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageExportCustomJob(ResourceType, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan BulkexportsV1ExportCustomJob, 1)

	go func() {
		for response != nil {
			for item := range response.Jobs {
				channel <- response.Jobs[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListExportCustomJobResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListExportCustomJobResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListExportCustomJobResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListExportCustomJobResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
