# Invalid Type

**Error Code:** `invalid-type`

**Category:** Semantic Validation

## Summary

A type reference is invalid, malformed, or uses incorrect format, preventing proper authorization evaluation.

## Description

This error occurs when type references don't follow OpenFGA's type specification rules. Common issues include:
- Malformed type names that don't follow naming conventions
- Invalid syntax in type restrictions
- Incorrect cross-module type references
- Type references that violate schema constraints

OpenFGA requires all type references to follow specific formatting and naming rules to ensure consistent parsing and evaluation across different implementations.

## Example

The following models would trigger this error:

### Invalid type name format:
```
model
  schema 1.1

type User          # Error: starts with uppercase
type my-type       # Error: contains hyphen  
type 123invalid    # Error: starts with number

type document
  relations
    define viewer: [User]        # Error: references invalid type name
    define editor: [my-type]     # Error: references invalid type name
```

### Invalid type reference syntax:
```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user.]       # Error: trailing dot
    define editor: [.user]       # Error: leading dot
    define admin: [user#]        # Error: incomplete relation reference
```

**Error Message:** `Invalid type format: 'User'. Type names must follow naming conventions`

## Resolution

Correct the type names and references to follow OpenFGA standards:

### Fix type naming:
```
model
  schema 1.1

type user          # Corrected from 'User' 
type my_type       # Corrected from 'my-type'
type valid_type    # Corrected from '123invalid'

type document
  relations
    define viewer: [user]        # Now references valid type
    define editor: [my_type]     # Now references valid type
```

### Fix type reference syntax:
```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]              # Clean type reference
    define editor: [user]              # Another valid reference
    define admin: [user] or viewer     # Valid computed userset
```

### Steps to fix:

1. **Identify the invalid type:**
   - Check the error message for the specific invalid type
   - Locate the type definition or reference causing the issue

2. **Apply correct naming conventions:**
   - Use lowercase letters, numbers, and underscores only
   - Start with a letter, not a number
   - Avoid special characters and spaces

3. **Fix syntax issues:**
   - Remove trailing dots, extra characters
   - Ensure proper cross-module reference format if applicable
   - Use clean, simple type names

4. **Update all references:**
   - Change the type definition if it's a definition issue
   - Update all places where the type is referenced
   - Verify no broken references remain

## Valid Type Patterns

### ✅ Correct type usage:
```
type user
type organization  
type user_profile
type api_key_123

type document
  relations
    define viewer: [user]                    # Simple type reference
    define editor: [user, organization]      # Multiple type references
    define admin: [user] or viewer           # Type with computed userset
```

### ❌ Invalid type usage:
```
type User              # Uppercase
type my-type           # Hyphen
type 123type           # Starts with number
type user@domain       # Special character

type document
  relations
    define viewer: [User.]      # Malformed reference
    define editor: [.user]      # Invalid syntax
    define admin: [user#]       # Incomplete reference
```

## TODO: Advanced Type Validation Guidelines

<!-- TODO: Add guidance for:
- Complex type naming strategies for large authorization models
- Cross-module type reference best practices
- Type versioning and compatibility considerations
- Performance implications of different type structures
- Migration strategies for fixing invalid type names in existing models
-->

## Related Errors

- [`invalid-name`](./invalid-name.md) - General naming convention issues
- [`undefined-type`](./undefined-type.md) - When referenced types don't exist
- [`reserved-type-keywords`](./reserved-type-keywords.md) - When type names use reserved words

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/semantic_validation.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java semantic validation package

The validation checks type name format using regular expressions and validates type reference syntax during parsing.
