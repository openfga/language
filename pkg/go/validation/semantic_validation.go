package validation

import (
	fgaSdk "github.com/openfga/go-sdk"
)

// SemanticValidator handles semantic validation of authorization models
// This is equivalent to the semantic validation logic in the JS implementation
type SemanticValidator struct {
	model       *fgaSdk.AuthorizationModel
	typeMap     map[string]*fgaSdk.TypeDefinition
	relationMap map[string]map[string]*fgaSdk.Userset
}

// NewSemanticValidator creates a new semantic validator for the given model.
func NewSemanticValidator(model *fgaSdk.AuthorizationModel) *SemanticValidator {
	validator := &SemanticValidator{
		model:       model,
		typeMap:     make(map[string]*fgaSdk.TypeDefinition),
		relationMap: make(map[string]map[string]*fgaSdk.Userset),
	}

	// Build type and relation maps for efficient lookups
	validator.buildMaps()
	return validator
}

// buildMaps constructs internal maps for efficient type and relation lookups.
func (sv *SemanticValidator) buildMaps() {
	if sv.model == nil {
		return
	}

	for i, typeDef := range sv.model.TypeDefinitions {
		sv.typeMap[typeDef.Type] = &sv.model.TypeDefinitions[i]

		if typeDef.HasRelations() {
			relations := typeDef.GetRelations()
			sv.relationMap[typeDef.Type] = make(map[string]*fgaSdk.Userset)

			for relationName, userset := range relations {
				sv.relationMap[typeDef.Type][relationName] = &userset
			}
		}
	}
}

// RelationDefined checks if a relation is defined for a given type
// This is equivalent to the relationDefined function in the JS implementation
func (sv *SemanticValidator) RelationDefined(typeName, relationName string) bool {
	if relations, exists := sv.relationMap[typeName]; exists {
		_, relationExists := relations[relationName]
		return relationExists
	}
	return false
}

// TypeDefined checks if a type is defined in the model.
func (sv *SemanticValidator) TypeDefined(typeName string) bool {
	_, exists := sv.typeMap[typeName]
	return exists
}

// GetTypeDefinition returns the type definition for a given type name.
func (sv *SemanticValidator) GetTypeDefinition(typeName string) *fgaSdk.TypeDefinition {
	return sv.typeMap[typeName]
}

// GetRelationUserset returns the userset definition for a specific relation.
func (sv *SemanticValidator) GetRelationUserset(typeName, relationName string) *fgaSdk.Userset {
	if relations, exists := sv.relationMap[typeName]; exists {
		return relations[relationName]
	}
	return nil
}

// ValidateRelationReferences validates that all relation references in the model are valid
// This checks that referenced types and relations actually exist
func ValidateRelationReferences(collector *ErrorCollector, model *fgaSdk.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}

	validator := NewSemanticValidator(model)

	for _, typeDef := range model.TypeDefinitions {
		// Validate relations in metadata (type restrictions)
		if typeDef.Metadata != nil && typeDef.Metadata.HasRelations() {
			relations := typeDef.Metadata.GetRelations()
			for relationName, relationMetadata := range relations {
				validateTypeRestrictions(collector, validator, typeDef.Type, relationName, relationMetadata, lines)
			}
		}

		// Validate userset definitions
		if typeDef.HasRelations() {
			relations := typeDef.GetRelations()
			for relationName, userset := range relations {
				validateUsersetReferences(collector, validator, typeDef.Type, relationName, userset, lines)
			}
		}
	}
}

// validateTypeRestrictions validates type restrictions in relation metadata
func validateTypeRestrictions(collector *ErrorCollector, validator *SemanticValidator, 
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

		// Check if the referenced type exists
		if !validator.TypeDefined(typeRestriction.Type) {
			relationLineIndex := GetRelationLineNumber(relationName, lines, nil)
			collector.RaiseUndefinedType(typeRestriction.Type, relationName, typeName, meta, relationLineIndex)
		}

		// If there's a relation specified, check if it exists on the referenced type
		if typeRestriction.Relation != nil && *typeRestriction.Relation != "" {
			if !validator.RelationDefined(typeRestriction.Type, *typeRestriction.Relation) {
				relationLineIndex := GetRelationLineNumber(relationName, lines, nil)
				collector.RaiseUndefinedRelation(*typeRestriction.Relation, typeRestriction.Type, relationName, typeName, meta, relationLineIndex)
			}
		}
	}
}

// validateUsersetReferences validates references in userset definitions
func validateUsersetReferences(collector *ErrorCollector, validator *SemanticValidator, 
	typeName, relationName string, userset fgaSdk.Userset, lines []string) {
	// Get file and module metadata from type definition
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

	// Validate computed userset
	if userset.ComputedUserset != nil && userset.ComputedUserset.HasRelation() {
		targetRelation := userset.ComputedUserset.GetRelation()
		if !validator.RelationDefined(typeName, targetRelation) {
			relationLineIndex := GetRelationLineNumber(relationName, lines, nil)
			collector.RaiseUndefinedRelation(targetRelation, typeName, relationName, typeName, meta, relationLineIndex)
		}
	}

	// Validate tuple-to-userset
	if userset.TupleToUserset != nil {
		// Validate the computed userset part
		if userset.TupleToUserset.ComputedUserset.HasRelation() {
			targetRelation := userset.TupleToUserset.ComputedUserset.GetRelation()
			if !validator.RelationDefined(typeName, targetRelation) {
				relationLineIndex := GetRelationLineNumber(relationName, lines, nil)
				collector.RaiseUndefinedRelation(targetRelation, typeName, relationName, typeName, meta, relationLineIndex)
			}
		}

		// Validate the tupleset part
		if userset.TupleToUserset.Tupleset.HasRelation() {
			tuplesetRelation := userset.TupleToUserset.Tupleset.GetRelation()
			if !validator.RelationDefined(typeName, tuplesetRelation) {
				relationLineIndex := GetRelationLineNumber(relationName, lines, nil)
				collector.RaiseUndefinedRelation(tuplesetRelation, typeName, relationName, typeName, meta, relationLineIndex)
			}
		}
	}

	// Recursively validate union operations
	if userset.Union != nil {
		for _, child := range userset.Union.Child {
			validateUsersetReferences(collector, validator, typeName, relationName, child, lines)
		}
	}

	// Recursively validate intersection operations
	if userset.Intersection != nil {
		for _, child := range userset.Intersection.Child {
			validateUsersetReferences(collector, validator, typeName, relationName, child, lines)
		}
	}

	// Recursively validate difference operations
	if userset.Difference != nil {
		validateUsersetReferences(collector, validator, typeName, relationName, userset.Difference.Base, lines)
		validateUsersetReferences(collector, validator, typeName, relationName, userset.Difference.Subtract, lines)
	}
}
