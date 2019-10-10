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
	@go build -a -o bin/bank1 -v $(PKG)/cmd/bank1
	@go build -a -o bin/bank2 -v $(PKG)/cmd/bank2
	@go build -a -o bin/bank3 -v $(PKG)/cmd/bank3
	@go build -a -o bin/bank4 -v $(PKG)/cmd/bank4
	@go build -a -o bin/bank5 -v $(PKG)/cmd/bank5
	@go build -a -o bin/bank6 -v $(PKG)/cmd/bank6

clean: ## Remove previous build
	@rm -f bin/bank1 bin/bank2 bin/bank3 bin/bank4 bin/bank5 bin/bank6

bench: ## Benchmark for code
	@go test ./... -bench=. -benchmem

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
