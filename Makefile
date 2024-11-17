# Variable definitions
GO := go
APP_NAME := vault-sync
BUILD_DIR := ./bin
LINT := golangci-lint

# Default goal
.DEFAULT_GOAL := build

# .PHONY declarations
.PHONY: wire build run test clean fmt lint

# wire target: Generate code using Wire
wire:
	@echo "Running wire to generate code..."
	@wire ./...

# build target: Build the Go project
build: clean
	@echo "Building the Go project..."
	$(GO) build -o $(BUILD_DIR)/$(APP_NAME) .

# run target: Run the Go project
run: build
	@echo "Running the Go project..."
	$(BUILD_DIR)/$(APP_NAME)

# test target: Run Go tests
test:
	@echo "Running tests..."
	$(GO) test -v ./... -p 4 # -p 4: Run tests in parallel to speed up execution

# clean target: Clean up generated files
clean:
	@echo "Cleaning up generated files..."
	rm -rf $(BUILD_DIR)

# fmt target: Format the Go code
fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...

# lint target: Run linters to check code quality
lint:
	@command -v $(LINT) >/dev/null 2>&1 || { echo "$(LINT) is not installed, please install it."; exit 1; }
	@echo "Running linters..."
	$(LINT) run
