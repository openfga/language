# Reserved Relation Keywords

**Error Code:** `reserved-relation-keywords`

**Category:** Naming Validation

## Summary

Relation names cannot use reserved keywords that have special meaning in OpenFGA's authorization language and system operations.

## Description

OpenFGA reserves certain keywords for internal operations, language constructs, and future feature expansion. Using these reserved words as relation names can cause:
- Parsing conflicts during DSL processing
- Ambiguous references in authorization evaluation
- Runtime errors during permission checks
- Incompatibility with future OpenFGA versions

Reserved relation keywords include system identifiers, built-in operations, and terms that have special meaning in OpenFGA's relation resolution logic.

## Example

The following models would trigger this error:

```
model
  schema 1.1

type user

type document
  relations
    define define: [user]     # Error: 'define' is a reserved keyword
    define relation: [user]   # Error: 'relation' is a reserved keyword
    define union: [user]      # Error: 'union' is a reserved keyword
    define from: [user]       # Error: 'from' is a reserved keyword
```

**Error Message:** `Relation name 'define' is a reserved keyword and cannot be used`

## Resolution

Choose different relation names that don't conflict with reserved keywords:

### Use descriptive, business-oriented names:

```
model
  schema 1.1

type user

type document
  relations
    define definition: [user]      # Instead of 'define'
    define association: [user]     # Instead of 'relation'
    define combined_access: [user] # Instead of 'union'
    define source_reference: [user] # Instead of 'from'
```

### Better domain-specific alternatives:

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]
    define editor: [user] or viewer
    define owner: [user] or editor
    define admin: [user] or owner
```

### Steps to fix:

1. **Identify the reserved keyword:**
   - Check the error message for the specific reserved word
   - Note which relation definition is using the reserved keyword

2. **Choose appropriate replacement:**
   - Select names that reflect the relation's authorization purpose
   - Use clear, business-domain terminology
   - Avoid other reserved keywords and system terms

3. **Update all references:**
   - Change the relation definition
   - Update any computed userset references to the renamed relation
   - Update tuple-to-userset operations that reference the relation
   - Verify no broken references remain

4. **Test the model:**
   - Validate the updated model
   - Ensure authorization logic still works correctly

## Common Reserved Keywords

| Reserved Word | Alternative Names | Purpose |
|---------------|-------------------|---------|
| `define` | `definition`, `specification`, `rule` | Business rule definitions |
| `relation` | `association`, `connection`, `link` | Business relationships |
| `union` | `combined`, `merged`, `aggregate` | Combined permissions |
| `intersection` | `shared`, `common`, `overlap` | Shared permissions |
| `difference` | `exclusive`, `except`, `minus` | Exclusive permissions |
| `from` | `via`, `through`, `source` | Relation traversal |
| `this` | `direct`, `assigned`, `explicit` | Direct assignment |
| `schema` | `version`, `structure`, `format` | Model structure |

## TODO: Complete Reserved Keywords List

<!-- TODO: Add comprehensive list of:
- All reserved relation keywords across OpenFGA versions
- Built-in operation keywords that cannot be used as relation names
- Future reserved keywords for compatibility
- Context-specific reserved terms
- Language construct keywords that affect relation parsing
-->

## Best Practices for Relation Naming

- **Use authorization verbs:** `can_view`, `can_edit`, `can_delete`
- **Use role-based names:** `viewer`, `editor`, `admin`, `owner`
- **Use business terms:** `subscriber`, `member`, `guest`, `participant`
- **Be descriptive:** Names should clearly indicate the permission granted
- **Stay consistent:** Use similar naming patterns across your model

## Related Errors

- [`reserved-type-keywords`](./reserved-type-keywords.md) - Reserved keywords for type names
- [`self-error`](./self-error.md) - Specific reserved words 'self' and 'this'
- [`invalid-name`](./invalid-name.md) - General invalid naming issues

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/name_validation.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java naming validation package

The validation maintains a list of reserved keywords and checks all relation names against this list during the parsing phase.
