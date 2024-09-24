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

func (wg *WeightedAuthorizationModelGraph) GetDrawingDirection() DrawingDirection {
	return wg.drawingDirection
}

func NewWeightedAuthorizationModelGraph(model *openfgav1.AuthorizationModel) (*WeightedAuthorizationModelGraph, error) {
	weightedGraphBuilder, err := NewWeightedAuthorizationModelGraphBuilder(model)
	if err != nil {
		return nil, err
	}

	multigraph, ok := weightedGraphBuilder.DirectedMultigraphBuilder.(*multi.DirectedGraph)
	if !ok {
		return nil, fmt.Errorf("%w: could not cast to DirectedGraph", ErrBuildingGraph)
	}

	wg := &WeightedAuthorizationModelGraph{multigraph, weightedGraphBuilder.drawingDirection}

	return wg, nil
}

var _ dot.Attributers = (*WeightedAuthorizationModelGraph)(nil)

func (wg *WeightedAuthorizationModelGraph) DOTAttributers() (encoding.Attributer, encoding.Attributer, encoding.Attributer) {
	return wg, nil, nil
}

func (wg *WeightedAuthorizationModelGraph) Attributes() []encoding.Attribute {
	rankdir := "BT" // bottom to top
	if wg.drawingDirection == DrawingDirectionCheck {
		rankdir = "TB" // top to bottom
	}

	return []encoding.Attribute{{
		Key:   "rankdir", // https://graphviz.org/docs/attrs/rankdir/
		Value: rankdir,
	}}
}

// GetDOT returns the DOT visualization. The output text is stable.
// It should only be used for debugging.
func (wg *WeightedAuthorizationModelGraph) GetDOT() string {
	dotRepresentation, err := dot.MarshalMulti(wg, "", "", "")
	if err != nil {
		return ""
	}

	return string(dotRepresentation)
}
