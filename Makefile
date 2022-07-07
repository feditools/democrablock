PROJECT_NAME=democrablock

.DEFAULT_GOAL := test

build-snapshot: clean
	goreleaser build --snapshot

bun-new-migration: export BUN_TIMESTAMP=$(shell date +%Y%m%d%H%M%S | head -c 14)
bun-new-migration:
	touch internal/db/bun/migrations/${BUN_TIMESTAMP}_new.go
	cat internal/db/bun/migrations/migration.go.tmpl > internal/db/bun/migrations/${BUN_TIMESTAMP}_new.go

check:
	golangci-lint run

check-fix:
	golangci-lint run --fix

clean:
	@echo cleaning up workspace
	@rm -Rvf coverage.txt dist democrablock
	@find . -name ".DS_Store" -exec rm -v {} \;
	@rm -Rvf web/static/css/default.min.css web/static/css/error.min.css

docker-pull:
	docker-compose --project-name ${PROJECT_NAME} -f deployments/docker-compose-test.yaml pull

docker-restart: docker-stop docker-start

docker-start:
	docker-compose --project-name ${PROJECT_NAME} -f deployments/docker-compose-test.yaml up -d

docker-stop:
	docker-compose --project-name ${PROJECT_NAME} -f deployments/docker-compose-test.yaml down

fmt:
	@echo formatting
	@go fmt $(shell go list ./... | grep -v /vendor/)

stage-static:
	minify web/static-src/css/default.css > web/static/css/default.min.css
	minify web/static-src/css/error.css > web/static/css/error.min.css
	minify web/static-src/css/login.css > web/static/css/login.min.css

test:  tidy fmt
	go test -race -cover ./...

test-ext: tidy fmt
	go test --tags=postgres,redis -cover ./...

tidy:
	go mod tidy -compat=1.17

vendor: tidy
	go mod vendor

.PHONY: build-snapshot bun-new-migration check check-fix docker-pull docker-restart docker-start docker-stop fmt stage-static test test-ext tidy vendor
