package graph

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"strings"

	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/multi"
)

const Infinite = math.MaxInt32

var ErrModelCycle = errors.New("model cycle")
var ErrInvalidModel = errors.New("invalid model")
var ErrTupleCycle = errors.New("tuple cycle")
var ErrContrainstTupleCycle = fmt.Errorf("%w: operands AND or BUT NOT cannot be involved in a cycle", ErrTupleCycle)

type WeightedAuthorizationModelGraph struct {
	*multi.DirectedGraph
	// edges is a convenience field to provide O(1) access to an edge.
	// The key is the unique label of a node and the value is the list of edges from that node.
	edges map[string][]*WeightedAuthorizationModelEdge
	// nodes is a convenience field to provide O(1) access to a node.
	// The key is the unique label of a node and the value is the Node struct.
	nodes            map[string]*WeightedAuthorizationModelNode
	drawingDirection DrawingDirection
}

func (wg *WeightedAuthorizationModelGraph) GetEdges() map[string][]*WeightedAuthorizationModelEdge {
	return wg.edges
}

func (wg *WeightedAuthorizationModelGraph) GetEdgesFromNode(node *WeightedAuthorizationModelNode) ([]*WeightedAuthorizationModelEdge, bool) {
	v, ok := wg.edges[node.uniqueLabel]
	return v, ok
}

func (wg *WeightedAuthorizationModelGraph) GetNodes() map[string]*WeightedAuthorizationModelNode {
	return wg.nodes
}

func (wg *WeightedAuthorizationModelGraph) GetNodeByLabel(uniqueLabel string) (*WeightedAuthorizationModelNode, bool) {
	v, ok := wg.nodes[uniqueLabel]
	return v, ok
}

func (wg *WeightedAuthorizationModelGraph) GetNodeByID(id int64) (*WeightedAuthorizationModelNode, error) {
	node := wg.Node(id)

	casted, ok := node.(*WeightedAuthorizationModelNode)
	if !ok {
		return nil, fmt.Errorf("%w: could not cast to AuthorizationModelNode", ErrQueryingGraph)
	}

	return casted, nil
}

func NewWeightedAuthorizationModelGraph() *WeightedAuthorizationModelGraph {
	return &WeightedAuthorizationModelGraph{
		DirectedGraph:    multi.NewDirectedGraph(),
		nodes:            make(map[string]*WeightedAuthorizationModelNode),
		edges:            make(map[string][]*WeightedAuthorizationModelEdge),
		drawingDirection: DrawingDirectionCheck,
	}
}

// AddNode adds a node to the graph with optional operationType and weight.
func (wg *WeightedAuthorizationModelGraph) AddNode(uniqueLabel, label string, nodeType NodeType) {
	innerNode := &AuthorizationModelNode{
		Node:        wg.NewNode(),
		label:       label,
		uniqueLabel: uniqueLabel,
		nodeType:    nodeType,
	}
	newNode := NewWeightedAuthorizationModelNode(innerNode)
	wg.DirectedGraph.AddNode(newNode)
	wg.nodes[uniqueLabel] = newNode
}

func (wg *WeightedAuthorizationModelGraph) AddEdge(fromID, toID string, edgeType EdgeType, condition string) {
	fromNode, _ := wg.GetNodeByLabel(fromID)
	toNode, _ := wg.GetNodeByLabel(toID)
	innerEdge := &AuthorizationModelEdge{
		edgeType:      edgeType,
		conditionedOn: condition,
	}

	innerEdge.Line = wg.NewLine(fromNode, toNode)
	newWeightedEdge := NewWeightedAuthorizationModelEdge(innerEdge)
	wg.DirectedGraph.SetLine(newWeightedEdge)
	wg.edges[fromID] = append(wg.edges[fromID], newWeightedEdge)
}

