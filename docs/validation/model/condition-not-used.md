# Condition Not Used

**Error Code:** `condition-not-used`

**Category:** Condition Validation

## Summary

A condition is defined, but no type restriction uses it.

## Description

Every condition defined in the model must be used by at least one type restriction, as in `[user with is_weekday]`.

## Example

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user with is_weekday]  # Only is_weekday is used
    define editor: [user]

condition is_weekday(current_time: timestamp) {
  current_time.getDayOfWeek() < 6
}

condition unused_condition(department: string) {  # Error: Never used
  department == "engineering"
}
```

**Error Location:** The `unused_condition` definition.

**Error Message:** `` `unused_condition` condition is not used in the model. ``

## Resolution

Use the condition in a type restriction, or remove it.

### Option 1: Use the condition in a relation

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user with is_weekday]
    define editor: [user with unused_condition]  # Now used
    define admin: viewer and editor  # Requires both conditions

condition is_weekday(current_time: timestamp) {
  current_time.getDayOfWeek() < 6
}

condition unused_condition(department: string) {
  department == "engineering"
}
```

### Option 2: Remove the unused condition

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user with is_weekday]
    define editor: [user]

condition is_weekday(current_time: timestamp) {
  current_time.getDayOfWeek() < 6
}

# unused_condition removed
```

## Related Errors

- [`condition-not-defined`](./condition-not-defined.md) - A type restriction uses a condition that is not defined
- [`different-nested-condition-name`](./different-nested-condition-name.md) - A JSON condition's `name` differs from its key

## Implementation Notes

- Go: `pkg/go/validation/condition_validation.go`
- JavaScript: `pkg/js/validator/validate-dsl.ts`
- Java: `pkg/java/src/main/java/dev/openfga/language/validation/ModelValidator.java`
