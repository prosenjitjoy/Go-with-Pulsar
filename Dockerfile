FROM golang:alpine AS builder
WORKDIR /src
COPY . .
RUN go build -o producer producer/main.go
RUN go build -o consumer consumer/main.go

FROM alpine:latest AS producer
WORKDIR /app
COPY --from=builder /src/producer/main producer
COPY config.yaml .
ENTRYPOINT [ "/app/producer" ]

FROM alpine:latest AS consumer
WORKDIR /app
COPY --from=builder /src/consumer/main consumer
COPY config.yaml .
ENTRYPOINT [ "/app/consumer" ]