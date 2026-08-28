package validation

import (
	"maps"
	"slices"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// ComplexOperationValidator handles validation of complex userset operations.
type ComplexOperationValidator struct {
	model     *openfgav1.AuthorizationModel
	validator *SemanticValidator
}

func NewComplexOperationValidator(model *openfgav1.AuthorizationModel) *ComplexOperationValidator {
	return newComplexOperationValidator(NewSemanticValidator(model))
}

func newComplexOperationValidator(validator *SemanticValidator) *ComplexOperationValidator {
	return &ComplexOperationValidator{
		model:     validator.model,
		validator: validator,
	}
}

// ValidateComplexOperations validates all complex operations in the model.
func ValidateComplexOperations(errs *ValidationErrors, model *openfgav1.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}
	validateComplexOperations(errs, NewSemanticValidator(model), lines)
}

func validateComplexOperations(errs *ValidationErrors, validator *SemanticValidator, lines []string) {
	model := validator.model
	if model == nil {
		return
	}
	opValidator := newComplexOperationValidator(validator)
	for _, typeDef := range model.GetTypeDefinitions() {
		relations := typeDef.GetRelations()
		for _, relationName := range slices.Sorted(maps.Keys(relations)) {
			opValidator.validateUsersetOperations(errs, typeDef.GetType(), relationName,
				relations[relationName], lines)
		}
	}
}

func (cov *ComplexOperationValidator) validateUsersetOperations(errs *ValidationErrors, typeName, relationName string, userset *openfgav1.Userset, lines []string) {
	cov.validateUsersetOperationsWithVisited(errs, typeName, relationName, userset, lines, make(map[string]bool))
}

func (cov *ComplexOperationValidator) validateUsersetOperationsWithVisited(errs *ValidationErrors, typeName, relationName string, userset *openfgav1.Userset, lines []string, visited map[string]bool) {
	if userset == nil {
		return
	}
	if union := userset.GetUnion(); union != nil {
		cov.validateUnionOperationWithVisited(errs, typeName, relationName, union, lines, visited)
	}
	if intersection := userset.GetIntersection(); intersection != nil {
		cov.validateIntersectionOperationWithVisited(errs, typeName, relationName, intersection, lines, visited)
	}
	if diff := userset.GetDifference(); diff != nil {
		cov.validateDifferenceOperationWithVisited(errs, typeName, relationName, diff, lines, visited)
	}
	cov.validateNestedOperationsWithVisited(errs, typeName, userset, lines, visited)
}

func (cov *ComplexOperationValidator) validateUnionOperationWithVisited(errs *ValidationErrors, typeName, relationName string, union *openfgav1.Usersets, lines []string, visited map[string]bool) {
	if union == nil || len(union.GetChild()) == 0 {
		return
	}
	cov.checkRedundantUnionMembers(errs, typeName, relationName, union, lines)
	for _, child := range union.GetChild() {
		cov.validateUsersetOperationsWithVisited(errs, typeName, relationName, child, lines, visited)
	}
	cov.validateUnionSemantics(errs, typeName, relationName, union, lines)
}

func (cov *ComplexOperationValidator) validateIntersectionOperationWithVisited(errs *ValidationErrors, typeName, relationName string, intersection *openfgav1.Usersets, lines []string, visited map[string]bool) {
	if intersection == nil || len(intersection.GetChild()) == 0 {
		return
	}
	cov.checkImpossibleIntersections(errs, typeName, relationName, intersection, lines)
	for _, child := range intersection.GetChild() {
		cov.validateUsersetOperationsWithVisited(errs, typeName, relationName, child, lines, visited)
	}
	cov.validateIntersectionSemantics(errs, typeName, relationName, intersection, lines)
}

func (cov *ComplexOperationValidator) validateDifferenceOperationWithVisited(errs *ValidationErrors, typeName, relationName string, difference *openfgav1.Difference, lines []string, visited map[string]bool) {
	if difference == nil {
		return
	}
	cov.validateUsersetOperationsWithVisited(errs, typeName, relationName, difference.GetBase(), lines, visited)
	cov.validateUsersetOperationsWithVisited(errs, typeName, relationName, difference.GetSubtract(), lines, visited)
	cov.validateDifferenceSemantics(errs, typeName, relationName, difference, lines)
}

func (cov *ComplexOperationValidator) checkRedundantUnionMembers(errs *ValidationErrors, typeName, relationName string, union *openfgav1.Usersets, lines []string) {
	seenOperations := make(map[string]bool)
	for _, child := range union.GetChild() {
		operationKey := cov.getUsersetOperationKey(child)
		if operationKey == "" {
			continue
		}
		if seenOperations[operationKey] {
			lineIndex := GetRelationLineNumber(relationName, lines, nil)
			meta := cov.getTypeMeta(typeName)
			errs.Add(newRedundantUnionMemberError(lines, operationKey, relationName, typeName, meta, lineIndex))
		}
		seenOperations[operationKey] = true
	}
}

