package validation

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// sharedDataDir is the shared test data directory, which sits at the repository root and
// is read by the Go, JS and Java implementations alike.
var sharedDataDir = filepath.Join("..", "..", "..", "tests", "data")

// TestSemanticValidationSharedCases runs every case in the shared semantic validation
// file. It fails if the file cannot be loaded or has no cases. Cases the file
// marks skip are reported as skipped.
func TestSemanticValidationSharedCases(t *testing.T) {
	t.Parallel()

	runner := NewYAMLTestRunner(sharedDataDir)

	suite, err := runner.LoadTestSuite("dsl-semantic-validation-cases.yaml")
	require.NoError(t, err)
	require.NotEmpty(t, suite.TestCases, "shared test file loaded no cases")

	for _, testCase := range suite.TestCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			result := runner.RunTestCase(testCase)
			if result.Status == caseSkipped {
				t.Skip("marked skip in the shared test file")
			}

			for _, problem := range result.Problems {
				t.Error(problem)
			}

			// Catches a failed status with no problems listed.
			require.Equal(t, casePass, result.Status)
		})
	}
}

// TestCompareWithExpected tests the comparison every shared case relies on. If it
// accepted a prefix of the message, or a finding with no position, every shared case
// would still pass.
func TestCompareWithExpected(t *testing.T) {
	t.Parallel()

	expected := YAMLExpectedError{
		Message: "the relation `viewer` does not exist.",
		Line:    &YAMLRange{Start: 4, End: 4},
		Column:  &YAMLRange{Start: 12, End: 18},
		Metadata: YAMLErrorMetadata{
			Symbol:    "viewer",
			ErrorType: string(MissingDefinition),
		},
	}

	finding := func(mutate func(*Finding)) *Finding {
		found := &Finding{
			Message: "the relation `viewer` does not exist.",
			Line:    &Range{Start: 4, End: 4},
			Column:  &Range{Start: 12, End: 18},
			Metadata: Metadata{
				Symbol: "viewer",
				Kind:   MissingDefinition,
			},
		}
		if mutate != nil {
			mutate(found)
		}

		return found
	}

	tests := []struct {
		name     string
		expected []YAMLExpectedError
		findings []*Finding
		problems int
	}{
		{
			name:     "match",
			expected: []YAMLExpectedError{expected},
			findings: []*Finding{finding(nil)},
		},
		{
			name:     "no errors expected and none found",
			expected: nil,
			findings: nil,
		},
		{
			// A message that only starts with the expected one is a different message, so
			// it satisfies nothing and is reported twice: the expectation went
			// unmatched and the finding was not expected.
			name:     "message is longer than the case states",
			expected: []YAMLExpectedError{expected},
			findings: []*Finding{finding(func(f *Finding) {
				f.Message += " Did you mean `view`?"
			})},
			problems: 2,
		},
		{
			name:     "wrong line",
			expected: []YAMLExpectedError{expected},
			findings: []*Finding{finding(func(f *Finding) {
				f.Line = &Range{Start: 5, End: 5}
			})},
			problems: 1,
		},
		{
			name:     "line end differs",
			expected: []YAMLExpectedError{expected},
			findings: []*Finding{finding(func(f *Finding) {
				f.Line = &Range{Start: 4, End: 6}
			})},
			problems: 1,
		},
		{
			name:     "no position at all",
			expected: []YAMLExpectedError{expected},
			findings: []*Finding{finding(func(f *Finding) {
				f.Line, f.Column = nil, nil
			})},
			problems: 1,
		},
		{
			name:     "wrong column",
			expected: []YAMLExpectedError{expected},
			findings: []*Finding{finding(func(f *Finding) {
				f.Column = &Range{Start: 12, End: 17}
			})},
			problems: 1,
		},
		{
			name:     "wrong symbol",
			expected: []YAMLExpectedError{expected},
			findings: []*Finding{finding(func(f *Finding) {
				f.Metadata.Symbol = "editor"
			})},
			problems: 1,
		},
		{
			name:     "wrong error type",
			expected: []YAMLExpectedError{expected},
			findings: []*Finding{finding(func(f *Finding) {
				f.Metadata.Kind = DuplicatedError
			})},
			problems: 1,
		},
		{
			name:     "one finding does not satisfy two expectations",
			expected: []YAMLExpectedError{expected, expected},
			findings: []*Finding{finding(nil)},
			problems: 1,
		},
		{
			name:     "finding the case does not expect",
			expected: []YAMLExpectedError{expected},
			findings: []*Finding{finding(nil), finding(func(f *Finding) {
				f.Message = "the relation `editor` does not exist."
			})},
			problems: 1,
		},
		{
			name:     "position the case leaves out is not compared",
			expected: []YAMLExpectedError{{Message: expected.Message}},
			findings: []*Finding{finding(func(f *Finding) {
				f.Line, f.Column = nil, nil
			})},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result := compareWithExpected(test.expected, test.findings)

			require.Len(t, result.Problems, test.problems, "problems: %v", result.Problems)

			if test.problems == 0 {
				require.Equal(t, casePass, result.Status)
			} else {
				require.Equal(t, caseFail, result.Status)
			}
		})
	}
}

// TestRunTestCaseUnparsableDSL covers the one outcome the shared cases cannot reach: a
// case whose DSL the transformer rejects tested nothing about validation, so it is
// reported as an error rather than passing for lack of findings.
func TestRunTestCaseUnparsableDSL(t *testing.T) {
	t.Parallel()

	result := NewYAMLTestRunner(sharedDataDir).RunTestCase(YAMLTestCase{
		Name: "not a model",
		DSL:  "type document\n  relations\n",
	})

	require.Equal(t, caseError, result.Status)
	require.Len(t, result.Problems, 1)
	require.Contains(t, result.Problems[0], "DSL does not parse")
}
