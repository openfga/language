package validation

import (
	"fmt"

	fgaSdk "github.com/openfga/go-sdk"
)

// ValidateWildcardUsage validates wildcard relation usage rules
// This ensures wildcards are used appropriately in type restrictions
func ValidateWildcardUsage(collector *ErrorCollector, model *fgaSdk.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}

	validator := NewSemanticValidator(model)

	for _, typeDef := range model.TypeDefinitions {
		if typeDef.Metadata == nil || !typeDef.Metadata.HasRelations() {
			continue
		}

		relations := typeDef.Metadata.GetRelations()
		for relationName, relationMetadata := range relations {
			validateWildcardInRelation(collector, validator, typeDef.Type, relationName, relationMetadata, lines)
		}
	}
}

// validateWildcardInRelation validates wildcard usage within a specific relation.
func validateWildcardInRelation(collector *ErrorCollector, validator *SemanticValidator,
	typeName, relationName string, relationMetadata fgaSdk.RelationMetadata, lines []string) {
	if !relationMetadata.HasDirectlyRelatedUserTypes() {
		return
	}

	// Get file and module metadata
	var file, module string
	if relationMetadata.HasSourceInfo() {
		sourceInfo := relationMetadata.GetSourceInfo()
		file = sourceInfo.GetFile()
	}
	if relationMetadata.HasModule() {
		module = relationMetadata.GetModule()
	}
	meta := &Meta{File: file, Module: module}

	userTypes := relationMetadata.GetDirectlyRelatedUserTypes()
	for _, typeRestriction := range userTypes {
		if typeRestriction.Type == "" {
			continue
		}

		// Check wildcard usage rules
		if typeRestriction.Wildcard != nil {
			// Validate that wildcard is used correctly
			validateWildcardRestriction(collector, validator, typeRestriction, relationName, typeName, meta, lines)
		}

		// Check that wildcard and relation aren't used together inappropriately
		if typeRestriction.Wildcard != nil && typeRestriction.Relation != nil && *typeRestriction.Relation != "" {
			relationLineIndex := GetRelationLineNumber(relationName, lines, nil)
			collector.RaiseInvalidWildcardUsage(typeRestriction.Type, relationName, typeName, "wildcard cannot be used with specific relation", meta, relationLineIndex)
		}
	}
}

// validateWildcardRestriction validates specific wildcard restriction rules.
func validateWildcardRestriction(collector *ErrorCollector, validator *SemanticValidator,
	typeRestriction fgaSdk.RelationReference, relationName, typeName string, meta *Meta, lines []string) {
	// Check that the referenced type exists
	if !validator.TypeDefined(typeRestriction.Type) {
		relationLineIndex := GetRelationLineNumber(relationName, lines, nil)
		collector.RaiseUndefinedType(typeRestriction.Type, relationName, typeName, meta, relationLineIndex)
		return
	}

	// Additional wildcard-specific validations could be added here
	// For example, checking schema version compatibility with wildcards
}

// ValidateTupleToUsersetRequirements validates tuple-to-userset usage requirements
// This ensures tuple-to-userset operations are used correctly
func ValidateTupleToUsersetRequirements(collector *ErrorCollector, model *fgaSdk.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}

	validator := NewSemanticValidator(model)

	for _, typeDef := range model.TypeDefinitions {
		if !typeDef.HasRelations() {
			continue
		}

		relations := typeDef.GetRelations()
		for relationName, userset := range relations {
			validateTupleToUsersetInUserset(collector, validator, typeDef.Type, relationName, userset, lines)
		}
	}
}

