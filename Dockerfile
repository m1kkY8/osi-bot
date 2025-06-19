FROM golang:1.24.4-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/bot

RUN go build -o /bot

FROM alpine:latest

RUN adduser -D appuser

WORKDIR /home/appuser

COPY --from=builder /bot .

COPY .env .

USER appuser

ENTRYPOINT ["./bot"]
