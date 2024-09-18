package graph

import (
	"fmt"
	"math"
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
	weights WeightToLeafs
}

var _ encoding.Attributer = (*WeightedAuthorizationModelEdge)(nil)

func (wn *WeightedAuthorizationModelEdge) Attributes() []encoding.Attribute {
	attrs := make([]encoding.Attribute, 0, len(wn.AuthorizationModelEdge.Attributes()))
	for _, attr := range wn.AuthorizationModelEdge.Attributes() {
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

func (wb *WeightedAuthorizationModelGraphBuilder) AddEdgeWithWeights(edge *AuthorizationModelEdge) error {
	from, ok := edge.From().(*AuthorizationModelNode)
	if !ok {
		return fmt.Errorf("%w: could not cast to WeightedAuthorizationModelNode", ErrBuildingGraph)
	}
	weightedFrom := &WeightedAuthorizationModelNode{from, make(WeightToLeafs)}
	to, ok := edge.To().(*AuthorizationModelNode)
	if !ok {
		return fmt.Errorf("%w: could not cast to WeightedAuthorizationModelNode", ErrBuildingGraph)
	}
	weightedTo := &WeightedAuthorizationModelNode{to, make(WeightToLeafs)}

	// Rewrite Line so that when we do WeightedAuthorizationModelEdge.From, it returns a WeightedAuthorizationModelNode
	// instead of an AuthorizationModelNode.
	edge.Line = wb.NewLine(weightedFrom, weightedTo)
	newWeightedEdge := &WeightedAuthorizationModelEdge{edge, make(WeightToLeafs)}
	wb.DirectedMultigraphBuilder.SetLine(newWeightedEdge)

	return nil
}

func NewWeightedAuthorizationModelGraph(model *openfgav1.AuthorizationModel) (*WeightedAuthorizationModelGraph, error) {
	g, err := NewAuthorizationModelGraph(model)
	if err != nil {
		return nil, err
	}
	g, err = g.Reversed()
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
	// DFS traversal
	seen := make(map[*WeightedAuthorizationModelNode]struct{})

	startFrom, err := wb.GetNodesWithNoIncomingEdges()
	if err != nil {
		return err
	}

	for node := range startFrom {
		seen[node] = struct{}{}
		err := wb.dfs(node, seen)
		if err != nil {
			return err
		}
	}

	return nil
}

//nolint:gocognit,cyclop
func (wb *WeightedAuthorizationModelGraph) dfs(curNode *WeightedAuthorizationModelNode, seen map[*WeightedAuthorizationModelNode]struct{}) error {
	// if _, seeen := seen[curNode]; seeen {
	//	// TODO
	//}
	seen[curNode] = struct{}{}

	neighborNodes := wb.From(curNode.ID())

	for neighborNodes.Next() {
		neighborNode, ok := neighborNodes.Node().(*WeightedAuthorizationModelNode)
		if !ok {
			return fmt.Errorf("%w: could not cast to WeightedAuthorizationModelNode", ErrBuildingGraph)
		}
		edges := wb.Lines(curNode.ID(), neighborNode.ID())
		for edges.Next() {
			edge, ok := edges.Line().(*WeightedAuthorizationModelEdge)
			if !ok {
				return fmt.Errorf("%w: could not cast to WeightedAuthorizationModelEdge", ErrBuildingGraph)
			}

			if len(edge.weights) > 0 {
				continue
			}
			if neighborNode.nodeType == SpecificType && !strings.Contains(neighborNode.label, "*") {
				edge.weights[neighborNode.label] = 1
				// continue
			}
			if err := wb.dfs(neighborNode, seen); err != nil {
				return err
			}
			if curNode == neighborNode {
				edge.weights[curNode.label] = math.MaxInt
				// continue
			}

			switch edge.edgeType {
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

			for k, v := range edge.weights {
				curNode.weights[k] = max(curNode.weights[k], v)
			}

			if curNode.nodeType == OperatorNode && (curNode.label == "union" || curNode.label == "intersection") && len(curNode.weights) > 1 {
				return fmt.Errorf("%w: operator node should have one and only one weight", ErrInvalidModel)
			}
		}
	}

	return nil
}

func (wb *WeightedAuthorizationModelGraph) GetNodesWithNoIncomingEdges() (map[*WeightedAuthorizationModelNode]struct{}, error) {
	if wb.drawingDirection != DrawingDirectionCheck {
		return nil, fmt.Errorf("%w: incorrect drawing direction: %v", ErrBuildingGraph, wb.drawingDirection)
	}
	res := make(map[*WeightedAuthorizationModelNode]struct{})
	iterNodes := wb.Nodes()
	for iterNodes.Next() {
		nextNode := iterNodes.Node()
		casted, ok := nextNode.(*WeightedAuthorizationModelNode)
		if !ok {
			return nil, fmt.Errorf("%w: could not cast to WeightedAuthorizationModelNode", ErrBuildingGraph)
		}

		if wb.To(casted.ID()).Len() == 0 {
			res[casted] = struct{}{}
		}
	}

	return res, nil
}
