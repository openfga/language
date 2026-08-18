package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	fgaerrors "github.com/openfga/language/pkg/go/errors"
)

// No validation emits a non-blocking severity yet: every errorInfoByType entry is
// SeverityError. The tests below build findings directly, which is the only way to
// pin the severity predicates while nothing emits one.

func finding(severity fgaerrors.Severity, message string) *ValidationError {
	return &ValidationError{
		Message:  message,
		Severity: severity,
		Metadata: &ErrorMetadata{ErrorType: RelationNoEntrypoint},
	}
}

func TestPredicatesCountBlockingOnly(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		findings      []*ValidationError
		wantHasErrors bool
		wantCount     int
		wantAllCount  int
	}{
		"nothing at all": {
			findings:      nil,
			wantHasErrors: false,
			wantCount:     0,
			wantAllCount:  0,
		},
		"advisory only": {
			findings:      []*ValidationError{finding(fgaerrors.SeverityAdvisory, "check may answer differently")},
			wantHasErrors: false,
			wantCount:     0,
			wantAllCount:  1,
		},
		"warning only": {
			findings:      []*ValidationError{finding(fgaerrors.SeverityWarning, "uses a construct a future version may reject")},
			wantHasErrors: false,
			wantCount:     0,
			wantAllCount:  1,
		},
		"one error among non-blocking": {
			findings: []*ValidationError{
				finding(fgaerrors.SeverityAdvisory, "advisory"),
				finding(fgaerrors.SeverityError, "real error"),
				finding(fgaerrors.SeverityWarning, "warning"),
			},
			wantHasErrors: true,
			wantCount:     1,
			wantAllCount:  3,
		},
		"severity unset counts as blocking": {
			findings:      []*ValidationError{{Message: "built by hand"}},
			wantHasErrors: true,
			wantCount:     1,
			wantAllCount:  1,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			validationErrors := NewValidationErrors(test.findings)

			assert.Equal(t, test.wantHasErrors, validationErrors.HasErrors())
			assert.Equal(t, test.wantCount, validationErrors.Count())
			assert.Equal(t, test.wantAllCount, validationErrors.CountAll())
			assert.Equal(t, test.wantAllCount > 0, validationErrors.HasFindings())
			assert.Len(t, validationErrors.GetErrors(), test.wantCount)
			assert.Len(t, validationErrors.AllFindings(), test.wantAllCount)
		})
	}
}

// TestValidModelWithAdvisoryIsStillValid checks an advisory does not fail a model.
// An advisory describes a model that is correct, so reporting one must not
// invalidate it.
func TestValidModelWithAdvisoryIsStillValid(t *testing.T) {
	t.Parallel()

	report := ValidationReport{
		ValidationErrors: NewValidationErrors([]*ValidationError{
			finding(fgaerrors.SeverityAdvisory, "a check against this model may answer differently"),
			finding(fgaerrors.SeverityWarning, "this model uses a construct a future version may reject"),
		}),
	}

	assert.True(t, report.IsValid(), "warnings and advisories must not make a model invalid")
	assert.True(t, report.ValidationErrors.HasFindings(), "but they must still be reported")
	assert.Equal(t, "no validation errors", report.ValidationErrors.Error(),
		"the error string describes blocking findings, and there are none")
}

// TestErrorOrNilIsNilWhenNothingBlocks is the entry-point half of the same rule:
// a model whose only findings are warnings or advisories is valid, so err != nil
// means the model is invalid and not that something was reported.
func TestErrorOrNilIsNilWhenNothingBlocks(t *testing.T) {
	t.Parallel()

	nonBlocking := NewValidationErrors([]*ValidationError{
		finding(fgaerrors.SeverityWarning, "this model uses a construct a future version may reject"),
		finding(fgaerrors.SeverityAdvisory, "a check against this model may answer differently"),
	})

	// A literal nil, not a nil *ValidationErrors: the latter is a non-nil error
	// however few findings it holds.
	require.NoError(t, nonBlocking.ErrorOrNil())
	assert.True(t, nonBlocking.HasFindings(), "the findings are still there to be reported")
}

func TestErrorOrNilCarriesEverythingWhenSomethingBlocks(t *testing.T) {
	t.Parallel()

	blocking := NewValidationErrors([]*ValidationError{
		finding(fgaerrors.SeverityWarning, "warning"),
		finding(fgaerrors.SeverityError, "boom"),
	})

	err := blocking.ErrorOrNil()
	require.Error(t, err)

	var recovered *ValidationErrors
	require.ErrorAs(t, err, &recovered)
	assert.Len(t, recovered.AllFindings(), 2, "the non-blocking finding travels with the blocking one")
}

