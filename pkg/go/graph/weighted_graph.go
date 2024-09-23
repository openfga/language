package graph

import (
	"fmt"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/multi"
)

type WeightedAuthorizationModelGraph struct {
	*multi.DirectedGraph
	drawingDirection DrawingDirection
}

// nolint: cyclop
func NewWeightedAuthorizationModelGraph(model *openfgav1.AuthorizationModel) (*WeightedAuthorizationModelGraph, error) {
	g, err := NewAuthorizationModelGraph(model)
	if err != nil {
		return nil, err
	}
	g, err = g.Reversed() // we want edges to have the direction of Check when doing the weight assignments later
	if err != nil {
		return nil, err
	}

	graphBuilder := &WeightedAuthorizationModelGraphBuilder{multi.NewDirectedGraph()}

	// Add all nodes
	iterNodes := g.Nodes()
	for iterNodes.Next() {
		nextNode := iterNodes.Node()
		node, ok := nextNode.(*AuthorizationModelNode)
		if !ok {
			return nil, fmt.Errorf("%w: could not cast to WeightedAuthorizationModelNode", ErrBuildingGraph)
		}
		newNode := NewWeightedAuthorizationModelNode(node, false)
		graphBuilder.AddNode(newNode)
	}

	// Add all the edges
	iterEdges := g.Edges()
	for iterEdges.Next() {
		nextEdge, ok := iterEdges.Edge().(multi.Edge)
		if !ok {
			return nil, fmt.Errorf("%w: could not cast %v to multi.Edge", ErrBuildingGraph, iterEdges.Edge())
		}

		// NOTE: because we use a multigraph, one edge can include multiple lines, so we need to add each line individually.
		iterLines := nextEdge.Lines
		for iterLines.Next() {
			nextLine := iterLines.Line()
			castedEdge, ok := nextLine.(*AuthorizationModelEdge)
			if !ok {
				return nil, fmt.Errorf("%w: could not cast %v to AuthorizationModelEdge", ErrBuildingGraph, nextLine)
			}

			err = graphBuilder.AddEdgeAndUpdateNodes(castedEdge)
			if err != nil {
				return nil, err
			}
		}
	}

	err = graphBuilder.AssignWeights()
	if err != nil {
		return nil, err
	}

	multigraph, ok := graphBuilder.DirectedMultigraphBuilder.(*multi.DirectedGraph)
	if !ok {
		return nil, fmt.Errorf("%w: could not cast to DirectedGraph", ErrBuildingGraph)
	}

	wg := &WeightedAuthorizationModelGraph{multigraph, g.drawingDirection}

	return wg, nil
}

var _ dot.Attributers = (*WeightedAuthorizationModelGraph)(nil)

func (wb *WeightedAuthorizationModelGraph) DOTAttributers() (encoding.Attributer, encoding.Attributer, encoding.Attributer) {
	return wb, nil, nil
}

func (wb *WeightedAuthorizationModelGraph) Attributes() []encoding.Attribute {
	rankdir := "BT" // bottom to top
	if wb.drawingDirection == DrawingDirectionCheck {
		rankdir = "TB" // top to bottom
	}

	return []encoding.Attribute{{
		Key:   "rankdir", // https://graphviz.org/docs/attrs/rankdir/
		Value: rankdir,
	}}
}

// GetDOT returns the DOT visualization. The output text is stable.
// It should only be used for debugging.
func (wb *WeightedAuthorizationModelGraph) GetDOT() string {
	dotRepresentation, err := dot.MarshalMulti(wb, "", "", "")
	if err != nil {
		return ""
	}

	return string(dotRepresentation)
}
