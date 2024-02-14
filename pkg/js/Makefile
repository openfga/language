BINARY_NAME=openfga-language
DOCKER_BINARY=docker

all: build

init-deps:
	npm ci

build: all-tests
	npm run build

run: build
	npm run start

clean:
	npm run clean

test: init-deps
	npm run test

lint: init-deps
	npm run lint;npm run format:check

audit: init-deps
	npm audit

format: init-deps
	npm run lint:fix;npm run format:fix

all-tests: audit lint test
	npx madge --circular --exclude '^dist|^gen' . --extensions ts,js