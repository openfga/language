package transformer_test

import (
	"errors"
	"testing"

	pb "github.com/openfga/api/proto/openfga/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/openfga/language/pkg/go/transformer"
)

func TestTransformModuleToJSON(t *testing.T) {
	t.Parallel()

	testCases, err := loadModuleTestCases()
	if err != nil {
		t.Fatal(err)
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			if testCase.Skip {
				t.Skip()
			}

			actual, err := transformer.TransformModuleFilesToModel(testCase.Modules)
			if len(testCase.ExpectedErrors) == 0 {
				require.NoError(t, err)

				expectedAuthModel := &pb.AuthorizationModel{}
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

					for i := 0; i < len(testCase.ExpectedErrors); i++ {
						errorDetails := testCase.ExpectedErrors[i]
						expected := errors[i]
						actual := &transformer.ModuleTransformationSingleError{
							Line:   errorDetails.Line.Start,
							Column: errorDetails.Column.Start,
							Msg:    errorDetails.Msg,
						}

						assert.Equal(t, expected, actual)
					}
				}
			}
		})
	}
}
