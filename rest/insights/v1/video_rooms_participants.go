/*
 * Twilio - Insights
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

// Get Video Log Analyzer data for a Room Participant.
func (c *ApiService) FetchVideoParticipantSummary(RoomSid string, ParticipantSid string) (*InsightsV1VideoRoomSummaryVideoParticipantSummary, error) {
	path := "/v1/Video/Rooms/{RoomSid}/Participants/{ParticipantSid}"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
	path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &InsightsV1VideoRoomSummaryVideoParticipantSummary{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListVideoParticipantSummary'
type ListVideoParticipantSummaryParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListVideoParticipantSummaryParams) SetPageSize(PageSize int) *ListVideoParticipantSummaryParams {
	params.PageSize = &PageSize
	return params
}

// Get a list of room participants.
func (c *ApiService) ListVideoParticipantSummary(RoomSid string, params *ListVideoParticipantSummaryParams) (*ListVideoParticipantSummaryResponse, error) {
	path := "/v1/Video/Rooms/{RoomSid}/Participants"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListVideoParticipantSummaryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
