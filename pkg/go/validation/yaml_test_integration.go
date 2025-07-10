package validation

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
	fgaSdk "github.com/openfga/go-sdk"
)

// YAMLTestCase represents a single test case from the YAML validation test files
type YAMLTestCase struct {
	Name           string                 `yaml:"name"`
	DSL            string                 `yaml:"dsl"`
	Skip           bool                   `yaml:"skip,omitempty"`
	ExpectedErrors []YAMLExpectedError    `yaml:"expected_errors,omitempty"`
	Metadata       map[string]interface{} `yaml:"metadata,omitempty"`
}

// YAMLExpectedError represents an expected validation error from YAML test files
type YAMLExpectedError struct {
	Message  string               `yaml:"msg"`
	Line     YAMLLineRange        `yaml:"line,omitempty"`
	Column   YAMLColumnRange      `yaml:"column,omitempty"`
	Metadata YAMLErrorMetadata    `yaml:"metadata,omitempty"`
}

// YAMLLineRange represents line start and end positions
type YAMLLineRange struct {
	Start int `yaml:"start"`
	End   int `yaml:"end"`
}

// YAMLColumnRange represents column start and end positions
type YAMLColumnRange struct {
	Start int `yaml:"start"`
	End   int `yaml:"end"`
}

// YAMLErrorMetadata represents error metadata from YAML test files
type YAMLErrorMetadata struct {
	Symbol    string `yaml:"symbol,omitempty"`
	ErrorType string `yaml:"errorType,omitempty"`
}

// YAMLTestSuite represents a collection of YAML test cases
type YAMLTestSuite struct {
	TestCases []YAMLTestCase
	FilePath  string
}

// YAMLTestRunner handles running YAML-based validation tests
type YAMLTestRunner struct {
	testDataPath string
	suites       map[string]*YAMLTestSuite
}

// NewYAMLTestRunner creates a new YAML test runner
func NewYAMLTestRunner(testDataPath string) *YAMLTestRunner {
	return &YAMLTestRunner{
		testDataPath: testDataPath,
		suites:       make(map[string]*YAMLTestSuite),
	}
}

// LoadTestSuite loads a YAML test suite from file
func (runner *YAMLTestRunner) LoadTestSuite(filename string) (*YAMLTestSuite, error) {
	filePath := filepath.Join(runner.testDataPath, filename)
	
	// Check if already loaded
	if suite, exists := runner.suites[filename]; exists {
		return suite, nil
	}
	
	// Read YAML file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read YAML file %s: %w", filePath, err)
	}
	
	// Parse YAML
	var testCases []YAMLTestCase
	if err := yaml.Unmarshal(data, &testCases); err != nil {
		return nil, fmt.Errorf("failed to parse YAML file %s: %w", filePath, err)
	}
	
	suite := &YAMLTestSuite{
		TestCases: testCases,
		FilePath:  filePath,
	}
	
	runner.suites[filename] = suite
	return suite, nil
}

// GetAvailableTestSuites returns all available YAML test suite files
func (runner *YAMLTestRunner) GetAvailableTestSuites() ([]string, error) {
	files, err := os.ReadDir(runner.testDataPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read test data directory: %w", err)
	}
	
	var yamlFiles []string
	for _, file := range files {
		if !file.IsDir() && (strings.HasSuffix(file.Name(), ".yaml") || strings.HasSuffix(file.Name(), ".yml")) {
			yamlFiles = append(yamlFiles, file.Name())
		}
	}
	
	return yamlFiles, nil
}

// RunTestCase runs a single YAML test case and compares results
func (runner *YAMLTestRunner) RunTestCase(testCase YAMLTestCase) (*YAMLTestResult, error) {
	if testCase.Skip {
		return &YAMLTestResult{
			TestCase: testCase,
			Status:   "SKIPPED",
			Message:  "Test case marked as skip in YAML",
		}, nil
	}
	
	// Parse DSL to create authorization model
	// Note: This would typically use a DSL parser, but for now we'll create a basic model
	model, err := runner.parseDSLToModel(testCase.DSL)
	if err != nil {
		return &YAMLTestResult{
			TestCase: testCase,
			Status:   "ERROR",
			Message:  fmt.Sprintf("Failed to parse DSL: %v", err),
		}, nil
	}
	
	// Run validation
	validationErrors := ValidateDSL(model, testCase.DSL, DefaultEngineOptions())
	
	// Compare results
	result := runner.compareResults(testCase, validationErrors)
	return result, nil
}

