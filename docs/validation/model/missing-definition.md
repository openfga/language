# Missing Definition

**Error Code:** `missing-definition`

**Category:** Semantic Validation

## Summary

A reference to a type, relation, or other construct exists in the model but the corresponding definition is missing, creating a broken reference.

## Description

This error is a general category for missing definitions that can occur in various contexts:
- Type references without corresponding type definitions
- Relation references without corresponding relation definitions
- Condition references without corresponding condition definitions
- Cross-module references to undefined constructs

Missing definitions prevent the authorization model from functioning correctly because the system cannot resolve references to undefined constructs during evaluation.

## Example

The following models would trigger this error:

### Missing type definition:
```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user, organization#member]  # Error: 'organization' not defined
```

### Missing relation definition:
```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]
    define editor: [user] or nonexistent  # Error: relation 'nonexistent' not defined
```

### Missing condition definition:
```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user with missing_condition]  # Error: condition not defined
```

**Error Message:** `The relation 'nonexistent' does not exist` or `Type 'organization' is not defined`

## Resolution

Define the missing construct or correct the reference:

### Option 1: Define the missing construct

```
model
  schema 1.1

type user

type organization  # Define the missing type
  relations
    define member: [user]

condition missing_condition(user: User) {  # Define the missing condition
  user.active == true
}

type document
  relations
    define viewer: [user, organization#member]
    define editor: [user] or viewer  # Reference existing relation
    define admin: [user with missing_condition]
```

### Option 2: Correct or remove the reference

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]              # Remove invalid reference
    define editor: [user] or viewer    # Reference existing relation
    define admin: [user]               # Remove condition reference
```

### Steps to fix:

1. **Identify the missing definition:**
   - Check the error message for the specific missing construct
   - Note the type of definition needed (type, relation, condition)

2. **Determine intended behavior:**
   - Verify if the construct should exist or if it's a typo
   - Review your authorization model design

3. **Choose resolution approach:**
   - Define the missing construct if it should exist
   - Correct the reference if it's a typo
   - Remove the reference if it's unnecessary

4. **Test the fix:**
   - Validate the model after making changes
   - Ensure the authorization logic still works as expected

## Common Patterns

### ✅ Valid references:
```
type user
type organization
  relations
    define member: [user]

type document
  relations
    define viewer: [user, organization#member]  # Both constructs defined
```

### ❌ Invalid references:
```
type user

type document
  relations
    define viewer: [user, undefined_type#member]  # Type not defined
    define editor: undefined_relation             # Relation not defined
```

## Related Errors

- [`undefined-type`](./undefined-type.md) - Specifically for missing type definitions
- [`undefined-relation`](./undefined-relation.md) - Specifically for missing relation definitions
- [`condition-not-defined`](./condition-not-defined.md) - Specifically for missing condition definitions

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/semantic_validation.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java semantic validation package

The validation performs comprehensive reference checking across all constructs in the authorization model to identify broken references.
