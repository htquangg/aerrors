.PHONY: generate-proto
generate-proto: ## generate proto
	buf generate

.PHONY: lint
lint: ## lint
	@echo "Linting"
	golangci-lint run \
		--timeout 10m \
		--config ./.golangci.yaml \
		--out-format=github-actions \
		--concurrency=$$(getconf _NPROCESSORS_ONLN)

.PHONY: test
test: ## run the go tests
	go test -coverprofile cover.out ./...

.PHONY: help
help: ## print help
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m\033[0m\n"} /^[$$()% 0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
