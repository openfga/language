package graph

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/openfga/language/pkg/go/transformer"
)

// driftModels returns the models the drift test runs through both builders,
// keyed by where each came from: the shared transformer models, the DSL of the
// shared semantic validation cases, and the cycle models in testdata.
func driftModels(t *testing.T) map[string]string {
	t.Helper()

	shared := filepath.Join("..", "..", "..", "tests", "data")

	transformerModels, err := filepath.Glob(filepath.Join(shared, "transformer", "*", "authorization-model.fga"))
	require.NoError(t, err)

	cycleModels, err := filepath.Glob(filepath.Join("testdata", "cycles", "*.fga"))
	require.NoError(t, err)

	models := map[string]string{}

	for _, file := range append(transformerModels, cycleModels...) {
		dsl, err := os.ReadFile(file)
		require.NoError(t, err)

		models[file] = string(dsl)
	}

	data, err := os.ReadFile(filepath.Join(shared, "dsl-semantic-validation-cases.yaml"))
	require.NoError(t, err)

	var cases []struct {
		Name string `yaml:"name"`
		DSL  string `yaml:"dsl"`
	}
	require.NoError(t, yaml.Unmarshal(data, &cases))

	for _, c := range cases {
		models["semantic validation case: "+c.Name] = c.DSL
	}

	return models
}

// graphShape lists what the weight walk decides about wg, one line per node and
// edge, sorted, so two graphs compare equal whatever order they were walked in.
func graphShape(wg *WeightedAuthorizationModelGraph) []string {
	var lines []string

	for id, node := range wg.nodes {
		lines = append(lines, fmt.Sprintf("node %s: type %d, weights %v, wildcards %v, recursive %q, tuple cycle %t",
			id, node.nodeType, node.weights, slices.Sorted(slices.Values(node.wildcards)),
			node.recursiveRelation, node.tupleCycle))
	}

	for from, edges := range wg.edges {
		for _, edge := range edges {
			lines = append(lines, fmt.Sprintf("edge %s -> %s: type %d, tupleset %q, weights %v, wildcards %v, recursive %q, tuple cycle %t, conditions %v",
				from, edge.to.uniqueLabel, edge.edgeType, edge.tuplesetRelation, edge.weights,
				slices.Sorted(slices.Values(edge.wildcards)), edge.recursiveRelation, edge.tupleCycle, edge.conditions))
		}
	}

	slices.Sort(lines)

	return lines
}

// TestRelationWalkWeighsWhatBuildWeighs holds the two weight walks to each
// other on every shared model: the walk Parse uses for a rejected model, which
// starts from the relations in order, accepts exactly the models AssignWeights
// accepts and weighs them the same. They share the weight calculation for each
// node, but walk the nodes in separate loops, so this is what catches one loop
// changing without the other.
func TestRelationWalkWeighsWhatBuildWeighs(t *testing.T) {
	t.Parallel()

	accepted, rejected := 0, 0

	for name, dsl := range driftModels(t) {
		model, err := transformer.TransformDSLToProto(dsl)
		if err != nil {
			continue
		}

		built, buildErr := NewWeightedAuthorizationModelGraphBuilder().Build(model)

		walked, walkErr := NewWeightedAuthorizationModelGraphBuilder().construct(model, nil)
		if walkErr == nil {
			walkErr = walked.assignWeightsByRelation()
		}

		if !assert.Equalf(t, buildErr == nil, walkErr == nil,
			"%s: only one walk accepts it (Build: %v, relation walk: %v)", name, buildErr, walkErr) {
			continue
		}

		if buildErr != nil {
			rejected++

			continue
		}

		accepted++

		assert.Equalf(t, graphShape(built), graphShape(walked), "%s: the two walks weigh it differently", name)
	}

	assert.Positive(t, accepted, "no shared model was accepted")
	assert.Positive(t, rejected, "no shared model was rejected")
}

// TestParseAcceptsWhatBuildAccepts pins Parse to Build on every shared model:
// for a model Build accepts, Parse returns the same graph, and for one Build
// rejects, Parse's error names at least one problem.
func TestParseAcceptsWhatBuildAccepts(t *testing.T) {
	t.Parallel()

	for name, dsl := range driftModels(t) {
		model, err := transformer.TransformDSLToProto(dsl)
		if err != nil {
			continue
		}

		built, buildErr := NewWeightedAuthorizationModelGraphBuilder().Build(model)
		parsed, parseErr := NewWeightedAuthorizationModelGraphBuilder().Parse(model)

		if buildErr == nil {
			require.NoErrorf(t, parseErr, "%s: Parse rejects a model Build accepts", name)
			assert.Equalf(t, graphShape(built), graphShape(parsed), "%s: Parse returns a different graph", name)

			continue
		}

		assert.Nilf(t, parsed, "%s: Parse returns a graph for a model Build rejects", name)
		assert.NotEmptyf(t, relationErrorsIn(parseErr), "%s: Parse's error names no problem", name)
	}
}
