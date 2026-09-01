package validation

import (
	"maps"
	"slices"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// validateDuplicates reports everything the model defines twice: a type
// declared twice, a type restriction repeated in a relation, and a partial
// relation definition repeated in a union, intersection or difference.
func validateDuplicates(model *openfgav1.AuthorizationModel, src source) Findings {
	var fs Findings

	seenTypes := make(map[string]bool)

	for _, typeDef := range model.GetTypeDefinitions() {
		typeName := typeDef.GetType()
		if typeName == "" {
			continue
		}

		file, module := typeMeta(typeDef)

		if seenTypes[typeName] {
			fs = append(fs, duplicateTypeName(typeName).at(src, src.typeLine(typeName)).in(file, module))
		}

		seenTypes[typeName] = true

		typeLine := src.typeLine(typeName)

		relationsMetadata := typeDef.GetMetadata().GetRelations()
		for _, relationName := range slices.Sorted(maps.Keys(relationsMetadata)) {
			fs = append(fs, duplicateRestrictionsIn(src, relationsMetadata[relationName],
				relationName, typeDef, typeLine)...)
			fs = append(fs, duplicateOperandsIn(src, typeDef, relationName, typeLine)...)
		}
	}

	return fs
}

// duplicateRestrictionsIn flags a type restriction repeated in one relation.
// Restrictions are compared as written: `user`, `user:*`, `user#member` and
// `user with cond` are all distinct.
func duplicateRestrictionsIn(src source, relationMetadata *openfgav1.RelationMetadata,
	relationName string, typeDef *openfgav1.TypeDefinition, typeLine int) Findings {
	if relationMetadata == nil {
		return nil
	}

	var fs Findings

	typeName := typeDef.GetType()
	file, module := typeMeta(typeDef)
	seen := make(map[string]bool)

	for _, restriction := range relationMetadata.GetDirectlyRelatedUserTypes() {
		if restriction.GetType() == "" {
			continue
		}

		written := restriction.GetType()
		if restriction.GetWildcard() != nil {
			written += ":*"
		} else if rel := restriction.GetRelation(); rel != "" {
			written += "#" + rel
		}

		if condition := restriction.GetCondition(); condition != "" {
			written += " with " + condition
		}

		if seen[written] {
			line := src.relationLine(relationName, typeLine)
			fs = append(fs, duplicateTypeRestriction(written, relationName, typeName).at(src, line).in(file, module))
		}

		seen[written] = true
	}

	return fs
}

// duplicateOperandsIn flags a partial relation definition repeated in a union
// or intersection, and a difference that subtracts an operand from itself.
func duplicateOperandsIn(src source, typeDef *openfgav1.TypeDefinition,
	relationName string, typeLine int) Findings {
	relation, ok := typeDef.GetRelations()[relationName]
	if !ok {
		return nil
	}

	file, module := relationMeta(typeDef, relationName)

	var fs Findings

	raise := func(operand string) {
		line := src.relationLine(relationName, typeLine)
		fs = append(fs, duplicatePartialRelation(operand, relationName, typeDef.GetType()).
			at(src, line).in(file, module))
	}

	// Union and intersection both store their members as a *openfgav1.Usersets
	// and treat a repeated member as a duplicate, so they share the check.
	for _, operands := range []*openfgav1.Usersets{relation.GetUnion(), relation.GetIntersection()} {
		if operands == nil {
			continue
		}

		seen := make(map[string]bool)

		for _, child := range operands.GetChild() {
			name := operandName(child)
			if name == "" {
				continue
			}

			if seen[name] {
				raise(name)
			}

			seen[name] = true
		}
	}

	if diff := relation.GetDifference(); diff != nil {
		base := operandName(diff.GetBase())
		if base != "" && base == operandName(diff.GetSubtract()) {
			raise(base)
		}
	}

	return fs
}

// operandName renders a union/intersection/difference member the way it was
// written: a computed userset as its relation, a tuple-to-userset as
// `target from tupleset`.
func operandName(userset *openfgav1.Userset) string {
	if userset == nil {
		return ""
	}

	if computed := userset.GetComputedUserset(); computed.GetRelation() != "" {
		return computed.GetRelation()
	}

	if ttu := userset.GetTupleToUserset(); ttu != nil {
		target := ttu.GetComputedUserset().GetRelation()
		from := ttu.GetTupleset().GetRelation()

		switch {
		case target != "" && from != "":
			return target + " from " + from
		case target != "":
			return target
		}
	}

	return ""
}
