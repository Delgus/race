PKG := "github.com/delgus/race"

.PHONY: all dep build clean test coverage coverhtml lint

all: build

fmt: ## gofmt all project
	@gofmt -l -s -w .

lint: ## Lint the files
	@golangci-lint run

test: ## Run unittests
	@go test -short ./... -coverprofile cover.out
	@go tool cover -func cover.out

build: ## Build the binary file
	@go build -race -a -o mybank -v $(PKG)/cmd/bank

clean: ## Remove previous build
	@rm -f mybank

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'