# Invalid Schema

**Error Code:** `invalid-schema`

**Category:** Schema Validation

## Summary

The model declares a schema version the validator does not recognise.

## Description

`1.1` and `1.2` are the supported versions. A version that parses but is neither of those is reported as `invalid-schema`, with the version itself as the symbol.

Two neighbouring conditions have codes of their own. Version `1.0` is recognised and retired, so it reports [`schema-version-unsupported`](./schema-version-unsupported.md), and a model carrying no version at all reports [`schema-version-required`](./schema-version-required.md).

A malformed `schema` line never reaches validation. `schema v1.1`, `schema 1.1.0`, `schema` with no version, and a file with no `model` declaration all fail during DSL transformation with a syntax error and no error code attached.

## Example

### DSL

```
model
  schema 0.9

type user
```

**Error Message:** `invalid schema 0.9`

The position covers the version itself: line 1, columns 9 to 12.

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
- [`schema-version-unsupported`](./schema-version-unsupported.md) - Version `1.0`, recognised but retired
- [`invalid-syntax`](./invalid-syntax.md) - Syntax problems, including a malformed `schema` line

## Implementation Notes

- Go: `ValidateSchemaVersion` in `pkg/go/validation/schema_validation.go`
- JavaScript: `validate-dsl.ts`, through `createInvalidSchemaVersionError` in `util/exceptions.ts`
- Java: `ModelValidator`, through `ValidationErrorsBuilder.raiseInvalidSchemaVersion`

All three tag the finding `invalid-schema`. The message and the code are pinned in `tests/data/dsl-semantic-validation-cases.yaml` and `tests/data/json-validation-cases.yaml`.
