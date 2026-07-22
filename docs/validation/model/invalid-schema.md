# Invalid Schema

**Error Code:** `invalid-schema`

**Category:** Schema Validation

## Summary

The overall schema structure is invalid or malformed, preventing proper model parsing and validation.

## Description

This error occurs when the authorization model's schema structure doesn't conform to OpenFGA's schema requirements. Unlike specific schema version errors, this represents fundamental structural problems with the schema that prevent basic parsing and validation.

Common schema structure issues include:
- Missing required schema components
- Malformed schema declarations
- Invalid schema syntax or formatting
- Structural inconsistencies that violate OpenFGA's schema rules

## Example

The following models would trigger this error:

### Missing model declaration:
```
schema 1.1  # Error: Missing 'model' declaration

type user

type document
  relations
    define viewer: [user]
```

### Malformed schema structure:
```
model
  # Error: Schema declaration without version
  schema

type user
```

### Invalid schema syntax:
```
model {  # Error: Invalid syntax for model declaration
  schema: 1.1
}

type user
```

**Error Message:** `Invalid schema structure: missing required model declaration`

## Resolution

Fix the schema structure to conform to OpenFGA's requirements:

### Option 1: Add missing model declaration

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]
```

### Option 2: Fix schema syntax

```
model
  schema 1.1  # Proper format: schema followed by version

type user

type document
  relations
    define viewer: [user]
```

### Steps to fix:

1. **Identify the structural issue:**
   - Check the error message for specific schema structure problems
   - Review the beginning of your model file for proper format

2. **Follow OpenFGA schema format:**
   - Start with `model` declaration
   - Follow with `schema X.Y` version specification
   - Use proper indentation and syntax

3. **Validate basic structure:**
   - Ensure model declaration comes first
   - Verify schema version is properly specified
   - Check that type definitions follow schema declaration

4. **Test the corrected structure:**
   - Validate the model after fixing schema structure
   - Ensure the model parses correctly

## Valid Schema Structure

### ✅ Correct schema format:
```
model
  schema 1.1

type user
  relations
    define profile_owner: [user]

type document
  relations
    define viewer: [user]
    define editor: [user] or viewer
```

### ❌ Invalid schema formats:
```
# Missing model declaration
schema 1.1
type user

# Wrong syntax
model {
  schema: 1.1
}
type user

# Missing schema version  
model
  schema
type user
```

## Related Errors

- [`schema-version-required`](./schema-version-required.md) - When schema version is missing
- [`invalid-schema-version`](./invalid-schema-version.md) - When version format is invalid
- [`invalid-syntax`](./invalid-syntax.md) - General syntax issues

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/schema_validation.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java schema validation package

The validation performs structural checks during the initial parsing phase to ensure the model follows OpenFGA's basic schema requirements.
