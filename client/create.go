package client

import (
	"fmt"
	"net/http"
)

//Request objects
type CreateRequestAttributes struct {
	Country      string `json:"country"`
	BaseCurrency string `json:"base_currency"`
	BankId       string `json:"bank_id"`
	BankIdCode   string `json:"bank_id_code"`
	Bic          string `json:"bic"`
}

type CreateRequest struct {
	Type           string                  `json:"type"`
	Id             string                  `json:"id"`
	OrganisationId string                  `json:"organisation_id"`
	Attributes     CreateRequestAttributes `json:"attributes"`
}

//Response objects
type CreateResponseAttributes struct {
	AlternativeBankAccountNames string `json:"alternative_bank_account_names"`
	BankId                      string `json:"bank_id"`
	BankIdCode                  string `json:"bank_id_code"`
	BaseCurrency                string `json:"base_currency"`
	Bic                         string `json:"bic"`
	Country                     string `json:"country"`
}

type CreateResponseData struct {
	Attributes     CreateResponseAttributes `json:"attributes"`
	CreatedOn      string                   `json:"created_on"`
	Id             string                   `json:"id"`
	ModifiedOn     string                   `json:"modified_on"`
	OrganisationId string                   `json:"organisation_id"`
	Type           string                   `json:"type"`
	Version        uint                     `json:"version"`
}

type CreateResponseLinks struct {
	Self string `json:"self"`
}

type CreateResponse struct {
	Data   CreateResponseData  `json:"data"`
	Links  CreateResponseLinks `json:"links"`
	Status Status              `json:"status"`
}

func (c *Client) Create(req CreateRequest) (*CreateResponse, error) {
	data := make(map[string]CreateRequest)
	data["data"] = req
	resp, err := c.post(apiPath, data, nil)
	return getCreateResponse(c, resp, err)
}

func getCreateResponse(c *Client, resp *http.Response, err error) (*CreateResponse, error) {
	if err != nil {
		return nil, err
	}
	var result CreateResponse

	if resp.StatusCode != 201 {
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
