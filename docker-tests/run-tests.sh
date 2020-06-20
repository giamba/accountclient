#!/bin/bash
set -e
set -m

check_api(){
    curl --silent http://$API_GATEWAY:8080 > dummy.txt  
    echo $?
}

until [[ $(check_api) = 0 ]]; do
    >&2 echo "...Waiting for api to be available" 
    sleep 1
done

>&2 echo "Api ready!" 
 
if [ -d /go/src/github.com/giamba/accountclient ]; then rm -Rf /go/src/github.com/giamba/accountclient; fi
 
>&2 echo "Downloading src from github.com/giamba/accountclient..."
go get github.com/giamba/accountclient
cd /go/src/github.com/giamba/accountclient

>&2 echo "Downloading dependencies"
go get -t ./...

>&2 echo "Running tests..."
cd client/
>&2 go test -ginkgo.v
