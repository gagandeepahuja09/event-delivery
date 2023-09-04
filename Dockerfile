FROM golang:latest AS build

WORKDIR /app

# Multi-stage build: dependencies are much less likely to change
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd

CMD ["./main"]