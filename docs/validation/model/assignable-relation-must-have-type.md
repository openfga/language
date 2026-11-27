# Assignable Relation Must Have Type

**Error Code:** `assignable-relation-must-have-type`

**Category:** Structural Validation

## Summary

A relation in a JSON model is a direct assignment, but lists no types that can be assigned to it.

## Description

In JSON, a direct assignment is written as `{ "this": {} }`, and the types that can be assigned are listed in `metadata.relations.<relation>.directly_related_user_types`. When that list is empty or missing, JavaScript and Java report `assignable-relation-must-have-type`.

A DSL model cannot have this error: a direct assignment in the DSL lists its types, as in `[user]`, and `[]` is a syntax error.

Go does not report this code. For the model below, Go reports [`relation-no-entry-point`](./relation-no-entry-point.md) instead: `` `viewer` is an impossible relation for `document` (no entrypoint). ``

## Example

This JSON model lists no types for the `viewer` relation:

```json
{
  "schema_version": "1.1",
  "type_definitions": [
    { "type": "user" },
    {
      "type": "document",
      "relations": { "viewer": { "this": {} } },
      "metadata": { "relations": { "viewer": { "directly_related_user_types": [] } } }
    }
  ]
}
```

**Error Message:** `assignable relation '' must have types`

The same error is reported when `viewer` has no entry in `metadata.relations`.

## Resolution

List the types that can be assigned to the relation. Each entry is a type (`{ "type": "user" }`), a wildcard (`{ "type": "user", "wildcard": {} }`) or a userset (`{ "type": "organization", "relation": "member" }`).

### Option 1: Add the types to the relation metadata

```json
{
  "schema_version": "1.1",
  "type_definitions": [
    { "type": "user" },
    {
      "type": "document",
      "relations": { "viewer": { "this": {} } },
      "metadata": { "relations": { "viewer": { "directly_related_user_types": [ { "type": "user" } ] } } }
    }
  ]
}
```

### Option 2: Write the model in the DSL

This DSL model is the same as the JSON in Option 1:

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]
```

## Related Errors

- [`relation-no-entry-point`](./relation-no-entry-point.md) - What Go reports for the same model

## Implementation Notes

- Go: not reported
- JavaScript: `pkg/js/validator/validate-dsl.ts`
- Java: `pkg/java/src/main/java/dev/openfga/language/validation/ModelValidator.java`
