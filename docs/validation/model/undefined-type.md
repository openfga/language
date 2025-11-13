# Undefined Type

**Error Code:** `undefined-type`

**Category:** Semantic Validation

## Summary

A type is referenced in the model but is not defined anywhere, creating a broken reference that would cause runtime errors.

## Description

This error occurs when a relation definition or type restriction references a type that doesn't exist in the model. OpenFGA requires all type references to be valid and resolvable at validation time to ensure the authorization model can function correctly.

Type references can occur in:
- Direct type restrictions: `[undefined_type]`
- Relation type restrictions: `[undefined_type#member]`
- Public type restrictions: `[undefined_type:*]`

## Example

The following model would trigger this error:

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user, undefined_type]  # Error: type not defined
    define editor: [organization#member]   # Error: 'organization' type not defined
    define admin: [employee:*]            # Error: 'employee' type not defined
```

**Error Location:** Line 7, `undefined_type` and `organization` in the relation definitions.

**Error Message:** `Type 'undefined_type' is not defined (referenced in relation 'viewer' of type 'document')`

## Resolution

Define the missing type or correct the reference to an existing type:

### Option 1: Define the missing type

```
model
  schema 1.1

type user

type undefined_type  # Define the missing type
  relations
    define member: [user]

type organization    # Define the missing type
  relations
    define member: [user]

type document
  relations
    define viewer: [user, undefined_type]
    define editor: [organization#member]
```

### Option 2: Correct the reference

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]              # Remove invalid reference
    define editor: [user]              # Use existing type
```

### Option 3: Import from another module (if applicable)

```
module current_module
  schema 1.1

import other_module

type user

type document
  relations
    define viewer: [user, other_module.organization#member]  # Reference imported type
```

### Steps to fix:

1. **Identify the undefined type:**
   - Check the error message for the exact type name
   - Note which relation and type contains the invalid reference

2. **Verify intended behavior:**
   - Determine if the type should exist or if it's a typo
   - Review your authorization model design

3. **Choose resolution approach:**
   - Define the missing type if it should exist
   - Correct the reference if it's a typo
   - Import the type if it exists in another module
   - Remove the reference if it's unnecessary

4. **Test the fix:**
   - Validate the model after making changes
   - Ensure the authorization logic still works as expected

## Common Causes

- **Typos in type names:** `usr` instead of `user`
- **Case sensitivity:** `User` instead of `user`
- **Missing type definitions:** Referencing planned but unimplemented types
- **Module issues:** Referencing types from unimported modules
- **Copy-paste errors:** Copying references from other models without definitions

## Related Errors

- [`undefined-relation`](./undefined-relation.md) - When a relation is referenced but not defined
- [`missing-definition`](./missing-definition.md) - General missing definition error
- [`invalid-type`](./invalid-type.md) - When type format is invalid

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/semantic_validation.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java semantic validation package

The validation builds a map of all defined types and cross-references them with type usage in relation definitions and type restrictions.
