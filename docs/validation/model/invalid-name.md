# Invalid Name

**Error Code:** `invalid-name`

**Category:** Naming Validation

## Summary

Type or relation names do not follow OpenFGA's naming conventions and format requirements, preventing proper parsing and evaluation.

## Description

OpenFGA enforces specific naming conventions for types and relations to ensure:
- Consistent parsing across different language implementations
- Predictable behavior in authorization evaluation
- Compatibility with OpenFGA's internal processing
- Clear, readable authorization models

Valid names must follow these rules:
- Start with a lowercase letter
- Contain only lowercase letters, numbers, and underscores
- Not contain spaces or special characters
- Be between 1 and 254 characters in length
- Not be empty or contain only whitespace

## Example

The following models would trigger this error:

### Invalid type names:
```
model
  schema 1.1

type User          # Error: starts with uppercase
type my-type       # Error: contains hyphen
type 123type       # Error: starts with number
type user type     # Error: contains space
type ""            # Error: empty name
```

### Invalid relation names:
```
model
  schema 1.1

type user

type document
  relations
    define Viewer: [user]       # Error: starts with uppercase
    define can-view: [user]     # Error: contains hyphen
    define 1viewer: [user]      # Error: starts with number
    define my viewer: [user]    # Error: contains space
```

**Error Message:** `Invalid name format: 'User'. Names must start with lowercase letter and contain only lowercase letters, numbers, and underscores`

## Resolution

Correct the names to follow OpenFGA naming conventions:

### Fix invalid type names:
```
model
  schema 1.1

type user          # Corrected from 'User'
type my_type       # Corrected from 'my-type'
type type_123      # Corrected from '123type'
type user_profile  # Corrected from 'user type'
type valid_name    # Corrected from empty name
```

### Fix invalid relation names:
```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]       # Corrected from 'Viewer'
    define can_view: [user]     # Corrected from 'can-view'
    define viewer_1: [user]     # Corrected from '1viewer'
    define my_viewer: [user]    # Corrected from 'my viewer'
```

### Steps to fix:

1. **Identify the invalid name:**
   - Check the error message for the specific invalid name
   - Note whether it's a type or relation name

2. **Apply naming conventions:**
   - Convert to lowercase
   - Replace spaces and special characters with underscores
   - Ensure the name starts with a letter
   - Keep length under 254 characters

3. **Update all references:**
   - Change the definition
   - Update any references to the renamed construct
   - Verify no broken references remain

4. **Test the model:**
   - Validate the updated model
   - Ensure authorization logic still works correctly

## Naming Convention Examples

### ✅ Valid names:
```
type user
type user_profile
type organization_member
type document_123
type api_key

type document
  relations
    define viewer
    define can_edit
    define owner_admin
    define level_1_access
```

### ❌ Invalid names:
```
type User              # Uppercase
type user-profile      # Hyphen
type organization member # Space
type 123document       # Starts with number
type user@domain       # Special character

type document
  relations
    define Viewer           # Uppercase
    define can-edit         # Hyphen
    define owner admin      # Space
    define 1st_level        # Starts with number
    define owner@company    # Special character
```

## Related Errors

- [`reserved-type-keywords`](./reserved-type-keywords.md) - When names use reserved keywords
- [`reserved-relation-keywords`](./reserved-relation-keywords.md) - When relation names use reserved keywords
- [`self-error`](./self-error.md) - Specific issues with 'self' and 'this'

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/name_validation.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java naming validation package

The validation uses regular expressions to check name format and applies consistent rules across all language implementations.