func (cov *ComplexOperationValidator) checkImpossibleIntersections(errs *ValidationErrors, typeName, relationName string, intersection *openfgav1.Usersets, lines []string) {
	typeRestrictions := make([]string, 0)
	for _, child := range intersection.GetChild() {
		if child.GetThis() != nil {
			typeRestrictions = append(typeRestrictions, "this")
		}
	}
	if len(typeRestrictions) <= 1 {
		return
	}
	uniqueTypes := make(map[string]bool)
	for _, t := range typeRestrictions {
		uniqueTypes[t] = true
	}
	if len(uniqueTypes) > 1 {
		lineIndex := GetRelationLineNumber(relationName, lines, nil)
		meta := cov.getTypeMeta(typeName)
		errs.Add(newImpossibleIntersectionError(lines, relationName, typeName, typeRestrictions, meta, lineIndex))
	}
}

func (cov *ComplexOperationValidator) validateUnionSemantics(errs *ValidationErrors, typeName, relationName string, union *openfgav1.Usersets, lines []string) {
	cov.checkSubsumingUnionMembers(errs, typeName, relationName, union, lines)
}

func (cov *ComplexOperationValidator) validateIntersectionSemantics(errs *ValidationErrors, typeName, relationName string, intersection *openfgav1.Usersets, lines []string) {
	cov.checkRedundantIntersectionMembers(errs, typeName, relationName, intersection, lines)
}

func (cov *ComplexOperationValidator) validateDifferenceSemantics(errs *ValidationErrors, typeName, relationName string, difference *openfgav1.Difference, lines []string) {
	baseKey := cov.getUsersetOperationKey(difference.GetBase())
	subtractKey := cov.getUsersetOperationKey(difference.GetSubtract())
	if baseKey != "" && baseKey == subtractKey {
		lineIndex := GetRelationLineNumber(relationName, lines, nil)
		meta := cov.getTypeMeta(typeName)
		errs.Add(newEmptyDifferenceError(lines, relationName, typeName, baseKey, meta, lineIndex))
	}
}

func (cov *ComplexOperationValidator) validateNestedOperationsWithVisited(errs *ValidationErrors, typeName string, userset *openfgav1.Userset, lines []string, visited map[string]bool) {
	if ttu := userset.GetTupleToUserset(); ttu != nil {
		if targetRelation := ttu.GetComputedUserset().GetRelation(); targetRelation != "" {
			key := typeName + "#" + targetRelation
			if visited[key] {
				return
			}
			visited[key] = true
			if targetUserset := cov.validator.GetRelationUserset(typeName, targetRelation); targetUserset != nil {
				cov.validateUsersetOperationsWithVisited(errs, typeName, targetRelation, targetUserset, lines, visited)
			}
		}
	}
}

func (cov *ComplexOperationValidator) getUsersetOperationKey(userset *openfgav1.Userset) string {
	if userset == nil {
		return ""
	}
	if userset.GetThis() != nil {
		return "this"
	}
	if cu := userset.GetComputedUserset(); cu != nil {
		if rel := cu.GetRelation(); rel != "" {
			return "computed:" + rel
		}
	}
	if ttu := userset.GetTupleToUserset(); ttu != nil {
		tuplesetRel := ttu.GetTupleset().GetRelation()
		computedRel := ttu.GetComputedUserset().GetRelation()
		return "ttu:" + tuplesetRel + ":" + computedRel
	}
	return ""
}

func (cov *ComplexOperationValidator) getTypeMeta(typeName string) *Meta {
	if typeDef := cov.validator.GetTypeDefinition(typeName); typeDef != nil {
		return &Meta{
			File:   typeDef.GetMetadata().GetSourceInfo().GetFile(),
			Module: typeDef.GetMetadata().GetModule(),
		}
	}
	return &Meta{}
}

func (cov *ComplexOperationValidator) checkSubsumingUnionMembers(_ *ValidationErrors, _, _ string, _ *openfgav1.Usersets, _ []string) {
	// Would check for cases like [user:*, user] where the wildcard subsumes the
	// specific relation. This requires detailed analysis of type restrictions.
}

func (cov *ComplexOperationValidator) checkRedundantIntersectionMembers(_ *ValidationErrors, _, _ string, _ *openfgav1.Usersets, _ []string) {
	// Would check for intersection members that don't restrict the result, e.g.
	// intersecting with `this`, which adds no restriction.
}
