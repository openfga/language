# Type Wildcard Relation

**Error Code:** `type-wildcard-relation`

**Category:** Wildcard Validation

## Summary

A type restriction cannot simultaneously specify both wildcard access and a specific relation, as this creates an ambiguous permission specification.

## Description

This error occurs when a type restriction attempts to combine wildcard access (`type:*`) with a specific relation reference (`type#relation`) in the same type restriction. This combination is invalid because:

- Wildcards grant access to all users of a type and are of the form `type:*`
- Relation references grant access to specific users through a relation and are of the form `type#relation`
- `type:*#relation` is not a valid syntax in OpenFGA nor is `type#relation:*`

OpenFGA requires clear, unambiguous permission specifications to ensure predictable authorization behavior.

## Example

The following model would trigger this error:

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
    define viewer: [organization:*#member]  # Error: wildcard + relation
    define editor: [organization#admin:*]   # Error: wildcard + relation
```

**Error Message:** `Type restriction cannot combine wildcard ':*' with relation reference '#member'`

## Resolution

Choose either wildcard access or relation-based access, but not both:

### Option 1: Use wildcard without relation, and vice versa

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
    define viewer: [user:*]              # Allows granting all users access to a particular document (once the tuple is written)
    define editor: [organization#member] # Allows granting members of an organization access to a particular document (once the tuple is written)
```

### Option 2: Use a combined approach if both are needed, but in separate restrictions

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
    define viewer: [user, organization#member]        # Direct users or org members
    define editor: [user:*, organization#admin] # All org users or specific admins
```

### Steps to fix:

1. **Identify the conflicting restriction:**
   - Check the error message for the specific type restriction
   - Locate the problematic wildcard + relation combination

2. **Determine intended access pattern:**
   - Decide if you want wildcard access (all users of type)
   - Or specific relation-based access (users through specific relation)

3. **Choose appropriate syntax:**
   - Use `[type:*]` for wildcard access to all users of that type
   - Use `[type#relation]` for access through specific relations
   - Use separate restrictions if you need both patterns

4. **Test the authorization logic:**
   - Verify the access pattern matches your intended permissions
   - Ensure the authorization behavior is clear and predictable

## Valid Wildcard and Relation Patterns

### ✅ Correct usage:
```
type user

type organization
  relations
    define member: [user]
    define admin: [user]

type document
  relations
    # Wildcard access (any user of type)
    define public_viewer: [user:*]
    
    # Relation access (specific users through relations)
    define member_viewer: [organization#member]
    define admin_editor: [organization#admin]
    
    # Combined in separate restrictions
    define flexible_access: [user:*, organization#member]
    
    # Mixed patterns
    define comprehensive_access: [user, user:*, organization#member]
```

### ❌ Invalid combinations:
```
type document
  relations
    # Cannot combine wildcard with relation in same restriction
    define viewer: [organization:*#member]    # Invalid
    define editor: [user#admin:*]             # Invalid
    define admin: [group:*#owner]             # Invalid
```

## Related Errors

- [`invalid-wildcard-error`](./invalid-wildcard-error.md) - General wildcard usage issues
- [`assignable-relation-must-have-type`](./assignable-relation-must-have-type.md) - Type requirement issues
- [`allowed-type-schema-10`](./allowed-type-schema-10.md) - Schema version compatibility

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/wildcard_validation.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java wildcard validation package

The validation parses type restrictions and checks for conflicting wildcard and relation syntax within the same restriction.
