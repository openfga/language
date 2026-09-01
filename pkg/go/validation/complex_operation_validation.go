package validation

import (
	"maps"
	"slices"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// validateComplexOperations walks every relation's rewrite tree and reports
// operations that are wrong by construction: a union repeating a member, an
// intersection of conflicting direct assignments, and a difference subtracting
// an operand from itself.
func validateComplexOperations(idx *index, src source) Findings {
	var fs Findings

	for _, typeDef := range idx.model.GetTypeDefinitions() {
		relations := typeDef.GetRelations()
		for _, relationName := range slices.Sorted(maps.Keys(relations)) {
			fs = append(fs, operationsIn(idx, src, typeDef.GetType(), relationName,
				relations[relationName], make(map[string]bool))...)
		}
	}

	return fs
}

// operationsIn checks one rewrite and recurses into its children. The visited
// map guards the hop a tuple-to-userset makes to its computed relation, so a
// pair of relations referring to each other terminates.
func operationsIn(idx *index, src source, typeName, relationName string,
	userset *openfgav1.Userset, visited map[string]bool) Findings {
	if userset == nil {
		return nil
	}

	var fs Findings

	if union := userset.GetUnion(); union != nil && len(union.GetChild()) > 0 {
		fs = append(fs, redundantUnionMembersIn(idx, src, typeName, relationName, union)...)

		for _, child := range union.GetChild() {
			fs = append(fs, operationsIn(idx, src, typeName, relationName, child, visited)...)
		}
	}

	if intersection := userset.GetIntersection(); intersection != nil && len(intersection.GetChild()) > 0 {
		fs = append(fs, impossibleIntersectionsIn(idx, src, typeName, relationName, intersection)...)

		for _, child := range intersection.GetChild() {
			fs = append(fs, operationsIn(idx, src, typeName, relationName, child, visited)...)
		}
	}

	if diff := userset.GetDifference(); diff != nil {
		fs = append(fs, operationsIn(idx, src, typeName, relationName, diff.GetBase(), visited)...)
		fs = append(fs, operationsIn(idx, src, typeName, relationName, diff.GetSubtract(), visited)...)
		fs = fs.add(emptyDifferenceIn(idx, src, typeName, relationName, diff))
	}

	if ttu := userset.GetTupleToUserset(); ttu != nil {
		if target := ttu.GetComputedUserset().GetRelation(); target != "" {
			key := typeName + "#" + target
			if !visited[key] {
				visited[key] = true

				if targetUserset := idx.userset(typeName, target); targetUserset != nil {
					fs = append(fs, operationsIn(idx, src, typeName, target, targetUserset, visited)...)
				}
			}
		}
	}

	return fs
}

// redundantUnionMembersIn flags a union member repeated within one union.
func redundantUnionMembersIn(idx *index, src source, typeName, relationName string,
	union *openfgav1.Usersets) Findings {
	var fs Findings

	seen := make(map[string]bool)

	for _, child := range union.GetChild() {
		key := operationKey(child)
		if key == "" {
			continue
		}

		if seen[key] {
			line := src.relationLine(relationName, -1)
			file, module := typeMeta(idx.typeDef(typeName))
			fs = append(fs, redundantUnionMember(key, relationName, typeName).at(src, line).in(file, module))
		}

		seen[key] = true
	}

	return fs
}

// impossibleIntersectionsIn flags an intersection whose direct-assignment
// members can never agree.
func impossibleIntersectionsIn(idx *index, src source, typeName, relationName string,
	intersection *openfgav1.Usersets) Findings {
	restrictions := make([]string, 0)

	for _, child := range intersection.GetChild() {
		if child.GetThis() != nil {
			restrictions = append(restrictions, "this")
		}
	}

	if len(restrictions) <= 1 {
		return nil
	}

	unique := make(map[string]bool)
	for _, restriction := range restrictions {
		unique[restriction] = true
	}

	if len(unique) <= 1 {
		return nil
	}

	line := src.relationLine(relationName, -1)
	file, module := typeMeta(idx.typeDef(typeName))

	return Findings{impossibleIntersection(relationName, typeName, restrictions).at(src, line).in(file, module)}
}

// emptyDifferenceIn flags a difference subtracting an operand from itself,
// which is empty by construction.
func emptyDifferenceIn(idx *index, src source, typeName, relationName string,
	diff *openfgav1.Difference) *Finding {
	base := operationKey(diff.GetBase())
	if base == "" || base != operationKey(diff.GetSubtract()) {
		return nil
	}

	line := src.relationLine(relationName, -1)
	file, module := typeMeta(idx.typeDef(typeName))

	return emptyDifference(relationName, typeName, base).at(src, line).in(file, module)
}

// operationKey names a rewrite for comparison: direct assignment, a computed
// relation, or a tuple-to-userset.
func operationKey(userset *openfgav1.Userset) string {
	if userset == nil {
		return ""
	}

	if userset.GetThis() != nil {
		return "this"
	}

	if computed := userset.GetComputedUserset(); computed.GetRelation() != "" {
		return "computed:" + computed.GetRelation()
	}

	if ttu := userset.GetTupleToUserset(); ttu != nil {
		return "ttu:" + ttu.GetTupleset().GetRelation() + ":" + ttu.GetComputedUserset().GetRelation()
	}

	return ""
}
