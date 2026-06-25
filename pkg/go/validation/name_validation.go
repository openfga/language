package validation

import (
	"fmt"
	"regexp"
	"strings"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// ValidationRegexRules contains the regex rules for validation
// These match the Rules from the JS implementation
var ValidationRegexRules = struct {
	Type      string
	Relation  string
	Condition string
	ID        string
	Object    string
}{
	Type:      "[^:#@\\*\\s]{1,254}",
	Relation:  "[^:#@\\*\\s]{1,50}",
	Condition: "[^\\*\\s]{1,50}",
	ID:        "[^#:\\s*][a-zA-Z0-9_|*@.+]*",
	Object:    "[^\\s]{2,256}",
}

// The anchored type, relation, and condition name rules are fixed, so compile
// them once. compiledNameRules caches them by their anchored rule string, which
// is also the clause reported in the error, so validateFieldValue can look up
// the compiled pattern without recompiling on every name.
var (
	typeNameRule      = fmt.Sprintf("^%s$", ValidationRegexRules.Type)
	relationNameRule  = fmt.Sprintf("^%s$", ValidationRegexRules.Relation)
	conditionNameRule = fmt.Sprintf("^%s$", ValidationRegexRules.Condition)

	compiledNameRules = map[string]*regexp.Regexp{
		typeNameRule:      regexp.MustCompile(typeNameRule),
		relationNameRule:  regexp.MustCompile(relationNameRule),
		conditionNameRule: regexp.MustCompile(conditionNameRule),
	}
)

// ValidateTypeName validates a type name with both regex and reserved keyword checking
// This enhances the basic regex validation with semantic checks
func ValidateTypeName(typeName string, collector *ErrorCollector, lineIndex *int, meta *Meta) bool {
	// First check if it's a reserved keyword
	if IsReservedTypeName(typeName) {
		collector.RaiseReservedTypeName(typeName, lineIndex, meta)
		return false
	}

	// Then check regex pattern. The clause passed to the error is the full
	// anchored rule, matching the reference implementation's reported rule.
	if !validateFieldValue(typeNameRule, typeName) {
		collector.RaiseInvalidName(typeName, typeNameRule, nil, lineIndex, meta)
		return false
	}

	return true
}

// ValidateRelationName validates a relation name with both regex and reserved keyword checking
// This enhances the basic regex validation with semantic checks
func ValidateRelationName(relationName, typeName string, collector *ErrorCollector, lineIndex *int, meta *Meta) bool {
	// First check if it's a reserved keyword
	if IsReservedRelationName(relationName) {
		collector.RaiseReservedRelationName(relationName, lineIndex, meta)
		return false
	}

	// Then check regex pattern. The clause passed to the error is the full
	// anchored rule, matching the reference implementation's reported rule.
	if !validateFieldValue(relationNameRule, relationName) {
		collector.RaiseInvalidName(relationName, relationNameRule, &typeName, lineIndex, meta)
		return false
	}

	return true
}

// ValidateConditionName validates a condition name with regex pattern.
func ValidateConditionName(conditionName string, collector *ErrorCollector, lineIndex *int, meta *Meta) bool {
	if !validateFieldValue(conditionNameRule, conditionName) {
		collector.RaiseInvalidName(conditionName, conditionNameRule, nil, lineIndex, meta)
		return false
	}

	return true
}

// validateFieldValue validates a field against a regex rule. Fixed rules are
// served from the precompiled cache; any other rule is compiled on demand.
func validateFieldValue(rule, value string) bool {
	if regex, ok := compiledNameRules[rule]; ok {
		return regex.MatchString(value)
	}
	regex, err := regexp.Compile(rule)
	if err != nil {
		return false
	}
	return regex.MatchString(value)
}

// GetTypeLineNumber finds the line number where a type is defined
// This is equivalent to the getTypeLineNumber function in JS
func GetTypeLineNumber(typeName string, lines []string, skipIndex *int) *int {
	if len(lines) == 0 {
		return nil
	}

	for i, line := range lines {
		// Skip the specified index if provided
		if skipIndex != nil && i == *skipIndex {
			continue
		}

		// Look for "type typeName" pattern
		trimmedLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimmedLine, "type ") {
			parts := strings.Fields(trimmedLine)
			if len(parts) >= 2 && parts[1] == typeName {
				return &i
			}
		}
	}

	return nil
}

