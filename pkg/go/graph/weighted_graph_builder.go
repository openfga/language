package graph

import (
	"fmt"
	"math"
	"sort"
	"strconv"
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
		formatV := strconv.Itoa(wt[k])
		if wt[k] == math.MaxInt {
			formatV = "+∞"
		}
		sb.WriteString(fmt.Sprintf("%v=%s,", k, formatV))
	}
	formattedWeights := sb.String()
	if len(formattedWeights) > 0 {
		formattedWeights = formattedWeights[:len(formattedWeights)-1]
	}

	return fmt.Sprintf("weights:[%v]", formattedWeights)
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

var _ graph.Node = (*WeightedAuthorizationModelNode)(nil)

type WeightedAuthorizationModelNode struct {
	Inner   *AuthorizationModelNode
	weights WeightToLeafs
}

func (wn *WeightedAuthorizationModelNode) ID() int64 {
	return wn.Inner.ID()
}

func (wn *WeightedAuthorizationModelNode) String() string {
	return fmt.Sprintf("%s [%v]", wn.Inner.uniqueLabel, wn.weights)
}

var _ encoding.Attributer = (*WeightedAuthorizationModelNode)(nil)

func (wn *WeightedAuthorizationModelNode) Attributes() []encoding.Attribute {
	weightsStr := wn.weights
	labelSet := false
	attrs := make([]encoding.Attribute, 0, len(wn.Inner.Attributes()))
	for _, attr := range wn.Inner.Attributes() {
		if attr.Key == "label" {
			labelSet = true
			if len(weightsStr) > 0 {
				attr.Value += fmt.Sprintf(" - %v", weightsStr)
			}
		}
		attrs = append(attrs, attr)
	}

	if !labelSet {
		attrs = append(attrs, encoding.Attribute{
			Key:   "label",
			Value: fmt.Sprintf("%v", weightsStr),
		})
	}

	return attrs
}

type WeightedAuthorizationModelEdge struct {
	Inner      *AuthorizationModelEdge
	weights    WeightToLeafs
	isSelfLoop bool
}

func (wn *WeightedAuthorizationModelEdge) From() graph.Node {
	return wn.Inner.From()
}

func (wn *WeightedAuthorizationModelEdge) To() graph.Node {
	return wn.Inner.To()
}

func (wn *WeightedAuthorizationModelEdge) ReversedLine() graph.Line {
	panic("implement me")
}

func (wn *WeightedAuthorizationModelEdge) ID() int64 {
	return wn.Inner.ID()
}

func (wn *WeightedAuthorizationModelEdge) String() string {
	return fmt.Sprintf("%d -> %d [%v]", wn.Inner.From().ID(), wn.Inner.To().ID(), wn.weights)
}

var (
	_ encoding.Attributer = (*WeightedAuthorizationModelEdge)(nil)
	_ graph.Line          = (*WeightedAuthorizationModelEdge)(nil)
)

func (wn *WeightedAuthorizationModelEdge) Attributes() []encoding.Attribute {
	weightsStr := wn.weights
	labelSet := false
	attrs := make([]encoding.Attribute, 0, len(wn.Inner.Attributes()))
	for _, attr := range wn.Inner.Attributes() {
		if attr.Key == "label" {
			labelSet = true
			if len(weightsStr) > 0 {
				attr.Value += fmt.Sprintf(" - %v", weightsStr)
			}
		}
		attrs = append(attrs, attr)
	}

	if !labelSet {
		attrs = append(attrs, encoding.Attribute{
			Key:   "label",
			Value: fmt.Sprintf("%v", weightsStr),
		})
	}

	return attrs
}

func (wb *WeightedAuthorizationModelGraphBuilder) AddEdgeWithWeights(edge *AuthorizationModelEdge) error {
	fromNode := wb.Node(edge.From().ID())
	if fromNode == nil {
		from, ok := edge.From().(*AuthorizationModelNode)
		if !ok {
			return fmt.Errorf("%w: could not cast to WeightedAuthorizationModelNode", ErrBuildingGraph)
		}
		fromNode = &WeightedAuthorizationModelNode{from, make(WeightToLeafs)}
		// if it's a self-edge, have both ends refer to the same object
		wb.AddNode(fromNode)
	}

	toNode := wb.Node(edge.To().ID())
	if toNode == nil {
		to, ok := edge.To().(*AuthorizationModelNode)
		if !ok {
			return fmt.Errorf("%w: could not cast to WeightedAuthorizationModelNode", ErrBuildingGraph)
		}
		toNode = &WeightedAuthorizationModelNode{to, make(WeightToLeafs)}
		wb.AddNode(toNode)
	}
	// Rewrite Line so that when we do WeightedAuthorizationModelEdge.From, it returns a WeightedAuthorizationModelNode
	// instead of an AuthorizationModelNode.
	edge.Line = wb.NewLine(fromNode, toNode)

	isSelfEdge := edge.From().ID() == edge.To().ID()
	newWeightedEdge := &WeightedAuthorizationModelEdge{edge, make(map[string]int), isSelfEdge}

	wb.DirectedMultigraphBuilder.SetLine(newWeightedEdge)

	return nil
}

