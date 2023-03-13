FROM golang:1.15-alpine AS builder

COPY . /app
WORKDIR /app

RUN go run main.go

EXPOSE 8000
