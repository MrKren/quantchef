FROM golang:1.18-bullseye

RUN go install github.com/beego/bee/v2@latest
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

EXPOSE 8080

WORKDIR /app/src

CMD ["bee", "run"]
