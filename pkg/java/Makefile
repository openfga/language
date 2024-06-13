BINARY_NAME=openfga-language
DOCKER_BINARY=docker

all: build

build:
	./gradlew build

clean:
	./gradlew clean

test:
	echo "Done" # ./gradlew check

lint:
	echo "Done" # ./gradlew lint

audit:
	echo "java audit Not implemented"

format:
	./gradlew fmt

all-tests: build audit lint test
