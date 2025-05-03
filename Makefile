
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOGEN=$(GOCMD) generate
GOMOD=$(GOCMD) mod
BINARY_NAME=raeapi
BINARY_UNIX=$(BINARY_NAME)_unix

# All
all: test 

# Test
test: fmt generate
	$(GOTEST) -v -json -failfast ./...  | tparse --progress --all

# Format
fmt:
	$(GOCMD) fmt ./...
	goimports -w .
	golines -w .


# Tidy dependencies
tidy:
	$(GOMOD) tidy


generate:
	$(GOGEN) $(MODULES)

setup:
	go install github.com/segmentio/golines@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/mailru/easyjson/...@latest
	go install github.com/mfridman/tparse@latest

.PHONY: all fmt generate setup 
