/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Messaging
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateAlphaSender'
type CreateAlphaSenderParams struct {
	// The Alphanumeric Sender ID string. Can be up to 11 characters long. Valid characters are A-Z, a-z, 0-9, space, hyphen `-`, plus `+`, underscore `_` and ampersand `&`. This value cannot contain only numbers.
	AlphaSender *string `json:"AlphaSender,omitempty"`
}

func (params *CreateAlphaSenderParams) SetAlphaSender(AlphaSender string) *CreateAlphaSenderParams {
	params.AlphaSender = &AlphaSender
	return params
}

//
func (c *ApiService) CreateAlphaSender(ServiceSid string, params *CreateAlphaSenderParams) (*MessagingV1AlphaSender, error) {
	path := "/v1/Services/{ServiceSid}/AlphaSenders"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AlphaSender != nil {
		data.Set("AlphaSender", *params.AlphaSender)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1AlphaSender{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//
func (c *ApiService) DeleteAlphaSender(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/AlphaSenders/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
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

//
func (c *ApiService) FetchAlphaSender(ServiceSid string, Sid string) (*MessagingV1AlphaSender, error) {
	path := "/v1/Services/{ServiceSid}/AlphaSenders/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1AlphaSender{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListAlphaSender'
type ListAlphaSenderParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListAlphaSenderParams) SetPageSize(PageSize int) *ListAlphaSenderParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListAlphaSenderParams) SetLimit(Limit int) *ListAlphaSenderParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of AlphaSender records from the API. Request is executed immediately.
func (c *ApiService) PageAlphaSender(ServiceSid string, params *ListAlphaSenderParams, pageToken, pageNumber string) (*ListAlphaSenderResponse, error) {
	path := "/v1/Services/{ServiceSid}/AlphaSenders"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

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

	ps := &ListAlphaSenderResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists AlphaSender records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListAlphaSender(ServiceSid string, params *ListAlphaSenderParams) ([]MessagingV1AlphaSender, error) {
	response, errors := c.StreamAlphaSender(ServiceSid, params)

	records := make([]MessagingV1AlphaSender, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams AlphaSender records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamAlphaSender(ServiceSid string, params *ListAlphaSenderParams) (chan MessagingV1AlphaSender, chan error) {
	if params == nil {
		params = &ListAlphaSenderParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan MessagingV1AlphaSender, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageAlphaSender(ServiceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamAlphaSender(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamAlphaSender(response *ListAlphaSenderResponse, params *ListAlphaSenderParams, recordChannel chan MessagingV1AlphaSender, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.AlphaSenders
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListAlphaSenderResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListAlphaSenderResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListAlphaSenderResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAlphaSenderResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
