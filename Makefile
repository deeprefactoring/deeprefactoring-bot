PACKAGE := github.com/deeprefactoring/deeprefactoring-bot
BINARY_NAME := deeprefactoring-bot

SHELL := /bin/bash

VERSION ?= $(shell git describe --long --tags --dirty --always)
DATE := $(shell date +%FT%T%z)

VERSION_FLAGS := -X main.buildVersion=$(VERSION) -X main.buildDate=$(DATE)

.PHONY: version
version:
	@echo $(VERSION)

.PHONY: format-check
format-check:
	@gofmt -l . | tee /dev/stderr | grep ^ && exit 1 || true > /dev/null 2>&1

.PHONY: vet
vet:
	@go vet ./...

.PHONY: lint
lint: format-check vet

.PHONY: build
build:
	go build -v -x \
	    -ldflags "-extldflags -static $(VERSION_FLAGS)" \
	    -o $(BINARY_NAME) \
	    $(PACKAGE)/cmd/app

.PHONY: test
test:
	go test -v \
	    -cover -coverprofile=coverage.out \
	    -race \
	    ./...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: bench
bench:
	go test -v \
	    -bench . -benchtime 5s \
	    $(PACKAGE)

.PHONY: package
package: lint test build
