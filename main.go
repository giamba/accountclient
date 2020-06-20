package main

import (
	"fmt"

	"github.com/giamba/accountclient/client"
	guuid "github.com/google/uuid"
)

func main() {
	fmt.Println("Start!")

	testList()

	testListAll()

	id := guuid.New().String()

	testCreate(id)

	testDelete(id)

	testFetch()

	fmt.Println("End.")
}

func testList() {
	var apiEndpoint = "http://localhost:8080"
	var accessToken = "8fb95528-57c6-422e-9722-d2147bcba8ed"

	var client = client.NewClient(apiEndpoint, accessToken)

	response, err := client.List(1, 3)
	if err != nil {
		panic(err)
	}
	for _, acc := range response.Data {
		fmt.Println(acc.Id)
	}
}

func testListAll() {
	var apiEndpoint = "http://localhost:8080"
	var accessToken = "8fb95528-57c6-422e-9722-d2147bcba8ed"

	var client = client.NewClient(apiEndpoint, accessToken)

	response, err := client.ListAll()
	if err != nil {
		panic(err)
	}
	for _, acc := range response.Data {
		fmt.Println(acc.Id)
	}
}

func testCreate(id string) {
	var apiEndpoint = "http://localhost:8080"
	var accessToken = "8fb95528-57c6-422e-9722-d2147bcba8ed"

	c := client.NewClient(apiEndpoint, accessToken)

	var reqData = client.CreateRequest{
		Type:           "accounts",
		Id:             id,
		OrganisationId: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Attributes: client.CreateRequestAttributes{
			Country:      "GB",
			BaseCurrency: "GBP",
			BankId:       "400300",
			BankIdCode:   "GBDSC",
			Bic:          "NWBKGB22"},
	}

	response, err := c.Create(reqData)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)

}

func testDelete(id string) {
	var apiEndpoint = "http://localhost:8080"
	var accessToken = "8fb95528-57c6-422e-9722-d2147bcba8ed"

	c := client.NewClient(apiEndpoint, accessToken)

	response, err := c.Delete(id, "0")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}

func testFetch() {
	var apiEndpoint = "http://localhost:8080"
	var accessToken = "8fb95528-57c6-422e-9722-d2147bcba8ed"

	c := client.NewClient(apiEndpoint, accessToken)

	resonse, err := c.Fetch("aca3bc81-5453-4246-b3d0-d621e79664ac")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resonse)
}
