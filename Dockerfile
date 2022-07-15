FROM golang:1.16.15

ARG APP_NAME

WORKDIR /build

COPY go.mod go.sum ./
COPY alta3research-gocert01.go ./alta3research-gocert01.go
COPY cmd ./cmd

RUN go build -o /app/$APP_NAME alta3research-gocert01.go

WORKDIR /app
