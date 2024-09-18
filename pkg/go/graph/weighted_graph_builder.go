package graph

import (
	"container/list"
	"fmt"
	"sort"
	"strings"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/multi"
)

// WeightToLeafs is a map of where the key is a type (e.g. folder, user) and the value is the weight or complexity of the relation.
type WeightToLeafs map[string]int

func (wt WeightToLeafs) String() string {
	if len(wt) == 0 {
		return ""
	}
	var sb strings.Builder

	// Extract keys and sort them
	keys := make([]string, 0, len(wt))
	for k := range wt {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		sb.WriteString(fmt.Sprintf("%v=%v,", k, wt[k]))
	}
	formattedWeights := sb.String()
	if len(formattedWeights) > 0 {
		formattedWeights = formattedWeights[:len(formattedWeights)-1]
	}

	return formattedWeights
}

type WeightedAuthorizationModelGraph struct {
	*multi.DirectedGraph
	drawingDirection DrawingDirection
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

type WeightedAuthorizationModelGraphBuilder struct {
	graph.DirectedMultigraphBuilder
}

type WeightedAuthorizationModelNode struct {
	*AuthorizationModelNode
	weights WeightToLeafs
}

var _ encoding.Attributer = (*WeightedAuthorizationModelNode)(nil)

func (wn *WeightedAuthorizationModelNode) Attributes() []encoding.Attribute {
	attrs := make([]encoding.Attribute, 0, len(wn.AuthorizationModelNode.Attributes()))
	for _, attr := range wn.AuthorizationModelNode.Attributes() {
		if attr.Key == "label" {
			weightsStr := wn.weights.String()
			if len(weightsStr) > 0 {
				attr.Value += fmt.Sprintf(" - weights:[%v]", weightsStr)
			}
		}
		attrs = append(attrs, attr)
	}

	return attrs
}

type WeightedAuthorizationModelEdge struct {
	*AuthorizationModelEdge
}

var _ encoding.Attributer = (*WeightedAuthorizationModelEdge)(nil)

func (we *WeightedAuthorizationModelEdge) Attributes() []encoding.Attribute {
	return we.AuthorizationModelEdge.Attributes()
}

func (wb *WeightedAuthorizationModelGraphBuilder) AddEdgeWithWeights(edge *AuthorizationModelEdge) error {
	from, ok := edge.From().(*AuthorizationModelNode)
	if !ok {
		return fmt.Errorf("%w: could not cast to WeightedAuthorizationModelNode", ErrBuildingGraph)
	}
	weightedFrom := &WeightedAuthorizationModelNode{from, make(map[string]int)}
	to, ok := edge.To().(*AuthorizationModelNode)
	if !ok {
		return fmt.Errorf("%w: could not cast to WeightedAuthorizationModelNode", ErrBuildingGraph)
	}
	weightedTo := &WeightedAuthorizationModelNode{to, make(map[string]int)}

	// Rewrite Line so that when we do WeightedAuthorizationModelEdge.From, it returns a WeightedAuthorizationModelNode
	// instead of an AuthorizationModelNode.
	edge.Line = wb.NewLine(weightedFrom, weightedTo)
	newWeightedEdge := &WeightedAuthorizationModelEdge{edge}
	wb.DirectedMultigraphBuilder.SetLine(newWeightedEdge)

	return nil
}

func NewWeightedAuthorizationModelGraph(model *openfgav1.AuthorizationModel) (*WeightedAuthorizationModelGraph, error) {
	g, err := NewAuthorizationModelGraph(model)
	if err != nil {
		return nil, err
	}

	graphBuilder := &WeightedAuthorizationModelGraphBuilder{multi.NewDirectedGraph()}

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

			err = graphBuilder.AddEdgeWithWeights(castedEdge)
			if err != nil {
				return nil, err
			}
		}
	}

	multigraph, ok := graphBuilder.DirectedMultigraphBuilder.(*multi.DirectedGraph)
	if !ok {
		return nil, fmt.Errorf("%w: could not cast to DirectedGraph", ErrBuildingGraph)
	}

	wg := &WeightedAuthorizationModelGraph{multigraph, g.drawingDirection}

	err = wg.AssignWeights()
	if err != nil {
		return nil, err
	}

	return wg, nil
}

// dfsState's goal is to avoid visiting the same edge more than once (it's okay to visit the same node more than once).
type dfsState struct {
	fromNodeID int64
	toNodeID   int64
	lineID     int64
}

//nolint:gocognit,cyclop
func (wb *WeightedAuthorizationModelGraph) AssignWeights() error {
	// DFS traversal
	stack := list.New()
	seen := make(map[dfsState]struct{})

	startFrom, err := wb.GetNodesWithNoIncomingEdges()
	if err != nil {
		return err
	}

	for node := range startFrom {
		stack.PushBack(node)
		seen[dfsState{fromNodeID: node.ID(), lineID: -1, toNodeID: -1}] = struct{}{}
	}

	for stack.Len() > 0 {
		curNode, ok := stack.Remove(stack.Back()).(*WeightedAuthorizationModelNode)
		if !ok {
			return fmt.Errorf("%w: could not cast to WeightedAuthorizationModelNode", ErrBuildingGraph)
		}

		neighbors := wb.From(curNode.ID())

		for neighbors.Next() {
			neighborNode, ok := neighbors.Node().(*WeightedAuthorizationModelNode)
			if !ok {
				return fmt.Errorf("%w: could not cast to WeightedAuthorizationModelNode", ErrBuildingGraph)
			}

			// There could be more than one edge between two nodes.
			edges := wb.Lines(curNode.ID(), neighborNode.ID())
			for edges.Next() {
				edge, ok := edges.Line().(*WeightedAuthorizationModelEdge)
				if !ok {
					return fmt.Errorf("%w: could not cast to WeightedAuthorizationModelEdge", ErrBuildingGraph)
				}

				cur := dfsState{fromNodeID: curNode.ID(), toNodeID: neighborNode.ID(), lineID: edge.ID()}

				if _, seeen := seen[cur]; !seeen {
					// Update weights
					if curNode.nodeType == SpecificType {
						neighborNode.weights[curNode.label] = 1
					}
					if neighborNode.nodeType == SpecificTypeAndRelation && edge.edgeType == ComputedEdge {
						neighborNode.weights = curNode.weights
					}

					// Continue DFS
					seen[cur] = struct{}{}
					stack.PushBack(neighborNode)
				}
			}
		}
	}

	return nil
}

func (wb *WeightedAuthorizationModelGraph) GetNodesWithNoIncomingEdges() (map[*WeightedAuthorizationModelNode]struct{}, error) {
	if wb.drawingDirection != DrawingDirectionListObjects {
		return nil, fmt.Errorf("%w: incorrect drawing direction: %v", ErrBuildingGraph, wb.drawingDirection)
	}
	leafs := make(map[*WeightedAuthorizationModelNode]struct{})
	iterNodes := wb.Nodes()
	for iterNodes.Next() {
		nextNode := iterNodes.Node()
		casted, ok := nextNode.(*WeightedAuthorizationModelNode)
		if !ok {
			return nil, fmt.Errorf("%w: could not cast to WeightedAuthorizationModelNode", ErrBuildingGraph)
		}

		if wb.To(casted.ID()).Len() == 0 {
			leafs[casted] = struct{}{}
		}
	}

	return leafs, nil
}
