# Error-message reconciliation: Go validation vs. canonical JS reference

## Status

Open. This is the work needed to make the Go semantic validator agree with the
canonical reference implementation. It is tracked separately from the structural
migration (go-sdk → proto) and the engine wiring, both of which are done.

## Why this exists

`pkg/js/validator/validate-dsl.ts` is the canonical, battle-tested semantic
validator for OpenFGA models. The Go validator in `pkg/go/validation/` is a port
of it. Both are exercised against the **same shared fixture file**:

```
tests/data/dsl-semantic-validation-cases.yaml   (82 semantic cases)
```

Each case pins an exact expected error: message text, error count, line/column,
symbol, and `errorType`. The Go integration harness
(`yaml_test_integration.go` / `yaml_integration_test.go`) loads these cases,
runs `ValidateDSL`, and compares actual errors against expected.

As of this writing the Go validator passes **6 / 82** of those cases. The
remaining ~76 fail almost entirely on **error-message wording and error
count** — not on whether a problem is detected. The validator usually finds the
right issue; it describes it in different words, or emits extra cascading errors
the reference does not.

> Note: today these mismatches do **not** fail `go test`. The harness logs
> `FAIL` via `t.Logf` and never calls `t.Fail()`. Making the suite enforce
> (so red is red) is a separate tracked item; it should be turned on only once
> the messages below are reconciled, otherwise the build goes red immediately.

## Source of truth

| | File |
|---|---|
| Canonical messages | `pkg/js/util/exceptions.ts` (the `raise*` methods) |
| Canonical logic | `pkg/js/validator/validate-dsl.ts` |
| Shared fixtures | `tests/data/dsl-semantic-validation-cases.yaml` |
| Go messages | `pkg/go/validation/error_collector.go` |
| Go logic | `pkg/go/validation/*.go` |

When in doubt, the JS string is correct. The fixture YAML is generated against
it, so matching `exceptions.ts` verbatim is the goal.

## The three failure classes

### 1. Message wording mismatch (the bulk of failures)

Same error detected, different text. Fix = align the Go `fmt.Sprintf` template
in `error_collector.go` to the JS string in `exceptions.ts`, character for
character (including backticks vs. straight quotes — the reference uses
backticks around symbols in many messages, Go currently uses single quotes).

Concrete mappings observed in the failing fixtures:

