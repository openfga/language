package validation

import (
	"errors"
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	fgaerrors "github.com/openfga/language/pkg/go/errors"
	"github.com/openfga/language/pkg/go/transformer"
)

// modelFromDSL parses a DSL string, failing the test if it does not parse: these
// tests are about semantic validation, so a syntax error in a fixture is a bug in
// the test rather than a result.
func modelFromDSL(t *testing.T, dsl string) *openfgav1.AuthorizationModel {
	t.Helper()

	model, err := transformer.TransformDSLToProto(dsl)
	require.NoError(t, err, "test DSL must parse; this test is about semantic validation")

	return model
}

// validateDSL runs the full validation path on a DSL string, as a consumer would,
// and recovers the collection behind the returned error.
func validateDSL(t *testing.T, dsl string) *ValidationErrors {
	t.Helper()

	return findingsFrom(ValidateDSL(modelFromDSL(t, dsl), dsl, DefaultEngineOptions()))
}

// TestErrorsIsThroughValidation checks errors.Is against findings from real
// validation, so a caller can identify what went wrong without matching message
// text.
func TestErrorsIsThroughValidation(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		dsl          string
		wantSentinel error
		wantCategory fgaerrors.ModelErrorKind
		wantScope    func(t *testing.T, err error)
	}{
		"undefined type in restriction": {
			dsl: `model
  schema 1.1
type document
  relations
    define viewer: [user]
`,
			wantSentinel: fgaerrors.ErrInvalidType,
			wantCategory: fgaerrors.ErrorKindObjectType,
			wantScope: func(t *testing.T, err error) {
				t.Helper()

				var scoped *fgaerrors.ErrObjectType
				require.ErrorAs(t, err, &scoped)
				assert.Equal(t, "user", scoped.ObjectType)
			},
		},
		"relation with no entrypoint": {
			dsl: `model
  schema 1.1
type user
type document
  relations
    define viewer: writer
    define writer: viewer
`,
			wantSentinel: fgaerrors.ErrNoEntrypoints,
			wantCategory: fgaerrors.ErrorKindRelation,
			wantScope: func(t *testing.T, err error) {
				t.Helper()

				var scoped *fgaerrors.ErrRelation
				require.ErrorAs(t, err, &scoped)
				assert.Equal(t, "document", scoped.ObjectType)
				assert.NotEmpty(t, scoped.Relation)
			},
		},
		"duplicate type": {
			dsl: `model
  schema 1.1
type user
type document
type document
`,
			wantSentinel: fgaerrors.ErrDuplicateDefinition,
			wantCategory: fgaerrors.ErrorKindObjectType,
			wantScope: func(t *testing.T, err error) {
				t.Helper()

				// An ErrObjectType has no Relation field, so a duplicate type
				// cannot arrive carrying one.
				var scoped *fgaerrors.ErrObjectType
				require.ErrorAs(t, err, &scoped)
				assert.Equal(t, "document", scoped.ObjectType)
			},
		},
		"condition defined but unused": {
			dsl: `model
  schema 1.1
type user
type document
  relations
    define viewer: [user]

condition inRegion(x: string) {
  x == "eu"
}
`,
			wantSentinel: fgaerrors.ErrConditionUnReferenced,
			wantCategory: fgaerrors.ErrorKindCondition,
			wantScope: func(t *testing.T, err error) {
				t.Helper()

				// A condition definition is not scoped to a type, and
				// ErrCondition has no field for one.
				var scoped *fgaerrors.ErrCondition
				require.ErrorAs(t, err, &scoped)
				assert.Equal(t, "inRegion", scoped.Condition)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			validationErrors := validateDSL(t, test.dsl)
			require.NotNil(t, validationErrors)
			require.True(t, validationErrors.HasErrors(), "expected this model to fail validation")

			// Every finding, so a severity this test does not expect fails on the
			// severity assertion below rather than by going missing here.
			var matched *ValidationError
			for _, candidate := range validationErrors.AllFindings() {
				if errors.Is(candidate, test.wantSentinel) {
					matched = candidate

					break
				}
			}

			require.NotNilf(t, matched,
				"no finding matched %v via errors.Is; got %v", test.wantSentinel, validationErrors.Error())

			assert.Equal(t, test.wantCategory, matched.Category)
			assert.Equal(t, fgaerrors.SeverityError, matched.Severity)
			assert.True(t, matched.Blocks())

			require.Error(t, matched.Unwrap(), "errors.As must have a scoped cause to reach")
			test.wantScope(t, error(matched))
		})
	}
}

// TestMetadataIsDerivedFromCause checks the serialised metadata and the errors.As
// payload describe the same scope, so the two cannot drift.
func TestMetadataIsDerivedFromCause(t *testing.T) {
	t.Parallel()

	validationErrors := validateDSL(t, `model
  schema 1.1
type user
type document
  relations
    define viewer: writer
    define writer: viewer
`)
	require.True(t, validationErrors.HasErrors())

	checked := 0

	for _, validationErr := range validationErrors.AllFindings() {
		if validationErr.Unwrap() == nil {
			continue
		}

		objectType, relation, condition := causeScope(error(validationErr))

		require.NotNil(t, validationErr.Metadata)
		assert.Equal(t, objectType, validationErr.Metadata.Type,
			"metadata type must match the cause it was derived from")
		assert.Equal(t, relation, validationErr.Metadata.Relation)
		assert.Equal(t, condition, validationErr.Metadata.Condition)

		checked++
	}

	assert.Positive(t, checked, "no finding carried a scoped cause; the derivation was not exercised")
}