// GetRelationLineNumber finds the line number where a relation is defined.
// skipIndex, when provided, is the index to begin searching from (inclusive) —
// matching the reference implementation's getRelationLineNumber, which slices
// the lines from skipIndex onward. This lets callers anchor the search to a
// specific type block so the correct occurrence is found when several types
// declare a relation of the same name.
func GetRelationLineNumber(relationName string, lines []string, skipIndex *int) *int {
	if len(lines) == 0 {
		return nil
	}

	start := 0
	if skipIndex != nil && *skipIndex > 0 {
		start = *skipIndex
	}

	for i := start; i < len(lines); i++ {
		// Look for "define relationName:" pattern
		trimmedLine := strings.TrimSpace(lines[i])
		if strings.HasPrefix(trimmedLine, "define ") {
			// Extract relation name from "define relationName:"
			definePart := strings.TrimPrefix(trimmedLine, "define ")
			colonIndex := strings.Index(definePart, ":")
			if colonIndex > 0 {
				relationInLine := strings.TrimSpace(definePart[:colonIndex])
				if relationInLine == relationName {
					return &i
				}
			}
		}
	}

	return nil
}

// GetConditionLineNumber finds the line number where a condition is defined
// This is equivalent to the geConditionLineNumber function in JS
func GetConditionLineNumber(conditionName string, lines []string, skipIndex *int) *int {
	if len(lines) == 0 {
		return nil
	}

	start := 0
	if skipIndex != nil && *skipIndex > 0 {
		start = *skipIndex
	}

	for i := start; i < len(lines); i++ {
		// Match the condition declaration itself, mirroring the reference's
		// `condition <name>` prefix check, so we don't match an unrelated line
		// that merely contains the condition name as a substring.
		trimmedLine := strings.TrimSpace(lines[i])
		if strings.HasPrefix(trimmedLine, "condition "+conditionName) {
			return &i
		}
	}

	return nil
}

// ValidateNameRules validates naming rules for types and relations in a model
// This is equivalent to the populateRelations function's naming validation in JS
func ValidateNameRules(collector *ErrorCollector, typeName string, relationNames []string,
	typeLineIndex *int, meta *Meta, lines []string) {
	// Validate type name
	ValidateTypeName(typeName, collector, typeLineIndex, meta)

	// Validate relation names
	for _, relationName := range relationNames {
		relationLineIndex := GetRelationLineNumber(relationName, lines, nil)
		ValidateRelationName(relationName, typeName, collector, relationLineIndex, meta)
	}
}

// ValidateNames checks every type, relation, and condition name in the model
// against the reserved-keyword and naming-rule constraints. It mirrors the name
// validation performed in the JS reference implementation's populateRelations.
func ValidateNames(collector *ErrorCollector, model *openfgav1.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}

	for _, typeDef := range model.GetTypeDefinitions() {
		typeName := typeDef.GetType()
		if typeName == "" {
			continue
		}
		meta := &Meta{
			File:   typeDef.GetMetadata().GetSourceInfo().GetFile(),
			Module: typeDef.GetMetadata().GetModule(),
		}

		typeLineIndex := GetTypeLineNumber(typeName, lines, nil)
		ValidateTypeName(typeName, collector, typeLineIndex, meta)

		for relationName := range typeDef.GetRelations() {
			relationLineIndex := GetRelationLineNumber(relationName, lines, typeLineIndex)
			ValidateRelationName(relationName, typeName, collector, relationLineIndex, meta)
		}
	}

	for conditionName, condition := range model.GetConditions() {
		conditionLineIndex := GetConditionLineNumber(conditionName, lines, nil)
		meta := &Meta{
			File:   condition.GetMetadata().GetSourceInfo().GetFile(),
			Module: condition.GetMetadata().GetModule(),
		}
		ValidateConditionName(conditionName, collector, conditionLineIndex, meta)
	}
}
