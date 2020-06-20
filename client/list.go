package client

import (
	"net/http"
	"strconv"
)

type ListAttributes struct {
	AlternativeBankAccountNames string `json:"alternative_bank_account_names"`
	BankId                      string `json:"bank_id"`
	BankIdCode                  string `json:"bank_id_code"`
	BaseCurrency                string `json:"base_currency"`
	Bic                         string `json:"bic"`
	Country                     string `json:"country"`
}

type ListData struct {
	Attributes     ListAttributes `json:"attributes"`
	CreatedOn      string         `json:"created_on"`
	Id             string         `json:"id"`
	ModifiedOn     string         `json:"modified_on"`
	OrganisationId string         `json:"organisation_id"`
	Type           string         `json:"type"`
	Version        uint           `json:"version"`
}

type ListResponseLinks struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Next  string `json:"next"`
	Self  string `json:"self"`
}

type ListStatus struct {
	StatusCode        int    `json:"status_code"`
	StatusDescription string `json:"status_descr"`
	ErrorMessage      string `json:"error_message"`
}

type ListResponse struct {
	Data   []ListData        `json:"data"`
	Links  ListResponseLinks `json:"links"`
	Status ListStatus        `json:"status"`
}

func (c *Client) List(pageNumber int, pageSize int) (*ListResponse, error) {
	pNumberStr := strconv.Itoa(pageNumber)
	pSizeStr := strconv.Itoa(pageSize)

	resp, err := c.get(apiPath + "?page[number]=" + pNumberStr + "&page[size]=" + pSizeStr)
	return getListResponse(c, resp, err)
}

func (c *Client) ListAll() (*ListResponse, error) {
	resp, err := c.get(apiPath)
	return getListResponse(c, resp, err)
}

func getListResponse(c *Client, resp *http.Response, err error) (*ListResponse, error) {
	if err != nil {
		return nil, err
	}

	var result ListResponse
	result.Status.StatusCode = resp.StatusCode
	result.Status.StatusDescription = resp.Status
	return &result, c.decodeJSON(resp, &result)
}
