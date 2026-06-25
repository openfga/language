package validation

import (
	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// ConditionValidator handles condition-related validation.
type ConditionValidator struct {
	model         *openfgav1.AuthorizationModel
	definedConds  map[string]*openfgav1.Condition
	usedConds     map[string]bool
	conditionRefs map[string][]ConditionReference
}

// ConditionReference tracks where a condition is referenced.
type ConditionReference struct {
	TypeName     string
	RelationName string
	Context      string
}

func NewConditionValidator(model *openfgav1.AuthorizationModel) *ConditionValidator {
	validator := &ConditionValidator{
		model:         model,
		definedConds:  make(map[string]*openfgav1.Condition),
		usedConds:     make(map[string]bool),
		conditionRefs: make(map[string][]ConditionReference),
	}
	validator.buildConditionMaps()
	return validator
}

func (cv *ConditionValidator) buildConditionMaps() {
	if cv.model == nil {
		return
	}
	for conditionName, condition := range cv.model.GetConditions() {
		cv.definedConds[conditionName] = condition
	}
	cv.scanForConditionUsage()
}

func (cv *ConditionValidator) scanForConditionUsage() {
	for _, typeDef := range cv.model.GetTypeDefinitions() {
		if metaProto := typeDef.GetMetadata(); metaProto != nil {
			for relationName, relationMetadata := range metaProto.GetRelations() {
				cv.scanRelationMetadataForConditions(typeDef.GetType(), relationName, relationMetadata)
			}
		}
		for relationName, userset := range typeDef.GetRelations() {
			cv.scanUsersetForConditions(typeDef.GetType(), relationName, userset)
		}
	}
}

func (cv *ConditionValidator) scanRelationMetadataForConditions(typeName, relationName string, rm *openfgav1.RelationMetadata) {
	if rm == nil {
		return
	}
	for _, typeRestriction := range rm.GetDirectlyRelatedUserTypes() {
		if cond := typeRestriction.GetCondition(); cond != "" {
			cv.usedConds[cond] = true
			cv.conditionRefs[cond] = append(cv.conditionRefs[cond], ConditionReference{
				TypeName:     typeName,
				RelationName: relationName,
				Context:      "type_restriction",
			})
		}
	}
}

func (cv *ConditionValidator) scanUsersetForConditions(typeName, relationName string, userset *openfgav1.Userset) {
	if userset == nil {
		return
	}
	if union := userset.GetUnion(); union != nil {
		for _, child := range union.GetChild() {
			cv.scanUsersetForConditions(typeName, relationName, child)
		}
	}
	if intersection := userset.GetIntersection(); intersection != nil {
		for _, child := range intersection.GetChild() {
			cv.scanUsersetForConditions(typeName, relationName, child)
		}
	}
	if diff := userset.GetDifference(); diff != nil {
		cv.scanUsersetForConditions(typeName, relationName, diff.GetBase())
		cv.scanUsersetForConditions(typeName, relationName, diff.GetSubtract())
	}
}

// ValidateUnusedConditions detects and reports unused condition definitions.
func ValidateUnusedConditions(collector *ErrorCollector, model *openfgav1.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}
	validateUnusedConditions(collector, NewConditionValidator(model), lines)
}

func validateUnusedConditions(collector *ErrorCollector, validator *ConditionValidator, lines []string) {
	for conditionName, condition := range validator.definedConds {
		if !validator.usedConds[conditionName] {
			lineIndex := GetConditionLineNumber(conditionName, lines, nil)
			meta := &Meta{
				File:   condition.GetMetadata().GetSourceInfo().GetFile(),
				Module: condition.GetMetadata().GetModule(),
			}
			collector.RaiseUnusedCondition(conditionName, meta, lineIndex)
		}
	}
}

// ValidateConditionReferences validates that all referenced conditions are defined.
func ValidateConditionReferences(collector *ErrorCollector, model *openfgav1.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}
	validateConditionReferences(collector, NewConditionValidator(model), lines)
}

func validateConditionReferences(collector *ErrorCollector, validator *ConditionValidator, lines []string) {
	model := validator.model
	for conditionName := range validator.usedConds {
		if _, exists := validator.definedConds[conditionName]; !exists {
			for _, ref := range validator.conditionRefs[conditionName] {
				lineIndex := GetRelationLineNumber(ref.RelationName, lines, nil)
				var file, module string
				for _, typeDef := range model.GetTypeDefinitions() {
					if typeDef.GetType() == ref.TypeName {
						file = typeDef.GetMetadata().GetSourceInfo().GetFile()
						module = typeDef.GetMetadata().GetModule()
						break
					}
				}
				meta := &Meta{File: file, Module: module}
				collector.RaiseInvalidConditionNameInParameter(conditionName, ref.TypeName, ref.RelationName, conditionName, meta, lineIndex)
			}
		}
	}
}

// ValidateConditionConsistency validates condition name consistency.
func ValidateConditionConsistency(collector *ErrorCollector, model *openfgav1.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}
	for _, condition := range model.GetConditions() {
		if condition.GetName() == "" {
			collector.RaiseDifferentNestedConditionName("", "anonymous_condition")
		}
	}
}

func (cv *ConditionValidator) GetDefinedConditions() []string {
	conditions := make([]string, 0, len(cv.definedConds))
	for name := range cv.definedConds {
		conditions = append(conditions, name)
	}
	return conditions
}

func (cv *ConditionValidator) GetUsedConditions() []string {
	conditions := make([]string, 0, len(cv.usedConds))
	for name := range cv.usedConds {
		conditions = append(conditions, name)
	}
	return conditions
}

func (cv *ConditionValidator) IsConditionDefined(conditionName string) bool {
	_, exists := cv.definedConds[conditionName]
	return exists
}

func (cv *ConditionValidator) IsConditionUsed(conditionName string) bool {
	return cv.usedConds[conditionName]
}

func (cv *ConditionValidator) GetConditionReferences(conditionName string) []ConditionReference {
	return cv.conditionRefs[conditionName]
}
