# OpenFGA Language - JS

Javascript implementation of ANTLR Grammar for the OpenFGA DSL and parser from and to the OpenFGA JSON Syntax

[![@openfga/syntax-transformer](https://img.shields.io/npm/v/%40openfga%2Fsyntax-transformer/beta?label=%40openfga%2Fsyntax-transformer&style=flat-square)](https://www.npmjs.com/package/@openfga/syntax-transformer)

## Installation

```bash
npm install @openfga/syntax-transformer
```

### Supported Node.js Versions

For details on the supported Node.js versions and our support policy, see [SUPPORTED_RUNTIMES.md](./SUPPORTED_RUNTIMES.md).

## Usage

### Transformer

Example transform DSL model to JSON, and from JSON to DSL.

```typescript
import { transformer } from "@openfga/syntax-transformer";

let dslString = `model
  schema 1.1
type user
type folder
  relations
    define viewer: [user]`;

// Transform from DSL model to a JSON object
const generatedJsonObject = transformer.transformDSLToJSONObject(dslString);

// Transform from DSL to a JSON string
const generatedJsonString = transformer.transformDSLToJSONString(dslString);

// Transform from a JSON string to DSL
const generatedDsl = transformer.transformJSONStringToDSL(generatedJsonString);
```

### Transform Mod File to JSON

```typescript
import { transformer } from "@openfga/syntax-transformer";

...

const modFileContents = `schema: "1.2"
contents:
  - core.fga
  - board.fga
  - wiki.fga`

// Transform from fga.mod to an object
const jsonModFile = transformer.TransformModFile(modFileContents)
```

### Transform set of Modules To Model

```typescript
import { transformer } from "@openfga/syntax-transformer";

...

const files: transformer.ModuleFile[] = [];
files.push({
    name: "core.fga",
    contents: `module core
    type user`
  },
  {
    name: "board.fga",
    contents: `module core
    type board`
  },
  {
    name: "wiki.fga",
		contents: `module core
    type wiki`
  }
);

// Compile module files into a complete model
const jsonModel = transformer.transformModuleFilesToModel(files, "1.2");

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

```typescript
import { errors, validator } from "@openfga/syntax-transformer";

...

let dslString = `model
  schema 1.2
type user
type folder
  relations
    define viewer: [user]`;

// Attempt validation of model
try {
  validator.validateDSL(dslString);
} catch (err) {
  if (err instanceof errors.BaseMultiError) {
    // Handle generated errors
  } else {
    console.error("Unhandled Exception: " + err);
  }
}

```

## License

This project is licensed under the Apache-2.0 license. See the [LICENSE](https://github.com/openfga/language/blob/main/LICENSE) file for more info.
