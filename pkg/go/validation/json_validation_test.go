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

// jsonTestCase is one case from tests/data/json-validation-cases.yaml, the shared
// cases for models supplied as JSON. A JSON model carries no source text to resolve a
// position against, so no case in it states a line or a column.
type jsonTestCase struct {
	Name           string              `yaml:"name"`
	JSON           string              `yaml:"json"`
	Skip           bool                `yaml:"skip,omitempty"`
	ExpectedErrors []YAMLExpectedError `yaml:"expected_errors,omitempty"`
}

// TestJSONValidationSharedCases runs the shared JSON cases against ValidateJSON. The
// same file is run by pkg/js and pkg/java, so what only a JSON model can reach, such as
// findings with no position, is checked the same way in all three.
func TestJSONValidationSharedCases(t *testing.T) {
	t.Parallel()

	data, err := os.ReadFile(filepath.Join(sharedDataDir, "json-validation-cases.yaml"))
	require.NoError(t, err)

	var cases []jsonTestCase
	require.NoError(t, yaml.Unmarshal(data, &cases))
	require.NotEmpty(t, cases, "shared test file loaded no cases")

	for _, testCase := range cases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			if testCase.Skip {
				t.Skip("marked skip in the shared test file")
			}

			model := &openfgav1.AuthorizationModel{}

			// DiscardUnknown: the test file is shared with implementations whose model
			// type may have fields this one does not, and an unknown field is not what
			// a case is testing.
			require.NoError(t,
				protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(testCase.JSON), model),
				"the case's JSON must parse as an authorization model")

			result := compareWithExpected(testCase.ExpectedErrors, ExtractAllAs[*Finding](ValidateJSON(model)))

			for _, problem := range result.Problems {
				t.Error(problem)
			}

			require.Equal(t, casePass, result.Status)
		})
	}
}
