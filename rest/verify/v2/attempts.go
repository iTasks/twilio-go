/*
 * Twilio - Verify
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
	"time"

	"github.com/twilio/twilio-go/client"
)

// Fetch a specific verification attempt.
func (c *ApiService) FetchVerificationAttempt(Sid string) (*VerifyV2VerificationAttempt, error) {
	path := "/v2/Attempts/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2VerificationAttempt{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListVerificationAttempt'
type ListVerificationAttemptParams struct {
	// Datetime filter used to query Verification Attempts created after this datetime.
	DateCreatedAfter *time.Time `json:"DateCreatedAfter,omitempty"`
	// Datetime filter used to query Verification Attempts created before this datetime.
	DateCreatedBefore *time.Time `json:"DateCreatedBefore,omitempty"`
	// Destination of a verification. Depending on the type of channel, it could be a phone number in E.164 format or an email address.
	ChannelDataTo *string `json:"ChannelData.To,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListVerificationAttemptParams) SetDateCreatedAfter(DateCreatedAfter time.Time) *ListVerificationAttemptParams {
	params.DateCreatedAfter = &DateCreatedAfter
	return params
}
func (params *ListVerificationAttemptParams) SetDateCreatedBefore(DateCreatedBefore time.Time) *ListVerificationAttemptParams {
	params.DateCreatedBefore = &DateCreatedBefore
	return params
}
func (params *ListVerificationAttemptParams) SetChannelDataTo(ChannelDataTo string) *ListVerificationAttemptParams {
	params.ChannelDataTo = &ChannelDataTo
	return params
}
func (params *ListVerificationAttemptParams) SetPageSize(PageSize int) *ListVerificationAttemptParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListVerificationAttemptParams) SetLimit(Limit int) *ListVerificationAttemptParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of VerificationAttempt records from the API. Request is executed immediately.
func (c *ApiService) PageVerificationAttempt(params *ListVerificationAttemptParams, pageToken, pageNumber string) (*ListVerificationAttemptResponse, error) {
	path := "/v2/Attempts"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreatedAfter", fmt.Sprint((*params.DateCreatedAfter).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedBefore != nil {
		data.Set("DateCreatedBefore", fmt.Sprint((*params.DateCreatedBefore).Format(time.RFC3339)))
	}
	if params != nil && params.ChannelDataTo != nil {
		data.Set("ChannelData.To", *params.ChannelDataTo)
	}
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

	ps := &ListVerificationAttemptResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists VerificationAttempt records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListVerificationAttempt(params *ListVerificationAttemptParams) ([]VerifyV2VerificationAttempt, error) {
	if params == nil {
		params = &ListVerificationAttemptParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageVerificationAttempt(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []VerifyV2VerificationAttempt

	for response != nil {
		records = append(records, response.Attempts...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListVerificationAttemptResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListVerificationAttemptResponse)
	}

	return records, err
}

// Streams VerificationAttempt records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamVerificationAttempt(params *ListVerificationAttemptParams) (chan VerifyV2VerificationAttempt, error) {
	if params == nil {
		params = &ListVerificationAttemptParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageVerificationAttempt(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan VerifyV2VerificationAttempt, 1)

	go func() {
		for response != nil {
			for item := range response.Attempts {
				channel <- response.Attempts[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListVerificationAttemptResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListVerificationAttemptResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListVerificationAttemptResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListVerificationAttemptResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
