# syntax=docker/dockerfile:1
FROM golang:alpine3.23
WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o /usr/local/bin/server ./cmd/server

CMD ["server"]

