package validation

import (
	"fmt"
	"slices"
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/openfga/language/pkg/go/graph"
	"github.com/openfga/language/pkg/go/transformer"
)

// sharedCaseModels returns the non-skipped cases of the shared semantic
// validation file with their models already parsed.
//
// It fails on a case that does not parse rather than dropping it. The shared file is
// the contract between the implementations, and a case silently excluded from a
// coverage count is worse than one that fails.
func sharedCaseModels(t *testing.T) []struct {
	Case  YAMLTestCase
	Model *openfgav1.AuthorizationModel
} {
	t.Helper()

	suite, err := NewYAMLTestRunner(sharedDataDir).LoadTestSuite("dsl-semantic-validation-cases.yaml")
	require.NoError(t, err)
	require.NotEmpty(t, suite.TestCases, "shared test file loaded no cases")

	var cases []struct {
		Case  YAMLTestCase
		Model *openfgav1.AuthorizationModel
	}

	for _, testCase := range suite.TestCases {
		if testCase.Skip {
			continue
		}

		model, err := transformer.TransformDSLToProto(testCase.DSL)
		require.NoErrorf(t, err, "shared case %q does not parse", testCase.Name)

		cases = append(cases, struct {
			Case  YAMLTestCase
			Model *openfgav1.AuthorizationModel
		}{Case: testCase, Model: model})
	}

	return cases
}

// TestGraphAgreesWithTraversalOnWhichModelsAreInvalid pins how the two paths
// divide the shared cases, and the claim that matters is graphOnly being zero: over
// every case available there is no model the graph rejects that validation calls
// valid. That is what makes the graph safe to put behind the same entry points.
//
// The other counts are here so a change to either path shows up as a number
// rather than as a case quietly moving between buckets.
func TestGraphAgreesWithTraversalOnWhichModelsAreInvalid(t *testing.T) {
	t.Parallel()

	var bothFlag, validationOnly, graphOnly, bothClean int

	var graphOnlyNames []string

	cases := sharedCaseModels(t)

	for _, entry := range cases {
		invalidByTraversal := len(ExtractAllAs[*Finding](ValidateDSL(entry.Model, entry.Case.DSL))) > 0

		_, buildErr := graph.NewWeightedAuthorizationModelGraphBuilder().Build(entry.Model)
		refusedByGraph := buildErr != nil

		switch {
		case invalidByTraversal && refusedByGraph:
			bothFlag++
		case invalidByTraversal:
			validationOnly++
		case refusedByGraph:
			graphOnly++

			graphOnlyNames = append(graphOnlyNames, entry.Case.Name)
		default:
			bothClean++
		}
	}

	// The one disagreement is the builder's stricter tupleset rule, covered by
	// graphDivergentSharedCases and reported through the unbuildable backstop.
	assert.Equal(t, graphDivergentSharedCases, graphOnlyNames,
		"the set of models the graph refuses and validation accepts has changed")

	assert.Len(t, cases, 102, "non-skipped shared cases")
	assert.Equal(t, 39, bothFlag, "cases both paths reject")
	assert.Equal(t, 55, validationOnly, "cases validation rejects and the graph builds")
	assert.Equal(t, 1, graphOnly, "cases the graph rejects and validation accepts")
	assert.Equal(t, 7, bothClean, "cases both paths accept")
}

// graphDivergentSharedCases are the shared cases that do not pass under the
// graph entry-point phase.
//
// The one divergence left is a builder gap: the weighted graph requires a
// tupleset target relation on every assignable type of the tupleset, while
// validation, pkg/js and pkg/java require it on one. The builder refuses the
// model, the walk and the cycle shapes find nothing to attribute the refusal
// to, and the case's empty expectation meets one graph-model-unbuildable
// finding.
//
// A refusal the rewrite-tree walk can attribute no longer diverges at all:
// graphRefusal reports the walk's own findings, so the cycle and
// no-entrypoint cases pass identically under both phases.
var graphDivergentSharedCases = []string{
	"from target relation needs to exist on only one of the assignable types",
}

// TestSharedCasesUnderGraphValidation runs every shared case with the graph entry-point
// phase selected and pins the divergence to exactly the cases above, by name.
//
// Naming them is the point. A count alone would stay green if one case started
// failing as another started passing, which is the shape a rule-parity regression
// takes.
func TestSharedCasesUnderGraphValidation(t *testing.T) {
	t.Parallel()

	var diverged []string

	for _, entry := range sharedCaseModels(t) {
		result := compareWithExpected(entry.Case.ExpectedErrors,
			ExtractAllAs[*Finding](ValidateDSLWithGraph(entry.Model, entry.Case.DSL)))

		if result.Status != casePass {
			diverged = append(diverged, entry.Case.Name)
		}
	}

	slices.Sort(diverged)

	assert.Equal(t, graphDivergentSharedCases, diverged,
		"the set of shared cases the graph path does not satisfy has changed")
}