// YAMLTestResult represents the result of running a YAML test case
type YAMLTestResult struct {
	TestCase       YAMLTestCase
	Status         string   // "PASS", "FAIL", "SKIPPED", "ERROR"
	Message        string
	ActualErrors   []*ValidationError
	ExpectedErrors []YAMLExpectedError
	ErrorDetails   []string
}

// compareResults compares actual validation results with expected results from YAML
func (runner *YAMLTestRunner) compareResults(testCase YAMLTestCase, validationErrors *ValidationErrors) *YAMLTestResult {
	result := &YAMLTestResult{
		TestCase:       testCase,
		ActualErrors:   validationErrors.GetErrors(),
		ExpectedErrors: testCase.ExpectedErrors,
	}
	
	actualCount := validationErrors.Count()
	expectedCount := len(testCase.ExpectedErrors)
	
	// Check error count match
	if actualCount != expectedCount {
		result.Status = "FAIL"
		result.Message = fmt.Sprintf("Error count mismatch: expected %d, got %d", expectedCount, actualCount)
		result.ErrorDetails = append(result.ErrorDetails, result.Message)
	}
	
	// If no errors expected and none found, test passes
	if expectedCount == 0 && actualCount == 0 {
		result.Status = "PASS"
		result.Message = "No errors expected and none found"
		return result
	}
	
	// Compare individual errors
	errorMatches := make([]bool, len(testCase.ExpectedErrors))
	for i, expectedError := range testCase.ExpectedErrors {
		matched := false
		for _, actualError := range validationErrors.GetErrors() {
			if runner.errorsMatch(expectedError, actualError) {
				matched = true
				break
			}
		}
		errorMatches[i] = matched
		if !matched {
			detail := fmt.Sprintf("Expected error not found: %s", expectedError.Message)
			result.ErrorDetails = append(result.ErrorDetails, detail)
		}
	}
	
	// Check for unexpected errors
	for _, actualError := range validationErrors.GetErrors() {
		matched := false
		for _, expectedError := range testCase.ExpectedErrors {
			if runner.errorsMatch(expectedError, actualError) {
				matched = true
				break
			}
		}
		if !matched {
			detail := fmt.Sprintf("Unexpected error found: %s", actualError.Message)
			result.ErrorDetails = append(result.ErrorDetails, detail)
		}
	}
	
	// Determine overall status
	if len(result.ErrorDetails) == 0 {
		result.Status = "PASS"
		result.Message = fmt.Sprintf("All %d errors matched correctly", expectedCount)
	} else {
		result.Status = "FAIL"
		result.Message = fmt.Sprintf("Found %d error mismatches", len(result.ErrorDetails))
	}
	
	return result
}

// errorsMatch checks if an expected error matches an actual validation error
func (runner *YAMLTestRunner) errorsMatch(expected YAMLExpectedError, actual *ValidationError) bool {
	// Check message content (allow partial matches for flexibility)
	if !strings.Contains(strings.ToLower(actual.Message), strings.ToLower(expected.Message)) {
		return false
	}
	
	// Check error type if specified
	if expected.Metadata.ErrorType != "" {
		expectedType := ValidationErrorType(expected.Metadata.ErrorType)
		if actual.Metadata.ErrorType != expectedType {
			return false
		}
	}
	
	// Check line numbers if specified
	if expected.Line.Start > 0 && actual.Line != nil {
		if actual.Line.Start != expected.Line.Start {
			return false
		}
	}
	
	// Check column numbers if specified
	if expected.Column.Start > 0 && actual.Column != nil {
		if actual.Column.Start != expected.Column.Start {
			return false
		}
	}
	
	return true
}

// parseDSLToModel converts DSL content to an AuthorizationModel
// This is a simplified implementation - in practice, this would use a proper DSL parser
func (runner *YAMLTestRunner) parseDSLToModel(dsl string) (*fgaSdk.AuthorizationModel, error) {
	// For now, create a basic model that can be used for testing
	// This would be replaced with actual DSL parsing logic
	
	lines := strings.Split(dsl, "\n")
	model := &fgaSdk.AuthorizationModel{
		SchemaVersion:   "1.1", // Default version
		TypeDefinitions: []fgaSdk.TypeDefinition{},
	}
	
	// Extract schema version if present
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "schema ") {
			version := strings.TrimSpace(strings.TrimPrefix(line, "schema "))
			model.SchemaVersion = version
			break
		}
	}
	
	// Basic type extraction (simplified)
	currentType := ""
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "type ") {
			currentType = strings.TrimSpace(strings.TrimPrefix(line, "type "))
			if currentType != "" {
				typeDef := fgaSdk.TypeDefinition{
					Type: currentType,
				}
				model.TypeDefinitions = append(model.TypeDefinitions, typeDef)
			}
		}
	}
	
	return model, nil
}

