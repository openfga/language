# OpenFGA Validation Error Troubleshooting Guide

Quick fixes for each validation error. Each row links to the error's page.

## Quick Reference

| Error | Quick Fix | Link |
|-------|-----------|------|
| `schema-version-required` | Add `"schema_version": "1.1"` to the JSON model | [Details](./schema-version-required.md) |
| `schema-version-unsupported` | Use version `1.1` or `1.2` | [Details](./schema-version-unsupported.md) |
| `invalid-schema` | Use version `1.1` or `1.2` | [Details](./invalid-schema.md) |
| `reserved-type-keywords` | Rename the type | [Details](./reserved-type-keywords.md) |
| `reserved-relation-keywords` | Rename the relation | [Details](./reserved-relation-keywords.md) |
| `invalid-name` | Shorten the name or remove `:`, `#`, `@`, `*` and whitespace | [Details](./invalid-name.md) |
| `duplicated-error` | Remove or merge the duplicates | [Details](./duplicated-error.md) |
| `missing-definition` | Define the relation, or fix the reference | [Details](./missing-definition.md) |
| `invalid-relation-type` | Reference a relation that exists on that type | [Details](./invalid-relation-type.md) |
| `invalid-type` | Define the type, or fix the reference | [Details](./invalid-type.md) |
| `relation-no-entry-point` | Add a direct assignment, such as `[user]` | [Details](./relation-no-entry-point.md) |
| `invalid-relation-on-tupleset` | Define the relation on the tupleset's type, or use one that exists | [Details](./invalid-relation-on-tupleset.md) |
| `tupleuserset-not-direct` | Make the tupleset relation a plain list of types, such as `[folder]` | [Details](./tupleuserset-not-direct.md) |
| `assignable-relation-must-have-type` | List the relation's types in `directly_related_user_types` | [Details](./assignable-relation-must-have-type.md) |
| `type-wildcard-relation` | Use a wildcard or a relation, not both | [Details](./type-wildcard-relation.md) |
| `condition-not-defined` | Define the condition, or fix the reference | [Details](./condition-not-defined.md) |
| `condition-not-used` | Use the condition, or remove it | [Details](./condition-not-used.md) |
| `different-nested-condition-name` | Make the condition's `name` equal its key | [Details](./different-nested-condition-name.md) |
| `multiple-modules-in-file` | Use one module per file in the model's metadata | [Details](./multiple-modules-in-file.md) |

## Validation Checklist

Before deploying a model, check that:

- [ ] The schema version is specified (`schema 1.1`, or `1.2` for modules)
- [ ] Type names are at most 254 characters, with no `:`, `#`, `@`, `*` or whitespace
- [ ] Relation names are at most 50 characters, with no `:`, `#`, `@`, `*` or whitespace
- [ ] No type or relation is named `self` or `this`
- [ ] All referenced types exist
- [ ] All referenced relations exist
- [ ] Every relation has an entry point
- [ ] Every condition that is used is defined, and every defined condition is used
- [ ] Each module file declares one module

## Using the FGA CLI

Get the CLI from https://github.com/openfga/cli.

1. [Validate a model](https://github.com/openfga/cli/#validate-an-authorization-model):
    ```bash
    fga model validate --file model.fga
    ```

2. [Run tests against a model](https://github.com/openfga/cli/#run-tests-on-an-authorization-model):
    ```bash
    fga model test --tests model.fga.yaml
    ```

## Getting Help

- **Examples**: the [validation test cases](../../../tests/data) in this repository, and the OpenFGA [sample stores](https://github.com/openfga/sample-stores)
- **Community**: [OpenFGA community](https://openfga.dev/community)
