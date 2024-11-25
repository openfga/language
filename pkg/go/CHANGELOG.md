# Changelog

## pkg/go/v0.2.0-beta.3

### [v0.2.0-beta.3](https://github.com/openfga/language/compare/pkg/go/v0.2.0-beta.2...pkg/go/v0.2.0-beta.3) (2024-11-25)

Added:

- Ability to visualize graph in reverse (i.e. the direction of a Check request) (#345)
- Model graph now has edges for computed usersets (#342)
- Model graph now includes cycle detection (#344)
- Model graph now exposes extra getters (#379)

Fixed:

- Error when a model file is provided as part of a modular model (#386) - thanks @fsedano!
- Model graphing now correctly tracks wildcard relation types (#356)
- Validation of Object IDs is aligned with the OpenFGA server (#348)

## pkg/go/v0.2.0-beta.2

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
