# Changelog

## [0.4.0](https://github.com/openfga/language/compare/pkg/go/v0.3.2...pkg/go/v0.4.0) (2026-11-17)


### ⚠ BREAKING CHANGES

* **go:** x/net v0.60.0 requires Go 1.26, so the minimum supported Go version for pkg/go is now 1.26 (was 1.25). The CI build matrix now tests Go 1.26 and 1.27.

### Fixed

* **go:** bump golang.org/x/net to v0.60.0 to resolve govulncheck findings ([#699](https://github.com/openfga/language/issues/699)) ([9248a88](https://github.com/openfga/language/commit/9248a88b64db3e4eb9e5f88a594e84ee87fd7ed4))
* **go:** handle proto direct assignments in DSL transform ([#621](https://github.com/openfga/language/issues/621)) ([ed62a11](https://github.com/openfga/language/commit/ed62a11e759a0faa1354404ead761f956469e149))


### Changed

* **graph:** use strings.IndexByte instead of strings.Contains for single char ([#694](https://github.com/openfga/language/issues/694)) ([2aed764](https://github.com/openfga/language/commit/2aed7640561f1bd836b68b3eb7c3e6fe8670ba4a))

## [0.3.2](https://github.com/openfga/language/compare/pkg/go/v0.3.1...pkg/go/v0.3.2) (2026-10-01)


### Added

* create deterministic name to the operator nodes ([#640](https://github.com/openfga/language/issues/640)) ([b8b1a5a](https://github.com/openfga/language/commit/b8b1a5ae1ff2c040ee3c84c23560ce74a01b3f24))
* support $expression reserved condition name on type restrictions ([#664](https://github.com/openfga/language/issues/664)) ([0d2ad7f](https://github.com/openfga/language/commit/0d2ad7fb7c40a9ade99ddbfaffc3cc12f5186456))


### Fixed

* flag multi-branch recursive relations as tuple cycles in the weighted graph ([#648](https://github.com/openfga/language/issues/648)) ([83fedf8](https://github.com/openfga/language/commit/83fedf8a4e70499aaa2f69d713084a410b687ad3))
* **graph:** prevent nil-pointer panic and corruption in parseThis ([#661](https://github.com/openfga/language/issues/661)) ([96eccfe](https://github.com/openfga/language/commit/96eccfe5f3b76097001c33ad2edaf319b64a554b))

## [0.3.1](https://github.com/openfga/language/compare/pkg/go/v0.3.0...pkg/go/v0.3.1) (2026-07-06)


### Added

* allow '-' and '/' in object IDs ([#624](https://github.com/openfga/language/issues/624)) ([e9a8cce](https://github.com/openfga/language/commit/e9a8cce18daf33b83d394a20daf96292c1bf9a54)), closes [#437](https://github.com/openfga/language/issues/437)


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
