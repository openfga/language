package graph

import (
	"fmt"
	"math"
	"strings"

	"gonum.org/v1/gonum/graph"
)

type WeightedAuthorizationModelGraphBuilder struct {
	graph.DirectedMultigraphBuilder
}

func (wb *WeightedAuthorizationModelGraphBuilder) AddEdgeWithWeights(edge *AuthorizationModelEdge) error {
	isLoop := edge.From().ID() == edge.To().ID()

	// create "from" node
	fromNode := wb.Node(edge.From().ID())
	if fromNode == nil {
		from, ok := edge.From().(*AuthorizationModelNode)
		if !ok {
			return fmt.Errorf("%w: could not cast to WeightedAuthorizationModelNode", ErrBuildingGraph)
		}
		fromNode = &WeightedAuthorizationModelNode{from, make(WeightMap), false}
		wb.AddNode(fromNode)
	}

	// create "to" node
	toNode := wb.Node(edge.To().ID())
	if toNode == nil {
		to, ok := edge.To().(*AuthorizationModelNode)
		if !ok {
			return fmt.Errorf("%w: could not cast to WeightedAuthorizationModelNode", ErrBuildingGraph)
		}
		toNode = &WeightedAuthorizationModelNode{to, make(WeightMap), false}
		wb.AddNode(toNode)
	}

	// update "isNested" field for the node
	if isLoop {
		selfNode, ok := fromNode.(*WeightedAuthorizationModelNode)
		if !ok {
			return fmt.Errorf("%w: could not cast to WeightedAuthorizationModelNode", ErrBuildingGraph)
		}
		selfNode.isNested = true
	}

	// add line
	edge.Line = wb.NewLine(fromNode, toNode)
	newWeightedEdge := &WeightedAuthorizationModelEdge{edge, make(map[string]int), isLoop}
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

//nolint:gocognit,cyclop
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

		if neighborNode.AuthorizationModelNode.nodeType == SpecificType && !strings.Contains(neighborNode.AuthorizationModelNode.label, "*") {
			edge.weights[neighborNode.AuthorizationModelNode.label] = 1
		}

		if err := wb.dfsToAssignWeights(neighborNode, seen); err != nil {
			return err
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
	}

	// second, now that all edge weights have been recursively assigned, assign weights to nodes
	for _, edge := range edgesArray {
		for k, v := range edge.weights {
			curNode.weights[k] = max(curNode.weights[k], v)
			if curNode.isNested {
				curNode.weights[k] = math.MaxInt
			}
		}
	}

	// third, update edges that are loops
	if curNode.isNested {
		for _, edge := range edgesArray {
			if edge.isNested {
				for k, v := range curNode.weights {
					edge.weights[k] = v
				}
			}
		}
	}

	// finally, make sure that intersections and exclusions are "correct"
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
