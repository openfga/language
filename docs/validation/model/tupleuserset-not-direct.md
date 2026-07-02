# Tuple-to-Userset Not Direct

**Error Code:** `tupleuserset-not-direct`

**Category:** Structural Validation

## Summary

A tuple-to-userset operation requires the tupleset relation to have direct assignment capability, but the referenced relation does not allow direct assignment.

## Description

Tuple-to-userset operations (e.g., `member from owner`) work by:
1. Finding tuples for the tupleset relation (`owner`) where the object is the current object being evaluated
2. Taking the user from those tuples
3. Finding tuples where each of those users are objects and the relation is the userset relation (`member`)

For this to work correctly, the tupleset relation must support direct assignment (have an entry point that is `[type]` assignments), because the operation needs to retrieve actual user objects from the tupleset.

## Example

The following model would trigger this error:

```
model
  schema 1.1

type user

type organization
  relations
    define member: [user]

type folder
  relations
    define owner: [organization] or admin  # Error: 'owner' has computed relation and is used in tuple-to-userset
    define reader: member from owner  # Error: 'owner' has computed relation and is used in tuple-to-userset
    
type photo
  relations
    define owner: [organization#member]  # Error: 'owner' has computed relation and is used in tuple-to-userset
    define reader: member from owner  # Error: 'owner' has computed relation and is used in tuple-to-userset

type resource
  relations
    define owner: [organization:*]  # Error: 'owner' has computed relation and is used in tuple-to-userset
    define reader: member from owner  # Error: 'owner' has computed relation and is used in tuple-to-userset

type document
  relations
    define owner: [organization]
    define admin: owner
    define reader: member from admin  # Error: 'owner' has computed relation and is used in tuple-to-userset
```

**Error Location:** Line 11, `member from owner` where `owner` lacks direct assignment.

**Error Message:** `Tuple-to-userset operation requires tupleset relation 'owner' to have direct assignment capability`

## Resolution

Add direct assignment capability to the tupleset relation:

### Option 1: Add direct assignment to the base relation and remove any indirections

```
model
  schema 1.1

type user
type organization
  relations
    define member: [user]

type document
  relations
    define owner: [organization]      # Only direct assignment with no indirection
    define reader: member from owner  # Now valid
```

### Steps to fix:

1. **Identify the problematic tupleset:**
   - Check which relation is used as the tupleset in the tuple-to-userset operation
   - Verify if that relation has direct assignment capability

2. **Analyze the intended behavior:**
   - Understand what types of objects should be directly assigned to the tupleset
   - Consider the authorization flow you want to achieve

3. **Choose resolution approach:**
   - Add direct assignment to the existing tupleset relation

4. **Validate the authorization logic:**
   - Ensure the changes maintain your intended authorization behavior
   - Test that the tuple-to-userset operation works as expected

## Related Errors

- [`relation-no-entry-point`](./relation-no-entry-point.md) - When relations lack entry points
- [`invalid-relation-on-tupleset`](./invalid-relation-on-tupleset.md) - Invalid tupleset relation usage
- [`undefined-relation`](./undefined-relation.md) - When tupleset relation doesn't exist

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/wildcard_validation.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java tuple-to-userset validation package

The validation analyzes the tupleset relation to ensure it has at least one direct assignment path that can produce concrete user objects.
