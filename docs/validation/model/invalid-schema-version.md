# Invalid Schema Version

**Error Code:** `invalid-schema-version`

**Category:** Schema Validation

## Summary

An unrecognised schema version is reported under [`invalid-schema`](./invalid-schema.md).

## Description

This code is declared but never emitted. A version other than `1.1` and `1.2` is reported as [`invalid-schema`](./invalid-schema.md), with the version as the symbol, so `schema 0.9` gives the message `invalid schema 0.9`. That page also covers version `1.0`, a missing version, and a `schema` line that does not parse.

## Resolution

Declare `1.1`, or `1.2` for a model split across files with `module` declarations. See [`invalid-schema`](./invalid-schema.md) for the full example and the compatibility matrix in [`schema-version-unsupported`](./schema-version-unsupported.md) for what each version supports.

## Related Errors

- [`invalid-schema`](./invalid-schema.md) - An unrecognised schema version
- [`schema-version-required`](./schema-version-required.md) - No schema version is declared
- [`schema-version-unsupported`](./schema-version-unsupported.md) - Version `1.0`, recognised but retired
