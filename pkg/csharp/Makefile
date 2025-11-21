.PHONY: all build clean test lint audit format all-tests
all: build

build:
	dotnet build src/OpenFga.Language/OpenFga.Language.csproj
	dotnet build test/OpenFga.Language.Tests/OpenFga.Language.Tests.csproj

clean:
	dotnet clean src/OpenFga.Language/OpenFga.Language.csproj
	dotnet clean test/OpenFga.Language.Tests/OpenFga.Language.Tests.csproj

test:
	dotnet test test/OpenFga.Language.Tests/OpenFga.Language.Tests.csproj

lint:
	dotnet format --verify-no-changes src/OpenFga.Language/OpenFga.Language.csproj
	dotnet format --verify-no-changes test/OpenFga.Language.Tests/OpenFga.Language.Tests.csproj

audit:
	echo "dotnet audit Not implemented"

format:
	dotnet format src/OpenFga.Language/OpenFga.Language.csproj
	dotnet format test/OpenFga.Language.Tests/OpenFga.Language.Tests.csproj

all-tests: build audit lint test