// RunAllTestSuites runs all available YAML test suites
func (runner *YAMLTestRunner) RunAllTestSuites() (map[string][]*YAMLTestResult, error) {
	suiteFiles, err := runner.GetAvailableTestSuites()
	if err != nil {
		return nil, err
	}
	
	results := make(map[string][]*YAMLTestResult)
	
	for _, suiteFile := range suiteFiles {
		suite, err := runner.LoadTestSuite(suiteFile)
		if err != nil {
			return nil, fmt.Errorf("failed to load test suite %s: %w", suiteFile, err)
		}
		
		var suiteResults []*YAMLTestResult
		for _, testCase := range suite.TestCases {
			result, err := runner.RunTestCase(testCase)
			if err != nil {
				result = &YAMLTestResult{
					TestCase: testCase,
					Status:   "ERROR",
					Message:  err.Error(),
				}
			}
			suiteResults = append(suiteResults, result)
		}
		
		results[suiteFile] = suiteResults
	}
	
	return results, nil
}

// GenerateTestReport generates a comprehensive test report
func (runner *YAMLTestRunner) GenerateTestReport(results map[string][]*YAMLTestResult) *YAMLTestReport {
	report := &YAMLTestReport{
		SuiteResults: results,
		Summary:      make(map[string]int),
	}
	
	totalTests := 0
	for _, suiteResults := range results {
		for _, result := range suiteResults {
			totalTests++
			report.Summary[result.Status]++
		}
	}
	
	report.Summary["TOTAL"] = totalTests
	
	// Calculate pass rate
	if totalTests > 0 {
		passCount := report.Summary["PASS"]
		report.PassRate = float64(passCount) / float64(totalTests) * 100.0
	}
	
	return report
}

// YAMLTestReport represents a comprehensive test report
type YAMLTestReport struct {
	SuiteResults map[string][]*YAMLTestResult
	Summary      map[string]int
	PassRate     float64
}

// PrintReport prints a formatted test report
func (report *YAMLTestReport) PrintReport() {
	fmt.Printf("📊 YAML Test Integration Report\n")
	fmt.Printf("================================\n\n")
	
	fmt.Printf("📈 Summary:\n")
	fmt.Printf("  Total Tests: %d\n", report.Summary["TOTAL"])
	fmt.Printf("  Passed: %d\n", report.Summary["PASS"])
	fmt.Printf("  Failed: %d\n", report.Summary["FAIL"])
	fmt.Printf("  Skipped: %d\n", report.Summary["SKIPPED"])
	fmt.Printf("  Errors: %d\n", report.Summary["ERROR"])
	fmt.Printf("  Pass Rate: %.1f%%\n\n", report.PassRate)
	
	// Print detailed results for each suite
	for suiteName, results := range report.SuiteResults {
		fmt.Printf("📋 Test Suite: %s\n", suiteName)
		fmt.Printf("  Tests: %d\n", len(results))
		
		passed := 0
		failed := 0
		skipped := 0
		errors := 0
		
		for _, result := range results {
			switch result.Status {
			case "PASS":
				passed++
			case "FAIL":
				failed++
			case "SKIPPED":
				skipped++
			case "ERROR":
				errors++
			}
		}
		
		fmt.Printf("  Passed: %d, Failed: %d, Skipped: %d, Errors: %d\n", passed, failed, skipped, errors)
		
		// Show failed tests
		if failed > 0 || errors > 0 {
			fmt.Printf("  ❌ Failed/Error Tests:\n")
			for _, result := range results {
				if result.Status == "FAIL" || result.Status == "ERROR" {
					fmt.Printf("    - %s: %s\n", result.TestCase.Name, result.Message)
					for _, detail := range result.ErrorDetails {
						fmt.Printf("      • %s\n", detail)
					}
				}
			}
		}
		
		fmt.Printf("\n")
	}
}
