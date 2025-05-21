# syntax=docker/dockerfile:1

# Tahap build
FROM golang:1.24 as builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
WORKDIR /build/cmd/app
RUN go build -o /build/app

# Tahap produksi
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /build/app .
CMD ["./app"]
