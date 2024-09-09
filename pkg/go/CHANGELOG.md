# Changelog

## pkg/go/v0.2.0-beta.1

### [v0.2.0-beta.2](https://github.com/openfga/language/compare/pkg/go/v0.2.0-beta.1...pkg/go/v0.2.0-beta.2) (2024-09-09)

Fixed:

- Updated generated antlr code to fix issue when parsing modules (#339)

## pkg/go/v0.2.0-beta.1

### [v0.2.0-beta.1](https://github.com/openfga/language/compare/pkg/go/v0.2.0-beta.0...pkg/go/v0.2.0-beta.1) (2024-09-06)

Added:
- Add `GetModuleForObjectTypeRelation` utility method (#336)
- Add `IsRelationAssignable` utility method (#336)
- Add utility validators for tuple fields and condition names (#294)
- Add an initial implementation of a model graph (#307,#308,#309,#310,#316,#317,#321,#322,#330)

Note: this version does not include the validation logic that the JS and Java ports have. See issue: https://github.com/openfga/language/issues/99

## pkg/go/v0.2.0-beta.0

### [v0.2.0-beta.0](https://github.com/openfga/language/tree/a3958b8187145f3a1f98f1d7334ba49411521cc8/pkg/go) (2024-06-12)

- Initial release

Note: this version does not include the validation logic that the JS and Java ports have. See issue: https://github.com/openfga/language/issues/99
