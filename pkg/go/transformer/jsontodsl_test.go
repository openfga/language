package transformer_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	language "github.com/openfga/language/pkg/go/transformer"
)

func TestTransformJSONProtoToDSL(t *testing.T) {
	t.Parallel()

	testCases, err := loadModuleTestCases()
	require.NoError(t, err)

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()
			if len(testCase.Modules) == 0 || testCase.Skip || len(testCase.ExpectedErrors) > 0 {
				t.Skip()
			}

			model, err := language.TransformModuleFilesToModel(testCase.Modules, "1.2")
			require.NoError(t, err)

			dsl, err := language.TransformJSONProtoToDSL(model, language.WithIncludeSourceInformation(false))
			require.NoError(t, err)
			require.Equal(t, testCase.DSL, dsl)

			dslWithSrc, err := language.TransformJSONProtoToDSL(model, language.WithIncludeSourceInformation(true))
			require.NoError(t, err)
			require.Equal(t, testCase.DSLWithSourceInfo, dslWithSrc)
		})

	}
}
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

func TestGetModulesFromJSON(t *testing.T) {
	t.Parallel()

	testCases, err := loadModuleTestCases()
	require.NoError(t, err)

	for _, testCase := range testCases {
		if testCase.ExpectedModules == nil {
			continue
		}

		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			if testCase.Skip {
				t.Skip()
			}

			model, err := language.LoadJSONStringToProto(testCase.JSON)
			require.NoError(t, err)

			modules, err := language.GetModulesFromJSON(model)
			require.NoError(t, err)
			require.Equal(t, testCase.ExpectedModules, modules)
		})
	}
}

func TestGetModulesFromJSONForEdgeCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name            string
		json            string
		expectedModules []string
		expectedError   string
	}{
		{
			name:            "schema 1.2 model without modules",
			json:            `{"schema_version":"1.2","type_definitions":[{"type":"user"}]}`,
			expectedModules: []string{},
		},
		{
			name: "modules declared on types, relations and conditions",
			json: `{
  "schema_version": "1.2",
  "type_definitions": [
    {"type": "user", "metadata": {"module": "core"}},
    {
      "type": "folder",
      "relations": {"viewer": {"this": {}}},
      "metadata": {
        "module": "core",
        "relations": {
          "viewer": {"directly_related_user_types": [{"type": "user", "condition": "is_open"}], "module": "drive"}
        }
      }
    }
  ],
  "conditions": {
    "is_open": {"name": "is_open", "expression": "open", "parameters": {"open": {"type_name": "TYPE_NAME_BOOL"}}, "metadata": {"module": "access"}}
  }
}`,
			expectedModules: []string{"access", "core", "drive"},
		},
		{
			name:          "schema 1.1 model",
			json:          `{"schema_version":"1.1","type_definitions":[{"type":"user"}]}`,
			expectedError: "model schema version 1.1 does not support modules",
		},
		{
			name:          "model without a schema version",
			json:          `{"type_definitions":[{"type":"user"}]}`,
			expectedError: "model schema version 1.1 does not support modules",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			model, err := language.LoadJSONStringToProto(testCase.json)
			require.NoError(t, err)

			modules, err := language.GetModulesFromJSON(model)
			if testCase.expectedError != "" {
				require.EqualError(t, err, testCase.expectedError)
				require.Nil(t, modules)

				return
			}

			require.NoError(t, err)
			require.Equal(t, testCase.expectedModules, modules)
		})
	}
}
