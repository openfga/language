package validation

import (
	"fmt"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// CycleDetector handles cycle detection and entry point validation.
type CycleDetector struct {
	validator    *SemanticValidator
	visitedNodes map[string]bool
	currentPath  map[string]bool
	entryPoints  map[string]bool
}

func NewCycleDetector(validator *SemanticValidator) *CycleDetector {
	return &CycleDetector{
		validator:    validator,
		visitedNodes: make(map[string]bool),
		currentPath:  make(map[string]bool),
		entryPoints:  make(map[string]bool),
	}
}

// ValidateCyclesAndEntryPoints validates that all relations have entry points and no cycles.
func ValidateCyclesAndEntryPoints(collector *ErrorCollector, model *openfgav1.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}
	validator := NewSemanticValidator(model)
	detector := NewCycleDetector(validator)

	for _, typeDef := range model.GetTypeDefinitions() {
		relations := typeDef.GetRelations()
		if len(relations) == 0 {
			continue
		}
		meta := &Meta{
			File:   typeDef.GetMetadata().GetSourceInfo().GetFile(),
			Module: typeDef.GetMetadata().GetModule(),
		}
		for relationName := range relations {
			relationKey := typeDef.GetType() + "#" + relationName
			detector.visitedNodes = make(map[string]bool)
			detector.currentPath = make(map[string]bool)
			detector.entryPoints = make(map[string]bool)

			hasCycle := detector.detectCycle(typeDef.GetType(), relationName, relationKey)
			hasEntryPoint := detector.hasEntryPoint(typeDef.GetType(), relationName)

			if hasCycle {
				lineIndex := GetRelationLineNumber(relationName, lines, nil)
				collector.RaiseCyclicRelation(relationName, typeDef.GetType(), meta, lineIndex)
			}
			if !hasEntryPoint {
				lineIndex := GetRelationLineNumber(relationName, lines, nil)
				collector.RaiseNoEntrypoint(relationName, typeDef.GetType(), meta, lineIndex)
			}
		}
	}
}

func (cd *CycleDetector) detectCycle(typeName, relationName, relationKey string) bool {
	if cd.currentPath[relationKey] {
		return true
	}
	if cd.visitedNodes[relationKey] {
		return false
	}
	cd.visitedNodes[relationKey] = true
	cd.currentPath[relationKey] = true

	if userset := cd.validator.GetRelationUserset(typeName, relationName); userset != nil {
		if cd.checkUsersetForCycles(typeName, userset) {
			return true
		}
	}
	delete(cd.currentPath, relationKey)
	return false
}

func (cd *CycleDetector) checkUsersetForCycles(typeName string, userset *openfgav1.Userset) bool {
	if userset == nil {
		return false
	}
	if cu := userset.GetComputedUserset(); cu != nil {
		if targetRelation := cu.GetRelation(); targetRelation != "" {
			if cd.detectCycle(typeName, targetRelation, typeName+"#"+targetRelation) {
				return true
			}
		}
	}
	if ttu := userset.GetTupleToUserset(); ttu != nil {
		if targetRelation := ttu.GetComputedUserset().GetRelation(); targetRelation != "" {
			if cd.detectCycle(typeName, targetRelation, typeName+"#"+targetRelation) {
				return true
			}
		}
	}
	if union := userset.GetUnion(); union != nil {
		for _, child := range union.GetChild() {
			if cd.checkUsersetForCycles(typeName, child) {
				return true
			}
		}
	}
	if intersection := userset.GetIntersection(); intersection != nil {
		for _, child := range intersection.GetChild() {
			if cd.checkUsersetForCycles(typeName, child) {
				return true
			}
		}
	}
	if diff := userset.GetDifference(); diff != nil {
		if cd.checkUsersetForCycles(typeName, diff.GetBase()) {
			return true
		}
		if cd.checkUsersetForCycles(typeName, diff.GetSubtract()) {
			return true
		}
	}
	return false
}

func (cd *CycleDetector) hasEntryPoint(typeName, relationName string) bool {
	userset := cd.validator.GetRelationUserset(typeName, relationName)
	if userset == nil {
		return false
	}
	return cd.checkUsersetForEntryPoint(userset, typeName, relationName)
}

func (cd *CycleDetector) checkUsersetForEntryPoint(userset *openfgav1.Userset, typeName, relationName string) bool {
	if userset == nil {
		return false
	}
	if userset.GetThis() != nil {
		return true
	}
	if typeDef := cd.validator.GetTypeDefinition(typeName); typeDef != nil {
		if rm, ok := typeDef.GetMetadata().GetRelations()[relationName]; ok {
			if len(rm.GetDirectlyRelatedUserTypes()) > 0 {
				return true
			}
		}
	}
	if union := userset.GetUnion(); union != nil {
		for _, child := range union.GetChild() {
			if cd.checkUsersetForEntryPoint(child, typeName, relationName) {
				return true
			}
		}
	}
	if intersection := userset.GetIntersection(); intersection != nil {
		if len(intersection.GetChild()) == 0 {
			return false
		}
		for _, child := range intersection.GetChild() {
			if !cd.checkUsersetForEntryPoint(child, typeName, relationName) {
				return false
			}
		}
		return true
	}
	if diff := userset.GetDifference(); diff != nil {
		return cd.checkUsersetForEntryPoint(diff.GetBase(), typeName, relationName)
	}
	return false
}

func (ec *ErrorCollector) RaiseCyclicRelation(relationName, typeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("Relation '%s' on type '%s' contains a cycle", relationName, typeName)
	ec.addError(message, CyclicError, relationName, lineIndex, meta, nil)
}

func (ec *ErrorCollector) RaiseNoEntrypoint(relationName, typeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("Relation '%s' on type '%s' has no entry point", relationName, typeName)
	ec.addError(message, RelationNoEntrypoint, relationName, lineIndex, meta, nil)
}
