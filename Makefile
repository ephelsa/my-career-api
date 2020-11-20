BUILD_NAME ?= "./bin/mycareer-api"
PKG_NAME ?= ""

lint-prepare:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	@echo "Running golangci lint"
	./bin/golangci-lint run ./...

.PHONY: lint-prepare lint

test:
	go test -v -cover -covermode=atomic ./...

.PHONY: test

database-run:
	@echo "Dev database"
	docker-compose --env-file ./internal/env/dev.env up

.PHONY: database-run

remove-build:
	@rm -f $(BUILD_NAME)

.PHONY: remove-build

build: remove-build
	@go build -trimpath -o $(BUILD_NAME) ./cmd/mycareer/

.PHONY: remove-build

new-pkg:
	@mkdir pkg/$(PKG_NAME)

	@mkdir pkg/$(PKG_NAME)/data
	@touch pkg/$(PKG_NAME)/data/$(PKG_NAME)_repository.go

	@mkdir pkg/$(PKG_NAME)/domain
	@touch pkg/$(PKG_NAME)/domain/$(PKG_NAME).go

	@mkdir pkg/$(PKG_NAME)/infraestructure
	@mkdir pkg/$(PKG_NAME)/infraestructure/database
	@touch pkg/$(PKG_NAME)/infraestructure/database/$(PKG_NAME)_database.go
	@touch pkg/$(PKG_NAME)/infraestructure/database/$(PKG_NAME)_database_test.go
	@mkdir pkg/$(PKG_NAME)/infraestructure/server
	@touch pkg/$(PKG_NAME)/infraestructure/server/$(PKG_NAME)_server.go
	@touch pkg/$(PKG_NAME)/infraestructure/server/$(PKG_NAME)_server_test.go
