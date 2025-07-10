# Self Error

**Error Code:** `self-error`

**Category:** Naming Validation

## Summary

Type or relation names cannot use reserved keywords 'self' or 'this' as they have special meaning in OpenFGA's authorization logic.

## Description

OpenFGA reserves the keywords 'self' and 'this' for special purposes in authorization models:

- **`this`**: Used in relation definitions to indicate direct assignment capability
- **`self`**: Reserved for potential future use and internal OpenFGA operations

Using these reserved words as type names or relation names can cause parsing conflicts, ambiguous references, and unexpected behavior in authorization checks.

## Example

The following models would trigger this error:

### Invalid type name:
```
model
  schema 1.1

type self  # Error: cannot use 'self' as type name
  relations
    define viewer: [user]
```

### Invalid relation name:
```
model
  schema 1.1

type user

type document
  relations
    define this: [user]  # Error: cannot use 'this' as relation name
```

**Error Message:** `A type cannot be named 'self' or 'this'` or `A relation cannot be named 'self' or 'this'`

## Resolution

Choose different names that don't conflict with reserved keywords:

### Fix type names:
```
model
  schema 1.1

type user  # Use descriptive, non-reserved names

type document
  relations
    define viewer: [user]
```

### Fix relation names:
```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]  # Use descriptive, non-reserved names
```

### Steps to fix:

1. **Identify reserved word usage:**
   - Check error message for specific location
   - Look for 'self' or 'this' used as names (not keywords)

2. **Choose appropriate replacements:**
   - Use descriptive names that reflect the purpose
   - Follow naming conventions (lowercase, no spaces)
   - Avoid other reserved words

3. **Update all references:**
   - Change the definition
   - Update any references to the renamed type/relation
   - Verify no broken references remain

4. **Test the model:**
   - Validate the updated model
   - Ensure authorization logic still works correctly

## Valid Usage vs Invalid Usage

### ❌ Invalid - 'this' as relation name:
```
type document
  relations
    define this: [user]  # 'this' cannot be a relation name
```

### ✅ Valid - descriptive names:
```
type user_profile  # Clear, descriptive type names
  relations
    define owner: [user]
    define viewer: [user] or owner
```

## Common Alternatives

Instead of reserved words, consider these alternatives:

| Reserved Word | Alternative Names |
|---------------|-------------------|
| `self` | `owner`, `user_profile`, `identity`, `principal` |
| `this` | `current`, `owner`, `direct`, `assigned` |

## Related Errors

- [`reserved-type-keywords`](./reserved-type-keywords.md) - Other reserved type keywords
- [`reserved-relation-keywords`](./reserved-relation-keywords.md) - Other reserved relation keywords  
- [`invalid-name`](./invalid-name.md) - General invalid naming issues

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/name_validation.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java naming validation package

The validation checks occur during the parsing phase before semantic analysis begins.
