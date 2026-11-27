# Invalid Type

**Error Code:** `invalid-type`

**Category:** Semantic Validation

## Summary

A type restriction names a type that is not defined in the model.

## Description

Every type in a type restriction, whether alone (`[team]`), as a wildcard (`[team:*]`) or with a relation (`[team#member]`), must be defined with a `type` declaration. Type names are case-sensitive, so `[User]` does not match `type user`. For `[team#member]`, JavaScript and Java also report [`invalid-relation-type`](./invalid-relation-type.md), while Go reports only `invalid-type`.

## Example

### Restriction names an undefined type

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user, team]  # Error: 'team' is not defined
```

### Restriction uses a different case or spelling

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [User]  # Error: only 'user' is defined
```

**Error Messages:**
- `` `team` is not a valid type. ``
- `` `User` is not a valid type. ``

## Resolution

Define the type, or correct the reference. In a modular model, also check that the file that declares the type is listed in `fga.mod`.

### Define the missing type

```
model
  schema 1.1

type user

type team

type document
  relations
    define viewer: [user, team]  # 'team' is now defined
```

### Fix the type reference

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]  # Matches the declared 'user' type
```

## Related Errors

- [`invalid-relation-type`](./invalid-relation-type.md) - The type exists, but the relation in `type#relation` does not
- [`invalid-name`](./invalid-name.md) - A type name does not match the naming rule
- [`reserved-type-keywords`](./reserved-type-keywords.md) - A type is named `self` or `this`

## Implementation Notes

- Go: `pkg/go/validation/semantic_validation.go`
- JavaScript: `pkg/js/validator/validate-dsl.ts`
- Java: `pkg/java/src/main/java/dev/openfga/language/validation/ModelValidator.java`
