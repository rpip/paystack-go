#!/bin/bash
export PAYSTACK_KEY=sk_test_b748a89ad84f35c2f1a8b81681f956274de048bb

go vet $(go list ./... | grep -v vendor)
test -z "$(golint ./... | grep -v vendor | tee /dev/stderr)"
test -z "$(gofmt -s -l . | grep -v vendor | tee /dev/stderr)"
go test -cover $(go list ./... | grep -v vendor)
go test -race $(go list ./... | grep -v vendor)
go test $(go list ./... | grep -v vendor)
