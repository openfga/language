package validation

import (
	"regexp"
	"strings"

	fgaSdk "github.com/openfga/go-sdk"
)

// Supported schema versions.
const (
	SchemaVersion11 = "1.1"
	SchemaVersion12 = "1.2"
)

// SupportedSchemaVersions contains all supported schema versions.
var SupportedSchemaVersions = map[string]bool{
	SchemaVersion11: true,
	SchemaVersion12: true,
}

// IsValidSchemaVersion checks if a schema version is supported.
func IsValidSchemaVersion(version string) bool {
	return SupportedSchemaVersions[version]
}

// GetSchemaLineNumber finds the line number where a schema version is defined
// This is equivalent to the getSchemaLineNumber function in JS
func GetSchemaLineNumber(schemaVersion string, lines []string) *int {
	if len(lines) == 0 {
		return nil
	}

	// Create regex pattern to match "schema X.X" with flexible whitespace
	// Equivalent to: line.trim().replace(/ {2,}/g, " ").match(`^schema ${schema}$`)
	pattern := `^\s*schema\s+` + regexp.QuoteMeta(schemaVersion) + `\s*$`
	regex := regexp.MustCompile(pattern)

	for i, line := range lines {
		// Normalize whitespace like the JS implementation
		normalizedLine := strings.TrimSpace(line)
		normalizedLine = regexp.MustCompile(`\s{2,}`).ReplaceAllString(normalizedLine, " ")

		if regex.MatchString(normalizedLine) {
			return &i
		}
	}

	return nil
}

// ValidateSchemaVersion validates the schema version of an authorization model
// This is equivalent to the schema validation logic in validateJSON
func ValidateSchemaVersion(collector *ErrorCollector, model *fgaSdk.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}

	schemaVersion := model.SchemaVersion

	// Check if schema version is missing
	if schemaVersion == "" {
		lineIndex := 0
		collector.RaiseSchemaVersionRequired("", &lineIndex)
		return
	}

	// Check if schema version is supported
	if !IsValidSchemaVersion(schemaVersion) {
		lineIndex := GetSchemaLineNumber(schemaVersion, lines)
		collector.RaiseInvalidSchemaVersion(schemaVersion, lineIndex)
		return
	}
}

// ValidateMultipleModulesInFile checks for multiple modules defined in single files
// This is equivalent to the multiple modules validation in validateJSON
func ValidateMultipleModulesInFile(collector *ErrorCollector, fileToModuleMap map[string]map[string]bool) {
	for file, moduleMap := range fileToModuleMap {
		if len(moduleMap) <= 1 {
			continue
		}

		// Convert module map to slice for error reporting
		modules := make([]string, 0, len(moduleMap))
		for module := range moduleMap {
			modules = append(modules, module)
		}

		collector.RaiseMultipleModulesInSingleFile(file, modules)
	}
}

// ValidateBasicModelStructure performs basic model structure validation
// This combines the fundamental validation checks before semantic analysis
func ValidateBasicModelStructure(collector *ErrorCollector, model *fgaSdk.AuthorizationModel,
	fileToModuleMap map[string]map[string]bool, lines []string) {
	// Schema version validation (required first)
	ValidateSchemaVersion(collector, model, lines)

	// Multiple modules in file validation
	ValidateMultipleModulesInFile(collector, fileToModuleMap)

	// If we have critical errors, don't proceed with further validation
	if collector.HasErrors() {
		return
	}
}
