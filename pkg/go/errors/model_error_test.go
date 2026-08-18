package errors_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	fgaerrors "github.com/openfga/language/pkg/go/errors"
)

// TestErrorsIsReachesCauseThroughScope checks errors.Is finds the sentinel through
// the scope type, so a caller branches on it without matching message text.
func TestErrorsIsReachesCauseThroughScope(t *testing.T) {
	t.Parallel()

	err := &fgaerrors.ErrRelation{
		ObjectType: "document",
		Relation:   "viewer",
		Cause:      fgaerrors.ErrNoEntrypoints,
	}

	require.ErrorIs(t, err, fgaerrors.ErrNoEntrypoints)
	assert.NotErrorIs(t, err, fgaerrors.ErrReservedKeywords,
		"must not match a sentinel it does not wrap")
}

// TestErrorsIsWorksThroughFurtherWrapping checks the cause survives a caller
// adding its own context, which is the normal way these errors travel.
func TestErrorsIsWorksThroughFurtherWrapping(t *testing.T) {
	t.Parallel()

	inner := &fgaerrors.ErrObjectType{
		ObjectType: "document",
		Cause:      fgaerrors.ErrDuplicateDefinition,
	}
	outer := fmt.Errorf("validating model: %w", inner)

	require.ErrorIs(t, outer, fgaerrors.ErrDuplicateDefinition)

	var objectTypeErr *fgaerrors.ErrObjectType
	require.ErrorAs(t, outer, &objectTypeErr, "errors.As must find the scope through fmt.Errorf")
	assert.Equal(t, "document", objectTypeErr.ObjectType)
}

// TestErrorsAsExposesScope checks errors.As recovers each scope type with its
// fields intact, and that the message names them.
func TestErrorsAsExposesScope(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		err               error
		wantMessageSubstr string
		wantCause         error
		wantScope         func(t *testing.T, err error)
	}{
		"object type": {
			err: &fgaerrors.ErrObjectType{
				ObjectType: "document",
				Cause:      fgaerrors.ErrReservedKeywords,
			},
			wantMessageSubstr: "the object type 'document'",
			wantCause:         fgaerrors.ErrReservedKeywords,
			wantScope: func(t *testing.T, err error) {
				t.Helper()

				var scoped *fgaerrors.ErrObjectType
				require.ErrorAs(t, err, &scoped)
				assert.Equal(t, "document", scoped.ObjectType)
			},
		},
		"relation": {
			err: &fgaerrors.ErrRelation{
				ObjectType: "document",
				Relation:   "viewer",
				Cause:      fgaerrors.ErrNoEntrypoints,
			},
			wantMessageSubstr: "relation 'viewer' of object type 'document'",
			wantCause:         fgaerrors.ErrNoEntrypoints,
			wantScope: func(t *testing.T, err error) {
				t.Helper()

				var scoped *fgaerrors.ErrRelation
				require.ErrorAs(t, err, &scoped)
				assert.Equal(t, "document", scoped.ObjectType)
				assert.Equal(t, "viewer", scoped.Relation)
			},
		},
		"relation condition": {
			err: &fgaerrors.ErrRelationCondition{
				ObjectType: "document",
				Relation:   "viewer",
				Condition:  "inRegion",
				Cause:      fgaerrors.ErrConditionUndefined,
			},
			wantMessageSubstr: "condition 'inRegion' of relation 'viewer' in object type 'document'",
			wantCause:         fgaerrors.ErrConditionUndefined,
			wantScope: func(t *testing.T, err error) {
				t.Helper()

				var scoped *fgaerrors.ErrRelationCondition
				require.ErrorAs(t, err, &scoped)
				assert.Equal(t, "document", scoped.ObjectType)
				assert.Equal(t, "viewer", scoped.Relation)
				assert.Equal(t, "inRegion", scoped.Condition)
			},
		},
		"condition": {
			err: &fgaerrors.ErrCondition{
				Condition: "inRegion",
				Cause:     fgaerrors.ErrConditionUnReferenced,
			},
			wantMessageSubstr: "condition 'inRegion'",
			wantCause:         fgaerrors.ErrConditionUnReferenced,
			wantScope: func(t *testing.T, err error) {
				t.Helper()

				var scoped *fgaerrors.ErrCondition
				require.ErrorAs(t, err, &scoped)
				assert.Equal(t, "inRegion", scoped.Condition)
			},
		},
		"model": {
			err:               &fgaerrors.ErrModel{Cause: fgaerrors.ErrMultipleModulesInFile},
			wantMessageSubstr: "error in authorization model",
			wantCause:         fgaerrors.ErrMultipleModulesInFile,
			wantScope: func(t *testing.T, err error) {
				t.Helper()

				var scoped *fgaerrors.ErrModel
				require.ErrorAs(t, err, &scoped)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test.wantScope(t, test.err)

			assert.Contains(t, test.err.Error(), test.wantMessageSubstr)

			// The cause must survive into the message as well as into errors.Is;
			// a message naming a relation but not the problem is not actionable.
			assert.Contains(t, test.err.Error(), test.wantCause.Error())
			assert.ErrorIs(t, test.err, test.wantCause)
		})
	}
}

