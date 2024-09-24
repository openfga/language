package graph

import (
	"fmt"
	"math"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/multi"
)

type WeightedAuthorizationModelGraphBuilder struct {
	graph.DirectedMultigraphBuilder
	drawingDirection DrawingDirection
}

//nolint: cyclop
func NewWeightedAuthorizationModelGraphBuilder(model *openfgav1.AuthorizationModel) (*WeightedAuthorizationModelGraphBuilder, error) {
	g, err := NewAuthorizationModelGraph(model)
	if err != nil {
		return nil, err
	}
	g, err = g.Reversed() // we want edges to have the direction of Check when doing the weight assignments later
	if err != nil {
		return nil, err
	}

	graphBuilder := &WeightedAuthorizationModelGraphBuilder{multi.NewDirectedGraph(), DrawingDirectionCheck}

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

	return graphBuilder, nil
}

func (wb *WeightedAuthorizationModelGraphBuilder) AddEdgeAndUpdateNodes(edge *AuthorizationModelEdge) error {
	fromNode := wb.Node(edge.From().ID())
	toNode := wb.Node(edge.To().ID())
	isNested := fromNode == toNode

	// update "isNested" field for the node
	if isNested {
		selfNode, ok := fromNode.(*WeightedAuthorizationModelNode)
		if !ok {
			return fmt.Errorf("%w: could not cast to WeightedAuthorizationModelNode", ErrBuildingGraph)
		}
		selfNode.isNested = true
	}

	// add line
	edge.Line = wb.NewLine(fromNode, toNode)
	newWeightedEdge := NewWeightedAuthorizationModelEdge(edge, isNested)
	wb.DirectedMultigraphBuilder.SetLine(newWeightedEdge)

	return nil
}

func (wb *WeightedAuthorizationModelGraphBuilder) AssignWeights() error {
	seen := make(map[int64]struct{})
	iterNodes := wb.Nodes()
	for iterNodes.Next() {
		nextNode, ok := iterNodes.Node().(*WeightedAuthorizationModelNode)
		if !ok {
			return fmt.Errorf("%w: could not cast to WeightedAuthorizationModelNode", ErrBuildingGraph)
		}

		err := wb.dfsToAssignWeights(nextNode, seen)
		if err != nil {
			return err
		}
	}

	return nil
}

func (wb *WeightedAuthorizationModelGraphBuilder) dfsToAssignWeights(curNode *WeightedAuthorizationModelNode, seen map[int64]struct{}) error {
	if _, seeen := seen[curNode.ID()]; seeen {
		return nil
	}
	seen[curNode.ID()] = struct{}{}

	edgesArray, err := wb.getOutgoingEdges(curNode)
	if err != nil {
		return err
	}

	// first, assign weights to neighbor edges
	for _, edge := range edgesArray {
		if len(edge.weights) > 0 {
			continue // already assigned
		}
		if edge.isNested {
			continue // will be assigned later
		}

		neighborNode, ok := edge.To().(*WeightedAuthorizationModelNode)
		if !ok {
			return fmt.Errorf("%w: could not cast to WeightedAuthorizationModelNode", ErrBuildingGraph)
		}

		// NOTE: type:* nodes don't have a weight, and neither will their edges
		if neighborNode.AuthorizationModelNode.nodeType == SpecificType {
			edge.weights[neighborNode.AuthorizationModelNode.label] = 1
		}

		// recursively assign weights to neighbor nodes
		if err := wb.dfsToAssignWeights(neighborNode, seen); err != nil {
			return err
		}

		if err := assignWeightsToEdge(edge); err != nil {
			return err
		}
	}

	// second, now that all edge weights have been recursively assigned, assign weights to node
	assignWeightsToNode(curNode, edgesArray)

	// third, update edges that are loops
	assignWeightsToLoopEdges(curNode, edgesArray)

	// finally, make sure that intersections and exclusions are "correct"
	return verifyIntersectionAndExclusionNodes(curNode, edgesArray)
}

// verifyIntersectionAndExclusionNodes checks that intersections and exclusions are correct. For example, an intersection operator that has
// a weight map such as "user:1, employee:2" may be valid, but if one of the edges/operands connecting it doesn't have a weight defined for
// "user", then the operator is invalid, because one of the operands of the intersection will never lead to a "user" type.
func verifyIntersectionAndExclusionNodes(curNode *WeightedAuthorizationModelNode, edgesArray []*WeightedAuthorizationModelEdge) error {
	if curNode.nodeType == OperatorNode && !(curNode.label == "union") {
		edgeWeights := make([]WeightMap, len(edgesArray))
		for i := 0; i < len(edgesArray); i++ {
			edgeWeights[i] = edgesArray[i].weights
		}
		intersect, err := IntersectionOfKeys(edgeWeights...)
		if err != nil {
			return err
		}

		if len(intersect) == 0 {
			return fmt.Errorf("%w: this operator has at leat one operand that cannot reach all types", ErrInvalidModel)
		}
	}

	return nil
}

func assignWeightsToLoopEdges(curNode *WeightedAuthorizationModelNode, edgesArray []*WeightedAuthorizationModelEdge) {
	if !curNode.isNested {
		return
	}
	for _, edge := range edgesArray {
		if edge.isNested {
			for k, v := range curNode.weights {
				edge.weights[k] = v
			}
		}
	}
}

func assignWeightsToNode(node *WeightedAuthorizationModelNode, neighborEdges []*WeightedAuthorizationModelEdge) {
	for _, edge := range neighborEdges {
		for k, v := range edge.weights {
			node.weights[k] = max(node.weights[k], v)
			if node.isNested {
				node.weights[k] = math.MaxInt
			}
		}
	}
}

func assignWeightsToEdge(edge *WeightedAuthorizationModelEdge) error {
	neighborNode, ok := edge.To().(*WeightedAuthorizationModelNode)
	if !ok {
		return fmt.Errorf("%w: could not cast to WeightedAuthorizationModelNode", ErrBuildingGraph)
	}

	switch edge.AuthorizationModelEdge.edgeType {
	case ComputedEdge, RewriteEdge:
		for k, v := range neighborNode.weights {
			edge.weights[k] = v
		}
	case DirectEdge, TTUEdge:
		for k, v := range neighborNode.weights {
			if v == math.MaxInt {
				edge.weights[k] = math.MaxInt
			} else {
				edge.weights[k] = v + 1
			}
		}
	}

	return nil
}

// getOutgoingEdges is nothing but a convenience function.
func (wb *WeightedAuthorizationModelGraphBuilder) getOutgoingEdges(node graph.Node) ([]*WeightedAuthorizationModelEdge, error) {
	edgesarray := make([]*WeightedAuthorizationModelEdge, 0)

	neighborNodes := wb.From(node.ID())
	for neighborNodes.Next() {
		neighborNode, ok := neighborNodes.Node().(*WeightedAuthorizationModelNode)
		if !ok {
			return nil, fmt.Errorf("%w: could not cast to WeightedAuthorizationModelNode", ErrBuildingGraph)
		}

		edges := wb.Lines(node.ID(), neighborNode.ID())
		for edges.Next() {
			edge, ok := edges.Line().(*WeightedAuthorizationModelEdge)
			if !ok {
				return nil, fmt.Errorf("%w: could not cast to WeightedAuthorizationModelEdge", ErrBuildingGraph)
			}

			edgesarray = append(edgesarray, edge)
		}
	}

	return edgesarray, nil
}
