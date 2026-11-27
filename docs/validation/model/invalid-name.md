# Invalid Name

**Error Code:** `invalid-name`

**Category:** Naming Validation

## Summary

A type, relation or condition name does not match its naming rule.

## Description

| Name | Rule | Meaning |
|------|------|---------|
| Type | `^[^:#@\*\s]{1,254}$` | 1 to 254 characters, none of them `:`, `#`, `@`, `*` or whitespace |
| Relation | `^[^:#@\*\s]{1,50}$` | 1 to 50 characters, none of them `:`, `#`, `@`, `*` or whitespace |
| Condition (Go only) | `^[^\*\s]{1,50}$` | 1 to 50 characters, none of them `*` or whitespace |

Uppercase letters, digits, hyphens and underscores are allowed. `self` and `this` are reported as [`reserved-type-keywords`](./reserved-type-keywords.md) or [`reserved-relation-keywords`](./reserved-relation-keywords.md) instead.

The DSL grammar is stricter. A DSL type or relation name uses only ASCII letters, digits, `_`, `-`, `.` and `/`, starts with a letter or `_`, and has each `.` or `/` between two letters, digits or `_` (as in `org.team`). Other names are syntax errors, so in the DSL this error is reported only for length: a type name over 254 characters, or a relation or condition name over 50. In JSON, it is also reported for disallowed characters.

## Example

### Invalid type names

Type names with whitespace or `@`, in a JSON model:

```json
{
  "schema_version": "1.1",
  "type_definitions": [
    { "type": "user" },
    { "type": "team member" },
    { "type": "user@domain" }
  ]
}
```

### Invalid relation names

```
model
  schema 1.1

type user

type document
  relations
    define can_view_documents_shared_with_the_entire_organization: [user]  # Error: 54 characters, the limit is 50
```

In JSON, a relation key such as `can view` or `owner@company` is reported the same way.

**Error Messages:**
- `type 'team member' does not match naming rule: '^[^:#@\*\s]{1,254}$'.`
- `type 'user@domain' does not match naming rule: '^[^:#@\*\s]{1,254}$'.`
- `relation 'can_view_documents_shared_with_the_entire_organization' of type 'document' does not match naming rule: '^[^:#@\*\s]{1,50}$'.`

### Names the DSL rejects as syntax errors

- `type 123type` and `define 1viewer: [user]` (start with a digit)
- `define my viewer: [user]` (contains a space)
- `type user@domain` and `define owner@company: [user]` (contain `@`)

## Resolution

Rename to a name that matches the rule.

### Fix invalid type names

```json
{
  "schema_version": "1.1",
  "type_definitions": [
    { "type": "user" },
    { "type": "team-member" },
    { "type": "user_domain" }
  ]
}
```

### Fix invalid relation names

```
model
  schema 1.1

type user

type document
  relations
    define can_view_shared_documents: [user]  # Shortened to 25 characters
```

## More Examples

### Valid names

```
model
  schema 1.1

type user
type User
type user_profile
type user-profile
type document_123

type document
  relations
    define viewer: [user]
    define can_edit: [user]
    define can-edit: [user]
    define Owner: [user]
    define level_1_access: [user]
```

### Invalid names

| Name | Problem | Reported as |
|------|---------|-------------|
| `team member` | Contains whitespace | `invalid-name` in JSON, syntax error in the DSL |
| `user@domain` | Contains `@` | `invalid-name` in JSON, syntax error in the DSL |
| `org:acme`, `doc#1`, `all*` | Contain `:`, `#` or `*` | `invalid-name` in JSON, syntax error in the DSL |
| A type name of 255 characters | Too long | `invalid-name` |
| A relation name of 51 characters | Too long | `invalid-name` |
| `123document` | Starts with a digit | Syntax error in the DSL, valid in JSON |
| `self`, `this` | Reserved | `reserved-type-keywords` or `reserved-relation-keywords` |

## Related Errors

- [`reserved-type-keywords`](./reserved-type-keywords.md) - A type is named `self` or `this`
- [`reserved-relation-keywords`](./reserved-relation-keywords.md) - A relation is named `self` or `this`

## Implementation Notes

- Go: `pkg/go/validation/name_validation.go`
- JavaScript: `pkg/js/validator/validate-dsl.ts`
- Java: `pkg/java/src/main/java/dev/openfga/language/validation/ModelValidator.java`

Go, JavaScript and Java apply the same type and relation rules. Only Go checks condition names.
