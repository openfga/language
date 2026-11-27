# Relation No Entry Point

**Error Code:** `relation-no-entry-point`

**Category:** Semantic Validation

## Summary

No user can ever have the relation, because it does not lead to a direct assignment.

## Description

A relation has an entry point when it leads to a direct assignment such as `[user]`. A computed userset, a tuple-to-userset or a `type#relation` restriction leads to one when the relation it refers to has one. An `or` needs one operand with an entry point, an `and` needs every operand to have one, and a `but not` needs both sides to have one.

In schema 1.1, direct assignment is written as a type restriction such as `[user]`. `define admin: this` is not direct assignment: it is read as a relation named `this` and reports [`missing-definition`](./missing-definition.md).

## Example

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

**Error Message:** `` `viewer` is an impossible relation for `document` (potential loop). ``

`editor` and `admin` report the same error. The message ends in `(potential loop).` when the check finds a cycle, and in `(no entrypoint).` otherwise.

## Resolution

Make at least one relation in the chain a direct assignment.

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

### Option 2: Remove the cycle

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user] or editor
    define editor: [user] or admin
    define admin: [user]  # The chain now ends in a direct assignment
```

## Related Errors

- [`missing-definition`](./missing-definition.md) - A computed userset names an undefined relation, including `this`

## Implementation Notes

- Go: `pkg/go/validation/cycle_detection.go`
- JavaScript: `pkg/js/validator/validate-dsl.ts`
- Java: `pkg/java/src/main/java/dev/openfga/language/validation/ModelValidator.java`
