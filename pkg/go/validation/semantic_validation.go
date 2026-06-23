package validation

import (
	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// SemanticValidator handles semantic validation of authorization models.
type SemanticValidator struct {
	model       *openfgav1.AuthorizationModel
	typeMap     map[string]*openfgav1.TypeDefinition
	relationMap map[string]map[string]*openfgav1.Userset
}

func NewSemanticValidator(model *openfgav1.AuthorizationModel) *SemanticValidator {
	validator := &SemanticValidator{
		model:       model,
		typeMap:     make(map[string]*openfgav1.TypeDefinition),
		relationMap: make(map[string]map[string]*openfgav1.Userset),
	}
	validator.buildMaps()
	return validator
}

func (sv *SemanticValidator) buildMaps() {
	if sv.model == nil {
		return
	}
	for _, typeDef := range sv.model.GetTypeDefinitions() {
		sv.typeMap[typeDef.GetType()] = typeDef
		if relations := typeDef.GetRelations(); len(relations) > 0 {
			sv.relationMap[typeDef.GetType()] = make(map[string]*openfgav1.Userset)
			for relationName, userset := range relations {
				sv.relationMap[typeDef.GetType()][relationName] = userset
			}
		}
	}
}

func (sv *SemanticValidator) RelationDefined(typeName, relationName string) bool {
	if relations, exists := sv.relationMap[typeName]; exists {
		_, relationExists := relations[relationName]
		return relationExists
	}
	return false
}

func (sv *SemanticValidator) TypeDefined(typeName string) bool {
	_, exists := sv.typeMap[typeName]
	return exists
}

func (sv *SemanticValidator) GetTypeDefinition(typeName string) *openfgav1.TypeDefinition {
	return sv.typeMap[typeName]
}

func (sv *SemanticValidator) GetRelationUserset(typeName, relationName string) *openfgav1.Userset {
	if relations, exists := sv.relationMap[typeName]; exists {
		return relations[relationName]
	}
	return nil
}

// ValidateRelationReferences validates that all relation references in the model are valid.
func ValidateRelationReferences(collector *ErrorCollector, model *openfgav1.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}
	validator := NewSemanticValidator(model)
	for _, typeDef := range model.GetTypeDefinitions() {
		if meta := typeDef.GetMetadata(); meta != nil {
			for relationName, relationMetadata := range meta.GetRelations() {
				validateTypeRestrictions(collector, validator, typeDef.GetType(), relationName, relationMetadata, lines)
			}
		}
		for relationName, userset := range typeDef.GetRelations() {
			validateUsersetReferences(collector, validator, typeDef.GetType(), relationName, userset, lines)
		}
	}
}

func validateTypeRestrictions(collector *ErrorCollector, validator *SemanticValidator,
	typeName, relationName string, relationMetadata *openfgav1.RelationMetadata, lines []string) {
	if relationMetadata == nil {
		return
	}
	meta := &Meta{
		File:   relationMetadata.GetSourceInfo().GetFile(),
		Module: relationMetadata.GetModule(),
	}
	for _, typeRestriction := range relationMetadata.GetDirectlyRelatedUserTypes() {
		if typeRestriction.GetType() == "" {
			continue
		}
		if !validator.TypeDefined(typeRestriction.GetType()) {
			lineIndex := GetRelationLineNumber(relationName, lines, nil)
			collector.RaiseUndefinedType(typeRestriction.GetType(), relationName, typeName, meta, lineIndex)
		}
		if rel := typeRestriction.GetRelation(); rel != "" {
			if !validator.RelationDefined(typeRestriction.GetType(), rel) {
				lineIndex := GetRelationLineNumber(relationName, lines, nil)
				collector.RaiseUndefinedRelation(rel, typeRestriction.GetType(), relationName, typeName, meta, lineIndex)
			}
		}
	}
}

func validateUsersetReferences(collector *ErrorCollector, validator *SemanticValidator,
	typeName, relationName string, userset *openfgav1.Userset, lines []string) {
	if userset == nil {
		return
	}
	var file, module string
	if typeDef := validator.GetTypeDefinition(typeName); typeDef != nil {
		file = typeDef.GetMetadata().GetSourceInfo().GetFile()
		module = typeDef.GetMetadata().GetModule()
	}
	meta := &Meta{File: file, Module: module}

	if cu := userset.GetComputedUserset(); cu != nil {
		if targetRelation := cu.GetRelation(); targetRelation != "" {
			if !validator.RelationDefined(typeName, targetRelation) {
				lineIndex := GetRelationLineNumber(relationName, lines, nil)
				collector.RaiseUndefinedRelation(targetRelation, typeName, relationName, typeName, meta, lineIndex)
			}
		}
	}

	if ttu := userset.GetTupleToUserset(); ttu != nil {
		if targetRelation := ttu.GetComputedUserset().GetRelation(); targetRelation != "" {
			if !validator.RelationDefined(typeName, targetRelation) {
				lineIndex := GetRelationLineNumber(relationName, lines, nil)
				collector.RaiseUndefinedRelation(targetRelation, typeName, relationName, typeName, meta, lineIndex)
			}
		}
		if tuplesetRelation := ttu.GetTupleset().GetRelation(); tuplesetRelation != "" {
			if !validator.RelationDefined(typeName, tuplesetRelation) {
				lineIndex := GetRelationLineNumber(relationName, lines, nil)
				collector.RaiseUndefinedRelation(tuplesetRelation, typeName, relationName, typeName, meta, lineIndex)
			}
		}
	}

	if union := userset.GetUnion(); union != nil {
		for _, child := range union.GetChild() {
			validateUsersetReferences(collector, validator, typeName, relationName, child, lines)
		}
	}
	if intersection := userset.GetIntersection(); intersection != nil {
		for _, child := range intersection.GetChild() {
			validateUsersetReferences(collector, validator, typeName, relationName, child, lines)
		}
	}
	if diff := userset.GetDifference(); diff != nil {
		validateUsersetReferences(collector, validator, typeName, relationName, diff.GetBase(), lines)
		validateUsersetReferences(collector, validator, typeName, relationName, diff.GetSubtract(), lines)
	}
}
