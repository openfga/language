package graph

import (
	"fmt"
	"math"
	"strings"
)

const Infinite = math.MaxInt32

type WeightedAuthorizationModelGraph struct {
	nodes map[string]*WeightedAuthorizationModelNode
	edges map[string][]*WeightedAuthorizationModelEdge
}

// NewWeightedAuthorizationModelGraph creates a new WeightedAuthorizationModelGraph.
func NewWeightedAuthorizationModelGraph() *WeightedAuthorizationModelGraph {
	return &WeightedAuthorizationModelGraph{
		nodes: make(map[string]*WeightedAuthorizationModelNode),
		edges: make(map[string][]*WeightedAuthorizationModelEdge),
	}
}

// AddNode adds a node to the graph with optional operationType and weight.
func (wg *WeightedAuthorizationModelGraph) AddNode(uniqueLabel, label string, nodeType NodeType) {
	wg.nodes[uniqueLabel] = &WeightedAuthorizationModelNode{uniqueLabel: uniqueLabel, label: label, nodeType: nodeType}
}

func (wg *WeightedAuthorizationModelGraph) AddEdge(fromID, toID string, edgeType EdgeType, condition string) {
	fromNode := wg.nodes[fromID]
	toNode := wg.nodes[toID]
	edge := &WeightedAuthorizationModelEdge{from: fromNode, to: toNode, edgeType: edgeType, conditionedOn: condition}
	wg.edges[fromID] = append(wg.edges[fromID], edge)
}

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
			for _, tupleCycle := range tupleCyles {
				fmt.Println("Tuple cycle is ", tupleCycle)
			}
			return fmt.Errorf("%d tuple cycles found without resolution", len(tupleCyles))
		}
	}
	return nil
}

func (wg *WeightedAuthorizationModelGraph) calculateNodeWeight(nodeID string, visited map[string]bool, ancestorPath []*WeightedAuthorizationModelEdge, tupleCycleDependencies map[string][]*WeightedAuthorizationModelEdge) ([]string, error) {
	tupleCycles := make([]string, 0)

	if visited[nodeID] {
		return nil, nil
	}

	if wg.nodes[nodeID].nodeType == SpecificType || wg.nodes[nodeID].nodeType == SpecificTypeWildcard {
		return nil, nil
	}

	visited[nodeID] = true

	for _, edge := range wg.edges[nodeID] {
		if len(edge.weights) != 0 {
			continue
		}
		if edge.to.nodeType == SpecificType || edge.to.nodeType == SpecificTypeWildcard {
			edge.weights = make(map[string]int)
			uniqueLabel := edge.to.uniqueLabel
			if edge.to.nodeType == SpecificTypeWildcard {
				uniqueLabel = uniqueLabel[:len(uniqueLabel)-2]
			}
			edge.weights[uniqueLabel] = 1
			continue
		}

		tcycle, err := wg.calculateEdgeWeight(edge, ancestorPath, visited, tupleCycleDependencies)
		if err != nil {
			tupleCycles = append(tupleCycles, tcycle...)
			return tupleCycles, err
		}
		if len(tcycle) > 0 {
			tupleCycles = append(tupleCycles, tcycle...) // verify if does not exist first
		}
	}

	var err error
	tupleCycles, err = wg.calculateNodeWeightFromTheEdges(nodeID, tupleCycleDependencies, tupleCycles)
	if err != nil {
		return tupleCycles, err
	}
	return tupleCycles, nil
}

