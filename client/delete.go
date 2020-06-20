package client

import (
	"fmt"
	"net/http"
)

type DeleteStatus struct {
	StatusCode        int    `json:"status_code"`
	StatusDescription string `json:"status_descr"`
	ErrorMessage      string `json:"error_message"`
}

func (c *Client) Delete(id string, version string) (*DeleteStatus, error) {
	resp, err := c.delete(apiPath + "/" + id + "?version=" + version)
	return getDeleteResponse(c, resp, err)
}

func getDeleteResponse(c *Client, resp *http.Response, err error) (*DeleteStatus, error) {
	if err != nil {
		return nil, err
	}

	var result DeleteStatus

	if resp.StatusCode != 204 {
		result.StatusCode = resp.StatusCode
		result.StatusDescription = resp.Status

		var errResult *ErrorObject
		if err := c.decodeJSON(resp, &errResult); err != nil {
			return nil, fmt.Errorf("Could not decode JSON error response: %v", err)
		}
		result.ErrorMessage = errResult.ErrorMessage
		return &result, nil
	}

	result.StatusCode = resp.StatusCode
	result.StatusDescription = resp.Status
	return &result, nil
}
