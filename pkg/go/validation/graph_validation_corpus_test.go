package validation

import (
	"fmt"
	"sort"
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/openfga/language/pkg/go/graph"
	"github.com/openfga/language/pkg/go/transformer"
)

// corpusModels returns the non-skipped cases of the shared semantic corpus with their
// models already parsed.
//
// It fails on a case that does not parse rather than dropping it. The corpus is the
// contract between the implementations, and a case silently excluded from a coverage
// count is worse than one that fails.
func corpusModels(t *testing.T) []struct {
	Case  YAMLTestCase
	Model *openfgav1.AuthorizationModel
} {
	t.Helper()

	suite, err := NewYAMLTestRunner(corpusDir).LoadTestSuite("dsl-semantic-validation-cases.yaml")
	require.NoError(t, err)
	require.NotEmpty(t, suite.TestCases, "corpus loaded no cases")

	var cases []struct {
		Case  YAMLTestCase
		Model *openfgav1.AuthorizationModel
	}

	for _, testCase := range suite.TestCases {
		if testCase.Skip {
			continue
		}

		model, err := transformer.TransformDSLToProto(testCase.DSL)
		require.NoErrorf(t, err, "corpus case %q does not parse", testCase.Name)

		cases = append(cases, struct {
			Case  YAMLTestCase
			Model *openfgav1.AuthorizationModel
		}{Case: testCase, Model: model})
	}

	return cases
}

// TestGraphAgreesWithTraversalOnWhichModelsAreInvalid pins how the two paths divide the
// corpus, and the claim that matters is graphOnly being zero: over every case available
// there is no model the graph rejects that validation calls valid. That is what makes
// the graph safe to put behind the same entry points.
//
// The other counts are here so a change to either path shows up as a number rather than
// as a corpus case quietly moving between buckets.
func TestGraphAgreesWithTraversalOnWhichModelsAreInvalid(t *testing.T) {
	t.Parallel()

	var bothFlag, validationOnly, graphOnly, bothClean int

	var graphOnlyNames []string

	cases := corpusModels(t)

	for _, entry := range cases {
		invalidByTraversal := findingsFrom(
			ValidateDSL(entry.Model, entry.Case.DSL, DefaultEngineOptions())).HasErrors()

		_, buildErr := graph.NewWeightedAuthorizationModelGraphBuilder().Build(entry.Model)
		refusedByGraph := buildErr != nil

		switch {
		case invalidByTraversal && refusedByGraph:
			bothFlag++
		case invalidByTraversal:
			validationOnly++
		case refusedByGraph:
			graphOnly++

			graphOnlyNames = append(graphOnlyNames, fmt.Sprintf("%s: %v", entry.Case.Name, buildErr))
		default:
			bothClean++
		}
	}

	assert.Emptyf(t, graphOnlyNames,
		"the graph refuses a model validation calls valid, so the two paths disagree on validity")

	assert.Len(t, cases, 89, "non-skipped corpus cases")
	assert.Equal(t, 38, bothFlag, "cases both paths reject")
	assert.Equal(t, 47, validationOnly, "cases validation rejects and the graph builds")
	assert.Equal(t, 0, graphOnly, "cases the graph rejects and validation accepts")
	assert.Equal(t, 4, bothClean, "cases both paths accept")
}

// TestCorpusUnderGraphValidation runs the whole corpus with the graph path selected and
// requires every case to pass, which is what makes the option safe to offer.
//
// Failures are named rather than counted. A count alone would stay green if one case
// started failing as another started passing, which is the shape a rule-parity regression
// takes.
func TestCorpusUnderGraphValidation(t *testing.T) {
	t.Parallel()

	var diverged []string

	for _, entry := range corpusModels(t) {
		result := compareWithCorpus(entry.Case.ExpectedErrors, findingsFrom(
			ValidateDSL(entry.Model, entry.Case.DSL, &EngineOptions{UseGraphValidation: true})))

		if result.Status != corpusPass {
			diverged = append(diverged, entry.Case.Name)
		}
	}

	sort.Strings(diverged)

	assert.Empty(t, diverged, "corpus cases the graph path does not satisfy")
}

