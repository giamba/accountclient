package client

import (
	"fmt"
	"net/http"
)

type FetchAttributes struct {
	AlternativeBankAccountNames string `json:"alternative_bank_account_names"`
	BankId                      string `json:"bank_id"`
	BankIdCode                  string `json:"bank_id_code"`
	BaseCurrency                string `json:"base_currency"`
	Bic                         string `json:"bic"`
	Country                     string `json:"country"`
}

type FetchData struct {
	Attributes     FetchAttributes `json:"attributes"`
	CreatedOn      string          `json:"created_on"`
	Id             string          `json:"id"`
	ModifiedOn     string          `json:"modified_on"`
	OrganisationId string          `json:"organisation_id"`
	Type           string          `json:"type"`
	Version        uint            `json:"version"`
}

type FetchLinks struct {
	Self string `json:"self"`
}

type FetchStatus struct {
	StatusCode        int    `json:"status_code"`
	StatusDescription string `json:"status_descr"`
	ErrorMessage      string `json:"error_message"`
}

type FetchResponse struct {
	Data   FetchData  `json:"data"`
	Links  FetchLinks `json:"links"`
	Status Status     `json:"status"`
}

func (c *Client) Fetch(id string) (*FetchResponse, error) {
	resp, err := c.get(apiPath + "/" + id)
	if err != nil {
		return nil, err
	}

	return getFetchResponse(c, resp, err)
}

func getFetchResponse(c *Client, resp *http.Response, err error) (*FetchResponse, error) {
	if err != nil {
		return nil, err
	}

	var result FetchResponse

	if resp.StatusCode != 200 {
		result.Status.StatusCode = resp.StatusCode
		result.Status.StatusDescription = resp.Status

		var errResult *ErrorObject
		if err := c.decodeJSON(resp, &errResult); err != nil {
			return nil, fmt.Errorf("Could not decode JSON error response: %v", err)
		}
		result.Status.ErrorMessage = errResult.ErrorMessage
		return &result, nil
	}

	result.Status.StatusCode = resp.StatusCode
	result.Status.StatusDescription = resp.Status
	return &result, c.decodeJSON(resp, &result)

}
