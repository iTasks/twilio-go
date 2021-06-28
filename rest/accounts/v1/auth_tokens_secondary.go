/*
 * Twilio - Accounts
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
	"net/url"
)

// Create a new secondary Auth Token
func (c *ApiService) CreateSecondaryAuthToken() (*AccountsV1SecondaryAuthToken, error) {
	path := "/v1/AuthTokens/Secondary"

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AccountsV1SecondaryAuthToken{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete the secondary Auth Token from your account
func (c *ApiService) DeleteSecondaryAuthToken() error {
	path := "/v1/AuthTokens/Secondary"

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
