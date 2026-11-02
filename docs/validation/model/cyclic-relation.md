# Cyclic Relation

**Error Code:** `cyclic-relation`

**Category:** Semantic Validation (Go graph validation only)

## Summary

The relation takes part in a cycle the resolver cannot work through, even though
every relation in the cycle may be satisfiable on its own.

## Description

Emitted only by the Go package's graph-based entry points, `ValidateDSLWithGraph`
and `ValidateJSONWithGraph`. The default `ValidateDSL`/`ValidateJSON` and the
JavaScript and Java validators do not run this check.

Two cycle shapes are reported, and the message ends with the one that applies:

- *no relation in it reads a tuple, so resolving it never terminates* — every
  step of the cycle rewrites another relation; going around it never consumes a
  tuple, so resolution never gets closer to an answer.
- *a relation in it is an operand of an `and` or a `but not`* — the cycle does
  read tuples, but a step of it is an operand of an intersection or exclusion,
  and the resolver cannot subtract or intersect a set it is still computing.

A relation is reported for the cycle it is in, not for anything wrong with the
relation itself: each one may have a valid entry point, which is why
[`relation-no-entry-point`](./relation-no-entry-point.md) stays silent on these
models.

## Example

### DSL

```
model
  schema 1.1
type user
type document
  relations
    define a: [user] or b
    define b: [user] or a
```

**Error Messages:**

```
`a` on `document` takes part in a cycle that cannot be resolved: no relation in it reads a tuple, so resolving it never terminates.
`b` on `document` takes part in a cycle that cannot be resolved: no relation in it reads a tuple, so resolving it never terminates.
```

Each finding points at its relation's `define` line, with the column covering
the relation name. Counting from zero, `a` is reported at line 5, columns 11 to 12.

## Resolution

Break the cycle. For the rewrite shape, remove one of the mutual references or
replace it with a direct type restriction. For the constrained shape, move the
cyclic step out of the `and`/`but not` operand, for example by defining the
subtracted set as its own relation with direct type restrictions.

## Related Errors

- [`relation-no-entry-point`](./relation-no-entry-point.md) - A relation nothing can satisfy; reported by all implementations
- [`graph-model-unbuildable`](./graph-model-unbuildable.md) - A graph refusal that cannot be attributed to relations

## Implementation Notes

- Go: `relationCycleShapes` in `pkg/go/validation/cycle_shape.go`, run only when
  the weighted-graph builder refuses the model and the rewrite-tree walk reports
  nothing. The graph stays the authority on whether a model is resolvable:
  `TestCycleShapesAgreeWithTheBuilder` holds the check to it.
- The expected messages, codes and positions are recorded in
  `tests/data/dsl-graph-validation-cases.yaml`, which an implementation adopts
  when it grows a graph path.
