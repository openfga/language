package validation

import (
	"regexp"
	"strings"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

const (
	SchemaVersion11 = "1.1"
	SchemaVersion12 = "1.2"
)

var SupportedSchemaVersions = map[string]bool{
	SchemaVersion11: true,
	SchemaVersion12: true,
}

// multiSpaceRegex collapses runs of whitespace when normalizing a DSL line for
// schema-version matching. Hoisted so it is compiled once, not per line.
var multiSpaceRegex = regexp.MustCompile(`\s{2,}`)

func IsValidSchemaVersion(version string) bool {
	return SupportedSchemaVersions[version]
}

func GetSchemaLineNumber(schemaVersion string, lines []string) *int {
	if len(lines) == 0 {
		return nil
	}
	// A trailing comment may follow the version, as in `schema 1.1 # note`. The `#`
	// has to be preceded by whitespace, so one written against the version is part of
	// the version and does not match here. This mirrors the reference's
	// getSchemaLineNumber; without it a commented schema line resolves to no
	// position and the finding reaches the caller with no line or column.
	pattern := `^\s*schema\s+` + regexp.QuoteMeta(schemaVersion) + `(\s+#.*)?\s*$`
	regex := regexp.MustCompile(pattern)
	for i, line := range lines {
		normalizedLine := strings.TrimSpace(line)
		normalizedLine = multiSpaceRegex.ReplaceAllString(normalizedLine, " ")
		if regex.MatchString(normalizedLine) {
			return &i
		}
	}
	return nil
}

// ValidateSchemaVersion validates the schema version of an authorization model.
func ValidateSchemaVersion(errs *ValidationErrors, model *openfgav1.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}
	schemaVersion := model.GetSchemaVersion()
	if schemaVersion == "" {
		lineIndex := 0
		errs.Add(newSchemaVersionRequiredError(lines, &lineIndex))
		return
	}
	switch schemaVersion {
	case SchemaVersion11, SchemaVersion12:
		// Supported — nothing to report.
	case "1.0":
		// Recognized but retired.
		errs.Add(newSchemaVersionUnsupportedError(lines, schemaVersion, GetSchemaLineNumber(schemaVersion, lines)))
	default:
		// Never a valid schema version.
		errs.Add(newInvalidSchemaVersionError(lines, schemaVersion, GetSchemaLineNumber(schemaVersion, lines)))
	}
}

// ValidateMultipleModulesInFile reports every file that declares more than one
// module.
//
// It reports the files, and each file's modules, in the order they were collected
// from the model, which is the order the reference reports them in and the order the
// shared corpus expects.
func ValidateMultipleModulesInFile(errs *ValidationErrors, files []FileInfo) {
	for _, file := range files {
		if len(file.Modules) <= 1 {
			continue
		}

		errs.Add(newMultipleModulesInSingleFileError(file.Path, file.Modules))
	}
}

// ValidateBasicModelStructure performs basic model structure validation.
func ValidateBasicModelStructure(errs *ValidationErrors, model *openfgav1.AuthorizationModel,
	files []FileInfo, lines []string) {
	ValidateSchemaVersion(errs, model, lines)
	ValidateMultipleModulesInFile(errs, files)
}
