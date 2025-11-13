# Reserved Type Keywords

**Error Code:** `reserved-type-keywords`

**Category:** Naming Validation

## Summary

Type names cannot use reserved keywords that have special meaning in OpenFGA's authorization language and system operations.

## Description

OpenFGA reserves certain keywords for internal operations, language constructs, and future feature expansion. Using these reserved words as type names can cause:
- Parsing conflicts and ambiguous references
- Runtime errors in authorization evaluation
- Incompatibility with future OpenFGA versions
- Unexpected behavior in authorization checks

Reserved type keywords include system identifiers, language constructs, and terms that have special meaning in the OpenFGA ecosystem.

## Example

The following models would trigger this error:

```
model
  schema 1.1

type model  # Error: 'model' is a reserved keyword
  relations
    define viewer: [user]

type schema  # Error: 'schema' is a reserved keyword
  relations
    define member: [user]

type relation  # Error: 'relation' is a reserved keyword
  relations
    define owner: [user]
```

**Error Message:** `Type name 'model' is a reserved keyword and cannot be used`

## Resolution

Choose different type names that don't conflict with reserved keywords:

### Use descriptive, domain-specific names:

```
model
  schema 1.1

type user

type data_model  # Instead of 'model'
  relations
    define viewer: [user]

type schema_definition  # Instead of 'schema'
  relations
    define member: [user]

type business_relation  # Instead of 'relation'
  relations
    define owner: [user]
```

### Steps to fix:

1. **Identify the reserved keyword:**
   - Check the error message for the specific reserved word
   - Note which type definition is using the reserved keyword

2. **Choose appropriate replacement:**
   - Select descriptive names that reflect the type's purpose
   - Avoid other reserved keywords and system terms
   - Use clear, domain-specific terminology

3. **Update all references:**
   - Change the type definition
   - Update any references to the renamed type in relations
   - Verify no broken references remain

4. **Test the model:**
   - Validate the updated model
   - Ensure authorization logic still works correctly

## Common Reserved Keywords

| Reserved Word | Alternative Names |
|---------------|-------------------|
| `model` | `data_model`, `authorization_model`, `business_model` |
| `schema` | `schema_definition`, `model_version`, `structure` |
| `relation` | `business_relation`, `relationship`, `association` |
| `type` | `entity_type`, `object_type`, `business_type` |
| `define` | `definition`, `specification`, `rule` |
| `condition` | `business_condition`, `rule_condition`, `constraint` |

## TODO: Complete Reserved Keywords List

<!-- TODO: Add comprehensive list of:
- All reserved keywords across OpenFGA versions
- System-level reserved terms
- Future reserved keywords for compatibility
- Language construct keywords
- Built-in function names that should be avoided
-->

## Best Practices

- **Use domain-specific names:** Choose names that reflect your business domain
- **Avoid system terms:** Stay away from technical OpenFGA terminology
- **Be descriptive:** Use clear, meaningful names that explain the type's purpose
- **Check documentation:** Verify names against OpenFGA keyword lists
- **Test thoroughly:** Validate models after renaming to ensure functionality

## Related Errors

- [`reserved-relation-keywords`](./reserved-relation-keywords.md) - Reserved keywords for relation names
- [`self-error`](./self-error.md) - Specific reserved words 'self' and 'this'
- [`invalid-name`](./invalid-name.md) - General invalid naming issues

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/name_validation.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java naming validation package

The validation maintains a list of reserved keywords and checks all type names against this list during the parsing phase.
