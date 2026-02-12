BINARY_NAME := hepic
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
BUILD_DATE := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
LDFLAGS := -ldflags "-X hepic-cli/cmd.Version=$(VERSION) -X hepic-cli/cmd.BuildDate=$(BUILD_DATE) -X hepic-cli/cmd.GitCommit=$(GIT_COMMIT)"

.PHONY: build test lint vet generate clean

build:
	go build $(LDFLAGS) -o $(BINARY_NAME) .

test:
	go test ./... -count=1

lint: vet
	@which golangci-lint > /dev/null 2>&1 || { echo "golangci-lint not installed, running go vet only"; exit 0; }
	golangci-lint run ./...

vet:
	go vet ./...

generate:
	go run tools/generate/main.go

clean:
	rm -f $(BINARY_NAME)
	go clean ./...
