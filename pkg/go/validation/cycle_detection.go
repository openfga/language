package validation

import (
	"maps"
	"slices"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// entryPointResult mirrors the reference implementation's { hasEntry, loop }.
type entryPointResult struct {
	hasEntry bool
	loop     bool
}

// validateEntryPoints reports relations that have no entry point. Such a
// relation is impossible: either it never reaches a concrete assignable type
// (no entrypoint) or it forms a rewrite loop (potential loop).
func validateEntryPoints(idx *index, src source) Findings {
	var fs Findings

	for _, typeDef := range idx.model.GetTypeDefinitions() {
		relations := typeDef.GetRelations()
		if len(relations) == 0 {
			continue
		}

		typeName := typeDef.GetType()
		typeLine := src.typeLine(typeName)

		for _, relationName := range slices.Sorted(maps.Keys(relations)) {
			result := hasEntryPointOrLoop(idx, typeName, relationName, relations[relationName],
				map[string]map[string]bool{})
			if result.hasEntry {
				continue
			}

			file, module := relationMeta(typeDef, relationName)
			line := src.relationLine(relationName, typeLine)

			finding := noEntryPoint(relationName, typeName)
			if result.loop {
				finding = noEntryPointLoop(relationName, typeName)
			}

			fs = append(fs, finding.at(src, line).in(file, module))
		}
	}

	return fs
}

// hasEntryPointOrLoop determines whether a rewrite reaches a concrete entry
// point. The visited map tracks type#relation pairs already on the current
// traversal.
//
// It is a port of the reference implementation's hasEntryPointOrLoop
// (pkg/js/validator/validate-dsl.ts): a single traversal per relation that
// yields exactly one outcome, rather than separate cycle and entry-point
// passes.
//
// Only the computed-userset branch turns a revisit into a reported loop. The
// direct type-relation and tuple-to-userset branches skip a reference already
// being resolved and answer loop: false, matching validate-dsl.ts, which reads
// hasEntry off those two recursive calls and discards their loop.
//
// Sibling branches (the this/ttu type loops, union/intersection children, and a
// difference's base/subtract) each get an isolated copy of visited so one
// branch's path can't poison another's loop check. The lone linear successor —
// the computed-userset tail call — shares visited directly: it accumulates down
// the chain to detect back-edges, and avoids an O(n²) copy on deep chains.
func hasEntryPointOrLoop(idx *index, typeName, relationName string,
	rewrite *openfgav1.Userset, visited map[string]map[string]bool) entryPointResult {
	if relationName == "" || rewrite == nil {
		return entryPointResult{}
	}

	if visited[typeName] == nil {
		visited[typeName] = map[string]bool{}
	}

	visited[typeName][relationName] = true

	if !idx.relationDefined(typeName, relationName) {
		return entryPointResult{}
	}

	switch rewrite.GetUserset().(type) {
	case *openfgav1.Userset_This:
		// A direct assignment has an entry point if any assignable type is a
		// concrete type or wildcard. A type#relation restriction only provides
		// an entry point if that referenced relation itself has one.
		for _, restriction := range idx.directTypeRestrictions(typeName, relationName) {
			restrictedType := restriction.GetType()
			restrictedRelation := restriction.GetRelation()

			if restrictedRelation == "" || restriction.GetWildcard() != nil {
				return entryPointResult{hasEntry: true}
			}

			assignable := idx.userset(restrictedType, restrictedRelation)
			if assignable == nil {
				// Matches validate-dsl.ts: returns on the first missing
				// reference rather than trying later types. Unreachable in
				// practice (the reference pass + cascade gate run first).
				return entryPointResult{}
			}

			if visited[restrictedType][restrictedRelation] {
				continue
			}

			if hasEntryPointOrLoop(idx, restrictedType, restrictedRelation, assignable, copyVisited(visited)).hasEntry {
				return entryPointResult{hasEntry: true}
			}
		}

		return entryPointResult{}

	case *openfgav1.Userset_ComputedUserset:
		computed := rewrite.GetComputedUserset().GetRelation()
		if computed == "" || !idx.relationDefined(typeName, computed) {
			return entryPointResult{}
		}

		if visited[typeName][computed] {
			return entryPointResult{loop: true}
		}

		// Linear successor: share visited so the chain accumulates (see above).
		return hasEntryPointOrLoop(idx, typeName, computed, idx.userset(typeName, computed), visited)

	case *openfgav1.Userset_TupleToUserset:
		ttu := rewrite.GetTupleToUserset()
		tupleset := ttu.GetTupleset().GetRelation()
		computed := ttu.GetComputedUserset().GetRelation()

		if tupleset == "" || computed == "" || !idx.relationDefined(typeName, tupleset) {
			return entryPointResult{}
		}

		for _, restriction := range idx.directTypeRestrictions(typeName, tupleset) {
			assignableType := restriction.GetType()

			assignable := idx.userset(assignableType, computed)
			if assignable == nil {
				continue
			}

			if visited[assignableType][computed] {
				continue
			}

			if hasEntryPointOrLoop(idx, assignableType, computed, assignable, copyVisited(visited)).hasEntry {
				return entryPointResult{hasEntry: true}
			}
		}

		return entryPointResult{}

	case *openfgav1.Userset_Union:
		hasLoop := false

		for _, child := range rewrite.GetUnion().GetChild() {
			result := hasEntryPointOrLoop(idx, typeName, relationName, child, copyVisited(visited))
			if result.hasEntry {
				return entryPointResult{hasEntry: true}
			}

			hasLoop = hasLoop || result.loop
		}

		return entryPointResult{loop: hasLoop}

	case *openfgav1.Userset_Intersection:
		for _, child := range rewrite.GetIntersection().GetChild() {
			result := hasEntryPointOrLoop(idx, typeName, relationName, child, copyVisited(visited))
			if !result.hasEntry {
				return entryPointResult{loop: result.loop}
			}
		}

		return entryPointResult{hasEntry: true}

	case *openfgav1.Userset_Difference:
		diff := rewrite.GetDifference()

		base := hasEntryPointOrLoop(idx, typeName, relationName, diff.GetBase(), copyVisited(visited))
		if !base.hasEntry {
			return entryPointResult{loop: base.loop}
		}

		subtract := hasEntryPointOrLoop(idx, typeName, relationName, diff.GetSubtract(), copyVisited(visited))
		if !subtract.hasEntry {
			return entryPointResult{loop: subtract.loop}
		}

		return entryPointResult{hasEntry: true}
	}

	return entryPointResult{}
}

// copyVisited deep-copies the visited map so sibling branches don't share state.
func copyVisited(src map[string]map[string]bool) map[string]map[string]bool {
	dst := make(map[string]map[string]bool, len(src))

	for typeName, relations := range src {
		inner := make(map[string]bool, len(relations))
		for relationName, visited := range relations {
			inner[relationName] = visited
		}

		dst[typeName] = inner
	}

	return dst
}
