# OpenFGA Language - Java

Java implementation of ANTLR Grammar for the OpenFGA DSL and parser from and to the OpenFGA JSON Syntax

[![openfga-language on Maven Central](https://img.shields.io/maven-central/v/dev.openfga/openfga-language?style=flat-square&label=openfga-language)](https://central.sonatype.com/artifact/dev.openfga/openfga-language)

## Installation

From [dev.openfga/openfga-language on MavenCentral](https://central.sonatype.com/artifact/dev.openfga/openfga-language)

## Usage

### Transformer

```java
import dev.openfga.language.DslToJsonTransformer;
import dev.openfga.language.JsonToDslTransformer;

...

var dslString = """
model
  schema 1.2
type user
type folder
  relations
    define viewer: [user]
""";

// Transform from DSL to a JSON string
var generatedJsonString = new DslToJsonTransformer().transform(dslString);

// Transform from a JSON string to DSL
var generatedDsl = new JsonToDslTransformer().transform(generatedJsonString);

// Parses DSL returning results DSL without throwing
var result = new DslToJsonTransformer().parseDsl(dslString);
```

### Transform Mod to JSON

```java
import dev.openfga.language.FgaModTransformer;

...

var modFileContents = """
schema: "1.2"
contents:
  - core.fga
  - board.fga
  - wiki.fga
""";

// Parse fga.mod to a FgaModFile object
var mod = new FgaModTransformer(modFileContents).parse();

// Parse fga.mod to JSON string
var jsonMod = new FgaModTransformer(modFileContents).transform();
```

### Transform set of Modules To Model

Not yet implemented, but the [enchancement is in our issue backlog](https://github.com/openfga/language/issues/279).

### Validation

```java
import dev.openfga.language.validation.ModelValidator;
import dev.openfga.language.errors.DslErrorsException;

...

var dslString = """
model
  schema 1.2
type user
type folder
  relations
    define viewer: [user]
""";

try {
  ModelValidator.validateDsl(dslString);
} catch (DslErrorsException e) {
  // Handle generated errors
  return e.getErrors();
} catch (IOException e) {
  throw new RuntimeException(e);
}

```

## License

This project is licensed under the Apache-2.0 license. See the [LICENSE](https://github.com/openfga/language/blob/main/LICENSE) file for more info.
