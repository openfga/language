package validation

import (
	"maps"
	"slices"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// ValidateWildcardUsage validates wildcard relation usage rules.
func ValidateWildcardUsage(errs *ValidationErrors, model *openfgav1.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}
	validateWildcardUsage(errs, NewSemanticValidator(model), lines)
}

func validateWildcardUsage(errs *ValidationErrors, validator *SemanticValidator, lines []string) {
	model := validator.model
	if model == nil {
		return
	}
	for _, typeDef := range model.GetTypeDefinitions() {
		if typeDef.GetMetadata() == nil {
			continue
		}
		relationsMetadata := typeDef.GetMetadata().GetRelations()
		for _, relationName := range slices.Sorted(maps.Keys(relationsMetadata)) {
			validateWildcardInRelation(errs, validator, typeDef.GetType(), relationName,
				relationsMetadata[relationName], lines)
		}
	}
}

func validateWildcardInRelation(errs *ValidationErrors, validator *SemanticValidator,
	typeName, relationName string, relationMetadata *openfgav1.RelationMetadata, lines []string) {
	if relationMetadata == nil {
		return
	}
	meta := &Meta{
		File:   relationMetadata.GetSourceInfo().GetFile(),
		Module: relationMetadata.GetModule(),
	}
	// Anchor relation line lookups to this type's declaration so the correct
	// `define` is found when several types share a relation name.
	typeLineIndex := GetTypeLineNumber(typeName, lines, nil)
	for _, typeRestriction := range relationMetadata.GetDirectlyRelatedUserTypes() {
		if typeRestriction.GetType() == "" {
			continue
		}
		if typeRestriction.GetWildcard() != nil {
			validateWildcardRestriction(errs, validator, typeRestriction, relationName, typeName, meta, lines, typeLineIndex)
			// wildcard and explicit relation together is invalid
			if typeRestriction.GetRelation() != "" {
				lineIndex := GetRelationLineNumber(relationName, lines, typeLineIndex)
				errs.Add(newInvalidWildcardUsageError(lines, typeRestriction.GetType(), relationName, typeName,
					"wildcard cannot be used with specific relation", meta, lineIndex))
			}
		}
	}
}

func validateWildcardRestriction(errs *ValidationErrors, validator *SemanticValidator,
	typeRestriction *openfgav1.RelationReference, relationName, typeName string, meta *Meta, lines []string, typeLineIndex *int) {
	if !validator.TypeDefined(typeRestriction.GetType()) {
		lineIndex := GetRelationLineNumber(relationName, lines, typeLineIndex)
		errs.Add(newUndefinedTypeError(lines, typeRestriction.GetType(), relationName, typeName, meta, lineIndex))
	}
}

// ValidateTupleToUsersetRequirements validates tuple-to-userset usage requirements.
func ValidateTupleToUsersetRequirements(errs *ValidationErrors, model *openfgav1.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}
	validateTupleToUsersetRequirements(errs, NewSemanticValidator(model), lines)
}

func validateTupleToUsersetRequirements(errs *ValidationErrors, validator *SemanticValidator, lines []string) {
	model := validator.model
	if model == nil {
		return
	}
	for _, typeDef := range model.GetTypeDefinitions() {
		relations := typeDef.GetRelations()
		for _, relationName := range slices.Sorted(maps.Keys(relations)) {
			validateTupleToUsersetInUserset(errs, validator, typeDef.GetType(), relationName,
				relations[relationName], lines)
		}
	}
}

func validateTupleToUsersetInUserset(errs *ValidationErrors, validator *SemanticValidator,
	typeName, relationName string, userset *openfgav1.Userset, lines []string) {
	if userset == nil {
		return
	}

	if ttu := userset.GetTupleToUserset(); ttu != nil {
		typeDef := validator.GetTypeDefinition(typeName)
		meta := &Meta{
			File:   typeDef.GetMetadata().GetSourceInfo().GetFile(),
			Module: typeDef.GetMetadata().GetModule(),
		}
		validateTupleToUsersetOperation(errs, validator, typeName, relationName, ttu, meta, lines)
	}
	if union := userset.GetUnion(); union != nil {
		for _, child := range union.GetChild() {
			validateTupleToUsersetInUserset(errs, validator, typeName, relationName, child, lines)
		}
	}
	if intersection := userset.GetIntersection(); intersection != nil {
		for _, child := range intersection.GetChild() {
			validateTupleToUsersetInUserset(errs, validator, typeName, relationName, child, lines)
		}
	}
	if diff := userset.GetDifference(); diff != nil {
		validateTupleToUsersetInUserset(errs, validator, typeName, relationName, diff.GetBase(), lines)
		validateTupleToUsersetInUserset(errs, validator, typeName, relationName, diff.GetSubtract(), lines)
	}
}

func validateTupleToUsersetOperation(errs *ValidationErrors, validator *SemanticValidator,
	typeName, relationName string, ttu *openfgav1.TupleToUserset, meta *Meta, lines []string) {
	tuplesetRelation := ttu.GetTupleset().GetRelation()
	if tuplesetRelation == "" {
		return
	}
	// Whether the tupleset/computed relations exist is validated in the
	// relation-reference pass (semantic_validation.go). Here we only check that
	// an existing tupleset relation is directly assignable.
	if !validator.RelationDefined(typeName, tuplesetRelation) {
		return
	}
	validateTuplesetDirectAssignment(errs, validator, typeName, tuplesetRelation, relationName, meta, lines)
}

func validateTuplesetDirectAssignment(errs *ValidationErrors, validator *SemanticValidator,
	typeName, tuplesetRelation, parentRelation string, meta *Meta, lines []string) {
	typeDef := validator.GetTypeDefinition(typeName)
	if typeDef == nil {
		return
	}
	if metaProto := typeDef.GetMetadata(); metaProto != nil {
		if rm, ok := metaProto.GetRelations()[tuplesetRelation]; ok {
			if len(rm.GetDirectlyRelatedUserTypes()) == 0 {
				lineIndex := GetRelationLineNumber(parentRelation, lines, nil)
				errs.Add(newTuplesetNotDirectError(lines, tuplesetRelation, typeName, parentRelation, meta, lineIndex))
			}
		}
	}
}
