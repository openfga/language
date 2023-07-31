package transformer_test

import (
	"log"
	"testing"

	pb "buf.build/gen/go/openfga/api/protocolbuffers/go/openfga/v1"
	language "github.com/openfga/language/pkg/go/transformer"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestDSLToJSONTransformer(t *testing.T) {
	testCases, err := LoadValidTransformerTestCases()
	if err != nil {
		log.Fatal(err)
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

	testCases2, err := LoadInvalidDslSyntaxTestCases()
	if err != nil {
		log.Fatal(err)
	}

	for _, testCase := range testCases2 {
		t.Run(testCase.Name, func(t *testing.T) {
			_, err := language.TransformDslToJSON(testCase.DSL)

			if testCase.Valid {
				require.NoError(t, err)
			} else {
				require.EqualErrorf(t, err, testCase.ErrorMessage, "")
			}
		})
	}
}
