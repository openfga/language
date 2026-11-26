# Different Nested Condition Name

**Error Code:** `different-nested-condition-name`

**Category:** Condition Validation

## Summary

In a JSON model, a condition's `name` field is missing or differs from its key in `conditions`.

> [!NOTE]
> Only JSON models can have this error. In the DSL, a condition's name is written once.

## Description

In JSON, conditions are a map from key to condition:

```jsonc
{
  "conditions": {
    "<condition_key>": {
      "name": "<condition_internal_name>",
      "expression": "...",
      "parameters": { ... }
    }
  }
}
```

The validator compares each key with its condition's `name` field and reports `different-nested-condition-name` when they differ, including when `name` is missing.

## Example

The condition key is `low_access_count`, but its `name` is `small_access_count`:

```json
{
  "schema_version": "1.1",
  "type_definitions": [
    { "type": "user" },
    {
      "type": "document",
      "relations": { "viewer": { "this": {} } },
      "metadata": {
        "relations": {
          "viewer": { "directly_related_user_types": [ { "type": "user", "condition": "low_access_count" } ] }
        }
      }
    }
  ],
  "conditions": {
    "low_access_count": {
      "name": "small_access_count",
      "expression": "access_count < 10",
      "parameters": { "access_count": { "type_name": "TYPE_NAME_INT" } }
    }
  }
}
```

**Error Message:** `` condition key is `low_access_count` but nested name property is small_access_count ``

If `name` is missing, the error is still reported, and the nested name in the message is empty (Go), `undefined` (JavaScript) or `null` (Java).

## Resolution

Make the key and `name` the same. Removing `name` does not fix the error.

### Option 1: Set `name` to the key

```json
{
  "schema_version": "1.1",
  "type_definitions": [
    { "type": "user" },
    {
      "type": "document",
      "relations": { "viewer": { "this": {} } },
      "metadata": {
        "relations": {
          "viewer": { "directly_related_user_types": [ { "type": "user", "condition": "low_access_count" } ] }
        }
      }
    }
  ],
  "conditions": {
    "low_access_count": {
      "name": "low_access_count",
      "expression": "access_count < 10",
      "parameters": { "access_count": { "type_name": "TYPE_NAME_INT" } }
    }
  }
}
```

### Option 2: Rename the key to match `name`

```json
{
  "schema_version": "1.1",
  "type_definitions": [
    { "type": "user" },
    {
      "type": "document",
      "relations": { "viewer": { "this": {} } },
      "metadata": {
        "relations": {
          "viewer": { "directly_related_user_types": [ { "type": "user", "condition": "small_access_count" } ] }
        }
      }
    }
  ],
  "conditions": {
    "small_access_count": {
      "name": "small_access_count",
      "expression": "access_count < 10",
      "parameters": { "access_count": { "type_name": "TYPE_NAME_INT" } }
    }
  }
}
```

If you rename the key, update every type restriction that uses the condition, as the `viewer` restriction above does.

## Related Errors

- [`condition-not-defined`](./condition-not-defined.md) - A type restriction uses a condition that is not defined
- [`condition-not-used`](./condition-not-used.md) - A condition is defined but never used

## Implementation Notes

- Go: `pkg/go/validation/condition_validation.go`
- JavaScript: `pkg/js/validator/validate-dsl.ts`
- Java: `pkg/java/src/main/java/dev/openfga/language/validation/ModelValidator.java`
