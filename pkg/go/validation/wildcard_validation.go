package validation

import (
	"fmt"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// ValidateWildcardUsage validates wildcard relation usage rules.
func ValidateWildcardUsage(collector *ErrorCollector, model *openfgav1.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}
	validateWildcardUsage(collector, NewSemanticValidator(model), lines)
}

func validateWildcardUsage(collector *ErrorCollector, validator *SemanticValidator, lines []string) {
	model := validator.model
	if model == nil {
		return
	}
	for _, typeDef := range model.GetTypeDefinitions() {
		if typeDef.GetMetadata() == nil {
			continue
		}
		for relationName, relationMetadata := range typeDef.GetMetadata().GetRelations() {
			validateWildcardInRelation(collector, validator, typeDef.GetType(), relationName, relationMetadata, lines)
		}
	}
}

func validateWildcardInRelation(collector *ErrorCollector, validator *SemanticValidator,
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
			validateWildcardRestriction(collector, validator, typeRestriction, relationName, typeName, meta, lines, typeLineIndex)
			// wildcard and explicit relation together is invalid
			if typeRestriction.GetRelation() != "" {
				lineIndex := GetRelationLineNumber(relationName, lines, typeLineIndex)
				collector.RaiseInvalidWildcardUsage(typeRestriction.GetType(), relationName, typeName,
					"wildcard cannot be used with specific relation", meta, lineIndex)
			}
		}
	}
}

func validateWildcardRestriction(collector *ErrorCollector, validator *SemanticValidator,
	typeRestriction *openfgav1.RelationReference, relationName, typeName string, meta *Meta, lines []string, typeLineIndex *int) {
	if !validator.TypeDefined(typeRestriction.GetType()) {
		lineIndex := GetRelationLineNumber(relationName, lines, typeLineIndex)
		collector.RaiseUndefinedType(typeRestriction.GetType(), relationName, typeName, meta, lineIndex)
	}
}

// ValidateTupleToUsersetRequirements validates tuple-to-userset usage requirements.
func ValidateTupleToUsersetRequirements(collector *ErrorCollector, model *openfgav1.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}
	validateTupleToUsersetRequirements(collector, NewSemanticValidator(model), lines)
}

func validateTupleToUsersetRequirements(collector *ErrorCollector, validator *SemanticValidator, lines []string) {
	model := validator.model
	if model == nil {
		return
	}
	for _, typeDef := range model.GetTypeDefinitions() {
		for relationName, userset := range typeDef.GetRelations() {
			validateTupleToUsersetInUserset(collector, validator, typeDef.GetType(), relationName, userset, lines)
		}
	}
}

func validateTupleToUsersetInUserset(collector *ErrorCollector, validator *SemanticValidator,
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
		validateTupleToUsersetOperation(collector, validator, typeName, relationName, ttu, meta, lines)
	}
	if union := userset.GetUnion(); union != nil {
		for _, child := range union.GetChild() {
			validateTupleToUsersetInUserset(collector, validator, typeName, relationName, child, lines)
		}
	}
	if intersection := userset.GetIntersection(); intersection != nil {
		for _, child := range intersection.GetChild() {
			validateTupleToUsersetInUserset(collector, validator, typeName, relationName, child, lines)
		}
	}
	if diff := userset.GetDifference(); diff != nil {
		validateTupleToUsersetInUserset(collector, validator, typeName, relationName, diff.GetBase(), lines)
		validateTupleToUsersetInUserset(collector, validator, typeName, relationName, diff.GetSubtract(), lines)
	}
}

func validateTupleToUsersetOperation(collector *ErrorCollector, validator *SemanticValidator,
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
	validateTuplesetDirectAssignment(collector, validator, typeName, tuplesetRelation, relationName, meta, lines)
}

func validateTuplesetDirectAssignment(collector *ErrorCollector, validator *SemanticValidator,
	typeName, tuplesetRelation, parentRelation string, meta *Meta, lines []string) {
	typeDef := validator.GetTypeDefinition(typeName)
	if typeDef == nil {
		return
	}
	if metaProto := typeDef.GetMetadata(); metaProto != nil {
		if rm, ok := metaProto.GetRelations()[tuplesetRelation]; ok {
			if len(rm.GetDirectlyRelatedUserTypes()) == 0 {
				lineIndex := GetRelationLineNumber(parentRelation, lines, nil)
				collector.RaiseTuplesetNotDirect(tuplesetRelation, typeName, parentRelation, meta, lineIndex)
			}
		}
	}
}

func (c *ErrorCollector) RaiseInvalidWildcardUsage(typeName, relationName, parentTypeName, reason string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("Invalid wildcard usage for type '%s' in relation '%s' of type '%s': %s",
		typeName, relationName, parentTypeName, reason)
	c.addError(message, InvalidWildcardError, typeName, lineIndex, meta, nil)
}

func (c *ErrorCollector) RaiseTuplesetNotDirect(tuplesetRelation, typeName, parentRelation string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("Tupleset relation '%s' on type '%s' must allow direct assignment (used in relation '%s')",
		tuplesetRelation, typeName, parentRelation)
	c.addError(message, TuplesetNotDirect, tuplesetRelation, lineIndex, meta, nil)
}
