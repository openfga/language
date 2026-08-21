# Cyclic Relation

**Error Code:** `cyclic-relation`

**Category:** Semantic Validation

## Summary

A relation takes part in a cycle the resolver cannot work through. The relation itself may
be perfectly satisfiable; it is reported for the cycle it belongs to.

## Description

Not every cycle between relations is a problem. `define member: [user, group#member]` is
the nested-group pattern every deployment has, and it terminates because each step around
the loop reads a tuple, so the set of groups to look at shrinks until it is empty.

Two shapes do not terminate, and this error reports the relations taking part in either.

**The cycle reads no tuple.** Every step is a rewrite, so going round the loop consumes
nothing and gets no closer to an answer:

```
define a: [user] or b
define b: [user] or a
```

**A step of the cycle is an operand of an `and` or a `but not`.** The cycle does read a
tuple, so it terminates, but the resolver cannot subtract or intersect a set it is still in
the middle of computing:

```
define member: [user, group#member] but not blocked
define blocked: [user, group#member]
```

Unlike [`relation-no-entry-point`](./relation-no-entry-point.md), this error is not about a
relation that can never be satisfied. In both examples above every relation holds a plain
`[user]`, so every one of them has a way in. That is exactly why a separate code exists:
the entrypoint check has nothing to say about these models, and telling someone to give
`a` an entrypoint it already has would send them looking for the wrong thing.

`errors.Is` on a finding with this code matches `errors.ErrRelationInUnresolvableCycle`. It
does not match `errors.ErrNoEntrypoints`.

It is raised only when graph-backed validation is enabled, which is not the default.

## Example

The following model reports this error:

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user] or editor
    define editor: admin
    define admin: viewer
```

**Error Message:** ``​`viewer` on `document` takes part in a cycle that cannot be resolved: no relation in it reads a tuple, so resolving it never terminates.``

One finding is reported per relation in the cycle, each with the line and column of its
`define`, so all three of `viewer`, `editor` and `admin` are reported here.

A cycle under an exclusion reports the other reason:

```
model
  schema 1.1

type user

type group
  relations
    define member: [user, group#member] but not blocked
    define blocked: [user, group#member]
```

**Error Message:** ``​`member` on `group` takes part in a cycle that cannot be resolved: a relation in it is an operand of an `and` or a `but not`.``

## Resolution

Break the cycle, or take it out of the operator.

### Option 1: give the cycle a rewrite-free step

For a cycle that reads no tuple, the fix is to stop one relation depending on another:

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user] or editor
    define editor: [user]
    define admin: [user] or editor
```

### Option 2: restructure into a hierarchy

Dependencies that all flow one way cannot close a loop:

```
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]
    define editor: [user] or viewer
    define admin: [user] or editor
```

### Option 3: move the recursion out of the operand

For a cycle under an `and` or a `but not`, the recursion is what has to leave the operator.
Making the other side non-recursive is not enough: `member: [user, group#member] but not
blocked` is still reported even when `blocked` is a plain `[user]`, because `member`'s own
recursion is the operand. Give the recursion its own relation and apply the operator to
that:

```
model
  schema 1.1

type user

type group
  relations
    define member: [user, group#member]
    define blocked: [user]
    define visible: member but not blocked
```

The same applies to an intersection:

```
model
  schema 1.1

type user

type group
  relations
    define admin: [user]
    define member: [user, group#member]
    define approved: member and admin
```

### Steps to fix:

1. **Read which reason it gives:** the clause after the colon says whether the cycle reads
   no tuple or sits under an operator. They call for different fixes.

2. **Use the positions:** every relation in the cycle is reported with its own line, so the
   findings together are the cycle.

3. **Check the recursion, not just the other operand:** for the operator case, a relation
   that refers to itself through a userset or a tupleset is a cycle on its own, and
   enclosing it in an `and` or a `but not` is what makes it unresolvable.

4. **Re-validate:** breaking one cycle can leave another, and a relation can sit in more
   than one.

## Related Errors

- [`relation-no-entry-point`](./relation-no-entry-point.md) - a relation nothing can
  satisfy, as against one that is satisfiable but caught in a cycle
- [`graph-model-unbuildable`](./graph-model-unbuildable.md) - a refused build that no
  per-relation check could account for
- [`cyclic-error`](./cyclic-error.md) - declared but not raised; a cycle with no entry
  point surfaces as `relation-no-entry-point`

## Implementation Notes

This code is specific to the Go implementation's graph-backed validation path. The
JavaScript and Java validators walk the rewrite tree, which answers has-an-entry-point for
both shapes above and reports nothing, so they have no equivalent.

- Go implementation: `pkg/go/validation/cycle_shape.go`

The weighted graph refuses both shapes, with `ErrModelCycle` and `ErrTupleCycle`, and names
no relation in either. The check reads the model to find the relations, and the graph stays
the authority on whether a model is resolvable: it may only report relations in a model the
builder refuses.
