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

// maxCheckedCycles caps the brute-force search, so a model too big to check fails.
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

// describeRelationErrors renders each error as its kind and relations.
func describeRelationErrors(relationErrs []*RelationError) []string {
	described := make([]string, 0, len(relationErrs))

	for _, relationErr := range relationErrs {
		var kind string

		switch {
		case errors.Is(relationErr, ErrModelCycle):
			kind = "model cycle"
		case errors.Is(relationErr, ErrTupleCycle):
			kind = "tuple cycle"
		case errors.Is(relationErr, ErrInvalidModel):
			kind = "invalid model"
		default:
			kind = relationErr.Err.Error()
		}

		described = append(described, kind+": "+strings.Join(relationErr.Relations, ", "))
	}

	return described
}

// cycleModelProblems parses a model from testdata/cycles and returns its RelationErrors.
func cycleModelProblems(t *testing.T, name string) RelationErrors {
	t.Helper()

	_, err := NewWeightedAuthorizationModelGraphBuilder().Parse(cycleModel(t, name))

	relationErrs := relationErrorsIn(err)
	require.NotEmpty(t, relationErrs, name)

	return relationErrs
}

// TestParseNamesEveryCycleInLargeModels checks models with seven to eleven rejected cycles.
// The relations organization#c and team#h_c are named because they share a component with one.
func TestParseNamesEveryCycleInLargeModels(t *testing.T) {
	t.Parallel()

	tests := map[string][]string{
		"file-storage-ten-loops": {
			"tuple cycle: drive#can_manage, organization#can_admin_drive",
			"model cycle: drive#editor, drive#writer",
			"model cycle: drive#r1, drive#r2, drive#r3, drive#r4",
			"model cycle: file#allowed, file#can_view",
			"model cycle: file#x, file#y, file#z",
			"tuple cycle: folder#can_view, folder#restricted",
			"tuple cycle: folder#reader",
			"tuple cycle: group#member, group#suspended",
			"tuple cycle: organization#a, organization#b, organization#c",
			"model cycle: organization#admin, organization#member",
		},
		"knowledge-base-seven-loops": {
			"tuple cycle: article#can_feature, container#featured",
			"tuple cycle: article#can_view, article#hidden, attachment#can_view",
			"model cycle: article#editor, article#reviewer",
			"tuple cycle: container#editor",
			"model cycle: container#l1, container#l2, container#l3, container#l4, container#l5, container#l6",
			"model cycle: group#p, group#q, group#r",
			"model cycle: organization#editor, organization#member",
		},
		"pattern-matrix-eleven-loops": {
			"tuple cycle: project#b4_x, team#b4_back",
			"tuple cycle: project#b5_x, team#b5_y",
			"model cycle: team#a1_x, team#a1_y",
			"model cycle: team#a2_x, team#a2_y",
			"model cycle: team#a3_x, team#a3_y",
			"model cycle: team#a4_x, team#a4_y",
			"tuple cycle: team#b1_member",
			"tuple cycle: team#b2_x, team#b2_y",
			"tuple cycle: team#b3_x",
			"tuple cycle: team#h_a, team#h_b, team#h_c",
			"model cycle: team#m_p, team#m_q, team#m_r",
		},
	}

	for name, want := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			for range 20 {
				require.Equal(t, want, describeRelationErrors(cycleModelProblems(t, name)))
			}
		})
	}
}

// TestParseReportsEveryProblem checks that a rejected cycle doesn't hide another problem.
func TestParseReportsEveryProblem(t *testing.T) {
	t.Parallel()

	tests := map[string][]string{
		// ghost needs ghost2, which has no entry point; the builder reports that as a model cycle.
		"mixed-no-entry-point": {
			"model cycle: drive#editor, drive#writer",
			"model cycle: folder#ghost",
		},
		"mixed-self-rewrite": {"model cycle: drive#editor, drive#writer"},
		"mixed-missing-from-target": {
			"model cycle: drive#editor, drive#writer",
			"invalid model: folder#can_preview",
		},
		"everything-at-once": {
			"tuple cycle: project#b4_x, team#b4_back",
			"tuple cycle: project#b5_x, team#b5_y",
			"invalid model: project#can_hold",
			"model cycle: team#a1_x, team#a1_y",
			"model cycle: team#a2_x, team#a2_y",
			"model cycle: team#a3_x, team#a3_y",
			"model cycle: team#a4_x, team#a4_y",
			"tuple cycle: team#b1_member",
			"tuple cycle: team#b2_x, team#b2_y",
			"tuple cycle: team#b3_x",
			"tuple cycle: team#h_a, team#h_b, team#h_c",
			"model cycle: team#m_p, team#m_q, team#m_r",
		},
	}

	for name, want := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, want, describeRelationErrors(cycleModelProblems(t, name)))
		})
	}
}