// TestRefusedCorpusModelsReportWhatTheWalkReports covers the cases the graph is never read
// for, which is 38 of the 89.
//
// Weight assignment stops at the first node it cannot weight, and it marks a node visited
// before walking that node's edges, so a graph it refused holds relations left unweighted
// because the walk had not reached them yet. Nothing distinguishes those from the relations
// nothing can satisfy, so a refused build is answered by the rewrite-tree walk and the
// findings have to be identical to what the traversal path gives.
//
// The count is asserted so the comparison cannot pass by covering nothing.
func TestRefusedCorpusModelsReportWhatTheWalkReports(t *testing.T) {
	t.Parallel()

	var refused int

	for _, entry := range corpusModels(t) {
		if _, err := graph.NewWeightedAuthorizationModelGraphBuilder().Build(entry.Model); err == nil {
			continue
		}

		refused++

		fromWalk := findingsFrom(ValidateDSL(entry.Model, entry.Case.DSL, DefaultEngineOptions())).GetErrors()
		viaGraph := findingsFrom(
			ValidateDSL(entry.Model, entry.Case.DSL, &EngineOptions{UseGraphValidation: true})).GetErrors()

		assert.Equalf(t, describeFindings(fromWalk), describeFindings(viaGraph),
			"the graph path reports something other than the walk for refused case %q", entry.Case.Name)
	}

	assert.Equal(t, 38, refused, "corpus cases the builder refuses")
}

// TestNoCorpusModelIsRefusedWithNothingToSay pins the gap this path leaves.
//
// A model the builder refuses and the walk has no finding for is reported as the refusal
// itself, which names a pattern but no relation and carries no position. No corpus case is
// in that state; the models that are have unit tests instead. A case arriving here means
// the corpus grew a model whose only finding a consumer cannot put on a line.
func TestNoCorpusModelIsRefusedWithNothingToSay(t *testing.T) {
	t.Parallel()

	var positionless []string

	for _, entry := range corpusModels(t) {
		for _, finding := range findingsFrom(
			ValidateDSL(entry.Model, entry.Case.DSL, &EngineOptions{UseGraphValidation: true})).GetErrors() {
			if finding.Metadata != nil && finding.Metadata.ErrorType == GraphModelUnbuildable {
				positionless = append(positionless, fmt.Sprintf("%s: %s", entry.Case.Name, finding.Message))
			}
		}
	}

	sort.Strings(positionless)

	assert.Empty(t, positionless, "corpus cases reporting a refusal with no relation and no position")
}

// TestEntrypointRuleMatchesTraversalWhereTheGraphBuilds is the parity guarantee behind
// the switch: wherever there is a graph to read, reading it names the same unreachable
// relations as walking the rewrite tree, with the same message and the same position.
//
// It compares whole findings rather than counts, so a rule that found the right
// relations and resolved them to the wrong line fails here.
func TestEntrypointRuleMatchesTraversalWhereTheGraphBuilds(t *testing.T) {
	t.Parallel()

	var built, firedOn int

	for _, entry := range corpusModels(t) {
		if _, err := graph.NewWeightedAuthorizationModelGraphBuilder().Build(entry.Model); err != nil {
			continue
		}

		built++

		fromTraversal := entrypointFindings(entry.Model, entry.Case.DSL, false)
		fromGraph := entrypointFindings(entry.Model, entry.Case.DSL, true)

		if len(fromGraph) > 0 {
			firedOn++
		}

		assert.Equalf(t, fromTraversal, fromGraph,
			"the two paths report different entrypoint findings for %q", entry.Case.Name)
	}

	assert.Equal(t, 51, built, "corpus cases the graph builds")
	assert.Equal(t, 7, firedOn,
		"cases the graph rule reports an unreachable relation for; zero here would make the "+
			"comparison above a comparison of empty lists")
}

// entrypointFindings returns the entrypoint findings for a model as sorted strings
// carrying the symbol, the position and the message, which is everything the corpus
// compares.
func entrypointFindings(model *openfgav1.AuthorizationModel, dsl string, useGraph bool) []string {
	findings := findingsFrom(ValidateDSL(model, dsl, &EngineOptions{UseGraphValidation: useGraph})).GetErrors()

	described := []string{}

	for _, finding := range findings {
		if finding.Metadata == nil || finding.Metadata.ErrorType != RelationNoEntrypoint {
			continue
		}

		described = append(described, fmt.Sprintf("%s %s %q",
			finding.Metadata.Symbol,
			describePosition(finding.Line, finding.Column),
			finding.Message))
	}

	sort.Strings(described)

	return described
}
