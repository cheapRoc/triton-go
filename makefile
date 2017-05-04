SHELL := /bin/bash
.SHELLFLAGS := -o pipefail -euc
.DEFAULT_GOAL := test

test-deps:
	@echo Installing dev/test dependencies
	go get -u github.com/golang/lint/golint
	go get -u github.com/golang/dep/...

deps: clean test-deps
	dep ensure

clean:
	rm -rf vendor

lint:
	@bash ./scripts/lint.sh

test:
	@go test -v
