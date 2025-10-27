# Cyclic Error

**Error Code:** `cyclic-error`

**Category:** Semantic Validation

## Summary

A general cycle detection error indicating circular dependencies in the authorization model that prevent proper evaluation of permissions.

## Description

This error represents the general category of cyclic dependencies in authorization models. Cycles can occur in various contexts:
- **Relation cycles:** Relations that reference each other in circular patterns
- **Type dependency cycles:** Types that depend on each other in circular ways
- **Computed userset cycles:** Circular references in permission computation

Cyclic dependencies prevent OpenFGA from resolving permissions to a finite set of users or conditions, making authorization evaluation impossible or leading to infinite loops.

## Example

The following models would trigger this error:

### Direct relation cycle:
```
model
  schema 1.1

type user

type document
  relations
    define viewer: editor
    define editor: viewer  # Direct cycle: viewer → editor → viewer
```

### Indirect relation cycle:
```
model
  schema 1.1

type user

type document
  relations
    define viewer: editor
    define editor: admin
    define admin: viewer  # Indirect cycle: viewer → editor → admin → viewer
```

### Complex permission cycle:
```
model
  schema 1.1

type user
type group
  relations
    define member: [user] or admin_member
    define admin_member: member  # Cycle in permission computation
```

**Error Message:** `Cyclic dependency detected in authorization model`

## Resolution

Break the circular dependency by restructuring the authorization model:

### Option 1: Create proper hierarchy

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]
    define editor: [user] or viewer  # Editor includes viewer
    define admin: [user] or editor   # Admin includes editor (and transitively viewer)
```

### Option 2: Add direct assignments to break cycles

```
model
  schema 1.1

type user
type group
  relations
    define member: [user]                    # Direct assignment breaks cycle
    define admin_member: [user] or member    # Clear hierarchy
```

### Option 3: Restructure authorization logic

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]
    define editor: [user]  # Independent permissions
    define admin: [user]   # No hierarchical dependencies
```

### Steps to fix:

1. **Identify the cycle:**
   - Trace the dependency chain to understand the circular reference
   - Map out which constructs depend on each other

2. **Analyze intended authorization flow:**
   - Determine the correct permission hierarchy
   - Identify which permissions should include others

3. **Restructure the model:**
   - Break the cycle by removing one problematic reference
   - Create clear hierarchical relationships
   - Add direct assignments where needed

4. **Validate the solution:**
   - Ensure all necessary permissions are still achievable
   - Test that the authorization logic meets requirements

## Common Authorization Patterns

### ✅ Valid hierarchical structures:
```
# Clear hierarchy
define viewer: [user]
define editor: [user] or viewer
define admin: [user] or editor

# Independent permissions
define read: [user]
define write: [user]
define delete: [user]

# Tuple-to-userset without cycles
define reader: [user] or member from owner
define owner: [organization#admin]
```

### ❌ Invalid cyclic structures:
```
# Direct cycle
define viewer: editor
define editor: viewer

# Indirect cycle
define a: b
define b: c
define c: a

# Self-referential cycle
define viewer: viewer or [user]
```

## Related Errors

- [`cyclic-relation`](./cyclic-relation.md) - Specific circular relation dependencies
- [`relation-no-entry-point`](./relation-no-entry-point.md) - When cycles prevent entry points
- [`undefined-relation`](./undefined-relation.md) - Missing relations in cycle chains

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/cycle_detection.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java cycle detection package

The validation uses depth-first search algorithms to detect cycles in the authorization model's dependency graph.
