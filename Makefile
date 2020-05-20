default: build

all: build test vet

build:
	go get ./...

test:
	go test ./...

vet:
	go vet ./...