func (wg *WeightedAuthorizationModelGraph) calculateEdgeWeight(edge *WeightedAuthorizationModelEdge, ancestorPath []*WeightedAuthorizationModelEdge, visited map[string]bool, tupleCycleDependencies map[string][]*WeightedAuthorizationModelEdge) ([]string, error) {
	if edge.from.uniqueLabel == edge.to.uniqueLabel {
		edge.weights = make(map[string]int)
		edge.weights["R#"+edge.to.uniqueLabel] = Infinite
		tupleCycleDependencies[edge.to.uniqueLabel] = append(tupleCycleDependencies[edge.to.uniqueLabel], edge)
		return []string{edge.from.uniqueLabel}, nil
	}

	ancestorPath = append(ancestorPath, edge)
	tupleCycle, err := wg.calculateNodeWeight(edge.to.uniqueLabel, visited, ancestorPath, tupleCycleDependencies)
	if err != nil {
		return tupleCycle, err
	}

	if len(edge.to.weights) == 0 {
		if wg.isTupleCycle(edge.to.uniqueLabel, ancestorPath) {
			edge.weights = make(map[string]int)
			edge.weights["R#"+edge.to.uniqueLabel] = Infinite
			tupleCycleDependencies[edge.to.uniqueLabel] = append(tupleCycleDependencies[edge.to.uniqueLabel], edge)
			tupleCycle = append(tupleCycle, edge.to.uniqueLabel)
			return tupleCycle, nil
		}
		return tupleCycle, fmt.Errorf("model cycle")
	}

	isTupleCycle := len(tupleCycle) > 0
	if isTupleCycle {
		for _, nodeID := range tupleCycle {
			tupleCycleDependencies[nodeID] = append(tupleCycleDependencies[nodeID], edge)
		}
	}

	weights := make(map[string]int)
	for key, value := range edge.to.weights {
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

// IsNodeVisited checks if a node with the given ID has been visited.
func (wg *WeightedAuthorizationModelGraph) isTupleCycle(nodeID string, ancestorPath []*WeightedAuthorizationModelEdge) bool {
	startTracking := false
	for _, edge := range ancestorPath {
		if !startTracking && edge.from.uniqueLabel == nodeID {
			startTracking = true
		}
		if startTracking {
			if edge.edgeType == TTUEdge || (edge.edgeType == DirectEdge && edge.to.nodeType == SpecificTypeAndRelation) {
				return true
			}
		}
	}
	return false
}

func (wg *WeightedAuthorizationModelGraph) calculateNodeWeightFromTheEdges(nodeID string, tupleCycleDependencies map[string][]*WeightedAuthorizationModelEdge, tupleCycles []string) ([]string, error) {
	var err error
	node := wg.nodes[nodeID]

	// if there is not cycle, we can calculate the weight without any reference
	if len(tupleCycles) == 0 {
		// for any node that is not and or but not, we can calculate the weight using the max strategy
		if node.nodeType != OperatorNode || node.label == UnionOperator {
			wg.calculateNodeWeightWithMaxStrategy(nodeID) // max weight strategy
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
		wg.calculateNodeWeightWithMaxStrategy(nodeID)
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
		wg.calculateNodeWeightWithMaxStrategy(nodeID)
		return tupleCycles, nil
	}

	// In the case of interception or exclussion involved in a cycle, is not allowed, so we return an error
	return tupleCycles, fmt.Errorf("operands AND or BUT NOT cannot be involved in a cycle")
}

func (wg *WeightedAuthorizationModelGraph) calculateNodeWeightWithMaxStrategy(nodeID string) {
	node := wg.nodes[nodeID]
	weights := make(map[string]int)

	for _, edge := range wg.edges[nodeID] {
		// the first time, take the weights of the edge
		if len(weights) == 0 {
			for key, value := range edge.weights {
				weights[key] = value
			}
			continue
		}
		for key, value := range edge.weights {
			if _, ok := weights[key]; !ok {
				weights[key] = value
			} else {
				weights[key] = int(math.Max(float64(weights[key]), float64(value)))
			}
		}
	}
	node.weights = weights
}

func (wg *WeightedAuthorizationModelGraph) calculateNodeWeightWithEnforceTypeStrategy(nodeID string) error {
	node := wg.nodes[nodeID]
	weights := make(map[string]int)

	for _, edge := range wg.edges[nodeID] {
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
		return fmt.Errorf("not all paths return the same type for the node %s", nodeID)
	}
	node.weights = weights
	return nil
}

func (wg *WeightedAuthorizationModelGraph) isNodeTupleCycleReference(nodeID string, tupleCycles []string) bool {
	for _, tupleCycle := range tupleCycles {
		if tupleCycle == nodeID {
			return true
		}
	}
	return false
}

func (wg *WeightedAuthorizationModelGraph) calculateNodeWeightAndFixDependencies(nodeID string, tupleCycleDependencies map[string][]*WeightedAuthorizationModelEdge) error {
	node := wg.nodes[nodeID]
	weights := make(map[string]int)
	referenceNodeID := "R#" + nodeID

	if (node.nodeType == OperatorNode && node.label != UnionOperator) || (node.nodeType != SpecificTypeAndRelation && node.nodeType != OperatorNode) {
		return fmt.Errorf("invalid node, reference node is not a union operator or a relation: %s", nodeID)
	}

	references := make([]string, 0)
	for _, edge := range wg.edges[nodeID] {
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

func (wg *WeightedAuthorizationModelGraph) fixDependantEdgesWeight(nodeCycle string, referenceNodeID string, references []string, tupleCycleDependencies map[string][]*WeightedAuthorizationModelEdge) {
	node := wg.nodes[nodeCycle]

	// We can take this approach because there is not an AND or a BUT NOT in the cycle
	// and all the nodes and edgs part of the cycle can apply the max strategy

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
	}
}

func (wg *WeightedAuthorizationModelGraph) fixDependantNodesWeight(nodeCycle string, referenceNodeID string, tupleCycleDependencies map[string][]*WeightedAuthorizationModelEdge) {
	node := wg.nodes[nodeCycle]

	// We can take this approach because there is not an AND or a BUT NOT in the cycle
	// and all the nodes and edgs part of the cycle can apply the max strategy

	for _, edge := range tupleCycleDependencies[nodeCycle] {
		fromNode := wg.nodes[edge.from.uniqueLabel]
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
	}
}

func (wg *WeightedAuthorizationModelGraph) removeNodeFromTupleCycles(nodeID string, tupleCycles []string) []string {
	result := make([]string, 0)
	for _, tupleCycle := range tupleCycles {
		if tupleCycle != nodeID {
			result = append(result, tupleCycle)
		}
	}
	return result
}
