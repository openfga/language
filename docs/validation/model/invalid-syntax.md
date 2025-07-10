# Invalid Syntax

**Error Code:** `invalid-syntax`

**Category:** Syntax Validation

## Summary

The DSL contains syntax errors that prevent proper parsing of the authorization model.

## Description

This error occurs when the OpenFGA DSL contains syntax that doesn't conform to the language specification. Syntax errors prevent the parser from understanding the model structure and must be fixed before semantic validation can occur.

Common syntax issues include:
- Incorrect indentation or formatting
- Missing or extra punctuation
- Malformed keywords or identifiers
- Invalid DSL structure or organization
- Unsupported language constructs

## Example

The following models would trigger this error:

### Incorrect indentation:
```
model
schema 1.1    # Error: missing indentation

type user
relations     # Error: missing indentation
define viewer: [user]  # Error: missing indentation
```

### Malformed keywords:
```
model
  schema 1.1

typ user      # Error: 'typ' instead of 'type'

type document
  relation    # Error: 'relation' instead of 'relations'
    defin viewer: [user]  # Error: 'defin' instead of 'define'
```

### Invalid structure:
```
model
  schema 1.1

type user {   # Error: invalid brace syntax
  relations
    define viewer: [user]
}
```

**Error Message:** `Syntax error at line 3: expected proper indentation`

## Resolution

Fix the syntax errors to conform to OpenFGA DSL specification:

### Fix indentation:
```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]
```

### Fix keywords:
```
model
  schema 1.1

type user    # Corrected from 'typ'

type document
  relations  # Corrected from 'relation'
    define viewer: [user]  # Corrected from 'defin'
```

### Fix structure:
```
model
  schema 1.1

type user

type document  # Removed invalid braces
  relations
    define viewer: [user]
```

### Steps to fix:

1. **Identify the syntax error:**
   - Check the error message for line number and specific issue
   - Review the problematic line and surrounding context

2. **Follow DSL specification:**
   - Use proper indentation (2 spaces per level)
   - Use correct keywords (`model`, `type`, `relations`, `define`)
   - Follow OpenFGA DSL structure rules

3. **Validate syntax incrementally:**
   - Fix one syntax error at a time
   - Validate after each fix to catch additional issues

4. **Test the corrected model:**
   - Ensure the model parses successfully
   - Proceed to semantic validation

## OpenFGA DSL Structure Reference

### ✅ Correct DSL syntax:
```
model
  schema 1.1

condition condition_name(param: Type) {
  param.field == "value"
}

type type_name
  relations
    define relation_name: [type] or other_relation
    define complex_relation: [type#relation, other_type:*]
```

### ❌ Invalid DSL syntax:
```
model {               # Wrong: no braces
  schema: 1.1         # Wrong: colon instead of space
}

typ user              # Wrong: 'typ' instead of 'type'
relation              # Wrong: 'relation' instead of 'relations'
define viewer [user]  # Wrong: missing colon
```

## Related Errors

- [`invalid-schema`](./invalid-schema.md) - When schema structure is invalid
- [`invalid-name`](./invalid-name.md) - When names don't follow format rules
- [`schema-version-required`](./schema-version-required.md) - When schema declaration is missing

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/` during parsing phase
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts` 
- Java implementation: Java DSL parser

The validation occurs during the initial parsing phase before semantic analysis begins, ensuring the model structure is valid before checking logical consistency.
