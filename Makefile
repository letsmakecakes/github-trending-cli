.PHONY: build test test-unit test-integration test-race clean install run lint coverage help

# Binary name
BINARY_NAME=trending-repos
BUILD_DIR=bin
MAIN_PATH=cmd/trending-repos/main.go

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOINSTALL=$(GOCMD) install

# Build flags
LDFLAGS=-ldflags "-s -w"

# Detect OS for race detection
ifeq ($(OS),Windows_NT)
	RACE_FLAG=
else
	RACE_FLAG=-race
endif

all: test build

## build: Build the application
build:
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

## build-all: Build for all platforms
build-all:
	@echo "Building for multiple platforms..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_PATH)
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)
	@echo "Build complete for all platforms"

## test: Run all tests
test:
	@echo "Running tests..."
	$(GOTEST) -v $(RACE_FLAG) -coverprofile=coverage.txt -covermode=atomic ./...

## test-race: Run all tests with race detection (requires CGO)
test-race:
	@echo "Running tests with race detection..."
	CGO_ENABLED=1 $(GOTEST) -v -race -coverprofile=coverage.txt -covermode=atomic ./...

## test-unit: Run only unit tests
test-unit:
	@echo "Running unit tests..."
	$(GOTEST) -v -short $(RACE_FLAG) ./...

## test-integration: Run only integration tests
test-integration:
	@echo "Running integration tests..."
	$(GOTEST) -v -tags=integration ./test/integration/...

## coverage: Generate coverage report
coverage: test
	@echo "Generating coverage report..."
	$(GOCMD) tool cover -html=coverage.txt -o coverage.html
	@echo "Coverage report generated: coverage.html"

## lint: Run linter
lint:
	@echo "Running linter..."
	@which golangci-lint > /dev/null || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	golangci-lint run ./...

## clean: Clean build files
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -f coverage.txt coverage.html
	@echo "Clean complete"

## install: Install the application
install: build
	@echo "Installing..."
	$(GOINSTALL) $(MAIN_PATH)
	@echo "Installation complete"

## run: Run the application with default parameters
run: build
	@echo "Running..."
	./$(BUILD_DIR)/$(BINARY_NAME) --duration week --limit 10

## run-day: Fetch daily trending repos
run-day: build
	./$(BUILD_DIR)/$(BINARY_NAME) --duration day --limit 10

## run-month: Fetch monthly trending repos
run-month: build
	./$(BUILD_DIR)/$(BINARY_NAME) --duration month --limit 20

## deps: Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

## fmt: Format code
fmt:
	@echo "Formatting code..."
	$(GOCMD) fmt ./...

## vet: Run go vet
vet:
	@echo "Running go vet..."
	$(GOCMD) vet ./...

## help: Display this help message
help:
	@echo "Available targets:"
	@grep -E '^## ' Makefile | sed 's/## /  /'