package transformer_test

import (
	"os"
	"path/filepath"
)

type TestCases struct {
	Cases []TestCase
}

type TestCase struct {
	Name string
	DSL  string
	JSON string
	Skip bool
}

func LoadTransformerTestCases() ([]TestCase, error) {
	testDataPath := filepath.Join("../../../tests", "data", "transformer")
	entries, err := os.ReadDir(testDataPath)
	if err != nil {
		return nil, err
	}

	testCases := []TestCase{}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		testCase := TestCase{Name: e.Name()}

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
