/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/url"
	"strings"
)

// Optional parameters for the method 'CreateUserDefinedMessage'
type CreateUserDefinedMessageParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created User Defined Message.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// A unique string value to identify API call. This should be a unique string value per API call and can be a randomly generated.
	Content *string `json:"Content,omitempty"`
	// A unique string value to identify API call. This should be a unique string value per API call and can be a randomly generated.
	IdempotencyKey *string `json:"IdempotencyKey,omitempty"`
}

func (params *CreateUserDefinedMessageParams) SetPathAccountSid(PathAccountSid string) *CreateUserDefinedMessageParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateUserDefinedMessageParams) SetContent(Content string) *CreateUserDefinedMessageParams {
	params.Content = &Content
	return params
}
func (params *CreateUserDefinedMessageParams) SetIdempotencyKey(IdempotencyKey string) *CreateUserDefinedMessageParams {
	params.IdempotencyKey = &IdempotencyKey
	return params
}

// Create a new User Defined Message for the given call sid.
func (c *ApiService) CreateUserDefinedMessage(CallSid string, params *CreateUserDefinedMessageParams) (*ApiV2010UserDefinedMessage, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/UserDefinedMessages.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Content != nil {
		data.Set("Content", *params.Content)
	}
	if params != nil && params.IdempotencyKey != nil {
		data.Set("IdempotencyKey", *params.IdempotencyKey)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010UserDefinedMessage{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
