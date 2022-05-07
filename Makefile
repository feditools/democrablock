PROJECT_NAME=democrablock

.DEFAULT_GOAL := test

check:
	golangci-lint run

check-fix:
	golangci-lint run --fix

fmt:
	@echo formatting
	@go fmt $(shell go list ./... | grep -v /vendor/)

test:  tidy check-fix fmt
	go test -cover ./...

tidy:
	go mod tidy -compat=1.17

vendor: tidy
	go mod vendor

.PHONY: check check-fix fmt test tidy vendor