// TestScopedErrorsDoNotMatchEachOther checks the scopes are mutually exclusive:
// errors.As for one scope does not match a different one.
func TestScopedErrorsDoNotMatchEachOther(t *testing.T) {
	t.Parallel()

	conditionErr := error(&fgaerrors.ErrCondition{
		Condition: "inRegion",
		Cause:     fgaerrors.ErrConditionUndefined,
	})

	var relationConditionErr *fgaerrors.ErrRelationCondition
	assert.False(t, errors.As(conditionErr, &relationConditionErr),
		"a condition definition finding must not pass for one about a relation's condition")

	var relationErr *fgaerrors.ErrRelation
	assert.False(t, errors.As(conditionErr, &relationErr))
}

// TestRelationErrorWithoutObjectType covers relation-scoped raise sites that
// have no object type to attach: the message must not name an empty object
// type.
func TestRelationErrorWithoutObjectType(t *testing.T) {
	t.Parallel()

	err := &fgaerrors.ErrRelation{
		Relation: "self",
		Cause:    fgaerrors.ErrReservedKeywords,
	}

	assert.NotContains(t, err.Error(), "object type ''")
	assert.Contains(t, err.Error(), "relation 'self'")
	assert.Contains(t, err.Error(), fgaerrors.ErrReservedKeywords.Error())
}

func TestSeverityBlocks(t *testing.T) {
	t.Parallel()

	assert.True(t, fgaerrors.SeverityError.Blocks(), "an error must fail validation")
	assert.False(t, fgaerrors.SeverityWarning.Blocks(), "a warning must not fail a valid model")
	assert.False(t, fgaerrors.SeverityAdvisory.Blocks(), "an advisory must not fail a valid model")

	// Only the two non-blocking severities answer false, so anything unrecognised
	// blocks rather than letting an invalid model pass.
	assert.True(t, fgaerrors.SeverityUnspecified.Blocks(),
		"a finding that never set a severity must block")
	assert.True(t, fgaerrors.Severity(99).Blocks(),
		"a severity this package does not declare must block")
}

// TestSeverityWireNames checks each severity marshals and unmarshals under its wire
// name. The names are API: the Go severity fixtures assert on them.
func TestSeverityWireNames(t *testing.T) {
	t.Parallel()

	tests := map[fgaerrors.Severity]string{
		fgaerrors.SeverityError:    `"error"`,
		fgaerrors.SeverityWarning:  `"warning"`,
		fgaerrors.SeverityAdvisory: `"advisory"`,
	}

	for severity, wantJSON := range tests {
		t.Run(severity.String(), func(t *testing.T) {
			t.Parallel()

			encoded, err := json.Marshal(severity)
			require.NoError(t, err)
			assert.JSONEq(t, wantJSON, string(encoded))

			var decoded fgaerrors.Severity
			require.NoError(t, json.Unmarshal(encoded, &decoded))
			assert.Equal(t, severity, decoded, "round trip must be lossless")

			assert.True(t, severity.IsValid())
		})
	}
}

