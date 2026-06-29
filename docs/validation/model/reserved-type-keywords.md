# Reserved Type Keywords

**Error Code:** `reserved-type-keywords`

**Category:** Naming Validation

## Summary

Type names cannot use the reserved keywords `self` or `this`, which have special meaning in OpenFGA's authorization language.

## Description

OpenFGA reserves `self` and `this` because they carry special meaning in relation definitions. Using either as a type name causes this error:

- **`this`**: refers to a relation's directly-assigned users.
- **`self`**: reserved alongside `this`.

These are the only reserved type names. Other identifiers that look like language constructs (`model`, `schema`, `type`, `define`, `relation`, etc.) are handled by the grammar and are not reported through this error.

## Example

The following model would trigger this error:

```
model
  schema 1.1

type self  # Error: a type cannot be named 'self' or 'this'
  relations
    define viewer: [user]
```

**Error Message:** `a type cannot be named 'self' or 'this'.`

## Resolution

Rename the type to anything other than `self` or `this`:

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]
```

### Steps to fix:

1. **Locate the offending type:**
   - The error symbol is the reserved name (`self` or `this`).

2. **Rename it:**
   - Choose a descriptive name that reflects the type's purpose.

3. **Update all references:**
   - Update any type restrictions or relations that referenced the renamed type.

4. **Re-validate the model.**

## Related Errors

- [`reserved-relation-keywords`](./reserved-relation-keywords.md) - Reserved keywords for relation names
- [`self-error`](./self-error.md) - Reserved error code, not currently emitted
- [`invalid-name`](./invalid-name.md) - General invalid naming issues

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/name_validation.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java naming validation package

The validation checks each type name against the reserved set (`self`, `this`) during the name-validation phase.
