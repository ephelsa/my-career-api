lint-prepare:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

test:
	go test -v -cover -covermode=atomic ./...

.PHONY: lint-prepare lint