// TestGraphDivergenceIsOnlyTheBuilderRefusal checks what the divergence above
// consists of, so the test names a behaviour rather than a list.
//
// For each divergent case: validation calls the model valid, and the graph path
// reports exactly one finding, the positionless unbuildable one. Anything else,
// a second finding or a different code, is a different problem than the
// builder refusing a model validation accepts.
func TestGraphDivergenceIsOnlyTheBuilderRefusal(t *testing.T) {
	t.Parallel()

	divergent := make(map[string]struct{}, len(graphDivergentSharedCases))
	for _, name := range graphDivergentSharedCases {
		divergent[name] = struct{}{}
	}

	var checked int

	for _, entry := range sharedCaseModels(t) {
		if _, ok := divergent[entry.Case.Name]; !ok {
			continue
		}

		checked++

		require.Emptyf(t, entry.Case.ExpectedErrors,
			"case %q expects findings, so it is not this divergence", entry.Case.Name)

		findings := ExtractAllAs[*Finding](ValidateDSLWithGraph(entry.Model, entry.Case.DSL))

		require.Lenf(t, findings, 1, "case %q reports more than the unbuildable finding", entry.Case.Name)
		assert.Equalf(t, GraphModelUnbuildable, findings[0].Metadata.Kind,
			"case %q diverges by reporting something other than the unbuildable finding", entry.Case.Name)
		assert.Nil(t, findings[0].Line)
	}

	assert.Len(t, graphDivergentSharedCases, checked, "every named case was found in the shared file")
}

// TestEntrypointRuleMatchesTraversalWhereTheGraphBuilds is the parity guarantee
// behind the switch: wherever there is a graph to read, reading it names the same
// unreachable relations as walking the rewrite tree, with the same message and
// the same position.
//
// It compares whole findings rather than counts, so a rule that found the right
// relations and resolved them to the wrong line fails here.
func TestEntrypointRuleMatchesTraversalWhereTheGraphBuilds(t *testing.T) {
	t.Parallel()

	var built, firedOn int

	for _, entry := range sharedCaseModels(t) {
		if _, err := graph.NewWeightedAuthorizationModelGraphBuilder().Build(entry.Model); err != nil {
			continue
		}

		built++

		fromTraversal := entrypointFindings(entry.Model, entry.Case.DSL, validateEntryPoints)
		fromGraph := entrypointFindings(entry.Model, entry.Case.DSL, validateWithGraph)

		if len(fromGraph) > 0 {
			firedOn++
		}

		assert.Equalf(t, fromTraversal, fromGraph,
			"the two paths report different entrypoint findings for %q", entry.Case.Name)
	}

	assert.Equal(t, 62, built, "shared cases the graph builds")
	assert.Equal(t, 7, firedOn,
		"cases the graph rule reports an unreachable relation for; zero here would make the "+
			"comparison above a comparison of empty lists")
}

// entrypointFindings returns the entrypoint findings for a model as sorted strings
// carrying the symbol, the position and the message, which is everything the
// shared-case compares.
func entrypointFindings(model *openfgav1.AuthorizationModel, dsl string, entryPoints entryPointPhase) []string {
	findings := ExtractAllAs[*Finding](validateWith(model, newSource(dsl), entryPoints))

	described := []string{}

	for _, finding := range findings {
		if finding.Metadata.Kind != RelationNoEntrypoint {
			continue
		}

		described = append(described, fmt.Sprintf("%s %s %q",
			finding.Metadata.Symbol,
			describePosition(finding.Line, finding.Column),
			finding.Message))
	}

	slices.Sort(described)

	return described
}

// TestGraphValidationSharedCases runs tests/data/dsl-graph-validation-cases.yaml
// through ValidateDSLWithGraph. The file is the contract for graph-based
// validation the way dsl-semantic-validation-cases.yaml is for the default
// path: it holds the models whose findings exist only under the graph, so an
// implementation that grows a graph path converges on the same messages,
// codes and positions by running it.
func TestGraphValidationSharedCases(t *testing.T) {
	t.Parallel()

	suite, err := NewYAMLTestRunner(sharedDataDir).LoadTestSuite("dsl-graph-validation-cases.yaml")
	require.NoError(t, err)
	require.NotEmpty(t, suite.TestCases, "shared test file loaded no cases")

	for _, testCase := range suite.TestCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			if testCase.Skip {
				t.Skip("marked skip in the shared test file")
			}

			model, err := transformer.TransformDSLToProto(testCase.DSL)
			require.NoError(t, err, "the case's DSL must parse")

			result := compareWithExpected(testCase.ExpectedErrors,
				ExtractAllAs[*Finding](ValidateDSLWithGraph(model, testCase.DSL)))

			for _, problem := range result.Problems {
				t.Error(problem)
			}

			require.Equal(t, casePass, result.Status)
		})
	}
}
