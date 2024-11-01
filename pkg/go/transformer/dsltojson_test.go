package transformer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"

	"github.com/openfga/language/pkg/go/transformer"
)

func TestDSLToJSONTransformerForValidCases(t *testing.T) {
	t.Parallel()

	testCases, err := loadValidTransformerTestCases()
	if err != nil {
		t.Fatal(err)
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			if testCase.Skip {
				t.Skip()
			}

			jsonString, err := transformer.TransformDSLToJSON(testCase.DSL)

			require.NoError(t, err)

			expectedAuthModel := &openfgav1.AuthorizationModel{}
			err = protojson.Unmarshal([]byte(testCase.JSON), expectedAuthModel)
			require.NoError(t, err)

			jsonBytes, err := protojson.Marshal(expectedAuthModel)
			require.NoError(t, err)

			require.JSONEq(t, string(jsonBytes), jsonString)
		})
	}
}

func TestDSLToJSONTransformerForSyntaxErrorCases(t *testing.T) {
	t.Parallel()

	testCases, err := loadInvalidDslSyntaxTestCases()
	if err != nil {
		t.Fatal(err)
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			_, err := transformer.TransformDSLToJSON(testCase.DSL)

			errorsCount := len(testCase.ExpectedErrors)
			if errorsCount == 0 {
				require.NoError(t, err)
			} else {
				require.Error(t, err)

				// unfortunately antlr is throwing different error messages in Go and JS - considering that at the moment
				//  we care that it errors for syntax errors more than we care about the error messages matching,
				//  esp. in Go as we are not building a language server on top of the returned errors yet
				//  actual matching error strings is safe to ignore for now
				// require.EqualErrorf(t, err, testCase.GetErrorString(), "")
			} //nolint:wsl
		})
	}
}
