package validation

import (
	fgaSdk "github.com/openfga/go-sdk"
)

// ComplexOperationValidator handles validation of complex userset operations
// This includes union, intersection, and difference operations with semantic analysis
type ComplexOperationValidator struct {
	model     *fgaSdk.AuthorizationModel
	validator *SemanticValidator
}

// NewComplexOperationValidator creates a new complex operation validator
func NewComplexOperationValidator(model *fgaSdk.AuthorizationModel) *ComplexOperationValidator {
	return &ComplexOperationValidator{
		model:     model,
		validator: NewSemanticValidator(model),
	}
}

// ValidateComplexOperations validates all complex operations in the model
// This includes union, intersection, and difference operations
func ValidateComplexOperations(collector *ErrorCollector, model *fgaSdk.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}
	
	opValidator := NewComplexOperationValidator(model)
	
	// Validate complex operations in each relation
	for _, typeDef := range model.TypeDefinitions {
		if !typeDef.HasRelations() {
			continue
		}
		
		relations := typeDef.GetRelations()
		for relationName, userset := range relations {
			opValidator.validateUsersetOperations(collector, typeDef.Type, relationName, userset, lines)
		}
	}
}

// validateUsersetOperations recursively validates operations in userset definitions
func (cov *ComplexOperationValidator) validateUsersetOperations(collector *ErrorCollector, typeName, relationName string, userset fgaSdk.Userset, lines []string) {
	// Validate union operations
	if userset.Union != nil {
		cov.validateUnionOperation(collector, typeName, relationName, userset.Union, lines)
	}
	
	// Validate intersection operations
	if userset.Intersection != nil {
		cov.validateIntersectionOperation(collector, typeName, relationName, userset.Intersection, lines)
	}
	
	// Validate difference operations
	if userset.Difference != nil {
		cov.validateDifferenceOperation(collector, typeName, relationName, userset.Difference, lines)
	}
	
	// Recursively validate nested operations
	cov.validateNestedOperations(collector, typeName, relationName, userset, lines)
}

// validateUnionOperation validates union operations for semantic correctness
func (cov *ComplexOperationValidator) validateUnionOperation(collector *ErrorCollector, typeName, relationName string, union *fgaSdk.Usersets, lines []string) {
	if union == nil || len(union.Child) == 0 {
		// Empty union is technically valid but might be a modeling issue
		return
	}
	
	// Check for redundant union members
	cov.checkRedundantUnionMembers(collector, typeName, relationName, union, lines)
	
	// Validate each child operation
	for _, child := range union.Child {
		cov.validateUsersetOperations(collector, typeName, relationName, child, lines)
	}
	
	// Check for semantic issues in union
	cov.validateUnionSemantics(collector, typeName, relationName, union, lines)
}

// validateIntersectionOperation validates intersection operations for semantic correctness
func (cov *ComplexOperationValidator) validateIntersectionOperation(collector *ErrorCollector, typeName, relationName string, intersection *fgaSdk.Usersets, lines []string) {
	if intersection == nil || len(intersection.Child) == 0 {
		// Empty intersection is technically valid but might be a modeling issue
		return
	}
	
	// Check for impossible intersections
	cov.checkImpossibleIntersections(collector, typeName, relationName, intersection, lines)
	
	// Validate each child operation
	for _, child := range intersection.Child {
		cov.validateUsersetOperations(collector, typeName, relationName, child, lines)
	}
	
	// Check for semantic issues in intersection
	cov.validateIntersectionSemantics(collector, typeName, relationName, intersection, lines)
}

// validateDifferenceOperation validates difference operations for semantic correctness
func (cov *ComplexOperationValidator) validateDifferenceOperation(collector *ErrorCollector, typeName, relationName string, difference *fgaSdk.Difference, lines []string) {
	if difference == nil {
		return
	}
	
	// Validate base and subtract operations
	cov.validateUsersetOperations(collector, typeName, relationName, difference.Base, lines)
	cov.validateUsersetOperations(collector, typeName, relationName, difference.Subtract, lines)
	
	// Check for semantic issues in difference
	cov.validateDifferenceSemantics(collector, typeName, relationName, difference, lines)
}

