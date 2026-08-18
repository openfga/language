# Invalid Schema Version

**Error Code:** `invalid-schema-version`

**Category:** Schema Validation

## Summary

An unrecognised schema version is reported under [`invalid-schema`](./invalid-schema.md).

## Description

`1.1` and `1.2` are the versions OpenFGA accepts. A version outside that set is reported as `invalid-schema`, with the version as the symbol, so `schema 0.9` gives the message `invalid schema 0.9`.

Two neighbouring conditions have codes of their own. Version `1.0` is recognised and retired, so it reports [`schema-version-unsupported`](./schema-version-unsupported.md), and a model carrying no version at all reports [`schema-version-required`](./schema-version-required.md).

A `schema` line that does not parse as a version at all, such as `schema v1.1`, `schema 1.1.0`, or `schema` on its own, fails during DSL transformation and returns a syntax error with no error code attached.

## Resolution

Declare `1.1`, or `1.2` for a model split across files with `module` declarations. See [`invalid-schema`](./invalid-schema.md) for the full example and the compatibility matrix in [`schema-version-unsupported`](./schema-version-unsupported.md) for what each version supports.

## Related Errors

- [`invalid-schema`](./invalid-schema.md) - An unrecognised schema version
- [`schema-version-required`](./schema-version-required.md) - No schema version is declared
- [`schema-version-unsupported`](./schema-version-unsupported.md) - Version `1.0`, recognised but retired
