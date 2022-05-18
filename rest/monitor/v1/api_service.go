/*
 * Twilio - Monitor
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	twilio "github.com/twilio/twilio-go/client"
)

type ApiService struct {
	baseURL        string
	requestHandler *twilio.RequestHandler
}

func NewApiService(requestHandler *twilio.RequestHandler) *ApiService {
	return &ApiService{
		requestHandler: requestHandler,
		baseURL:        "https://monitor.twilio.com",
	}
}

func NewApiServiceWithClient(client twilio.BaseClient) *ApiService {
	return NewApiService(twilio.NewRequestHandler(client))
}
