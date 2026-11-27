# Invalid Relation on Tupleset

**Error Code:** `invalid-relation-on-tupleset`

**Category:** Structural Validation

## Summary

In a tuple-to-userset such as `member from owner`, the relation `member` does not exist on any type that `owner` allows.

## Description

`member from owner` reads the objects assigned to `owner`, then checks `member` on each of them. So `member` must exist on at least one type that `owner` allows. When it exists on none, the error is reported once for each allowed type.

A tupleset relation that is not defined on the type, as in `member from nope`, is [`invalid-relation-type`](./invalid-relation-type.md).

## Example

```
model
  schema 1.1

type user
type organization
  relations
    define member: [user]
    # Note: 'admin' relation is not defined

type document
  relations
    define owner: [organization]
    define reader: admin from owner  # Error: 'admin' doesn't exist on organization
```

**Error Message:** `` the `admin` relation definition on type `document` is not valid: `admin` does not exist on `owner`, which is of type `organization`. ``

## Resolution

Define the relation on the tupleset's type, or use one that exists.

### Option 1: Define the missing relation

```
model
  schema 1.1

type user
type organization
  relations
    define member: [user]
    define admin: [user]      # Define the missing relation

type document
  relations
    define owner: [organization]
    define reader: admin from owner  # Now valid
```

### Option 2: Use an existing relation

```
model
  schema 1.1

type user
type organization
  relations
    define member: [user]

type document
  relations
    define owner: [organization]
    define reader: member from owner  # Use existing relation
```

## More Examples

### Valid

`member from parent` is valid because `member` exists on `organization`, even though `folder` does not have it:

```
model
  schema 1.1

type user
type folder
type organization
  relations
    define member: [user]
    define admin: [user]

type document
  relations
    define owner: [organization]
    define parent: [folder, organization]
    define reader: member from owner     # Valid: 'member' exists on organization
    define editor: admin from owner      # Valid: 'admin' exists on organization
    define viewer: member from parent    # Valid: 'member' exists on organization, though not on folder
```

### Invalid

```
model
  schema 1.1

type user
type folder
type organization
  relations
    define member: [user]
    # 'admin' relation not defined

type document
  relations
    define owner: [organization]
    define parent: [folder, organization]
    define reader: admin from owner      # Error: 'admin' doesn't exist on organization
    define viewer: admin from parent     # Error: 'admin' exists on neither folder nor organization
```

## Related Errors

- [`tupleuserset-not-direct`](./tupleuserset-not-direct.md) - The tupleset relation is not a plain direct assignment
- [`invalid-relation-type`](./invalid-relation-type.md) - The tupleset relation after `from` is not defined
- [`missing-definition`](./missing-definition.md) - A computed userset names an undefined relation

## Implementation Notes

- Go: `pkg/go/validation/semantic_validation.go`
- JavaScript: `pkg/js/validator/validate-dsl.ts`
- Java: `pkg/java/src/main/java/dev/openfga/language/validation/ModelValidator.java`
