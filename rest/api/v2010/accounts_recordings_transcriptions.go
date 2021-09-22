/*
 * Twilio - Api
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

// Optional parameters for the method 'DeleteRecordingTranscription'
type DeleteRecordingTranscriptionParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteRecordingTranscriptionParams) SetPathAccountSid(PathAccountSid string) *DeleteRecordingTranscriptionParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

func (c *ApiService) DeleteRecordingTranscription(RecordingSid string, Sid string, params *DeleteRecordingTranscriptionParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Recordings/{RecordingSid}/Transcriptions/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"RecordingSid"+"}", RecordingSid, -1)
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

// Optional parameters for the method 'FetchRecordingTranscription'
type FetchRecordingTranscriptionParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchRecordingTranscriptionParams) SetPathAccountSid(PathAccountSid string) *FetchRecordingTranscriptionParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

func (c *ApiService) FetchRecordingTranscription(RecordingSid string, Sid string, params *FetchRecordingTranscriptionParams) (*ApiV2010RecordingTranscription, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Recordings/{RecordingSid}/Transcriptions/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"RecordingSid"+"}", RecordingSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010RecordingTranscription{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRecordingTranscription'
type ListRecordingTranscriptionParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListRecordingTranscriptionParams) SetPathAccountSid(PathAccountSid string) *ListRecordingTranscriptionParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListRecordingTranscriptionParams) SetPageSize(PageSize int) *ListRecordingTranscriptionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListRecordingTranscriptionParams) SetLimit(Limit int) *ListRecordingTranscriptionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of RecordingTranscription records from the API. Request is executed immediately.
func (c *ApiService) PageRecordingTranscription(RecordingSid string, params *ListRecordingTranscriptionParams, pageToken, pageNumber string) (*ListRecordingTranscriptionResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Recordings/{RecordingSid}/Transcriptions.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"RecordingSid"+"}", RecordingSid, -1)

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

	ps := &ListRecordingTranscriptionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists RecordingTranscription records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListRecordingTranscription(RecordingSid string, params *ListRecordingTranscriptionParams) ([]ApiV2010RecordingTranscription, error) {
	if params == nil {
		params = &ListRecordingTranscriptionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageRecordingTranscription(RecordingSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010RecordingTranscription

	for response != nil {
		records = append(records, response.Transcriptions...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListRecordingTranscriptionResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListRecordingTranscriptionResponse)
	}

	return records, err
}

// Streams RecordingTranscription records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamRecordingTranscription(RecordingSid string, params *ListRecordingTranscriptionParams) (chan ApiV2010RecordingTranscription, error) {
	if params == nil {
		params = &ListRecordingTranscriptionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageRecordingTranscription(RecordingSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010RecordingTranscription, 1)

	go func() {
		for response != nil {
			for item := range response.Transcriptions {
				channel <- response.Transcriptions[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListRecordingTranscriptionResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListRecordingTranscriptionResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListRecordingTranscriptionResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRecordingTranscriptionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
