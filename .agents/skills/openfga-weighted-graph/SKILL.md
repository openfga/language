---
name: openfga-weighted-graph
description: Analyze whether an OpenFGA authorization model can be loaded by the weighted graph (pkg/go/graph WeightedAuthorizationModelGraphBuilder), explain each incompatibility, propose model changes, and plan the tuple/data migration. Use when a model is valid for OpenFGA but the weighted graph build fails, when you see errors like "type does not have defined X relation", "not all paths return the same type", "operands AND or BUT NOT cannot be involved in a cycle", or "tuple cycle", or when OpenFGA Check/ListObjects is unexpectedly slow because the weighted-graph optimizations are disabled for a model.
---

# OpenFGA weighted graph compatibility

The OpenFGA server builds a weighted graph for every model and **silently discards build errors** (`typesystem.attachGraphs`). The model is still accepted and answers stay correct, but every weighted-graph optimization (weight-2 Check resolvers, optimized ListObjects) is turned off for the whole model. So "OpenFGA accepts it" does not mean "the weighted graph accepts it".

## Workflow

1. **Run the checker** (from the repo root; pass an absolute path to a `.fga`, `.json` or `fga.mod` model):
   ```bash
   go -C .agents/skills/openfga-weighted-graph/scripts/wgcheck run . /abs/path/model.fga
   ```
   It validates with OpenFGA's `typesystem.NewAndValidate`, then builds the weighted graph using this repo's `pkg/go` (via a `replace` in its go.mod). It patches each finding in memory and rebuilds, so a single run lists **all** incompatibilities, tagged `[A]`, `[B]`, `[C]`. Exit code 0 means compatible.
   - If OpenFGA rejects the model, stop: fix that first.
   - `[?]` means an error not covered below; read `pkg/go/graph/weighted_graph.go` / `weighted_graph_builder.go` for the error text.
