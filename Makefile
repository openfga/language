DOCKER_BINARY=docker
ANTLR_DOCKER_IMAGE=docker.io/rhamzeh/antlr
ANTLR_CMD=${DOCKER_BINARY} run -t --rm -v ${PWD}:/app ${ANTLR_DOCKER_IMAGE}

#### Global #####

.PHONY: all
all: build

.PHONY: build
build: build-go build-js

.PHONY: test
test: test-go test-js

.PHONY: lint
lint: lint-go lint-js

#### Go #####

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

#### TypeScript #####

.PHONY: antlr-gen-js
antlr-gen-js:
	${ANTLR_CMD} -Dlanguage=TypeScript -o src/js/gen /app/OpenFGA.g4

.PHONY: build-js
build-js: antlr-gen-js
	$(MAKE) -C src/js build

.PHONY: run-js
run-js: antlr-gen-js
	$(MAKE) -C src/js run

.PHONY: clean-js
clean-js:
	$(MAKE) -C src/js clean

.PHONY: test-js
test-js: antlr-gen-js
	$(MAKE) -C src/js test

.PHONY: lint-js
lint-js: antlr-gen-js
	$(MAKE) -C src/js lint

.PHONY: audit-js
audit-js: antlr-gen-js
	$(MAKE) -C src/js audit

.PHONY: format-js
format-js: antlr-gen-js
	$(MAKE) -C src/js format