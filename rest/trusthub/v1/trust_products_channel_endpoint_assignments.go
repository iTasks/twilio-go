/*
 * Twilio - Trusthub
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

// Optional parameters for the method 'CreateTrustProductChannelEndpointAssignment'
type CreateTrustProductChannelEndpointAssignmentParams struct {
	// The SID of an channel endpoint
	ChannelEndpointSid *string `json:"ChannelEndpointSid,omitempty"`
	// The type of channel endpoint. eg: phone-number
	ChannelEndpointType *string `json:"ChannelEndpointType,omitempty"`
}

func (params *CreateTrustProductChannelEndpointAssignmentParams) SetChannelEndpointSid(ChannelEndpointSid string) *CreateTrustProductChannelEndpointAssignmentParams {
	params.ChannelEndpointSid = &ChannelEndpointSid
	return params
}
func (params *CreateTrustProductChannelEndpointAssignmentParams) SetChannelEndpointType(ChannelEndpointType string) *CreateTrustProductChannelEndpointAssignmentParams {
	params.ChannelEndpointType = &ChannelEndpointType
	return params
}

// Create a new Assigned Item.
func (c *ApiService) CreateTrustProductChannelEndpointAssignment(TrustProductSid string, params *CreateTrustProductChannelEndpointAssignmentParams) (*TrusthubV1TrustProductTrustProductChannelEndpointAssignment, error) {
	path := "/v1/TrustProducts/{TrustProductSid}/ChannelEndpointAssignments"
	path = strings.Replace(path, "{"+"TrustProductSid"+"}", TrustProductSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ChannelEndpointSid != nil {
		data.Set("ChannelEndpointSid", *params.ChannelEndpointSid)
	}
	if params != nil && params.ChannelEndpointType != nil {
		data.Set("ChannelEndpointType", *params.ChannelEndpointType)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1TrustProductTrustProductChannelEndpointAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove an Assignment Item Instance.
func (c *ApiService) DeleteTrustProductChannelEndpointAssignment(TrustProductSid string, Sid string) error {
	path := "/v1/TrustProducts/{TrustProductSid}/ChannelEndpointAssignments/{Sid}"
	path = strings.Replace(path, "{"+"TrustProductSid"+"}", TrustProductSid, -1)
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

// Fetch specific Assigned Item Instance.
func (c *ApiService) FetchTrustProductChannelEndpointAssignment(TrustProductSid string, Sid string) (*TrusthubV1TrustProductTrustProductChannelEndpointAssignment, error) {
	path := "/v1/TrustProducts/{TrustProductSid}/ChannelEndpointAssignments/{Sid}"
	path = strings.Replace(path, "{"+"TrustProductSid"+"}", TrustProductSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1TrustProductTrustProductChannelEndpointAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListTrustProductChannelEndpointAssignment'
type ListTrustProductChannelEndpointAssignmentParams struct {
	// The SID of an channel endpoint
	ChannelEndpointSid *string `json:"ChannelEndpointSid,omitempty"`
	// comma separated list of channel endpoint sids
	ChannelEndpointSids *string `json:"ChannelEndpointSids,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListTrustProductChannelEndpointAssignmentParams) SetChannelEndpointSid(ChannelEndpointSid string) *ListTrustProductChannelEndpointAssignmentParams {
	params.ChannelEndpointSid = &ChannelEndpointSid
	return params
}
func (params *ListTrustProductChannelEndpointAssignmentParams) SetChannelEndpointSids(ChannelEndpointSids string) *ListTrustProductChannelEndpointAssignmentParams {
	params.ChannelEndpointSids = &ChannelEndpointSids
	return params
}
func (params *ListTrustProductChannelEndpointAssignmentParams) SetPageSize(PageSize int) *ListTrustProductChannelEndpointAssignmentParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Assigned Items for an account.
func (c *ApiService) ListTrustProductChannelEndpointAssignment(TrustProductSid string, params *ListTrustProductChannelEndpointAssignmentParams) (*ListTrustProductChannelEndpointAssignmentResponse, error) {
	path := "/v1/TrustProducts/{TrustProductSid}/ChannelEndpointAssignments"
	path = strings.Replace(path, "{"+"TrustProductSid"+"}", TrustProductSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ChannelEndpointSid != nil {
		data.Set("ChannelEndpointSid", *params.ChannelEndpointSid)
	}
	if params != nil && params.ChannelEndpointSids != nil {
		data.Set("ChannelEndpointSids", *params.ChannelEndpointSids)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTrustProductChannelEndpointAssignmentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
