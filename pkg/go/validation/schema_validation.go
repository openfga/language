package validation

import (
	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// validateSchemaVersion reports a missing, retired, or never-valid schema
// version. A model with no version at all is reported at line zero, matching
// the reference.
func validateSchemaVersion(model *openfgav1.AuthorizationModel, src source) Findings {
	version := model.GetSchemaVersion()

	switch version {
	case "":
		return Findings{schemaVersionRequired().at(src, 0)}
	case "1.1", "1.2":
		return nil
	case "1.0":
		// Recognized but retired.
		return Findings{schemaVersionUnsupported(version).at(src, src.schemaLine(version))}
	default:
		// Never a valid schema version.
		return Findings{invalidSchemaVersion(version).at(src, src.schemaLine(version))}
	}
}
