FROM golang:1.24.4-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git upx

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/bot

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o /bot
RUN upx --best --lzma /bot

FROM scratch
COPY --from=builder /bot /bot
ENTRYPOINT ["/bot"]
