FROM golang:1.24

# Install golang-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/ && \
    chmod +x /usr/local/bin/migrate

WORKDIR /app

COPY go.* .
RUN go mod tidy
RUN go mod verify
RUN go mod download

COPY . .

CMD ["go", "run", "cmd/app/main.go"]

# RUN go build -o /build/app
# CMD ["./app"]