// AssignWeights assigns weights to all the edges and nodes of the graph.
func (wg *WeightedAuthorizationModelGraph) AssignWeights() error {
	visited := make(map[string]bool)
	ancestorPath := make([]*WeightedAuthorizationModelEdge, 0)
	tupleCycleDependencies := make(map[string][]*WeightedAuthorizationModelEdge)

	for node := range wg.nodes {
		if visited[node] {
			continue
		}

		tupleCyles, err := wg.calculateNodeWeight(node, visited, ancestorPath, tupleCycleDependencies)
		if err != nil {
			return err
		}
		if len(tupleCyles) > 0 {
			return fmt.Errorf("%w: %d tuple cycles found without resolution", ErrTupleCycle, len(tupleCyles))
		}
	}
	return nil
}

func (wg *WeightedAuthorizationModelGraph) calculateEdgeWildcards(edge *WeightedAuthorizationModelEdge) {
	if len(edge.wildcards) > 0 {
		return
	}
	toNode, _ := wg.GetNodeByID(edge.AuthorizationModelEdge.To().ID())
	nodeWildcards := wg.nodes[toNode.uniqueLabel].wildcards
	if len(nodeWildcards) == 0 {
		return
	}
	edge.wildcards = append(edge.wildcards, nodeWildcards...)
}

func (wg *WeightedAuthorizationModelGraph) addReferentialWildcardsToEdge(edge *WeightedAuthorizationModelEdge, referentialNodeID string) {
	referentialNode := wg.nodes[referentialNodeID]
	for _, wildcard := range referentialNode.wildcards {
		if !slices.Contains(edge.wildcards, wildcard) {
			edge.wildcards = append(edge.wildcards, wildcard)
		}
	}
}

func (wg *WeightedAuthorizationModelGraph) addReferentialWildcardsToNode(nodeID string, referentialNodeID string) {
	referentialNode := wg.nodes[referentialNodeID]
	node := wg.nodes[nodeID]
	for _, wildcard := range referentialNode.wildcards {
		if !slices.Contains(node.wildcards, wildcard) {
			node.wildcards = append(node.wildcards, wildcard)
		}
	}
}

func (wg *WeightedAuthorizationModelGraph) addEdgeWildcardsToNode(nodeID string, edge *WeightedAuthorizationModelEdge) {
	node := wg.nodes[nodeID]
	if len(edge.wildcards) == 0 {
		return
	}

	for _, wildcard := range edge.wildcards {
		if !slices.Contains(node.wildcards, wildcard) {
			node.wildcards = append(node.wildcards, wildcard)
		}
	}
}

// Calculate the weight of the node based on the weights of the edges that are connected to the node.
func (wg *WeightedAuthorizationModelGraph) calculateNodeWeight(nodeID string, visited map[string]bool, ancestorPath []*WeightedAuthorizationModelEdge, tupleCycleDependencies map[string][]*WeightedAuthorizationModelEdge) ([]string, error) {
	tupleCycles := make([]string, 0)

	if visited[nodeID] {
		return nil, nil
	}

	// if the node is a specific type or a specific type wildcard, we can set the weight to 1 and return
	if wg.nodes[nodeID].nodeType == SpecificType || wg.nodes[nodeID].nodeType == SpecificTypeWildcard {
		return nil, nil
	}

	visited[nodeID] = true

	for _, edge := range wg.edges[nodeID] {
		if len(edge.weights) != 0 {
			continue
		}
		toNode, _ := wg.GetNodeByID(edge.AuthorizationModelEdge.To().ID())
		if toNode.nodeType == SpecificType || toNode.nodeType == SpecificTypeWildcard {
			edge.weights = make(map[string]int)
			uniqueLabel := toNode.uniqueLabel
			if toNode.nodeType == SpecificTypeWildcard {
				uniqueLabel = uniqueLabel[:len(uniqueLabel)-2]
				edge.wildcards = append(edge.wildcards, uniqueLabel)
				wg.addEdgeWildcardsToNode(nodeID, edge)
			}
			edge.weights[uniqueLabel] = 1
			continue
		}

		tcycle, err := wg.calculateEdgeWeight(edge, ancestorPath, visited, tupleCycleDependencies)
		wg.calculateEdgeWildcards(edge)
		wg.addEdgeWildcardsToNode(nodeID, edge)
		if err != nil {
			tupleCycles = append(tupleCycles, tcycle...)
			return tupleCycles, err
		}
		if len(tcycle) > 0 {
			tupleCycles = append(tupleCycles, tcycle...) // verify if does not exist first
		}
	}

	return wg.calculateNodeWeightFromTheEdges(nodeID, tupleCycleDependencies, tupleCycles)
}

