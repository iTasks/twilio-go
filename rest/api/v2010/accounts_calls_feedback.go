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
)

// Optional parameters for the method 'FetchCallFeedback'
type FetchCallFeedbackParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchCallFeedbackParams) SetPathAccountSid(PathAccountSid string) *FetchCallFeedbackParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch a Feedback resource from a call
func (c *ApiService) FetchCallFeedback(CallSid string, params *FetchCallFeedbackParams) (*ApiV2010CallFeedback, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Feedback.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010CallFeedback{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateCallFeedback'
type UpdateCallFeedbackParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// One or more issues experienced during the call. The issues can be: `imperfect-audio`, `dropped-call`, `incorrect-caller-id`, `post-dial-delay`, `digits-not-captured`, `audio-latency`, `unsolicited-call`, or `one-way-audio`.
	Issue *[]string `json:"Issue,omitempty"`
	// The call quality expressed as an integer from `1` to `5` where `1` represents very poor call quality and `5` represents a perfect call.
	QualityScore *int `json:"QualityScore,omitempty"`
}

func (params *UpdateCallFeedbackParams) SetPathAccountSid(PathAccountSid string) *UpdateCallFeedbackParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateCallFeedbackParams) SetIssue(Issue []string) *UpdateCallFeedbackParams {
	params.Issue = &Issue
	return params
}
func (params *UpdateCallFeedbackParams) SetQualityScore(QualityScore int) *UpdateCallFeedbackParams {
	params.QualityScore = &QualityScore
	return params
}

// Update a Feedback resource for a call
func (c *ApiService) UpdateCallFeedback(CallSid string, params *UpdateCallFeedbackParams) (*ApiV2010CallFeedback, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Feedback.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Issue != nil {
		for _, item := range *params.Issue {
			data.Add("Issue", item)
		}
	}
	if params != nil && params.QualityScore != nil {
		data.Set("QualityScore", fmt.Sprint(*params.QualityScore))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010CallFeedback{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
