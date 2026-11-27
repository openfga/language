# Schema Version Required

**Error Code:** `schema-version-required`

**Category:** Schema Validation

## Summary

A JSON model has no `schema_version`, or a module file is validated on its own.

## Description

The schema version decides which features a model can use. `1.1` and `1.2` are supported, and modules need `1.2`.

## Example

### JSON

This JSON model has no `schema_version` field:

```json
{
  "type_definitions": [
    { "type": "user" },
    {
      "type": "document",
      "relations": { "viewer": { "this": {} } },
      "metadata": { "relations": { "viewer": { "directly_related_user_types": [ { "type": "user" } ] } } }
    }
  ]
}
```

**Error Message:** `schema version required`

### DSL

In the DSL, the `schema` line is part of the syntax. A model without it fails to parse with a syntax error (`mismatched input 'type' expecting 'schema'`) and never reaches this check:

```
model
type user
```

The error is also reported when a single module file is validated on its own, because module files have no `schema` line: the version is set in `fga.mod`. Validate the combined model instead.

## Resolution

Add the schema version:
- In JSON, add `"schema_version"` at the top level, as below
- In the DSL, add `schema 1.1` on the line after `model`
- For modules, set `schema: '1.2'` in `fga.mod`

```json
{
  "schema_version": "1.1",
  "type_definitions": [
    { "type": "user" },
    {
      "type": "document",
      "relations": { "viewer": { "this": {} } },
      "metadata": { "relations": { "viewer": { "directly_related_user_types": [ { "type": "user" } ] } } }
    }
  ]
}
```

## Related Errors

- [`schema-version-unsupported`](./schema-version-unsupported.md) - The model declares version `1.0`
- [`invalid-schema`](./invalid-schema.md) - The model declares a version other than `1.0`, `1.1` or `1.2`

## Implementation Notes

- Go: `pkg/go/validation/schema_validation.go`
- JavaScript: `pkg/js/validator/validate-dsl.ts`
- Java: `pkg/java/src/main/java/dev/openfga/language/validation/ModelValidator.java`