// Calculate the weight of the edge based on the type of edge and the weight of the node that is connected to.
func (wg *WeightedAuthorizationModelGraph) calculateEdgeWeight(edge *WeightedAuthorizationModelEdge, ancestorPath []*WeightedAuthorizationModelEdge, visited map[string]bool, tupleCycleDependencies map[string][]*WeightedAuthorizationModelEdge) ([]string, error) {
	fromNode, _ := wg.GetNodeByID(edge.AuthorizationModelEdge.From().ID())
	toNode, _ := wg.GetNodeByID(edge.AuthorizationModelEdge.To().ID())
	// if it is a recursive edge, we need to set the weight to infinite and add the edge to the tuple cycle dependencies
	if fromNode.uniqueLabel == toNode.uniqueLabel {
		edge.weights = make(map[string]int)
		edge.weights["R#"+toNode.uniqueLabel] = Infinite
		tupleCycleDependencies[toNode.uniqueLabel] = append(tupleCycleDependencies[toNode.uniqueLabel], edge)
		return []string{fromNode.uniqueLabel}, nil
	}

	// calculate the weight of the node that is connected to the edge
	ancestorPath = append(ancestorPath, edge)
	tupleCycle, err := wg.calculateNodeWeight(toNode.uniqueLabel, visited, ancestorPath, tupleCycleDependencies)
	if err != nil {
		return tupleCycle, err
	}

	// if the node that is connected to the edge does not have any weight, we need to check if is a tuple cycle or a model cycle
	if len(toNode.weights) == 0 {
		if wg.isTupleCycle(toNode.uniqueLabel, ancestorPath) {
			edge.weights = make(map[string]int)
			edge.weights["R#"+toNode.uniqueLabel] = Infinite
			tupleCycleDependencies[toNode.uniqueLabel] = append(tupleCycleDependencies[toNode.uniqueLabel], edge)
			tupleCycle = append(tupleCycle, toNode.uniqueLabel)
			return tupleCycle, nil
		}
		return tupleCycle, ErrModelCycle
	}

	isTupleCycle := len(tupleCycle) > 0
	if isTupleCycle {
		for _, nodeID := range tupleCycle {
			tupleCycleDependencies[nodeID] = append(tupleCycleDependencies[nodeID], edge)
		}
	}

	weights := make(map[string]int)
	for key, value := range toNode.weights {
		if !isTupleCycle && strings.HasPrefix(key, "R#") {
			nodeDependency := strings.TrimPrefix(key, "R#")
			tupleCycleDependencies[nodeDependency] = append(tupleCycleDependencies[nodeDependency], edge)
			tupleCycle = append(tupleCycle, nodeDependency)
		}
		weights[key] = value
	}
	edge.weights = weights

	if edge.edgeType == TTUEdge || edge.edgeType == DirectEdge {
		for key, value := range edge.weights {
			if edge.weights[key] == Infinite {
				continue
			}
			edge.weights[key] = value + 1
		}
	}
	return tupleCycle, nil
}

// This function is called when the nodeID is already in the visited path and does not have a weight associated to it.
// In this case we need to know if in the ancestor path to the nodeID is there any edge that is a TTU or a userset.
// If exists we can conclude that is a tuple cycle otherwise is a model cycle.
func (wg *WeightedAuthorizationModelGraph) isTupleCycle(nodeID string, ancestorPath []*WeightedAuthorizationModelEdge) bool {
	startTracking := false
	for _, edge := range ancestorPath {
		fromNode, _ := wg.GetNodeByID(edge.AuthorizationModelEdge.From().ID())
		toNode, _ := wg.GetNodeByID(edge.AuthorizationModelEdge.To().ID())
		if !startTracking && fromNode.uniqueLabel == nodeID {
			startTracking = true
		}
		if startTracking {
			if edge.edgeType == TTUEdge || (edge.edgeType == DirectEdge && toNode.nodeType == SpecificTypeAndRelation) {
				return true
			}
		}
	}
	return false
}

