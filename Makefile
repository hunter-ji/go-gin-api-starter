# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=go-gin-api-starter
BINARY_UNIX=$(BINARY_NAME)_linux
MAIN_PATH=cmd/api/main.go

# Build flags
BUILD_FLAGS=-v

.PHONY: all build run clean test coverage deps lint help build-linux build-linux-mac-intel build-linux-mac-arm build-linux-win

all: test build

build:
	$(GOBUILD) $(BUILD_FLAGS) -o $(BINARY_NAME) $(MAIN_PATH)

run:
	$(GORUN) $(MAIN_PATH)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

test:
	$(GOTEST) -v ./...

coverage:
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out

deps:
	$(GOGET) -v -t -d ./...
	$(GOMOD) tidy

lint:
	golangci-lint run

# Build for Linux on Linux
build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(BUILD_FLAGS) -o $(BINARY_UNIX) $(MAIN_PATH)

# Build for Linux on macOS (Intel)
build-linux-mac-intel:
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(BUILD_FLAGS) -o $(BINARY_UNIX) $(MAIN_PATH)

# Build for Linux on macOS (Apple Silicon)
build-linux-mac-arm:
	GOOS=linux GOARCH=arm64 $(GOBUILD) $(BUILD_FLAGS) -o $(BINARY_UNIX)_arm64 $(MAIN_PATH)

# Build for Linux on Windows
build-linux-win:
	SET GOOS=linux
	SET GOARCH=amd64
	$(GOBUILD) $(BUILD_FLAGS) -o $(BINARY_UNIX) $(MAIN_PATH)

help:
	@echo "Available commands:"
	@echo "  make build                  - Build the binary for current OS"
	@echo "  make run                    - Run the application"
	@echo "  make clean                  - Remove binary and cache"
	@echo "  make test                   - Run tests"
	@echo "  make coverage               - Run tests with coverage"
	@echo "  make deps                   - Download dependencies"
	@echo "  make lint                   - Run linter"
	@echo "  make build-linux            - Build for Linux on Linux"
	@echo "  make build-linux-mac-intel  - Build for Linux on macOS (Intel)"
	@echo "  make build-linux-mac-arm    - Build for Linux on macOS (Apple Silicon)"
	@echo "  make build-linux-win        - Build for Linux on Windows"
