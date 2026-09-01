package validation

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/openfga/language/pkg/go/transformer"
)

// Outcomes of a corpus case. ERROR is a case whose DSL no longer parses, which is
// distinct from FAIL: nothing about validation was tested.
const (
	corpusPass    = "PASS"
	corpusFail    = "FAIL"
	corpusSkipped = "SKIPPED"
	corpusError   = "ERROR"
)

// findingsOf recovers the findings behind a validation error; nil in, none out.
func findingsOf(err error) Findings {
	var findings Findings
	errors.As(err, &findings)

	return findings
}

// YAMLTestCase is one case from the shared validation corpus under tests/data.
type YAMLTestCase struct {
	Name           string              `yaml:"name"`
	DSL            string              `yaml:"dsl"`
	Skip           bool                `yaml:"skip,omitempty"`
	ExpectedErrors []YAMLExpectedError `yaml:"expected_errors,omitempty"`
}

// YAMLExpectedError is one error a corpus case expects.
//
// Line and Column are pointers so a case that states no position is told apart from
// one that states position 0. The corpus counts lines from zero, so the zero value is
// a real position and a value type would make the two indistinguishable.
type YAMLExpectedError struct {
	Message  string            `yaml:"msg"`
	Line     *YAMLRange        `yaml:"line,omitempty"`
	Column   *YAMLRange        `yaml:"column,omitempty"`
	Metadata YAMLErrorMetadata `yaml:"metadata,omitempty"`
}

// YAMLRange is a start and end position, used for both the line and the column.
type YAMLRange struct {
	Start int `yaml:"start"`
	End   int `yaml:"end"`
}

// YAMLErrorMetadata is the metadata a corpus case pins: the offending symbol and the
// error type.
type YAMLErrorMetadata struct {
	Symbol    string `yaml:"symbol,omitempty"`
	ErrorType string `yaml:"errorType,omitempty"`
}

// YAMLTestSuite is a corpus file's cases.
type YAMLTestSuite struct {
	TestCases []YAMLTestCase
	FilePath  string
}

// YAMLTestRunner loads corpus files from testDataPath and runs their cases.
type YAMLTestRunner struct {
	testDataPath string
	suites       map[string]*YAMLTestSuite
}

func NewYAMLTestRunner(testDataPath string) *YAMLTestRunner {
	return &YAMLTestRunner{
		testDataPath: testDataPath,
		suites:       make(map[string]*YAMLTestSuite),
	}
}

// LoadTestSuite reads and parses a corpus file, caching it by name.
func (runner *YAMLTestRunner) LoadTestSuite(filename string) (*YAMLTestSuite, error) {
	if suite, exists := runner.suites[filename]; exists {
		return suite, nil
	}

	filePath := filepath.Join(runner.testDataPath, filename)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read YAML file %s: %w", filePath, err)
	}

	var testCases []YAMLTestCase
	if err := yaml.Unmarshal(data, &testCases); err != nil {
		return nil, fmt.Errorf("failed to parse YAML file %s: %w", filePath, err)
	}

	suite := &YAMLTestSuite{TestCases: testCases, FilePath: filePath}
	runner.suites[filename] = suite

	return suite, nil
}

// YAMLTestResult is the outcome of one corpus case.
type YAMLTestResult struct {
	Status string

	// Problems lists every divergence from the case, one per line, and is empty for
	// a case that passed.
	Problems []string
}

// RunTestCase validates a case's DSL and compares the findings with what the case
// expects.
func (runner *YAMLTestRunner) RunTestCase(testCase YAMLTestCase) *YAMLTestResult {
	if testCase.Skip {
		return &YAMLTestResult{Status: corpusSkipped}
	}

	model, err := transformer.TransformDSLToProto(testCase.DSL)
	if err != nil {
		return &YAMLTestResult{
			Status:   corpusError,
			Problems: []string{fmt.Sprintf("DSL does not parse: %v", err)},
		}
	}

	return compareWithCorpus(testCase.ExpectedErrors, findingsOf(ValidateDSL(model, testCase.DSL)))
}

