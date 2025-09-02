.DEFAULT_GOAL := test

.PHONY: fmt vet test build run-server run-fetcher clean

# Format and vet
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

# Run unit tests
test:
	go test ./...

# Build binaries
build:
	go build -o bin/server ./cmd/server
	go build -o bin/fetcher ./cmd/fetcher

# Run server or fetcher directly
run-server:
	go run ./cmd/server

run-fetcher:
	go run ./cmd/fetcher

# Clean artifacts
clean:
	go clean
	rm -rf bin/
	rm -f coverage.out coverage.html
