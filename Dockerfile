FROM golang:1.17.0-alpine

RUN apk update && apk add bash

COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

COPY . /go/src/github.com/dirkarnez/golang-hello-world
WORKDIR /go/src/github.com/dirkarnez/golang-hello-world

RUN go build -o app

EXPOSE 5000

ENTRYPOINT ["./wait-for-it.sh","mariadb:3306","--","./app"]
