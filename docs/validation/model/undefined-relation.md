# Undefined Relation

**Error Code:** `undefined-relation`

**Category:** Semantic Validation

## Summary

A relation is referenced in the model but is not defined anywhere, creating a broken reference that would cause runtime errors.

## Description

This error occurs when a relation definition references another relation that doesn't exist within the same type or any other type in the model. OpenFGA requires all relation references to be valid and resolvable at validation time to ensure the authorization model can function correctly.

Relation references can occur in:
- Computed usersets (`viewer`)
- Tuple-to-userset operations (`member from owner`)
- Complex operations (unions, intersections, differences)
- Grouping and mixing operators (`(((viewer or editor) and (admin or guest)) but not blocked)`)

## Example

The following model would trigger this error:

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]
    define editor: [user] or nonexistent_relation
```

**Error Location:** Line 8, `nonexistent_relation` in the editor relation definition.

**Error Message:** `Relation 'nonexistent_relation' is not defined on type 'document' (referenced in relation 'editor' of type 'document')`

## Resolution

Define the missing relation or correct the reference to an existing relation:

### Option 1: Define the missing relation

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]
    define nonexistent_relation: [user]  # Define the missing relation
    define editor: [user] or nonexistent_relation
```

### Option 2: Correct the reference

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]
    define editor: [user] or viewer  # Reference existing relation
```

### Steps to fix:

1. **Identify the undefined relation:**
   - Check the error message for the exact relation name
   - Note which type and relation contains the invalid reference

2. **Verify intended behavior:**
   - Determine if the relation should exist or if it's a typo
   - Review your authorization model design

3. **Choose resolution approach:**
   - Define the missing relation if it should exist
   - Correct the reference if it's a typo
   - Remove the reference if it's unnecessary

4. **Test the fix:**
   - Validate the model after making changes
   - Ensure the authorization logic still works as expected

## Common Causes

- **Typos in relation names:** `veiwer` instead of `viewer`
- **Case sensitivity:** `Viewer` instead of `viewer`
- **Missing relation definitions:** Referencing planned but unimplemented relations
- **Copy-paste errors:** Copying references from other models without definitions

## Related Errors

- [`undefined-type`](./undefined-type.md) - When a type is referenced but not defined
- [`missing-definition`](./missing-definition.md) - General missing definition error
- [`invalid-relation-type`](./invalid-relation-type.md) - When relation type is invalid

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/semantic_validation.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java semantic validation package

The validation performs deep traversal of all relation definitions to build a complete reference graph and identify broken links.
