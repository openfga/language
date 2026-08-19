package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	fgaerrors "github.com/openfga/language/pkg/go/errors"
)

// TestCriticalImpliesBlocking checks no code is critical without also being
// blocking. Criticality and severity are fields on the same errorInfo entry, so a
// code cannot claim to invalidate the whole model and not fail validation.
func TestCriticalImpliesBlocking(t *testing.T) {
	t.Parallel()

	for errorType, info := range errorInfoByType {
		if !info.Critical {
			continue
		}

		assert.Equalf(t, fgaerrors.SeverityError, info.Severity,
			"%q is critical but its severity is %q: a finding cannot invalidate the whole "+
				"model and leave it valid", errorType, info.Severity)
	}
}

// TestCriticalErrorTypesAreEmitted checks criticality is only claimed for a code
// some Raise* method raises.
//
// It reads the raise sites in the collector, not the callers of those methods, so a
// code raised only by a Raise* method that nothing calls still counts here.
func TestCriticalErrorTypesAreEmitted(t *testing.T) {
	t.Parallel()

	emitted := emittedErrorTypes(t)
	require.NotEmpty(t, emitted)

	for errorType, info := range errorInfoByType {
		if !info.Critical {
			continue
		}

		name := errorTypeConstantName(t, errorType)
		_, ok := emitted[name]
		assert.Truef(t, ok, "%s is marked critical but no Raise* method raises it", name)
	}
}

// TestUnemittedErrorTypesAreNotCritical checks the same from the other side, so a
// change that starts emitting one of these codes has to decide its criticality
// rather than inherit one.
func TestUnemittedErrorTypesAreNotCritical(t *testing.T) {
	t.Parallel()

	for errorType := range unemittedErrorTypes {
		assert.Falsef(t, isCriticalErrorType(errorType),
			"%q is not emitted, so calling it critical asserts nothing", errorType)
	}
}

// TestCriticalityOfEveryEmittedCode pins the criticality of every declared code, so
// a change to errorInfo that alters one has to be made here as well.
func TestCriticalityOfEveryEmittedCode(t *testing.T) {
	t.Parallel()

	wantCritical := map[ValidationErrorType]bool{
		RelationNoEntrypoint:  true,
		UndefinedType:         true,
		UndefinedRelation:     true,
		InvalidRelationType:   true,
		DuplicatedError:       true,
		InvalidSchema:         true,
		MultipleModulesInFile: true,
		GraphModelUnbuildable: true,
	}

	// Nothing raises these two, so they are held at not-critical rather than listed
	// above.
	neverRaised := map[ValidationErrorType]struct{}{
		CyclicRelation:       {},
		InvalidSchemaVersion: {},
	}

	for _, errorType := range allErrorTypes {
		if _, unemitted := neverRaised[errorType]; unemitted {
			assert.Falsef(t, isCriticalErrorType(errorType),
				"%q is never raised and must not be claimed critical", errorType)

			continue
		}

		assert.Equalf(t, wantCritical[errorType], isCriticalErrorType(errorType),
			"criticality of %q does not match the list in this test", errorType)
	}
}

// TestHasCriticalErrorsThroughValidation checks criticality end to end, and that a
// non-critical error leaves the flag unset. Otherwise the field would be
// indistinguishable from HasErrors.
func TestHasCriticalErrorsThroughValidation(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		dsl          string
		wantCritical bool
		wantValid    bool
	}{
		"undefined type is an error but not critical": {
			dsl: `model
  schema 1.1
type document
  relations
    define viewer: [user]
`,
			wantCritical: false,
			wantValid:    false,
		},
		"duplicate type is critical": {
			dsl: `model
  schema 1.1
type user
type document
type document
`,
			wantCritical: true,
			wantValid:    false,
		},
		"relation with no entrypoint is critical": {
			dsl: `model
  schema 1.1
type user
type document
  relations
    define viewer: writer
    define writer: viewer
`,
			wantCritical: true,
			wantValid:    false,
		},
		"valid model has neither": {
			dsl: `model
  schema 1.1
type user
type document
  relations
    define viewer: [user]
`,
			wantCritical: false,
			wantValid:    true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			report := CreateValidationReport(modelFromDSL(t, test.dsl), test.dsl, DefaultEngineOptions())

			assert.Equal(t, test.wantCritical, report.HasCriticalErrors())
			assert.Equal(t, test.wantValid, report.IsValid())

			if test.wantCritical {
				require.False(t, report.IsValid(), "a critical finding must also block")
			}
		})
	}
}
