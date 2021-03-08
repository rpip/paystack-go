#!/bin/bash

go vet $(go list ./... | grep -v vendor)
test -z "$(golint ./... | grep -v vendor | tee /dev/stderr)"
test -z "$(gofmt -s -l . | grep -v vendor | tee /dev/stderr)"
go test -cover $(go list ./... | grep -v vendor)
go test -race $(go list ./... | grep -v vendor)
go test $(go list ./... | grep -v vendor)
