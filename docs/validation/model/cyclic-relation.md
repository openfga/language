# Cyclic Relation

**Error Code:** `cyclic-relation`

**Category:** Semantic Validation

## Summary

A circular dependency has been detected in relation definitions, creating an infinite loop that prevents proper authorization evaluation.

## Description

This error occurs when relations reference each other in a circular pattern, creating an infinite loop during authorization evaluation. OpenFGA must be able to resolve all relation dependencies to a finite set of directly assigned users or computed values.

Circular dependencies can occur:
- **Direct cycles:** `A → B → A`
- **Indirect cycles:** `A → B → C → A` 
- **Self-referential cycles:** `A → A`

Unlike [`relation-no-entry-point`](./relation-no-entry-point.md), this error specifically focuses on detecting cycles in the relation dependency graph, even when entry points exist.

## Example

The following model would trigger this error:

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user] or editor
    define editor: admin
    define admin: viewer  # Creates cycle: viewer → editor → admin → viewer
```

**Error Location:** The cycle involves multiple relations forming a circular dependency.

**Error Message:** `Cyclic relation dependency detected involving relations: viewer, editor, admin`

## Resolution

Break the circular dependency by removing or restructuring one of the relation references:

### Option 1: Remove problematic reference

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user] or editor
    define editor: [user]  # Remove reference to admin
    define admin: [user] or editor
```

### Option 2: Restructure hierarchy

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

### Steps to fix:

1. **Identify the cycle:**
   - Review the error message to see which relations form the cycle
   - Map out the dependency chain

2. **Analyze intended authorization hierarchy:**
   - Determine the correct permission hierarchy
   - Identify which direction relationships should flow

3. **Break the cycle:**
   - Remove one problematic reference
   - Restructure to create a proper hierarchy
   - Ensure the authorization logic still meets requirements

4. **Validate the solution:**
   - Check that all necessary permissions are still achievable
   - Verify no new cycles are introduced

## Common Authorization Patterns

### ✅ Valid hierarchical structure:
```
define viewer: [user]
define editor: [user] or viewer
define admin: [user] or editor
define owner: [user] or admin
```

### ❌ Invalid circular structure:
```
define viewer: editor
define editor: admin  
define admin: viewer  # Creates cycle
```

## Related Errors

- [`relation-no-entry-point`](./relation-no-entry-point.md) - When cycles prevent any entry points
- [`cyclic-error`](./cyclic-error.md) - General cyclic dependency error
- [`undefined-relation`](./undefined-relation.md) - When relations in cycle don't exist

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/cycle_detection.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java semantic validation package

The cycle detection uses depth-first search with visited node tracking to identify circular dependencies in the relation graph.
