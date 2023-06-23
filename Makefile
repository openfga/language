DOCKER_BINARY=docker
ANTLR_DOCKER_IMAGE=docker.io/rhamzeh/antlr
ANTLR_CMD=${DOCKER_BINARY} run -t --rm -v ${PWD}:/app ${ANTLR_DOCKER_IMAGE}

.PHONY: all
all: build-go

.PHONY: antlr-gen-go
antlr-gen-go:
	${ANTLR_CMD} -Dlanguage=Go -o src/go/gen /app/OpenFGA.g4

.PHONY: build-go
build-go: antlr-gen-go
	$(MAKE) -C src/go build

.PHONY: run-go
run-go: antlr-gen-go
	$(MAKE) -C src/go run

.PHONY: clean-go
clean-go:
	$(MAKE) -C src/go clean

.PHONY: test-go
test-go: antlr-gen-go
	$(MAKE) -C src/go test

.PHONY: lint-go
lint-go: antlr-gen-go
	$(MAKE) -C src/go lint

.PHONY: audit-go
audit-go: antlr-gen-go
	$(MAKE) -C src/go audit

.PHONY: format-go
format-go: antlr-gen-go
	$(MAKE) -C src/go format

.PHONY: antlr-gen-js
antlr-gen-js:
	${ANTLR_CMD} -Dlanguage=JavaScript -o src/js/gen /app/OpenFGA.g4