// Calculate the node weight base on the weights of all the edges that are connected from the node.
func (wg *WeightedAuthorizationModelGraph) calculateNodeWeightFromTheEdges(nodeID string, tupleCycleDependencies map[string][]*WeightedAuthorizationModelEdge, tupleCycles []string) ([]string, error) {
	var err error
	node := wg.nodes[nodeID]

	// if there is not cycle, we can calculate the weight without taking in consideration the cycle complexity.
	if len(tupleCycles) == 0 {
		// for any node that is not AND or BUT NOT, we can calculate the weight using the max strategy
		if node.nodeType != OperatorNode || node.label == UnionOperator {
			err = wg.calculateNodeWeightWithMaxStrategy(nodeID)
			if err != nil {
				return tupleCycles, err
			}
		} else {
			// for and and but not, we need to enforce the type strategy, meaning all edges require to return the same type
			err = wg.calculateNodeWeightWithEnforceTypeStrategy(nodeID)
			if err != nil {
				return tupleCycles, err
			}
		}
		return tupleCycles, nil
	}

	// recursive case
	if node.nodeType == SpecificTypeAndRelation && wg.isNodeTupleCycleReference(nodeID, tupleCycles) {
		// calculate the weight of the node and fix all the dependencies that are in the tuple cycle.
		err := wg.calculateNodeWeightAndFixDependencies(nodeID, tupleCycleDependencies)
		if err != nil {
			return tupleCycles, err
		}
		tupleCycles = wg.removeNodeFromTupleCycles(nodeID, tupleCycles)
		return tupleCycles, nil
	}

	if node.nodeType != OperatorNode {
		// even when there is a cycle, if the relation is not recursive then we calculate the weight using the max strategy
		err = wg.calculateNodeWeightWithMaxStrategy(nodeID)
		if err != nil {
			return tupleCycles, err
		}
		return tupleCycles, nil
	}

	// if the node is an union operator and there is a cycle
	if node.nodeType == OperatorNode && node.label == UnionOperator {
		// if the node is the reference node of the cycle, recalculate the weight, solve the depencies and remove the node from the tuple cycle
		if wg.isNodeTupleCycleReference(nodeID, tupleCycles) {
			err := wg.calculateNodeWeightAndFixDependencies(nodeID, tupleCycleDependencies)
			if err != nil {
				return tupleCycles, err
			}
			tupleCycles = wg.removeNodeFromTupleCycles(nodeID, tupleCycles)
			return tupleCycles, nil
		}
		// otherwise even when there is a cycle but the reference node is not the tuple cycle, we can calculate the weight
		err = wg.calculateNodeWeightWithMaxStrategy(nodeID)
		if err != nil {
			return tupleCycles, err
		}
		return tupleCycles, nil
	}

	// In the case of interception or exclussion involved in a cycle, is not allowed, so we return an error
	return tupleCycles, ErrContrainstTupleCycle
}

// The max weight strategy is to take all the types for all the edges and get the max value
// if more than one edge have the same type in their weights.
func (wg *WeightedAuthorizationModelGraph) calculateNodeWeightWithMaxStrategy(nodeID string) error {
	node := wg.nodes[nodeID]
	weights := make(map[string]int)
	edges := wg.edges[nodeID]

	if len(edges) == 0 && node.nodeType != SpecificType && node.nodeType != SpecificTypeWildcard {
		return fmt.Errorf("%w: %s node does not have any terminal type to reach to", ErrInvalidModel, node.uniqueLabel)
	}

	for _, edge := range edges {
		for key, value := range edge.weights {
			if _, ok := weights[key]; !ok {
				weights[key] = value
			} else {
				weights[key] = int(math.Max(float64(weights[key]), float64(value)))
			}
		}
	}
	node.weights = weights
	return nil
}

