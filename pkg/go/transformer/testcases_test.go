package transformer_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type validTestCase struct {
	Name string
	DSL  string
	JSON string
	Skip bool
}

type startEnd struct {
	Start int `json:"start,omitempty" yaml:"start,omitempty"`
	End   int `json:"end,omitempty" yaml:"end,omitempty"`
}

type expectedError struct {
	Msg    string   `json:"msg" yaml:"msg"`
	Line   startEnd `json:"line" yaml:"line"`
	Column startEnd `json:"column" yaml:"column"`
}

func (testCase *invalidDslSyntaxTestCase) GetErrorString() string {
	pluralS := "s"
	if len(testCase.ExpectedErrors) == 1 {
		pluralS = ""
	}

	errorLines := []string{}
	for _, e := range testCase.ExpectedErrors {
		errorLines = append(errorLines, fmt.Sprintf(`syntax error at line=%d, column=%d: %s`, e.Line.Start, e.Column.Start, e.Msg))
	}

	return fmt.Sprintf("%d error%s occurred:\n\t* %v\n\n", len(testCase.ExpectedErrors), pluralS, strings.Join(errorLines, "\n\t* "))
}

type invalidDslSyntaxTestCase struct {
	Name           string          `json:"name" yaml:"name"`
	DSL            string          `json:"dsl" yaml:"dsl"`
	Valid          bool            `json:"valid" yaml:"valid"`
	ExpectedErrors []expectedError `json:"expected_errors" yaml:"expected_errors"`
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

func loadInvalidDslSyntaxTestCases() ([]invalidDslSyntaxTestCase, error) {
	data, err := os.ReadFile(filepath.Join("../../../tests", "data", "dsl-syntax-validation-cases.yaml"))
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
