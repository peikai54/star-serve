FROM golang:latest AS builder
WORKDIR /app
COPY . /app
ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn,direct
RUN go build main.go


FROM alpine AS runner
WORKDIR /app
COPY --from=builder /app/main /app

RUN chmod +x main
EXPOSE 8000

ENTRYPOINT ["./main"]