# Reserved Relation Keywords

**Error Code:** `reserved-relation-keywords`

**Category:** Naming Validation

## Summary

Relation names cannot use the reserved keywords `self` or `this`, which have special meaning in OpenFGA's authorization language.

## Description

OpenFGA reserves `self` and `this` because they carry special meaning in relation definitions (`this` refers to a relation's directly-assigned users). Using either as a relation name triggers this error.

These are the only reserved relation names. Operators that appear in the DSL grammar (`or`, `and`, `but not`, `from`, `define`, etc.) are handled by the parser and are not reported through this error.

## Example

The following model would trigger this error:

```
model
  schema 1.1

type user

type document
  relations
    define this: [user]  # Error: a relation cannot be named 'self' or 'this'
```

**Error Message:** `a relation cannot be named 'self' or 'this'.`

## Resolution

Rename the relation to anything other than `self` or `this`:

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]
    define editor: [user] or viewer
    define owner: [user] or editor
```

### Steps to fix:

1. **Locate the offending relation:**
   - The error symbol is the reserved name (`self` or `this`).

2. **Rename it:**
   - Choose a descriptive name that reflects the permission granted.

3. **Update all references:**
   - Update any computed-userset or tuple-to-userset rewrites that referenced the renamed relation.

4. **Re-validate the model.**

## Best Practices for Relation Naming

- **Use authorization verbs:** `can_view`, `can_edit`, `can_delete`
- **Use role-based names:** `viewer`, `editor`, `admin`, `owner`
- **Use business terms:** `subscriber`, `member`, `guest`, `participant`
- **Be descriptive:** Names should clearly indicate the permission granted
- **Stay consistent:** Use similar naming patterns across your model

## Related Errors

- [`reserved-type-keywords`](./reserved-type-keywords.md) - Reserved keywords for type names
- [`self-error`](./self-error.md) - Reserved error code, not currently emitted
- [`invalid-name`](./invalid-name.md) - General invalid naming issues

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/name_validation.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java naming validation package

The validation checks each relation name against the reserved set (`self`, `this`) during the name-validation phase.