// TestEverySemanticFindingCarriesErrorInfo sweeps a range of broken models and
// asserts no finding escapes without severity, category and a matchable cause.
// A gap here means some code path bypasses the table.
func TestEverySemanticFindingCarriesErrorInfo(t *testing.T) {
	t.Parallel()

	models := []string{
		`model
  schema 1.1
type document
  relations
    define viewer: [user]
`,
		`model
  schema 1.1
type user
type document
type document
`,
		`model
  schema 1.1
type user
type document
  relations
    define viewer: writer
    define writer: viewer
`,
		`model
  schema 1.1
type user
type document
  relations
    define viewer: [user]

condition inRegion(x: string) {
  x == "eu"
}
`,
		`model
  schema 1.1
type user
type document
  relations
    define parent: [document]
    define viewer: viewer from parent
`,
	}

	total := 0

	for index, dsl := range models {
		model, err := transformer.TransformDSLToProto(dsl)

		// Not skipped: a model that stops parsing drops silently out of the sweep,
		// and the total below would still pass on the models that remain.
		require.NoErrorf(t, err, "model %d no longer parses", index)

		validationErrors := findingsFrom(ValidateDSL(model, dsl, DefaultEngineOptions()))

		for _, validationErr := range validationErrors.AllFindings() {
			total++

			require.NotNilf(t, validationErr.Metadata, "model %d: finding without metadata", index)

			errorType := validationErr.Metadata.ErrorType

			assert.NotEmptyf(t, validationErr.Severity,
				"model %d: %q has no severity", index, errorType)

			if _, classified := errorInfoByType[errorType]; classified {
				require.Errorf(t, validationErr.Unwrap(),
					"model %d: %q is in the errorInfoByType but carries no cause", index, errorType)
			}
		}
	}

	assert.Positive(t, total, "no findings produced; this test asserted nothing")
}

// TestNonBlockingTableEntryReachesTheCaller closes the gap the other severity tests
// leave: they assert what errorInfoByType holds, or build findings by hand, and every
// entry is SeverityError today, so nothing follows a non-blocking severity from the
// table through addScopedError and out of an entry point. This downgrades one entry and
// does exactly that.
//
// It must not call t.Parallel: it mutates errorInfoByType, and Go runs a sequential
// test only with other sequential tests.
func TestNonBlockingTableEntryReachesTheCaller(t *testing.T) {
	original := errorInfoByType[InvalidName]
	downgraded := original
	downgraded.Severity = fgaerrors.SeverityWarning
	errorInfoByType[InvalidName] = downgraded

	t.Cleanup(func() { errorInfoByType[InvalidName] = original })

	// A name the DSL parser would reject, so the model is built as the proto a JSON
	// caller would supply.
	model := &openfgav1.AuthorizationModel{
		SchemaVersion:   "1.1",
		TypeDefinitions: []*openfgav1.TypeDefinition{{Type: "Bad Type Name"}},
	}

	require.NoError(t, ValidateModelJSON(model),
		"a model whose only finding is a warning is valid, so the entry point returns nil")

	engine := NewValidationEngine(model, "")
	collection := engine.RunAllValidations(DefaultEngineOptions())

	require.Equal(t, 1, collection.CountAll(),
		"this model must raise exactly one finding, or the counts below are ambiguous")
	assert.Equal(t, 0, collection.Count())
	assert.False(t, collection.HasErrors())
	assert.True(t, collection.HasFindings())
	assert.Empty(t, collection.GetErrors())
	require.NoError(t, collection.ErrorOrNil())

	finding := collection.AllFindings()[0]
	assert.Equal(t, fgaerrors.SeverityWarning, finding.Severity, "the severity came from the table")
	assert.False(t, finding.Blocks())

	// Severity is independent of the cause: a warning still carries its sentinel and
	// its scope.
	require.ErrorIs(t, error(finding), fgaerrors.ErrInvalidName)

	var scoped *fgaerrors.ErrObjectType
	require.ErrorAs(t, error(finding), &scoped)
	assert.Equal(t, "Bad Type Name", scoped.ObjectType)

	summary := engine.GetValidationSummary()
	assert.Equal(t, 0, summary.TotalErrors)
	assert.Equal(t, 1, summary.TotalFindings)
	assert.Equal(t, 1, summary.FindingsBySeverity[fgaerrors.SeverityWarning])
	assert.False(t, summary.HasCriticalErrors)

	report := CreateValidationReport(model, "", DefaultEngineOptions())
	assert.True(t, report.IsValid(), "a warning does not invalidate the model")
	assert.Len(t, report.GetErrorsByType(InvalidName), 1,
		"GetErrorsByType names a code, so it returns the finding whatever its severity")
}
