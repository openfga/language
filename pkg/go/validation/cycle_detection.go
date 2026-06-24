package validation

import (
	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// entryPointResult mirrors the reference implementation's { hasEntry, loop }.
type entryPointResult struct {
	hasEntry bool
	loop     bool
}

// CycleDetector walks relation rewrites to determine whether each relation has a
// concrete entry point (a directly-assignable type that is not itself a relation
// reference) or is otherwise impossible — either because it bottoms out with no
// entry point, or because it loops back on itself.
//
// This is a port of the reference implementation's hasEntryPointOrLoop
// (pkg/js/validator/validate-dsl.ts): a single traversal per relation that
// yields exactly one outcome, rather than separate cycle and entry-point passes.
type CycleDetector struct {
	validator *SemanticValidator
}

func NewCycleDetector(validator *SemanticValidator) *CycleDetector {
	return &CycleDetector{validator: validator}
}

// ValidateCyclesAndEntryPoints reports relations that have no entry point. A
// relation with no entry point is impossible: either it never reaches a concrete
// assignable type (no entrypoint) or it forms a rewrite loop (potential loop).
func ValidateCyclesAndEntryPoints(collector *ErrorCollector, model *openfgav1.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}
	validator := NewSemanticValidator(model)
	detector := NewCycleDetector(validator)

	for _, typeDef := range model.GetTypeDefinitions() {
		relations := typeDef.GetRelations()
		if len(relations) == 0 {
			continue
		}
		typeName := typeDef.GetType()
		typeLineIndex := GetTypeLineNumber(typeName, lines, nil)
		for relationName, userset := range relations {
			meta := relationMeta(typeDef, relationName)
			result := detector.hasEntryPointOrLoop(typeName, relationName, userset, map[string]map[string]bool{})
			if !result.hasEntry {
				lineIndex := GetRelationLineNumber(relationName, lines, typeLineIndex)
				if result.loop {
					collector.RaiseNoEntryPointLoop(relationName, typeName, meta, lineIndex)
				} else {
					collector.RaiseNoEntryPoint(relationName, typeName, meta, lineIndex)
				}
			}
		}
	}
}

// relationMeta resolves the file/module for a relation, falling back to the type.
func relationMeta(typeDef *openfgav1.TypeDefinition, relationName string) *Meta {
	if rm, ok := typeDef.GetMetadata().GetRelations()[relationName]; ok {
		file := rm.GetSourceInfo().GetFile()
		module := rm.GetModule()
		if file == "" {
			file = typeDef.GetMetadata().GetSourceInfo().GetFile()
		}
		if module == "" {
			module = typeDef.GetMetadata().GetModule()
		}
		return &Meta{File: file, Module: module}
	}
	return &Meta{
		File:   typeDef.GetMetadata().GetSourceInfo().GetFile(),
		Module: typeDef.GetMetadata().GetModule(),
	}
}

