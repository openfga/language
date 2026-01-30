# AGENTS.md

This file provides guidance to AI coding agents working with code in this repository.

## Project Overview

OpenFGA Language is a multi-language library providing ANTLR grammar for the OpenFGA DSL with parsers, transformers (DSL↔JSON), and validators. It's part of the OpenFGA fine-grained authorization ecosystem.

**Published packages:**
- npm: `@openfga/syntax-transformer`
- Go: `github.com/openfga/language/pkg/go`
- Maven: `dev.openfga/openfga-language`

## Build Commands

From repository root (requires Docker for ANTLR generation):

```bash
make build          # Build all packages (Go, JS, Java)
make test           # Run tests for all packages
make lint           # Lint all packages and test data
make antlr-gen      # Generate ANTLR parsers for all languages
```

### Go Package (`pkg/go/`)

```bash
make test-go        # Run Go tests with race detection
make lint-go        # Lint with golangci-lint (auto-fix enabled)
make audit-go       # Check vulnerabilities with govulncheck

# Run specific test
cd pkg/go && go test ./transformer/... -run TestDslToJson -v
```

### JavaScript Package (`pkg/js/`)

```bash
make test-js        # Run Jest tests
make lint-js        # ESLint + Prettier check
make audit-js       # npm audit

# Run specific test
cd pkg/js && npx jest tests/validator.test.ts -v
```

### Java Package (`pkg/java/`)

```bash
make test-java      # Run Gradle tests
make lint-java      # Lint Java code

# Run specific test
cd pkg/java && ./gradlew test --tests "dev.openfga.language.DslToJsonShould"
```

## Architecture

```
├── OpenFGALexer.g4, OpenFGAParser.g4   # ANTLR grammar (source of truth)
├── pkg/
│   ├── go/                              # Go implementation
│   │   ├── gen/                         # Generated ANTLR parser
│   │   ├── transformer/                 # DSL↔JSON transformation
│   │   ├── graph/                       # Graph-based utilities (Go only)
│   │   └── validation/                  # Validation rules
│   ├── js/                              # TypeScript implementation
│   │   ├── gen/                         # Generated ANTLR parser
│   │   ├── transformer/                 # DSL↔JSON transformation
│   │   └── validator/                   # Syntactic + semantic validation
│   └── java/                            # Java implementation
│       └── src/main/gen/                # Generated ANTLR parser
└── tests/data/                          # Shared test fixtures (YAML/JSON)
```

**Feature parity:**
| Feature | Go | JS | Java |
|---------|----|----|------|
| DSL↔JSON Transformer | ✓ | ✓ | ✓ |
| Syntactic Validation | ✓ | ✓ | ✓ |
| Semantic Validation | ✗ | ✓ | ✓ |
| Graph Utilities | ✓ | ✗ | ✗ |

## Critical: Graph Package (`pkg/go/graph/`)

The `pkg/go/graph/` directory contains graph-based utilities used by the main OpenFGA server (`github.com/openfga/openfga`) for its typesystem, Check algorithm, and ListObjects algorithm to establish connections based on the authorization model.

**Changes to this directory require special care:**
- Test changes thoroughly with `go test ./graph/... -v`
- Verify that `github.com/openfga/openfga` still works correctly with the updated library
- Breaking changes here can affect core authorization functionality

## Development Notes

- **ANTLR code generation** requires Docker (`make build-antlr-container` runs automatically)
- **Shared test data** in `tests/data/` is used by all three language implementations
- **Generated code** in `*/gen/` directories should not be manually edited
- Changes to grammar files (`*.g4`) require running `make antlr-gen` before building

## Commit Convention

Follow [Conventional Commits](https://www.conventionalcommits.org/): `feat`, `fix`, `docs`, `chore`, `test`, `refactor`, `perf`, `build`, `ci`, `style`