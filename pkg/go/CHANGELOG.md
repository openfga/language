# Changelog

## Unreleased

### [Unreleased](https://github.com/openfga/language/compare/pkg/go/v0.2.1...HEAD)

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
