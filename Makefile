.DEFAULT_GOAL := build

.PHONY: build fmt vet run

fmt:
	echo "Running go fmt"
	go fmt ./...

vet:
	echo "Running go vet"
	go vet ./...

build: fmt vet
	echo "Building the binary"
	go build -o bin/ ./...

# for development
run: fmt vet
	echo "Running the server"
	go run cmd/api-server/main.go