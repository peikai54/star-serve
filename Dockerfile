FROM golang:latest AS builder

COPY . /app
WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct
    
RUN go build main.go

RUN chmod +x main

EXPOSE 8000

ENTRYPOINT ["./main"]