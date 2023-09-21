# Changelog

## v0.2.0-beta.2

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