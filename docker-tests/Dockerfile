FROM golang:latest

RUN apt-get update && apt-get install -y dos2unix

COPY run-tests.sh .

ENV API_GATEWAY=172.28.5.254

RUN dos2unix run-tests.sh && chmod +x run-tests.sh 

CMD [ "./run-tests.sh" ]