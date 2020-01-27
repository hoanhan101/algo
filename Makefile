#
# algo cmd helpers
#

.PHONY: count
count:  ## Count the number of questions
	go run count.go

.PHONY: clean
clean:  ## Clean test cache
	go clean -testcache
	rm -f coverage.out

.PHONY: cover
cover: test  ## Run unit tests and open the coverage report
	go tool cover -html=coverage.out

.PHONY: lint
lint:  ## Lint source files
	golangci-lint run

.PHONY: test
test:  ## Run unit tests
	go test -short -race -coverprofile=coverage.out -covermode=atomic ./...

.PHONY: test-verbose
test-verbose:  ## Run tests verbosely
	go test -v ./...

.PHONY: ready
ready: clean lint test count  ## Clean up, lint source files, run tests and be ready for a push

.PHONY: help
help:  ## Print usage information
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

.DEFAULT_GOAL := help
