BINARY_NAME=openfga-language
DOCKER_BINARY=docker

all: build

build:
	./gradlew build

clean:
	./gradlew clean

test:
	./gradlew check

lint:
	./gradlew lint

audit:
	echo "java audit Not implemented"

format:
	./gradlew fmt

all-tests: build audit lint test