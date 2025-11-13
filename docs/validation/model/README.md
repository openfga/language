# OpenFGA Model Validation Errors

This directory contains comprehensive documentation for all OpenFGA model validation errors. These error codes and messages are consistent across Go, JavaScript, and Java implementations to provide a unified validation experience.

## Overview

OpenFGA model validation ensures that authorization models are syntactically correct, semantically valid, and follow best practices. When validation fails, specific error codes and messages help identify and resolve issues.

## Error Categories

- **Schema Validation**: Issues with model schema versions and structure
- **Name Validation**: Problems with type and relation naming
- **Semantic Validation**: Logical issues like undefined references and cycles
- **Structural Validation**: Problems with model structure and relationships
- **Condition Validation**: Issues with condition definitions and usage
- **Multi-file Validation**: Problems with module consistency across files

## Complete Error Reference

| Error Code | Error Type | Summary | Documentation |
|------------|------------|---------|---------------|
| `schema-version-required` | Schema | Schema version must be specified | [schema-version-required.md](./schema-version-required.md) |
| `schema-version-unsupported` | Schema | Unsupported schema version | [schema-version-unsupported.md](./schema-version-unsupported.md) |
| `invalid-schema-version` | Schema | Invalid schema version format | [invalid-schema-version.md](./invalid-schema-version.md) |
| `reserved-type-keywords` | Naming | Type name uses reserved keyword | [reserved-type-keywords.md](./reserved-type-keywords.md) |
| `reserved-relation-keywords` | Naming | Relation name uses reserved keyword | [reserved-relation-keywords.md](./reserved-relation-keywords.md) |
| `self-error` | Naming | Invalid use of 'self' or 'this' | [self-error.md](./self-error.md) |
| `invalid-name` | Naming | Invalid type or relation name format | [invalid-name.md](./invalid-name.md) |
| `duplicated-error` | Structure | Duplicate type or relation definition | [duplicated-error.md](./duplicated-error.md) |
| `missing-definition` | Semantic | Referenced type or relation not defined | [missing-definition.md](./missing-definition.md) |
| `undefined-type` | Semantic | Type is referenced but not defined | [undefined-type.md](./undefined-type.md) |
| `undefined-relation` | Semantic | Relation is referenced but not defined | [undefined-relation.md](./undefined-relation.md) |
| `invalid-relation-type` | Semantic | Invalid relation type in reference | [invalid-relation-type.md](./invalid-relation-type.md) |
| `invalid-type` | Semantic | Invalid type in relation definition | [invalid-type.md](./invalid-type.md) |
| `relation-no-entry-point` | Semantic | Relation has no entry point for assignment | [relation-no-entry-point.md](./relation-no-entry-point.md) |
| `cyclic-error` | Semantic | Circular dependency in relations | [cyclic-error.md](./cyclic-error.md) |
| `cyclic-relation` | Semantic | Circular relation dependency detected | [cyclic-relation.md](./cyclic-relation.md) |
| `invalid-relation-on-tupleset` | Structure | Invalid relation in tuple-to-userset | [invalid-relation-on-tupleset.md](./invalid-relation-on-tupleset.md) |
| `tupleuserset-not-direct` | Structure | Tuple-to-userset must have direct assignment | [tupleuserset-not-direct.md](./tupleuserset-not-direct.md) |
| `invalid-wildcard-error` | Wildcard | Invalid wildcard usage in relation | [invalid-wildcard-error.md](./invalid-wildcard-error.md) |
| `assignable-relation-must-have-type` | Wildcard | Assignable relation must specify type | [assignable-relation-must-have-type.md](./assignable-relation-must-have-type.md) |
| `type-wildcard-relation` | Wildcard | Type restriction cannot have wildcard and relation | [type-wildcard-relation.md](./type-wildcard-relation.md) |
| `condition-not-defined` | Condition | Referenced condition is not defined | [condition-not-defined.md](./condition-not-defined.md) |
| `condition-not-used` | Condition | Defined condition is never used | [condition-not-used.md](./condition-not-used.md) |
| `different-nested-condition-name` | Condition | Condition name mismatch in nested structure | [different-nested-condition-name.md](./different-nested-condition-name.md) |
| `multiple-modules-in-file` | Multi-file | Multiple modules detected in single file | [multiple-modules-in-file.md](./multiple-modules-in-file.md) |
| `module-split-across-files` | Multi-file | Module definition split across multiple files | [module-split-across-files.md](./module-split-across-files.md) |
| `cross-module-reference` | Multi-file | Reference crosses module boundaries | [cross-module-reference.md](./cross-module-reference.md) |
| `invalid-schema` | Schema | Invalid schema structure | [invalid-schema.md](./invalid-schema.md) |
| `invalid-syntax` | Syntax | Invalid DSL syntax | [invalid-syntax.md](./invalid-syntax.md) |

## Usage

Each error documentation includes:

- **Error Code**: The unique identifier for the error
- **Summary**: Brief description of what causes the error
- **Description**: Detailed explanation of the validation rule
- **Example**: Code example that would trigger this error
- **Resolution**: How to fix the error (step-by-step guidance)

## Implementation Notes

These validation errors are implemented consistently across:

- **Go**: `pkg/go/validation/` package
- **JavaScript**: `pkg/js/validator/` package  
- **Java**: Java validation implementation

Error codes, messages, and validation logic are synchronized to ensure identical behavior across all language implementations.

## Contributing

When adding new validation rules:

1. Add the error code to all implementations (Go, JS, Java)
2. Create documentation following the template format
3. Add the error to this index table
4. Include test cases in the appropriate test suites
5. Update the YAML test files for cross-language validation

## Support

For questions about validation errors or to report inconsistencies between language implementations, please file an issue in the [OpenFGA Language repository](https://github.com/openfga/language).
