# Changelog

## [0.3.0](https://github.com/openfga/language/compare/pkg/java/v0.2.1...pkg/java/v0.3.0) (2026-11-17)


### ⚠ BREAKING CHANGES

* openfga-language now requires Java 17. openfga-sdk 0.11.2 is compiled for Java 17, so Java 11 is no longer supported. The SDK also brings Jackson 3 (tools.jackson 3.2.4) and the OpenTelemetry API 1.67 onto the runtime classpath, and no longer brings jackson-datatype-jsr310 or jackson-databind-nullable.

### Added

* require Java 17 and bump openfga-sdk to 0.11.2 ([#725](https://github.com/openfga/language/issues/725)) ([102d5c1](https://github.com/openfga/language/commit/102d5c19dae738e26f3849006b2f08311e67ab47))
* support $expression reserved condition name on type restrictions ([#664](https://github.com/openfga/language/issues/664)) ([0d2ad7f](https://github.com/openfga/language/commit/0d2ad7fb7c40a9ade99ddbfaffc3cc12f5186456))


### Fixed

* correct error locations when type/schema lines have trailing comments ([#628](https://github.com/openfga/language/issues/628)) ([92d7b52](https://github.com/openfga/language/commit/92d7b5278a316ca9efa2aa38bf024191ce9cc37a))
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
