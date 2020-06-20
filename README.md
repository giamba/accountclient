
## Intro
Hi, I'm Giambattista Pieretti and this is my code assignment.
This is my first project in Go.

## Pull 
`go get github.com/giamba/accountclient`

## Run on docker
```
1) docker network create \
  --driver=bridge \
  --subnet=172.28.0.0/16 \
  --ip-range=172.28.5.0/24 \
  --gateway=172.28.5.254 \
  test-network

2) docker-compose up --build 
```
## Run locally
```
1) Comment client-tests service section in docker-compose.yml

2) docker network create \
  --driver=bridge \
  --subnet=172.28.0.0/16 \
  --ip-range=172.28.5.0/24 \
  --gateway=172.28.5.254 \
  test-network

3) docker-compose up --build  

4) Run the tests: cd client/ && go test -ginkgo.v
```

## Usage
```go   
package main
import (
	"fmt"
	"bitbucket.org/giamba/accountclient/client"
)

func main() {

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
```
See main.go for other examples

## Possible Improvements
- Docker network investigate for a easier solution
- Optimize docker test image, size and time 
- Improve Ginkgo tests readability and understand better the framework
- Improve test quality and test coverage
- Better packages and test organization
- Add unit tests  
- Improve my Go skills