// checkRedundantUnionMembers checks for redundant or duplicate members in union operations
func (cov *ComplexOperationValidator) checkRedundantUnionMembers(collector *ErrorCollector, typeName, relationName string, union *fgaSdk.Usersets, lines []string) {
	// Track seen operations to detect duplicates
	seenOperations := make(map[string]bool)
	
	for _, child := range union.Child {
		operationKey := cov.getUsersetOperationKey(child)
		if operationKey != "" {
			if seenOperations[operationKey] {
				// Found duplicate operation in union
				relationLineIndex := GetRelationLineNumber(relationName, lines, nil)
				
				// Get metadata for error reporting
				var file, module string
				if typeDef := cov.validator.GetTypeDefinition(typeName); typeDef != nil && typeDef.Metadata != nil {
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
				
				meta := &Meta{File: file, Module: module}
				collector.RaiseRedundantUnionMember(operationKey, relationName, typeName, meta, relationLineIndex)
			}
			seenOperations[operationKey] = true
		}
	}
}

// checkImpossibleIntersections checks for operations that cannot possibly intersect
func (cov *ComplexOperationValidator) checkImpossibleIntersections(collector *ErrorCollector, typeName, relationName string, intersection *fgaSdk.Usersets, lines []string) {
	// Check for intersections that are logically impossible
	// For example: intersection of two completely disjoint type restrictions
	
	typeRestrictions := make([]string, 0)
	
	for _, child := range intersection.Child {
		if cov.isTypeRestriction(child) {
			restrictionType := cov.getTypeFromRestriction(child)
			if restrictionType != "" {
				typeRestrictions = append(typeRestrictions, restrictionType)
			}
		}
	}
	
	// If we have multiple different type restrictions in intersection, it's impossible
	if len(typeRestrictions) > 1 {
		uniqueTypes := make(map[string]bool)
		for _, t := range typeRestrictions {
			uniqueTypes[t] = true
		}
		
		if len(uniqueTypes) > 1 {
			relationLineIndex := GetRelationLineNumber(relationName, lines, nil)
			
			var file, module string
			if typeDef := cov.validator.GetTypeDefinition(typeName); typeDef != nil && typeDef.Metadata != nil {
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
			
			meta := &Meta{File: file, Module: module}
			collector.RaiseImpossibleIntersection(relationName, typeName, typeRestrictions, meta, relationLineIndex)
		}
	}
}

// validateUnionSemantics performs semantic validation on union operations
func (cov *ComplexOperationValidator) validateUnionSemantics(collector *ErrorCollector, typeName, relationName string, union *fgaSdk.Usersets, lines []string) {
	// Check if union contains operations that subsume others
	// For example: [user:*, user] where user:* already includes user
	cov.checkSubsumingUnionMembers(collector, typeName, relationName, union, lines)
}

// validateIntersectionSemantics performs semantic validation on intersection operations
func (cov *ComplexOperationValidator) validateIntersectionSemantics(collector *ErrorCollector, typeName, relationName string, intersection *fgaSdk.Usersets, lines []string) {
	// Check for semantically redundant intersections
	// For example: intersection with 'this' (which doesn't restrict anything)
	cov.checkRedundantIntersectionMembers(collector, typeName, relationName, intersection, lines)
}

// validateDifferenceSemantics performs semantic validation on difference operations
func (cov *ComplexOperationValidator) validateDifferenceSemantics(collector *ErrorCollector, typeName, relationName string, difference *fgaSdk.Difference, lines []string) {
	// Check if difference base and subtract are the same (results in empty set)
	baseKey := cov.getUsersetOperationKey(difference.Base)
	subtractKey := cov.getUsersetOperationKey(difference.Subtract)
	
	if baseKey != "" && baseKey == subtractKey {
		relationLineIndex := GetRelationLineNumber(relationName, lines, nil)
		
		var file, module string
		if typeDef := cov.validator.GetTypeDefinition(typeName); typeDef != nil && typeDef.Metadata != nil {
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
		
		meta := &Meta{File: file, Module: module}
		collector.RaiseEmptyDifference(relationName, typeName, baseKey, meta, relationLineIndex)
	}
}

// validateNestedOperations recursively validates nested complex operations
func (cov *ComplexOperationValidator) validateNestedOperations(collector *ErrorCollector, typeName, relationName string, userset fgaSdk.Userset, lines []string) {
	// Handle tuple-to-userset operations
	if userset.TupleToUserset != nil {
		// Validate the computed userset part which might contain complex operations
		if userset.TupleToUserset.ComputedUserset.HasRelation() {
			targetRelation := userset.TupleToUserset.ComputedUserset.GetRelation()
			if targetUserset := cov.validator.GetRelationUserset(typeName, targetRelation); targetUserset != nil {
				cov.validateUsersetOperations(collector, typeName, targetRelation, *targetUserset, lines)
			}
		}
	}
}

// Helper methods for operation analysis

// getUsersetOperationKey generates a unique key for a userset operation for comparison
func (cov *ComplexOperationValidator) getUsersetOperationKey(userset fgaSdk.Userset) string {
	if userset.This != nil {
		return "this"
	}
	
	if userset.ComputedUserset != nil && userset.ComputedUserset.HasRelation() {
		return "computed:" + userset.ComputedUserset.GetRelation()
	}
	
	if userset.TupleToUserset != nil {
		var tuplesetRel, computedRel string
		if userset.TupleToUserset.Tupleset.HasRelation() {
			tuplesetRel = userset.TupleToUserset.Tupleset.GetRelation()
		}
		if userset.TupleToUserset.ComputedUserset.HasRelation() {
			computedRel = userset.TupleToUserset.ComputedUserset.GetRelation()
		}
		return "ttu:" + tuplesetRel + ":" + computedRel
	}
	
	// For complex nested operations, return empty string (requires more sophisticated comparison)
	return ""
}

// isTypeRestriction checks if a userset represents a type restriction
func (cov *ComplexOperationValidator) isTypeRestriction(userset fgaSdk.Userset) bool {
	// Type restrictions are typically represented as 'this' operations
	// in the context of relation metadata DirectlyRelatedUserTypes
	return userset.This != nil
}

// getTypeFromRestriction extracts the type name from a type restriction
func (cov *ComplexOperationValidator) getTypeFromRestriction(userset fgaSdk.Userset) string {
	// This would need to be enhanced based on how type restrictions are represented
	// in the specific userset context
	return ""
}

// checkSubsumingUnionMembers checks for union members that subsume others
func (cov *ComplexOperationValidator) checkSubsumingUnionMembers(collector *ErrorCollector, typeName, relationName string, union *fgaSdk.Usersets, lines []string) {
	// Implementation would check for cases like [user:*, user] where the wildcard subsumes the specific relation
	// This requires detailed analysis of type restrictions which would be implemented based on specific requirements
}

// checkRedundantIntersectionMembers checks for redundant members in intersection operations
func (cov *ComplexOperationValidator) checkRedundantIntersectionMembers(collector *ErrorCollector, typeName, relationName string, intersection *fgaSdk.Usersets, lines []string) {
	// Implementation would check for intersection members that don't actually restrict the result
	// For example, intersecting with 'this' which doesn't add any restriction
}
