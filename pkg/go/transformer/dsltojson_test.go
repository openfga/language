package transformer_test

import (
	"testing"

	pb "github.com/openfga/api/proto/openfga/v1"
	language "github.com/openfga/language/pkg/go/transformer"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestDSLToJSONTransformerForValidCases(t *testing.T) {
	testCases, err := loadValidTransformerTestCases()
	if err != nil {
		t.Fatal(err)
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.Skip {
				t.Skip()
			}

			model, err := language.TransformDslToJSON(testCase.DSL)

			require.NoError(t, err)

			bytes, err := protojson.Marshal(model)
			require.NoError(t, err)

			expectedAuthModel := &pb.AuthorizationModel{}
			err = protojson.Unmarshal([]byte(testCase.JSON), expectedAuthModel)
			require.NoError(t, err)

			jsonBytes, err := protojson.Marshal(expectedAuthModel)
			require.NoError(t, err)

			require.JSONEq(t, string(jsonBytes), string(bytes))
		})
	}
}

func TestDSLToJSONTransformerForSyntaxErrorCases(t *testing.T) {
	testCases, err := loadInvalidDslSyntaxTestCases()
	if err != nil {
		t.Fatal(err)
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			_, err := language.TransformDslToJSON(testCase.DSL)

			errorsCount := len(testCase.ExpectedErrors)
			if errorsCount == 0 {
				require.NoError(t, err)
			} else {
				require.EqualErrorf(t, err, testCase.GetErrorString(), "")
			}
		})
	}
}
