# Graph Model Unbuildable

**Error Code:** `graph-model-unbuildable`

**Category:** Semantic Validation (Go graph validation only)

## Summary

The weighted authorization-model graph builder refused the model, and the
refusal cannot be attributed to specific relations.

## Description

Emitted only by the Go package's graph-based entry points, `ValidateDSLWithGraph`
and `ValidateJSONWithGraph`, as the last of three steps. A refused model is
first answered by the rewrite-tree walk, which names unreachable relations as
[`relation-no-entry-point`](./relation-no-entry-point.md); where the walk is
silent, relations in an unresolvable cycle are named as
[`cyclic-relation`](./cyclic-relation.md). Only a refusal neither can attribute
is reported as this single finding.

The reason is a fixed phrase for the kind of failure, not the builder's own
error text. The builder names whichever broken relation it happened to meet
first, which can change between runs on the same model, so its text is not
quoted. The reason is one of:

- `invalid model`
- `model cycle`
- `tuple cycle`
- `tuple cycle: operands AND or BUT NOT cannot be involved in a cycle`

The finding is about the whole model, so it carries no line or column.

## Example

### DSL

```
model
  schema 1.1
type user
type org
type folder
  relations
    define viewer: [user]
type document
  relations
    define parent: [folder, org]
    define viewer: viewer from parent
```

**Error Message:** `the model cannot be built into a weighted graph: invalid model`

The builder requires the tuple-to-userset target relation (`viewer`) to exist on
every type assignable to the tupleset (`folder` and `org`); validation and the
other implementations require it on at least one. This is the only known model
shape where the two paths disagree, and the shared test cases pin it.

## Resolution

Define the target relation on every type the tupleset relation can hold, or
remove the types that cannot carry it. If the reason is a cycle, break it as
described under [`cyclic-relation`](./cyclic-relation.md).

## Related Errors

- [`cyclic-relation`](./cyclic-relation.md) - The same refusal, attributed to the relations in the cycle
- [`relation-no-entry-point`](./relation-no-entry-point.md) - Unreachable relations, reported by all implementations

## Implementation Notes

- Go: `graphRefusal` in `pkg/go/validation/graph_validation.go`; the stable
  wording comes from `graphRefusalReasons`.
- The expected messages, codes and positions are recorded in
  `tests/data/dsl-graph-validation-cases.yaml`, which an implementation adopts
  when it grows a graph path.
