package graph

import (
	"errors"
	"reflect"
	"slices"
	"sort"

	"github.com/openfga/language/pkg/go/utils"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/multi"
	"gonum.org/v1/gonum/graph/topo"
)

var ErrBuildingGraph = errors.New("cannot build graph")

type AuthorizationModelGraph struct {
	*multi.DirectedGraph
}

var _ dot.Attributers = (*AuthorizationModelGraph)(nil)

func (g *AuthorizationModelGraph) DOTAttributers() (encoding.Attributer, encoding.Attributer, encoding.Attributer) {
	return g, nil, nil
}

func (g *AuthorizationModelGraph) Attributes() []encoding.Attribute {
	// https://graphviz.org/docs/attrs/rankdir/ - bottom to top
	return []encoding.Attribute{{
		Key:   "rankdir",
		Value: "BT",
	}}
}

// GetDOT returns the DOT visualization. The output text is stable.
// It should only be used for debugging.
func (g *AuthorizationModelGraph) GetDOT() string {
	dotRepresentation, err := dot.MarshalMulti(g, "", "", "")
	if err != nil {
		return ""
	}

	return string(dotRepresentation)
}

func sortAndRemoveDuplicateAndAlgebraic(orig [][]string) [][]string {
	newSlices := make(utils.SlicesOfSlices, 0, len(orig))
	for _, slice := range orig {
		newSlice := make([]string, 0, len(slice))
		for _, item := range slice {
			if item != union && item != intersection && item != exclusion && !slices.Contains(newSlice, item) {
				newSlice = append(newSlice, item)
			}
		}
		sort.Strings(newSlice)
		if !slicesOfSliceContains(newSlices, newSlice) {
			newSlices = append(newSlices, newSlice)
		}
	}
	// now, sort the slices according to size
	sort.Sort(newSlices)

	return newSlices
}

// CycleInformation encapsulates whether the graph has cycles.
type CycleInformation struct {
	// If hasCyclesAtCompileTime is non-empty, we should block this model from ever being written.
	// This is because we are trying to perform a Check on it will cause a stack overflow no matter what the tuples are.
	hasCyclesAtCompileTime [][]string

	// If canHaveCyclesAtRuntime is non-empty, there could exist tuples that introduce a cycle.
	canHaveCyclesAtRuntime [][]string
}

// slicesOfSliceContains returns true if the newSlice is deeply equal to any of the origSlices.
func slicesOfSliceContains(origSlices [][]string, newSlice []string) bool {
	for _, origSlice := range origSlices {
		if reflect.DeepEqual(newSlice, origSlice) {
			return true
		}
	}

	return false
}

// SortedHasCyclesAtCompileTime returns a sorted HasCyclesAtCompileTime which removed algebraic operation (such as union/intersection/exclusion)
// The []string are sorted by length. If []string has the same length, it will return if the first/second/third/.. item is smallest
// Within each []string, it is sorted by alphabet.  In addition, the duplicate node is removed.
func (c *CycleInformation) SortedHasCyclesAtCompileTime() [][]string {
	return sortAndRemoveDuplicateAndAlgebraic(c.hasCyclesAtCompileTime)
}

// SortedCanHaveCyclesAtRuntime returns a sorted HasCyclesAtCompileTime which removed algebraic operation (such as union/intersection/exclusion)
// The []string are sorted by length. If []string has the same length, it will return if the first/second/third/.. item is smallest
// Within each []string, it is sorted by alphabet.  In addition, the duplicate node is removed.
func (c *CycleInformation) SortedCanHaveCyclesAtRuntime() [][]string {
	return sortAndRemoveDuplicateAndAlgebraic(c.canHaveCyclesAtRuntime)
}

func (g *AuthorizationModelGraph) nodeListHasNonComputedEdge(nodeList []graph.Node) bool {
	for i, nodeI := range nodeList {
		for _, nodeJ := range nodeList[i+1:] {
			allEdges := g.Lines(nodeI.ID(), nodeJ.ID())
			for allEdges.Next() {
				edge, ok := allEdges.Line().(*AuthorizationModelEdge)
				if ok && (edge.edgeType == TTUEdge || edge.edgeType == DirectEdge) {
					return true
				}
			}
		}
	}

	return false
}

func nodeListIdentifier(nodeList []graph.Node) []string {
	labels := make([]string, 0, len(nodeList))
	for _, node := range nodeList {
		auth, ok := node.(*AuthorizationModelNode)
		if ok {
			labels = append(labels, auth.label)
		}
	}

	return labels
}

func (g *AuthorizationModelGraph) GetCycles() CycleInformation {
	var nodesWithCyclesAtCompileTime [][]string
	var nodesWithCyclesAtRuntime [][]string

	// TODO: investigate whether len(1) should be identified as cycle

	nodes := topo.DirectedCyclesIn(g)
	for _, nodeList := range nodes {
		if g.nodeListHasNonComputedEdge(nodeList) {
			nodesWithCyclesAtRuntime = append(nodesWithCyclesAtRuntime, nodeListIdentifier(nodeList))
		} else {
			nodesWithCyclesAtCompileTime = append(nodesWithCyclesAtCompileTime, nodeListIdentifier(nodeList))
		}
	}

	return CycleInformation{
		hasCyclesAtCompileTime: nodesWithCyclesAtCompileTime,
		canHaveCyclesAtRuntime: nodesWithCyclesAtRuntime,
	}
}

// TODO add graph traversals, etc.
