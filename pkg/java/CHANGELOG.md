# Changelog

## pkg/java/v0.2.0-beta.2

## [0.2.1-beta.2](https://github.com/openfga/language/compare/pkg/java/v0.2.0-beta.2...pkg/java/v0.2.1-beta.2) (2026-06-12)


### Added

* allow '/' and '.' in type and relation names ([#254](https://github.com/openfga/language/issues/254)) ([8862530](https://github.com/openfga/language/commit/88625300b642f9b51fd0140f7190e702658abd09))
* **language:** added unit test for java ([2e0fb5d](https://github.com/openfga/language/commit/2e0fb5d18ee3d66911ff33ac7ae3329e870507fb))
* **lanuage:** added objectId regex for go and java ([6bb8f13](https://github.com/openfga/language/commit/6bb8f13568e42cd98d473b924a747537e5dc32cd))


### Fixed

* fix validation of object IDs ([#348](https://github.com/openfga/language/issues/348)) ([dc0d9f1](https://github.com/openfga/language/commit/dc0d9f1e83a164a95e2a0352128a3a8fb373250f))
* **language:** fix java lint issue ([23f369d](https://github.com/openfga/language/commit/23f369d63fddc45653f67177d8d42c9168e4f8d8))

### [v0.2.0-beta.2](https://github.com/openfga/language/compare/pkg/java/v0.2.0-beta.1...pkg/java/v0.2.0-beta.2) (2024-09-06)

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
