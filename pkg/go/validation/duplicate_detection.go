package validation

import (
	fgaSdk "github.com/openfga/go-sdk"
)

// DuplicateTypeTracker tracks type names to detect duplicates.
type DuplicateTypeTracker struct {
	typeNames map[string]bool
}

// NewDuplicateTypeTracker creates a new duplicate type tracker.
func NewDuplicateTypeTracker() *DuplicateTypeTracker {
	return &DuplicateTypeTracker{
		typeNames: make(map[string]bool),
	}
}

// CheckAndAddType checks if a type name is duplicate and adds it to the tracker.
func (dt *DuplicateTypeTracker) CheckAndAddType(typeName string, collector *ErrorCollector,
	meta *Meta, lines []string) bool {
	if dt.typeNames[typeName] {
		// Type name is duplicate
		typeLineIndex := GetTypeLineNumber(typeName, lines, nil)
		collector.RaiseDuplicateTypeName(typeName, meta, typeLineIndex)
		return false
	}

	// Add type name to tracker
	dt.typeNames[typeName] = true
	return true
}

// CheckForDuplicateTypeNamesInRelation checks for duplicate type names in relation definitions
// This is equivalent to the checkForDuplicatesTypeNamesInRelation function in JS
func CheckForDuplicateTypeNamesInRelation(collector *ErrorCollector, relationMetadata *fgaSdk.RelationMetadata, 
	relationName, typeName string, meta *Meta, lines []string) {
	if relationMetadata == nil || relationMetadata.DirectlyRelatedUserTypes == nil {
		return
	}

	// Track type restrictions to detect duplicates
	typeRestrictions := make(map[string]bool)

	for _, typeRestriction := range *relationMetadata.DirectlyRelatedUserTypes {
		if typeRestriction.Type == "" {
			continue
		}

		// Build type restriction string (type:* or type#relation or type with condition)
		typeRestrictionString := typeRestriction.Type

		// Check for wildcard
		if typeRestriction.Wildcard != nil {
			typeRestrictionString += ":*"
		} else if typeRestriction.Relation != nil && *typeRestriction.Relation != "" {
			typeRestrictionString += "#" + *typeRestriction.Relation
		}

		if typeRestriction.Condition != nil && *typeRestriction.Condition != "" {
			typeRestrictionString += " with " + *typeRestriction.Condition
		}

		// Check for duplicate
		if typeRestrictions[typeRestrictionString] {
			relationLineIndex := GetRelationLineNumber(relationName, lines, nil)
			collector.RaiseDuplicateTypeRestriction(typeRestrictionString, relationName, typeName, meta, relationLineIndex)
		} else {
			typeRestrictions[typeRestrictionString] = true
		}
	}
}

// CheckForDuplicatesInRelation checks for duplicate relations in type definitions
// This is equivalent to the checkForDuplicatesInRelation function in JS
func CheckForDuplicatesInRelation(collector *ErrorCollector, typeDef *fgaSdk.TypeDefinition, 
	relationName string, lines []string) {
	if typeDef == nil || !typeDef.HasRelations() {
		return
	}

	relations := typeDef.GetRelations()
	relation, exists := relations[relationName]
	if !exists {
		return
	}

	// Get file and module metadata
	var file, module string
	if typeDef.Metadata != nil && typeDef.Metadata.HasRelations() {
		relationMetadata := typeDef.Metadata.GetRelations()
		if relationMeta, exists := relationMetadata[relationName]; exists {
			if relationMeta.HasSourceInfo() {
				sourceInfo := relationMeta.GetSourceInfo()
				file = sourceInfo.GetFile()
			}
			if relationMeta.HasModule() {
				module = relationMeta.GetModule()
			}
		}
	}

	// If no file from relation metadata, try type metadata
	if file == "" && typeDef.Metadata != nil && typeDef.Metadata.HasSourceInfo() {
		sourceInfo := typeDef.Metadata.GetSourceInfo()
		file = sourceInfo.GetFile()
	}

	// If no module from relation metadata, try type metadata
	if module == "" && typeDef.Metadata != nil && typeDef.Metadata.HasModule() {
		module = typeDef.Metadata.GetModule()
	}

	meta := &Meta{File: file, Module: module}

	// Check for duplicate operations in complex usersets
	if relation.Union != nil {
		checkDuplicatesInUnion(collector, relation.Union, relationName, typeDef.Type, meta, lines)
	}

	if relation.Intersection != nil {
		checkDuplicatesInIntersection(collector, relation.Intersection, relationName, typeDef.Type, meta, lines)
	}

	if relation.Difference != nil {
		checkDuplicatesInDifference(collector, relation.Difference, relationName, typeDef.Type, meta, lines)
	}
}

