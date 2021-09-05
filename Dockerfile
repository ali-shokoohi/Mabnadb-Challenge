FROM golang:alpine

ENV GO111MODULE auto

WORKDIR /go/src/gitlab.com/shokoohi/mabnadp-challenge

COPY . .

RUN go get -d -v ./...

RUN go install gitlab.com/shokoohi/mabnadp-challenge

ENTRYPOINT /go/bin/mabnadp-challenge

EXPOSE 8000
