PACKAGE := github.com/deeprefactoring/deeprefactoring-bot

GOFILES_NOVENDOR = $(shell find . -type f -name '*.go' | grep -v /vendor/ | grep -v /.glide/)

SHELL := /bin/bash

VERSION ?= $(shell git describe --long --tags --dirty --always)
DATE := $(shell date +%FT%T%z)

VERSION_FLAGS := -X main.buildVersion=$(VERSION) -X main.buildDate=$(DATE)

.PHONY: version
version:
	@echo $(VERSION)

.PHONY: deps
deps:
	glide install

.PHONY: format-check
format-check:
	@gofmt -l -d $(GOFILES_NOVENDOR) | tee /dev/stderr | grep ^ && exit 1 || true > /dev/null 2>&1

.PHONY: vet
vet:
	@go tool vet -all $(GOFILES_NOVENDOR)

.PHONY: lint
lint: format-check vet

.PHONY: build-static
build-static:
	go build -v -x -race \
	    -ldflags "-linkmode external -extldflags -static $(VERSION_FLAGS)" \
	    -o deeprefactoring-bot \
	    $(PACKAGE)/cmd/app

.PHONY: build
build:
	go build -v -ldflags "$(VERSION_FLAGS)" $(PACKAGE)/cmd/app

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
package: lint deps test build-static


#
# Following targets related to CI and docker
#

# run inside docker

# entry point for CI
.PHONY: ci
ci:
	docker run -it \
	    -u $(id -u):$(id -g) \
	    -v `pwd`/Makefile:/go/Makefile \
	    -v `pwd`:/go/src/$(PACKAGE) \
	    -w /go/src/$(PACKAGE) \
	    ykhrustalev/go-glide:1.8.0 \
	    make package
