package graph

import (
	"fmt"

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

	outgoingEdgesOfNode, err := wb.getOutgoingEdges(curNode)
	if err != nil {
		return err
	}

	// first, assign weights to neighbor edges
	for _, edge := range outgoingEdgesOfNode {
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

		if err := edge.assignWeightsToEdge(); err != nil {
			return err
		}
	}

	// second, now that all edge weights have been recursively assigned, assign weights to node
	curNode.assignWeightsToNode(outgoingEdgesOfNode)

	// third, update edges that are loops
	assignWeightsToLoopEdges(curNode, outgoingEdgesOfNode)

	// finally, make sure that intersections and exclusions are "correct"
	return curNode.verifyNodeIsValid(outgoingEdgesOfNode)
}

func assignWeightsToLoopEdges(curNode *WeightedAuthorizationModelNode, outgoingEdges []*WeightedAuthorizationModelEdge) {
	if !curNode.isNested {
		return
	}
	for _, edge := range outgoingEdges {
		if edge.isNested {
			for k, v := range curNode.weights {
				edge.weights[k] = v
			}
		}
	}
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
