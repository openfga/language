# OpenFGA Model Validation Errors

Each page in this directory covers one validation error code: what triggers it, an example, and how to fix it.

## Complete Error Reference

| Error Code | Error Type | Summary | Documentation |
|------------|------------|---------|---------------|
| `schema-version-required` | Schema | Schema version must be specified | [schema-version-required.md](./schema-version-required.md) |
| `schema-version-unsupported` | Schema | Schema version `1.0`, which is no longer supported | [schema-version-unsupported.md](./schema-version-unsupported.md) |
| `invalid-schema` | Schema | Schema version other than `1.0`, `1.1` or `1.2` | [invalid-schema.md](./invalid-schema.md) |
| `reserved-type-keywords` | Naming | Type name uses reserved keyword | [reserved-type-keywords.md](./reserved-type-keywords.md) |
| `reserved-relation-keywords` | Naming | Relation name uses reserved keyword | [reserved-relation-keywords.md](./reserved-relation-keywords.md) |
| `invalid-name` | Naming | Type, relation or condition name does not match its naming rule | [invalid-name.md](./invalid-name.md) |
| `duplicated-error` | Structure | Duplicate type, type restriction or rewrite operand | [duplicated-error.md](./duplicated-error.md) |
| `missing-definition` | Semantic | Computed userset references an undefined relation | [missing-definition.md](./missing-definition.md) |
| `invalid-relation-type` | Semantic | Relation in `type#relation` or tupleset relation not defined | [invalid-relation-type.md](./invalid-relation-type.md) |
| `invalid-type` | Semantic | Type restriction references an undefined type | [invalid-type.md](./invalid-type.md) |
| `relation-no-entry-point` | Semantic | Relation does not lead to a direct assignment | [relation-no-entry-point.md](./relation-no-entry-point.md) |
| `invalid-relation-on-tupleset` | Structure | Tuple-to-userset relation missing on every tupleset type | [invalid-relation-on-tupleset.md](./invalid-relation-on-tupleset.md) |
| `tupleuserset-not-direct` | Structure | Tupleset relation must be a plain direct assignment of types | [tupleuserset-not-direct.md](./tupleuserset-not-direct.md) |
| `assignable-relation-must-have-type` | Structure | Direct assignment lists no types (JSON; JS and Java only) | [assignable-relation-must-have-type.md](./assignable-relation-must-have-type.md) |
| `type-wildcard-relation` | Wildcard | Type restriction cannot have wildcard and relation (JSON; JS and Java only) | [type-wildcard-relation.md](./type-wildcard-relation.md) |
| `condition-not-defined` | Condition | Referenced condition is not defined | [condition-not-defined.md](./condition-not-defined.md) |
| `condition-not-used` | Condition | Defined condition is never used | [condition-not-used.md](./condition-not-used.md) |
| `different-nested-condition-name` | Condition | Condition `name` differs from its key (JSON) | [different-nested-condition-name.md](./different-nested-condition-name.md) |
| `multiple-modules-in-file` | Multi-file | Model metadata assigns several modules to one file | [multiple-modules-in-file.md](./multiple-modules-in-file.md) |

## Implementation Notes

- **Go**: `pkg/go/validation/`
- **JavaScript**: `pkg/js/validator/`
- **Java**: `pkg/java/src/main/java/dev/openfga/language/validation/`

Go does not report `assignable-relation-must-have-type` or `type-wildcard-relation`.

## Contributing

When adding a validation rule:

1. Add the error code to all implementations (Go, JS, Java)
2. Add a page for it, following the existing pages
3. Add the error to this index table
4. Add test cases to the shared YAML files in `tests/data`, which all three implementations run

## Support

For questions about validation errors, or to report a difference between implementations, file an issue in the [OpenFGA Language repository](https://github.com/openfga/language).
