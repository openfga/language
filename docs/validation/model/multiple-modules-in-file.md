# Multiple Modules in File

**Error Code:** `multiple-modules-in-file`

**Category:** Multi-file Validation

## Summary

A single file contains definitions for multiple modules, violating OpenFGA's module organization requirements where each file should contain only one module.

## Description

OpenFGA's module system requires clean separation between modules to maintain:
- Clear module boundaries and dependencies
- Predictable import/export behavior
- Maintainable code organization
- Consistent module resolution

When multiple modules are defined within a single file, it creates ambiguity about module ownership, makes dependency tracking difficult, and can lead to unexpected behavior during module resolution.

## Example

The following file structure would trigger this error:

**single-file.fga:**
```
module user_management
  schema 1.1

type user
  relations
    define profile_owner: [user]

module document_management  # Error: Second module in same file
  schema 1.1

type document
  relations
    define owner: [user_management.user]
```

**Error Message:** `Multiple modules detected in file 'single-file.fga': user_management, document_management`

## Resolution

Split the modules into separate files:

### Option 1: Create separate files for each module

**user-management.fga:**
```
module user_management
  schema 1.1

type user
  relations
    define profile_owner: [user]
```

**document-management.fga:**
```
module document_management
  schema 1.1

import user_management

type document
  relations
    define owner: [user_management.user]
```

### Option 2: Combine into single module (if appropriate)

**combined.fga:**
```
module application
  schema 1.1

type user
  relations
    define profile_owner: [user]

type document
  relations
    define owner: [user]
```

### Steps to fix:

1. **Identify all modules in the file:**
   - Review the error message for list of modules
   - Locate all `module` declarations in the file

2. **Plan module organization:**
   - Decide if modules should remain separate or be combined
   - Consider module dependencies and relationships

3. **Create separate files:**
   - Create one file per module with descriptive names
   - Move each module's content to its own file
   - Update any cross-module references

4. **Update imports:**
   - Add necessary import statements for cross-module references
   - Verify all module dependencies are properly declared

## TODO: Module Organization Best Practices

<!-- TODO: Add guidance for:
- Best practices for module naming and organization
- How to handle complex module dependencies
- Strategies for refactoring large single-file models into modules
- Performance implications of module structure
- Module versioning and compatibility considerations
-->

## Common Causes

- **Copy-paste errors:** Copying entire modules without creating separate files
- **Refactoring mistakes:** Merging files without removing module declarations
- **Template usage:** Using module templates without proper file separation
- **Migration issues:** Converting single-module models to multi-module incorrectly

## Related Errors

- [`module-split-across-files`](./module-split-across-files.md) - When single module spans multiple files
- [`cross-module-reference`](./cross-module-reference.md) - Invalid references between modules
- [`invalid-syntax`](./invalid-syntax.md) - Syntax issues in module declarations

## Implementation Notes

This validation is enforced consistently across:
- Go implementation: `pkg/go/validation/multi_file_validation.go`
- JavaScript implementation: `pkg/js/validator/validate-dsl.ts`
- Java implementation: Java multi-file validation package

The validation tracks module declarations across all files and reports conflicts when multiple modules are found in a single file.