// TestSeverityRejectsUnknownValues covers both ends of the mapping: a number with
// no name must not reach a consumer as a bare integer, and a name this package
// does not declare must not decode to a severity.
func TestSeverityRejectsUnknownValues(t *testing.T) {
	t.Parallel()

	_, err := json.Marshal(fgaerrors.Severity(99))
	require.ErrorIs(t, err, fgaerrors.ErrUnknownSeverity,
		"an undeclared severity must fail to marshal rather than ship as 99")

	var decoded fgaerrors.Severity
	require.ErrorIs(t, json.Unmarshal([]byte(`"critical"`), &decoded),
		fgaerrors.ErrUnknownSeverity)

	assert.False(t, fgaerrors.Severity(99).IsValid())
	assert.False(t, fgaerrors.SeverityUnspecified.IsValid())
}

// TestSeverityUnspecifiedIsNotASeverity checks the zero value is not a severity.
// The constants count from one so a finding that never set a severity does not read
// as an error.
func TestSeverityUnspecifiedIsNotASeverity(t *testing.T) {
	t.Parallel()

	assert.NotEqual(t, fgaerrors.SeverityUnspecified, fgaerrors.SeverityError,
		"the first real severity must not be the zero value")
	assert.Equal(t, 0, int(fgaerrors.SeverityUnspecified))
	assert.Empty(t, fgaerrors.SeverityUnspecified.String())

	// omitempty drops it, so a finding with no severity ships without the field
	// rather than claiming to be an error.
	encoded, err := json.Marshal(struct {
		Severity fgaerrors.Severity `json:"severity,omitempty"`
	}{})
	require.NoError(t, err)
	assert.JSONEq(t, `{}`, string(encoded))
}

// TestSeverityIsAWireNameAsAMapKey covers ValidationSummary.FindingsBySeverity,
// which is keyed by Severity. For a map key the encoding/json package consults
// neither String nor MarshalJSON, only MarshalText, so the counts stay keyed by
// name rather than by number.
func TestSeverityIsAWireNameAsAMapKey(t *testing.T) {
	t.Parallel()

	encoded, err := json.Marshal(map[fgaerrors.Severity]int{
		fgaerrors.SeverityError:    3,
		fgaerrors.SeverityWarning:  1,
		fgaerrors.SeverityAdvisory: 2,
	})
	require.NoError(t, err)
	assert.JSONEq(t, `{"error":3,"warning":1,"advisory":2}`, string(encoded))

	// Decoding a map key takes the same route: the key type has to implement
	// UnmarshalText for encoding/json to reach it at all, otherwise an integer key
	// is parsed as a number and "error" fails.
	var decoded map[fgaerrors.Severity]int
	require.NoError(t, json.Unmarshal(encoded, &decoded))
	assert.Equal(t, map[fgaerrors.Severity]int{
		fgaerrors.SeverityError:    3,
		fgaerrors.SeverityWarning:  1,
		fgaerrors.SeverityAdvisory: 2,
	}, decoded)
}

// TestModelErrorKindWireNames checks each category marshals and unmarshals under
// its wire name. The names are API: the Go severity fixtures assert on them.
func TestModelErrorKindWireNames(t *testing.T) {
	t.Parallel()

	tests := map[fgaerrors.ModelErrorKind]string{
		fgaerrors.ErrorKindObjectType:        `"object-type"`,
		fgaerrors.ErrorKindRelation:          `"relation"`,
		fgaerrors.ErrorKindRelationCondition: `"relation-condition"`,
		fgaerrors.ErrorKindCondition:         `"condition"`,
		fgaerrors.ErrorKindInvalidModel:      `"invalid-model"`,
	}

	for errorType, wantJSON := range tests {
		t.Run(errorType.String(), func(t *testing.T) {
			t.Parallel()

			encoded, err := json.Marshal(errorType)
			require.NoError(t, err)
			assert.JSONEq(t, wantJSON, string(encoded))

			var decoded fgaerrors.ModelErrorKind
			require.NoError(t, json.Unmarshal(encoded, &decoded))
			assert.Equal(t, errorType, decoded, "round trip must be lossless")

			assert.True(t, errorType.IsValid())
		})
	}
}

