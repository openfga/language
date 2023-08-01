package transformer_test

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type validTestCase struct {
	Name string
	DSL  string
	JSON string
	Skip bool
}

func loadValidTransformerTestCases() ([]validTestCase, error) {
	testDataPath := filepath.Join("../../../tests", "data", "transformer")
	entries, err := os.ReadDir(testDataPath)
	if err != nil {
		return nil, err
	}

	testCases := []validTestCase{}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		testCase := validTestCase{Name: e.Name()}

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

type invalidDslSyntaxTestCase struct {
	Name         string `json:"name"`
	DSL          string `json:"dsl"`
	Valid        bool   `json:"valid"`
	ErrorMessage string `json:"error_message"`
}

func loadInvalidDslSyntaxTestCases() ([]invalidDslSyntaxTestCase, error) {
	data, err := os.ReadFile(filepath.Join("../../../tests", "data", "dsl-syntax-validation.json"))
	if err != nil {
		return nil, err
	}

	testCases := []invalidDslSyntaxTestCase{}
	err = json.Unmarshal(data, &testCases)

	return testCases, err
}
