package transformer_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	language "github.com/openfga/language/pkg/go/transformer"
)

func TestJSONToDSLTransformer(t *testing.T) {
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

			dsl, err := language.TransformJSONStringToDSL(testCase.JSON)
			require.NoError(t, err)

			require.Equal(t, testCase.DSL, *dsl)
		})
	}
}

func TestJSONToDSLTransformerForSyntaxErrorCases(t *testing.T) {
	t.Parallel()

	testCases, err := loadInvalidJSONSyntaxTestCases()
	if err != nil {
		t.Fatal(err)
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			_, err := language.TransformJSONStringToDSL(testCase.JSON)

			if testCase.ErrorMessage == "" {
				require.NoError(t, err)
			} else {
				require.EqualErrorf(t, err, testCase.ErrorMessage, "")
			}
		})
	}
}

func TestJSONToDSLTransformerForModularModelCases(t *testing.T) {
	t.Parallel()

	testCases, err := loadModuleTestCases()
	require.NoError(t, err)

	for _, testCase := range testCases {
		if testCase.DSL == "" {
			continue
		}

		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			dsl, err := language.TransformJSONStringToDSL(testCase.JSON)
			require.NoError(t, err)

			require.Equal(t, testCase.DSL, *dsl)
		})
	}
}

func TestJSONToDSLTransformerForModularModelCasesWithSourceInfo(t *testing.T) {
	t.Parallel()

	testCases, err := loadModuleTestCases()
	require.NoError(t, err)

	for _, testCase := range testCases {
		if testCase.DSLWithSourceInfo == "" {
			continue
		}

		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			dsl, err := language.TransformJSONStringToDSL(testCase.JSON, language.WithIncludeSourceInformation(true))
			require.NoError(t, err)

			require.Equal(t, testCase.DSLWithSourceInfo, *dsl)
		})
	}
}
