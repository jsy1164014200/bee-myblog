FROM golang

MAINTAINER jasonjiang <gng@bingyan.net>
WORKDIR /go/src/app
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep && dep ensure && go build
