package graph

import (
	"errors"
	"slices"
	"strings"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/multi"
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

func (g *AuthorizationModelGraph) AllPathsFromSourceToTarget(source int64, target int64) [][]int64 {
	ans := make([][]int64, 0)
	curr := []int64{source}
	g.backtrack(curr, &ans, target)

	return ans
}

func (g *AuthorizationModelGraph) getPath(p []int64) string {
	var sb strings.Builder
	for i, element := range p {
		n, _ := g.Node(element).(*AuthorizationModelNode)
		sb.WriteString(n.uniqueLabel)
		if i != len(p)-1 {
			sb.WriteString("->")
		}
	}

	return sb.String()
}

func (g *AuthorizationModelGraph) backtrack(curr []int64, ans *[][]int64, target int64) {
	if curr[len(curr)-1] == target {
		currCopy := make([]int64, len(curr))
		copy(currCopy, curr)
		*ans = append(*ans, currCopy)

		return
	}

	neighbors := g.From(curr[len(curr)-1])
	for neighbors.Next() {
		neigh := neighbors.Node().ID()
		if slices.Contains(curr, neigh) {
			// avoid cycles
			continue
		}
		curr = append(curr, neigh)
		g.backtrack(curr, ans, target)
		curr = curr[:len(curr)-1]
	}
}

func (g *AuthorizationModelGraph) GetAllTruthyPaths() string {
	var allPathsString strings.Builder
	nonLeafs := g.GetNonLeafNodes()
	leafs := g.GetLeafNodeIDs()

	for _, source := range nonLeafs {
		for _, target := range leafs {
			allPaths := g.AllPathsFromSourceToTarget(source.ID(), target)
			if len(allPaths) > 0 {
				for _, path := range allPaths {
					allPathsString.WriteString(g.getPath(path) + "\n")
				}
			}
		}
	}

	return allPathsString.String()
}

func (g *AuthorizationModelGraph) GetNonLeafNodes() []graph.Node {
	var nonLeafNodeIDs []graph.Node
	leafNodes := g.GetLeafNodeIDs()
	iterNodes := g.Nodes()
	for iterNodes.Next() {
		node := iterNodes.Node()
		if _, ok := slices.BinarySearch(leafNodes, node.ID()); ok {
			continue
		}
		converted, _ := node.(*AuthorizationModelNode)
		if converted.nodeType == OperatorNode {
			continue
		}
		nonLeafNodeIDs = append(nonLeafNodeIDs, node)
	}

	return nonLeafNodeIDs
}

func (g *AuthorizationModelGraph) GetLeafNodeIDs() []int64 {
	var noOutGoingEdges []int64
	iterNodes := g.Nodes()
	for iterNodes.Next() {
		node := iterNodes.Node()
		if g.From(node.ID()).Len() == 0 {
			noOutGoingEdges = append(noOutGoingEdges, node.ID())
		}
	}

	slices.Sort(noOutGoingEdges)

	return noOutGoingEdges
}

// TODO add graph traversals, cycle detection, etc.
