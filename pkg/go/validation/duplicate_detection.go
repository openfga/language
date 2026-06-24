package validation

import (
	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// DuplicateTypeTracker tracks type names to detect duplicates.
type DuplicateTypeTracker struct {
	typeNames map[string]bool
}

func NewDuplicateTypeTracker() *DuplicateTypeTracker {
	return &DuplicateTypeTracker{typeNames: make(map[string]bool)}
}

func (dt *DuplicateTypeTracker) CheckAndAddType(typeName string, collector *ErrorCollector,
	meta *Meta, lines []string) bool {
	if dt.typeNames[typeName] {
		typeLineIndex := GetTypeLineNumber(typeName, lines, nil)
		collector.RaiseDuplicateTypeName(typeName, meta, typeLineIndex)
		return false
	}
	dt.typeNames[typeName] = true
	return true
}

// CheckForDuplicateTypeNamesInRelation checks for duplicate type restrictions within a relation.
func CheckForDuplicateTypeNamesInRelation(collector *ErrorCollector, relationMetadata *openfgav1.RelationMetadata,
	relationName, typeName string, meta *Meta, typeLineIndex *int, lines []string) {
	if relationMetadata == nil {
		return
	}
	typeRestrictions := make(map[string]bool)
	for _, typeRestriction := range relationMetadata.GetDirectlyRelatedUserTypes() {
		if typeRestriction.GetType() == "" {
			continue
		}
		typeRestrictionString := typeRestriction.GetType()
		if typeRestriction.GetWildcard() != nil {
			typeRestrictionString += ":*"
		} else if rel := typeRestriction.GetRelation(); rel != "" {
			typeRestrictionString += "#" + rel
		}
		if cond := typeRestriction.GetCondition(); cond != "" {
			typeRestrictionString += " with " + cond
		}
		if typeRestrictions[typeRestrictionString] {
			lineIndex := GetRelationLineNumber(relationName, lines, typeLineIndex)
			collector.RaiseDuplicateTypeRestriction(typeRestrictionString, relationName, typeName, meta, lineIndex)
		} else {
			typeRestrictions[typeRestrictionString] = true
		}
	}
}

// CheckForDuplicatesInRelation checks for duplicate relations in type definitions.
func CheckForDuplicatesInRelation(collector *ErrorCollector, typeDef *openfgav1.TypeDefinition,
	relationName string, typeLineIndex *int, lines []string) {
	if typeDef == nil {
		return
	}
	relations := typeDef.GetRelations()
	relation, exists := relations[relationName]
	if !exists {
		return
	}

	var file, module string
	if meta := typeDef.GetMetadata(); meta != nil {
		if rm, ok := meta.GetRelations()[relationName]; ok {
			file = rm.GetSourceInfo().GetFile()
			module = rm.GetModule()
		}
		if file == "" {
			file = meta.GetSourceInfo().GetFile()
		}
		if module == "" {
			module = meta.GetModule()
		}
	}
	meta := &Meta{File: file, Module: module}

	if union := relation.GetUnion(); union != nil {
		checkDuplicatesInUnion(collector, union, relationName, typeDef.GetType(), meta, typeLineIndex, lines)
	}
	if intersection := relation.GetIntersection(); intersection != nil {
		checkDuplicatesInIntersection(collector, intersection, relationName, typeDef.GetType(), meta, typeLineIndex, lines)
	}
	if diff := relation.GetDifference(); diff != nil {
		checkDuplicatesInDifference(collector, diff, relationName, typeDef.GetType(), meta, typeLineIndex, lines)
	}
}

func checkDuplicatesInUnion(collector *ErrorCollector, union *openfgav1.Usersets,
	relationName, typeName string, meta *Meta, typeLineIndex *int, lines []string) {
	if union == nil {
		return
	}
	relationDefs := make(map[string]bool)
	for _, child := range union.GetChild() {
		if relationDef := getRelationDefName(child); relationDef != "" {
			if relationDefs[relationDef] {
				lineIndex := GetRelationLineNumber(relationName, lines, typeLineIndex)
				collector.RaiseDuplicateType(relationDef, relationName, typeName, meta, lineIndex)
			} else {
				relationDefs[relationDef] = true
			}
		}
	}
}

func checkDuplicatesInIntersection(collector *ErrorCollector, intersection *openfgav1.Usersets,
	relationName, typeName string, meta *Meta, typeLineIndex *int, lines []string) {
	if intersection == nil {
		return
	}
	relationDefs := make(map[string]bool)
	for _, child := range intersection.GetChild() {
		if relationDef := getRelationDefName(child); relationDef != "" {
			if relationDefs[relationDef] {
				lineIndex := GetRelationLineNumber(relationName, lines, typeLineIndex)
				collector.RaiseDuplicateType(relationDef, relationName, typeName, meta, lineIndex)
			} else {
				relationDefs[relationDef] = true
			}
		}
	}
}

func checkDuplicatesInDifference(collector *ErrorCollector, difference *openfgav1.Difference,
	relationName, typeName string, meta *Meta, typeLineIndex *int, lines []string) {
	if difference == nil {
		return
	}
	baseName := getRelationDefName(difference.GetBase())
	subtractName := getRelationDefName(difference.GetSubtract())
	if baseName != "" && baseName == subtractName {
		lineIndex := GetRelationLineNumber(relationName, lines, typeLineIndex)
		collector.RaiseDuplicateType(baseName, relationName, typeName, meta, lineIndex)
	}
}

func getRelationDefName(userset *openfgav1.Userset) string {
	if userset == nil {
		return ""
	}
	if cu := userset.GetComputedUserset(); cu != nil {
		if rel := cu.GetRelation(); rel != "" {
			return rel
		}
	}
	if ttu := userset.GetTupleToUserset(); ttu != nil {
		target := ttu.GetComputedUserset().GetRelation()
		from := ttu.GetTupleset().GetRelation()
		if target != "" && from != "" {
			return target + " from " + from
		}
		if target != "" {
			return target
		}
	}
	return ""
}

// ValidateDuplicates performs comprehensive duplicate detection on a model.
func ValidateDuplicates(collector *ErrorCollector, model *openfgav1.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}
	typeTracker := NewDuplicateTypeTracker()
	for _, typeDef := range model.GetTypeDefinitions() {
		typeName := typeDef.GetType()
		if typeName == "" {
			continue
		}
		meta := &Meta{
			File:   typeDef.GetMetadata().GetSourceInfo().GetFile(),
			Module: typeDef.GetMetadata().GetModule(),
		}
		typeTracker.CheckAndAddType(typeName, collector, meta, lines)
		typeLineIndex := GetTypeLineNumber(typeName, lines, nil)
		if metaProto := typeDef.GetMetadata(); metaProto != nil {
			for relationName, relationMetadata := range metaProto.GetRelations() {
				CheckForDuplicateTypeNamesInRelation(collector, relationMetadata, relationName, typeName, meta, typeLineIndex, lines)
				CheckForDuplicatesInRelation(collector, typeDef, relationName, typeLineIndex, lines)
			}
		}
	}
}
