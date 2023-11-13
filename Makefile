docker_binary=docker

ANTLR_DOCKER_IMAGE=docker.io/openfga_utils/antlr
ANTLR_CMD=${docker_binary} run -t --rm -v ${PWD}:/app:Z ${ANTLR_DOCKER_IMAGE}

#### Global #####

.PHONY: all
all: build

.PHONY: antlr-gen
antlr-gen: antlr-gen-go antlr-gen-js

.PHONY: build
build: build-go build-js

.PHONY: test
test: test-go test-js

.PHONY: lint
lint: lint-go lint-js

#### Go #####

.PHONY: antlr-gen-go
antlr-gen-go:
	$(MAKE) antlr-gen-base language=Go packageName=go

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

#### TypeScript #####

.PHONY: antlr-gen-js
antlr-gen-js:
	$(MAKE) antlr-gen-base language=TypeScript packageName=js

.PHONY: build-js
build-js: antlr-gen-js
	$(MAKE) -C pkg/js build

.PHONY: run-js
run-js: antlr-gen-js
	$(MAKE) -C pkg/js run

.PHONY: clean-js
clean-js:
	$(MAKE) -C pkg/js clean

.PHONY: test-js
test-js: antlr-gen-js
	$(MAKE) -C pkg/js test

.PHONY: lint-js
lint-js: antlr-gen-js
	$(MAKE) -C pkg/js lint

.PHONY: audit-js
audit-js: antlr-gen-js
	$(MAKE) -C pkg/js audit

.PHONY: format-js
format-js: antlr-gen-js
	$(MAKE) -C pkg/js format

.PHONY: all-tests-js
all-tests-js: antlr-gen-js
	$(MAKE) -C pkg/js all-tests

#### Util ####

.PHONY: build-antlr-container
build-antlr-container:
	${docker_binary} build -f antlr.Containerfile -t ${ANTLR_DOCKER_IMAGE} .

.PHONY: antlr-gen-base
antlr-gen-base: build-antlr-container
	${ANTLR_CMD} -Dlanguage=${language} -o pkg/${packageName}/gen /app/OpenFGALexer.g4 /app/OpenFGAParser.g4
