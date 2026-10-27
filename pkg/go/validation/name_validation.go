package validation

import (
	"errors"
	"fmt"
	"maps"
	"regexp"
	"slices"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// The anchored name rules, compiled once. Each rule string doubles as the
// clause quoted in the invalid-name message, so the shared test cases fix the values.
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
// naming rule, returning the finding as an error or nil. It knows nothing about
// the source text; the caller recovers the finding to stamp position.
func validateTypeName(name string) error {
	switch {
	case reservedKeywords[name]:
		return reservedTypeName(name)
	case !typeNameRegex.MatchString(name):
		return invalidTypeName(name)
	}

	return nil
}

// validateRelationName checks one relation name, returning the finding as an
// error or nil.
func validateRelationName(name, typeName string) error {
	switch {
	case reservedKeywords[name]:
		return reservedRelationName(name, typeName)
	case !relationNameRegex.MatchString(name):
		return invalidRelationName(name, typeName)
	}

	return nil
}

// validateConditionName checks one condition name, returning the finding as an
// error or nil. Go-only: pkg/js and pkg/java accept any condition name the
// grammar parses.
func validateConditionName(name string) error {
	if !conditionNameRegex.MatchString(name) {
		return invalidConditionName(name)
	}

	return nil
}

// validateNames checks every type, relation and condition name in the model
// against the reserved keywords and the naming rules. Neither pkg/js nor
// pkg/java checks condition names.
func validateNames(model *openfgav1.AuthorizationModel, src source) error {
	var fs []*Finding

	for _, typeDef := range model.GetTypeDefinitions() {
		typeName := typeDef.GetType()
		if typeName == "" {
			continue
		}

		file, module := typeMeta(typeDef)

		typeLine := src.typeLine(typeName)

		var finding *Finding
		if errors.As(validateTypeName(typeName), &finding) {
			finding.at(src, typeLine)
			finding.File, finding.Metadata.Module = file, module
			fs = append(fs, finding)
		}

		// Relations reach us in a proto map, which has no order, so they are
		// walked in name order here and in every other phase to report the same
		// model's findings in the same order from run to run.
		for _, relationName := range slices.Sorted(maps.Keys(typeDef.GetRelations())) {
			relationLine := src.relationLine(relationName, typeLine)
			var finding *Finding
			if errors.As(validateRelationName(relationName, typeName), &finding) {
				finding.at(src, relationLine)
				finding.File, finding.Metadata.Module = file, module
				fs = append(fs, finding)
			}
		}
	}

	conditions := model.GetConditions()
	for _, conditionName := range slices.Sorted(maps.Keys(conditions)) {
		condition := conditions[conditionName]
		file := condition.GetMetadata().GetSourceInfo().GetFile()
		module := condition.GetMetadata().GetModule()

		var finding *Finding
		if errors.As(validateConditionName(conditionName), &finding) {
			finding.at(src, src.conditionLine(conditionName))
			finding.File, finding.Metadata.Module = file, module
			fs = append(fs, finding)
		}
	}

	return joinFindings(fs...)
}
