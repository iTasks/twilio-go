package twilio

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// TaskRouterActivity describes the current status of a Worker,
// which determines whether they are eligible to receive task assignments.
// refer: https://www.twilio.com/docs/taskrouter/api/activity
type TaskRouterActivity struct {
	AccountSID   *string    `json:"account_sid"`
	Available    *bool      `json:"available"`
	DateCreated  *time.Time `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated"`
	FriendlyName *string    `json:"friendly_name"`
	SID          *string    `json:"sid"`
	WorkspaceSID *string    `json:"workspace_sid"`
	URL          *string    `json:"url"`
}

// TaskRouterActivityParams activity params to create/update activity.
type TaskRouterActivityParams struct {
	Available    *bool   `form:",omitempty"`
	FriendlyName *string `form:",omitempty"`
}

// TaskRouterActivityList struct to parse response of activity read.
type TaskRouterActivityList struct {
	Activities []*TaskRouterActivity `json:"activities"`
	Meta       *Meta                 `json:"meta,omitempty"`
}

// TaskRouterActivityQueryParams query params to read workspaces.
type TaskRouterActivityQueryParams struct {
	FriendlyName *string `form:",omitempty"`
	Available    *string `form:",omitempty"`
	PageSize     *int    `form:",omitempty"`
}

// taskRouterActivityClient is the entrypoint for activity CRUD.
type taskRouterActivityClient struct {
	baseURL string
	client  *Twilio
}

// newTaskRouterActivityClient constructs a new Activity Client.
func newTaskRouterActivityClient(client *Twilio) *taskRouterActivityClient {
	c := new(taskRouterActivityClient)
	c.client = client
	c.baseURL = fmt.Sprintf("https://taskrouter.%s/v1", client.BaseURL)

	return c
}

// Create creates activity with the given the config.
func (c *taskRouterActivityClient) Create(
	workspaceSID string,
	activityParams *TaskRouterActivityParams,
) (*TaskRouterActivity, error) {
	url := c.url(fmt.Sprintf("/Workspaces/%s/Activities", workspaceSID))

	if len(*activityParams.FriendlyName) == 0 {
		return nil, errors.New("friendlyname is required in activityparams")
	}

	resp, err := c.client.Post(url, activityParams)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	activity := &TaskRouterActivity{}

	if err = json.NewDecoder(resp.Body).Decode(activity); err != nil {
		return nil, err
	}

	return activity, nil
}

// Fetch fetches activity for the given activity SID.
func (c *taskRouterActivityClient) Fetch(workspaceSID string, sid string) (*TaskRouterActivity, error) {
	url := c.url(fmt.Sprintf("/Workspaces/%s/Activities/%s", workspaceSID, sid))
	resp, err := c.client.Get(url, nil)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	activity := &TaskRouterActivity{}

	if err = json.NewDecoder(resp.Body).Decode(activity); err != nil {
		return nil, err
	}

	return activity, nil
}

// Read returns all existing activities for a workspace.
func (c *taskRouterActivityClient) Read(workspaceSID string) (*TaskRouterActivityList, error) {
	url := c.url(fmt.Sprintf("/Workspaces/%s/Activities", workspaceSID))

	resp, err := c.client.Get(url, nil)

	if err != nil {
		return nil, err
	}

	activities := &TaskRouterActivityList{}

	if err = json.NewDecoder(resp.Body).Decode(&activities); err != nil {
		return nil, err
	}

	return activities, nil
}

// Update updates activity with given config.
func (c *taskRouterActivityClient) Update(
	workspaceSID string,
	sid string,
	activityParams *TaskRouterActivityParams,
) (*TaskRouterActivity, error) {
	url := c.url(fmt.Sprintf("/Workspaces/%s/Activities/%s", workspaceSID, sid))

	resp, err := c.client.Post(url, activityParams)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	activity := &TaskRouterActivity{}

	if err = json.NewDecoder(resp.Body).Decode(activity); err != nil {
		return nil, err
	}

	return activity, nil
}

// Delete deletes workflow for given SID.
func (c *taskRouterActivityClient) Delete(workspaceSID string, sid string) error {
	url := c.url(fmt.Sprintf("/Workspaces/%s/Activities/%s", workspaceSID, sid))

	resp, err := c.client.Delete(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}

func (c *taskRouterActivityClient) url(path string) string {
	if c.client.defaultbaseURL != nil {
		return *c.client.defaultbaseURL + path
	}

	return c.baseURL + path
}