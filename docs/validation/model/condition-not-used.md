# Condition Not Used

**Error Code:** `condition-not-used`

**Category:** Condition Validation

## Summary

A condition is defined in the model but is never referenced or used in any relation, creating unnecessary code that should be removed or utilized.

## Description

This error occurs when a condition is defined but never referenced in any relation definitions. Unused conditions:
- Add unnecessary complexity to the model
- Can indicate incomplete implementation or refactoring artifacts
- May confuse developers about the model's intended behavior
- Consume resources during model processing

OpenFGA validates that all defined conditions serve a purpose in the authorization model to maintain clean, efficient authorization logic.

## Example

The following model would trigger this error:

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

condition unused_condition(user: User) {  # Error: Never used
  user.department == "engineering"
}
```

**Error Location:** The `unused_condition` definition that is never referenced.

**Error Message:** `Condition 'unused_condition' is defined but not used in any relation`

## Resolution

Either use the condition in a relation or remove it from the model:

### Option 1: Use the condition in a relation

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user with is_weekday]
    define editor: [user with unused_condition]  # Now used
    define admin: [user with unused_condition and is_weekday]  # Combined usage

condition is_weekday(current_time: timestamp) {
  current_time.getDayOfWeek() < 6
}

condition unused_condition(user: User) {
  user.department == "engineering"
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

### Steps to fix:

1. **Identify the unused condition:**
   - Check the error message for the specific condition name
   - Locate the condition definition in your model

2. **Determine intended usage:**
   - Review if the condition was meant to be used somewhere
   - Check if this is leftover from refactoring or incomplete implementation

3. **Choose resolution approach:**
   - Use the condition in appropriate relations if it serves a purpose
   - Remove the condition if it's no longer needed
   - Document why the condition exists if it's for future use

4. **Test the authorization logic:**
   - Ensure removing/using the condition doesn't break intended behavior
   - Verify conditional authorization works as expected

## Common Causes

- **Refactoring artifacts:** Conditions left over after removing relations
- **Copy-paste errors:** Copying conditions from other models without usage
- **Incomplete implementation:** Conditions defined but relations not yet updated
- **Development process:** Conditions created during development but never utilized

## Related Errors

- [`condition-not-defined`](./condition-not-defined.md) - When conditions are referenced but not defined
- [`different-nested-condition-name`](./different-nested-condition-name.md) - Condition name mismatches
- [`missing-definition`](./missing-definition.md) - General missing definition patterns

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/condition_validation.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java condition validation package

The validation builds maps of defined conditions and used conditions, then identifies any conditions that exist in the defined set but not in the used set.
