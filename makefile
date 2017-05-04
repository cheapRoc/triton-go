SHELL := /bin/bash

test-deps:
	@echo Installing dev/test dependencies
	go get -u github.com/golang/lint/golint
	go get -u github.com/golang/dep/...

deps: clean test-deps
	dep ensure

clean:
	rm -rf vendor

lint:
	@golint

test:
	@go test -v