// TestErrorOrNilOnANilCollection covers the nil receiver, since the entry points
// call it on whatever RunAllValidations handed back.
func TestErrorOrNilOnANilCollection(t *testing.T) {
	t.Parallel()

	var nilCollection *ValidationErrors

	assert.NoError(t, nilCollection.ErrorOrNil())
}

// TestUnwrapReachesEveryFinding checks errors.Is sees a non-blocking finding too.
// Unwrap answers "was this condition reported", which is a different question from
// "does the model still validate".
func TestUnwrapReachesEveryFinding(t *testing.T) {
	t.Parallel()

	warning := finding(fgaerrors.SeverityWarning, "warned")
	warning.Cause = &fgaerrors.ErrRelation{ObjectType: "document", Relation: "viewer", Cause: fgaerrors.ErrNoEntrypoints}

	collection := NewValidationErrors([]*ValidationError{
		finding(fgaerrors.SeverityError, "boom"),
		warning,
	})

	require.ErrorIs(t, collection, fgaerrors.ErrNoEntrypoints,
		"the sentinel of a warning must still be reachable")

	var scoped *fgaerrors.ErrRelation
	require.ErrorAs(t, collection, &scoped)
	assert.Equal(t, "viewer", scoped.Relation)
}

// TestUnwrapSkipsNilFindings checks a directly-constructed collection holding a nil
// entry does not panic: a nil *ValidationError handed to errors.Is as a non-nil
// error would dereference nil on Unwrap.
func TestUnwrapSkipsNilFindings(t *testing.T) {
	t.Parallel()

	collection := NewValidationErrors([]*ValidationError{nil, finding(fgaerrors.SeverityError, "boom")})

	assert.Len(t, collection.Unwrap(), 1)
	assert.NotErrorIs(t, collection, fgaerrors.ErrNoEntrypoints)
}

func TestBlockingFindingMakesModelInvalid(t *testing.T) {
	t.Parallel()

	report := ValidationReport{
		ValidationErrors: NewValidationErrors([]*ValidationError{
			finding(fgaerrors.SeverityAdvisory, "advisory"),
			finding(fgaerrors.SeverityError, "boom"),
		}),
	}

	assert.False(t, report.IsValid())
	assert.Contains(t, report.ValidationErrors.Error(), "1 error occurred",
		"the count must agree with Count(), not with len(Errors)")
	assert.NotContains(t, report.ValidationErrors.Error(), "advisory")
}

// TestCascadeGateIgnoresNonBlockingFindings checks the gate in RunAllValidations
// counts only blocking findings. If it counted advisories, one advisory raised early
// would skip every gated phase and hide the errors they would have found.
func TestCascadeGateIgnoresNonBlockingFindings(t *testing.T) {
	t.Parallel()

	collector := NewErrorCollector(nil)
	collector.errors = append(collector.errors,
		finding(fgaerrors.SeverityAdvisory, "advisory"),
		finding(fgaerrors.SeverityWarning, "warning"),
	)

	require.False(t, collector.HasErrors(),
		"a collector holding only non-blocking findings must not close the cascade gate")
	assert.Equal(t, 0, collector.Count())
	assert.Equal(t, 2, collector.CountAll())
	assert.Len(t, collector.AllFindings(), 2,
		"the collector is the raw record and filters nothing")

	collector.errors = append(collector.errors, finding(fgaerrors.SeverityError, "real error"))
	assert.True(t, collector.HasErrors(), "a blocking finding must close the gate")
}

// TestSummarySplitsBySeverity checks the summary reports both totals, so a
// consumer can say "3 findings, 1 of which fails the model".
func TestSummarySplitsBySeverity(t *testing.T) {
	t.Parallel()

	engine := &ValidationEngine{collector: NewErrorCollector(nil)}
	engine.collector.errors = append(engine.collector.errors,
		finding(fgaerrors.SeverityError, "error"),
		finding(fgaerrors.SeverityWarning, "warning"),
		finding(fgaerrors.SeverityAdvisory, "advisory"),
	)

	summary := engine.GetValidationSummary()

	assert.Equal(t, 1, summary.TotalErrors, "only blocking findings are errors")
	assert.Equal(t, 3, summary.TotalFindings)
	assert.Equal(t, 1, summary.FindingsBySeverity[fgaerrors.SeverityError])
	assert.Equal(t, 1, summary.FindingsBySeverity[fgaerrors.SeverityWarning])
	assert.Equal(t, 1, summary.FindingsBySeverity[fgaerrors.SeverityAdvisory])
	assert.Equal(t, 3, summary.ErrorsByType[RelationNoEntrypoint],
		"the by-type breakdown covers every finding, so it sums to TotalFindings")
}

