package validation

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// corpusDir is the shared test data directory, which sits at the repository root and
// is read by the Go, JS and Java implementations alike.
var corpusDir = filepath.Join("..", "..", "..", "tests", "data")

// TestSemanticValidationCorpus runs every case in the shared semantic validation
// corpus.
//
// A corpus that cannot be loaded fails the test rather than skipping it: a corpus that
// silently does not run looks exactly like one that passes. The cases the corpus itself
// marks skip are skipped, so they stay visible in the output.
func TestSemanticValidationCorpus(t *testing.T) {
	t.Parallel()

	runner := NewYAMLTestRunner(corpusDir)

	suite, err := runner.LoadTestSuite("dsl-semantic-validation-cases.yaml")
	require.NoError(t, err)
	require.NotEmpty(t, suite.TestCases, "corpus loaded no cases")

	for _, testCase := range suite.TestCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			result := runner.RunTestCase(testCase)
			if result.Status == corpusSkipped {
				t.Skip("marked skip in the corpus")
			}

			for _, problem := range result.Problems {
				t.Error(problem)
			}

			// The problems above are the readable failure; this catches a status that
			// reported none.
			require.Equal(t, corpusPass, result.Status)
		})
	}
}

// TestCompareWithCorpus covers the comparison itself. It is what decides whether the
// corpus passes, so a change that loosened it — comparing a prefix of the message,
// treating a finding with no position as a match — would take the whole corpus green
// with it.
func TestCompareWithCorpus(t *testing.T) {
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
			// A message that only starts with the corpus's is a different message, so
			// it satisfies nothing and is reported twice: the expectation went
			// unmatched and the finding was not expected.
			name:     "message is longer than the corpus states",
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
				f.Metadata.Kind = UndefinedRelation
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
			name:     "finding the corpus does not expect",
			expected: []YAMLExpectedError{expected},
			findings: []*Finding{finding(nil), finding(func(f *Finding) {
				f.Message = "the relation `editor` does not exist."
			})},
			problems: 1,
		},
		{
			name:     "position the corpus leaves out is not compared",
			expected: []YAMLExpectedError{{Message: expected.Message}},
			findings: []*Finding{finding(func(f *Finding) {
				f.Line, f.Column = nil, nil
			})},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result := compareWithCorpus(test.expected, test.findings)

			require.Len(t, result.Problems, test.problems, "problems: %v", result.Problems)

			if test.problems == 0 {
				require.Equal(t, corpusPass, result.Status)
			} else {
				require.Equal(t, corpusFail, result.Status)
			}
		})
	}
}

// TestRunTestCaseUnparsableDSL covers the one outcome the corpus cases cannot reach: a
// case whose DSL the transformer rejects tested nothing about validation, so it is
// reported as an error rather than passing for lack of findings.
func TestRunTestCaseUnparsableDSL(t *testing.T) {
	t.Parallel()

	result := NewYAMLTestRunner(corpusDir).RunTestCase(YAMLTestCase{
		Name: "not a model",
		DSL:  "type document\n  relations\n",
	})

	require.Equal(t, corpusError, result.Status)
	require.Len(t, result.Problems, 1)
	require.Contains(t, result.Problems[0], "DSL does not parse")
}
