package validation

import (
	"maps"
	"slices"

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
			// Relations in name order: the references collected here are reported in
			// the order they were appended, so ranging the map would vary it.
			relationsMetadata := metaProto.GetRelations()
			for _, relationName := range slices.Sorted(maps.Keys(relationsMetadata)) {
				cv.scanRelationMetadataForConditions(typeDef.GetType(), relationName,
					relationsMetadata[relationName])
			}
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

// ValidateUnusedConditions detects and reports unused condition definitions.
func ValidateUnusedConditions(errs *ValidationErrors, model *openfgav1.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}
	validateUnusedConditions(errs, NewConditionValidator(model), lines)
}

func validateUnusedConditions(errs *ValidationErrors, validator *ConditionValidator, lines []string) {
	for _, conditionName := range slices.Sorted(maps.Keys(validator.definedConds)) {
		if !validator.usedConds[conditionName] {
			condition := validator.definedConds[conditionName]
			lineIndex := GetConditionLineNumber(conditionName, lines, nil)
			meta := &Meta{
				File:   condition.GetMetadata().GetSourceInfo().GetFile(),
				Module: condition.GetMetadata().GetModule(),
			}
			errs.Add(newUnusedConditionError(lines, conditionName, meta, lineIndex))
		}
	}
}

// ValidateConditionReferences validates that all referenced conditions are defined.
func ValidateConditionReferences(errs *ValidationErrors, model *openfgav1.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}
	validateConditionReferences(errs, NewConditionValidator(model), lines)
}

func validateConditionReferences(errs *ValidationErrors, validator *ConditionValidator, lines []string) {
	model := validator.model
	for _, conditionName := range slices.Sorted(maps.Keys(validator.usedConds)) {
		if _, exists := validator.definedConds[conditionName]; !exists {
			for _, ref := range validator.conditionRefs[conditionName] {
				// Anchor the relation line lookup to the referencing type's
				// declaration so the correct `define` is found when several types
				// share a relation name, matching the reference.
				typeLineIndex := GetTypeLineNumber(ref.TypeName, lines, nil)
				lineIndex := GetRelationLineNumber(ref.RelationName, lines, typeLineIndex)
				var file, module string
				for _, typeDef := range model.GetTypeDefinitions() {
					if typeDef.GetType() == ref.TypeName {
						file = typeDef.GetMetadata().GetSourceInfo().GetFile()
						module = typeDef.GetMetadata().GetModule()
						break
					}
				}
				meta := &Meta{File: file, Module: module}
				errs.Add(newInvalidConditionNameInParameterError(lines, conditionName, ref.TypeName, ref.RelationName, conditionName, meta, lineIndex))
			}
		}
	}
}

// ValidateConditionConsistency checks that each condition's nested name property
// matches its map key, mirroring the reference (validate-dsl.ts): the nested name
// is compared to the key and any difference is reported.
func ValidateConditionConsistency(errs *ValidationErrors, model *openfgav1.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}
	conditions := model.GetConditions()
	for _, conditionKey := range slices.Sorted(maps.Keys(conditions)) {
		condition := conditions[conditionKey]
		if condition == nil {
			continue
		}
		if condition.GetName() != conditionKey {
			errs.Add(newDifferentNestedConditionNameError(conditionKey, condition.GetName()))
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