// TestRejectedCyclesMatchTheChecker holds rejectedCycles to a brute-force search over every
// simple cycle of the shared and testdata models, and checks that Parse reports each one.
func TestRejectedCyclesMatchTheChecker(t *testing.T) {
	t.Parallel()

	accepted, refused, withCycles := 0, 0, 0

	for name, dsl := range checkerModels(t) {
		model, err := transformer.TransformDSLToProto(dsl)
		if err != nil {
			continue
		}

		wg, _ := NewWeightedAuthorizationModelGraphBuilder().construct(model, func(*RelationError) {})
		check := checkCycles(t, wg)
		cycles, _ := wg.rejectedCycles(wg.stronglyConnectedComponents(nil))

		_, err = NewWeightedAuthorizationModelGraphBuilder().Parse(model)
		if err == nil {
			accepted++

			assert.Emptyf(t, check.modelCycle, "%s: accepted, but the search finds a model cycle", name)
			assert.Emptyf(t, check.andOrButNot, "%s: accepted, but the search finds a tuple cycle through an and or a but not", name)
			assert.Emptyf(t, cycles, "%s: accepted, but the scan names a cycle", name)

			continue
		}

		refused++

		reported := map[string]bool{}
		for _, relationErr := range problemsOf(t, err) {
			reported[relationErr.Error()] = true
		}

		named := map[string]bool{} // relation: named for a model cycle
		for _, cycle := range cycles {
			assert.Truef(t, reported[cycle.Error()], "%s: Parse does not report %q", name, cycle.Error())

			for _, relation := range cycle.Relations {
				named[relation] = errors.Is(cycle, ErrModelCycle)
			}
		}

		if len(named) > 0 {
			withCycles++
		}

		for relation := range check.modelCycle {
			modelCycle, ok := named[relation]
			assert.Truef(t, ok && modelCycle, "%s: %s is on a model cycle", name, relation)
		}

		for relation := range check.andOrButNot {
			modelCycle, ok := named[relation]
			assert.Truef(t, ok && !modelCycle, "%s: %s is on a tuple cycle through an and or a but not", name, relation)
		}

		for relation, modelCycle := range named {
			if check.modelCycle[relation] || check.andOrButNot[relation] {
				continue
			}

			assert.Falsef(t, modelCycle, "%s: %s is named for a model cycle it is not on", name, relation)
			assert.Truef(t, check.groupedWithOperator(relation), "%s: %s is named but is not grouped with an and or a but not", name, relation)
		}
	}

	assert.Positive(t, accepted, "no accepted model was checked")
	assert.Positive(t, withCycles, "no refused model with a cycle was checked")

	t.Logf("checked %d accepted and %d refused models, %d of them with cycles", accepted, refused, withCycles)
}

// checkerModels returns the models the checker runs over, keyed by source.
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

// cycleCheck is what a brute-force search over every simple cycle finds, written apart
// from rejectedCycles so the two can be compared.
type cycleCheck struct {
	modelCycle  map[string]bool // relations on a cycle with no TTU or userset edge
	andOrButNot map[string]bool // relations on a cycle through an and or a but not, less those above

	parent    map[string]string // nodes joined by the cycles they share
	operators []string          // the and and but-not nodes found on a cycle
}

// checkCycles enumerates every simple cycle once, from its smallest node, and records each.
func checkCycles(t *testing.T, wg *WeightedAuthorizationModelGraph) cycleCheck {
	t.Helper()

	check := cycleCheck{modelCycle: map[string]bool{}, andOrButNot: map[string]bool{}, parent: map[string]string{}}

	labels := slices.Sorted(maps.Keys(wg.nodes))
	rank := make(map[string]int, len(labels))

	for i, label := range labels {
		rank[label] = i
	}

	cycles := 0

	var (
		path          []string // from the cycle's start, start included
		ttuOrUsersets []bool   // whether each step along path is a TTU or userset edge
	)

	var search func(at string)
	search = func(at string) {
		for _, edge := range wg.edges[at] {
			to := edge.to.uniqueLabel
			// isTTUOrUsersetEdge, written out so the checker doesn't share the code it checks.
			ttuOrUserset := edge.edgeType == TTUEdge || (edge.edgeType == DirectEdge && edge.to.nodeType == SpecificTypeAndRelation)

			switch {
			case to == path[0]:
				cycles++
				require.LessOrEqual(t, cycles, maxCheckedCycles, "too many simple cycles to check")
				check.record(wg, path, ttuOrUserset || slices.Contains(ttuOrUsersets, true))
			case rank[to] > rank[path[0]] && !slices.Contains(path, to):
				path, ttuOrUsersets = append(path, to), append(ttuOrUsersets, ttuOrUserset)
				search(to)
				path, ttuOrUsersets = path[:len(path)-1], ttuOrUsersets[:len(ttuOrUsersets)-1]
			}
		}
	}

	for _, start := range labels {
		path, ttuOrUsersets = []string{start}, nil
		search(start)
	}

	for relation := range check.modelCycle {
		delete(check.andOrButNot, relation)
	}

	return check
}

// record files a closed cycle under the rule it breaks, if any. A relation pointing at
// itself breaks neither.
func (c *cycleCheck) record(wg *WeightedAuthorizationModelGraph, cycle []string, throughTTUOrUserset bool) {
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

	if throughTTUOrUserset && !throughOperator {
		return
	}

	for _, label := range cycle {
		if wg.nodes[label].nodeType != SpecificTypeAndRelation {
			continue
		}

		if throughTTUOrUserset {
			c.andOrButNot[label] = true
		} else {
			c.modelCycle[label] = true
		}
	}
}

// groupedWithOperator reports whether relation shares a group of joined cycles with an and or but not.
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
