package validation

import (
	"maps"
	"slices"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// validateRelationReferences checks that every type and relation a relation
// refers to — in its type restrictions and in its rewrites — exists in the
// model.
func validateRelationReferences(idx *index, src source) Findings {
	var fs Findings

	for _, typeDef := range idx.model.GetTypeDefinitions() {
		typeName := typeDef.GetType()

		// When a type is declared more than once only the winning (last)
		// definition is validated: the reference resolves a shadowed duplicate's
		// relations against the winning type, so validating the winner alone
		// avoids spurious reference errors. The duplicate itself is reported by
		// the duplicates phase.
		if winning := idx.typeDef(typeName); winning != nil && winning != typeDef {
			continue
		}

		// Anchor relation line lookups to the type's declaration so the correct
		// `define` occurrence is found when several types declare a relation of
		// the same name.
		typeLine := src.typeLine(typeName)

		if meta := typeDef.GetMetadata(); meta != nil {
			relationsMetadata := meta.GetRelations()
			for _, relationName := range slices.Sorted(maps.Keys(relationsMetadata)) {
				fs = append(fs, validateTypeRestrictions(idx, src, typeDef, relationName,
					relationsMetadata[relationName], typeLine)...)
			}
		}

		relations := typeDef.GetRelations()
		for _, relationName := range slices.Sorted(maps.Keys(relations)) {
			fs = append(fs, validateUsersetReferences(idx, src, typeDef, relationName,
				relations[relationName], typeLine)...)
		}
	}

	return fs
}

// validateTypeRestrictions checks a relation's directly-related user types:
// each restriction must name a defined type, and a `type#relation` restriction
// a relation defined on that type.
func validateTypeRestrictions(idx *index, src source, typeDef *openfgav1.TypeDefinition,
	relationName string, relationMetadata *openfgav1.RelationMetadata, typeLine int) Findings {
	if relationMetadata == nil {
		return nil
	}

	var fs Findings

	typeName := typeDef.GetType()
	file := relationMetadata.GetSourceInfo().GetFile()
	module := relationMetadata.GetModule()

	for _, restriction := range relationMetadata.GetDirectlyRelatedUserTypes() {
		restrictedType := restriction.GetType()
		if restrictedType == "" {
			continue
		}

		// A directly-related type that doesn't exist: `X` is not a valid type.
		if !idx.typeDefined(restrictedType) {
			line := src.relationLine(relationName, typeLine)
			fs = append(fs, invalidType(restrictedType).atRestriction(src, line).in(file, module))

			continue
		}

		// A type#relation restriction whose relation doesn't exist on that
		// type: `rel` is not a valid relation for `X`.
		if rel := restriction.GetRelation(); rel != "" && !idx.relationDefined(restrictedType, rel) {
			line := src.relationLine(relationName, typeLine)
			fs = append(fs, invalidTypeRelation(restrictedType+"#"+rel, restrictedType, relationName, rel, typeName).
				at(src, line).in(file, module))
		}
	}

	return fs
}

// validateUsersetReferences checks the relations a rewrite names: a computed
// userset must exist on the type, and a tuple-to-userset must satisfy
// validateTupleToUsersetReferences. Union, intersection and difference are
// walked into.
func validateUsersetReferences(idx *index, src source, typeDef *openfgav1.TypeDefinition,
	relationName string, userset *openfgav1.Userset, typeLine int) Findings {
	if userset == nil {
		return nil
	}

	var fs Findings

	typeName := typeDef.GetType()
	file, module := typeMeta(typeDef)

	if computed := userset.GetComputedUserset(); computed != nil {
		// `define a: b` where b is not a relation on this type.
		if target := computed.GetRelation(); target != "" && !idx.relationDefined(typeName, target) {
			line := src.relationLine(relationName, typeLine)
			fs = append(fs, missingRelation(target, typeName, relationName).at(src, line).in(file, module))
		}
	}

	if ttu := userset.GetTupleToUserset(); ttu != nil {
		fs = append(fs, validateTupleToUsersetReferences(idx, src, typeDef, relationName, ttu, typeLine)...)
	}

	if union := userset.GetUnion(); union != nil {
		for _, child := range union.GetChild() {
			fs = append(fs, validateUsersetReferences(idx, src, typeDef, relationName, child, typeLine)...)
		}
	}

	if intersection := userset.GetIntersection(); intersection != nil {
		for _, child := range intersection.GetChild() {
			fs = append(fs, validateUsersetReferences(idx, src, typeDef, relationName, child, typeLine)...)
		}
	}

	if diff := userset.GetDifference(); diff != nil {
		fs = append(fs, validateUsersetReferences(idx, src, typeDef, relationName, diff.GetBase(), typeLine)...)
		fs = append(fs, validateUsersetReferences(idx, src, typeDef, relationName, diff.GetSubtract(), typeLine)...)
	}

	return fs
}

// validateTupleToUsersetReferences validates a `target from tupleset` rewrite,
// mirroring the reference implementation:
//   - the tupleset relation must exist on the current type;
//   - the tupleset relation must be a plain direct assignment whose assignable
//     types are concrete (no wildcard, no type#relation);
//   - the computed target relation must exist on at least one of the types the
//     tupleset relation is assignable to.
func validateTupleToUsersetReferences(idx *index, src source, typeDef *openfgav1.TypeDefinition,
	relationName string, ttu *openfgav1.TupleToUserset, typeLine int) Findings {
	fromRelation := ttu.GetTupleset().GetRelation()
	targetRelation := ttu.GetComputedUserset().GetRelation()

	if fromRelation == "" || targetRelation == "" {
		return nil
	}

	typeName := typeDef.GetType()
	file, module := typeMeta(typeDef)
	line := src.relationLine(relationName, typeLine)
	symbol := targetRelation + " from " + fromRelation

	// 1. The tupleset relation must exist on the current type.
	if !idx.relationDefined(typeName, fromRelation) {
		return Findings{invalidTypeRelation(symbol, typeName, relationName, fromRelation, typeName).
			at(src, line).in(file, module)}
	}

	// 2. The tupleset relation must be a single direct assignment.
	fromTypes, isDirect := idx.directlyAssignableTypes(typeName, fromRelation)
	if !isDirect || len(fromTypes) == 0 {
		return Findings{tupleUsersetRequiresDirect(fromRelation, typeName, relationName).
			atFromClause(src, line).in(file, module)}
	}

	// 3. Each assignable type of the tupleset relation must be a concrete type
	//    (no wildcard, no type#relation), and the computed target must exist on
	//    at least one of them.
	var fs Findings

	notValid := make([]*openfgav1.RelationReference, 0, len(fromTypes))

	for _, restriction := range fromTypes {
		if restriction.GetWildcard() != nil || restriction.GetRelation() != "" {
			// A wildcard or type#relation cannot be used as a tupleset target.
			fs = append(fs, tupleUsersetRequiresDirect(fromRelation, typeName, relationName).
				atFromClause(src, line).in(file, module))

			continue
		}

		targetType := restriction.GetType()
		if !idx.typeDefined(targetType) || !idx.relationDefined(targetType, targetRelation) {
			notValid = append(notValid, restriction)
		}
	}

	// If the target is missing on every assignable type, report it per type.
	if len(notValid) == len(fromTypes) {
		for _, restriction := range notValid {
			fs = append(fs, invalidRelationOnTupleset(symbol, targetRelation, typeName, fromRelation,
				restriction.GetType(), relationName).at(src, line).in(file, module))
		}
	}

	return fs
}
