package graph

import (
	"errors"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/openfga/language/pkg/go/transformer"
)

// maxCheckedCycles caps the brute-force search, so a model that grows past what
// it can enumerate fails loudly instead of being checked partly.
const maxCheckedCycles = 200_000

// cycleModel parses a model from testdata/cycles.
func cycleModel(t *testing.T, name string) *openfgav1.AuthorizationModel {
	t.Helper()

	dsl, err := os.ReadFile(filepath.Join("testdata", "cycles", name+".fga"))
	require.NoError(t, err)

	model, err := transformer.TransformDSLToProto(string(dsl))
	require.NoError(t, err)

	return model
}

// describeCycles renders each cycle as its kind and members, so a whole refusal
// compares as one value.
func describeCycles(cycles []Cycle) []string {
	described := make([]string, 0, len(cycles))

	for _, cycle := range cycles {
		kind := "and or but not"
		if errors.Is(cycle.Kind, ErrModelCycle) {
			kind = "no stored data"
		}

		described = append(described, kind+": "+strings.Join(cycle.Members, ", "))
	}

	return described
}

// TestBuildWithCyclesNamesEveryLoopInLargeModels covers models with seven to
// eleven broken loops next to legal ones: every broken loop comes back with all
// of its relations, no legal loop does, and every build agrees.
//
// The relations organization#c and team#h_c sit on a legal loop that shares a
// relation with a broken one. The scan names the whole group, so they are named
// too.
func TestBuildWithCyclesNamesEveryLoopInLargeModels(t *testing.T) {
	t.Parallel()

	tests := map[string][]string{
		"file-storage-ten-loops": {
			"and or but not: drive#can_manage, organization#can_admin_drive",
			"no stored data: drive#editor, drive#writer",
			"no stored data: drive#r1, drive#r2, drive#r3, drive#r4",
			"no stored data: file#allowed, file#can_view",
			"no stored data: file#x, file#y, file#z",
			"and or but not: folder#can_view, folder#restricted",
			"and or but not: folder#reader",
			"and or but not: group#member, group#suspended",
			"and or but not: organization#a, organization#b, organization#c",
			"no stored data: organization#admin, organization#member",
		},
		"knowledge-base-seven-loops": {
			"and or but not: article#can_feature, container#featured",
			"and or but not: article#can_view, article#hidden, attachment#can_view",
			"no stored data: article#editor, article#reviewer",
			"and or but not: container#editor",
			"no stored data: container#l1, container#l2, container#l3, container#l4, container#l5, container#l6",
			"no stored data: group#p, group#q, group#r",
			"no stored data: organization#editor, organization#member",
		},
		"pattern-matrix-eleven-loops": {
			"and or but not: project#b4_x, team#b4_back",
			"and or but not: project#b5_x, team#b5_y",
			"no stored data: team#a1_x, team#a1_y",
			"no stored data: team#a2_x, team#a2_y",
			"no stored data: team#a3_x, team#a3_y",
			"no stored data: team#a4_x, team#a4_y",
			"and or but not: team#b1_member",
			"and or but not: team#b2_x, team#b2_y",
			"and or but not: team#b3_x",
			"and or but not: team#h_a, team#h_b, team#h_c",
			"no stored data: team#m_p, team#m_q, team#m_r",
		},
	}

	for name, want := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			model := cycleModel(t, name)

			for range 20 {
				_, err := NewWeightedAuthorizationModelGraphBuilder().BuildWithCycles(model)

				var cycleErr *CycleError
				require.ErrorAs(t, err, &cycleErr)
				require.Equal(t, want, describeCycles(cycleErr.Cycles))
			}
		})
	}
}

