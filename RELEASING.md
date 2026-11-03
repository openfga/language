# Releasing

Releases are automated with [release-please](https://github.com/googleapis/release-please). This repo hosts three independently versioned packages: `pkg/go`, `pkg/js` and `pkg/java`.

## Cutting a release

1. Go to **Actions → Release Please → Run workflow**. Choose the target package (or `all`) and a bump type (`auto` follows conventional commits), or pick `explicit` and enter the version in `release-version`.
   - `explicit` with `all` applies the same version to every package, so target a single package instead.
   - For a pre-release, use `explicit` with `X.Y.Z-alpha.N`, `-beta.N` or `-rc.N`. The GitHub release is marked as a pre-release and npm publishes under the matching dist-tag. Other suffixes fail the `pkg/js` publish.
2. release-please opens one PR per package, titled `release: pkg/<name> vX.Y.Z`, with the version bump and that package's updated `CHANGELOG.md`. Review it. Each run closes all open release PRs, including other packages', so merge any pending one first.
3. Squash-merge it, keeping the `release:` prefix in the squash commit subject. CI tags the commit (`pkg/<name>/vX.Y.Z`) and publishes the release automatically. The GitHub release stays a draft until the package's `Release (Go)`, `Release (JS)` or `Release (Java)` run succeeds; if it fails, re-run its failed jobs.

## Version bumps

Bumps follow [conventional commits](https://www.conventionalcommits.org/), with `bump-minor-pre-major` and `bump-patch-for-minor-pre-major` enabled:

| Commit | Before v1.0.0 | v1.0.0 and later |
|---|---|---|
| `feat:` | patch | minor |
| `fix:`, `perf:`, `refactor:`, `revert:`, `docs:` | patch | patch |
| `type!:` or `BREAKING CHANGE:` footer | minor | major |

## Changelog

Built from commit messages since the last release: `feat:` → Added, `fix:` → Fixed, `perf:`/`refactor:` → Changed, `revert:` → Removed, `docs:` → Documentation. `test:`, `ci:` and `chore:` commits are hidden.

### Overriding an entry

If a commit message isn't user-facing, add this block to the PR description. This works after merging too; re-run the workflow to refresh an open release PR. release-please uses it instead of the commit message, but package attribution still follows the files the PR changed:

```
BEGIN_COMMIT_OVERRIDE
feat: user-facing description of the change
END_COMMIT_OVERRIDE
```

### Breaking changes

Mark the type with `!` and add a `BREAKING CHANGE:` footer. PRs are squash-merged, and the squash title defaults to the PR title only for multi-commit PRs, so before merging check that the squash commit title has the `!` and its body has the footer:

```
feat!: remove the Foo option

BREAKING CHANGE: the Foo option has been removed. Use Bar instead.
```

This gets its own changelog section and triggers a major bump (minor before v1.0.0).
