# The Weighted Authorization Graph

*The invariants of the graph that governs authorization*

Consider a directed multigraph `G = (V, E)`. It admits loops, so an edge may return to its own vertex, and it admits parallel edges, so several edges may join the same pair of vertices. The standard properties of such graphs are assumed; only the axioms specific to the present construction are set down here.

Every vertex has exactly one kind, drawn from `{SpecificType, SpecificTypeAndRelation, OperatorNode, SpecificTypeWildcard, LogicalDirectGrouping, LogicalTTUGrouping}`.

Every edge has exactly one kind, drawn from `{DirectEdge, RewriteEdge, TTUEdge, ComputedEdge, DirectLogicalEdge, TTULogicalEdge}`.

## Definitions

- A *sink* is a vertex with no outgoing edge.
- The *weight* is a function _w(v, s)_ that takes exactly one vertex _v_ and exactly one sink _s_ as inputs and outputs a single value. The output value must strictly belong to the set of non-negative integers or infinity: {0, 1, 2, ...} ⋃ {∞}.
- A sink _s_ is *reachable* from a vertex _v_ if a path leads from _v_ to _s_.
- An edge *resolves a tuple* if it is a `DirectEdge` or a `TTUEdge`.
- A *recursive* cycle is a directed cycle whose vertices and edges are all confined to a single relation.
- A *tuple cycle* is a directed cycle that spans multiple distinct relations, or in which a single relation recurses through more than one branch.

## Axioms

### 1. Structure

1. Each vertex carries a label unique to it, determined entirely by the model.
2. A vertex has an outdegree of zero if and only if its kind is `SpecificType` or `SpecificTypeWildcard`.
3. No edge is without a condition; where none is imposed, a placeholder token stands in its place.
4. For each represented combination of source, target, kind, and tupleset relation, exactly one edge exists. Its condition set is the accumulation of all specified conditions for that combination.
5. An edge has a non-empty tupleset relation if and only if its kind is `TTUEdge` or `TTULogicalEdge`.
6. Every `OperatorNode` has an indegree of exactly one, and its incoming edge must be a `RewriteEdge`.
7. Every `OperatorNode` vertex represents exactly one of three specific operations: union, intersection, or exclusion.
8. An exclusion `OperatorNode` vertex has an outdegree of exactly two. These outgoing edges form an ordered sequence consisting of a _base_ edge followed by a _subtract_ edge.
9. A `DirectEdge` can only originate from a `SpecificTypeAndRelation` or `LogicalDirectGrouping` vertex, and can only terminate at a `SpecificType`, `SpecificTypeWildcard`, or `SpecificTypeAndRelation` vertex.
10. A `ComputedEdge` must originate from a `SpecificTypeAndRelation` vertex and terminate at a `SpecificTypeAndRelation` vertex.
11. A `RewriteEdge` can only originate from a `SpecificTypeAndRelation` or `OperatorNode` vertex, and can only terminate at an `OperatorNode` or `SpecificTypeAndRelation` vertex.
12. A `TTUEdge` can only originate from a `SpecificTypeAndRelation` or `LogicalTTUGrouping` vertex, and can only terminate at a `SpecificTypeAndRelation` vertex.
13. A `DirectLogicalEdge` must originate from an `OperatorNode` vertex and terminate at a `LogicalDirectGrouping` vertex.
14. A `TTULogicalEdge` must originate from an `OperatorNode` vertex and terminate at a `LogicalTTUGrouping` vertex.
15. Every `LogicalDirectGrouping` or `LogicalTTUGrouping` vertex must have an outdegree of two or more.
16. A `LogicalDirectGrouping` or `LogicalTTUGrouping` vertex exists if and only if a direct set or a tuple-to-userset expands to more than one target beneath an operator.

### 2. Weight

1. A path exists from a vertex to a sink if and only if the weight between them is greater than zero.
2. If a path uses an edge that resolves a tuple, then the required further tuple must be found before the edge's target vertex is reached.
3. For any edge connecting a source vertex to a target vertex, the weight from the source to a given sink is never less than the weight from the target to that same sink. Furthermore, if the edge resolves a tuple, the weight from the source must be strictly greater than the finite weight from the target.
4. A vertex has a weight of ∞ to a sink if and only if its reachability relies on a cyclic path.
5. Excluding intersection and exclusion nodes, the weight of a non-sink vertex to a given sink is the maximum of the weights carried by its outgoing edges to that sink.
6. The weight of an intersection vertex to a given sink is exactly zero if any of its outgoing edges cannot reach that sink; otherwise, its weight is equal to the maximum weight carried by its outgoing edges toward that sink.
7. A sink is reachable from an exclusion vertex if and only if the base edge reaches it; if the base edge fails, the final weight is zero. When the base edge succeeds, the final weight is the maximum value between the base edge's weight and the subtract edge's weight.
8. Every non-sink vertex must reach at least one sink.

### 3. Cycles

1. Every directed cycle in the graph must contain at least one tuple-resolving edge.
2. No recursive or tuple cycle may pass through an intersection or exclusion vertex.
