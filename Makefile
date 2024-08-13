.PHONY: gen-proto
gen-proto: ## generate proto
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
test: ## run the tests
	go test -coverprofile cover.out ./...

.PHONY: test-bench
test-bench: ## run the tests bench
	go test ./... -test.run=NONE -test.bench=. -test.benchmem | tee ./assets/out.dat ; \
	awk '/Benchmark/{count ++; gsub(/BenchmarkTest/,""); printf("%d,%s,%s,%s\n",count,$$1,$$2,$$3)}' ./assets/out.dat > ./assets/final.dat ; \
		gnuplot -e "file_path='./assets/final.dat'" -e "graphic_file_name='./assets/operations.png'" -e "y_label='number of operations'" -e "y_range_min='000000000''" -e "y_range_max='400000000'" -e "column_1=1" -e "column_2=3" ./assets/performance.gp ; \
		gnuplot -e "file_path='./assets/final.dat'" -e "graphic_file_name='./assets/time_operations.png'" -e "y_label='each operation in nanoseconds'" -e "y_range_min='000''" -e "y_range_max='45000'" -e "column_1=1" -e "column_2=4" ./assets/performance.gp ; \
		rm -f ./assets/out.dat ./assets/final.dat ; \
		echo "'assets/operations.png' and 'assets/time_operations.png' graphics were generated."


.PHONY: help
help: ## print help
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m\033[0m\n"} /^[$$()% 0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
