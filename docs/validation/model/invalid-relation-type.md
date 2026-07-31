# Invalid Relation Type

**Error Code:** `invalid-relation-type`

**Category:** Semantic Validation

## Summary

A relation reference specifies an invalid or incorrectly formatted relation type, preventing proper authorization evaluation.

## Description

This error occurs when relation definitions reference relations in an invalid format or context. Common issues include:
- Malformed relation references in type restrictions
- Invalid syntax in computed userset operations
- Incorrect relation specification in tuple-to-userset operations
- Missing or improperly formatted relation qualifiers

OpenFGA requires relation references to follow specific syntax rules to ensure proper parsing and evaluation during authorization checks.

## Example

The following models would trigger this error:

### Invalid relation reference syntax:
```
model
  schema 1.1

type user
type organization
  relations
    define member: [user]

type document
  relations
    define viewer: [organization#]        # Error: empty relation after #
    define editor: [organization##member] # Error: double ##
    define admin: [organization#member#]  # Error: trailing #
```

### Invalid relation in computed userset:
```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]
    define editor: invalid.relation       # Error: invalid relation format
    define admin: viewer.                 # Error: trailing dot
```

**Error Message:** `Invalid relation type format in reference: 'organization#'`

## Resolution

Correct the relation reference syntax:

### Fix relation reference format:
```
model
  schema 1.1

type user
type organization
  relations
    define member: [user]

type document
  relations
    define viewer: [organization#member]  # Correct format: type#relation
    define editor: [user, organization#member]
    define admin: [user] or viewer
```

### Fix computed userset references:
```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]
    define editor: viewer                 # Correct computed userset
    define admin: [user] or editor
```

### Steps to fix:

1. **Identify the invalid relation reference:**
   - Check the error message for the specific invalid syntax
   - Locate the problematic relation reference in your model

2. **Understand the correct syntax:**
   - Type restrictions: `[type#relation]` or `[type]`
   - Computed usersets: `relation_name`
   - Tuple-to-userset: `relation from tupleset_relation`

3. **Correct the syntax:**
   - Remove extra characters (trailing #, dots, etc.)
   - Ensure proper type#relation format for cross-type references
   - Use simple relation names for computed usersets

4. **Test the model:**
   - Validate the corrected model
   - Ensure authorization logic works as expected

## Valid Relation Reference Patterns

### ✅ Correct syntax:
```
type organization
  relations
    define member: [user]
    define admin: [user]

type document
  relations
    # Type restrictions
    define viewer: [user]                    # Direct type reference
    define editor: [user, organization#member] # Cross-type relation reference
    
    # Computed usersets
    define collaborator: viewer              # Simple relation reference
    define manager: editor                   # Another relation reference
    
    # Complex operations
    define contributor: viewer or editor     # Union of relations
    
    # Tuple-to-userset
    define reader: member from owner         # Relation traversal
    define owner: [organization#admin]       # Owner assignment
```

### ❌ Incorrect syntax:
```
# Malformed cross-type references
define viewer: [organization#]          # Missing relation
define editor: [organization##member]   # Double ##
define admin: [#member]                 # Missing type

# Invalid computed usersets
define collaborator: viewer.            # Trailing dot
define manager: .editor                 # Leading dot
define contributor: viewer..editor      # Double dots

# Invalid operations
define reader: [user]#member           # Wrong placement of #
define owner: organization#            # Incomplete reference
```

## Related Errors

- [`undefined-relation`](./undefined-relation.md) - When referenced relations don't exist
- [`invalid-type`](./invalid-type.md) - When type references are invalid
- [`missing-definition`](./missing-definition.md) - When definitions are missing

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/semantic_validation.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java semantic validation package

The validation uses parsing rules to check relation reference syntax and ensures all references follow OpenFGA's specification.
