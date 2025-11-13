# Invalid Relation on Tupleset

**Error Code:** `invalid-relation-on-tupleset`

**Category:** Structural Validation

## Summary

A tuple-to-userset operation references an invalid or non-existent relation on the tupleset, preventing proper authorization evaluation.

## Description

Tuple-to-userset operations (e.g., `member from owner`) work by:
1. Finding tuples for the tupleset relation (`owner`)
2. Taking the user/object from those tuples  
3. Applying the computed userset relation (`member`) to that user/object

This error occurs when the relation specified in the tuple-to-userset operation doesn't exist on the expected type, or when the relation reference is malformed or invalid.

## Example

The following models would trigger this error:

### Non-existent relation on tupleset:
```
model
  schema 1.1

type user
type organization
  relations
    define member: [user]
    # Note: 'admin' relation is not defined

type document
  relations
    define owner: [organization]
    define reader: admin from owner  # Error: 'admin' doesn't exist on organization
```

### Malformed relation reference:
```
model
  schema 1.1

type user
type organization
  relations
    define member: [user]

type document
  relations
    define owner: [organization]
    define reader: member. from owner     # Error: malformed relation reference
    define viewer: .member from owner     # Error: invalid syntax
```

**Error Message:** `Invalid relation 'admin' on tupleset type 'organization' in tuple-to-userset operation`

## Resolution

Ensure the relation exists on the tupleset type and is properly referenced:

### Option 1: Define the missing relation

```
model
  schema 1.1

type user
type organization
  relations
    define member: [user]
    define admin: [user]      # Define the missing relation

type document
  relations
    define owner: [organization]
    define reader: admin from owner  # Now valid
```

### Option 2: Use an existing relation

```
model
  schema 1.1

type user
type organization
  relations
    define member: [user]

type document
  relations
    define owner: [organization]
    define reader: member from owner  # Use existing relation
```

### Option 3: Fix malformed syntax

```
model
  schema 1.1

type user
type organization
  relations
    define member: [user]

type document
  relations
    define owner: [organization]
    define reader: member from owner  # Clean, correct syntax
```

### Steps to fix:

1. **Identify the problematic relation reference:**
   - Check the error message for the specific relation and tupleset type
   - Locate the tuple-to-userset operation in your model

2. **Verify the tupleset type:**
   - Check if the relation exists on the tupleset type
   - Review the type definition to see available relations

3. **Choose resolution approach:**
   - Define the missing relation if it should exist
   - Use an existing relation if there was a typo
   - Fix syntax issues in the relation reference

4. **Test the authorization logic:**
   - Ensure the tuple-to-userset operation works as intended
   - Verify the authorization flow makes sense

## Valid Tuple-to-Userset Patterns

### ✅ Correct usage:
```
type user
type organization
  relations
    define member: [user]
    define admin: [user]

type document
  relations
    define owner: [organization]
    define reader: member from owner     # Valid: 'member' exists on organization
    define editor: admin from owner      # Valid: 'admin' exists on organization
```

### ❌ Invalid usage:
```
type organization
  relations
    define member: [user]
    # 'admin' relation not defined

type document
  relations
    define owner: [organization]
    define reader: admin from owner      # Error: 'admin' doesn't exist
    define editor: member. from owner    # Error: malformed syntax
    define viewer: from owner            # Error: missing relation
```

## Related Errors

- [`tupleuserset-not-direct`](./tupleuserset-not-direct.md) - When tupleset lacks direct assignment
- [`undefined-relation`](./undefined-relation.md) - When relations don't exist
- [`invalid-relation-type`](./invalid-relation-type.md) - When relation syntax is invalid

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/wildcard_validation.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java tuple-to-userset validation package

The validation checks that all relations referenced in tuple-to-userset operations exist on their respective types and follow proper syntax rules.
