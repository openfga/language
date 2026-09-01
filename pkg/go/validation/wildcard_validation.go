package validation

import (
	"maps"
	"slices"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// validateWildcards checks every wildcard type restriction: the type it
// restricts must exist, and a restriction cannot carry both a wildcard and a
// relation.
func validateWildcards(idx *index, src source) Findings {
	var fs Findings

	for _, typeDef := range idx.model.GetTypeDefinitions() {
		if typeDef.GetMetadata() == nil {
			continue
		}

		typeName := typeDef.GetType()

		// Anchor relation line lookups to this type's declaration so the correct
		// `define` is found when several types share a relation name.
		typeLine := src.typeLine(typeName)

		relationsMetadata := typeDef.GetMetadata().GetRelations()
		for _, relationName := range slices.Sorted(maps.Keys(relationsMetadata)) {
			relationMetadata := relationsMetadata[relationName]
			if relationMetadata == nil {
				continue
			}

			file := relationMetadata.GetSourceInfo().GetFile()
			module := relationMetadata.GetModule()

			for _, restriction := range relationMetadata.GetDirectlyRelatedUserTypes() {
				if restriction.GetType() == "" || restriction.GetWildcard() == nil {
					continue
				}

				if !idx.typeDefined(restriction.GetType()) {
					line := src.relationLine(relationName, typeLine)
					fs = append(fs, undefinedType(restriction.GetType(), relationName, typeName).
						at(src, line).in(file, module))
				}

				// A wildcard and an explicit relation together is invalid.
				if restriction.GetRelation() != "" {
					line := src.relationLine(relationName, typeLine)
					fs = append(fs, invalidWildcardUsage(restriction.GetType(), relationName, typeName,
						"wildcard cannot be used with specific relation").at(src, line).in(file, module))
				}
			}
		}
	}

	return fs
}

// validateTupleToUsersets checks that every tupleset relation used in a
// `target from tupleset` rewrite allows direct assignment. Whether the tupleset
// and computed relations exist is the reference phase's job; here an existing
// tupleset relation with no assignable types is reported.
func validateTupleToUsersets(idx *index, src source) Findings {
	var fs Findings

	for _, typeDef := range idx.model.GetTypeDefinitions() {
		relations := typeDef.GetRelations()
		for _, relationName := range slices.Sorted(maps.Keys(relations)) {
			fs = append(fs, tuplesetsIn(idx, src, typeDef.GetType(), relationName, relations[relationName])...)
		}
	}

	return fs
}

// tuplesetsIn walks one relation's rewrite tree and reports each
// tuple-to-userset whose tupleset relation is not directly assignable.
func tuplesetsIn(idx *index, src source, typeName, relationName string, userset *openfgav1.Userset) Findings {
	if userset == nil {
		return nil
	}

	var fs Findings

	if ttu := userset.GetTupleToUserset(); ttu != nil {
		fs = fs.add(tuplesetNotAssignable(idx, src, typeName, relationName, ttu))
	}

	if union := userset.GetUnion(); union != nil {
		for _, child := range union.GetChild() {
			fs = append(fs, tuplesetsIn(idx, src, typeName, relationName, child)...)
		}
	}

	if intersection := userset.GetIntersection(); intersection != nil {
		for _, child := range intersection.GetChild() {
			fs = append(fs, tuplesetsIn(idx, src, typeName, relationName, child)...)
		}
	}

	if diff := userset.GetDifference(); diff != nil {
		fs = append(fs, tuplesetsIn(idx, src, typeName, relationName, diff.GetBase())...)
		fs = append(fs, tuplesetsIn(idx, src, typeName, relationName, diff.GetSubtract())...)
	}

	return fs
}

// tuplesetNotAssignable reports a defined tupleset relation that declares no
// directly-related user types, or nil when it declares some or does not exist.
func tuplesetNotAssignable(idx *index, src source, typeName, relationName string,
	ttu *openfgav1.TupleToUserset) *Finding {
	tuplesetRelation := ttu.GetTupleset().GetRelation()
	if tuplesetRelation == "" || !idx.relationDefined(typeName, tuplesetRelation) {
		return nil
	}

	typeDef := idx.typeDef(typeName)

	relationMetadata, ok := typeDef.GetMetadata().GetRelations()[tuplesetRelation]
	if !ok || len(relationMetadata.GetDirectlyRelatedUserTypes()) > 0 {
		return nil
	}

	file, module := typeMeta(typeDef)
	line := src.relationLine(relationName, -1)

	return tuplesetNotDirect(tuplesetRelation, typeName, relationName).at(src, line).in(file, module)
}
