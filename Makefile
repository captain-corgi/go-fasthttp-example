.PHONY: build run test clean lint help

APP_NAME=go-fasthttp-example
MAIN_PATH=./cmd/server

help:
	@echo "Available commands:"
	@echo "  tidy    - Run go mod tidy and vendor"
	@echo "  build   - Build the application"
	@echo "  run     - Run the application"
	@echo "  test    - Run tests"
	@echo "  clean   - Remove build artifacts"
	@echo "  lint    - Run linter"

tidy:
	go mod tidy
	go mod vendor

build:
	go build -o bin/$(APP_NAME) $(MAIN_PATH)

run:
	go run $(MAIN_PATH)/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/

lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint is not installed. Please install it using: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi