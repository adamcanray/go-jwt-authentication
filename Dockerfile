# syntax=docker/dockerfile:1

FROM golang:1.18.2-alpine

WORKDIR /app

ADD . /app

RUN go build -o /go-jwt-authentication

EXPOSE 8080

CMD [ "/go-jwt-authentication" ]