// This strategy is used in AND or BUT NOT operations, and enforces that all the edges return the same type
// if an edge does not return the same type, the key is removed from the weights.
// While doing that process we does not have any type to get in the weight it means that for this operation
// not all paths return the same type and the model is not valid.
func (wg *WeightedAuthorizationModelGraph) calculateNodeWeightWithEnforceTypeStrategy(nodeID string) error {
	node := wg.nodes[nodeID]
	weights := make(map[string]int)
	edges := wg.edges[nodeID]

	if len(edges) == 0 && node.nodeType != SpecificType && node.nodeType != SpecificTypeWildcard {
		return fmt.Errorf("%w: %s node does not have any terminal type to reach to", ErrInvalidModel, node.uniqueLabel)
	}

	for _, edge := range edges {
		// the first time, take the weights of the edge
		if len(weights) == 0 {
			for key, value := range edge.weights {
				weights[key] = value
			}
			continue
		}

		// for AndOperation or ButnotOperation, remove the key if it is not in the edge, not all edges return the same type
		for key := range weights {
			if value, ok := edge.weights[key]; !ok {
				delete(weights, key)
			} else {
				weights[key] = int(math.Max(float64(weights[key]), float64(value)))
			}
		}
	}
	if len(weights) == 0 {
		return fmt.Errorf("%w: not all paths return the same type for the node %s", ErrInvalidModel, nodeID)
	}
	node.weights = weights
	return nil
}

// This is a comodity function to check if the node is the root of any tuple cycle,
// meaning in the dependencies list there is a reference to this node.
// Finding if a node is the root of a tuple cycle allows to fix the problem in the node and all its dependencies.
func (wg *WeightedAuthorizationModelGraph) isNodeTupleCycleReference(nodeID string, tupleCycles []string) bool {
	for _, tupleCycle := range tupleCycles {
		if tupleCycle == nodeID {
			return true
		}
	}
	return false
}

// This function will calculate the weight of the node by eliminating the reference node of itself and
// fixing the dependencies on the edges and the nodes that are part of the tuple cycle.
// Once all the dependencies are fixed, the node is removed from the tuple cycle list.
func (wg *WeightedAuthorizationModelGraph) calculateNodeWeightAndFixDependencies(nodeID string, tupleCycleDependencies map[string][]*WeightedAuthorizationModelEdge) error {
	node := wg.nodes[nodeID]
	weights := make(map[string]int)
	referenceNodeID := "R#" + nodeID
	edges := wg.edges[nodeID]

	if (node.nodeType == OperatorNode && node.label != UnionOperator) || (node.nodeType != SpecificTypeAndRelation && node.nodeType != OperatorNode) {
		return fmt.Errorf("%w: invalid node, reference node is not a union operator or a relation: %s", ErrTupleCycle, nodeID)
	}

	if len(edges) == 0 && node.nodeType != SpecificType && node.nodeType != SpecificTypeWildcard {
		return fmt.Errorf("%w: %s node does not have any terminal type to reach to", ErrInvalidModel, node.uniqueLabel)
	}

	references := make([]string, 0)
	for _, edge := range edges {
		for key := range edge.weights {
			if key == referenceNodeID {
				continue
			}
			if strings.HasPrefix(key, "R#") {
				references = append(references, key)
			}
			weights[key] = Infinite
		}
	}
	node.weights = weights

	wg.fixDependantEdgesWeight(nodeID, referenceNodeID, references, tupleCycleDependencies)
	wg.fixDependantNodesWeight(nodeID, referenceNodeID, tupleCycleDependencies)
	delete(tupleCycleDependencies, nodeID)
	return nil
}