// TestBuildWithCyclesBesideOtherProblems covers broken loops in models that have
// another problem too.
func TestBuildWithCyclesBesideOtherProblems(t *testing.T) {
	t.Parallel()

	// A relation with no entry point, or one that only rewrites itself, is not a
	// loop the builder refuses, so only the broken loop is named.
	t.Run("the loop is named next to other relations", func(t *testing.T) {
		t.Parallel()

		for _, name := range []string{"mixed-no-entry-point", "mixed-self-rewrite"} {
			_, err := NewWeightedAuthorizationModelGraphBuilder().BuildWithCycles(cycleModel(t, name))

			var cycleErr *CycleError
			require.ErrorAs(t, err, &cycleErr, name)
			assert.Equal(t, []string{"no stored data: drive#editor, drive#writer"}, describeCycles(cycleErr.Cycles), name)
		}
	})

	// A missing from target stops the build while the graph is still being
	// constructed, before any loop is looked for, so the loops in these models
	// are not named. This pins today's behavior.
	t.Run("a missing from target hides the loops", func(t *testing.T) {
		t.Parallel()

		for _, name := range []string{"mixed-missing-from-target", "everything-at-once"} {
			_, err := NewWeightedAuthorizationModelGraphBuilder().BuildWithCycles(cycleModel(t, name))
			require.ErrorIs(t, err, ErrInvalidModel, name)

			var cycleErr *CycleError
			assert.NotErrorAs(t, err, &cycleErr, name)
		}
	})
}

// TestUnresolvableCyclesMatchTheChecker holds the structural scan to a
// brute-force search over every simple cycle, in both directions, over the
// models in testdata/cycles, the transformer fixtures and the shared validation
// cases. On a model the builder accepts, neither finds a broken loop. On a model
// it refuses for a loop, the scan names every relation the search finds, with
// the same kind, and names nothing else but relations grouped with an and or a
// but not.
func TestUnresolvableCyclesMatchTheChecker(t *testing.T) {
	t.Parallel()

	accepted, refused := 0, 0

	for name, dsl := range checkerModels(t) {
		model, err := transformer.TransformDSLToProto(dsl)
		if err != nil {
			continue
		}

		wg, err := NewWeightedAuthorizationModelGraphBuilder().construct(model)
		if err != nil {
			continue
		}

		check := checkCycles(t, wg)

		_, err = NewWeightedAuthorizationModelGraphBuilder().BuildWithCycles(model)

		switch {
		case err == nil:
			accepted++

			assert.Emptyf(t, check.noStoredData, "%s: accepted, but the search finds a loop with no stored data", name)
			assert.Emptyf(t, check.andOrButNot, "%s: accepted, but the search finds a loop through an and or a but not", name)
			assert.Emptyf(t, wg.unresolvableCycles(), "%s: accepted, but the scan names a loop", name)
		case errors.Is(err, ErrModelCycle) || errors.Is(err, ErrTupleCycle):
			refused++

			var cycleErr *CycleError
			require.ErrorAsf(t, err, &cycleErr, "%s", name)
			assert.Positivef(t, len(check.noStoredData)+len(check.andOrButNot), "%s: refused, but the search finds no broken loop", name)

			named := map[string]bool{} // relation: named for a loop with no stored data
			for _, cycle := range cycleErr.Cycles {
				for _, member := range cycle.Members {
					named[member] = errors.Is(cycle.Kind, ErrModelCycle)
				}
			}

			for relation := range check.noStoredData {
				noStoredData, ok := named[relation]
				assert.Truef(t, ok && noStoredData, "%s: %s is on a loop with no stored data", name, relation)
			}

			for relation := range check.andOrButNot {
				noStoredData, ok := named[relation]
				assert.Truef(t, ok && !noStoredData, "%s: %s is on a loop through an and or a but not", name, relation)
			}

			for relation, noStoredData := range named {
				if check.noStoredData[relation] || check.andOrButNot[relation] {
					continue
				}

				assert.Falsef(t, noStoredData, "%s: %s is named for a loop with no stored data it is not on", name, relation)
				assert.Truef(t, check.groupedWithOperator(relation), "%s: %s is named but is not grouped with an and or a but not", name, relation)
			}
		}
	}

	assert.Positive(t, accepted, "no accepted model was checked")
	assert.Positive(t, refused, "no refused model was checked")

	t.Logf("checked %d accepted and %d refused models", accepted, refused)
}

// checkerModels gathers the models the checker runs over, keyed by where they
// come from.
func checkerModels(t *testing.T) map[string]string {
	t.Helper()

	models := map[string]string{}

	patterns := []string{
		filepath.Join("testdata", "cycles", "*.fga"),
		filepath.Join("..", "..", "..", "tests", "data", "transformer", "*", "authorization-model.fga"),
	}

	for _, pattern := range patterns {
		paths, err := filepath.Glob(pattern)
		require.NoError(t, err)

		for _, path := range paths {
			dsl, err := os.ReadFile(path)
			require.NoError(t, err)

			models[path] = string(dsl)
		}
	}

	raw, err := os.ReadFile(filepath.Join("..", "..", "..", "tests", "data", "dsl-semantic-validation-cases.yaml"))
	require.NoError(t, err)

	var cases []struct {
		Name string `yaml:"name"`
		DSL  string `yaml:"dsl"`
		Skip bool   `yaml:"skip"`
	}
	require.NoError(t, yaml.Unmarshal(raw, &cases))

	for _, c := range cases {
		if !c.Skip {
			models["shared case: "+c.Name] = c.DSL
		}
	}

	require.NotEmpty(t, models)

	return models
}

