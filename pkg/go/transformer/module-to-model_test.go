package transformer_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/openfga/language/pkg/go/transformer"
)

func TestTransformModuleToJSON(t *testing.T) {
	t.Parallel()

	testCases, err := loadModuleTestCases()
	if err != nil {
		t.Fatal(err)
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			if testCase.Skip {
				t.Skip()
			}

			if testCase.Modules == nil {
				return
			}

			actual, err := transformer.TransformModuleFilesToModel(testCase.Modules, "1.2")
			if len(testCase.ExpectedErrors) == 0 {
				require.NoError(t, err)

				expectedAuthModel := &openfgav1.AuthorizationModel{}
				err = protojson.Unmarshal([]byte(testCase.JSON), expectedAuthModel)
				require.NoError(t, err)

				expectedJSONBytes, err := protojson.Marshal(expectedAuthModel)
				require.NoError(t, err)

				actualJSONBytes, err := protojson.Marshal(actual)
				require.NoError(t, err)

				require.JSONEq(t, string(expectedJSONBytes), string(actualJSONBytes))
			} else {
				require.Error(t, err)

				errorMessage := testCase.GetErrorString()
				assert.Equal(t, errorMessage, err.Error())

				var verr *transformer.ModuleValidationMultipleError
				if errors.As(err, &verr) {
					errors := verr.Errors

					require.Equal(t, len(errors), len(testCase.ExpectedErrors), "did not get the expected amount of errors")

					for i := 0; i < len(testCase.ExpectedErrors); i++ {
						errorDetails := testCase.ExpectedErrors[i]
						actual := errors[i]

						errorType := "transformation"
						if errorDetails.Type != "" {
							errorType = errorDetails.Type
						}

						assert.Equal(t,
							fmt.Sprintf("%s error at line=%d, column=%d: %s",
								errorType,
								errorDetails.Line.Start,
								errorDetails.Column.Start,
								errorDetails.Msg,
							),
							actual.Error(),
						)
					}
				}
			}
		})
	}
}

func TestSchemaVersion(t *testing.T) {
	t.Parallel()

	actual, err := transformer.TransformModuleFilesToModel([]transformer.ModuleFile{
		{
			Name: "core.fga",
			Contents: `module core
  type user`,
		},
	}, "1.1")

	require.NoError(t, err)
	assert.Equal(t, "1.1", actual.GetSchemaVersion())
}
