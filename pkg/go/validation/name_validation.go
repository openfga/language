package validation

import (
	"fmt"
	"maps"
	"regexp"
	"slices"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// The anchored name rules, compiled once. Each rule string doubles as the
// clause quoted in the invalid-name message, so the values are corpus-pinned.
var (
	typeNameRule      = fmt.Sprintf("^%s$", RuleType)
	relationNameRule  = fmt.Sprintf("^%s$", RuleRelation)
	conditionNameRule = fmt.Sprintf("^%s$", RuleCondition)

	typeNameRegex      = regexp.MustCompile(typeNameRule)
	relationNameRegex  = regexp.MustCompile(relationNameRule)
	conditionNameRegex = regexp.MustCompile(conditionNameRule)
)

// reservedKeywords are the words a type or relation cannot be named.
var reservedKeywords = map[string]bool{
	"self": true,
	"this": true,
}

// validateTypeName checks one type name against the reserved keywords and the
// naming rule, returning the finding or nil. It knows nothing about the source
// text; the caller stamps position.
func validateTypeName(name string) *Finding {
	switch {
	case reservedKeywords[name]:
		return reservedTypeName(name)
	case !typeNameRegex.MatchString(name):
		return invalidTypeName(name)
	}

	return nil
}

// validateRelationName checks one relation name, returning the finding or nil.
func validateRelationName(name, typeName string) *Finding {
	switch {
	case reservedKeywords[name]:
		return reservedRelationName(name, typeName)
	case !relationNameRegex.MatchString(name):
		return invalidRelationName(name, typeName)
	}

	return nil
}

// validateConditionName checks one condition name, returning the finding or nil.
func validateConditionName(name string) *Finding {
	if !conditionNameRegex.MatchString(name) {
		return invalidConditionName(name)
	}

	return nil
}

// validateNames checks every type, relation, and condition name in the model
// against the reserved-keyword and naming-rule constraints. It mirrors the name
// validation performed in the JS reference implementation's populateRelations.
func validateNames(model *openfgav1.AuthorizationModel, src source) error {
	var fs []*Finding

	for _, typeDef := range model.GetTypeDefinitions() {
		typeName := typeDef.GetType()
		if typeName == "" {
			continue
		}

		file := typeDef.GetMetadata().GetSourceInfo().GetFile()
		module := typeDef.GetMetadata().GetModule()

		typeLine := src.typeLine(typeName)
		fs = append(fs, validateTypeName(typeName).at(src, typeLine).in(file, module))

		// Relations reach us in a proto map, which has no order, so they are
		// walked in name order here and in every other phase to report the same
		// model's findings in the same order from run to run.
		for _, relationName := range slices.Sorted(maps.Keys(typeDef.GetRelations())) {
			relationLine := src.relationLine(relationName, typeLine)
			fs = append(fs, validateRelationName(relationName, typeName).at(src, relationLine).in(file, module))
		}
	}

	conditions := model.GetConditions()
	for _, conditionName := range slices.Sorted(maps.Keys(conditions)) {
		condition := conditions[conditionName]
		file := condition.GetMetadata().GetSourceInfo().GetFile()
		module := condition.GetMetadata().GetModule()

		fs = append(fs, validateConditionName(conditionName).at(src, src.conditionLine(conditionName)).in(file, module))
	}

	return joinFindings(fs...)
}
