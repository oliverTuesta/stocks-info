# Minimal multi-stage Dockerfile for the backend REST API
# Builds the Go server in a lightweight builder and copies the static binary

FROM golang:1.24.5-alpine AS builder
WORKDIR /src

# Install git for module downloads if needed and set up build deps
RUN apk add --no-cache git

# Only copy go.mod/go.sum first to leverage layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source and build the server
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags='-s -w' -o /server ./cmd/server

FROM alpine:3.18
# Ensure TLS trust store is present so the binary can make HTTPS requests
RUN apk add --no-cache ca-certificates \
    && update-ca-certificates

RUN addgroup -S app && adduser -S -G app app

# Copy the statically linked binary from the builder
COPY --from=builder /server /server
RUN chmod +x /server

USER app
EXPOSE 8080
ENTRYPOINT ["/server"]
