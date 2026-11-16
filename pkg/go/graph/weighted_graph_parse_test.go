package graph

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
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

// TestComponentWalkWeighsWhatBuildWeighs holds the two weight walks to each
// other on every shared model: the walk Parse uses for a rejected model, which
// goes through the strongly connected components with what each depends on
// first, accepts exactly the models AssignWeights accepts and weighs them the
// same. They share the weight calculation for each node, but walk the nodes in
// separate loops, so this is what catches one loop changing without the other.
func TestComponentWalkWeighsWhatBuildWeighs(t *testing.T) {
	t.Parallel()

	accepted, rejected := 0, 0

	for name, dsl := range driftModels(t) {
		model, err := transformer.TransformDSLToProto(dsl)
		if err != nil {
			continue
		}

		built, buildErr := NewWeightedAuthorizationModelGraphBuilder().Build(model)

		walked, constructErr := NewWeightedAuthorizationModelGraphBuilder().construct(model, nil)

		var walkErr error
		if constructErr != nil {
			walkErr = constructErr
		} else if walkErrs := walked.assignWeightsByComponent(walked.components(nil), map[string]bool{}); len(walkErrs) > 0 {
			walkErr = walkErrs
		}

		if !assert.Equalf(t, buildErr == nil, walkErr == nil,
			"%s: only one walk accepts it (Build: %v, component walk: %v)", name, buildErr, walkErr) {
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
// rejects, Parse's error names at least one problem, and every problem names
// the relations it involves.
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

		relationErrs := relationErrorsIn(parseErr)
		assert.NotEmptyf(t, relationErrs, "%s: Parse's error names no problem", name)

		for _, relationErr := range relationErrs {
			assert.NotEmptyf(t, relationErr.Relations, "%s: %q names no relation", name, relationErr.Error())
		}
	}
}

// rejectedTheSameWay parses dsl, which the builder must reject, 200 times, and
// returns its problems once every run has reported the same ones.
func rejectedTheSameWay(t *testing.T, dsl string) RelationErrors {
	t.Helper()

	model, err := transformer.TransformDSLToProto(dsl)
	require.NoError(t, err)

	_, err = NewWeightedAuthorizationModelGraphBuilder().Parse(model)
	relationErrs := problemsOf(t, err)

	for range 199 {
		_, again := NewWeightedAuthorizationModelGraphBuilder().Parse(model)
		require.EqualError(t, again, err.Error())
	}

	return relationErrs
}

// TestParseReportsEveryRootCause pins that the weight walk does not stop at its
// first error: every relation that fails on its own is reported with its
// relations, and a relation that fails only because it depends on one that
// fails is not.
func TestParseReportsEveryRootCause(t *testing.T) {
	t.Parallel()

	missingRelations := driftModels(t)["semantic validation case: multiple missing relation references with module comments on every type line"]
	require.NotEmpty(t, missingRelations, "the shared validation case is missing")

	tests := map[string]struct {
		dsl  string
		want []string
	}{
		"invalid relations in two types": {
			dsl: `model
  schema 1.1
type user
type employee
type alpha
  relations
    define a: [user]
    define b: [employee]
    define c: a and b
type zeta
  relations
    define a: [user]
    define b: [employee]
    define c: a and b
`,
			want: []string{"invalid model: alpha#c", "invalid model: zeta#c"},
		},
		// p needs q, which needs z. Only z is invalid on its own.
		"a chain of relations that need an invalid one": {
			dsl: `model
  schema 1.1
type user
type employee
type document
  relations
    define p: q
    define q: z
    define z: [user] and y
    define y: [employee]
`,
			want: []string{"invalid model: document#z"},
		},
		"a relation that needs a cycle": {
			dsl: `model
  schema 1.1
type user
type document
  relations
    define a: [user] or b
    define b: [user] or a
    define r: [user] or a
`,
			want: []string{"model cycle: document#a, document#b"},
		},
		// owner and reporter each need a member relation that does not exist.
		"missing relations in two types": {
			dsl:  missingRelations,
			want: []string{"invalid model: page#member", "invalid model: ticket#member"},
		},
		// a and b reach folder#viewer, which has no weights, through the node
		// that groups viewer from parent for both of them. Whichever is walked
		// first leaves that node without weights, and both fail on it either way.
		"relations that share a tuple to userset with no weights": {
			dsl: `model
  schema 1.1
type user
type drive
  relations
    define viewer: [user]
type folder
  relations
    define parent: [folder, drive]
    define parent2: [folder]
    define viewer: viewer from parent2
    define a: [user] or viewer from parent
    define b: [user] and viewer from parent
`,
			want: []string{"model cycle: folder#a", "model cycle: folder#b"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, test.want, describeRelationErrors(rejectedTheSameWay(t, test.dsl)))
		})
	}
}

// TestParseReportsTheSameProblemsWhateverTheTypeNames pins that a cycle does
// not hide an invalid relation in another type, whichever type sorts first.
func TestParseReportsTheSameProblemsWhateverTheTypeNames(t *testing.T) {
	t.Parallel()

	reported := func(cycleType, invalidType string) []string {
		relationErrs := rejectedTheSameWay(t, fmt.Sprintf(`model
  schema 1.1
type user
type employee
type %s
  relations
    define a: [user] or b
    define b: [user] or a
type %s
  relations
    define x: [user] and y
    define y: [employee]
`, cycleType, invalidType))

		// Name the types for what is wrong in them, so both orders compare equal.
		byRole := strings.NewReplacer(cycleType+"#", "cycle#", invalidType+"#", "invalid#")

		described := describeRelationErrors(relationErrs)
		for i := range described {
			described[i] = byRole.Replace(described[i])
		}

		slices.Sort(described)

		return described
	}

	want := []string{"invalid model: invalid#x", "model cycle: cycle#a, cycle#b"}

	assert.Equal(t, want, reported("aaa", "zzz"), "the cycle sorts first")
	assert.Equal(t, want, reported("zzz", "aaa"), "the invalid relation sorts first")
}
