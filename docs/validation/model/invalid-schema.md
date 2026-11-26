# Invalid Schema

**Error Code:** `invalid-schema`

**Category:** Schema Validation

## Summary

The model declares a schema version other than `1.0`, `1.1` or `1.2`.

## Description

`1.1` and `1.2` are the supported versions. Version `1.0` reports [`schema-version-unsupported`](./schema-version-unsupported.md), a missing version reports [`schema-version-required`](./schema-version-required.md), and any other version reports `invalid-schema`.

In the DSL, a malformed version such as `schema v1.1` or `schema 1.1.0` is a syntax error.

## Example

### DSL

```
model
  schema 0.9

type user
```

**Error Location:** The version, `0.9`.

**Error Message:** `invalid schema 0.9`

### JSON

```json
{
  "schema_version": "1.3"
}
```

**Error Message:** `invalid schema 1.3`

## Resolution

Declare a supported version:

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]
```

Use `1.2` if the model is split across files with `module` declarations, `1.1` otherwise.

## Related Errors

- [`schema-version-required`](./schema-version-required.md) - No schema version is declared
- [`schema-version-unsupported`](./schema-version-unsupported.md) - The model declares version `1.0`

## Implementation Notes

- Go: `pkg/go/validation/schema_validation.go`
- JavaScript: `pkg/js/validator/validate-dsl.ts`
- Java: `pkg/java/src/main/java/dev/openfga/language/validation/ModelValidator.java`
