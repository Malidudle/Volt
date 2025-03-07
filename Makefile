# Volt Makefile
# This Makefile provides commands for common development tasks

# Variables
PORT ?= 8080
GO = go
APP_NAME = volt

# Colors for output
BLUE = \033[1;34m
GREEN = \033[1;32m
YELLOW = \033[1;33m
RESET = \033[0m

.PHONY: help run generate dev build test clean vet fmt install-air

# Default target - shows help
help:
	@echo "Volt Commands:"
	@echo "run        - Run API server"
	@echo "generate   - Generate route definitions"
	@echo "dev        - Run with hot reload (requires Air)"
	@echo "build      - Build application binary"
	@echo "test       - Run tests"
	@echo "fmt        - Format code"
	@echo "vet        - Run Go vet"
	@echo "clean      - Clean build artifacts"
	@echo "install-air - Install Air for hot reload"
	@echo "Use PORT=xxxx to specify custom port (default: 8080)"

# Run the server
run:
	@echo "${BLUE}Starting server on port $(PORT)...${RESET}"
	@PORT=$(PORT) $(GO) run main.go

# Generate route definitions
generate:
	@echo "${BLUE}Generating routes...${RESET}"
	@$(GO) run cmd/generate/main.go

# Run development server with hot reload
dev: generate
	@echo "${BLUE}Looking for air...${RESET}"
	@if command -v air > /dev/null; then \
		echo "${GREEN}Air found in PATH${RESET}"; \
		PORT=$(PORT) air; \
	elif [ -f ~/go/bin/air ]; then \
		echo "${GREEN}Air found in ~/go/bin${RESET}"; \
		PORT=$(PORT) ~/go/bin/air; \
	elif [ -f $${GOPATH}/bin/air ]; then \
		echo "${GREEN}Air found in GOPATH${RESET}"; \
		PORT=$(PORT) $${GOPATH}/bin/air; \
	else \
		echo "${YELLOW}Air not found. Install with 'make install-air'${RESET}"; \
		PORT=$(PORT) $(GO) run main.go; \
	fi

# Build the application
build: generate
	@echo "${BLUE}Building $(APP_NAME)...${RESET}"
	@$(GO) build -o $(APP_NAME) main.go

# Run tests
test:
	@$(GO) test ./... -v

# Format the code
fmt:
	@$(GO) fmt ./...

# Run Go vet
vet:
	@$(GO) vet ./...

# Clean build artifacts
clean:
	@rm -f $(APP_NAME)
	@rm -rf tmp

# Install air for hot reloading
install-air:
	@echo "${BLUE}Installing Air...${RESET}"
	@$(GO) install github.com/air-verse/air@latest
	@echo "${GREEN}Air installed. Set up your PATH:${RESET}"
	@echo "Bash/zsh: export PATH=\$$PATH:\$$HOME/go/bin"
	@echo "Windows: set PATH=%PATH%;%USERPROFILE%\\go\\bin" 