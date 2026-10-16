package validation

import (
	"regexp"
	"strings"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"

	"github.com/openfga/language/pkg/go/utils"
)

const (
	SchemaVersion11 = "1.1"
	SchemaVersion12 = "1.2"
)

var SupportedSchemaVersions = map[string]bool{
	SchemaVersion11: true,
	SchemaVersion12: true,
}

func IsValidSchemaVersion(version string) bool {
	return SupportedSchemaVersions[version]
}

func GetSchemaLineNumber(schemaVersion string, lines []string) *int {
	if len(lines) == 0 {
		return nil
	}
	// A trailing comment may follow the version (`schema 1.5 # note`), matching the
	// reference's `(\s+#.*)?`.
	pattern := `^\s*schema\s+` + regexp.QuoteMeta(schemaVersion) + `(\s+#.*)?$`
	regex := regexp.MustCompile(pattern)
	for i, line := range lines {
		// Fold inline whitespace with the helper the other line-number lookups
		// use, so `schema\t1.5` resolves like `schema 1.5`.
		normalizedLine := utils.NormalizeWhitespace(strings.TrimSpace(line))
		if regex.MatchString(normalizedLine) {
			return &i
		}
	}
	return nil
}

// ValidateSchemaVersion validates the schema version of an authorization model.
func ValidateSchemaVersion(collector *ErrorCollector, model *openfgav1.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}
	schemaVersion := model.GetSchemaVersion()
	if schemaVersion == "" {
		lineIndex := 0
		collector.RaiseSchemaVersionRequired("", &lineIndex)
		return
	}
	switch schemaVersion {
	case SchemaVersion11, SchemaVersion12:
		// Supported — nothing to report.
	case "1.0":
		// Recognized but retired.
		collector.RaiseSchemaVersionUnsupported(schemaVersion, GetSchemaLineNumber(schemaVersion, lines))
	default:
		// Never a valid schema version.
		collector.RaiseInvalidSchemaVersion(schemaVersion, GetSchemaLineNumber(schemaVersion, lines))
	}
}

// ValidateMultipleModulesInFile checks for multiple modules defined in single files.
func ValidateMultipleModulesInFile(collector *ErrorCollector, fileToModuleMap map[string]map[string]bool) {
	for file, moduleMap := range fileToModuleMap {
		if len(moduleMap) <= 1 {
			continue
		}
		modules := make([]string, 0, len(moduleMap))
		for module := range moduleMap {
			modules = append(modules, module)
		}
		collector.RaiseMultipleModulesInSingleFile(file, modules)
	}
}

// ValidateBasicModelStructure performs basic model structure validation.
func ValidateBasicModelStructure(collector *ErrorCollector, model *openfgav1.AuthorizationModel,
	fileToModuleMap map[string]map[string]bool, lines []string) {
	ValidateSchemaVersion(collector, model, lines)
	ValidateMultipleModulesInFile(collector, fileToModuleMap)
}
