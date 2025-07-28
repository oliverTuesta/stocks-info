.DEFAULT_GOAL := build

.PHONY:fmt vet build
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

test:
	go test

build: test vet
	go build

clean:
	go clean

