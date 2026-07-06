# Changelog

## [0.4.0](https://github.com/openfga/language/compare/pkg/go/v0.3.0...pkg/go/v0.4.0) (2026-07-06)


### Added

* allow '-' and '/' in object IDs ([#624](https://github.com/openfga/language/issues/624)) ([e9a8cce](https://github.com/openfga/language/commit/e9a8cce18daf33b83d394a20daf96292c1bf9a54)), closes [#437](https://github.com/openfga/language/issues/437)


### Miscellaneous

* **pkg/go:** release 0.4.0 ([0680d0c](https://github.com/openfga/language/commit/0680d0c0c4f0fc084d5be421e83a171796c4a6a8))

## [0.3.0](https://github.com/openfga/language/compare/pkg/go/v0.2.1...pkg/go/v0.3.0) (2026-06-16)

> [!WARNING]  
> **BREAKING CHANGES**:
> 
> **pkg/go/transformer:** TransformDSLToProto and TransformModularDSLToProto now return *OpenFgaDslSyntaxMultipleError instead of *multierror.Error.

### Changed

* **pkg/go/transformer:** replace go-multierror with stdlib ([#601](https://github.com/openfga/language/issues/601)) ([7fd286c](https://github.com/openfga/language/commit/7fd286c295788871e8333f42bb5e0ee2e27074e4))


## pkg/go/v0.2.1

### [v0.2.1](https://github.com/openfga/language/compare/pkg/go/v0.2.0-beta.2...pkg/go/v0.2.1) (2026-03-25)

Added:
- Functionality to validate the write relation between a node and a UserType (#518)
- Functionality to validate write tuples (#517)
- Locking for UserSet weight writes (#514)
- UserSet weight to edge (#509)
- UserSet weights (#505)
- Relation definition for all the edges (#494)
- Differentiating "recursive" cycles from tuple cycles (#454)
- Allow '/' and '.' in type and relation names (#254)
- Deduplicate edges when conditions exist and add conditions to the edges (#422)

Changed:
- Removed the overhead in openfga to rewrite NoCond from "" to "None" and vice versa (#508)

Fixed:
- Recursive nested tuple cycles assignments (#497)
- Correctly assign weights for intersections with multiple direct types (#455)
- Exclusion weight where type is not in the excluded relation (#443)

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
