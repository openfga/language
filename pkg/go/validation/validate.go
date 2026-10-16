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

// ValidateDSLWithGraph runs the same validations as ValidateDSL but resolves
// entry-point reachability from the weighted authorization-model graph rather
// than by walking the rewrite tree. Everywhere the graph builds, it reports the
// same relation-no-entry-point findings as ValidateDSL: same message, code and
// position. A model the graph builder refuses yields a single positionless
// finding coded GraphModelUnbuildable, which stands in for the per-relation
// findings the tree walk would report and has no pkg/js or pkg/java counterpart.
//
// ValidateDSL remains the default; this path is opt-in until the graph is the
// source of truth for reachability.
func ValidateDSLWithGraph(model *openfgav1.AuthorizationModel, dsl string) error {
	return validateWith(model, newSource(dsl), validateWithGraph)
}

// ValidateJSONWithGraph is ValidateDSLWithGraph for a model that reached the
// caller as JSON, so with no DSL source text behind it. Findings carry a nil
// Line and Column; the messages and metadata match what ValidateDSLWithGraph
// reports for the same model. The name matches ValidateJSON.
func ValidateJSONWithGraph(model *openfgav1.AuthorizationModel) error {
	return validateWith(model, source{}, validateWithGraph)
}

// validate runs every validation phase over model, resolving entry points by
// walking the rewrite tree.
func validate(model *openfgav1.AuthorizationModel, src source) error {
	return validateWith(model, src, validateEntryPoints)
}

// entryPointPhase reports relations that can never be satisfied. The default
// validateEntryPoints walks the rewrite tree; validateWithGraph reads the same
// question off the weighted graph. Passing it to validateWith as a parameter
// lets the two share every other phase, the phase order and the cascade gate,
// and differ only in how they answer that one question.
type entryPointPhase func(idx *index, src source) error

// validateWith runs the validation phases in the reference implementation's
// order, resolving entry points through the given phase.
//
// Schema, name and reference validation always run. The later structural
// phases are gated on nothing having been found yet: a model with bad
// references or duplicates would otherwise produce a cascade of derived
// entry-point and operation errors for the same root cause. This mirrors the
// reference's modelValidation, which skips the later passes once any error has
// been recorded. Multi-file and condition checks are independent of the cascade
// and always run, matching the reference's handling of conditions.
func validateWith(model *openfgav1.AuthorizationModel, src source, entryPoints entryPointPhase) error {
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
		add(entryPoints(idx, src))
		add(validateTupleToUsersets(idx, src))
		add(validateComplexOperations(idx, src))
		add(validateWildcards(idx, src))
	}

	add(validateMultiFile(model))
	add(validateConditions(model, src))

	return errors.Join(errs...)
}