// TestModelErrorKindRejectsUnknownValues covers both ends of the mapping: a
// number with no name must not reach a consumer as a bare integer, and a name
// this package does not declare must not decode to a category.
func TestModelErrorKindRejectsUnknownValues(t *testing.T) {
	t.Parallel()

	_, err := json.Marshal(fgaerrors.ModelErrorKind(99))
	require.ErrorIs(t, err, fgaerrors.ErrUnknownModelErrorKind,
		"an undeclared category must fail to marshal rather than ship as 99")

	var decoded fgaerrors.ModelErrorKind
	require.ErrorIs(t, json.Unmarshal([]byte(`"not-a-category"`), &decoded),
		fgaerrors.ErrUnknownModelErrorKind)

	assert.False(t, fgaerrors.ModelErrorKind(99).IsValid())
	assert.False(t, fgaerrors.ModelErrorKindUnspecified.IsValid())
}

// TestModelErrorKindUnspecifiedIsNotACategory checks the zero value is not a
// category. The constants count from one so a finding that never set one does not
// read as being about an object type.
func TestModelErrorKindUnspecifiedIsNotACategory(t *testing.T) {
	t.Parallel()

	assert.NotEqual(t, fgaerrors.ModelErrorKindUnspecified, fgaerrors.ErrorKindObjectType,
		"the first real category must not be the zero value")
	assert.Equal(t, 0, int(fgaerrors.ModelErrorKindUnspecified))
	assert.Empty(t, fgaerrors.ModelErrorKindUnspecified.String())

	// omitempty drops it, so a finding with no category ships without the field
	// rather than with a wrong one.
	encoded, err := json.Marshal(struct {
		Category fgaerrors.ModelErrorKind `json:"category,omitempty"`
	}{})
	require.NoError(t, err)
	assert.JSONEq(t, `{}`, string(encoded))
}

// TestSentinelsAreDistinct guards against a copy-paste leaving two names pointing
// at one value, which would make errors.Is match the wrong condition.
func TestSentinelsAreDistinct(t *testing.T) {
	t.Parallel()

	sentinels := map[string]error{
		"ErrInvalidSchemaVersion":               fgaerrors.ErrInvalidSchemaVersion,
		"ErrSchemaVersionUnsupported":           fgaerrors.ErrSchemaVersionUnsupported,
		"ErrSchemaVersionRequired":              fgaerrors.ErrSchemaVersionRequired,
		"ErrReservedKeywords":                   fgaerrors.ErrReservedKeywords,
		"ErrInvalidName":                        fgaerrors.ErrInvalidName,
		"ErrDuplicateDefinition":                fgaerrors.ErrDuplicateDefinition,
		"ErrObjectTypeUndefined":                fgaerrors.ErrObjectTypeUndefined,
		"ErrRelationUndefined":                  fgaerrors.ErrRelationUndefined,
		"ErrInvalidType":                        fgaerrors.ErrInvalidType,
		"ErrInvalidRelationType":                fgaerrors.ErrInvalidRelationType,
		"ErrInvalidRelationOnTupleset":          fgaerrors.ErrInvalidRelationOnTupleset,
		"ErrInvalidRelationOnTuplesetNotDirect": fgaerrors.ErrInvalidRelationOnTuplesetNotDirect,
		"ErrNoEntrypoints":                      fgaerrors.ErrNoEntrypoints,
		"ErrDirectlyAssignableRelation":         fgaerrors.ErrDirectlyAssignableRelation,
		"ErrInvalidWildcard":                    fgaerrors.ErrInvalidWildcard,
		"ErrConditionUndefined":                 fgaerrors.ErrConditionUndefined,
		"ErrConditionUnReferenced":              fgaerrors.ErrConditionUnReferenced,
		"ErrConditionNameMismatch":              fgaerrors.ErrConditionNameMismatch,
		"ErrMultipleModulesInFile":              fgaerrors.ErrMultipleModulesInFile,
	}

	seenMessages := make(map[string]string, len(sentinels))

	for name, sentinel := range sentinels {
		require.Errorf(t, sentinel, "%s is nil", name)

		for otherName, other := range sentinels {
			if name == otherName {
				continue
			}

			require.NotErrorIsf(t, sentinel, other,
				"%s and %s are the same value; errors.Is cannot tell them apart", name, otherName)
		}

		if previous, duplicate := seenMessages[sentinel.Error()]; duplicate {
			t.Errorf("%s and %s have the identical message %q", name, previous, sentinel.Error())
		}

		seenMessages[sentinel.Error()] = name
	}
}
