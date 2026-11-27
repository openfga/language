# Multiple Modules in File

**Error Code:** `multiple-modules-in-file`

**Category:** Multi-file Validation

## Summary

A model's metadata assigns more than one module to the same file. A file can hold only one module, so the model cannot be written back as modular DSL.

## Description

In a modular model, each `.fga` file starts with one `module <name>` line, and an `fga.mod` file sets the schema version and lists the files. When the files are combined, each type and condition records its module and file in `metadata` (`module` and `source_info.file`), and relations added with `extend type` record their own.

The validator groups this metadata by file and reports each file that has more than one module. In a DSL module file, a second `module` line is a syntax error, not this error.

## Example

Both types point to `single-file.fga`, but belong to different modules:

```json
{
  "schema_version": "1.2",
  "type_definitions": [
    {
      "type": "user",
      "metadata": { "module": "user_management", "source_info": { "file": "single-file.fga" } }
    },
    {
      "type": "document",
      "relations": { "owner": { "this": {} } },
      "metadata": {
        "module": "document_management",
        "source_info": { "file": "single-file.fga" },
        "relations": { "owner": { "directly_related_user_types": [ { "type": "user" } ] } }
      }
    }
  ]
}
```

**Error Message:** `file single-file.fga would contain multiple module definitions (user_management, document_management) when transforming to DSL. Only one module can be defined per file.`

## Resolution

Give each module its own file, or use one module for the file.

### Option 1: Create separate files for each module

```json
{
  "schema_version": "1.2",
  "type_definitions": [
    {
      "type": "user",
      "metadata": { "module": "user_management", "source_info": { "file": "user-management.fga" } }
    },
    {
      "type": "document",
      "relations": { "owner": { "this": {} } },
      "metadata": {
        "module": "document_management",
        "source_info": { "file": "document-management.fga" },
        "relations": { "owner": { "directly_related_user_types": [ { "type": "user" } ] } }
      }
    }
  ]
}
```

As modular DSL, this is two module files and an `fga.mod` that lists them:

**fga.mod:**
```yaml
schema: '1.2'
contents:
  - user-management.fga
  - document-management.fga
```

**user-management.fga:**
```
module user_management

type user
```

**document-management.fga:**
```
module document_management

type document
  relations
    define owner: [user]
```

Module files have no `schema` line; the version comes from `fga.mod`. A type from another module, such as `user`, is referenced by its name.

### Option 2: Use one module for the file

Use the same `module` name, for example `application`, for everything in the file. As modular DSL, the file is:

**single-file.fga:**
```
module application

type user

type document
  relations
    define owner: [user]
```

## Related Errors

- [`schema-version-required`](./schema-version-required.md) - Reported when a single module file is validated on its own

## Implementation Notes

- Go: `pkg/go/validation/multi_file_validation.go`
- JavaScript: `pkg/js/validator/validate-dsl.ts`
- Java: `pkg/java/src/main/java/dev/openfga/language/validation/ModelValidator.java`
