# Missing Definition

**Error Code:** `missing-definition`

**Category:** Semantic Validation

## Summary

A relation's rewrite refers to a relation that is not defined on the same type.

## Description

This error is reported when a computed userset names a relation the type does not have, on its own (`define editor: owner`) or as an operand of `or`, `and` or `but not`.

Other undefined references have their own codes:
- An undefined type in a type restriction, such as `[organization]` with no `type organization`, is [`invalid-type`](./invalid-type.md).
- An undefined condition, such as `[user with missing_condition]`, is [`condition-not-defined`](./condition-not-defined.md).
- An undefined relation in a `type#relation` restriction, or an undefined tupleset relation after `from`, is [`invalid-relation-type`](./invalid-relation-type.md).
- A relation before `from` that is missing on every type the tupleset relation allows is [`invalid-relation-on-tupleset`](./invalid-relation-on-tupleset.md).

## Example

### Missing relation definition

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]
    define editor: [user] or owner  # Error: relation 'owner' not defined
```

### Using `this` for direct assignment

In schema 1.1, `this` is not direct assignment. It is read as a relation named `this`:

```
model
  schema 1.1

type user

type document
  relations
    define viewer: this  # Error: relation 'this' not defined
```

**Error Messages:**
- `` the relation `owner` does not exist. ``
- `` the relation `this` does not exist. ``

## Resolution

Define the relation, or correct the reference. For direct assignment, use a type restriction such as `[user]`.

### Option 1: Define the missing relation

```
model
  schema 1.1

type user

type document
  relations
    define owner: [user]  # Define the missing relation
    define viewer: [user]
    define editor: [user] or owner
```

### Option 2: Correct or remove the reference

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]              # Direct assignment instead of 'this'
    define editor: [user] or viewer    # Reference existing relation
```

## Related Errors

- [`invalid-type`](./invalid-type.md) - An undefined type in a type restriction
- [`condition-not-defined`](./condition-not-defined.md) - An undefined condition
- [`invalid-relation-type`](./invalid-relation-type.md) - An undefined relation in a `type#relation` restriction, or an undefined tupleset relation
- [`invalid-relation-on-tupleset`](./invalid-relation-on-tupleset.md) - An undefined relation before `from`

## Implementation Notes

- Go: `pkg/go/validation/semantic_validation.go`
- JavaScript: `pkg/js/validator/validate-dsl.ts`
- Java: `pkg/java/src/main/java/dev/openfga/language/validation/ModelValidator.java`
