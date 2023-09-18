package transformer_test

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
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
	Name         string `json:"name" yaml:"name"`
	DSL          string `json:"dsl" yaml:"dsl"`
	Valid        bool   `json:"valid" yaml:"valid"`
	ErrorMessage string `json:"error_message" yaml:"error_message"`
}

func loadInvalidDslSyntaxTestCases() ([]invalidDslSyntaxTestCase, error) {
	data, err := os.ReadFile(filepath.Join("../../../tests", "data", "dsl-syntax-validation-cases.json"))
	if err != nil {
		return nil, err
	}

	testCases := []invalidDslSyntaxTestCase{}
	err = yaml.Unmarshal(data, &testCases)

	return testCases, err
}

type invalidJsonSyntaxTestCase struct {
	Name         string `json:"name" yaml:"name"`
	JSON         string `json:"json" yaml:"json"`
	ErrorMessage string `json:"error_message" yaml:"error_message"`
	Skip         bool   `json:"skip" yaml:"skip"`
}

func loadInvalidJsonSyntaxTestCases() ([]invalidJsonSyntaxTestCase, error) {
	data, err := os.ReadFile(filepath.Join("../../../tests", "data", "json-syntax-transformer-validation-cases.yaml"))
	if err != nil {
		return nil, err
	}

	testCases := []invalidJsonSyntaxTestCase{}
	err = yaml.Unmarshal(data, &testCases)

	return testCases, err
}
