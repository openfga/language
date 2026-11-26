# Duplicated Error

**Error Code:** `duplicated-error`

**Category:** Structural Validation

## Summary

A model declares a type twice, repeats a type restriction in a relation, or repeats an operand in a relation's rewrite.

## Description

The validator reports `duplicated-error` for:
- **Type names** declared more than once
- **Type restrictions** repeated in one relation, such as `[user, user]`
- **Rewrite operands** repeated at the top level of a union or intersection, such as `editor or editor`, or a difference that subtracts an operand from itself, such as `editor but not editor`

A relation defined twice in one type is a syntax error, not this error (see [Duplicate relation names](#duplicate-relation-names)).

Duplicates are checked only when the earlier checks find no errors. A model with a duplicate and an undefined type reports `invalid-type`, and reports the duplicate only after that is fixed.

## Example

### Duplicate type names

```
model
  schema 1.1

type user

type organization
  relations
    define member: [user]

type organization  # Error: Duplicate type name
  relations
    define admin: [user]
```

### Duplicate type restrictions

```
model
  schema 1.1

type user
type organization
  relations
    define member: [user]

type document
  relations
    define viewer: [user, organization#member, user]  # Error: Duplicate 'user' restriction
```

### Duplicate rewrite operands

```
model
  schema 1.1

type user

type document
  relations
    define editor: [user]
    define viewer: editor or editor  # Error: Duplicate 'editor' operand
```

**Error Messages:** (one per example above)
- `` the type `organization` is a duplicate. ``
- `` the type restriction `user` is a duplicate in the relation `viewer`. ``
- `` the partial relation definition `editor` is a duplicate in the relation `viewer`. ``

### Duplicate relation names

Defining the same relation twice in one type is a syntax error:

```
model
  schema 1.1

type user
type organization
  relations
    define member: [user]

type document
  relations
    define viewer: [user]
    define editor: [user]
    define viewer: [organization#member]  # Syntax error: 'viewer' is already defined in 'document'
```

> [!NOTE]
> In JSON, relations are a map, so a relation name cannot repeat within a type.

## Resolution

Remove or merge the duplicates.

### Fix duplicate types

```
model
  schema 1.1

type user

type organization  # Merged into a single declaration
  relations
    define member: [user]
    define admin: [user]
```

### Fix duplicate type restrictions

```
model
  schema 1.1

type user
type organization
  relations
    define member: [user]

type document
  relations
    define viewer: [user, organization#member]  # Remove duplicate 'user'
```

### Fix duplicate rewrite operands

```
model
  schema 1.1

type user

type document
  relations
    define editor: [user]
    define viewer: editor  # Repeated operand removed
```

For a relation defined twice, merge the two definitions into one, for example `define viewer: [user, organization#member]`.

## Related Errors

- [`invalid-type`](./invalid-type.md) - Reported instead of this error when a restriction also names an undefined type

## Implementation Notes

- Go: `pkg/go/validation/duplicate_detection.go`
- JavaScript: `pkg/js/validator/validate-dsl.ts`
- Java: `pkg/java/src/main/java/dev/openfga/language/validation/ModelValidator.java`
