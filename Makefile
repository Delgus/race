PKG := "github.com/delgus/race"

.PHONY: all fmt lint test build clean bench help

all: build

fmt: ## gofmt all project
	@gofmt -l -s -w .

lint: ## Lint the files
	@golangci-lint run

test: ## Run unittests
	@go test -race -short ./... -coverprofile=coverage.txt

build: ## Build the binary file
	@go build -a -o mybank -v $(PKG)/cmd/bank

clean: ## Remove previous build
	@rm -f mybank

bench: ## Benchmark for code
	@go test ./... -bench=. -benchmem

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
