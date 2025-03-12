# OpenFGA Language - GoLang

GoLang implementation of ANTLR Grammar for the OpenFGA DSL and parser from and to the OpenFGA JSON Syntax

[![Go Reference](https://pkg.go.dev/badge/github.com/openfga/language.svg)](https://pkg.go.dev/github.com/openfga/language/pkg/go)
[![GitHub Release](https://img.shields.io/github/v/release/openfga/language?include_prereleases&filter=pkg%2Fgo%2Fv*&label=openfga-language)](https://github.com/openfga/language/tree/main/pkg/go)

## Installation

```bash
go get github.com/openfga/language/pkg/go
```

## Usage

### Transformer

```go
import "github.com/openfga/language/pkg/go/transformer"

...

dslString := `model
  schema 1.1
type user
type folder
  relations
    define viewer: [user]`

// Transform from DSL syntax to the OpenFGA Authorization Model Protobuf format
generatedProto, err := transformer.TransformDSLToProto(dslString)

// Transform from DSL to a JSON string
generatedJsonString, err := transformer.TransformDSLToJSONString(dslString)

// Transform from a JSON string to DSL
generatedDsl, err := transformer.TransformJSONStringToDSL(generatedJsonString)
```

### Transform Mod to JSON

```go
import "github.com/openfga/language/pkg/go/transformer"

...

modFileContents := `schema: "1.2"
contents:
  - core.fga
  - board.fga
  - wiki.fga`

// Transform from fga.mod to a JSON object for parsing
jsonModFile, err := transformer.TransformModFile(modFileContents)
```

### Transform Modules To Model

```go
import "github.com/openfga/language/pkg/go/transformer"

...

// Transform an array of modules, and their contents, into a singular model
model, err := transformer.TransformModuleFilesToModel([]transformer.ModuleFile{
		{
			Name: "core.fga",
			Contents: `module core
  type user`,
		},
    {
			Name: "board.fga",
			Contents: `module core
  type board`,
		},
    {
			Name: "wiki.fga",
			Contents: `module core
  type wiki`,
		},

	}, "1.2")


jsonString, _ := json.MarshalIndent(model, "", "  ")

/*
{
  "schema_version": "1.2",
  "type_definitions": [
    {
      "type": "user",
      "metadata": {
        "module": "core",
        "source_info": {
          "file": "core.fga"
        }
      }
    },
    {
      "type": "board",
      "metadata": {
        "module": "core",
        "source_info": {
          "file": "board.fga"
        }
      }
    },
    {
      "type": "wiki",
      "metadata": {
        "module": "core",
        "source_info": {
          "file": "wiki.fga"
        }
      }
    }
  ]
}
*/

```

### Validation

Not yet implemented, but the [feature is in our issue backlog](https://github.com/openfga/language/issues/99).

## License

This project is licensed under the Apache-2.0 license. See the [LICENSE](https://github.com/openfga/language/blob/main/LICENSE) file for more info.
