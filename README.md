# OpenFGA Language

ANTLR Grammar for the OpenFGA DSL and parser from and to the OpenFGA JSON Syntax

[![Go Reference](https://pkg.go.dev/badge/github.com/openfga/language.svg)](https://pkg.go.dev/github.com/openfga/language)
[![Release](https://img.shields.io/github/v/release/openfga/language?sort=semver&color=green)](https://github.com/openfga/language/releases)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](./LICENSE)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fopenfga%2Flanguage.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fopenfga%2Flanguage?ref=badge_shield)
[![Discord Server](https://img.shields.io/discord/759188666072825867?color=7289da&logo=discord "Discord Server")](https://discord.com/channels/759188666072825867/930524706854031421)
[![Twitter](https://img.shields.io/twitter/follow/openfga?color=%23179CF0&logo=twitter&style=flat-square "@openfga on Twitter")](https://twitter.com/openfga)

## Table of Contents
- [About OpenFGA](#about)
- [Resources](#resources)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)


## About
[OpenFGA](https://openfga.dev) is an open source Fine-Grained Authorization solution inspired by [Google's Zanzibar paper](https://research.google/pubs/pub48190/). It was created by the FGA team at [Auth0](https://auth0.com) based on [Auth0 Fine-Grained Authorization (FGA)](https://fga.dev), available under [a permissive license (Apache-2)](https://github.com/openfga/rfcs/blob/main/LICENSE) and welcomes community contributions.

OpenFGA is designed to make it easy for application builders to model their permission layer, and to add and integrate fine-grained authorization into their applications. OpenFGA’s design is optimized for reliability and low latency at a high scale.

## Resources

- [OpenFGA Documentation](https://openfga.dev/docs)
- [OpenFGA API Documentation](https://openfga.dev/api/service)
- [Twitter](https://twitter.com/openfga)
- [OpenFGA Discord Community](https://discord.gg/8naAwJfWN6)
- [Zanzibar Academy](https://zanzibar.academy)
- [Google's Zanzibar Paper (2019)](https://research.google/pubs/pub48190/)

## About This Repo
This repo contains everything needed to interact with the OpenFGA Authorization Models, in multiple languages (currently, Go and JS are supported).

It is planned to include:
- [ ] ANTLR Grammar for the DSL (in progress)
- [x] Transformer from the DSL to JSON and from JSON to DSL
- [ ] Validation functions that allow ensuring the model is valid
- [ ] Graphing Utility functions that allow visualizing the model
- [ ] Model complexity calculator to understand how complex the model is

## Installation
TBD

## Usage
TBD

## Contributing

See [CONTRIBUTING](https://github.com/openfga/.github/blob/main/CONTRIBUTING.md).

## Author

[OpenFGA](https://github.com/openfga)

## License

This project is licensed under the Apache-2.0 license. See the [LICENSE](https://github.com/openfga/cli/blob/main/LICENSE) file for more info.
