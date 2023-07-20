package transformer_test

import (
	"log"
	"testing"

	language "github.com/openfga/language/pkg/go/transformer"
	"github.com/stretchr/testify/require"
)

func TestJSONToDSLTransformer(t *testing.T) {
	testCases, err := LoadTransformerTestCases()
	if err != nil {
		log.Fatal(err)
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.Skip {
				t.Skip()
			}

			dsl, err := language.TransformJSONStringToDSL(testCase.JSON)
			require.NoError(t, err)

			require.Equal(t, testCase.DSL, *dsl)
		})
	}
}
