# Assignable Relation Must Have Type

**Error Code:** `assignable-relation-must-have-type`

**Category:** Wildcard Validation

## Summary

An assignable relation in a type restriction must specify a type, as wildcard-only references without type context are invalid.

## Description

This error occurs when relation references in type restrictions lack the required type specification. OpenFGA requires clear type context for all assignable relations to ensure:
- Proper authorization evaluation and user resolution
- Clear semantic meaning in permission checks
- Consistent behavior across different implementations
- Predictable relation traversal during authorization

When a relation reference lacks type context, the authorization system cannot determine which type's relation should be evaluated, leading to ambiguous authorization behavior.

## Example

The following models would trigger this error:

### Missing type in relation reference:
```
model
  schema 1.1

type user
type organization
  relations
    define member: [user]

type document
  relations
    define viewer: [#member]        # Error: missing type before #member
    define editor: [user, #admin]   # Error: missing type for #admin reference
```

### Incomplete type restriction:
```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user, :*]       # Error: missing type before wildcard
    define editor: [#]              # Error: incomplete type#relation reference
```

**Error Message:** `Assignable relation '#member' must specify a type`

## Resolution

Add the required type specification to relation references:

### Option 1: Add missing type to relation references

```
model
  schema 1.1

type user
type organization
  relations
    define member: [user]
    define admin: [user]

type document
  relations
    define viewer: [organization#member]     # Added 'organization' type
    define editor: [user, organization#admin] # Added 'organization' type
```

### Option 2: Use direct type assignments instead

```
model
  schema 1.1

type user
type organization
  relations
    define member: [user]

type document
  relations
    define viewer: [user, organization#member] # Clear type specifications
    define editor: [user]                      # Simple direct assignment
```

### Steps to fix:

1. **Identify incomplete relation references:**
   - Check the error message for specific relation references missing types
   - Locate all `#relation` references without type prefixes

2. **Determine the intended type:**
   - Review your authorization model to understand which type should be referenced
   - Consider the business logic and permission flow

3. **Add type specifications:**
   - Add the appropriate type prefix: `type#relation`
   - Ensure the referenced type and relation exist in your model

4. **Validate the authorization logic:**
   - Ensure the type#relation combination makes sense for your use case
   - Test that the authorization behavior matches expectations

## Valid Type and Relation Patterns

### ✅ Correct type#relation usage:
```
type user
type organization
  relations
    define member: [user]
    define admin: [user]

type group
  relations
    define member: [user]

type document
  relations
    define viewer: [user, organization#member]          # Clear type#relation
    define editor: [organization#admin, group#member]   # Multiple clear references
    define owner: [user]                                # Direct type reference
```

### ❌ Invalid incomplete references:
```
type document
  relations
    define viewer: [#member]                # Missing type
    define editor: [user, #admin]           # Missing type
    define admin: [organization#]           # Missing relation
    define owner: [#]                       # Missing both type and relation
```

## Related Errors

- [`invalid-wildcard-error`](./invalid-wildcard-error.md) - General wildcard usage issues
- [`type-wildcard-relation`](./type-wildcard-relation.md) - Conflicting wildcard and relation usage
- [`undefined-type`](./undefined-type.md) - When referenced types don't exist

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/wildcard_validation.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java wildcard validation package

The validation parses type restrictions and ensures all relation references include proper type specifications.
