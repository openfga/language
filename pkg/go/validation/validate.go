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

// ValidateDSLWithGraph runs the same validations as ValidateDSL but resolves
// entry-point reachability from the weighted authorization-model graph rather
// than by walking the rewrite tree. Where the graph builds, it reports the same
// relation-no-entry-point findings as ValidateDSL: same message, code and
// position.
//
// A model the graph builder refuses is reported by the first of these that
// applies: the findings the rewrite-tree walk reports for it; cyclic-relation
// findings naming the relations of a cycle the resolver cannot work through;
// or a single positionless graph-model-unbuildable finding. The last two codes
// are Go-only; pkg/js and pkg/java have no graph path.
func ValidateDSLWithGraph(model *openfgav1.AuthorizationModel, dsl string) error {
	return validateWith(model, newSource(dsl), validateWithGraph)
}

// ValidateJSONWithGraph is ValidateDSLWithGraph for a model that has no DSL
// source. Findings carry no Line or Column; messages and metadata match what
// ValidateDSLWithGraph reports for the same model. The name matches
// ValidateJSON.
func ValidateJSONWithGraph(model *openfgav1.AuthorizationModel) error {
	return validateWith(model, source{}, validateWithGraph)
}

// entryPointPhase answers which relations can never be satisfied. The default
// validateEntryPoints walks the rewrite tree; validateWithGraph reads the same
// question off the weighted graph. The two share every other phase, the phase
// order and the cascade gate through validateWith.
type entryPointPhase func(idx *index, src source) error

// validate runs every validation phase over model, resolving entry points by
// walking the rewrite tree.
func validate(model *openfgav1.AuthorizationModel, src source) error {
	return validateWith(model, src, validateEntryPoints)
}

// validateWith runs the validation phases over model and joins their findings.
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
	}

	add(validateMultiFile(model))
	add(validateConditions(model, src))

	return errors.Join(errs...)
}
