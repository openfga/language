# Changelog

## pkg/js/v0.2.0-beta.20

### [v0.2.0-beta.20](https://github.com/openfga/language/compare/pkg/js/v0.2.0-beta.19...pkg/js/v0.2.0-beta.20) (2024-09-06)

Added:

- Add `getModuleForObjectTypeRelation` utility method (#336)
- Add `isRelationAssignable` utility method (#336)

Fixed:

- correct reported location of errors in some cases (#127)
- `tupleuserset-not-direct` is now prioritized above `no-entrypoint` error (#314)
- condition name is now correctly restrict to 50 characters (#320)

## pkg/js/v0.2.0-beta.19

### [v0.2.0-beta.19](https://github.com/openfga/language/compare/pkg/js/v0.2.0-beta.18...pkg/js/v0.2.0-beta.19) (2024-07-25)

Added:

- migrated validation code for store files (#286)
- getModulesFromJSON: function to get modules from json (#287)

## pkg/js/v0.2.0-beta.18

### [v0.2.0-beta.18](https://github.com/openfga/language/compare/34a1a85f8307547cfd44d270472c317cc64962dd...6bf7c3e4910b7acc1ff80d8c66a8e1277e949fe6) (2024-06-12)

Fixed:

- Incorrectly assigned error for invalid-relation-type (#260)

Added:

- Implement fga.mod transformer (#243)

## pkg/js/v0.2.0-beta.17

### [v0.2.0-beta.17](https://github.com/openfga/language/compare/f45e80a9f169370f7cbeb5da9d147e5ff0e7305e...7b0b7f08393b13d8c53509447ff6e3847542eefc) (2024-04-12)

Fixes:

- Throw syntax error for missing module name (#233)
- Validate that only one module is defined per file in json (#232)
- Only allow a type to be extended once per file (#232)
- Restrict contents items to be underneath or alongside fga.mod file (#228) - thanks @Siddhant-K-code
- Remove circular dependency (#229)

## v0.2.0-beta.16

### [v0.2.0-beta.16](https://github.com/openfga/language/compare/0a0d96a126f87daf34d00514cc9c7ceb8a959993...7484f693f40aeed47ca1e2a2c6109758876a4ae3) (2024-04-04)

Refactor:

- Make line and column numbers zero based (https://github.com/openfga/language/pull/223)

## v0.2.0-beta.15

### [v0.2.0-beta.15](https://github.com/openfga/language/compare/78ce3fed297528d1ef8083b4bef1500b26dada78...5ead374fcd3d6c0678303d9020ff5eacc4f7ed63) (2024-04-01)

Fixed:

- Incorrect range for 'condition-not-defined' (https://github.com/openfga/language/pull/219)

## v0.2.0-beta.14

### [v0.2.0-beta.14](https://github.com/openfga/language/compare/9974b568edcc3e2c5bdde5143b5a4a13b45c7968...f9816aad7efa32a6ab0b9f004c7831628906de18) (2024-03-27)

Fixed:

- Better validation for [Modular Models](https://github.com/openfga/rfcs/blob/main/20231212-modular-models.md)

## v0.2.0-beta.13

### [v0.2.0-beta.13](https://github.com/openfga/language/compare/152c063895fb866f0e8da9d4c3707f8c0d01e28c...3d05a4fde617d54129c1c9f8eed1c9541b261096) (2024-03-22)

Added:

- Expose better validation and transformation for [Modular Models](https://github.com/openfga/rfcs/blob/main/20231212-modular-models.md)

## v0.2.0-beta.12

### [v0.2.0-beta.12](https://github.com/openfga/language/compare/323aac545377c0a26486db1323ce26ca5e4e1443...8a7e04f4da52f8be1c9321ff4b125a8fef1e6102) (2024-03-19)

Added:

- Initial support for [Modular Models](https://github.com/openfga/rfcs/blob/main/20231212-modular-models.md) (#148)

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
