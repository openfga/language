BINARY_NAME=openfga-language
DOCKER_BINARY=docker
GO_BIN ?= $(shell go env GOPATH)/bin

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

test: ## Run tests
	go test ./... -count=1

lint: $(GO_BIN)/golangci-lint ## Lint Go source files
	@echo "==> Linting Go source files"
	@golangci-lint run -v --fix -c .golangci.yaml ./...

audit: $(GO_BIN)/govulncheck ## Audit Go source files
	@echo "==> Checking Go source files for vulnerabilities"
	govulncheck ./...

format: $(GO_BIN)/gofumpt ## Format Go source files
	@echo "==> Formatting project files"
	gofumpt -w transformer/ errors/

all-tests: build audit lint test