FROM golang:latest

COPY . /app
WORKDIR /app

RUN go run main.go

EXPOSE 8000
