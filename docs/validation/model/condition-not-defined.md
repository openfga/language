# Condition Not Defined

**Error Code:** `condition-not-defined`

**Category:** Condition Validation

## Summary

A relation references a condition that is not defined in the model, creating a broken reference that would cause runtime errors.

## Description

This error occurs when a relation definition includes a condition reference that doesn't exist in the model's condition definitions. OpenFGA requires all condition references to be valid and resolvable at validation time.

Conditions are used to add dynamic evaluation logic to authorization checks, such as time-based access, IP restrictions, or custom business logic. When a condition is referenced but not defined, the authorization system cannot evaluate the conditional logic.

## Example

The following model would trigger this error:

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

**Error Location:** Line 7, `undefined_condition` in the viewer relation definition.

**Error Message:** `Condition 'undefined_condition' is not defined (referenced in relation 'viewer' of type 'document')`

## Resolution

Define the missing condition or correct the reference to an existing condition:

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

condition undefined_condition(user: User) {
  user.department == "engineering"
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

### Steps to fix:

1. **Identify the undefined condition:**
   - Check the error message for the exact condition name
   - Note which relation and type contains the invalid reference

2. **Verify intended behavior:**
   - Determine if the condition should exist or if it's a typo
   - Review your conditional authorization requirements

3. **Choose resolution approach:**
   - Define the missing condition if it should exist
   - Correct the reference if it's a typo
   - Remove the reference if conditional logic isn't needed

4. **Test the condition:**
   - Validate the model after making changes
   - Test the conditional logic in your application

## Common Causes

- **Typos in condition names:** `is_weekday` vs `is_week_day` vs `is_wekday`
- **Case sensitivity:** `IsWeekday` instead of `is_weekday`
- **Missing condition definitions:** Referencing planned but unimplemented conditions
- **Refactoring errors:** Renaming conditions without updating references

## Related Errors

- [`condition-not-used`](./condition-not-used.md) - When conditions are defined but never used
- [`different-nested-condition-name`](./different-nested-condition-name.md) - Condition name mismatches
- [`undefined-relation`](./undefined-relation.md) - Similar pattern for relation references

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/condition_validation.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java condition validation package

The validation builds a map of all defined conditions and cross-references them with condition usage in relation definitions.
