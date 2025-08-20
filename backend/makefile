.DEFAULT_GOAL := build

.PHONY: fmt vet test build clean \
        build-fetcher build-server \
        run-fetcher run-server

# Format and lint
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

test:
	go test ./...

# Build both binaries
build: test vet build-fetcher build-server

# Build fetcher binary
build-fetcher:
	go build -o bin/fetcher ./cmd/fetcher

# Build server binary
build-server:
	go build -o bin/server ./cmd/server

# Run fetcher directly (without building binary)
run-fetcher:
	go run ./cmd/fetcher

# Run server directly (without building binary)
run-server:
	go run ./cmd/server

# Clean build artifacts
clean:
	go clean
	rm -rf bin/
