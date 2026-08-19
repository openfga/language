# Graph Model Unbuildable

**Error Code:** `graph-model-unbuildable`

**Category:** Semantic Validation

## Summary

The model could not be built into a weighted authorization model graph, and nothing else
had a finding for it either. This is the last resort finding: it names the class of
problem and no relation, so it appears only when no other check could say something more
useful.

## Description

Entrypoint and cycle validation can be answered two ways: by walking the model's
rewrite tree, or by building the same weighted graph the server builds and reading it.
The second route needs a graph, and the builder refuses to produce one for some
models.

A refused build is not reported as this code straight away. Three things are tried
first, and this code is what is left when all three come up empty:

1. The rewrite-tree walk runs, which reports
   [`relation-no-entry-point`](./relation-no-entry-point.md) per relation with a line
   and column. Most refused models are answered here.
2. If the walk finds nothing, relations taking part in a cycle the resolver cannot work
   through are reported as [`cyclic-relation`](./cyclic-relation.md), again per relation
   with a position.
3. Only if neither has anything to say is the refusal itself reported, as this code.

It is raised only when graph-backed validation is enabled, which is not the default.

The message ends with the sentinel the graph builder returned, one of `model cycle`,
`tuple cycle`, `tuple cycle: operands AND or BUT NOT cannot be involved in a cycle`, or
`invalid model`.

That is the sentinel's own text and nothing more. The builder's own message often goes
on to name a type or a relation, and that part is deliberately left out: the builder
stops at the first problem it meets and picks which one that is by ranging over a map,
so on a model with several independent problems it names one of them and names a
different one on the next run. Every such message is true, and none is stable enough to
put in a finding. The builder's full text stays reachable by unwrapping.

A finding with this code wraps two errors. `errors.Is` matches
`errors.ErrModelNotBuildable` for the refusal itself, and it also matches whichever
sentinel the graph package returned, so a caller can branch on the specific reason
without parsing the message:

```go
var findings *validation.ValidationErrors
if errors.As(err, &findings) {
    for _, f := range findings.GetErrors() {
        if errors.Is(f, fgaerrors.ErrModelNotBuildable) {
            // the build was refused
        }
        if errors.Is(f, graph.ErrTupleCycle) {
            // and this is why
        }
        // f.Unwrap() reaches the builder's own message, first problem and all
    }
}
```

## Example

The following model reports this error:

```
model
  schema 1.1

type user

type folder
  relations
    define viewer: [user]

type team

type document
  relations
    define parent: [folder, team]
    define viewer: viewer from parent
```

**Error Message:** `the model cannot be built into a weighted graph: invalid model`

`parent` accepts a `folder` or a `team`, and `viewer from parent` needs a `viewer`
relation on both. `folder` has one and `team` does not. Every relation here has a way in,
so the rewrite-tree walk reports nothing, and there is no cycle, so nothing names a
relation. The refusal is all there is to report.

Cycles do not reach this code. Both of the following are refused by the builder, and both
are reported per relation with a position instead:

```
model
  schema 1.1

type user

type document
  relations
    define admin: [user]
    define viewer: admin and editor
    define editor: viewer
```

reports `relation-no-entry-point` for `viewer` and `editor`, because neither can be
entered.

```
model
  schema 1.1

type user

type document
  relations
    define a: [user] or b
    define b: [user] or a
```

reports `cyclic-relation` for `a` and `b`. Both hold a plain `[user]`, so both can be
entered and the walk is silent, but going round the cycle reads no tuple, so resolving
either one never terminates.

## Resolution

The model is invalid and the sentinel names the class of problem. Unwrap the finding for
the builder's own message if you want the first specific instance it hit, remembering it
is one of possibly several.

For `invalid model` on a tupleset, give the computed relation to every type the tupleset
accepts:

```
model
  schema 1.1

type user

type folder
  relations
    define viewer: [user]

type team
  relations
    define viewer: [user]

type document
  relations
    define parent: [folder, team]
    define viewer: viewer from parent
```

or narrow the tupleset to the types that have it:

```
model
  schema 1.1

type user

type folder
  relations
    define viewer: [user]

type team

type document
  relations
    define parent: [folder]
    define viewer: viewer from parent
```

### Steps to fix:

1. **Read the sentinel:** the text after the last colon names the class of problem.

2. **Unwrap for the detail:** the builder's own message names the first problem it met.
   Treat it as one instance rather than the whole list, and re-validate after fixing it.

3. **Expect no position:** this finding is model-level. The builder stops at the first
   problem and returns no graph to locate it in, and anything that could have been
   located has already been reported as `relation-no-entry-point` or `cyclic-relation`.

4. **Re-validate:** a model the builder accepts still gets the graph-backed entrypoint
   checks run over it, so a clean build is not on its own a clean model.

## Related Errors

- [`cyclic-relation`](./cyclic-relation.md) - a relation in a cycle the resolver cannot
  work through, reported per relation with a position
- [`relation-no-entry-point`](./relation-no-entry-point.md) - what most refused models
  report instead, per relation and with positions
- [`invalid-relation-on-tupleset`](./invalid-relation-on-tupleset.md) - a tupleset whose
  computed relation is missing everywhere, rather than on some types only

## Implementation Notes

This code is specific to the Go implementation's graph-backed validation path. The
JavaScript and Java validators walk the rewrite tree and have no equivalent.

- Go implementation: `pkg/go/validation/graph_validation.go`
- Graph builder: `pkg/go/graph/weighted_graph_builder.go`

Validation calls the exported builder rather than reimplementing it, so a model refused
here is refused for anything else that builds the same graph.