// compareWithCorpus pairs each expected error with a distinct finding, so a case
// expecting two errors is not satisfied by one finding that matches both.
func compareWithCorpus(expectedErrors []YAMLExpectedError, findings Findings) *YAMLTestResult {
	result := &YAMLTestResult{}
	claimed := make([]bool, len(findings))
	matched := make([]bool, len(expectedErrors))

	// Whole matches are paired first. Pairing on the message alone up front would let
	// one expectation take the finding that another one matches outright.
	for i, expected := range expectedErrors {
		for j, finding := range findings {
			if !claimed[j] && describeMismatch(expected, finding) == "" {
				claimed[j], matched[i] = true, true

				break
			}
		}
	}

	// What is left pairs on the message alone, so a finding that resolved to the wrong
	// line is reported as that one field rather than as an expectation with nothing
	// behind it plus an unexplained extra finding.
	for i, expected := range expectedErrors {
		if matched[i] {
			continue
		}

		paired := false

		for j, finding := range findings {
			if !claimed[j] && finding.Message == expected.Message {
				claimed[j], paired = true, true
				result.Problems = append(result.Problems,
					fmt.Sprintf("expected %s: %s", describeExpected(expected), describeMismatch(expected, finding)))

				break
			}
		}

		if !paired {
			result.Problems = append(result.Problems,
				fmt.Sprintf("no finding matches expected %s", describeExpected(expected)))
		}
	}

	for j, finding := range findings {
		if !claimed[j] {
			result.Problems = append(result.Problems, fmt.Sprintf("unexpected finding %s", describeFinding(finding)))
		}
	}

	if len(result.Problems) > 0 {
		result.Status = corpusFail
	} else {
		result.Status = corpusPass
	}

	return result
}

// describeMismatch returns what keeps finding from satisfying expected, or "" when
// nothing does.
//
// The message has to be equal rather than merely contain the expected text: the corpus
// is the contract between the implementations, so a message that only starts with the
// reference's is a divergence, not a pass.
func describeMismatch(expected YAMLExpectedError, finding *Finding) string {
	if finding.Message != expected.Message {
		return fmt.Sprintf("message %q, want %q", finding.Message, expected.Message)
	}

	if expected.Metadata.ErrorType != "" &&
		string(finding.Metadata.Kind) != expected.Metadata.ErrorType {
		return fmt.Sprintf("errorType %q, want %q", finding.Metadata.Kind, expected.Metadata.ErrorType)
	}

	if expected.Metadata.Symbol != "" && finding.Metadata.Symbol != expected.Metadata.Symbol {
		return fmt.Sprintf("symbol %q, want %q", finding.Metadata.Symbol, expected.Metadata.Symbol)
	}

	// A position the corpus states has to be reached, both ends of it. A finding
	// with no position at all must not satisfy an expectation that states one:
	// resolving to nowhere in the source is a failure mode of the line searches,
	// not a match.
	if problem := describeRangeMismatch("line", expected.Line, finding.Line); problem != "" {
		return problem
	}

	return describeRangeMismatch("column", expected.Column, finding.Column)
}

func describeRangeMismatch(name string, expected *YAMLRange, actual *Range) string {
	if expected == nil {
		return ""
	}

	if actual == nil {
		return "no " + name
	}

	if actual.Start != expected.Start || actual.End != expected.End {
		return fmt.Sprintf("%s %d-%d, want %d-%d", name, actual.Start, actual.End, expected.Start, expected.End)
	}

	return ""
}

func describeExpected(expected YAMLExpectedError) string {
	return fmt.Sprintf("%q [%s]%s", expected.Message, expected.Metadata.ErrorType,
		describePosition(rangeOf(expected.Line), rangeOf(expected.Column)))
}

func describeFinding(finding *Finding) string {
	return fmt.Sprintf("%q [%s]%s", finding.Message, finding.Metadata.Kind,
		describePosition(finding.Line, finding.Column))
}

func describePosition(line, column *Range) string {
	described := ""
	if line != nil {
		described += fmt.Sprintf(" line %d-%d", line.Start, line.End)
	}

	if column != nil {
		described += fmt.Sprintf(" column %d-%d", column.Start, column.End)
	}

	return described
}

func rangeOf(yamlRange *YAMLRange) *Range {
	if yamlRange == nil {
		return nil
	}

	return &Range{Start: yamlRange.Start, End: yamlRange.End}
}
