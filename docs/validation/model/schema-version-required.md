# Schema Version Required

**Error Code:** `schema-version-required`

**Category:** Schema Validation

## Summary

Every OpenFGA authorization model must specify a schema version to ensure proper parsing and validation behavior.

## Description

OpenFGA requires all authorization models to explicitly declare a schema version. The schema version determines which features are available and how the model should be interpreted. Without a schema version, the validation system cannot determine which validation rules to apply or which language features are supported.

Schema versions follow semantic versioning (e.g., "1.1", "1.2") and each version may introduce new capabilities or modify existing behavior.

## Example

The following model would trigger this error:

```
model
type user

type document
  relations
    define viewer: [user]
```

**Error Location:** The model declaration without schema specification.

## Resolution

Add a schema version declaration to your model:

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]
```

### Steps to fix:

1. **Choose appropriate schema version:**
   - Do not use `1.0` it is no longer supported by OpenFGA
   - Use `1.1` or `1.2` for new models (recommended)
   - Use `1.2` when using modules

2. **Add schema declaration:**
   - Place the schema declaration immediately after the `model` keyword
   - Use proper indentation (2 spaces)

3. **Verify compatibility:**
   - Ensure your model features are supported in the chosen schema version
   - Test your model with the updated schema

## Related Errors

- [`schema-version-unsupported`](./schema-version-unsupported.md) - When an unsupported version is specified
- [`invalid-schema-version`](./invalid-schema-version.md) - When the version format is invalid

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/schema_validation.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java validation package