func NewWeightedAuthorizationModelGraph(model *openfgav1.AuthorizationModel) (*WeightedAuthorizationModelGraph, error) {
	g, err := NewAuthorizationModelGraph(model)
	if err != nil {
		return nil, err
	}
	g, err = g.Reversed() // we want direction of Check
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

func (wb *WeightedAuthorizationModelGraph) AssignWeights() error {
	seen := make(map[int64]struct{})
	iterNodes := wb.Nodes()
	for iterNodes.Next() {
		nextNode, ok := iterNodes.Node().(*WeightedAuthorizationModelNode)
		if !ok {
			return fmt.Errorf("%w: could not cast to WeightedAuthorizationModelNode", ErrBuildingGraph)
		}

		err := wb.dfs(nextNode, seen)
		if err != nil {
			return err
		}
	}

	//err := wb.ReassignWeightsForNodesWithSelfLoopsToInfinity()
	//if err != nil {
	//	return err
	//}

	return nil
}

//nolint:gocognit,cyclop
func (wb *WeightedAuthorizationModelGraph) dfs(curNode *WeightedAuthorizationModelNode, seen map[int64]struct{}) error {
	if _, seeen := seen[curNode.ID()]; seeen {
		return nil
	}
	seen[curNode.ID()] = struct{}{}

	edgesArray, err := wb.getNeighboringEdges(curNode)
	if err != nil {
		return err
	}

	for _, edge := range edgesArray {
		if len(edge.weights) > 0 || edge.isSelfLoop {
			continue
		}

		neighborNode, ok := edge.To().(*WeightedAuthorizationModelNode)
		if !ok {
			return fmt.Errorf("%w: could not cast to WeightedAuthorizationModelNode", ErrBuildingGraph)
		}

		if neighborNode.Inner.nodeType == SpecificType && !strings.Contains(neighborNode.Inner.label, "*") {
			edge.weights[neighborNode.Inner.label] = 1
		}

		if err := wb.dfs(neighborNode, seen); err != nil {
			return err
		}

		switch edge.Inner.edgeType {
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

		fmt.Println("updated edge weights", edge)
	}

	if len(edgesArray) > 0 {
		// copy the values of the weights from the first edge to the node
		for k, v := range edgesArray[0].weights {
			curNode.weights[k] = v
		}

		// now process the rest of the edges
		for i := 1; i < len(edgesArray); i++ {
			curEdge := edgesArray[i]
			if curEdge.isSelfLoop {
				// special case for self loops
				for k := range curNode.weights {
					curNode.weights[k] = math.MaxInt
					curEdge.weights[k] = math.MaxInt
				}
			} else {
				// pass the weights from the edges to the node
				for k, v := range curEdge.weights {
					if _, ok := curNode.weights[k]; !ok {
						if curNode.Inner.nodeType == OperatorNode && (curNode.Inner.label == "exclusion" || curNode.Inner.label == "intersection") {
							return fmt.Errorf("%w: operator node should have one and only one weight", ErrInvalidModel)
						}
					}
					curNode.weights[k] = max(curNode.weights[k], v)
				}
			}

		}
	}

	fmt.Println("updated node", curNode)

	return nil
}

func (wb *WeightedAuthorizationModelGraph) getNeighboringEdges(node *WeightedAuthorizationModelNode) ([]*WeightedAuthorizationModelEdge, error) {
	edgesarray := make([]*WeightedAuthorizationModelEdge, 0)
	selfEdgesArray := make([]*WeightedAuthorizationModelEdge, 0)

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

			if edge.isSelfLoop {
				selfEdgesArray = append(selfEdgesArray, edge)
			} else {
				edgesarray = append(edgesarray, edge)
			}
		}
	}

	edgesarray = append(edgesarray, selfEdgesArray...)

	return edgesarray, nil
}
