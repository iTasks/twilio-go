/*
 * Twilio - Video
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

// Returns a single Track resource represented by &#x60;track_sid&#x60;.  Note: This is one resource with the Video API that requires a SID, be Track Name on the subscriber side is not guaranteed to be unique.
func (c *ApiService) FetchRoomParticipantSubscribedTrack(RoomSid string, ParticipantSid string, Sid string) (*VideoV1RoomParticipantSubscribedTrack, error) {
	path := "/v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/SubscribedTracks/{Sid}"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
	path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1RoomParticipantSubscribedTrack{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRoomParticipantSubscribedTrack'
type ListRoomParticipantSubscribedTrackParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListRoomParticipantSubscribedTrackParams) SetPageSize(PageSize int) *ListRoomParticipantSubscribedTrackParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListRoomParticipantSubscribedTrackParams) SetLimit(Limit int) *ListRoomParticipantSubscribedTrackParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of RoomParticipantSubscribedTrack records from the API. Request is executed immediately.
func (c *ApiService) PageRoomParticipantSubscribedTrack(RoomSid string, ParticipantSid string, params *ListRoomParticipantSubscribedTrackParams, pageToken, pageNumber string) (*ListRoomParticipantSubscribedTrackResponse, error) {
	path := "/v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/SubscribedTracks"

	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
	path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)

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

	ps := &ListRoomParticipantSubscribedTrackResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists RoomParticipantSubscribedTrack records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListRoomParticipantSubscribedTrack(RoomSid string, ParticipantSid string, params *ListRoomParticipantSubscribedTrackParams) ([]VideoV1RoomParticipantSubscribedTrack, error) {
	response, errors := c.StreamRoomParticipantSubscribedTrack(RoomSid, ParticipantSid, params)

	records := make([]VideoV1RoomParticipantSubscribedTrack, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams RoomParticipantSubscribedTrack records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamRoomParticipantSubscribedTrack(RoomSid string, ParticipantSid string, params *ListRoomParticipantSubscribedTrackParams) (chan VideoV1RoomParticipantSubscribedTrack, chan error) {
	if params == nil {
		params = &ListRoomParticipantSubscribedTrackParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan VideoV1RoomParticipantSubscribedTrack, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageRoomParticipantSubscribedTrack(RoomSid, ParticipantSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamRoomParticipantSubscribedTrack(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamRoomParticipantSubscribedTrack(response *ListRoomParticipantSubscribedTrackResponse, params *ListRoomParticipantSubscribedTrackParams, recordChannel chan VideoV1RoomParticipantSubscribedTrack, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.SubscribedTracks
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListRoomParticipantSubscribedTrackResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListRoomParticipantSubscribedTrackResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListRoomParticipantSubscribedTrackResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRoomParticipantSubscribedTrackResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
