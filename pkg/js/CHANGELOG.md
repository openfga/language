# Changelog

## v0.2.0-beta.11

### [v0.2.0-beta.11](https://github.com/openfga/language/releases/tag/vv0.2.0-beta.10...v0.2.0-beta.11) (2024-02-22)

Fixed:
- Allow transforming JSON where `this` is not the first element (#167)

## v0.2.0-beta.10

### [v0.2.0-beta.10](https://github.com/openfga/language/releases/tag/vv0.2.0-beta.9...v0.2.0-beta.10) (2024-01-23)

Fixed:
- Re-allow `list` and `map` to be accepted as identifiers and relation names (#134)

## v0.2.0-beta.9

### [v0.2.0-beta.9](https://github.com/openfga/language/releases/tag/vv0.2.0-beta.8...v0.2.0-beta.9) (2024-01-08)

Fixed:
- Support models with comments (#131)

## v0.2.0-beta.8

### [v0.2.0-beta.8](https://github.com/openfga/language/releases/tag/vv0.2.0-beta.7...v0.2.0-beta.8) (2024-01-03)

Fixed:
- Fixed exported types (#129)

## v0.2.0-beta.7

### [v0.2.0-beta.7](https://github.com/openfga/language/releases/tag/vv0.2.0-beta.6...v0.2.0-beta.7) (2023-12-13)

Fixed:
- Issue in validating entrypoint or loop with some models (#120)

## v0.2.0-beta.6

### [v0.2.0-beta.6](https://github.com/openfga/language/releases/tag/vv0.2.0-beta.5...v0.2.0-beta.6) (2023-12-11)

last commit: 8b692a44e937beae8693cc155b205ffc5b732fbe

Added:
- Initial limited support for mixing operators [#107](https://github.com/openfga/language/pull/107)

## v0.2.0-beta.5

### [v0.2.0-beta.5](https://github.com/openfga/language/releases/tag/vv0.2.0-beta.4...v0.2.0-beta.5) (2023-10-04)

Added:
- Initial support for Conditions when transforming from DSL to JSON (https://github.com/openfga/language/pull/75)

## v0.2.0-beta.4

### [v0.2.0-beta.4](https://github.com/openfga/language/releases/tag/vv0.2.0-beta.3...v0.2.0-beta.4) (2023-10-04)

Added:
- Initial support for ABAC when transforming from DSL to JSON (https://github.com/openfga/language/pull/75)

## v0.2.0-beta.2, v0.2.0-beta.3

### [v0.2.0-beta.2](https://github.com/openfga/language/releases/tag/vv0.2.0-beta.1...v0.2.0-beta.2) (2023-09-21)

Fixes:
- Fixed improper offset on the duplicate relations error (https://github.com/openfga/language/pull/70)
- Fixed other errors triggering before duplicate type (https://github.com/openfga/language/pull/67)
- Add README to the published SDK (https://github.com/openfga/language/pull/65)

## v0.2.0-beta.1

### [v0.2.0-beta.1](https://github.com/openfga/language/releases/tag/v0.2.0-language) (2023-09-19)

Changed:
[BREAKING]
- `friendlySyntaxToApiSyntax` is now `transformer.transformDSLToJSON`
- `apiSyntaxToFriendlySyntax` is now `transformer.transformJSONToDSL`
- `checkDsl` is now `validator.validateDSL`
- All non-transformer functions have been removed

You can read the pre-v0.2.0 changelog [here](https://github.com/openfga/syntax-transformer/blob/main/CHANGELOG.md)