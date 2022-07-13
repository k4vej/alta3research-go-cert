FROM golang:1.16.15

WORKDIR /tmp

COPY main.go ./main.go

RUN go build -o /app/ main.go

WORKDIR /app
