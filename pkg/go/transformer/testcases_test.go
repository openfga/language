package transformer_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/openfga/language/pkg/go/transformer"
)

type validTestCase struct {
	Name string
	DSL  string
	JSON string
	Skip bool
}

type startEnd struct {
	Start int `json:"start,omitempty" yaml:"start,omitempty"`
	End   int `json:"end,omitempty"   yaml:"end,omitempty"`
}

type meta struct {
	ErrorType string `json:"errorType" yaml:"errorType"` //nolint:tagliatelle
}

type expectedError struct {
	Msg      string   `json:"msg"      yaml:"msg"`
	Line     startEnd `json:"line"     yaml:"line"`
	Column   startEnd `json:"column"   yaml:"column"`
	File     string   `json:"file"     yaml:"file"`
	Metadata meta     `json:"metadata" yaml:"metadata"`
	Type     string   `json:"type"     yaml:"type"`
}

func (testCase *invalidDslSyntaxTestCase) GetErrorString() string {
	pluralS := "s"
	if len(testCase.ExpectedErrors) == 1 {
		pluralS = ""
	}

	errorLines := []string{}
	for _, e := range testCase.ExpectedErrors {
		errorLines = append(
			errorLines,
			fmt.Sprintf(`syntax error at line=%d, column=%d: %s`, e.Line.Start, e.Column.Start, e.Msg),
		)
	}

	return fmt.Sprintf(
		"%d error%s occurred:\n\t* %v\n\n",
		len(testCase.ExpectedErrors),
		pluralS,
		strings.Join(errorLines, "\n\t* "),
	)
}

type invalidDslSyntaxTestCase struct {
	Name           string          `json:"name"            yaml:"name"`
	DSL            string          `json:"dsl"             yaml:"dsl"`
	Valid          bool            `json:"valid"           yaml:"valid"`
	ExpectedErrors []expectedError `json:"expected_errors" yaml:"expected_errors"`
}

func loadValidTransformerTestCases() ([]validTestCase, error) {
	testDataPath := filepath.Join("../../../tests", "data", "transformer")

	entries, err := os.ReadDir(testDataPath)
	if err != nil {
		return nil, err //nolint:wrapcheck
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
			return nil, err //nolint:wrapcheck
		}

		testCase.JSON = string(jsonData)

		dslData, err := os.ReadFile(filepath.Join(testDataPath, testCase.Name, "authorization-model.fga"))
		if err != nil {
			return nil, err //nolint:wrapcheck
		}

		testCase.DSL = string(dslData)
		testCases = append(testCases, testCase)
	}

	return testCases, nil
}

func loadInvalidDslSyntaxTestCases() ([]invalidDslSyntaxTestCase, error) {
	data, err := os.ReadFile(filepath.Join("../../../tests", "data", "dsl-syntax-validation-cases.yaml"))
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	testCases := []invalidDslSyntaxTestCase{}
	err = yaml.Unmarshal(data, &testCases)

	return testCases, err //nolint:wrapcheck
}

type invalidJSONSyntaxTestCase struct {
	Name         string `json:"name"          yaml:"name"`
	JSON         string `json:"json"          yaml:"json"`
	ErrorMessage string `json:"error_message" yaml:"error_message"`
	Skip         bool   `json:"skip"          yaml:"skip"`
}

func loadInvalidJSONSyntaxTestCases() ([]invalidJSONSyntaxTestCase, error) {
	data, err := os.ReadFile(filepath.Join("../../../tests", "data", "json-syntax-transformer-validation-cases.yaml"))
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	testCases := []invalidJSONSyntaxTestCase{}
	err = yaml.Unmarshal(data, &testCases)

	return testCases, err //nolint:wrapcheck
}

type fgdModFileTestCase struct {
	Name           string          `json:"name"            yaml:"name"`
	ModFile        string          `json:"modFile"         yaml:"modFile"` //nolint:tagliatelle
	JSON           string          `json:"json"            yaml:"json"`
	Skip           bool            `json:"skip"            yaml:"skip"`
	ExpectedErrors []expectedError `json:"expected_errors" yaml:"expected_errors"`
}

