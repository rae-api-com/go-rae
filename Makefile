
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOGEN=$(GOCMD) generate
GOMOD=$(GOCMD) mod
GOFMT=gofmt

# Build info
VERSION?=$(shell git describe --tags --always --dirty)
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

.PHONY: all build test clean deps fmt vet help ci

all: deps fmt vet test build ## Run all checks and build

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the application
	$(GOBUILD) -v ./...

test: ## Run tests with coverage
	$(GOTEST) -v -race -coverprofile=coverage.out ./...

test-json: fmt generate ## Run tests with JSON output (original behavior)
	$(GOTEST) -v -json -failfast ./... | tparse --progress --all

test-coverage: test ## Generate coverage report
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

clean: ## Clean build artifacts
	$(GOCLEAN)
	rm -f coverage.out coverage.html

deps: ## Download and tidy dependencies
	$(GOMOD) download
	$(GOMOD) tidy 

fmt: ## Format code
	$(GOCMD) fmt ./...
	goimports -w .
	golines -w .

vet: ## Run go vet
	$(GOCMD) vet ./...

check: fmt vet ## Run format and vet checks

ci: deps check test ## Run CI checks locally

generate: ## Generate code
	$(GOGEN) $(MODULES)

setup: ## Setup development environment
	go install github.com/segmentio/golines@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/mailru/easyjson/...@latest
	go install github.com/mfridman/tparse@latest
	@echo "Development environment setup complete"

benchmark: ## Run benchmarks
	$(GOTEST) -bench=. -benchmem ./...

# Release helpers
tag-release: ## Tag a new release (usage: make tag-release VERSION=v1.0.0)
ifndef VERSION
	$(error VERSION is required. Usage: make tag-release VERSION=v1.0.0)
endif
	git tag -a $(VERSION) -m "Release $(VERSION)"
	git push origin $(VERSION)

# Documentation
docs: ## Show documentation info
	@echo "Documentation available at: https://pkg.go.dev/github.com/rae-api-com/go-rae"
	@echo "Local docs: godoc -http=:6060"

# Legacy targets (maintain compatibility)
tidy: deps ## Alias for deps 
