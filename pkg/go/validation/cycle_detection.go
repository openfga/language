package validation

import (
	"fmt"

	fgaSdk "github.com/openfga/go-sdk"
)

// CycleDetector handles cycle detection and entry point validation
// This is equivalent to the hasEntryPointOrLoop logic in the JS implementation
type CycleDetector struct {
	validator    *SemanticValidator
	visitedNodes map[string]bool
	currentPath  map[string]bool
	entryPoints  map[string]bool
}

// NewCycleDetector creates a new cycle detector for the given semantic validator.
func NewCycleDetector(validator *SemanticValidator) *CycleDetector {
	return &CycleDetector{
		validator:    validator,
		visitedNodes: make(map[string]bool),
		currentPath:  make(map[string]bool),
		entryPoints:  make(map[string]bool),
	}
}

// ValidateCyclesAndEntryPoints validates that all relations have entry points and no cycles
// This is equivalent to the entry point and cycle validation in the JS implementation
func ValidateCyclesAndEntryPoints(collector *ErrorCollector, model *fgaSdk.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}

	validator := NewSemanticValidator(model)
	detector := NewCycleDetector(validator)

	// Check each relation in each type for cycles and entry points
	for _, typeDef := range model.TypeDefinitions {
		if !typeDef.HasRelations() {
			continue
		}

		relations := typeDef.GetRelations()
		for relationName := range relations {
			relationKey := typeDef.Type + "#" + relationName

			// Reset detector state for each relation check
			detector.visitedNodes = make(map[string]bool)
			detector.currentPath = make(map[string]bool)
			detector.entryPoints = make(map[string]bool)

			// Check for cycles and collect entry points
			hasCycle := detector.detectCycle(typeDef.Type, relationName, relationKey)
			hasEntryPoint := detector.hasEntryPoint(typeDef.Type, relationName)

			// Get file and module metadata
			var file, module string
			if typeDef.Metadata != nil {
				if typeDef.Metadata.HasSourceInfo() {
					sourceInfo := typeDef.Metadata.GetSourceInfo()
					file = sourceInfo.GetFile()
				}
				if typeDef.Metadata.HasModule() {
					module = typeDef.Metadata.GetModule()
				}
			}
			meta := &Meta{File: file, Module: module}

			// Report cycle if found
			if hasCycle {
				relationLineIndex := GetRelationLineNumber(relationName, lines, nil)
				collector.RaiseCyclicRelation(relationName, typeDef.Type, meta, relationLineIndex)
			}

			// Report missing entry point if no direct assignment found
			if !hasEntryPoint {
				relationLineIndex := GetRelationLineNumber(relationName, lines, nil)
				collector.RaiseNoEntrypoint(relationName, typeDef.Type, meta, relationLineIndex)
			}
		}
	}
}

// detectCycle detects if there's a cycle in relation definitions using DFS.
func (cd *CycleDetector) detectCycle(typeName, relationName, relationKey string) bool {
	// If we've already visited this node in the current path, we found a cycle
	if cd.currentPath[relationKey] {
		return true
	}

	// If we've already fully processed this node, no cycle from here
	if cd.visitedNodes[relationKey] {
		return false
	}

	// Mark as visited and add to current path
	cd.visitedNodes[relationKey] = true
	cd.currentPath[relationKey] = true

	// Get the userset for this relation
	userset := cd.validator.GetRelationUserset(typeName, relationName)
	if userset != nil {
		// Check all referenced relations for cycles
		if cd.checkUsersetForCycles(typeName, *userset) {
			return true
		}
	}

	// Remove from current path (backtrack)
	delete(cd.currentPath, relationKey)
	return false
}

