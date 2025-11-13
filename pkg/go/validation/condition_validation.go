package validation

import (
	fgaSdk "github.com/openfga/go-sdk"
)

// ConditionValidator handles condition-related validation
// This includes unused condition detection and condition consistency validation
type ConditionValidator struct {
	model         *fgaSdk.AuthorizationModel
	definedConds  map[string]*fgaSdk.Condition    // conditions defined in the model
	usedConds     map[string]bool                 // conditions referenced in relations
	conditionRefs map[string][]ConditionReference // where each condition is used
}

// ConditionReference tracks where a condition is referenced
type ConditionReference struct {
	TypeName     string
	RelationName string
	Context      string // e.g., "type_restriction", "computed_userset"
}

// NewConditionValidator creates a new condition validator for the given model
func NewConditionValidator(model *fgaSdk.AuthorizationModel) *ConditionValidator {
	validator := &ConditionValidator{
		model:         model,
		definedConds:  make(map[string]*fgaSdk.Condition),
		usedConds:     make(map[string]bool),
		conditionRefs: make(map[string][]ConditionReference),
	}
	
	validator.buildConditionMaps()
	return validator
}

// buildConditionMaps constructs internal maps for condition tracking
func (cv *ConditionValidator) buildConditionMaps() {
	if cv.model == nil {
		return
	}
	
	// Build map of defined conditions
	if cv.model.Conditions != nil {
		for conditionName, condition := range *cv.model.Conditions {
			cv.definedConds[conditionName] = &condition
		}
	}
	
	// Scan model for condition usage
	cv.scanForConditionUsage()
}

// scanForConditionUsage scans the model to track condition usage
func (cv *ConditionValidator) scanForConditionUsage() {
	for _, typeDef := range cv.model.TypeDefinitions {
		// Check conditions in relation metadata (type restrictions)
		if typeDef.Metadata != nil && typeDef.Metadata.HasRelations() {
			relations := typeDef.Metadata.GetRelations()
			for relationName, relationMetadata := range relations {
				cv.scanRelationMetadataForConditions(typeDef.Type, relationName, relationMetadata)
			}
		}
		
		// Check conditions in userset definitions
		if typeDef.HasRelations() {
			relations := typeDef.GetRelations()
			for relationName, userset := range relations {
				cv.scanUsersetForConditions(typeDef.Type, relationName, userset)
			}
		}
	}
}

// scanRelationMetadataForConditions scans relation metadata for condition references
func (cv *ConditionValidator) scanRelationMetadataForConditions(typeName, relationName string, relationMetadata fgaSdk.RelationMetadata) {
	if !relationMetadata.HasDirectlyRelatedUserTypes() {
		return
	}
	
	userTypes := relationMetadata.GetDirectlyRelatedUserTypes()
	for _, typeRestriction := range userTypes {
		if typeRestriction.Condition != nil && *typeRestriction.Condition != "" {
			conditionName := *typeRestriction.Condition
			cv.usedConds[conditionName] = true
			
			ref := ConditionReference{
				TypeName:     typeName,
				RelationName: relationName,
				Context:      "type_restriction",
			}
			cv.conditionRefs[conditionName] = append(cv.conditionRefs[conditionName], ref)
		}
	}
}

// scanUsersetForConditions recursively scans userset definitions for condition references
func (cv *ConditionValidator) scanUsersetForConditions(typeName, relationName string, userset fgaSdk.Userset) {
	// Check computed userset conditions
	if userset.ComputedUserset != nil {
		// Note: computed usersets typically don't have direct condition references
		// but we scan for completeness
	}
	
	// Check tuple-to-userset conditions
	if userset.TupleToUserset != nil {
		// Scan both tupleset and computed userset parts
		if userset.TupleToUserset.Tupleset.HasRelation() {
			// Tupleset relations can have conditions in their definitions
			// This would be handled when we process that relation's metadata
		}
	}
	
	// Recursively scan union operations
	if userset.Union != nil {
		for _, child := range userset.Union.Child {
			cv.scanUsersetForConditions(typeName, relationName, child)
		}
	}
	
	// Recursively scan intersection operations
	if userset.Intersection != nil {
		for _, child := range userset.Intersection.Child {
			cv.scanUsersetForConditions(typeName, relationName, child)
		}
	}
	
	// Recursively scan difference operations
	if userset.Difference != nil {
		cv.scanUsersetForConditions(typeName, relationName, userset.Difference.Base)
		cv.scanUsersetForConditions(typeName, relationName, userset.Difference.Subtract)
	}
}