| Concept | Canonical (JS, `exceptions.ts`) | Current (Go, `error_collector.go`) |
|---|---|---|
| Reserved type name | `` a type cannot be named 'self' or 'this'. `` | matches (`RaiseReservedTypeName`) — keep ✓ |
| Reserved relation name | `` a relation cannot be named 'self' or 'this'. `` | matches (`RaiseReservedRelationName`) — keep ✓ |
| Invalid name (naming rule) | `` ...does not match naming rule: '${clause}'. `` where `clause` is the **anchored** regex `^[^:#@\*\s]{1,254}$` | `type '%s' does not match naming rule: '%s'.` but passes the **unanchored** `clause` (`[^:#@\*\s]{1,254}`) — add `^...$` |
| Undefined relation | `` `allowed` does not exist. `` (raiseInvalidRelationError → `the relation \`%s\` does not exist.`) | `Relation 'allowed' is not defined on type 'group' (referenced in relation 'reader' of type 'group')` — wrong template **and** wrong raise method |
| Invalid relation for type | `` `org` is not a valid relation for `group`. `` | `Relation 'org' is not defined on type 'group' (...)` |
| Invalid type | `` `customer` is not a valid type. `` | `type '%s' is not defined.` |
| No entrypoint | `` `viewer` is an impossible relation for `group` (no entrypoint). `` | `Relation 'viewer' on type 'group' has no entry point` |
| No entrypoint (loop) | `` `x` is an impossible relation for `resource` (potential loop). `` | `Relation 'x' on type 'resource' contains a cycle` + `...has no entry point` (two errors, see class 3) |
| Tupleset not direct | `` `parent` relation used inside from allows only direct relation. `` | matches (`RaiseTupleUsersetRequiresDirect`) — keep |
| Duplicate type restriction | `` the type restriction \`viewer\` is a duplicate in the relation \`editor\`. `` | `the type 'viewer' is defined more than once in relation 'editor' of type 'document'.` |
| Duplicate partial relation | `` the partial relation definition \`viewer\` is a duplicate in the relation \`editor\`. `` | not emitted / different text |
| Schema version required | `schema version required` | `a schema version is required in the model.` |
| Schema version unsupported | `schema version no longer supported` | `the schema version '%s' is not supported.` |
| Max one direct relationship | `each relationship must have at most 1 set of direct relations defined.` | `the relation '%s' can have at most one direct relationship.` |
| Undefined condition (param) | `` `allowed_ip` is not a defined condition in the model. `` | `condition parameter name '%s' is invalid.` |

This table is representative, not exhaustive — every `raise*` method in
`exceptions.ts` should be diffed against its Go counterpart in
`error_collector.go`. There are ~26 raise methods.

**Recommended approach:** treat `exceptions.ts` as the spec. Go method by
method, copy the JS template string into the Go `Sprintf`, preserving backticks
and argument order. This is mechanical and low-risk once started.

### 2. Wrong raise method / errorType (semantic, not just text)

Some Go validators reach for a different error category than the reference. The
clearest case is "relation referenced but not defined":

- Reference distinguishes:
  - `raiseInvalidRelationError` → `` the relation \`x\` does not exist. ``
    (a relation referenced in a rewrite that doesn't exist on the type)
  - `raiseInvalidTypeRelation` → `` the \`reader\` relation definition on type
    \`group\` is not valid: \`reader\` does not exist on \`parent\`... ``
    (a `X from Y` where the computed relation is missing on the target type)
- Go currently funnels several of these into a single
  `RaiseUndefinedRelation` / `RaiseInvalidRelationError` with a generic
  "is not defined on type ... (referenced in ...)" template.

Fixing class 1 wording is not enough here; the **errorType** and the **choice
of raise method** must match so the fixture's `metadata.errorType` assertion
passes. This requires reading `validate-dsl.ts` around lines 500–600 (the
`from`/computed-relation validation) to mirror which error fires when.

### 3. Cascading / duplicate errors (error-count mismatch)

The reference emits **one** error per root cause. The Go validator often emits
several for the same underlying problem. Example from fixture case 1:

```
expected: 1  → the relation `allowed` does not exist.
got:       2 → Relation 'allowed' is not defined on type 'group' (...)
               Relation 'reader' on type 'group' has no entry point
```

And the no-entrypoint cases routinely emit both `contains a cycle` **and**
`has no entry point` for one relation, where the reference emits a single
`(no entrypoint)` or `(potential loop)`.

Root cause: the Go semantic phase and cycle/entrypoint phase both fire on the
same broken relation without coordination. The reference suppresses the
downstream error once the root error is recorded (a relation that references a
non-existent relation should not *also* be reported as having no entry point —
the first error subsumes it).

Fix is logic, not text: when a relation is already flagged (undefined
reference, etc.), skip the derived entrypoint/cycle complaint for that same
relation. Look at how `validate-dsl.ts` orders and short-circuits its checks.

## Suggested order of work

1. **Class 1 (wording)** first — mechanical, unblocks the most fixtures, and
   makes the remaining diffs readable. Diff every `raise*` in `exceptions.ts`
   against `error_collector.go`.
2. **Class 2 (raise method / errorType)** — needs `validate-dsl.ts` reading but
   is well-scoped to the `from`/computed-relation and undefined-relation paths.
3. **Class 3 (cascading)** — the hardest; requires cross-phase coordination so
   one root cause yields one error. Do last, measuring fixture pass-count after
   each change.
4. Only once the pass count is high, flip the YAML harness from `t.Logf` to a
   real `t.Fail()`/`require` so regressions are caught.

## How to measure progress

```
cd pkg/go
go test ./validation/... -v -run TestYAMLTestRunner_ComprehensiveValidation 2>&1 \
  | grep -E "Total Tests:|Passed:|Failed:"
```

Baseline at time of writing: **6 / 82 passing**. Each reconciliation step should
move this number up; if it drops, the change regressed a previously-passing
case.

## What is explicitly NOT in scope here

- The go-sdk → proto migration (done).
- Wiring name/reserved-keyword validation into the engine (done; that is why the
  4 reserved-name cases now pass).
- The two intentional stubs `checkSubsumingUnionMembers` /
  `checkRedundantIntersectionMembers` — these are unimplemented advanced checks
  in both the original Go branch and are not required by the fixtures.
