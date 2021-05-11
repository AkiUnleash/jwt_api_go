FROM golang:1.15.2-alpine

RUN apk update && apk add git

RUN mkdir /go/src/api
WORKDIR /go/src/api

ADD ./api /go/src/api

# RUN ./api/main