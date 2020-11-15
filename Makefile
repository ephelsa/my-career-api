lint-prepare:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

.PHONY: lint-prepare lint

test:
	go test -v -cover -covermode=atomic ./...

.PHONY: test

database-run:
	docker-compose --env-file ./internal/env/dev.env up

.PHONY: database-run
