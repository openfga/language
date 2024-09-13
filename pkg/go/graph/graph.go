package graph

import (
	"errors"
	"fmt"

	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/multi"
)

var ErrBuildingGraph = errors.New("cannot build graph")

type DrawingDirection bool

const (
	DrawingDirectionListObjects DrawingDirection = true
	DrawingDirectionCheck       DrawingDirection = false
)

type AuthorizationModelGraph struct {
	*multi.DirectedGraph
	drawingDirection DrawingDirection
}

// Reversed returns a full copy of the graph, but with the direction of the arrows flipped.
func (g *AuthorizationModelGraph) Reversed() (*AuthorizationModelGraph, error) {
	graphBuilder := &AuthorizationModelGraphBuilder{
		multi.NewDirectedGraph(), map[string]int64{},
	}

	// Add all nodes as-is.
	iterNodes := g.Nodes()
	for iterNodes.Next() {
		nextNode := iterNodes.Node()
		graphBuilder.AddNode(nextNode)
	}

	// Add all edges as-is, but with their From and To flipped.
	iterEdges := g.Edges()
	for iterEdges.Next() {
		nextEdge, ok := iterEdges.Edge().(multi.Edge)
		if !ok {
			return nil, fmt.Errorf("%w: could not cast to multi.Edge", ErrBuildingGraph)
		}
		// NOTE: because we use a multigraph, one edge can include multiple lines, so we need to add each line individually.
		iterLines := nextEdge.Lines
		for iterLines.Next() {
			nextLine := iterLines.Line()
			casted, ok := nextLine.(*AuthorizationModelEdge)
			if !ok {
				return nil, fmt.Errorf("%w: could not cast to AuthorizationModelEdge", ErrBuildingGraph)
			}
			graphBuilder.AddEdge(nextLine.To(), nextLine.From(), casted.edgeType, casted.conditionedOn)
		}
	}

	multigraph, ok := graphBuilder.DirectedMultigraphBuilder.(*multi.DirectedGraph)
	if ok {
		return &AuthorizationModelGraph{multigraph, !g.drawingDirection}, nil
	}

	return nil, fmt.Errorf("%w: could not cast to directed graph", ErrBuildingGraph)
}

var _ dot.Attributers = (*AuthorizationModelGraph)(nil)

func (g *AuthorizationModelGraph) DOTAttributers() (encoding.Attributer, encoding.Attributer, encoding.Attributer) {
	return g, nil, nil
}

func (g *AuthorizationModelGraph) Attributes() []encoding.Attribute {
	// https://graphviz.org/docs/attrs/rankdir/
	if g.drawingDirection == DrawingDirectionCheck {
		return []encoding.Attribute{{
			Key:   "rankdir",
			Value: "TB",
		}}
	}

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

// TODO add graph traversals, cycle detection, etc.
