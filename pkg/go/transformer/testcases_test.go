package transformer_test

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type ValidTestCase struct {
	Name string
	DSL  string
	JSON string
	Skip bool
}

func LoadValidTransformerTestCases() ([]ValidTestCase, error) {
	testDataPath := filepath.Join("../../../tests", "data", "transformer")
	entries, err := os.ReadDir(testDataPath)
	if err != nil {
		return nil, err
	}

	testCases := []ValidTestCase{}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		testCase := ValidTestCase{Name: e.Name()}

		skipFile, _ := os.ReadFile(filepath.Join(testDataPath, testCase.Name, "test.skip"))
		if skipFile != nil {
			testCase.Skip = true
		}
		jsonData, err := os.ReadFile(filepath.Join(testDataPath, testCase.Name, "authorization-model.json"))
		if err != nil {
			return nil, err
		}
		testCase.JSON = string(jsonData)

		dslData, err := os.ReadFile(filepath.Join(testDataPath, testCase.Name, "authorization-model.fga"))
		if err != nil {
			return nil, err
		}

		testCase.DSL = string(dslData)
		testCases = append(testCases, testCase)
	}

	return testCases, nil
}

type InvalidDslSyntaxTestCase struct {
	Name         string `json:"name"`
	DSL          string `json:"dsl"`
	Valid        bool   `json:"valid"`
	ErrorMessage string `json:"error_message"`
}

func LoadInvalidDslSyntaxTestCases() ([]InvalidDslSyntaxTestCase, error) {
	data, err := os.ReadFile(filepath.Join("../../../tests", "data", "dsl-syntax-validation.json"))
	if err != nil {
		return nil, err
	}

	testCases := []InvalidDslSyntaxTestCase{}
	err = json.Unmarshal(data, &testCases)

	return testCases, err
}