2. **Read the model** and, for each finding, pick a fix from the catalog below. Prefer fixes that need no tuple migration and no application change.
3. **Prove the effective permissions are unchanged** (or document exactly what changes). Reason about every path that reaches the changed relation, then write a `.fga.yaml` test (see [Verifying equivalence](#verifying-equivalence)) and run it against both the old and new model.
4. **Re-run the checker** on the new model until it exits 0.
5. **Produce the migration plan** (see [Migration playbook](#migration-playbook)) if tuples or Check targets change.

## Incompatibility catalog

### [A] Tuple-to-userset where some parent types lack the relation

```
model
  schema 1.1
type user
type organization
  relations
    define admin: [user]                                  # no "viewer"
type folder
  relations
    define admin: [user]
    define viewer: [user] or admin
type document
  relations
    define parent: [folder, organization]
    define owner: [user] or admin from parent
    define viewer: [user] or owner or viewer from parent  # OpenFGA: OK if ANY parent type has viewer
```
The weighted graph requires **every** type in the tupleset (`parent`) to define the computed relation.

| Fix | Model change | Data migration |
|---|---|---|
| A1. Add the relation to the missing type as directly assignable and never write tuples to it | Add relation | None |
| A2. Split the tupleset per type | Rename/split tupleset | Rewrite `document:x#parent@folder:y` → `document:x#parent_folder@folder:y` and `document:x#parent@organization:y` → `document:x#parent_organization@organization:y` |

A1 is the equivalent fix. In the original model, `viewer from parent` with an `organization` parent always resolves to nothing, and an empty `organization#viewer` keeps it that way:
```
model
  schema 1.1
type user
type organization
  relations
    define admin: [user]
    define viewer: [user]       # never written: keeps 'viewer from parent' empty for organizations
type folder
  relations
    define admin: [user]
    define viewer: [user] or admin
type document
  relations
    define parent: [folder, organization]
    define owner: [user] or admin from parent
    define viewer: [user] or owner or viewer from parent
```

A2:
```
model
  schema 1.1
type user
type organization
  relations
    define admin: [user]
type folder
  relations
    define admin: [user]
    define viewer: [user] or admin
type document
  relations
    define parent_folder: [folder]
    define parent_organization: [organization]
    define owner: [user] or admin from parent_folder or admin from parent_organization
    define viewer: [user] or owner or viewer from parent_folder
```

Prefer A1. Do not define the missing relation in terms of existing ones (e.g. `define viewer: admin`): any non-empty definition grants access the original model never granted, even if it happens to be redundant with another path today. A1's only risk is that the relation is writable; make sure the application never writes `organization#viewer` tuples. Use A2 when the tupleset split is wanted anyway.

### [B] Intersection whose operands share no user type

```
model
  schema 1.1
type user
type organization
  relations
    define member: [user]
type document
  relations
    define editor: [user]
    define publisher: [organization]
    define can_publish: editor and publisher  # user vs organization: always false
```
The weighted graph reports "not all paths return the same type". Usually a modeling bug; find what was intended:

| Fix | Model change | Data migration |
|---|---|---|
| B1. Follow the stored organization to its members | Rewrite `can_publish` | None (existing `publisher` tuples become the tupleset) |
| B2. Store usersets instead of objects | Change `publisher` type restriction | Rewrite `document:x#publisher@organization:y` → `document:x#publisher@organization:y#member` |
| B3. Remove `can_publish` if unused | Delete relation | Remove callers |

B1:
```
model
  schema 1.1
type user
type organization
  relations
    define member: [user]
type document
  relations
    define editor: [user]
    define publisher: [organization]
    define can_publish: editor and member from publisher
```

B2:
```
model
  schema 1.1
type user
type organization
  relations
    define member: [user]
type document
  relations
    define editor: [user]
    define publisher: [organization#member]
    define can_publish: editor and publisher
```

### [C] `and` / `but not` inside a recursive cycle

The weighted graph only allows cycles made of `or`, direct usersets and TTUs. Fix: keep the recursive relation union-only (this is where tuples live) and apply the constraint in a new **permission** relation that nothing references recursively.

C1, cycle through exclusion via a userset:
```
model
  schema 1.1
type user
type document
  relations
    define viewer: [user] but not restricted
    define restricted: [user, document#viewer]
```
Fixed:
```
model
  schema 1.1
type user
type document
  relations
    define viewer: [user]
    define restricted: [user, document#viewer]
    define can_view: viewer but not restricted
```

C2, cycle through intersection:
```
model
  schema 1.1
type user
type document
  relations
    define allowed: [user]
    define viewer: [user, document#viewer] and allowed
```
Fixed:
```
model
  schema 1.1
type user
type document
  relations
    define allowed: [user]
    define viewer: [user, document#viewer]
    define can_view: viewer and allowed
```

C3, cycle through a TTU inside an exclusion (inherited folder access minus blocked users):
```
model
  schema 1.1
type user
type folder
  relations
    define parent: [folder]
    define blocked: [user] or blocked from parent
    define viewer: ([user] or viewer from parent) but not blocked
```
Fixed:
```
model
  schema 1.1
type user
type folder
  relations
    define parent: [folder]
    define blocked: [user] or blocked from parent
    define viewer: [user] or viewer from parent
    define can_view: viewer but not blocked
```

- **Data:** no tuple migration; existing tuples keep their relation.
- **Application:** Check/ListObjects/ListUsers must target the new permission (`can_view`) instead of `viewer`.
- **Semantics:** the constraint is applied once, at the object being checked, not at each hop.
  - C1: `document:y#viewer` used as a userset no longer excludes people restricted on `y`.
  - C2: `allowed` is no longer required on intermediate documents.
  - C3: equivalent, because `blocked` is inherited along the same `parent` path, so anyone blocked at an intermediate folder is also blocked at the folder being checked. Look for this pattern: when the subtracted/intersected relation propagates the same way as the recursive one, the fix preserves access.
  - Always call out differences. If per-hop constraints are a hard requirement, the model cannot be made weighted-graph compatible; the user must accept losing the optimizations.

## Verifying equivalence

Write one `.fga.yaml` per model version with **identical** `tuples` and `tests`, differing only in `model_file`, and run both with `fga model test --tests <file>`. Cover each changed relation with:
- a user who should gain nothing new through the changed path (assert `false` where it was `false`),
- a user who keeps access through the existing path (assert `true`),
- `list_objects` for the affected types.

```yaml
model_file: ./model-new.fga
tuples:
  - user: user:anne
    relation: admin
    object: organization:acme
  - user: organization:acme
    relation: parent
    object: document:roadmap
tests:
  - name: organization admin access is unchanged
    check:
      - user: user:anne
        object: document:roadmap
        assertions:
          owner: true
          viewer: true
      - user: user:bob
        object: document:roadmap
        assertions:
          viewer: false
    list_objects:
      - user: user:anne
        type: document
        assertions:
          viewer: [document:roadmap]
```

When semantics intentionally change (catalog C), write tests that pin the new behavior and list the differences for the user.

## Migration playbook

Only needed when tuples move (A2, B2) or callers switch relations (C). Model-only fixes (A1, B1) are a single model write.

1. **Expand:** write a model that contains both old and new relations/permissions.
2. **Dual-write:** the application writes tuples to old and new relations.
3. **Backfill:** `Read` existing tuples page by page and `Write` their new form (batches ≤ the server's max tuples per write; make it idempotent with `on_duplicate: ignore`).
4. **Verify:** run the `.fga.yaml` suites against both models; spot-check production with Check on sampled users.
5. **Switch reads:** point Check/ListObjects/ListUsers at the new relations, pinning the new `authorization_model_id`.
6. **Contract:** write a model without the old relations, stop dual-writing, delete old tuples.
7. Re-run the checker on the final model; it must exit 0.

## Report format

For each finding give: tag and location (`[A] document#viewer`), why the weighted graph rejects it (one line), the recommended fix as a DSL diff, whether effective permissions change, and the migration steps (or "model-only change"). End with the checker output for the proposed model.