// validateTupleToUsersetInUserset recursively validates tuple-to-userset usage.
func validateTupleToUsersetInUserset(collector *ErrorCollector, validator *SemanticValidator,
	typeName, relationName string, userset fgaSdk.Userset, lines []string) {
	// Get file and module metadata
	var file, module string
	if typeDef := validator.GetTypeDefinition(typeName); typeDef != nil && typeDef.Metadata != nil {
		if typeDef.Metadata.HasSourceInfo() {
			sourceInfo := typeDef.Metadata.GetSourceInfo()
			file = sourceInfo.GetFile()
		}
		if typeDef.Metadata.HasModule() {
			module = typeDef.Metadata.GetModule()
		}
	}
	meta := &Meta{File: file, Module: module}

	// Validate tuple-to-userset operations
	if userset.TupleToUserset != nil {
		validateTupleToUsersetOperation(collector, validator, typeName, relationName, *userset.TupleToUserset, meta, lines)
	}

	// Recursively check union operations
	if userset.Union != nil {
		for _, child := range userset.Union.Child {
			validateTupleToUsersetInUserset(collector, validator, typeName, relationName, child, lines)
		}
	}

	// Recursively check intersection operations
	if userset.Intersection != nil {
		for _, child := range userset.Intersection.Child {
			validateTupleToUsersetInUserset(collector, validator, typeName, relationName, child, lines)
		}
	}

	// Recursively check difference operations
	if userset.Difference != nil {
		validateTupleToUsersetInUserset(collector, validator, typeName, relationName, userset.Difference.Base, lines)
		validateTupleToUsersetInUserset(collector, validator, typeName, relationName, userset.Difference.Subtract, lines)
	}
}

// validateTupleToUsersetOperation validates a specific tuple-to-userset operation.
func validateTupleToUsersetOperation(collector *ErrorCollector, validator *SemanticValidator,
	typeName, relationName string, tupleToUserset fgaSdk.TupleToUserset, meta *Meta, lines []string) {
	// Validate that the tupleset relation exists and is appropriate
	if tupleToUserset.Tupleset.HasRelation() {
		tuplesetRelation := tupleToUserset.Tupleset.GetRelation()

		// Check that the tupleset relation exists
		if !validator.RelationDefined(typeName, tuplesetRelation) {
			relationLineIndex := GetRelationLineNumber(relationName, lines, nil)
			collector.RaiseUndefinedRelation(tuplesetRelation, typeName, relationName, typeName, meta, relationLineIndex)
			return
		}

		// Validate that the tupleset relation allows direct assignment
		// (it should have type restrictions that allow tuple-to-userset to work)
		validateTuplesetDirectAssignment(collector, validator, typeName, tuplesetRelation, relationName, meta, lines)
	}

	// Validate that the computed userset relation exists
	if tupleToUserset.ComputedUserset.HasRelation() {
		computedRelation := tupleToUserset.ComputedUserset.GetRelation()

		if !validator.RelationDefined(typeName, computedRelation) {
			relationLineIndex := GetRelationLineNumber(relationName, lines, nil)
			collector.RaiseUndefinedRelation(computedRelation, typeName, relationName, typeName, meta, relationLineIndex)
		}
	}
}

// validateTuplesetDirectAssignment ensures tupleset relations support direct assignment.
func validateTuplesetDirectAssignment(collector *ErrorCollector, validator *SemanticValidator,
	typeName, tuplesetRelation, parentRelation string, meta *Meta, lines []string) {
	// Get the type definition to check relation metadata
	if typeDef := validator.GetTypeDefinition(typeName); typeDef != nil {
		if typeDef.Metadata != nil && typeDef.Metadata.HasRelations() {
			relations := typeDef.Metadata.GetRelations()
			if relationMeta, exists := relations[tuplesetRelation]; exists {
				// Check if the tupleset relation has direct type restrictions
				if !relationMeta.HasDirectlyRelatedUserTypes() {
					relationLineIndex := GetRelationLineNumber(parentRelation, lines, nil)
					collector.RaiseTuplesetNotDirect(tuplesetRelation, typeName, parentRelation, meta, relationLineIndex)
				}
			}
		}
	}
}

// Additional error raising methods for wildcard and tuple-to-userset validation.
func (ec *ErrorCollector) RaiseInvalidWildcardUsage(typeName, relationName, parentTypeName, reason string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("Invalid wildcard usage for type '%s' in relation '%s' of type '%s': %s", typeName, relationName, parentTypeName, reason)
	ec.addError(message, InvalidWildcardError, typeName, lineIndex, meta, nil)
}

func (ec *ErrorCollector) RaiseTuplesetNotDirect(tuplesetRelation, typeName, parentRelation string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("Tupleset relation '%s' on type '%s' must allow direct assignment (used in relation '%s')", tuplesetRelation, typeName, parentRelation)
	ec.addError(message, TuplesetNotDirect, tuplesetRelation, lineIndex, meta, nil)
}
