# Invalid Wildcard Error

**Error Code:** `invalid-wildcard-error`

**Category:** Wildcard Validation

## Summary

Wildcard usage in type restrictions is invalid or incorrectly formatted, violating OpenFGA's wildcard syntax rules.

## Description

OpenFGA supports wildcard type restrictions using the `*` symbol to indicate "any user of this type." However, wildcards have specific usage rules and constraints:

- Wildcards must be used correctly within type restrictions
- Certain combinations of wildcards and relations are not allowed
- Schema version compatibility affects wildcard availability
- Wildcards cannot be used in contexts where specific user identity is required

## Example

The following models would trigger this error:

### Invalid wildcard syntax:
```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user:*]  # Error: Invalid wildcard syntax
```

### Wildcard in wrong context:
```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user*]   # Error: Malformed wildcard
    define editor: [*]       # Error: Wildcard without type
```

**Error Message:** `Invalid wildcard usage in type restriction`

## Resolution

Use correct wildcard syntax according to OpenFGA specifications:

### Correct wildcard usage:
```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user, user:*]  # Valid: specific users or any user
    define editor: [user]          # Valid: no wildcard needed
```

### Alternative without wildcards:
```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]  # Simple user restriction without wildcards
    define editor: [user]
```

### Steps to fix:

1. **Identify the invalid wildcard usage:**
   - Check the error message for the specific location
   - Review the wildcard syntax in your type restrictions

2. **Understand wildcard requirements:**
   - Verify schema version supports wildcards
   - Check if wildcards are appropriate for your use case

3. **Correct the syntax:**
   - Use proper wildcard format: `[type:*]`
   - Ensure wildcards are used in valid contexts
   - Consider if wildcards are necessary for your authorization model

4. **Test the authorization logic:**
   - Verify wildcard behavior matches expectations
   - Test with actual authorization data

## TODO: Complete Wildcard Usage Guidelines

<!-- TODO: Add comprehensive guidance for:
- All valid wildcard syntax patterns
- Schema version compatibility requirements
- When to use wildcards vs specific user restrictions
- Performance implications of wildcard usage
- Advanced wildcard patterns and combinations
- Wildcard behavior in different relation types
-->

## Common Wildcard Patterns

### ✅ Valid patterns:
```
define viewer: [user:*]              # Any user
define editor: [user, org#member:*]  # Specific users or any org member
define admin: [user]                 # No wildcard needed
```

### ❌ Invalid patterns:
```
define viewer: [*]          # Missing type
define editor: [user*]      # Wrong syntax
define admin: [user:*, *]   # Mixed invalid syntax
```

## Related Errors

- [`type-wildcard-relation`](./type-wildcard-relation.md) - Wildcard and relation conflicts
- [`allowed-type-schema-10`](./allowed-type-schema-10.md) - Schema version requirements
- [`assignable-relation-must-have-type`](./assignable-relation-must-have-type.md) - Type requirement issues

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/wildcard_validation.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java wildcard validation package

The validation checks wildcard syntax during relation parsing and verifies compatibility with schema version and authorization model structure.
