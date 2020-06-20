package client_test

import (
	"os"

	"bitbucket.org/giamba/accountclient/client"
	guuid "github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var apiUrl = getApiUrl()
var token = "8fb95528-57c6-422e-9722-d2147bcba8ed"

func getApiUrl() string {
	ip, exists := os.LookupEnv("API_GATEWAY")
	if exists {
		return "http://" + ip + ":8080"
	}
	return "http://localhost:8080"
}

//Create Tests
var _ = Describe("Given an account to create", func() {
	Context("When submit a new valid account", func() {
		newAccountId := guuid.New().String()
		c := client.NewClient(apiUrl, token)
		result, err := c.Create(client.CreateRequest{
			Type:           "accounts",
			Id:             newAccountId,
			OrganisationId: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Attributes: client.CreateRequestAttributes{
				Country:      "GB",
				BaseCurrency: "GBP",
				BankId:       "400300",
				BankIdCode:   "GBDSC",
				Bic:          "NWBKGB22"},
		})

		It("status code should be 201", func() {
			Expect(result.Status.StatusCode).To(Equal(201))
			Expect(err).NotTo(HaveOccurred())
		})

		It("the accountId shold be correct", func() {
			Expect(result.Data.Id).To(Equal(newAccountId))
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("When submit an existing account", func() {
		accountId := guuid.New().String()
		createAccountUtil(accountId)
		c := client.NewClient(apiUrl, token)

		result, err := c.Create(client.CreateRequest{
			Type:           "accounts",
			Id:             accountId,
			OrganisationId: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Attributes: client.CreateRequestAttributes{
				Country:      "GB",
				BaseCurrency: "GBP",
				BankId:       "400300",
				BankIdCode:   "GBDSC",
				Bic:          "NWBKGB22"},
		})

		It("status code should be 409 Conflict", func() {
			Expect(result.Status.StatusCode).To(Equal(409))
			Expect(err).NotTo(HaveOccurred())
		})

		It("the accountId shold be empty", func() {
			Expect(result.Data.Id).To(Equal(""))
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("When submit a not valid account id", func() {
		c := client.NewClient(apiUrl, token)
		result, err := c.Create(client.CreateRequest{
			Type:           "accounts",
			Id:             "xxxxxxxxxxxxxxxxxxxxxxx",
			OrganisationId: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Attributes: client.CreateRequestAttributes{
				Country:      "GB",
				BaseCurrency: "GBP",
				BankId:       "400300",
				BankIdCode:   "GBDSC",
				Bic:          "NWBKGB22"},
		})

		It("status code should be 400 Conflict", func() {
			Expect(result.Status.StatusCode).To(Equal(400))
			Expect(err).NotTo(HaveOccurred())
		})

		It("the accountId shold be empty", func() {
			Expect(result.Data.Id).To(Equal(""))
			Expect(err).NotTo(HaveOccurred())
		})
	})
})

//Delete tests
var _ = Describe("Given an account to delete", func() {
	Context("When submit a delete request", func() {
		accountId := guuid.New().String()
		createAccountUtil(accountId)

		c := client.NewClient(apiUrl, token)
		result, err := c.Delete(accountId, "0")

		It("status code should be 204 No Content", func() {
			Expect(result.StatusCode).To(Equal(204))
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("When submit an invalid delete request", func() {
		accountId := "xxxxxxxxxxxxxxxxxxxxxx"

		c := client.NewClient(apiUrl, token)
		result, err := c.Delete(accountId, "0")

		It("status code should be 400 Bad Request", func() {
			Expect(result.StatusCode).To(Equal(400))
			Expect(err).NotTo(HaveOccurred())
		})
	})
})

//Fetch tests
var _ = Describe("Given an account to retrieve", func() {
	Context("When submit a valid fetch request", func() {
		accountId := guuid.New().String()
		createAccountUtil(accountId)

		c := client.NewClient(apiUrl, token)
		result, err := c.Fetch(accountId)

		It("status code should be 200 OK", func() {
			Expect(result.Status.StatusCode).To(Equal(200))
			Expect(err).NotTo(HaveOccurred())
		})

		It("the accountId shold be correct", func() {
			Expect(result.Data.Id).To(Equal(accountId))
		})
	})
	Context("When fetch an invalid uuid", func() {
		accountId := "xxxxxxxxxxxxxxxxxxxxxx"

		c := client.NewClient(apiUrl, token)
		result, err := c.Fetch(accountId)

		It("status code should be 400 Bad Request", func() {
			Expect(result.Status.StatusCode).To(Equal(400))
			Expect(err).NotTo(HaveOccurred())
		})
	})
	Context("When fetch a non existing account", func() {
		accountId := guuid.New().String()

		c := client.NewClient(apiUrl, token)
		result, err := c.Fetch(accountId)

		It("status code should be 404 Not Found", func() {
			Expect(result.Status.StatusCode).To(Equal(404))
			Expect(err).NotTo(HaveOccurred())
		})
	})
})

//ListAll tests
var _ = Describe("Given accounts in database", func() {
	Context("When submit listAll and there are 3 accounts in the db", func() {
		deleteAllUtil()
		createAccountUtil(guuid.New().String())
		createAccountUtil(guuid.New().String())
		createAccountUtil(guuid.New().String())

		c := client.NewClient(apiUrl, token)
		result, err := c.ListAll()

		It("status code should be 200 OK", func() {
			Expect(result.Status.StatusCode).To(Equal(200))
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return 3 accounts", func() {
			Expect(len(result.Data)).To(Equal(3))
		})
	})

	Context("When submit listAll and there are no accounts in the db", func() {
		deleteAllUtil()

		c := client.NewClient(apiUrl, token)
		result, err := c.ListAll()

		It("status code should be 200 OK", func() {
			Expect(result.Status.StatusCode).To(Equal(200))
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return 0 accounts", func() {
			Expect(len(result.Data)).To(Equal(0))
		})
	})
})

//List tests
var _ = Describe("Given accounts in database", func() {
	Context("When submit list(0,2) and there are 3 accounts in the db", func() {
		deleteAllUtil()
		createAccountUtil(guuid.New().String())
		createAccountUtil(guuid.New().String())
		createAccountUtil(guuid.New().String())

		c := client.NewClient(apiUrl, token)
		result, err := c.List(0, 2)

		It("status code should be 200 OK", func() {
			Expect(result.Status.StatusCode).To(Equal(200))
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return 2 accounts", func() {
			Expect(len(result.Data)).To(Equal(2))
		})
	})
})

// Utilities functions
func createAccountUtil(id string) {
	c := client.NewClient(apiUrl, token)

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

	_, err := c.Create(reqData)

	if err != nil {
		panic(err)
	}
}

func deleteAccountUtil(id string) {
	c := client.NewClient(apiUrl, token)

	_, err := c.Delete(id, "0")

	if err != nil {
		panic(err)
	}
}

func countAccountsUtil() int {
	c := client.NewClient(apiUrl, token)
	result, _ := c.ListAll()
	return len(result.Data)
}

func deleteAllUtil() {
	c := client.NewClient(apiUrl, token)
	result, _ := c.ListAll()

	for _, acc := range result.Data {
		deleteAccountUtil(acc.Id)
	}
}