// cycleCheck is what a brute-force search over every simple cycle of a
// constructed graph finds. It is written apart from unresolvableCycles so the
// two can be held to each other.
type cycleCheck struct {
	noStoredData map[string]bool // relations on a cycle with no tuple-reading edge
	andOrButNot  map[string]bool // relations on a cycle through an and or a but not, less those above

	parent    map[string]string // nodes joined by the cycles they share
	operators []string          // the and and but-not nodes found on a cycle
}

// checkCycles enumerates every simple cycle once, from its smallest node, and
// applies the two rules to each.
func checkCycles(t *testing.T, wg *WeightedAuthorizationModelGraph) cycleCheck {
	t.Helper()

	check := cycleCheck{noStoredData: map[string]bool{}, andOrButNot: map[string]bool{}, parent: map[string]string{}}

	labels := slices.Sorted(maps.Keys(wg.nodes))
	rank := make(map[string]int, len(labels))

	for i, label := range labels {
		rank[label] = i
	}

	cycles := 0

	var (
		path  []string // from the cycle's start, start included
		reads []bool   // whether each step along path reads a tuple
	)

	var search func(at string)
	search = func(at string) {
		for _, edge := range wg.edges[at] {
			to := edge.to.uniqueLabel
			// The weight walk's own test for a step that reads a tuple.
			step := edge.edgeType == TTUEdge || (edge.edgeType == DirectEdge && edge.to.nodeType == SpecificTypeAndRelation)

			switch {
			case to == path[0]:
				cycles++
				require.LessOrEqual(t, cycles, maxCheckedCycles, "too many simple cycles to check")
				check.record(wg, path, step || slices.Contains(reads, true))
			case rank[to] > rank[path[0]] && !slices.Contains(path, to):
				path, reads = append(path, to), append(reads, step)
				search(to)
				path, reads = path[:len(path)-1], reads[:len(reads)-1]
			}
		}
	}

	for _, start := range labels {
		path, reads = []string{start}, nil
		search(start)
	}

	for relation := range check.noStoredData {
		delete(check.andOrButNot, relation)
	}

	return check
}

// record files one closed cycle under the rule it breaks, if any. A single
// relation pointing at itself is the builder's recursive-edge case and breaks
// neither rule.
func (c *cycleCheck) record(wg *WeightedAuthorizationModelGraph, cycle []string, readsTuple bool) {
	for _, node := range cycle[1:] {
		c.join(cycle[0], node)
	}

	if len(cycle) == 1 {
		return
	}

	throughOperator := false

	for _, label := range cycle {
		if node := wg.nodes[label]; node.nodeType == OperatorNode && (node.label == IntersectionOperator || node.label == ExclusionOperator) {
			throughOperator = true
			c.operators = append(c.operators, label)
		}
	}

	if readsTuple && !throughOperator {
		return
	}

	for _, label := range cycle {
		if wg.nodes[label].nodeType != SpecificTypeAndRelation {
			continue
		}

		if readsTuple {
			c.andOrButNot[label] = true
		} else {
			c.noStoredData[label] = true
		}
	}
}

// groupedWithOperator reports whether a relation shares a group of joined cycles
// with an and or a but not.
func (c *cycleCheck) groupedWithOperator(relation string) bool {
	if _, ok := c.parent[relation]; !ok {
		return false
	}

	return slices.ContainsFunc(c.operators, func(operator string) bool {
		return c.find(operator) == c.find(relation)
	})
}

func (c *cycleCheck) join(a, b string) {
	if rootA, rootB := c.find(a), c.find(b); rootA != rootB {
		c.parent[rootA] = rootB
	}
}

func (c *cycleCheck) find(node string) string {
	parent, ok := c.parent[node]
	if !ok {
		c.parent[node] = node

		return node
	}

	if parent == node {
		return node
	}

	root := c.find(parent)
	c.parent[node] = root

	return root
}
