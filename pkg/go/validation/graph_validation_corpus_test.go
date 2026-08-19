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

	assert.Len(t, cases, 84, "non-skipped corpus cases")
	assert.Equal(t, 33, bothFlag, "cases both paths reject")
	assert.Equal(t, 47, validationOnly, "cases validation rejects and the graph builds")
	assert.Equal(t, 0, graphOnly, "cases the graph rejects and validation accepts")
	assert.Equal(t, 4, bothClean, "cases both paths accept")
}

// graphDivergentCorpusCases are the corpus cases that do not pass under
// UseGraphValidation, and every one of them diverges the same way: the case expects one
// relation-no-entry-point per unreachable relation, and the graph path reports a single
// graph-model-unbuildable for the model.
//
// The cause is the builder, which stops at the first problem and returns no graph with
// its error, so there is nothing left to enumerate relations from. Twenty-three
// positioned findings across these eleven cases become eleven positionless ones.
//
// Every other case whose build is refused is stopped by the cascade gate before the
// entrypoint phase runs, so its refusal is never reported and the case still passes.
var graphDivergentCorpusCases = []string{
	"cycle 1 should fail",
	"cycle 2 should fail",
	"cyclic loop",
	"exclusion base not allow to reference itself in TTU",
	"intersection child not allow to reference itself in TTU",
	"no entry point exclusion that relates to itself",
	"no entry point intersection that relates to itself",
	"no_entrypoint_1 should fail",
	"no_entrypoint_2 should fail",
	"no_entrypoint_3a should fail",
	"no_entrypoint_3b should fail",
}

// TestCorpusUnderGraphValidation runs the whole corpus with the graph path selected and
// pins the divergence to exactly the cases above, by name.
//
// Naming them is the point. A count alone would stay green if one case started failing
// as another started passing, which is the shape a rule-parity regression takes.
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

	assert.Equal(t, graphDivergentCorpusCases, diverged,
		"the set of corpus cases the graph path does not satisfy has changed")
}

// TestGraphDivergenceIsOnlyTheUnbuildableSubstitution checks what the divergence above
// consists of, so the test names a behaviour rather than a list.
//
// For each divergent case: the case expects nothing but entrypoint findings, and the
// graph path reports exactly one finding, the unbuildable one. Anything else, a second
// finding or a different code, is a different problem than fail-fast Build.
func TestGraphDivergenceIsOnlyTheUnbuildableSubstitution(t *testing.T) {
	t.Parallel()

	divergent := make(map[string]struct{}, len(graphDivergentCorpusCases))
	for _, name := range graphDivergentCorpusCases {
		divergent[name] = struct{}{}
	}

	var expectedEntrypointFindings, checked int

	for _, entry := range corpusModels(t) {
		if _, ok := divergent[entry.Case.Name]; !ok {
			continue
		}

		checked++

		for _, expected := range entry.Case.ExpectedErrors {
			require.Equalf(t, string(RelationNoEntrypoint), expected.Metadata.ErrorType,
				"case %q expects a code other than the entrypoint one, so it is not this divergence",
				entry.Case.Name)

			expectedEntrypointFindings++
		}

		findings := findingsFrom(
			ValidateDSL(entry.Model, entry.Case.DSL, &EngineOptions{UseGraphValidation: true})).GetErrors()

		require.Lenf(t, findings, 1, "case %q reports more than the unbuildable finding", entry.Case.Name)
		require.NotNil(t, findings[0].Metadata)
		assert.Equalf(t, GraphModelUnbuildable, findings[0].Metadata.ErrorType,
			"case %q diverges by reporting something other than the unbuildable finding", entry.Case.Name)
	}

	assert.Len(t, graphDivergentCorpusCases, checked, "every named case was found in the corpus")
	assert.Equal(t, 23, expectedEntrypointFindings,
		"positioned entrypoint findings the graph path gives up, one per unreachable relation")
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