// ValidateUnusedConditions detects and reports unused condition definitions
func ValidateUnusedConditions(collector *ErrorCollector, model *fgaSdk.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}
	
	validator := NewConditionValidator(model)
	
	// Check each defined condition
	for conditionName, condition := range validator.definedConds {
		if !validator.usedConds[conditionName] {
			// Condition is defined but not used
			conditionLineIndex := GetConditionLineNumber(conditionName, lines, nil)
			
			// Get metadata for error reporting
			var file, module string
			if condition.Metadata != nil {
				if condition.Metadata.HasSourceInfo() {
					sourceInfo := condition.Metadata.GetSourceInfo()
					if sourceInfo.HasFile() {
						file = sourceInfo.GetFile()
					}
				}
				if condition.Metadata.HasModule() {
					module = condition.Metadata.GetModule()
				}
			}
			
			meta := &Meta{File: file, Module: module}
			collector.RaiseUnusedCondition(conditionName, meta, conditionLineIndex)
		}
	}
}

// ValidateConditionReferences validates that all referenced conditions are defined
func ValidateConditionReferences(collector *ErrorCollector, model *fgaSdk.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}
	
	validator := NewConditionValidator(model)
	
	// Check each used condition to ensure it's defined
	for conditionName := range validator.usedConds {
		if _, exists := validator.definedConds[conditionName]; !exists {
			// Condition is used but not defined
			refs := validator.conditionRefs[conditionName]
			for _, ref := range refs {
				relationLineIndex := GetRelationLineNumber(ref.RelationName, lines, nil)
				
				// Get metadata from the type that references the undefined condition
				var file, module string
				for _, typeDef := range model.TypeDefinitions {
					if typeDef.Type == ref.TypeName {
						if typeDef.Metadata != nil {
							if typeDef.Metadata.HasSourceInfo() {
								sourceInfo := typeDef.Metadata.GetSourceInfo()
								if sourceInfo.HasFile() {
									file = sourceInfo.GetFile()
								}
							}
							if typeDef.Metadata.HasModule() {
								module = typeDef.Metadata.GetModule()
							}
						}
						break
					}
				}
				
				meta := &Meta{File: file, Module: module}
				collector.RaiseInvalidConditionNameInParameter(conditionName, ref.TypeName, ref.RelationName, conditionName, meta, relationLineIndex)
			}
		}
	}
}

// ValidateConditionConsistency validates condition name consistency
// This checks for conditions with mismatched names in nested structures
func ValidateConditionConsistency(collector *ErrorCollector, model *fgaSdk.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}
	
	// Check each condition definition for internal consistency
	if model.Conditions != nil {
		for _, condition := range *model.Conditions {
		// Check if condition has nested condition references with different names
		// This is a more complex validation that would need to inspect condition expressions
		// For now, we implement a basic version that checks the condition name itself
		
		if condition.Name == "" {
			// Anonymous condition - this might be invalid depending on schema version
			_ = GetConditionLineNumber("", lines, nil)
			collector.RaiseDifferentNestedConditionName("", "anonymous_condition")
		}
		}
	}
}

// GetDefinedConditions returns all condition names defined in the model
func (cv *ConditionValidator) GetDefinedConditions() []string {
	conditions := make([]string, 0, len(cv.definedConds))
	for name := range cv.definedConds {
		conditions = append(conditions, name)
	}
	return conditions
}

// GetUsedConditions returns all condition names used in the model
func (cv *ConditionValidator) GetUsedConditions() []string {
	conditions := make([]string, 0, len(cv.usedConds))
	for name := range cv.usedConds {
		conditions = append(conditions, name)
	}
	return conditions
}

// IsConditionDefined checks if a condition is defined in the model
func (cv *ConditionValidator) IsConditionDefined(conditionName string) bool {
	_, exists := cv.definedConds[conditionName]
	return exists
}

// IsConditionUsed checks if a condition is used in the model
func (cv *ConditionValidator) IsConditionUsed(conditionName string) bool {
	return cv.usedConds[conditionName]
}

// GetConditionReferences returns all references to a specific condition
func (cv *ConditionValidator) GetConditionReferences(conditionName string) []ConditionReference {
	return cv.conditionRefs[conditionName]
}
