# Different Nested Condition Name

**Error Code:** `different-nested-condition-name`

**Category:** Condition Validation

## Summary

A JSON authorization model declares a condition whose map key (the condition id) does not match the `name` field inside the condition object, or contains an internally nested condition object whose `name` field differs from the key under which it appears. This indicates an inconsistency that can break condition reference resolution.

> [!NOTE]
> This error only applies to the JSON authorization model format. The DSL does **not** embed a secondary `name` property for conditions, so this mismatch cannot occur when using the DSL.

## Description

In the JSON format, conditions are supplied as an object map:

```json5
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

The validator expects the `<condition_key>` and the `name` field (if provided) to be identical. A mismatch can introduce ambiguity about the canonical condition identifier used by relations or other tooling.

You will receive `different-nested-condition-name` when:
- A top‑level condition entry's key and its `name` field differ.

This validation does not concern typos in relation references (see [`condition-not-defined`](./condition-not-defined.md)) or unused conditions ([`condition-not-used`](./condition-not-used.md)). It is strictly about internal condition name consistency inside the JSON structure.

## Examples

### Mismatched top-level condition key vs internal name
```json5
{
  "schema_version": "1.1",
  "type_definitions": [
    { "type": "user", "relations": {}, "metadata": null },
    { "type": "document", "relations": { "viewer": { "this": {} } }, "metadata": { "relations": { "viewer": { "directly_related_user_types": [ { "type": "user", "condition": "low_access_count" } ] } } } }
  ],
  "conditions": {
    "low_access_count": {
      "name": "small_access_count", // Mismatch
      "expression": "access_count < 10",
      "parameters": { "access_count": { "type_name": "TYPE_NAME_INT" } }
    }
  }
}
```
**Error Message:** `Condition name mismatch: expected 'low_access_count' but found 'small_access_count'`

## Resolution

Ensure the condition key and the internal `name` (when present) are identical. You have a few options:

### Option 1: Align the internal `name` with the key
```json5
"conditions": {
  "low_access_count": {
    "name": "low_access_count",
    "expression": "access_count < 10",
    "parameters": { "access_count": { "type_name": "TYPE_NAME_INT" } }
  }
}
```

### Option 2: Rename the condition key to match the intended `name`
```json5
"conditions": {
  "small_access_count": {  // Key updated
    "name": "small_access_count",
    "expression": "access_count < 10",
    "parameters": { "access_count": { "type_name": "TYPE_NAME_INT" } }
  }
}
```
Be sure to update all relation references now using `small_access_count`.

## Steps to Fix

1. Identify mismatch:
   - Read the error message for: expected <key> vs found <internal-name>.
   - Locate the offending condition entry in the JSON.
2. Decide canonical name:
   - Pick either the key or the internal `name` as the authoritative identifier.
3. Normalize:
   - Make them identical, or remove `name` to rely solely on the key.
4. Update references:
   - Adjust any relation metadata or condition usages to the canonical name if you changed the key.
5. Re‑validate:
   - Re-run the validator to ensure the error disappears and no new `condition-not-defined` issues were introduced.

## Common Condition Naming Issues (JSON Format)

### ✅ Consistent
```json5
"conditions": {
  "business_access": {
    "name": "business_access",
    "expression": "is_weekday && is_business_hours",
    "parameters": {}
  }
}
```

### ❌ Mismatch
```json5
"conditions": {
  "business_access": {
    "name": "biz_access", // Mismatch
    "expression": "is_weekday && is_business_hours",
    "parameters": {}
  }
}
```

## Related Errors

- [`condition-not-defined`](./condition-not-defined.md) - Referenced conditions that don't exist
- [`condition-not-used`](./condition-not-used.md) - Conditions defined but never referenced
- [`invalid-name`](./invalid-name.md) - General naming format violations

## Implementation Notes

Current behavior:
- The validator compares the condition map key to the `name` field (when present) in JSON models.
- The DSL representation does not produce this error because there is no distinct internal `name` field separate from the declared identifier.

Code references:
- Go: `pkg/go/validation/condition_validation.go` (consistency logic)
- JavaScript: `pkg/js/util/exceptions.ts` and usage in validators
- Java: Corresponding condition validation package (JSON path)

Internally the validator raises `different-nested-condition-name` to prevent ambiguity about which identifier should be authoritative.
