# Project name
PROJECT = paystack

# Set an output prefix, which is the local directory if not specified
PREFIX?=$(shell pwd)
BUILDTAGS=
GLIDE = $(shell which glide)

.PHONY: clean all fmt vet lint build test static deps docker
.DEFAULT: default

all: clean build fmt lint test vet

build:
	@echo "+ $@"
	@go build -tags "$(BUILDTAGS) cgo" .

static:
	@echo "+ $@"
	CGO_ENABLED=1 go build -tags "$(BUILDTAGS) cgo static_build" -ldflags "-w -extldflags -static" -o reg .

fmt:
	@echo "+ $@"
	@gofmt -s -l . | grep -v vendor | tee /dev/stderr

lint:
	@echo "+ $@"
	@golint ./... | grep -v vendor | tee /dev/stderr

test: fmt lint vet
	@echo "+ $@"
	@PAYSTACK_KEY=$(PAYSTACK_KEY) go test -v -tags "$(BUILDTAGS) cgo" $(shell go list ./... | grep -v vendor)

vet:
	@echo "+ $@"
	@go vet $(shell go list ./... | grep -v vendor)

clean:
	@echo "+ $@"
	@rm -rf reg

deps:
	@echo "Installing dependencies..."
	@$(GLIDE) install

docker:
	@docker build . -t $(PROJECT) -f Dockerfile.test
