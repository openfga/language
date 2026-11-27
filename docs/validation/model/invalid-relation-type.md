# Invalid Relation Type

**Error Code:** `invalid-relation-type`

**Category:** Semantic Validation

## Summary

A relation reference names a relation that does not exist on the type it points to.

## Description

This error is reported for:
- **`type#relation` restrictions**: in `[organization#member]`, `organization` is defined but has no relation `member`.
- **Tupleset relations**: in `owner from parent`, `parent` is not defined on the current type.

Similar problems have other codes:
- An undefined type in a restriction, as in `[team]`, is [`invalid-type`](./invalid-type.md). For `[team#member]` with no `type team`, Go reports only `invalid-type`, while JavaScript and Java report `invalid-type` and `invalid-relation-type`.
- A computed userset that names an undefined relation, as in `define editor: nonexistent`, is [`missing-definition`](./missing-definition.md).
- A relation before `from` that is missing on every type the tupleset relation allows, as in `admin from owner`, is [`invalid-relation-on-tupleset`](./invalid-relation-on-tupleset.md).

## Example

### Undefined relation in a type restriction

```
model
  schema 1.1

type user
type organization
  relations
    define member: [user]

type document
  relations
    define viewer: [organization#admin]  # Error: 'organization' has no 'admin' relation
```

### Undefined tupleset relation

```
model
  schema 1.1

type user

type document
  relations
    define owner: [user]
    define viewer: owner from parent  # Error: 'document' has no 'parent' relation
```

**Error Messages:**
- `` `admin` is not a valid relation for `organization`. ``
- `` `parent` is not a valid relation for `document`. ``

## Resolution

Reference a relation that exists, or define it.

### Fix the type#relation reference

```
model
  schema 1.1

type user
type organization
  relations
    define member: [user]
    define admin: [user]

type document
  relations
    define viewer: [organization#admin]  # 'admin' is now defined on organization
    define editor: [user, organization#member]
```

### Fix the tupleset relation

```
model
  schema 1.1

type user

type document
  relations
    define owner: [user]
    define parent: [document]
    define viewer: owner from parent  # 'parent' is now defined on document
```

## Related Errors

- [`invalid-type`](./invalid-type.md) - The type in a restriction is not defined
- [`missing-definition`](./missing-definition.md) - A computed userset names an undefined relation
- [`invalid-relation-on-tupleset`](./invalid-relation-on-tupleset.md) - The relation before `from` is missing on the tupleset's types

## Implementation Notes

- Go: `pkg/go/validation/semantic_validation.go`
- JavaScript: `pkg/js/validator/validate-dsl.ts`
- Java: `pkg/java/src/main/java/dev/openfga/language/validation/ModelValidator.java`
