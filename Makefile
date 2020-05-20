default: build

.PHONY: all build test vet
all: build test vet

build:
	go mod download
	go get ./...

test:
	go test ./...

vet:
	go vet ./...
