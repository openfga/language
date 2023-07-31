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
	testCases, err := LoadTransformerTestCases()
	if err != nil {
		log.Fatal(err)
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.Skip {
				t.Skip()
			}

			model := language.TransformDslToJSON(testCase.DSL)

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
