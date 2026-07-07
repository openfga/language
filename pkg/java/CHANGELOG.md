# Changelog

## [0.2.2](https://github.com/openfga/language/compare/pkg/java/v0.2.1...pkg/java/v0.2.2) (2026-07-07)


### Fixed

* **java:** derive publish coordinates from project instead of hardcoded values ([67c43b7](https://github.com/openfga/language/commit/67c43b7e506ed97bc1d79d5af7e627193001111a))

## [0.2.1](https://github.com/openfga/language/compare/pkg/java/v0.2.0-beta.2...pkg/java/v0.2.1) (2026-07-06)


### Added

* allow '-' and '/' in object IDs ([#624](https://github.com/openfga/language/issues/624)) ([e9a8cce](https://github.com/openfga/language/commit/e9a8cce18daf33b83d394a20daf96292c1bf9a54)), closes [#437](https://github.com/openfga/language/issues/437)
* allow '/' and '.' in type and relation names ([#254](https://github.com/openfga/language/issues/254)) ([8862530](https://github.com/openfga/language/commit/88625300b642f9b51fd0140f7190e702658abd09))


## [v0.2.0-beta.2](https://github.com/openfga/language/compare/pkg/java/v0.2.0-beta.1...pkg/java/v0.2.0-beta.2) (2024-09-06)

Added:
- Add `getModuleForObjectTypeRelation` utility method (#336)
- Add `isRelationAssignable` utility method (#336)
- Add utility validators for tuple fields and condition names (#294)

- Fixed:
- `tupleuserset-not-direct` is now prioritized above `no-entrypoint` error (#314)
- correct based index for reported errors that was causing the wrong location to be highlighted (#331)

## pkg/java/v0.2.0-beta.1

### [v0.2.0-beta.1](https://github.com/openfga/language/tree/7b8d22c70355fb7a0796b17d18eafaaa6360759b/pkg/java) (2024-06-13)

- Initial release
