# Graph Model Unbuildable

**Error Code:** `graph-model-unbuildable`

**Category:** Semantic Validation

## Summary

The model could not be built into a weighted authorization model graph, so the checks
that read that graph had nothing to read. The message carries the reason the graph
builder gave.

## Description

Entrypoint and cycle validation can be answered two ways: by walking the model's
rewrite tree, or by building the same weighted graph the server builds and reading it.
The second route needs a graph, and the builder refuses to produce one for some
models. This code reports that refusal.

It is raised only when graph-backed validation is enabled, which is not the default.
With the default options the models below report
[`relation-no-entry-point`](./relation-no-entry-point.md) instead, one finding per
affected relation, each carrying a line and column.

The builder refuses a model for reasons of its own, and the message ends with the one
it gave:

- `model cycle`, when relations refer to each other in a loop that admits no
  assignment
- `tuple cycle: ...`, when a cycle cannot be resolved, most often because it runs
  through an intersection or an exclusion
- `invalid model: ...`, when the graph cannot be constructed from the definitions
  given, with the specific reason following the colon

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
    }
}
```

## Example

The following model would trigger this error:

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

**Error Message:** `the model cannot be built into a weighted graph: model cycle`

`viewer` requires `editor`, `editor` is `viewer`, and neither can be entered. The
builder stops there.

A cycle that runs through an exclusion is refused with a different reason:

```
model
  schema 1.1

type user

type folder
  relations
    define parent: [folder]
    define viewer: [user] but not banned
    define banned: viewer from parent
```

**Error Message:** `the model cannot be built into a weighted graph: tuple cycle: operands AND or BUT NOT cannot be involved in a cycle`

## Resolution

The model is invalid, and the fix is the same fix the reason calls for. Read the
message after the colon and treat it as the finding.

For a cycle, break the loop, so that one relation in it stops depending on the rest.
Giving a relation in the loop a way to be entered is not enough on its own, because the
loop still resolves to itself:

```
model
  schema 1.1

type user

type document
  relations
    define admin: [user]
    define viewer: admin and editor
    define editor: [user]
```

For a cycle through `and` or `but not`, take the cyclic relation out of the operand:

```
model
  schema 1.1

type user

type folder
  relations
    define parent: [folder]
    define banned: [user]
    define viewer: [user] but not banned
```

### Steps to fix:

1. **Read the reason:** the text after the colon is the graph builder's own error, and
   it names the class of problem.

2. **Find the relations involved:** this finding is model-level and carries no
   position, because the builder stops at the first problem and returns no graph to
   locate it in. Running validation with the default options reports the same model as
   `relation-no-entry-point`, once per affected relation, with a line and column for
   each.

3. **Fix the model, not the finding:** every reason the builder gives is a model that
   cannot answer a check. There is no configuration that makes one of these models
   valid.

4. **Re-validate:** a model the builder accepts still gets the graph-backed
   entrypoint checks run over it, so a clean build is not on its own a clean model.

## Related Errors

- [`relation-no-entry-point`](./relation-no-entry-point.md) - what the default
  validation path reports for these models, per relation and with positions
- [`cyclic-error`](./cyclic-error.md) - circular dependency between relations
- [`invalid-relation-type`](./invalid-relation-type.md) - a relation that is not valid
  for the type it is used with

## Implementation Notes

This code is specific to the Go implementation's graph-backed validation path. The
JavaScript and Java validators walk the rewrite tree and have no equivalent.

- Go implementation: `pkg/go/validation/graph_validation.go`
- Graph builder: `pkg/go/graph/weighted_graph_builder.go`

Validation calls the exported builder rather than reimplementing it, so a model refused
here is refused for anything else that builds the same graph.
