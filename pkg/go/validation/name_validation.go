package validation

import (
	"fmt"
	"regexp"
	"strings"
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

// ValidateTypeName validates a type name with both regex and reserved keyword checking
// This enhances the basic regex validation with semantic checks
func ValidateTypeName(typeName string, collector *ErrorCollector, lineIndex *int, meta *Meta) bool {
	// First check if it's a reserved keyword
	if IsReservedTypeName(typeName) {
		collector.RaiseReservedTypeName(typeName, lineIndex, meta)
		return false
	}

	// Then check regex pattern
	if !validateFieldValue(fmt.Sprintf("^%s$", ValidationRegexRules.Type), typeName) {
		collector.RaiseInvalidName(typeName, ValidationRegexRules.Type, nil, lineIndex, meta)
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

	// Then check regex pattern
	if !validateFieldValue(fmt.Sprintf("^%s$", ValidationRegexRules.Relation), relationName) {
		collector.RaiseInvalidName(relationName, ValidationRegexRules.Relation, &typeName, lineIndex, meta)
		return false
	}

	return true
}

// ValidateConditionName validates a condition name with regex pattern.
func ValidateConditionName(conditionName string, collector *ErrorCollector, lineIndex *int, meta *Meta) bool {
	if !validateFieldValue(fmt.Sprintf("^%s$", ValidationRegexRules.Condition), conditionName) {
		collector.RaiseInvalidName(conditionName, ValidationRegexRules.Condition, nil, lineIndex, meta)
		return false
	}

	return true
}

// validateFieldValue validates a field against a regex pattern
// This is equivalent to the validateFieldValue function in the JS implementation
func validateFieldValue(rule, value string) bool {
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

// GetRelationLineNumber finds the line number where a relation is defined
// This is equivalent to the getRelationLineNumber function in JS
func GetRelationLineNumber(relationName string, lines []string, skipIndex *int) *int {
	if len(lines) == 0 {
		return nil
	}

	for i, line := range lines {
		// Skip the specified index if provided
		if skipIndex != nil && i == *skipIndex {
			continue
		}

		// Look for "define relationName:" pattern
		trimmedLine := strings.TrimSpace(line)
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

	for i, line := range lines {
		// Skip the specified index if provided
		if skipIndex != nil && i == *skipIndex {
			continue
		}

		// Look for condition definitions in various formats
		trimmedLine := strings.TrimSpace(line)

		// Check for condition name in the line
		if strings.Contains(trimmedLine, conditionName) {
			// Additional validation to ensure it's actually a condition definition
			// This could be enhanced based on the specific DSL format
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
