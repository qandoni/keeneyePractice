FROM golang:1.25-alpine AS builder
RUN apk add --no-cache \
    gcc \
    g++ \
    musl-dev \
    pkgconf \
    librdkafka-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux \
    go build \
    -tags musl \
    -ldflags="-s -w" \
    -o myapp ./cmd/myapp

FROM alpine:3.22
RUN apk add --no-cache librdkafka
WORKDIR /app
COPY --from=builder /app/myapp .
CMD ["./myapp"]