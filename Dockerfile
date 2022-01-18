FROM golang:1.17.0-alpine

RUN apk update && apk add bash

COPY . /go/src/github.com/dirkarnez/golang-hello-world
WORKDIR /go/src/github.com/dirkarnez/golang-hello-world

RUN chmod +x /wait-for-it.sh

RUN go build -o app

EXPOSE 5000

ENTRYPOINT ["./wait-for-it.sh","mariadb:3306","--","./app"]
