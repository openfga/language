# Condition Not Defined

**Error Code:** `condition-not-defined`

**Category:** Condition Validation

## Summary

A type restriction uses a condition that the model does not define.

## Description

Every condition named in a type restriction, as in `[user with is_weekday]`, must be defined with a `condition` block.

## Example

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user with undefined_condition]  # Error: condition not defined
    define editor: [user with is_weekday]          # Valid: condition exists

condition is_weekday(current_time: timestamp) {
  current_time.getDayOfWeek() < 6
}
```

**Error Location:** `undefined_condition` in the `viewer` relation.

**Error Message:** `` `undefined_condition` is not a defined condition in the model. ``

## Resolution

Define the condition, correct the reference, or remove it.

### Option 1: Define the missing condition

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user with undefined_condition]  # Now valid
    define editor: [user with is_weekday]

condition is_weekday(current_time: timestamp) {
  current_time.getDayOfWeek() < 6
}

condition undefined_condition(department: string) {
  department == "engineering"
}
```

### Option 2: Correct the reference

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user with is_weekday]  # Reference existing condition
    define editor: [user with is_weekday]

condition is_weekday(current_time: timestamp) {
  current_time.getDayOfWeek() < 6
}
```

### Option 3: Remove the condition reference

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]  # Remove condition reference
    define editor: [user with is_weekday]

condition is_weekday(current_time: timestamp) {
  current_time.getDayOfWeek() < 6
}
```

## Related Errors

- [`condition-not-used`](./condition-not-used.md) - A condition is defined but never used
- [`different-nested-condition-name`](./different-nested-condition-name.md) - A JSON condition's `name` differs from its key

## Implementation Notes

- Go: `pkg/go/validation/condition_validation.go`
- JavaScript: `pkg/js/validator/validate-dsl.ts`
- Java: `pkg/java/src/main/java/dev/openfga/language/validation/ModelValidator.java`
