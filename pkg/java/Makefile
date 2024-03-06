BINARY_NAME=openfga-language
DOCKER_BINARY=docker

all: build

build:
	./gradlew build

clean:
	gradlew clean

test:
	gradlew check

lint:
	echo "java lint Not implemented"

audit:
	echo "java audit Not implemented"

format:
	echo "java format Not implemented"

all-tests: build audit lint test