# Tuple-to-Userset Not Direct

**Error Code:** `tupleuserset-not-direct`

**Category:** Structural Validation

## Summary

The tupleset relation of a tuple-to-userset is not a plain direct assignment of types.

## Description

A tuple-to-userset such as `member from owner` reads the objects assigned to `owner`, then checks `member` on each of them. So `owner` must be a direct assignment of types only, such as `define owner: [folder, organization]`. The validator rejects a tupleset relation that:
- has a rewrite, such as `[organization] or admin`
- is an alias of another relation, such as `define parent: owner`
- allows a wildcard, such as `[organization:*]`
- allows a userset, such as `[organization#member]`

## Example

```
model
  schema 1.1

type user

type organization
  relations
    define member: [user]

type folder
  relations
    define admin: [user]
    define owner: [organization] or admin  # Not a plain direct assignment: has a rewrite
    define reader: member from owner  # Error: 'owner' is used in tuple-to-userset

type photo
  relations
    define owner: [organization#member]  # Not a plain direct assignment: allows a userset
    define reader: member from owner  # Error: 'owner' is used in tuple-to-userset

type resource
  relations
    define owner: [organization:*]  # Not a plain direct assignment: allows a wildcard
    define reader: member from owner  # Error: 'owner' is used in tuple-to-userset

type document
  relations
    define owner: [organization]
    define admin: owner  # Not a plain direct assignment: an alias of 'owner'
    define reader: member from admin  # Error: 'admin' is used in tuple-to-userset
```

**Error Location:** The relation after `from` in each `reader` definition.

**Error Message:** `` `owner` relation used inside from allows only direct relation. ``

The `folder`, `photo` and `resource` types each report this message. The `document` type reports it for `admin`.

## Resolution

Make the tupleset relation a direct assignment of types:
- Replace a wildcard (`[organization:*]`) or userset (`[organization#member]`) with the type (`[organization]`)
- Move a rewrite (`or`, `and`, `but not`) out of the tupleset relation and into the relation that uses it
- Replace an alias with the relation it points to, for example `member from owner` instead of `member from admin` when `define admin: owner`

### Option 1: Use a direct assignment of types

```
model
  schema 1.1

type user
type organization
  relations
    define member: [user]

type document
  relations
    define owner: [organization]      # Only direct assignment with no indirection
    define reader: member from owner  # Now valid
```

### Option 2: Move the extra logic out of the tupleset relation

```
model
  schema 1.1

type user

type organization
  relations
    define member: [user]

type folder
  relations
    define admin: [user]
    define owner: [organization]
    define reader: member from owner or admin  # 'or admin' moved out of 'owner'
```

## Related Errors

- [`invalid-relation-type`](./invalid-relation-type.md) - The tupleset relation after `from` is not defined
- [`invalid-relation-on-tupleset`](./invalid-relation-on-tupleset.md) - The relation before `from` is missing on the tupleset's types

## Implementation Notes

- Go: `pkg/go/validation/semantic_validation.go`
- JavaScript: `pkg/js/validator/validate-dsl.ts`
- Java: `pkg/java/src/main/java/dev/openfga/language/validation/ModelValidator.java`
