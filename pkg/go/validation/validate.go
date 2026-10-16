package validation

import (
	"errors"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// ValidateDSL runs every validation over model, using dsl — the source text the
// model was parsed from — to resolve each finding's position. The model is the
// already-parsed proto; nothing here parses.
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

// ValidateJSON runs every validation over a model that reached the caller as
// JSON, so with no DSL source text behind it. Findings carry a nil Line and
// Column; the messages and metadata are what ValidateDSL reports for the same
// model. The name matches pkg/js's validateJSON and pkg/java's validateJson.
func ValidateJSON(model *openfgav1.AuthorizationModel) error {
	return validate(model, source{})
}

// validate runs the validation phases in the reference implementation's order.
//
// Schema, name and reference validation always run. The later structural
// phases are gated on nothing having been found yet: a model with bad
// references or duplicates would otherwise produce a cascade of derived
// entry-point and operation errors for the same root cause. This mirrors the
// reference's modelValidation, which skips the later passes once any error has
// been recorded. Multi-file and condition checks are independent of the cascade
// and always run, matching the reference's handling of conditions.
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
		add(validateTupleToUsersets(idx, src))
		add(validateComplexOperations(idx, src))
		add(validateWildcards(idx, src))
	}

	add(validateMultiFile(model))
	add(validateConditions(model, src))

	return errors.Join(errs...)
}
