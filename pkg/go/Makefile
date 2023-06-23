BINARY_NAME=openfga-language
DOCKER_BINARY=docker

all: build

build: lint audit format
	go build ./...

clean:
	go clean

test:
	go test ./...

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && golangci-lint run

audit:
	go install golang.org/x/vuln/cmd/govulncheck@latest && govulncheck ./...;

format:
	go install mvdan.cc/gofumpt@latest && gofumpt -w .
	go install github.com/daixiang0/gci@latest && gci write -s standard -s default .