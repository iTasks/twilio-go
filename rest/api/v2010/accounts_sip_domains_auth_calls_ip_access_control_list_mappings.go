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

// Optional parameters for the method 'CreateSipAuthCallsIpAccessControlListMapping'
type CreateSipAuthCallsIpAccessControlListMappingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The SID of the IpAccessControlList resource to map to the SIP domain.
	IpAccessControlListSid *string `json:"IpAccessControlListSid,omitempty"`
}

func (params *CreateSipAuthCallsIpAccessControlListMappingParams) SetPathAccountSid(PathAccountSid string) *CreateSipAuthCallsIpAccessControlListMappingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateSipAuthCallsIpAccessControlListMappingParams) SetIpAccessControlListSid(IpAccessControlListSid string) *CreateSipAuthCallsIpAccessControlListMappingParams {
	params.IpAccessControlListSid = &IpAccessControlListSid
	return params
}

// Create a new IP Access Control List mapping
func (c *ApiService) CreateSipAuthCallsIpAccessControlListMapping(DomainSid string, params *CreateSipAuthCallsIpAccessControlListMappingParams) (*ApiV2010SipAuthCallsIpAccessControlListMapping, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/IpAccessControlListMappings.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"DomainSid"+"}", DomainSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.IpAccessControlListSid != nil {
		data.Set("IpAccessControlListSid", *params.IpAccessControlListSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010SipAuthCallsIpAccessControlListMapping{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteSipAuthCallsIpAccessControlListMapping'
type DeleteSipAuthCallsIpAccessControlListMappingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IpAccessControlListMapping resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteSipAuthCallsIpAccessControlListMappingParams) SetPathAccountSid(PathAccountSid string) *DeleteSipAuthCallsIpAccessControlListMappingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete an IP Access Control List mapping from the requested domain
func (c *ApiService) DeleteSipAuthCallsIpAccessControlListMapping(DomainSid string, Sid string, params *DeleteSipAuthCallsIpAccessControlListMappingParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/IpAccessControlListMappings/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"DomainSid"+"}", DomainSid, -1)
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

// Optional parameters for the method 'FetchSipAuthCallsIpAccessControlListMapping'
type FetchSipAuthCallsIpAccessControlListMappingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IpAccessControlListMapping resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchSipAuthCallsIpAccessControlListMappingParams) SetPathAccountSid(PathAccountSid string) *FetchSipAuthCallsIpAccessControlListMappingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch a specific instance of an IP Access Control List mapping
func (c *ApiService) FetchSipAuthCallsIpAccessControlListMapping(DomainSid string, Sid string, params *FetchSipAuthCallsIpAccessControlListMappingParams) (*ApiV2010SipAuthCallsIpAccessControlListMapping, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/IpAccessControlListMappings/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"DomainSid"+"}", DomainSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010SipAuthCallsIpAccessControlListMapping{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSipAuthCallsIpAccessControlListMapping'
type ListSipAuthCallsIpAccessControlListMappingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IpAccessControlListMapping resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSipAuthCallsIpAccessControlListMappingParams) SetPathAccountSid(PathAccountSid string) *ListSipAuthCallsIpAccessControlListMappingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListSipAuthCallsIpAccessControlListMappingParams) SetPageSize(PageSize int) *ListSipAuthCallsIpAccessControlListMappingParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSipAuthCallsIpAccessControlListMappingParams) SetLimit(Limit int) *ListSipAuthCallsIpAccessControlListMappingParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SipAuthCallsIpAccessControlListMapping records from the API. Request is executed immediately.
func (c *ApiService) PageSipAuthCallsIpAccessControlListMapping(DomainSid string, params *ListSipAuthCallsIpAccessControlListMappingParams, pageToken, pageNumber string) (*ListSipAuthCallsIpAccessControlListMappingResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/IpAccessControlListMappings.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"DomainSid"+"}", DomainSid, -1)

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

	ps := &ListSipAuthCallsIpAccessControlListMappingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SipAuthCallsIpAccessControlListMapping records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSipAuthCallsIpAccessControlListMapping(DomainSid string, params *ListSipAuthCallsIpAccessControlListMappingParams) ([]ApiV2010SipAuthCallsIpAccessControlListMapping, error) {
	if params == nil {
		params = &ListSipAuthCallsIpAccessControlListMappingParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSipAuthCallsIpAccessControlListMapping(DomainSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010SipAuthCallsIpAccessControlListMapping

	for response != nil {
		records = append(records, response.Contents...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSipAuthCallsIpAccessControlListMappingResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListSipAuthCallsIpAccessControlListMappingResponse)
	}

	return records, err
}

// Streams SipAuthCallsIpAccessControlListMapping records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSipAuthCallsIpAccessControlListMapping(DomainSid string, params *ListSipAuthCallsIpAccessControlListMappingParams) (chan ApiV2010SipAuthCallsIpAccessControlListMapping, error) {
	if params == nil {
		params = &ListSipAuthCallsIpAccessControlListMappingParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSipAuthCallsIpAccessControlListMapping(DomainSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010SipAuthCallsIpAccessControlListMapping, 1)

	go func() {
		for response != nil {
			for item := range response.Contents {
				channel <- response.Contents[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSipAuthCallsIpAccessControlListMappingResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListSipAuthCallsIpAccessControlListMappingResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListSipAuthCallsIpAccessControlListMappingResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSipAuthCallsIpAccessControlListMappingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
