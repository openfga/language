# Schema Version Unsupported

**Error Code:** `schema-version-unsupported`

**Category:** Schema Validation

## Summary

The model declares schema version `1.0`, which is no longer supported.

## Description

`1.1` and `1.2` are the supported versions. `1.0` reports this error, and any other version reports [`invalid-schema`](./invalid-schema.md).

A model written in 1.0 syntax, such as `define viewer as self`, fails to parse with a syntax error before this check. This error is reported for a DSL model that declares `schema 1.0` but is otherwise written in 1.1 syntax, and for a JSON model with `"schema_version": "1.0"`.

## Example

```
model
  schema 1.0  # Error: Version 1.0 is no longer supported

type user

type document
  relations
    define viewer: [user]
```

**Error Location:** The version, `1.0`.

**Error Message:** `schema version no longer supported`

## Resolution

Declare version `1.1`, or `1.2` if the model uses modules:

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]
```

### Rewriting a 1.0 model

- Change `schema 1.0` to `schema 1.1`
- Replace `as` with `:`
- Replace `self` with the types that can be assigned to the relation, such as `[user, group#member, user:*]`

This 1.0 model:

```
model
  schema 1.0

type user

type group
  relations
    define member as self

type folder
  relations
    define parent as self
    define viewer as self or viewer from parent

type document
  relations
    define parent as self
    define viewer as self
    define can_read as viewer or viewer from parent
```

is written in 1.1 as:

```
model
  schema 1.1

type user

type group
  relations
    define member: [user]

type folder
  relations
    define parent: [folder]
    define viewer: [user] or viewer from parent

type document
  relations
    define parent: [folder]
    define viewer: [user]
    define can_read: viewer or viewer from parent
```

Version 1.1 also changes which tuples are valid. For example, every user needs a type (`user:bob`, not `bob`), and a wildcard is written `user:*`, not `*`. The full 1.0 to 1.1 migration guide is in the [openfga.dev repository history](https://github.com/openfga/openfga.dev/blob/da151c064105eea5dddaf0594c4e499106140689/docs/content/modeling/migrating/migrating-schema-1-1.mdx).

## Related Errors

- [`schema-version-required`](./schema-version-required.md) - No schema version is declared
- [`invalid-schema`](./invalid-schema.md) - The model declares a version other than `1.0`, `1.1` or `1.2`

## Implementation Notes

- Go: `pkg/go/validation/schema_validation.go`
- JavaScript: `pkg/js/validator/validate-dsl.ts`
- Java: `pkg/java/src/main/java/dev/openfga/language/validation/ModelValidator.java`
