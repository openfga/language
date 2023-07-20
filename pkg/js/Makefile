BINARY_NAME=openfga-language
DOCKER_BINARY=docker

all: build

build: lint audit format
	npm run build

run: build
	npm run start

clean:
	npm run clean

test:
	npm run test

lint:
	npm run lint

audit:
	npm run audit

format:
	npm run format:fix