// checkDuplicatesInUnion checks for duplicates in union operations.
func checkDuplicatesInUnion(collector *ErrorCollector, union *fgaSdk.Usersets,
	relationName, typeName string, meta *Meta, lines []string) {
	if union == nil {
		return
	}

	// Track relation definitions to detect duplicates
	relationDefs := make(map[string]bool)

	for _, child := range union.Child {
		relationDef := getRelationDefName(child)
		if relationDef != "" {
			if relationDefs[relationDef] {
				relationLineIndex := GetRelationLineNumber(relationName, lines, nil)
				collector.RaiseDuplicateType(relationDef, relationName, typeName, meta, relationLineIndex)
			} else {
				relationDefs[relationDef] = true
			}
		}
	}
}

// checkDuplicatesInIntersection checks for duplicates in intersection operations.
func checkDuplicatesInIntersection(collector *ErrorCollector, intersection *fgaSdk.Usersets,
	relationName, typeName string, meta *Meta, lines []string) {
	if intersection == nil {
		return
	}

	// Track relation definitions to detect duplicates
	relationDefs := make(map[string]bool)

	for _, child := range intersection.Child {
		relationDef := getRelationDefName(child)
		if relationDef != "" {
			if relationDefs[relationDef] {
				relationLineIndex := GetRelationLineNumber(relationName, lines, nil)
				collector.RaiseDuplicateType(relationDef, relationName, typeName, meta, relationLineIndex)
			} else {
				relationDefs[relationDef] = true
			}
		}
	}
}

// checkDuplicatesInDifference checks for duplicates in difference operations
func checkDuplicatesInDifference(collector *ErrorCollector, difference *fgaSdk.Difference, 
	relationName, typeName string, meta *Meta, lines []string) {
	if difference == nil {
		return
	}

	// Check base for duplicates
	baseName := getRelationDefName(difference.Base)
	if baseName != "" {
		// For difference operations, we need to check if base appears multiple times
		// This is a simplified check - more complex logic may be needed
		relationLineIndex := GetRelationLineNumber(relationName, lines, nil)

		// Check if subtract also has the same relation (which would be a duplicate)
		subtractName := getRelationDefName(difference.Subtract)
		if subtractName == baseName {
			collector.RaiseDuplicateType(baseName, relationName, typeName, meta, relationLineIndex)
		}
	}
}

// getRelationDefName extracts the relation definition name from a userset
// This is equivalent to the getRelationDefName function in JS
func getRelationDefName(userset fgaSdk.Userset) string {
	if userset.ComputedUserset != nil && userset.ComputedUserset.HasRelation() {
		return userset.ComputedUserset.GetRelation()
	}

	if userset.TupleToUserset != nil {
		var target, from string

		if userset.TupleToUserset.ComputedUserset.HasRelation() {
			target = userset.TupleToUserset.ComputedUserset.GetRelation()
		}

		if userset.TupleToUserset.Tupleset.HasRelation() {
			from = userset.TupleToUserset.Tupleset.GetRelation()
		}

		if target != "" && from != "" {
			return target + " from " + from
		}

		if target != "" {
			return target
		}
	}

	return ""
}

// ValidateDuplicates performs comprehensive duplicate detection on a model
// This combines all duplicate detection logic
func ValidateDuplicates(collector *ErrorCollector, model *fgaSdk.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}

	// Create duplicate type tracker
	typeTracker := NewDuplicateTypeTracker()

	for _, typeDef := range model.TypeDefinitions {
		if typeDef.Type == "" {
			continue
		}

		typeName := typeDef.Type

		// Get type metadata
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

		// Check for duplicate type names
		typeTracker.CheckAndAddType(typeName, collector, meta, lines)

		// Check for duplicates in relations
		if typeDef.Metadata != nil && typeDef.Metadata.HasRelations() {
			relations := typeDef.Metadata.GetRelations()
			for relationName, relationMetadata := range relations {
				// Check for duplicate type names in relation
				CheckForDuplicateTypeNamesInRelation(collector, &relationMetadata, relationName, typeName, meta, lines)

				// Check for duplicate relations
				CheckForDuplicatesInRelation(collector, &typeDef, relationName, lines)
			}
		}
	}
}
