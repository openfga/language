# Invalid Schema Version

**Error Code:** `invalid-schema-version`

**Category:** Schema Validation

## Summary

The schema version format is invalid or malformed, preventing proper model validation and execution.

## Description

OpenFGA schema versions must follow a specific format to ensure proper parsing and feature detection. Valid schema versions:
- Follow semantic versioning format (e.g., "1.1", "1.2")
- Use numeric values separated by dots
- Contain only supported version numbers
- Cannot be empty or contain invalid characters

Invalid schema version formats prevent the validation system from determining which features are available and which validation rules to apply.

## Example

The following models would trigger this error:

### Invalid version formats:
```
model
  schema v1.1        # Error: contains 'v' prefix
  
model
  schema 1.1.0.0     # Error: too many version parts

model
  schema 1.x         # Error: non-numeric version part

model
  schema ""          # Error: empty version string

model
  schema 1.1-beta    # Error: contains suffix
```

**Error Message:** `Invalid schema version format: 'v1.1'. Schema version must be in format 'X.Y'`

## Resolution

Use proper schema version format:

### Correct schema version formats:
```
model
  schema 1.1    # Valid: current recommended version

model
  schema 1.2    # Valid: current recommended version + module support
```

### Steps to fix:

1. **Identify the format issue:**
   - Check the error message for the specific format problem
   - Review the schema version declaration in your model

2. **Use correct format:**
   - Remove any prefixes (v, version, etc.)
   - Use only numeric values separated by a single dot
   - Remove any suffixes or additional version parts

3. **Choose appropriate version:**
   - Use `1.1` or `1.2` for new models (recommended)
   - Use `1.2` when using modules

4. **Update and validate:**
   - Correct the schema version format
   - Ensure your model features are compatible with the chosen version

## Valid Schema Version Examples

### ✅ Correct formats:
```
model
  schema 1.1

model
  schema 1.2
```

### ❌ Invalid formats:
```
model
  schema v1.1          # Prefix not allowed
  
model
  schema 1.1.0         # Too many parts
  
model
  schema 1.x           # Non-numeric
  
model
  schema 1.1-beta      # Suffix not allowed
  
model
  schema version 1.1   # Extra text
```

## Feature Compatibility Matrix

| Feature                              | Schema 1.0 | Schema 1.1 | Schema 1.2 |
|--------------------------------------|------------|------------|------------|
| Supported                            | ❌          | ✅          | ✅          |
| Basic relations                      | ✅          | ✅          | ✅          |
| Simple wildcards                     | ✅          | ✅          | ✅          |
| Wildcards                            | ✅          | ✅          | ✅          |
| Basic operations                     | ✅          | ✅          | ✅          |
| Type Restrictions                    | ❌          | ✅          | ✅          |
| Conditions                           | ❌          | ✅          | ✅          |
| Operator grouping `(a or (b and c))` | ❌          | ✅          | ✅          |
| Modules                              | ❌          | ❌          | ✅          |

> [!WARNING]
> Schema version `1.0` is no longer supported by OpenFGA. Models using this version must be updated to `1.1` or `1.2`. See the [Schema 1.1 Migration Guide](../migrations/schema1.0-to-schema1.1.md) for assistance.

## Related Errors

- [`schema-version-required`](./schema-version-required.md) - When no schema version is specified
- [`schema-version-unsupported`](./schema-version-unsupported.md) - When version is not supported
- [`invalid-schema`](./invalid-schema.md) - General schema structure issues

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/schema_validation.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java schema validation package

The validation uses regular expressions to check version format and ensures consistency across all language implementations.