// checkUsersetForCycles recursively checks userset definitions for cycles.
func (cd *CycleDetector) checkUsersetForCycles(typeName string, userset fgaSdk.Userset) bool {
	// Check computed userset (direct relation reference)
	if userset.ComputedUserset != nil && userset.ComputedUserset.HasRelation() {
		targetRelation := userset.ComputedUserset.GetRelation()
		relationKey := typeName + "#" + targetRelation

		if cd.detectCycle(typeName, targetRelation, relationKey) {
			return true
		}
	}

	// Check tuple-to-userset
	if userset.TupleToUserset != nil {
		// Check the computed userset part
		if userset.TupleToUserset.ComputedUserset.HasRelation() {
			targetRelation := userset.TupleToUserset.ComputedUserset.GetRelation()
			relationKey := typeName + "#" + targetRelation

			if cd.detectCycle(typeName, targetRelation, relationKey) {
				return true
			}
		}

		// Note: tupleset doesn't create cycles in the same type, it references related objects
	}

	// Check union operations
	if userset.Union != nil {
		for _, child := range userset.Union.Child {
			if cd.checkUsersetForCycles(typeName, child) {
				return true
			}
		}
	}

	// Check intersection operations
	if userset.Intersection != nil {
		for _, child := range userset.Intersection.Child {
			if cd.checkUsersetForCycles(typeName, child) {
				return true
			}
		}
	}

	// Check difference operations
	if userset.Difference != nil {
		if cd.checkUsersetForCycles(typeName, userset.Difference.Base) {
			return true
		}
		if cd.checkUsersetForCycles(typeName, userset.Difference.Subtract) {
			return true
		}
	}

	return false
}

// hasEntryPoint checks if a relation has a direct entry point (direct assignment)
// This is equivalent to checking for "this" or direct type restrictions
func (cd *CycleDetector) hasEntryPoint(typeName, relationName string) bool {
	userset := cd.validator.GetRelationUserset(typeName, relationName)
	if userset == nil {
		return false
	}

	return cd.checkUsersetForEntryPoint(*userset, typeName, relationName)
}

// checkUsersetForEntryPoint recursively checks if a userset has an entry point.
func (cd *CycleDetector) checkUsersetForEntryPoint(userset fgaSdk.Userset, typeName, relationName string) bool {
	// Direct assignment via "this" is an entry point
	if userset.This != nil {
		return true
	}

	// Check if there are direct type restrictions (these are entry points)
	if typeDef := cd.validator.GetTypeDefinition(typeName); typeDef != nil {
		if typeDef.Metadata != nil && typeDef.Metadata.HasRelations() {
			relations := typeDef.Metadata.GetRelations()
			if relationMeta, exists := relations[relationName]; exists {
				if relationMeta.HasDirectlyRelatedUserTypes() {
					userTypes := relationMeta.GetDirectlyRelatedUserTypes()
					if len(userTypes) > 0 {
						return true // Direct type restrictions are entry points
					}
				}
			}
		}
	}

	// Check union operations - if any child has an entry point, the union has an entry point
	if userset.Union != nil {
		for _, child := range userset.Union.Child {
			if cd.checkUsersetForEntryPoint(child, typeName, relationName) {
				return true
			}
		}
	}

	// Check intersection operations - all children must have entry points for intersection to have one
	if userset.Intersection != nil {
		hasAllEntryPoints := true
		for _, child := range userset.Intersection.Child {
			if !cd.checkUsersetForEntryPoint(child, typeName, relationName) {
				hasAllEntryPoints = false
				break
			}
		}
		if hasAllEntryPoints && len(userset.Intersection.Child) > 0 {
			return true
		}
	}

	// Check difference operations - base must have entry point
	if userset.Difference != nil {
		return cd.checkUsersetForEntryPoint(userset.Difference.Base, typeName, relationName)
	}

	// Computed userset and tuple-to-userset don't provide direct entry points
	// They rely on other relations having entry points
	return false
}

// Additional error raising methods for cycle detection.
func (ec *ErrorCollector) RaiseCyclicRelation(relationName, typeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("Relation '%s' on type '%s' contains a cycle", relationName, typeName)
	ec.addError(message, CyclicError, relationName, lineIndex, meta, nil)
}

func (ec *ErrorCollector) RaiseNoEntrypoint(relationName, typeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("Relation '%s' on type '%s' has no entry point", relationName, typeName)
	ec.addError(message, RelationNoEntrypoint, relationName, lineIndex, meta, nil)
}
