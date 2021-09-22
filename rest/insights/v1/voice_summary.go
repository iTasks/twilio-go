/*
 * Twilio - Insights
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
	"net/url"

	"strings"
)

// Optional parameters for the method 'FetchSummary'
type FetchSummaryParams struct {
	//
	ProcessingState *string `json:"ProcessingState,omitempty"`
}

func (params *FetchSummaryParams) SetProcessingState(ProcessingState string) *FetchSummaryParams {
	params.ProcessingState = &ProcessingState
	return params
}

func (c *ApiService) FetchSummary(CallSid string, params *FetchSummaryParams) (*InsightsV1Summary, error) {
	path := "/v1/Voice/{CallSid}/Summary"
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ProcessingState != nil {
		data.Set("ProcessingState", *params.ProcessingState)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &InsightsV1Summary{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
