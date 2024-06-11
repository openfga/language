# OpenFGA Language

ANTLR Grammar for the OpenFGA DSL and parser from and to the OpenFGA JSON Syntax

[![Go Reference](https://pkg.go.dev/badge/github.com/openfga/language.svg)](https://pkg.go.dev/github.com/openfga/language/pkg/go)
[![Release](https://img.shields.io/github/v/release/openfga/language?sort=semver&color=green)](https://github.com/openfga/language/releases)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](./LICENSE)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fopenfga%2Flanguage.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fopenfga%2Flanguage?ref=badge_shield)
[![Join our community](https://img.shields.io/badge/slack-cncf_%23openfga-40abb8.svg?logo=slack)](https://openfga.dev/community)
[![Twitter](https://img.shields.io/twitter/follow/openfga?color=%23179CF0&logo=twitter&style=flat-square "@openfga on Twitter")](https://twitter.com/openfga)

## Table of Contents
- [About OpenFGA](#about)
- [Resources](#resources)
- [Packages](#packages)
- [About This Repo](#about-this-repo)
- [Installation](#installation)
- [Usage](#usage)
- [CLI](#cli)
- [Community Parsers](#community-parsers)
- [Contributing](#contributing)
- [License](#license)


## About
[OpenFGA](https://openfga.dev) is an open source Fine-Grained Authorization solution inspired by [Google's Zanzibar paper](https://research.google/pubs/pub48190/). It was created by the FGA team at [Auth0](https://auth0.com) based on [Auth0 Fine-Grained Authorization (FGA)](https://fga.dev), available under [a permissive license (Apache-2)](https://github.com/openfga/rfcs/blob/main/LICENSE) and welcomes community contributions.

OpenFGA is designed to make it easy for application builders to model their permission layer, and to add and integrate fine-grained authorization into their applications. OpenFGA’s design is optimized for reliability and low latency at a high scale.

## Resources

- [OpenFGA Documentation](https://openfga.dev/docs)
- [OpenFGA API Documentation](https://openfga.dev/api/service)
- [Twitter](https://twitter.com/openfga)
- [OpenFGA Community](https://openfga.dev/community)
- [Zanzibar Academy](https://zanzibar.academy)
- [Google's Zanzibar Paper (2019)](https://research.google/pubs/pub48190/)

## Packages

| Packages | Package Manager |
| -------- | --------------- |
| JS       | [![@openfga/syntax-transformer on npm](https://img.shields.io/npm/v/@openfga/syntax-transformer.svg?label=@openfga/syntax-transformer&style=flat-square)](https://www.npmjs.com/package/@openfga/syntax-transformer) |
| Go       | [![OpenFGA Language (Go) on GitHub](https://img.shields.io/github/v/release/openfga/language/pkg/go?label=openfga-language&style=flat-square)](https://github.com/openfga/language/tree/main/pkg/go) |
| Java     | [![openfga-language on Maven Central](https://img.shields.io/maven-central/v/dev.openfga/openfga-language?style=flat-square&label=openfga-language)](https://central.sonatype.com/artifact/dev.openfga/openfga-language) |

## About This Repo
This repo contains everything needed to interact with the OpenFGA Authorization Models schema versions 1.1+, in multiple languages (currently, Go and JS are supported).

| Feature    | Implemented in ANTLR |
|------------|----------------------|
| Basic DSL  | ✅                    |
| Nesting    | ✅ (partial, see [#113](https://github.com/openfga/language/issues/113)) |
| Conditions | ✅                    |

| Feature                                               | Go                   | JS           | Java                  |
|-------------------------------------------------------|----------------------|--------------|-----------------------|
| Transformer from the DSL to JSON and from JSON to DSL | ✅                    | ✅           | ✅                    |
| Syntactic Model Validations                           | ✅                    | ✅           | ✅                    |
| Schema 1.0                                            | ❌                    | ❌ (supported in [v0.1.6](https://www.npmjs.com/package/@openfga/syntax-transformer/v/0.1.6) and below) | ❌ |
| Schema 1.1                                            | ✅                    | ✅ (supported in [v0.2.0](https://www.npmjs.com/package/@openfga/syntax-transformer) and above) | ✅         |
| Schema 1.2 (Modular Models)                           | ✅                    | ✅           | ✅                     |
| Semantic Model Validations                            | ❌ (planned, see [#99](https://github.com/openfga/language/issues/99)) | ✅      | ✅      |
| Graphing & Utility Methods                            | ❌ (planned)          | ❌ (planned) | ❌ (planned)           |


## Installation
### Go
```bash
go get github.com/openfga/language/pkg/go
```

### Node
```bash
npm install @openfga/syntax-transformer
```

### Java
From [dev.openfga/openfga-language on MavenCentral](https://central.sonatype.com/artifact/dev.openfga/openfga-language)

## Usage

### Transformer

#### Go
```go
import "github.com/openfga/language/pkg/go/transformer"

dslString := `model
  schema 1.1
type user
type folder
  relations
    define viewer: [user]`

// Transform from DSL to a JSON string
generatedJsonString, err := transformer.TransformDSLToJSONString(dslString)

// Transform from a JSON string to DSL
generatedDsl, err := transformer.TransformJSONStringToDSL(generatedJsonString)
```

#### Node
```typescript
import { transformer } from "@openfga/syntax-transformer"

let dslString = `model
  schema 1.1
type user
type folder
  relations
    define viewer: [user]`;

// Transform from DSL to a JSON string
const generatedJsonString = transformer.transformDSLToJSONString(dslString)

// Transform from a JSON string to DSL
const generatedDsl = transformer.transformJSONStringToDSL(generatedJsonString)
```

#### Java
```java
import dev.openfga.language.DslToJsonTransformer;
import dev.openfga.language.JsonToDslTransformer;

var dslString = """model
  schema 1.1
type user
type folder
  relations
    define viewer: [user]""";

// Transform from DSL to a JSON string
var generatedJsonString = new DslToJsonTransformer().transform(dslString);

// Transform from a JSON string to DSL
var generatedDsl = new JsonToDslTransformer().transform(generatedJsonString);
```


## CLI

Use the [FGA CLI](https://github.com/openfga/cli)

## Community Parsers

| Repo                                                                         | License                                                                            | Maintainers                                                                           | Language                                  | Schema v1.0                                                                 | Schema v1.1                                                                                          | Package Managers                                                                                                                                                                                                                                                                                               | Other Links                                                                                                                    |
|------------------------------------------------------------------------------|------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------|-------------------------------------------|-----------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------|
| [openfga language (syntax-transformer)](https://github.com/openfga/language) | [Apache-2.0](https://github.com/openfga/language/blob/main/LICENSE)                | [@openfga](https://github.com/orgs/openfga/people)                                    | ANTLR (Go and Typescript implementations) | <[0.1.5](https://www.npmjs.com/package/@openfga/syntax-transformer/v/0.1.5) | Yes  (v0.0.8+)                                                                                       | [![npm:@openfga/syntax-transformer](https://img.shields.io/npm/v/@openfga/syntax-transformer.svg?style=flat)](https://www.npmjs.com/package/@openfga/syntax-transformer) - ![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/openfga/language?label=go)                                |                                                                                                                                |
| [openfga-dsl-parser](https://github.com/maxmindlin/openfga-dsl-parser)       | [Apache-2.0](https://github.com/maxmindlin/openfga-dsl-parser/blob/master/LICENSE) | [@maxmindlin](https://github.com/maxmindlin) - [@dblclik](https://github.com/dblclik) | Rust                                      | Yes                                                                         | No                                                                                                   | [![crates:openfga-dsl-parser](https://img.shields.io/crates/v/openfga-dsl-parser.svg?style=flat)](https://crates.io/crates/openfga-dsl-parser)[![pypi:openfga-dsl-parser-python](https://img.shields.io/pypi/v/openfga-dsl-parser-python.svg?style=flat)](https://pypi.org/project/openfga-dsl-parser-python/) | [WASM](https://github.com/dblclik/openfga-dsl-parser-wasm) - [Python](https://github.com/maxmindlin/openfga-dsl-parser-python) |
| [openfga-rs](https://github.com/iammathew/openfga-rs)                        | [Apache-2.0](https://github.com/iammathew/openfga-rs/blob/main/LICENSE.md)         | [@iammathew](https://github.com/iammathew)                                            | Rust                                      | Yes                                                                         | No                                                                                                   |                                                                                                                                                                                                                                                                                                                |                                                                                                                                |
| [openfga-dsl-parser](https://github.com/craigpastro/openfga-dsl-parser)      | [Apache-2.0](https://github.com/craigpastro/openfga-dsl-parser/blob/main/LICENSE)  | [@craigpastro](https://github.com/craigpastro)                                        | ANTLR & Go                                | Yes                                                                         | Partial (requires self). [Supports nesting](https://github.com/openfga/syntax-transformer/issues/34) | ![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/craigpastro/openfga-dsl-parser?label=go)                                                                                                                                                                                             |                                                                                                                                |

## Community Wrapper
| Repo                                                                  | License                                                                 | Maintainers                                        | Language   | Schema v1.0 | Schema v1.1 | Package Managers                                                                                                                                                           | Other Links |
|-----------------------------------------------------------------------|-------------------------------------------------------------------------|----------------------------------------------------|------------|-------------|-------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|
| [fga cli](https://github.com/openfga/cli)                             | [Apache-2.0](https://github.com/openfga/cli/blob/main/LICENSE)          | [@openfga](https://github.com/orgs/openfga/people) | Go         | No          | Yes         | ![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/openfga/cli?label=go)                                                                            |             |
| [fga-transformer-cli](https://github.com/ozee-io/fga-transformer-cli) | [MIT](https://github.com/ozee-io/fga-transformer-cli/blob/main/LICENSE) | [@ozee-io](https://github.com/orgs/ozee-io/people) | Javascript | Yes         | Yes         | [![npm:@openfga/syntax-transformer](https://img.shields.io/npm/v/@ozee-io/fga-transformer-cli.svg?style=flat)](https://www.npmjs.com/package/@ozee-io/fga-transformer-cli) |             |


## Contributing

See [CONTRIBUTING](https://github.com/openfga/.github/blob/main/CONTRIBUTING.md).

## Author

[OpenFGA](https://github.com/openfga)

## License

This project is licensed under the Apache-2.0 license. See the [LICENSE](https://github.com/openfga/cli/blob/main/LICENSE) file for more info.
