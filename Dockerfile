FROM golang:1.16.15

ARG APP_NAME

WORKDIR /build

COPY go.mod go.sum ./
COPY main.go ./main.go
COPY cmd ./cmd

RUN go build -o /app/$APP_NAME main.go

WORKDIR /app
