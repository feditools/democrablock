PROJECT_NAME=democrablock

.DEFAULT_GOAL := test

bun-new-migration: export BUN_TIMESTAMP=$(shell date +%Y%m%d%H%M%S | head -c 14)
bun-new-migration:
	touch internal/db/bun/migrations/${BUN_TIMESTAMP}_new.go
	cat internal/db/bun/migrations/migration.go.tmpl > internal/db/bun/migrations/${BUN_TIMESTAMP}_new.go

check:
	golangci-lint run

check-fix:
	golangci-lint run --fix

fmt:
	@echo formatting
	@go fmt $(shell go list ./... | grep -v /vendor/)

test:  tidy fmt
	go test -race -cover ./...

tidy:
	go mod tidy -compat=1.17

vendor: tidy
	go mod vendor

.PHONY: bun-new-migration check check-fix fmt test tidy vendor