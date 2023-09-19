# Changelog

## v0.2.0-language

### [0.2.0-language](https://github.com/openfga/language/releases/tag/v0.2.0-language) (2023-09-19)

Changed:
[BREAKING]
- `friendlySyntaxToApiSyntax` is now `transformer.transformDslToJSON`
- `apiSyntaxToFriendlySyntax` is now `transformer.transformJSONToDSL`
- `checkDsl` is now `validator.validateDsl`
- All non-transformer functions have been removed

You can read the pre-v0.2.0 changelog [here](https://github.com/openfga/syntax-transformer/blob/main/CHANGELOG.md)