# syntax=docker/dockerfile:1

# Tahap build
FROM golang:1.24

WORKDIR /app

COPY go.* .
RUN go mod tidy
RUN go mod verify
RUN go mod download

COPY . .

CMD ["go", "run", "cmd/app/main.go"]

# RUN go build -o /build/app
# CMD ["./app"]
