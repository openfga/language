package graph

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/graph"
)

type WeightedAuthorizationModelGraphBuilder struct {
	graph.DirectedMultigraphBuilder
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

	edgesArray, err := wb.getNeighboringEdges(curNode)
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

		if err := wb.dfsToAssignWeights(neighborNode, seen); err != nil {
			return err
		}

		wb.assignWeightsToEdge(edge, neighborNode)
	}

	// second, now that all edge weights have been recursively assigned, assign weights to node
	wb.assignWeightsToNode(curNode, edgesArray)

	// third, update edges that are loops
	wb.assignWeightsToLoopEdges(curNode, edgesArray)

	// finally, make sure that intersections and exclusions are "correct"
	return wb.verifyOperators(curNode, edgesArray)
}

func (wb *WeightedAuthorizationModelGraphBuilder) verifyOperators(curNode *WeightedAuthorizationModelNode, edgesArray []*WeightedAuthorizationModelEdge) error {
	if curNode.nodeType == OperatorNode && !(curNode.label == "union") {
		edgeWeights := make([]WeightMap, len(edgesArray))
		for i := 0; i < len(edgesArray); i++ {
			edgeWeights[i] = edgesArray[i].weights
		}
		intersect, err := Intersection(edgeWeights...)
		if err != nil {
			return err
		}

		if len(intersect) == 0 {
			return fmt.Errorf("%w: this operator cannot reach any type", ErrInvalidModel)
		}
	}

	return nil
}

func (wb *WeightedAuthorizationModelGraphBuilder) assignWeightsToLoopEdges(curNode *WeightedAuthorizationModelNode, edgesArray []*WeightedAuthorizationModelEdge) {
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

func (wb *WeightedAuthorizationModelGraphBuilder) assignWeightsToNode(node *WeightedAuthorizationModelNode, neighborEdges []*WeightedAuthorizationModelEdge) {
	for _, edge := range neighborEdges {
		for k, v := range edge.weights {
			node.weights[k] = max(node.weights[k], v)
			if node.isNested {
				node.weights[k] = math.MaxInt
			}
		}
	}
}

func (wb *WeightedAuthorizationModelGraphBuilder) assignWeightsToEdge(edge *WeightedAuthorizationModelEdge, neighborNode *WeightedAuthorizationModelNode) {
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
}

// getNeighboringEdges is nothing but a convenience function.
func (wb *WeightedAuthorizationModelGraphBuilder) getNeighboringEdges(node *WeightedAuthorizationModelNode) ([]*WeightedAuthorizationModelEdge, error) {
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
