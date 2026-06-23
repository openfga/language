# `pkg/go` Maintainer Notes — Gotchas & Flows

> Onboarding notes for a new maintainer. Focus is `pkg/go`, especially the **graph** package.
> This is a working map of how the code fits together and the traps that aren't obvious from reading a single file. It is not a substitute for the code; line references are accurate as of the time of writing but drift, so re-grep before relying on a specific line.

---

## 1. What this repo is

`github.com/openfga/language` is a **multi-language** toolkit for the OpenFGA DSL (the authorization-model language). The single source of truth is the ANTLR grammar at the repo root:

- `OpenFGALexer.g4`, `OpenFGAParser.g4` — grammar. **Change these and you must regenerate parsers for all three languages.**
- `pkg/go`, `pkg/js`, `pkg/java` — three independent implementations that must stay behaviour-compatible.
- `tests/data/` — **shared** YAML/JSON fixtures consumed by all three implementations. This is the cross-language contract.

Each `pkg/<lang>` is released independently (see §9).

### Feature parity matrix (the important asymmetry)

| Feature | Go | JS | Java |
|---|---|---|---|
| DSL↔JSON transformer | ✓ | ✓ | ✓ |
| Syntactic validation | ✓ | ✓ | ✓ |
| Semantic validation | ✗ (issue #99) | ✓ | ✓ |
| **Graph utilities** | **✓ (Go ONLY)** | ✗ | ✗ |

**The graph package is Go-only and has no sibling implementation to cross-check against.** That is the single most important fact for you. There is no JS/Java reference for graph behaviour — the Go tests *are* the spec.

---

## 2. The graph package is load-bearing for the OpenFGA server

`pkg/go/graph/` is consumed directly by `github.com/openfga/openfga` (the main server) for its **typesystem, Check, and ListObjects** algorithms. Per `AGENTS.md`:

- Breaking changes here can break core authorization in production.
- Always run `go test ./graph/... -v` and consider how the change lands in the server before merging.
- Treat the exported surface (`Get*`, `Weighted*`, node/edge types) as a **public API**. Renames/signature changes are breaking even when "just a cleanup."

There is already a concrete example of API-compat care in the code: `GetEdgesFromNodeId` is kept as a `// Deprecated:` shim that forwards to `GetEdgesFromNodeID` (`weighted_graph.go:40`). Follow that pattern — don't delete exported methods, deprecate and forward.

> **Commit `refactor(pkg/go/...)!` ⇒ breaking.** The `!` and a `BREAKING CHANGE:` footer drive a major (or pre-1.0 minor) bump via release-please. See git history e.g. `7fd286c refactor(pkg/go/transformer)!`.

---

## 3. Two different "graphs" — don't confuse them

There are **two parallel, mostly independent** graph implementations in the same package. This is the #1 source of confusion.

### A. `AuthorizationModelGraph` ("plain graph") — `graph.go`, `graph_node.go`, `graph_edge.go`, `graph_builder.go`

- Built with **gonum** (`gonum.org/v1/gonum/graph/multi.DirectedGraph`). It embeds a gonum multigraph and leans on gonum for topology (`topo.PathExistsIn`, `topo.DirectedCyclesIn`).
- Entry point: `NewAuthorizationModelGraph(model) (*AuthorizationModelGraph, error)`.
- Purpose: structural reachability + visualization. Key methods: `PathExists`, `Reversed`, `GetCycles`, `GetDOT`, `GetNodeByLabel`.
- Node lookup is by **label** through an `ids map[string]int64` (`NodeLabelsToIDs`). gonum assigns numeric IDs; this map is how you go label→id→node in O(1).
- **No weights.** Conditions are tracked on edges but "Conditions are not encoded in the graph" for traversal purposes.

### B. `WeightedAuthorizationModelGraph` ("weighted graph") — `weighted_graph.go`, `weighted_graph_node.go`, `weighted_graph_edge.go`, `weighted_graph_builder.go`

- **Does NOT use gonum for storage.** It is hand-rolled: two plain maps,
  ```go
  nodes map[string]*WeightedAuthorizationModelNode  // keyed by uniqueLabel
  edges map[string][]*WeightedAuthorizationModelEdge // keyed by from-node uniqueLabel (adjacency list)
  ```
  (The *builder* `WeightedAuthorizationModelGraphBuilder` embeds a gonum multigraph, but the resulting `WeightedAuthorizationModelGraph` it produces uses the maps. Don't be misled by the embed.)
- Entry point: `NewWeightedAuthorizationModelGraphBuilder().Build(model)`.
- Purpose: the **weight** algorithm that the server uses to plan Check/ListObjects (which path is cheapest, recursion detection, etc.).
- This is where ~80% of the package's complexity and line count lives (`weighted_graph.go` ~1250 lines, `weighted_graph_builder.go`, plus ~7000 lines of tests).

**Gotcha:** the two share `NodeType`, `EdgeType`, operator-name constants, and `DrawingDirection`, but **not** node/edge structs, not the builder, not the storage model. A fix in one is *not* automatically reflected in the other. When you change graph-building logic, check whether the same logic exists (duplicated) in both `graph_builder.go` and `weighted_graph_builder.go` — the parsing functions (`parseThis`, `parseComputed`, `parseTupleToUserset`, `checkRewrite`/`parseRewrite`) are near-duplicates with subtle differences.

---

## 4. Graph vocabulary (NodeType / EdgeType)

Defined in `weighted_graph_node.go` and `weighted_graph_edge.go`. You will read these constantly.

**NodeType:**
| Const | Meaning | Example |
|---|---|---|
| `SpecificType` | a type | `group` |
| `SpecificTypeAndRelation` | a relation on a type | `group#member` |
| `OperatorNode` | union/intersection/exclusion | label is `"union"`/`"intersection"`/`"exclusion"` |
| `SpecificTypeWildcard` | typed wildcard | `group:*` |
| `LogicalDirectGrouping` | grouping of multiple direct assignments under an operator | `[user, employee, type1#rel]` |
| `LogicalTTUGrouping` | grouping of a TTU with multiple possible parent types | `member from parent` where `parent` has several types |

**EdgeType:**
| Const | Meaning |
|---|---|
| `DirectEdge` | direct assignment `define rel: [user]` |
| `RewriteEdge` | operator→relation, or relation→operator wiring |
| `TTUEdge` | tuple-to-userset `define rel: admin from parent` |
| `ComputedEdge` | `define rel1: rel2` (relation→relation) |
| `DirectLogicalEdge` | operator→`LogicalDirectGrouping` |
| `TTULogicalEdge` | operator→`LogicalTTUGrouping` |

### Operator-node uniqueness gotcha — ULIDs
Operator nodes get a **unique label built from a fresh ULID**: `fmt.Sprintf("%s:%s", operator, ulid.Make().String())` (`graph_builder.go:109`, `weighted_graph_builder.go:93`). So `union:01J54...`. This means:
- `label` (e.g. `"union"`) is *not* unique; `uniqueLabel` is.
- **`ulid.Make()` is nondeterministic** (it embeds a timestamp + randomness). The graph is deliberately built with **types and relations sorted** (`slices.SortFunc` by type name, `slices.Sort` on relations) so that *structure* is stable, but the ULID strings on operator nodes are NOT stable across runs. **Never assert on the exact ULID in a test or snapshot** — match on `label`/structure instead. `GetDOT()` output is "stable" only in structure, and the doc comment says it's for debugging only.

### `LogicalDirectGrouping` / `LogicalTTUGrouping` — the "only when >1 child under an operator" rule
These grouping nodes are created **only** when a direct-assignment list or a multi-parent TTU appears as a child of an operator node AND there is more than one related type. Read the conditions carefully:

- `parseThis` (weighted, `weighted_graph_builder.go:185`): grouping node created only if `parentNode.nodeType != SpecificTypeAndRelation && len(directlyRelated) > 1`. I.e. directly on a relation (`define rel: [a,b,c]`) → three direct edges straight off the relation node, **no** grouping node. Inside an `or`/`and`/`but not` → grouping node.
- The edge-type comments explicitly note that **under a `not`/exclusion the grouping node is intentionally NOT created** to avoid unnecessary nesting.

If you add/inspect grouping logic, this asymmetry (relation-level vs operator-level, and the not-operator exception) is the thing to keep straight.

---

## 5. The weight algorithm (`weighted_graph.go`) — the deep end

This is the hardest code in the repo. High-level model:

- A node's `weights` is a `map[string]int` from **terminal type** (e.g. `"user"`) to an integer weight (roughly path length / cost), where `Infinite = math.MaxInt32` signals recursion.
- `AssignWeights()` is the driver. It walks nodes, computing edge weights then node weights bottom-up.

### Key invariants & traps

1. **Operator/logical nodes are deferred.** `AssignWeights` *skips* operator and logical nodes in its first pass (`weighted_graph.go:227`, `isLogicalOperator`) and resolves them later in `fixDependantNodesWeight`. The comment says this is "for more deterministic behavior." Don't "optimize" by computing them inline — order matters here.

2. **Three weight-combination strategies, chosen by operator:**
   - `calculateNodeWeightWithMaxStrategy` — used for non-operator nodes and **Union**. A type is valid if it appears in *any* branch → take max.
   - `calculateNodeWeightWithEnforceTypeStrategy` — **Intersection (AND)**. A type is valid only if it appears in *all* branches → intersect keys, drop any key missing from a branch. **If the result is empty → `ErrInvalidModel` ("not all paths return the same type").**
   - `calculateNodeWeightWithMixedStrategy` — **Exclusion (A but not B)**. `A` is max; `B` only constrains types already in `A`. Asserts exactly 2 edges.

3. **Cycle classification is central.** `isTupleCycle` (`weighted_graph.go:546`) distinguishes:
   - **Model cycle** (`ErrModelCycle`): a pure computed/rewrite cycle with no TTU or userset edge → the model is invalid, would stack-overflow at Check time regardless of tuples.
   - **Tuple cycle** (`ErrTupleCycle`): cycle goes through a `TTUEdge` or a `DirectEdge` into a `SpecificTypeAndRelation` → can only cycle *if tuples exist*. This is allowed (weight becomes `Infinite`), and tracked via `recursiveRelation` / `tupleCycle` metadata on nodes+edges.
   - **`ErrContrainstTupleCycle`**: AND/BUT-NOT operands cannot be part of a tuple cycle → invalid model.

4. **`R#` reference keys.** During cycle handling, weights temporarily carry keys prefixed `"R#"` + nodeID (e.g. `R#group#member`) meaning "this weight depends on resolving the cycle rooted at that node." `calculateNodeWeightAndFixDependencies` + `fixDependantEdgesWeight` + `fixDependantNodesWeight` later substitute the real weights and strip the `R#` placeholders. If you ever see `R#...` leaking into a final node's `weights`, the dependency-fix pass didn't run / a cycle wasn't rooted correctly. That's a bug, not expected output.

5. **`tupleCycleDependencies map[string][]*edge`** is the bookkeeping that lets the algorithm fix all edges/nodes that pointed at a cycle once the cycle's root is resolved. The "responsible node" (`isNodeTupleCycleReference`) is the one that fixes and then `removeNodeFromTupleCycles`.

6. **Userset weights are computed lazily and cached in `sync.Map`.** `getWeightForUserset` / `GetEdgeWeight` compute weights for `type#relation` keys on demand and cache them in `node.usersetWeights` / `edge.usersetWeights` (both `sync.Map`). This is why those fields are `sync.Map` and not plain maps — see §6 on concurrency. There's an "early pruning" check before traversal (`weighted_graph.go:333`) using the rule that a direct edge adds +1, so a reachable userset must have weight strictly greater (unless `Infinite`).

7. **`directAssigns` on a node** is a flat list of the unique labels a relation can be directly assigned (filled in `parseThis`). It exists so the server can answer "is this write/contextual tuple valid?" in O(1). `GetDirectEdgesAssignation` / `GetDirectEdgeForUserType` reconstruct the actual edges, which is non-trivial because direct edges may sit under a `LogicalDirectGrouping` or deeper under operator nodes (hence the traversal loops).

**Practical advice:** before touching the weight algorithm, read the tests first (`weighted_graph_test.go`, `weighted_graph_builder_test.go`, `weighted_graph_userset_test.go`). They encode the intended weights for dozens of model shapes and are far more readable than the algorithm itself. Add a failing test that captures your case, *then* change code.

---

## 6. Concurrency gotcha (graph)

`WeightedAuthorizationModelNode` and `WeightedAuthorizationModelEdge` contain `usersetWeights sync.Map`. The server calls weight queries concurrently, so:

- **Read-only query methods may still mutate** (they populate the `sync.Map` cache). That's intentional and the `sync.Map` makes it safe.
- **The maps `nodes`/`edges` themselves are NOT synchronized.** They are populated during `Build`/`AssignWeights` and then treated as immutable. Do not add code that writes to `wg.nodes` / `wg.edges` after build — that would race with concurrent readers.
- There is a dedicated `weighted_graph_concurrent_test.go` (`TestConcurrentUsersetWeights`). **Run the whole suite with `-race`** (the Makefile `test` target already passes `-race`). If you add caching or memoization, add/extend a concurrent test.

---

## 7. The transformer package — entry points & flows

`pkg/go/transformer/`. Public API (grep `^func [A-Z]`):

DSL → JSON/proto (`dsltojson.go`):
- `TransformDSLToProto(data) (*openfgav1.AuthorizationModel, error)` — the main one.
- `MustTransformDSLToProto`, `TransformDSLToJSON`, `MustTransformDSLToJSON`.
- `TransformModularDSLToProto` — modular models (returns the model + a map of type-def extensions).
- `ParseDSL(data) (*OpenFgaDslListener, *OpenFgaDslErrorListener)` — lower-level; gives you the ANTLR listener + error listener.

JSON/proto → DSL (`jsontodsl.go`):
- `TransformJSONProtoToDSL(model, opts...)`, `TransformJSONStringToDSL(string, opts...)`, `LoadJSONStringToProto`.
- Options pattern: `WithIncludeSourceInformation(bool)` (`TransformOption`).

Modules (schema 1.2 / modular models):
- `mod-to-json.go`: `TransformModFile(data) (*ModFile, error)` — parses an `fga.mod` file.
- `module-to-model.go`: `TransformModuleFilesToModel([]ModuleFile, schemaVersion)` — stitches multiple `.fga` module files into one model.

### Transformer gotchas

- **ANTLR listener pattern, mutable state.** `OpenFgaDslListener` accumulates state (`currentTypeDef`, `currentRelation`, `rewriteStack`, etc.) as ANTLR walks the parse tree. It is **not reusable / not concurrency-safe** — `newOpenFgaDslListener()` per parse. The `rewriteStack` is how nested `or`/`and`/`but not` expressions are assembled (`ParseExpression`).
- **Proto, not hand-rolled JSON.** Models are `openfgav1.AuthorizationModel` from `github.com/openfga/api/proto`. JSON conversion goes through `protojson`, not `encoding/json`. If a field "disappears" in JSON, suspect proto field names / `protojson` casing, not your code.
- **Errors are user-facing and line/column-aware.** Syntax errors come from `OpenFgaDslErrorListener` and are matched character-for-character by the shared fixtures (see §8). Changing an error string is a cross-language contract change.
- **`MustTransform*` panic.** Only use in tests / where input is known-good.

---

## 8. Shared test fixtures — the cross-language contract

`tests/data/` (at repo root, three levels up from `pkg/go/transformer`) holds YAML/JSON cases used by Go, JS, and Java. The Go side loads them with relative paths like `filepath.Join("../../../tests", "data", "transformer")` (`transformer/testcases_test.go:70`).

Files you'll touch:
- `transformer/`, `transformer-module/` — DSL↔JSON case dirs.
- `dsl-syntax-validation-cases.yaml`, `json-syntax-transformer-validation-cases.yaml`, `dsl-semantic-validation-cases.yaml`, `json-validation-cases.yaml`, `fga-mod-transformer-cases.yaml`.
- `stores/` — store-level fixtures.

**Gotchas:**
- A test case has a `Skip` flag (`validTestCase.Skip`) — used to land a fixture before all three languages implement it. If you add a case Go can't yet handle but JS can, you may need to coordinate skips across languages.
- These fixtures are linted with **prettier** (`make lint-tests`, run from `pkg/js`). A fixture-only change can fail CI on formatting. Run `make format-tests` before pushing.
- **Adding a fixture affects JS and Java CI too.** A "Go change" that edits shared fixtures is really a three-language change. If you only intend to change Go behaviour, prefer a Go-only test (e.g. in `graph/`) over editing shared fixtures.
- Error-case fixtures match the *exact* error string and line/column. See `GetErrorString()` in `testcases_test.go` for the format the Go side expects (`N error(s) occurred:\n\t* syntax error at line=L, column=C: msg`).

---

## 9. Build, generate, lint, release

### ANTLR generation needs Docker
- `make antlr-gen-go` (and `-js`, `-java`) run an **ANTLR Docker container** (`make build-antlr-container`). You cannot regenerate parsers without Docker.
- Generated code lives in `pkg/go/gen/` (`openfga_parser.go`, `openfga_lexer.go`, listeners, `.tokens`, `.interp`). **Never hand-edit `gen/`.** Edit the `.g4` files at the repo root, then regenerate.
- Most `make` targets depend on `antlr-gen-*` first, so a normal `make build-go` will try to invoke Docker. If you only changed Go source (not grammar), you can work inside `pkg/go` directly with plain `go` commands to skip Docker (see below).

### Day-to-day Go commands (from `pkg/go/`, no Docker needed)
```bash
go build ./...
go test ./... -count=1 -race          # == make test (note: -race always)
go test ./graph/... -v                # graph only
go test ./transformer/... -run TestDslToJson -v
```
Or via root Makefile (these prepend antlr-gen → Docker):
- `make test-go`, `make lint-go` (golangci-lint **with --fix**), `make audit-go` (govulncheck), `make format-go` (gofumpt — only formats `transformer/` and `errors/`, per the Makefile).

**Gotcha:** `make lint-go` runs `golangci-lint --fix`, which **modifies your files**. Don't be surprised by working-tree changes after linting. Config is `.golangci.yaml`.

**Gotcha:** `go.mod` is `go 1.25.0`. The module path is `github.com/openfga/language/pkg/go` — note the package is *inside* `pkg/go`, so import paths are `.../pkg/go/graph`, `.../pkg/go/transformer`.

### Releases — release-please, per-package
- `release-please-config.json` + `.release-please-manifest.json` drive **independent** releases per package. Current: `pkg/go 0.3.0`, `pkg/js 0.2.1`, `pkg/java 0.2.0-beta.2`.
- Tags use `include-component-in-tag` + `tag-separator: "/"` ⇒ Go tags look like `pkg/go/v0.3.0`.
- Conventional-commit type → changelog section is configured (`feat`→Added, `fix`→Fixed, `perf`/`refactor`→Changed, etc.). `chore`/`ci`/`test`/`release` are hidden.
- **Pre-1.0 bumping:** `bump-minor-pre-major` + `bump-patch-for-minor-pre-major` mean a `feat` bumps the *patch* (not minor) while under 1.0, and a breaking change bumps the *minor*. Plan version expectations accordingly.
- There is a `pkg/go/.release-please-trigger` file used to force/trigger a release.

### Commit convention (enforced)
Conventional Commits: `feat|fix|docs|chore|test|refactor|perf|build|ci|style`. Use `!` / `BREAKING CHANGE:` for breaking. Scope helps: `feat(pkg/go/graph): ...`.
**(Repo/house rule already in your memory: no `Co-Authored-By`/Claude attribution lines in commits.)**

---

## 10. `DrawingDirection` — easy to get backwards

`type DrawingDirection bool` (`graph.go:19`):
- `DrawingDirectionListObjects = true` — terminal types have **outgoing** edges, no incoming (bottom-up). This is the **default** for `NewAuthorizationModelGraph` (`graph_builder.go:32`).
- `DrawingDirectionCheck = false` — terminal types have **incoming** edges, no outgoing (top-down). This is the default for the **weighted** builder (`weighted_graph_builder.go:20`).

So the two graphs are built pointing in **opposite directions by default.** `Reversed()` flips the whole graph and negates the direction (`!g.drawingDirection`). DOT `rankdir` is derived from it (`BT` for ListObjects, `TB` for Check). When you reason about "from/to" on an edge, always confirm which graph and which direction you're in.

---

## 11. Smaller packages

- `pkg/go/validation/validation-rules.go` — pure **regex** validators for identifiers (type/relation/condition/object/userset/wildcard). Stateless, simple. The regex rules (`RuleType`, `RuleRelation`, etc.) define the legal identifier charset; they must match the grammar and the other languages.
- `pkg/go/utils/model_utils.go` — `GetModuleForObjectTypeRelation` (module resolution for modular models, falls back type-module → relation-module), `IsRelationAssignable` (does a rewrite contain `this`? recursive over union/intersection/difference).
- `pkg/go/utils/line-numbers.go` — line-number helpers for error reporting.
- `pkg/go/errors/errors.go` — a couple of formatted error constructors; note `//nolint:goerr113` because they intentionally build dynamic errors. `ErrUnsupportedJSON` and the DSL-nesting error tie to the "Nesting: partial" caveat (issue #113) in the README.

---

## 12. Known limitations / planned work (context for "what to build")

- **Semantic validation is NOT implemented in Go** (issue #99) — only syntactic. JS/Java have it. This is a likely parity gap to be asked about.
- **Nesting is partial** (issue #113) — some deeply-nested rewrites aren't expressible in DSL; `UnsupportedDSLNestingError` is raised.
- **Graph utilities are Go-only and explicitly "planned" for JS/Java** per the README parity table — but today they live only here.
- `graph.go` ends with `// TODO add graph traversals, etc.` and `GetCycles` has `// TODO: investigate whether len(1) should be identified as cycle`. `graph_builder.go:298` notes `typeAndRelationExists` is O(n) and "should be made faster, ideally typeDefs is a map." These are sanctioned starting points if you're looking for graph work.

---

## 13. Checklist before merging a `graph` change

1. `cd pkg/go && go test ./graph/... -v -race` green.
2. Did you touch logic that is **duplicated** in both `graph_builder.go` and `weighted_graph_builder.go`? Update both or document why not.
3. Did you change an **exported** symbol? If so it's a breaking change → `!` commit, deprecate-don't-delete where possible.
4. Did you change weight/cycle semantics? Add a model-shaped test and reason about the impact on `github.com/openfga/openfga` (typesystem/Check/ListObjects).
5. No `ulid`/timestamp/map-iteration assumptions leaked into deterministic output.
6. Ran `-race` (covers the `sync.Map` userset caches).
7. If you edited `tests/data/`, ran `make format-tests` and remembered it affects JS + Java CI.
8. Commit message follows Conventional Commits, correct scope, no attribution footer.
