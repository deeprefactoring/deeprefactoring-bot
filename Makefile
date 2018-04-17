PACKAGE := github.com/deeprefactoring/deeprefactoring-bot
BINARY_NAME := deeprefactoring-bot

GOFILES_NOVENDOR = $(shell find . -type f -name '*.go' | grep -v /vendor/ )

SHELL := /bin/bash

VERSION ?= $(shell git describe --long --tags --dirty --always)
DATE := $(shell date +%FT%T%z)

VERSION_FLAGS := -X main.buildVersion=$(VERSION) -X main.buildDate=$(DATE)

.PHONY: version
version:
	@echo $(VERSION)

.PHONY: deps
deps:
	dep ensure

.PHONY: format-check
format-check:
	@gofmt -l -d $(GOFILES_NOVENDOR) | tee /dev/stderr | grep ^ && exit 1 || true > /dev/null 2>&1

.PHONY: vet
vet:
	@go tool vet -all $(GOFILES_NOVENDOR)

.PHONY: lint
lint: format-check vet

.PHONY: build-static
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
	    $(PACKAGE)
	go tool cover -html=coverage.out -o coverage.html

.PHONY: bench
bench:
	go test -v \
	    -bench . -benchtime 5s \
	    $(PACKAGE)

.PHONY: package
package: lint deps test build
