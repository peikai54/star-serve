FROM golang:1.15-alpine AS builder

COPY . /app
WORKDIR /app

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

RUN go run main.go

EXPOSE 8000
