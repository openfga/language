BINARY_NAME=openfga-language
DOCKER_BINARY=docker
GO_BIN ?= $(shell go env GOPATH)/bin
.DEFAULT_GOAL := help

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


#-----------------------------------------------------------------------------------------------------------------------
# Rules (https://www.gnu.org/software/make/manual/html_node/Rule-Introduction.html#Rule-Introduction)
#-----------------------------------------------------------------------------------------------------------------------
$(GO_BIN)/golangci-lint:
	@echo "==> Installing golangci-lint within "${GO_BIN}""
	@go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@latest

$(GO_BIN)/govulncheck:
	@echo "==> Installing govulncheck within "${GO_BIN}""
	@go install -v golang.org/x/vuln/cmd/govulncheck@latest

$(GO_BIN)/gofumpt:
	@echo "==> Installing gofumpt within "${GO_BIN}""
	@go install -v mvdan.cc/gofumpt@latest

#-----------------------------------------------------------------------------------------------------------------------
# Phony Rules(https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html)
#-----------------------------------------------------------------------------------------------------------------------
all: build

build:
	go build ./...

clean: ## Clean project files
	go clean

test: ## Run Go tests
	go test ./... -count=1 -race

lint: $(GO_BIN)/golangci-lint ## Lint Go source files
	@echo "==> Linting Go source files"
	@golangci-lint run -v --fix -c .golangci.yaml ./...

audit: $(GO_BIN)/govulncheck ## Audit Go source files
	@echo "==> Checking Go source files for vulnerabilities"
	govulncheck ./...

format: $(GO_BIN)/gofumpt ## Format Go source files
	@echo "==> Formatting project files"
	gofumpt -w transformer/ errors/

all-tests: build audit lint test ## Run linting and tests