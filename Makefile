default: build

.PHONY: all build test vet
all: build test vet

build:
	go mod download
	go get ./...
	go build ./...

test:
	go test ./...

vet:
	go vet ./...
