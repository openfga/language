package validation

import (
	"errors"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// ValidateDSL validates model, which was parsed from dsl, and uses dsl only to
// find each finding's line and column.
//
// It returns nil for a valid model. Otherwise the error joins every finding in
// the order raised, which ExtractAllAs recovers:
//
//	for _, finding := range validation.ExtractAllAs[*validation.Finding](err) {
//	    ...
//	}
func ValidateDSL(model *openfgav1.AuthorizationModel, dsl string) error {
	return validate(model, newSource(dsl))
}

// ValidateJSON validates a model that has no DSL source, such as one received
// as JSON. Its findings have no Line or Column but otherwise match what
// ValidateDSL reports for the same model. The name matches pkg/js's
// validateJSON and pkg/java's validateJson.
func ValidateJSON(model *openfgav1.AuthorizationModel) error {
	return validate(model, source{})
}

// validate runs the validation phases over model and joins their findings.
//
// Schema, name and relation-reference validation always run. Duplicate
// detection runs only if they found nothing, and entry-point detection only if
// nothing has been found at all: a model with bad references or duplicates
// would otherwise produce a cascade of derived entry-point errors for the same
// root cause. Multi-file and condition checks always run.
//
// The gating in pkg/js differs, so for some models it reports fewer findings
// than this does. A schema or name finding skips every later check but the
// multi-file one, and an undefined condition, which pkg/js reports with the
// relation references, skips duplicate and entry-point detection. A skipped
// check also leaves the modules of the relations or conditions it would have
// visited out of the multi-file check. Findings are also not always in
// pkg/js's order, which this package's shared-case test does not check.
func validate(model *openfgav1.AuthorizationModel, src source) error {
	if model == nil {
		return nil
	}

	idx := newIndex(model)

	var errs []error
	add := func(err error) {
		if err != nil {
			errs = append(errs, err)
		}
	}

	add(validateSchemaVersion(model, src))
	add(validateNames(model, src))
	add(validateRelationReferences(idx, src))

	if len(errs) == 0 {
		add(validateDuplicates(model, src))
	}

	if len(errs) == 0 {
		add(validateEntryPoints(idx, src))
	}

	add(validateMultiFile(model))
	add(validateConditions(model, src))

	return errors.Join(errs...)
}
