# Relation No Entry Point

**Error Code:** `relation-no-entry-point`

**Category:** Semantic Validation

## Summary

A relation definition has no entry point, meaning there's no way to directly assign users to this relation, making it unreachable and non-functional.

## Description

Every relation in an OpenFGA model must have at least one "entry point" - a way for users to be assigned to that relation. Entry points can be:

1. **Direct assignment:** `[user]` or `[user, organization#member]`
2. **This keyword:** Allowing direct assignment through `this`
3. **Computed usersets with entry points:** Relations that eventually resolve to direct assignments

A relation without entry points creates a "dead end" in your authorization model where no users can ever satisfy the relation, regardless of the authorization data.

## Example

The following model would trigger this error:

```
model
  schema 1.1

type user

type document
  relations
    define viewer: editor
    define editor: admin  
    define admin: viewer  # Circular reference with no entry point
```

**Error Location:** All three relations (`viewer`, `editor`, `admin`) form a cycle with no entry point.

**Error Message:** `Relation 'viewer' on type 'document' has no entry point`

## Resolution

Add a direct assignment entry point to break the cycle and make relations reachable:

### Option 1: Add direct assignment

```
model
  schema 1.1

type user

type document
  relations
    define viewer: editor
    define editor: admin  
    define admin: [user] or viewer  # Add direct assignment entry point
```

### Option 2: Use 'this' for direct assignment

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user] or editor
    define editor: admin  
    define admin: this  # Allow direct assignment
```

### Steps to fix:

1. **Identify the problematic relation:**
   - Check which relation is reported as having no entry point
   - Trace the relation dependencies to understand the issue

2. **Analyze relation dependencies:**
   - Map out how relations reference each other
   - Look for circular dependencies or dead ends

3. **Add appropriate entry points:**
   - Add `[user]` or other type restrictions for direct assignment
   - Use `this` if the relation should allow direct assignment
   - Ensure at least one relation in any chain has an entry point

4. **Verify authorization logic:**
   - Ensure the entry points match your intended authorization model
   - Test that users can be properly assigned and checked

## Common Patterns

### Valid patterns with entry points:

```
# Direct assignment
define viewer: [user]

# Computed with entry point
define editor: [user] or viewer

# Tuple-to-userset with entry point
define reader: [user] or member from owner
```

### Invalid patterns (no entry points):

```
# Circular reference
define viewer: editor
define editor: viewer

# Chain with no entry
define viewer: editor  
define editor: admin
define admin: super_admin
define super_admin: viewer
```

## Related Errors

- [`cyclic-error`](./cyclic-error.md) - When relations form circular dependencies
- [`cyclic-relation`](./cyclic-relation.md) - Specific circular relation detection
- [`undefined-relation`](./undefined-relation.md) - When referenced relations don't exist

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/cycle_detection.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts` 
- Java implementation: Java semantic validation package

The validation uses depth-first search to detect cycles and analyze reachability from direct assignment entry points.