// This function will fix the weight of the edges that are dependent on the reference node
// We can take this max weight strategy to remove the dependecies in the edges
// because AND or a BUT NOT are not allowed in a tuple cycle.
func (wg *WeightedAuthorizationModelGraph) fixDependantEdgesWeight(nodeCycle string, referenceNodeID string, references []string, tupleCycleDependencies map[string][]*WeightedAuthorizationModelEdge) {
	node := wg.nodes[nodeCycle]

	// for each edge recorded to be dependent on the reference node, we need to update the weight
	for _, edge := range tupleCycleDependencies[nodeCycle] {
		edgeWeights := make(map[string]int)
		// for each weight in the edge, we need to update the weight
		for key1, value1 := range edge.weights {
			// when the key in the weight slice is the reference node, we substitute that weight with the weight of the reference node
			if key1 == referenceNodeID {
				for key2, value2 := range node.weights {
					// if the key does not exist in the edge, we add it
					if _, ok := edgeWeights[key2]; !ok {
						edgeWeights[key2] = value2
						// in case that there is more than one tuple cycle, and now this edge will be dependant on resolving another cycle,
						// add the edge to the dependencies for the new reference node
						if len(references) > 0 && strings.HasPrefix(key2, "R#") {
							nodeDependency := strings.TrimPrefix(key2, "R#")
							tupleCycleDependencies[nodeDependency] = append(tupleCycleDependencies[nodeDependency], edge)
						}
					} else {
						// if the key already exists, we take the max value
						edgeWeights[key2] = int(math.Max(float64(edgeWeights[key2]), float64(value2)))
					}
				}
			} else {
				if _, ok := edgeWeights[key1]; !ok {
					edgeWeights[key1] = value1
				} else {
					edgeWeights[key1] = int(math.Max(float64(edgeWeights[key1]), float64(value1)))
				}
			}
		}
		edge.weights = edgeWeights
		wg.addReferentialWildcardsToEdge(edge, nodeCycle)
	}
}

// This function will fix the weight of the nodes that are dependent on the reference node
// We can take this max weight strategy to remove the dependecies in the nodes
// because AND or a BUT NOT are not allowed in a tuple cycle.
func (wg *WeightedAuthorizationModelGraph) fixDependantNodesWeight(nodeCycle string, referenceNodeID string, tupleCycleDependencies map[string][]*WeightedAuthorizationModelEdge) {
	node := wg.nodes[nodeCycle]

	for _, edge := range tupleCycleDependencies[nodeCycle] {
		fromNode, _ := wg.GetNodeByID(edge.AuthorizationModelEdge.From().ID())
		nodeWeights := make(map[string]int)
		for key1, value1 := range fromNode.weights {
			if key1 == referenceNodeID {
				for key2, value2 := range node.weights {
					if _, ok := nodeWeights[key2]; !ok {
						nodeWeights[key2] = value2
					} else {
						nodeWeights[key2] = int(math.Max(float64(nodeWeights[key2]), float64(value2)))
					}
				}
			} else {
				if _, ok := nodeWeights[key1]; !ok {
					nodeWeights[key1] = value1
				} else {
					nodeWeights[key1] = int(math.Max(float64(nodeWeights[key1]), float64(value1)))
				}
			}
		}
		fromNode.weights = nodeWeights
		wg.addReferentialWildcardsToNode(fromNode.uniqueLabel, nodeCycle)
	}
}

// This function will remove the node from the tuple cycle list.
func (wg *WeightedAuthorizationModelGraph) removeNodeFromTupleCycles(nodeID string, tupleCycles []string) []string {
	result := make([]string, 0)
	for _, tupleCycle := range tupleCycles {
		if tupleCycle != nodeID {
			result = append(result, tupleCycle)
		}
	}
	return result
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

// GetDOT returns the DOT visualization. It should only be used for debugging.
func (wg *WeightedAuthorizationModelGraph) GetDOT() string {
	dotRepresentation, err := dot.MarshalMulti(wg, "", "", "")
	if err != nil {
		return ""
	}

	return string(dotRepresentation)
}
