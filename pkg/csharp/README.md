# OpenFGA Language - C#

C# implementation of ANTLR Grammar for the OpenFGA DSL and parser from and to the OpenFGA JSON Syntax

[![OpenFga.Language on NuGet](https://img.shields.io/nuget/v/OpenFga.Language?style=flat-square&label=OpenFga.Language)](https://www.nuget.org/packages/OpenFga.Language)

## Installation

From [OpenFga.Language on NuGet](https://www.nuget.org/packages/OpenFga.Language)

```bash
dotnet add package OpenFga.Language
```

## Usage

### Transformer

```csharp
using OpenFga.Language;

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
var generatedJsonString = new DslToJsonTransformer().Transform(dslString);

// Transform from a JSON string to DSL
var generatedDsl = new JsonToDslTransformer().Transform(generatedJsonString);

// Parse DSL returning results without throwing
var result = new DslToJsonTransformer().ParseDsl(dslString);
```

### Transform Mod to JSON

```csharp
using OpenFga.Language;

...

var modFileContents = """
schema: "1.2"
contents:
  - core.fga
  - board.fga
  - wiki.fga
""";

// Parse fga.mod to a FgaModFile object
var mod = new FgaModTransformer(modFileContents).Parse();

// Parse fga.mod to JSON string
var jsonMod = new FgaModTransformer(modFileContents).Transform();
```

### Transform set of Modules To Model

```csharp
using OpenFga.Language;

...

// Transform an array of modules, and their contents, into a singular model
var model = new DslToJsonTransformer().TransformModuleFilesToModel(new List<DslToJsonTransformer.ModuleFile>
{
    new DslToJsonTransformer.ModuleFile
    {
        Name = "core.fga",
        Contents = @"module core
  type user"
    },
    new DslToJsonTransformer.ModuleFile
    {
        Name = "board.fga",
        Contents = @"module board
  type board"
    },
    new DslToJsonTransformer.ModuleFile
    {
        Name = "wiki.fga",
        Contents = @"module wiki
  type wiki"
    }
}, "1.2");

var jsonString = Json.Stringify(model);
```

### Validation

```csharp
using OpenFga.Language.Validation;
using OpenFga.Language.Errors;

...

var dslString = """
model
  schema 1.2
type user
type folder
  relations
    define viewer: [user]
""";

try 
{
  ModelValidator.ValidateDsl(dslString);
  Console.WriteLine("Model is valid!");
}
catch (DslErrorsException ex)
{
  Console.WriteLine($"Validation failed: {ex.Message}");
}
```

## License

This project is licensed under the Apache-2.0 license.