# Reserved Relation Keywords

**Error Code:** `reserved-relation-keywords`

**Category:** Naming Validation

## Summary

A relation is named `self` or `this`, which are reserved.

## Description

`self` and `this` are the only reserved names. In JSON, `this` marks a direct assignment (`{ "this": {} }`), and the schema 1.0 DSL used `self` for the same purpose.

Other keywords are not reported with this code. `model`, `schema`, `type`, `relation`, `module` and `extend` are valid relation names, and names such as `define`, `relations`, `condition`, `and`, `or`, `from` and `with` are syntax errors.

## Example

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

Rename the relation, and update the rewrites that use it:

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

## Related Errors

- [`reserved-type-keywords`](./reserved-type-keywords.md) - A type is named `self` or `this`
- [`invalid-name`](./invalid-name.md) - A name does not match the naming rule

## Implementation Notes

- Go: `pkg/go/validation/name_validation.go`
- JavaScript: `pkg/js/validator/validate-dsl.ts`
- Java: `pkg/java/src/main/java/dev/openfga/language/validation/ModelValidator.java`