func (testCase *fgdModFileTestCase) GetErrorString() string {
	pluralS := ""
	if len(testCase.ExpectedErrors) > 1 {
		pluralS = "s"
	}

	errorsString := []string{}
	for _, err := range testCase.ExpectedErrors {
		errorsString = append(
			errorsString,
			fmt.Sprintf("validation error at line=%d, column=%d: %s", err.Line.Start, err.Column.Start, err.Msg),
		)
	}

	return fmt.Sprintf(
		"%d error%s occurred:\n\t* %s\n\n",
		len(testCase.ExpectedErrors),
		pluralS,
		strings.Join(errorsString, "\n\t* "),
	)
}

func loadModFileTestCases() ([]fgdModFileTestCase, error) {
	data, err := os.ReadFile(filepath.Join("../../../tests", "data", "fga-mod-transformer-cases.yaml"))
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	testCases := []fgdModFileTestCase{}
	err = yaml.Unmarshal(data, &testCases)

	return testCases, err //nolint:wrapcheck
}

type moduleTestCase struct {
	Name           string `json:"name" yaml:"name"`
	Modules        []transformer.ModuleFile
	JSON           string `json:"json" yaml:"json"`
	DSL            string
	Skip           bool
	ExpectedErrors []expectedError `json:"expected_errors" yaml:"expected_errors"`
}

func (testCase *moduleTestCase) GetErrorString() string {
	pluralS := ""
	if len(testCase.ExpectedErrors) > 1 {
		pluralS = "s"
	}

	errorsString := []string{}

	for _, err := range testCase.ExpectedErrors {
		errorType := "transformation"
		if err.Type != "" {
			errorType = err.Type
		}

		errorsString = append(
			errorsString,
			fmt.Sprintf("%s error at line=%d, column=%d: %s", errorType, err.Line.Start, err.Column.Start, err.Msg),
		)
	}

	return fmt.Sprintf(
		"%d error%s occurred:\n\t* %s\n\n",
		len(testCase.ExpectedErrors),
		pluralS,
		strings.Join(errorsString, "\n\t* "),
	)
}

func loadModuleTestCases() ([]moduleTestCase, error) { //nolint:cyclop
	testDataPath := filepath.Join("../../../tests", "data", "transformer-module")

	entries, err := os.ReadDir(testDataPath)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	testCases := []moduleTestCase{}

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}

		testCase := moduleTestCase{Name: e.Name()}

		skipFile, _ := os.ReadFile(filepath.Join(testDataPath, testCase.Name, "test.skip"))
		if skipFile != nil {
			testCase.Skip = true
		}

		modelFile := filepath.Join(testDataPath, testCase.Name, "authorization-model.json")
		if jsonData, err := os.ReadFile(modelFile); err == nil {
			testCase.JSON = string(jsonData)
		}

		dslFile := filepath.Join(testDataPath, testCase.Name, "combined.fga")
		if dslData, err := os.ReadFile(dslFile); err == nil {
			testCase.DSL = string(dslData)
		}

		errorsFile := filepath.Join(testDataPath, testCase.Name, "expected_errors.json")
		if errorsData, err := os.ReadFile(errorsFile); err == nil {
			err := json.Unmarshal(errorsData, &testCase.ExpectedErrors)
			if err != nil {
				return nil, err //nolint:wrapcheck
			}

			errors := []expectedError{}

			// Ignore any errors that are from validation performed on the constructed model
			for _, err := range testCase.ExpectedErrors {
				if err.Metadata.ErrorType == "" {
					errors = append(errors, err)
				}
			}

			testCase.ExpectedErrors = errors
		}

		moduleDirectory := filepath.Join(testDataPath, testCase.Name, "module")

		moduleFiles, err := os.ReadDir(moduleDirectory)
		if err != nil {
			return nil, err //nolint:wrapcheck
		}

		modules := []transformer.ModuleFile{}

		for _, file := range moduleFiles {
			if file.IsDir() || !strings.HasSuffix(file.Name(), ".fga") {
				continue
			}

			moduleFile, _ := os.ReadFile(filepath.Join(moduleDirectory, file.Name()))

			modules = append(modules, transformer.ModuleFile{
				Name:     file.Name(),
				Contents: string(moduleFile),
			})
		}

		testCase.Modules = modules
		testCases = append(testCases, testCase)
	}

	return testCases, nil
}
