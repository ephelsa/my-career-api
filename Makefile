BUILD_NAME ?= "./bin/mycareer-api"

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
