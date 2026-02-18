# Makefile for NullDB backend

build:
	@echo "Building..."
	@go build -o bin/main main.go

fmt:
	@go fmt ./...
	@go mod tidy

run:
	@go run main.go

clean:
	@echo "Cleaning..."
	@rm -rf bin

test:
	@echo "Testing..."
	@go test ./... -v

.PHONY: build run clean test
