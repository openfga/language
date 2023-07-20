DOCKER_BINARY=docker
ANTLR_DOCKER_IMAGE=docker.io/openfga_utils/antlr
ANTLR_CMD=${DOCKER_BINARY} run -t --rm -v ${PWD}:/app ${ANTLR_DOCKER_IMAGE}

#### Go #####

.PHONY: antlr-gen-go
antlr-gen-go: build-antlr-container
	${ANTLR_CMD} -Dlanguage=Go -o pkg/go/gen /app/OpenFGA.g4

.PHONY: build-go
build-go: antlr-gen-go
	$(MAKE) -C pkg/go build

.PHONY: run-go
run-go: antlr-gen-go
	$(MAKE) -C pkg/go run

.PHONY: clean-go
clean-go:
	$(MAKE) -C pkg/go clean

.PHONY: test-go
test-go: antlr-gen-go
	$(MAKE) -C pkg/go test

.PHONY: lint-go
lint-go: antlr-gen-go
	$(MAKE) -C pkg/go lint

.PHONY: audit-go
audit-go: antlr-gen-go
	$(MAKE) -C pkg/go audit

.PHONY: format-go
format-go: antlr-gen-go
	$(MAKE) -C pkg/go format

.PHONY: all-tests-go
all-tests-go: antlr-gen-go
	$(MAKE) -C pkg/go all-tests

#### Util ####

.PHONY: build-antlr-container
build-antlr-container:
	docker build -f antlr.Containerfile -t ${ANTLR_DOCKER_IMAGE} .