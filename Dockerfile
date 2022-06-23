FROM golang:1.18-bullseye

RUN go install github.com/beego/bee/v2@latest

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

EXPOSE 8080

WORKDIR /app
COPY . ./app
WORKDIR /app/src

CMD ["bee", "run", "quantchef"]