// hasEntryPointOrLoop determines whether a rewrite reaches a concrete entry point.
// visited tracks type#relation pairs already on the current traversal so that a
// rewrite referencing a relation already being resolved is reported as a loop.
func (cd *CycleDetector) hasEntryPointOrLoop(typeName, relationName string,
	rewrite *openfgav1.Userset, visitedRecords map[string]map[string]bool) entryPointResult {
	if relationName == "" || rewrite == nil {
		return entryPointResult{}
	}

	visited := copyVisited(visitedRecords)
	if visited[typeName] == nil {
		visited[typeName] = map[string]bool{}
	}
	visited[typeName][relationName] = true

	if !cd.validator.RelationDefined(typeName, relationName) {
		return entryPointResult{}
	}

	switch rewrite.GetUserset().(type) {
	case *openfgav1.Userset_This:
		// A direct assignment has an entry point if any assignable type is a
		// concrete type or wildcard. A type#relation restriction only provides an
		// entry point if that referenced relation itself has one.
		for _, tr := range cd.directTypeRestrictions(typeName, relationName) {
			decodedType := tr.GetType()
			decodedRelation := tr.GetRelation()
			isWildcard := tr.GetWildcard() != nil

			if decodedRelation == "" || isWildcard {
				return entryPointResult{hasEntry: true}
			}
			assignable := cd.validator.GetRelationUserset(decodedType, decodedRelation)
			if assignable == nil {
				return entryPointResult{}
			}
			if visited[decodedType][decodedRelation] {
				continue
			}
			if cd.hasEntryPointOrLoop(decodedType, decodedRelation, assignable, visited).hasEntry {
				return entryPointResult{hasEntry: true}
			}
		}
		return entryPointResult{}

	case *openfgav1.Userset_ComputedUserset:
		computed := rewrite.GetComputedUserset().GetRelation()
		if computed == "" || !cd.validator.RelationDefined(typeName, computed) {
			return entryPointResult{}
		}
		if visited[typeName][computed] {
			return entryPointResult{loop: true}
		}
		return cd.hasEntryPointOrLoop(typeName, computed, cd.validator.GetRelationUserset(typeName, computed), visited)

	case *openfgav1.Userset_TupleToUserset:
		ttu := rewrite.GetTupleToUserset()
		tupleset := ttu.GetTupleset().GetRelation()
		computed := ttu.GetComputedUserset().GetRelation()
		if tupleset == "" || computed == "" {
			return entryPointResult{}
		}
		if !cd.validator.RelationDefined(typeName, tupleset) {
			return entryPointResult{}
		}
		for _, tr := range cd.directTypeRestrictions(typeName, tupleset) {
			assignableType := tr.GetType()
			assignable := cd.validator.GetRelationUserset(assignableType, computed)
			if assignable == nil {
				continue
			}
			if visited[assignableType][computed] {
				continue
			}
			if cd.hasEntryPointOrLoop(assignableType, computed, assignable, visited).hasEntry {
				return entryPointResult{hasEntry: true}
			}
		}
		return entryPointResult{}

	case *openfgav1.Userset_Union:
		hasLoop := false
		for _, child := range rewrite.GetUnion().GetChild() {
			res := cd.hasEntryPointOrLoop(typeName, relationName, child, copyVisited(visited))
			if res.hasEntry {
				return entryPointResult{hasEntry: true}
			}
			hasLoop = hasLoop || res.loop
		}
		return entryPointResult{loop: hasLoop}

	case *openfgav1.Userset_Intersection:
		for _, child := range rewrite.GetIntersection().GetChild() {
			res := cd.hasEntryPointOrLoop(typeName, relationName, child, copyVisited(visited))
			if !res.hasEntry {
				return entryPointResult{loop: res.loop}
			}
		}
		return entryPointResult{hasEntry: true}

	case *openfgav1.Userset_Difference:
		diff := rewrite.GetDifference()
		base := cd.hasEntryPointOrLoop(typeName, relationName, diff.GetBase(), visited)
		if !base.hasEntry {
			return entryPointResult{loop: base.loop}
		}
		subtract := cd.hasEntryPointOrLoop(typeName, relationName, diff.GetSubtract(), visited)
		if !subtract.hasEntry {
			return entryPointResult{loop: subtract.loop}
		}
		return entryPointResult{hasEntry: true}
	}

	return entryPointResult{}
}

// directTypeRestrictions returns the directly-related user types declared for a
// relation in its metadata.
func (cd *CycleDetector) directTypeRestrictions(typeName, relationName string) []*openfgav1.RelationReference {
	typeDef := cd.validator.GetTypeDefinition(typeName)
	if typeDef == nil {
		return nil
	}
	rm, ok := typeDef.GetMetadata().GetRelations()[relationName]
	if !ok {
		return nil
	}
	return rm.GetDirectlyRelatedUserTypes()
}

// copyVisited deep-copies the visited map so sibling branches don't share state.
func copyVisited(src map[string]map[string]bool) map[string]map[string]bool {
	dst := make(map[string]map[string]bool, len(src))
	for typeName, relations := range src {
		inner := make(map[string]bool, len(relations))
		for relationName, v := range relations {
			inner[relationName] = v
		}
		dst[typeName] = inner
	}
	return dst
}
