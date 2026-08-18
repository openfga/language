package validation

import (
	"os"
	"path/filepath"
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
	"gopkg.in/yaml.v3"
)

// jsonCorpusCase is one case from tests/data/json-validation-cases.yaml, the corpus for
// models supplied as JSON. A JSON model carries no source text to resolve a position
// against, so no case in it states a line or a column.
type jsonCorpusCase struct {
	Name           string              `yaml:"name"`
	JSON           string              `yaml:"json"`
	Skip           bool                `yaml:"skip,omitempty"`
	ExpectedErrors []YAMLExpectedError `yaml:"expected_errors,omitempty"`
}

// TestJSONValidationCorpus runs the JSON corpus against ValidateJSON. The JS and Java
// validators both consume this file and nothing in Go read it, so a rule that only a
// JSON model can reach was pinned in the other two SDKs and free to drift here.
func TestJSONValidationCorpus(t *testing.T) {
	t.Parallel()

	data, err := os.ReadFile(filepath.Join(corpusDir, "json-validation-cases.yaml"))
	require.NoError(t, err, "the corpus must be readable; a missing file is not a pass")

	var cases []jsonCorpusCase
	require.NoError(t, yaml.Unmarshal(data, &cases))
	require.NotEmpty(t, cases, "an empty corpus would assert nothing")

	for _, testCase := range cases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			if testCase.Skip {
				t.Skip("the corpus marks this case skipped")
			}

			model := &openfgav1.AuthorizationModel{}

			// DiscardUnknown: the corpus is shared with implementations whose model
			// type may have fields this one does not, and an unknown field is not what
			// a case is testing.
			require.NoError(t,
				protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(testCase.JSON), model),
				"the case's JSON must parse as an authorization model")

			result := compareWithCorpus(testCase.ExpectedErrors,
				findingsFrom(ValidateJSON(model, DefaultEngineOptions())))

			for _, problem := range result.Problems {
				t.Error(problem)
			}

			require.Equal(t, corpusPass, result.Status)
		})
	}
}
