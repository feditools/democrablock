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

clean:
	@echo cleaning up workspace
	@rm -Rvf coverage.txt dist democrablock
	@find . -name ".DS_Store" -exec rm -v {} \;
	@rm -Rvf web/static/css/default.min.css web/static/css/error.min.css

fmt:
	@echo formatting
	@go fmt $(shell go list ./... | grep -v /vendor/)

stage-static:
	minify web/static-src/css/default.css > web/static/css/default.min.css
	minify web/static-src/css/error.css > web/static/css/error.min.css

test:  tidy fmt
	go test -race -cover ./...

test-docker-pull:
	docker-compose --project-name ${PROJECT_NAME} -f deployments/docker-compose-test.yaml pull

test-docker-restart: test-docker-stop test-docker-start

test-docker-start:
	docker-compose --project-name ${PROJECT_NAME} -f deployments/docker-compose-test.yaml up -d

test-docker-stop:
	docker-compose --project-name ${PROJECT_NAME} -f deployments/docker-compose-test.yaml down

test-ext: tidy fmt
	go test --tags=postgres,redis -cover ./...

tidy:
	go mod tidy -compat=1.17

vendor: tidy
	go mod vendor

.PHONY: bun-new-migration check check-fix fmt stage-static test test-docker-restart test-docker-start test-docker-stop tidy vendor