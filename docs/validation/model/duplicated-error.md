# Duplicated Error

**Error Code:** `duplicated-error`

**Category:** Structural Validation

## Summary

Duplicate definitions of types or relations within the same scope are not allowed as they create ambiguity in the authorization model.

## Description

OpenFGA requires unique names within their respective scopes:
- **Type names** must be unique across the entire model
- **Relation names** must be unique within each type
- **Type restrictions** within a relation must not be duplicated

Duplicate definitions create ambiguity about which definition should be used and can lead to inconsistent authorization behavior.

## Example

The following models would trigger this error:

### Duplicate type names:
```
model
  schema 1.1

type user
  relations
    define viewer: [organization#member]

type user  # Error: Duplicate type name
  relations
    define admin: [organization#member]
```

### Duplicate relation names:

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]
    define editor: [user]
    define viewer: [organization#member]  # Error: Duplicate relation name
```

> [!NOTE]
> This only occurs in DSL validation, not in JSON - as in JSON the relations are represented as a map, which inherently prevents duplicates.

### Duplicate type restrictions:
```
model
  schema 1.1

type user
type organization
  relations
    define member: [user]

type document
  relations
    define viewer: [user, organization#member, user]  # Error: Duplicate 'user' restriction
```

**Error Message:** `Duplicate definition found: type 'user'` or `Duplicate relation 'viewer' in type 'document'`

## Resolution

Remove or rename the duplicate definitions:

### Fix duplicate types:
```
model
  schema 1.1

type user
  relations
    define viewer: [organization#member]

type admin_user  # Rename to make unique
  relations
    define admin: [organization#member]
```

### Fix duplicate relations:
```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user, organization#member]  # Combine into single definition
    define editor: [user]
```

### Fix duplicate type restrictions:
```
model
  schema 1.1

type user
type organization
  relations
    define member: [user]

type document
  relations
    define viewer: [user, organization#member]  # Remove duplicate 'user'
```

### Steps to fix:

1. **Identify the duplicate:**
   - Check error message for specific duplicate name and location
   - Locate both definitions in your model

2. **Determine intended behavior:**
   - Decide if definitions should be merged or renamed
   - Consider if this was a copy-paste error

3. **Choose resolution approach:**
   - **Merge**: Combine definitions if they serve the same purpose
   - **Rename**: Give unique names if they serve different purposes
   - **Remove**: Delete if one is unnecessary

4. **Update references:**
   - If renaming, update all references to use the new name
   - Verify no broken references remain

## Common Causes

- **Copy-paste errors:** Duplicating code blocks without renaming
- **Refactoring mistakes:** Moving definitions without removing originals
- **Merge conflicts:** Git merges creating duplicate definitions
- **Template usage:** Using templates without customizing names

## Related Errors

- [`invalid-name`](./invalid-name.md) - When names don't follow proper format
- [`reserved-type-keywords`](./reserved-type-keywords.md) - When using reserved words
- [`undefined-relation`](./undefined-relation.md) - Broken references after renaming

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/duplicate_detection.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java duplicate detection package

The validation builds maps of all defined names and checks for conflicts during the parsing phase.
