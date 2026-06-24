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

// GetRelationNames returns the names of every relation defined on a type.
func (sv *SemanticValidator) GetRelationNames(typeName string) []string {
	relations := sv.relationMap[typeName]
	names := make([]string, 0, len(relations))
	for name := range relations {
		names = append(names, name)
	}
	return names
}

// GetDirectlyAssignableTypes returns the type names a relation is directly
// assignable to, but only when that relation is a single direct assignment
// (i.e. `define r: [a, b]` rather than a rewrite). The bool reports whether the
// relation is such a single direct assignment. This mirrors the reference
// implementation's allowableTypes helper used for tuple-to-userset validation.
func (sv *SemanticValidator) GetDirectlyAssignableTypes(typeName, relationName string) ([]string, bool) {
	userset := sv.GetRelationUserset(typeName, relationName)
	if userset == nil || userset.GetThis() == nil {
		return nil, false
	}
	typeDef := sv.typeMap[typeName]
	if typeDef == nil {
		return nil, false
	}
	relMeta := typeDef.GetMetadata().GetRelations()[relationName]
	types := make([]string, 0)
	for _, tr := range relMeta.GetDirectlyRelatedUserTypes() {
		if tr.GetType() != "" {
			types = append(types, tr.GetType())
		}
	}
	return types, true
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
		restrictedType := typeRestriction.GetType()
		if restrictedType == "" {
			continue
		}
		// A directly-related type that doesn't exist: `X` is not a valid type.
		if !validator.TypeDefined(restrictedType) {
			lineIndex := GetRelationLineNumber(relationName, lines, nil)
			collector.RaiseInvalidType(restrictedType, typeName, relationName, meta, lineIndex)
			continue
		}
		// A type#relation restriction whose relation doesn't exist on that type:
		// `rel` is not a valid relation for `X`.
		if rel := typeRestriction.GetRelation(); rel != "" {
			if !validator.RelationDefined(restrictedType, rel) {
				lineIndex := GetRelationLineNumber(relationName, lines, nil)
				symbol := restrictedType + "#" + rel
				collector.RaiseInvalidTypeRelation(symbol, restrictedType, relationName, rel, restrictedType, lineIndex, meta)
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
		// `define a: b` where b is not a relation on this type.
		if targetRelation := cu.GetRelation(); targetRelation != "" {
			if !validator.RelationDefined(typeName, targetRelation) {
				lineIndex := GetRelationLineNumber(relationName, lines, nil)
				validRelations := validator.GetRelationNames(typeName)
				collector.RaiseInvalidRelationError(targetRelation, typeName, relationName, validRelations, lineIndex, meta)
			}
		}
	}

	if ttu := userset.GetTupleToUserset(); ttu != nil {
		validateTupleToUsersetReferences(collector, validator, typeName, relationName, ttu, meta, lines)
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

// validateTupleToUsersetReferences validates a `target from from` rewrite,
// mirroring the reference implementation:
//   - the `from` (tupleset) relation must exist on the current type;
//   - the `from` relation must be a plain direct assignment whose assignable
//     types are concrete (no wildcard, no type#relation);
//   - the computed `target` relation must exist on at least one of the types the
//     `from` relation is assignable to.
func validateTupleToUsersetReferences(collector *ErrorCollector, validator *SemanticValidator,
	typeName, relationName string, ttu *openfgav1.TupleToUserset, meta *Meta, lines []string) {
	fromRelation := ttu.GetTupleset().GetRelation()
	targetRelation := ttu.GetComputedUserset().GetRelation()
	if fromRelation == "" || targetRelation == "" {
		return
	}
	lineIndex := GetRelationLineNumber(relationName, lines, nil)
	symbol := targetRelation + " from " + fromRelation

	// 1. The `from` relation must exist on the current type.
	if !validator.RelationDefined(typeName, fromRelation) {
		collector.RaiseInvalidTypeRelation(symbol, typeName, relationName, fromRelation, typeName, lineIndex, meta)
		return
	}

	// 2. The `from` relation must be a single direct assignment.
	fromTypes, isValid := validator.GetDirectlyAssignableTypes(typeName, fromRelation)
	if !isValid || len(fromTypes) == 0 {
		collector.RaiseTupleUsersetRequiresDirect(fromRelation, typeName, relationName, meta, lineIndex)
		return
	}

	// 3. The computed `target` must exist on at least one assignable type.
	notValid := 0
	for _, fromType := range fromTypes {
		if !validator.TypeDefined(fromType) || !validator.RelationDefined(fromType, targetRelation) {
			notValid++
		}
	}
	if notValid == len(fromTypes) {
		for _, fromType := range fromTypes {
			collector.RaiseInvalidRelationOnTupleset(symbol, fromType, typeName, relationName, targetRelation, fromRelation, lineIndex, meta)
		}
	}
}
