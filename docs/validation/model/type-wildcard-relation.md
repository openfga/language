# Type Wildcard Relation

**Error Code:** `type-wildcard-relation`

**Category:** Wildcard Validation

## Summary

A type restriction combines a wildcard and a relation, as in `organization:*#member`.

## Description

A type restriction is a type (`user`), a wildcard (`user:*`) or a userset (`organization#member`), never a wildcard and a relation together.

In the DSL, `[organization:*#member]` and `[organization#admin:*]` are syntax errors, so a DSL model cannot have this error. JavaScript and Java report it for a JSON restriction whose `type` is `organization:*` and that also sets `relation`, or whose `type` is `organization#member` and that also sets `wildcard`.

Go does not report this code. For the model below, Go reports `invalid-type`, because `organization:*` is not a defined type.

A restriction that sets both the `relation` and `wildcard` fields is not reported with this code either. JavaScript and Java keep the wildcard and ignore the relation, and Go rejects the JSON because the two fields cannot both be set.

## Example

```json
{
  "schema_version": "1.1",
  "type_definitions": [
    { "type": "user" },
    {
      "type": "organization",
      "relations": { "member": { "this": {} } },
      "metadata": { "relations": { "member": { "directly_related_user_types": [ { "type": "user" } ] } } }
    },
    {
      "type": "document",
      "relations": { "viewer": { "this": {} } },
      "metadata": {
        "relations": {
          "viewer": { "directly_related_user_types": [ { "type": "organization:*", "relation": "member" } ] }
        }
      }
    }
  ]
}
```

**Error Message:** `` type restriction `organization:*#member` cannot contain both wildcard and relation ``

A restriction of `{ "type": "organization#member", "wildcard": {} }` is reported the same way, with `organization#member:*` in the message.

## Resolution

Use either a wildcard or a relation, in separate restrictions if you need both. In JSON, keep `type` a plain type name and set either `wildcard` or `relation`, for example `{ "type": "user", "wildcard": {} }` or `{ "type": "organization", "relation": "member" }`.

### Option 1: Use a wildcard or a relation

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
    define viewer: [user:*]              # Allows granting all users access to a particular document (once the tuple is written)
    define editor: [organization#member] # Allows granting members of an organization access to a particular document (once the tuple is written)
```

### Option 2: Use both, in separate restrictions

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
    define viewer: [user, organization#member]        # Direct users or org members
    define editor: [user:*, organization#admin] # All users or org admins
```

## Related Errors

- [`invalid-type`](./invalid-type.md) - What Go reports for the same model
- [`assignable-relation-must-have-type`](./assignable-relation-must-have-type.md) - A JSON direct assignment lists no types

## Implementation Notes

- Go: not reported
- JavaScript: `pkg/js/validator/validate-dsl.ts`
- Java: `pkg/java/src/main/java/dev/openfga/language/validation/ModelValidator.java`