// TestGetErrorsByTypeIgnoresSeverity checks the deliberate exception: the caller
// asked for a specific code, so filtering by severity as well would drop matches it
// explicitly requested.
func TestGetErrorsByTypeIgnoresSeverity(t *testing.T) {
	t.Parallel()

	report := ValidationReport{
		ValidationErrors: NewValidationErrors([]*ValidationError{
			finding(fgaerrors.SeverityAdvisory, "advisory"),
			finding(fgaerrors.SeverityError, "error"),
		}),
	}

	assert.Len(t, report.GetErrorsByType(RelationNoEntrypoint), 2)
	assert.Empty(t, report.GetErrorsByType(UndefinedType))
}

// TestRealValidationStillFails checks the severity predicates do not stop real
// errors from counting. Every errorInfoByType entry is blocking, so real validation
// must still fail.
func TestRealValidationStillFails(t *testing.T) {
	t.Parallel()

	validationErrors := validateDSL(t, `model
  schema 1.1
type document
  relations
    define viewer: [user]
`)

	require.True(t, validationErrors.HasErrors(), "an undefined type must still fail validation")
	assert.Positive(t, validationErrors.Count())
	assert.Equal(t, validationErrors.CountAll(), validationErrors.Count(),
		"nothing emits a non-blocking severity yet, so the two counts must agree")
}

// TestPredicatesSurviveAWholeCollectionOfNothing pins the reads that a caller can
// reach without going through the collector: a nil collection, and one holding a nil
// finding. Every read goes through ValidationErrors.findings, so this covers the set.
func TestPredicatesSurviveAWholeCollectionOfNothing(t *testing.T) {
	t.Parallel()

	var absent *ValidationErrors

	assert.False(t, absent.HasErrors())
	assert.False(t, absent.HasFindings())
	assert.Equal(t, 0, absent.Count())
	assert.Equal(t, 0, absent.CountAll())
	assert.Empty(t, absent.GetErrors())
	assert.Empty(t, absent.AllFindings())
	assert.Empty(t, absent.Unwrap())
	require.NoError(t, absent.ErrorOrNil())
	assert.Equal(t, "no validation errors", absent.Error())

	// A nil finding is not a finding: every read drops it, so the counts agree with
	// each other and nothing hands a caller an entry that dereferences nil.
	held := NewValidationErrors([]*ValidationError{nil, finding(fgaerrors.SeverityError, "real")})

	assert.True(t, held.HasErrors())
	assert.True(t, held.HasFindings())
	assert.Equal(t, 1, held.Count())
	assert.Equal(t, 1, held.CountAll(), "the nil entry is not a finding to count")
	assert.Len(t, held.GetErrors(), 1)
	assert.Len(t, held.AllFindings(), 1)
	assert.Len(t, held.Unwrap(), 1)
	assert.Contains(t, held.Error(), "real")

	// Every entry AllFindings returns is safe to dereference, which is why the nil is
	// dropped rather than counted.
	for _, f := range held.AllFindings() {
		assert.Equal(t, fgaerrors.SeverityError, f.Severity)
		assert.Contains(t, f.String(), "real")
	}

	// A collection of nothing but nil reports nothing, rather than reporting a count
	// while Unwrap and Error report none.
	onlyNil := NewValidationErrors([]*ValidationError{nil})

	assert.False(t, onlyNil.HasFindings(), "a nil entry is not something reported")
	assert.Equal(t, 0, onlyNil.CountAll())
	assert.Empty(t, onlyNil.AllFindings())
	require.NoError(t, onlyNil.ErrorOrNil())

	// Add is the other way in.
	added := NewValidationErrors(nil)
	added.Add(nil)
	assert.False(t, added.HasFindings())
	assert.Equal(t, 0, added.CountAll())

	// A zero report reaches a nil collection through IsValid.
	var report ValidationReport

	assert.True(t, report.IsValid())
	assert.False(t, report.HasCriticalErrors())
	assert.Empty(t, report.GetErrorsByType(UndefinedType))

	// GetErrorsByType reads the code off metadata, which a hand-built finding can omit.
	withoutMetadata := ValidationReport{
		ValidationErrors: NewValidationErrors([]*ValidationError{nil, {Message: "no metadata"}}),
	}
	assert.Empty(t, withoutMetadata.GetErrorsByType(UndefinedType))
}
