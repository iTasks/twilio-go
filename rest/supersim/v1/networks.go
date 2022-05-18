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
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Fetch a Network resource.
func (c *ApiService) FetchNetwork(Sid string) (*SupersimV1Network, error) {
	path := "/v1/Networks/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1Network{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListNetwork'
type ListNetworkParams struct {
	// The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the Network resources to read.
	IsoCountry *string `json:"IsoCountry,omitempty"`
	// The 'mobile country code' of a country. Network resources with this `mcc` in their `identifiers` will be read.
	Mcc *string `json:"Mcc,omitempty"`
	// The 'mobile network code' of a mobile operator network. Network resources with this `mnc` in their `identifiers` will be read.
	Mnc *string `json:"Mnc,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListNetworkParams) SetIsoCountry(IsoCountry string) *ListNetworkParams {
	params.IsoCountry = &IsoCountry
	return params
}
func (params *ListNetworkParams) SetMcc(Mcc string) *ListNetworkParams {
	params.Mcc = &Mcc
	return params
}
func (params *ListNetworkParams) SetMnc(Mnc string) *ListNetworkParams {
	params.Mnc = &Mnc
	return params
}
func (params *ListNetworkParams) SetPageSize(PageSize int) *ListNetworkParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListNetworkParams) SetLimit(Limit int) *ListNetworkParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Network records from the API. Request is executed immediately.
func (c *ApiService) PageNetwork(params *ListNetworkParams, pageToken, pageNumber string) (*ListNetworkResponse, error) {
	path := "/v1/Networks"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.IsoCountry != nil {
		data.Set("IsoCountry", *params.IsoCountry)
	}
	if params != nil && params.Mcc != nil {
		data.Set("Mcc", *params.Mcc)
	}
	if params != nil && params.Mnc != nil {
		data.Set("Mnc", *params.Mnc)
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

	ps := &ListNetworkResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Network records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListNetwork(params *ListNetworkParams) ([]SupersimV1Network, error) {
	response, errors := c.StreamNetwork(params)

	records := make([]SupersimV1Network, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Network records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamNetwork(params *ListNetworkParams) (chan SupersimV1Network, chan error) {
	if params == nil {
		params = &ListNetworkParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan SupersimV1Network, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageNetwork(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamNetwork(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamNetwork(response *ListNetworkResponse, params *ListNetworkParams, recordChannel chan SupersimV1Network, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Networks
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListNetworkResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListNetworkResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListNetworkResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListNetworkResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
