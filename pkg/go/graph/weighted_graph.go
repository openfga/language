package graph

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"strings"
)

const Infinite = math.MaxInt32

var ErrModelCycle = errors.New("model cycle")
var ErrInvalidModel = errors.New("invalid model")
var ErrTupleCycle = errors.New("tuple cycle")
var ErrContrainstTupleCycle = fmt.Errorf("%w: operands AND or BUT NOT cannot be involved in a cycle", ErrTupleCycle)

type WeightedAuthorizationModelGraph struct {
	edges map[string][]*WeightedAuthorizationModelEdge
	nodes map[string]*WeightedAuthorizationModelNode
}

// GetEdges returns the edges map.
func (wg *WeightedAuthorizationModelGraph) GetEdges() map[string][]*WeightedAuthorizationModelEdge {
	return wg.edges
}

func (wg *WeightedAuthorizationModelGraph) GetEdgesFromNode(node *WeightedAuthorizationModelNode) ([]*WeightedAuthorizationModelEdge, bool) {
	v, ok := wg.edges[node.uniqueLabel]
	return v, ok
}

// GetNodes returns the nodes map.
func (wg *WeightedAuthorizationModelGraph) GetNodes() map[string]*WeightedAuthorizationModelNode {
	return wg.nodes
}

func (wg *WeightedAuthorizationModelGraph) GetNodeByID(uniqueLabel string) (*WeightedAuthorizationModelNode, bool) {
	v, ok := wg.nodes[uniqueLabel]
	return v, ok
}

// NewWeightedAuthorizationModelGraph creates a new WeightedAuthorizationModelGraph.
func NewWeightedAuthorizationModelGraph() *WeightedAuthorizationModelGraph {
	return &WeightedAuthorizationModelGraph{
		nodes: make(map[string]*WeightedAuthorizationModelNode),
		edges: make(map[string][]*WeightedAuthorizationModelEdge),
	}
}

// AddNode adds a node to the graph with optional nodeType and weight.
func (wg *WeightedAuthorizationModelGraph) AddNode(uniqueLabel, label string, nodeType NodeType) {
	var wildcards []string
	if nodeType == SpecificTypeWildcard {
		wildcards = []string{uniqueLabel[:len(uniqueLabel)-2]}
	}
	wg.nodes[uniqueLabel] = &WeightedAuthorizationModelNode{uniqueLabel: uniqueLabel, label: label, nodeType: nodeType, wildcards: wildcards}
}

// AddNode adds a node to the graph with optional nodeType and weight.
func (wg *WeightedAuthorizationModelGraph) GetOrAddNode(uniqueLabel, label string, nodeType NodeType) *WeightedAuthorizationModelNode {
	if existingNode := wg.nodes[uniqueLabel]; existingNode != nil {
		return existingNode
	}
	var wildcards []string
	if nodeType == SpecificTypeWildcard {
		wildcards = []string{uniqueLabel[:len(uniqueLabel)-2]}
	}
	wg.nodes[uniqueLabel] = &WeightedAuthorizationModelNode{uniqueLabel: uniqueLabel, label: label, nodeType: nodeType, wildcards: wildcards}
	return wg.nodes[uniqueLabel]
}

func (wg *WeightedAuthorizationModelGraph) AddEdge(fromID, toID string, edgeType EdgeType, relationDefinition string, tuplesetRelation string, conditions []string) {
	fromNode := wg.nodes[fromID]
	toNode := wg.nodes[toID]
	if len(conditions) == 0 {
		conditions = []string{NoCond}
	}
	edge := &WeightedAuthorizationModelEdge{from: fromNode, to: toNode, edgeType: edgeType, tuplesetRelation: tuplesetRelation, wildcards: nil, conditions: conditions, relationDefinition: relationDefinition}
	wg.edges[fromID] = append(wg.edges[fromID], edge)
}

func (wg *WeightedAuthorizationModelGraph) UpsertEdge(fromNode, toNode *WeightedAuthorizationModelNode, edgeType EdgeType, relationDefinition string, tuplesetRelation string, condition string) error {
	if fromNode == nil || toNode == nil {
		return fmt.Errorf("%w: Model cannot be parsed", ErrInvalidModel)
	}

	if condition == "" {
		condition = NoCond
	}

	edges := wg.edges[fromNode.uniqueLabel]
	for _, edge := range edges {
		if edge.to.uniqueLabel == toNode.uniqueLabel && edge.edgeType == edgeType && edge.tuplesetRelation == tuplesetRelation {
			for _, cond := range edge.conditions {
				if cond == condition {
					return nil
				}
			}
			edge.conditions = append(edge.conditions, condition)
			return nil
		}
	}

	conditions := []string{condition}
	edge := &WeightedAuthorizationModelEdge{from: fromNode, to: toNode, edgeType: edgeType, tuplesetRelation: tuplesetRelation, wildcards: nil, conditions: conditions, relationDefinition: relationDefinition}
	wg.edges[fromNode.uniqueLabel] = append(wg.edges[fromNode.uniqueLabel], edge)
	return nil
}

func (wg *WeightedAuthorizationModelGraph) HasEdge(fromNode, toNode *WeightedAuthorizationModelNode, edgeType EdgeType, tuplesetRelation string) bool {
	if fromNode == nil || toNode == nil {
		return false
	}
	edges := wg.edges[fromNode.uniqueLabel]
	for _, edge := range edges {
		if edge.to.uniqueLabel == toNode.uniqueLabel && edge.edgeType == edgeType && edge.tuplesetRelation == tuplesetRelation {
			return true
		}
	}

	return false
}

// AssignWeights assigns weights to all the edges and nodes of the graph.
func (wg *WeightedAuthorizationModelGraph) AssignWeights() error {
	visited := make(map[string]bool)
	ancestorPath := make([]*WeightedAuthorizationModelEdge, 0)
	tupleCycleDependencies := make(map[string][]*WeightedAuthorizationModelEdge)

	for nodeID, node := range wg.nodes {
		if wg.isLogicalOperator(node) {
			// Inititally defer weight assignment of operator nodes and the logical nodes to later in `fixDependantNodesWeight()`.
			// This enables more deterministic behavior of intermediate functions.
			continue
		}

		if visited[nodeID] {
			continue
		}

		tupleCyles, err := wg.calculateNodeWeight(nodeID, visited, ancestorPath, tupleCycleDependencies)
		if err != nil {
			return err
		}
		if len(tupleCyles) > 0 {
			return fmt.Errorf("%w: %d tuple cycles found without resolution", ErrTupleCycle, len(tupleCyles))
		}
	}
	return nil
}

func (wg *WeightedAuthorizationModelGraph) isLogicalOperator(node *WeightedAuthorizationModelNode) bool {
	// a logical ttu is when a ttu has more than one edges due to having multiple terminal types as usersets,
	// and should always be treated as a union among all those ttu edges as they belong to the same logical ttu
	// a logical userset is when we have multiple asignations to terminal types or usersets, we need to treat that
	// as a union for all the direct edges
	nodeType := node.GetNodeType()
	return nodeType == OperatorNode || nodeType == LogicalTTUGrouping || nodeType == LogicalDirectGrouping
}

func (wg *WeightedAuthorizationModelGraph) isLogicalUnionOperator(node *WeightedAuthorizationModelNode) bool {
	// a logical ttu is when a ttu has more than one edges due to having multiple terminal types as usersets,
	// and should always be treated as a union among all those ttu edges as they belong to the same logical ttu
	// a logical userset is when we have multiple asignations to terminal types or usersets, we need to treat that
	// as a union for all the direct edges
	nodeType := node.GetNodeType()
	return (nodeType == OperatorNode && node.GetLabel() == UnionOperator) || nodeType == LogicalTTUGrouping || nodeType == LogicalDirectGrouping
}

func (wg *WeightedAuthorizationModelGraph) calculateEdgeWildcards(edge *WeightedAuthorizationModelEdge) {
	// if wildcards already exist, we don't need to calculate them
	if len(edge.wildcards) > 0 {
		return
	}
	nodeWildcards := wg.nodes[edge.to.uniqueLabel].wildcards
	// if the node does not have any wildcards, we don't need to calculate them
	if len(nodeWildcards) == 0 {
		return
	}
	// otherwise add the node wildcards
	edge.wildcards = nodeWildcards
}

func (wg *WeightedAuthorizationModelGraph) addReferentialWildcardsToEdge(edge *WeightedAuthorizationModelEdge, referentialNodeID string) {
	referentialNode := wg.nodes[referentialNodeID]
	// if the node does not have any wildcards, we don't need to calculate them
	if len(referentialNode.wildcards) == 0 {
		return
	}
	// if the edge does not have any wildcards, we can add the referential node wildcards
	if len(edge.wildcards) == 0 {
		edge.wildcards = referentialNode.wildcards
		return
	}
	// otherwise add the referential node wildcards to the existing edge wildcards only if the wildcard does not exist in the slice
	// this is to avoid duplicates
	// and to ensure that the wildcards are unique
	for _, wildcard := range referentialNode.wildcards {
		if !slices.Contains(edge.wildcards, wildcard) {
			edge.wildcards = append(edge.wildcards, wildcard)
		}
	}
}

func (wg *WeightedAuthorizationModelGraph) addReferentialWildcardsToNode(nodeID string, referentialNodeID string) {
	referentialNode := wg.nodes[referentialNodeID]
	node := wg.nodes[nodeID]
	// if the node does not have any wildcards, we can add the referential node wildcards
	if len(node.wildcards) == 0 {
		node.wildcards = referentialNode.wildcards
		return
	}
	// otherwise add the referential node wildcards to the existing node wildcards only if the wildcard does not exist in the slice
	// this is to avoid duplicates
	// and to ensure that the wildcards are unique
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

	// if the node does not have any wildcards, we can add the edge wildcards
	if len(node.wildcards) == 0 {
		node.wildcards = edge.wildcards
		return
	}
	// otherwise add the edge wildcards to the existing node wildcards only if the wildcard does not exist in the slice
	// this is to avoid duplicates
	// and to ensure that the wildcards are unique
	for _, wildcard := range edge.wildcards {
		if !slices.Contains(node.wildcards, wildcard) {
			node.wildcards = append(node.wildcards, wildcard)
		}
	}
}

func (wg *WeightedAuthorizationModelGraph) addWildcardToEdge(wildcardType string, edge *WeightedAuthorizationModelEdge) {
	if len(edge.wildcards) == 0 {
		edge.wildcards = []string{wildcardType}
		return
	}

	if !slices.Contains(edge.wildcards, wildcardType) {
		edge.wildcards = append(edge.wildcards, wildcardType)
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
		if edge.to.nodeType == SpecificType || edge.to.nodeType == SpecificTypeWildcard {
			edge.weights = make(map[string]int)
			uniqueLabel := edge.to.uniqueLabel
			if edge.to.nodeType == SpecificTypeWildcard {
				uniqueLabel = uniqueLabel[:len(uniqueLabel)-2]
				wg.addWildcardToEdge(uniqueLabel, edge)
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
	// if it is a recursive edge, we need to set the weight to infinite and add the edge to the tuple cycle dependencies
	if edge.from.uniqueLabel == edge.to.uniqueLabel {
		edge.weights = make(map[string]int)
		edge.weights["R#"+edge.to.uniqueLabel] = Infinite
		tupleCycleDependencies[edge.to.uniqueLabel] = append(tupleCycleDependencies[edge.to.uniqueLabel], edge)
		assignRecursiveCycleMetadata([]*WeightedAuthorizationModelEdge{edge}, edge.from.uniqueLabel)
		return []string{edge.from.uniqueLabel}, nil
	}

	// calculate the weight of the node that is connected to the edge
	ancestorPath = append(ancestorPath, edge)
	tupleCycle, err := wg.calculateNodeWeight(edge.to.uniqueLabel, visited, ancestorPath, tupleCycleDependencies)
	if err != nil {
		return tupleCycle, err
	}

	// if the node that is connected to the edge does not have any weight, we need to check if is a tuple cycle or a model cycle
	if len(edge.to.weights) == 0 {
		if isTupleCycle(edge.to.uniqueLabel, ancestorPath) {
			edge.weights = make(map[string]int)
			edge.weights["R#"+edge.to.uniqueLabel] = Infinite
			tupleCycleDependencies[edge.to.uniqueLabel] = append(tupleCycleDependencies[edge.to.uniqueLabel], edge)
			tupleCycle = append(tupleCycle, edge.to.uniqueLabel)
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

// This function is called when the nodeID is already in the visited path and does not have a weight associated to it.
// In this case we need to know if in the ancestor path to the nodeID is there any edge that is a TTU or a userset.
// If exists we can conclude that is a tuple cycle otherwise is a model cycle.
func isTupleCycle(nodeID string, ancestorPath []*WeightedAuthorizationModelEdge) bool {
	startTracking := false
	tupleCycle := false
	recursiveRelation := ""
	recursion := true
	cyclePath := make([]*WeightedAuthorizationModelEdge, 0)

	for _, edge := range ancestorPath {
		if !startTracking && edge.from.uniqueLabel == nodeID {
			startTracking = true
			if edge.from.nodeType == SpecificTypeAndRelation {
				recursiveRelation = nodeID
			}
		}
		if startTracking {
			cyclePath = append(cyclePath, edge)
			if edge.edgeType == TTUEdge || (edge.edgeType == DirectEdge && edge.to.nodeType == SpecificTypeAndRelation) {
				tupleCycle = true
			}
			if edge.to.nodeType == SpecificTypeAndRelation {
				if len(recursiveRelation) == 0 {
					recursiveRelation = edge.to.GetUniqueLabel()
				} else if edge.to.uniqueLabel != recursiveRelation {
					recursion = false
				}
			}
		}
	}
	if tupleCycle {
		if recursion {
			assignRecursiveCycleMetadata(cyclePath, recursiveRelation)
		} else {
			assignTupleCycleMetadata(cyclePath)
		}
	}

	return tupleCycle
}

// Calculate the node weight base on the weights of all the edges that are connected from the node.
func (wg *WeightedAuthorizationModelGraph) calculateNodeWeightFromTheEdges(nodeID string, tupleCycleDependencies map[string][]*WeightedAuthorizationModelEdge, tupleCycles []string) ([]string, error) {
	var err error
	node := wg.nodes[nodeID]

	// if there is not a cycle, we can calculate the weight without taking in consideration the cycle complexity.
	if len(tupleCycles) == 0 {
		// for any non-operator node, we can calculate the weight using the max strategy because
		// a type is valid as long as it appears in one of the branch.
		if node.nodeType != OperatorNode {
			err := wg.calculateNodeWeightWithMaxStrategy(nodeID)
			if err != nil {
				return tupleCycles, err
			}
		} else {
			switch node.label {
			case UnionOperator:
				// union requires max strategy because if a type appears in any of the "OR" branch,
				// it is valid.
				err := wg.calculateNodeWeightWithMaxStrategy(nodeID)
				if err != nil {
					return tupleCycles, err
				}
			case IntersectionOperator:
				// intersection requires enforce type strategy because a type is valid only if it appears in
				// all the "AND" branches.
				err := wg.calculateNodeWeightWithEnforceTypeStrategy(nodeID)
				if err != nil {
					return tupleCycles, err
				}
			case ExclusionOperator:
				// exclusion (A but not B) requires mix strategy where A is max but B requires enforce type because
				// if a type appears in any of the A branch, it is valid.  However, if a type appears in any of the B
				// branch, it is only valid if it is also in A.
				err := wg.calculateNodeWeightWithMixedStrategy(nodeID)
				if err != nil {
					return tupleCycles, err
				}
			}

		}
		return tupleCycles, nil
	}

	// tuple cycle where the node is responsible for the cycle and it needs to fix the cycle
	if node.nodeType == SpecificTypeAndRelation && wg.isNodeTupleCycleReference(nodeID, tupleCycles) {
		// calculate the weight of the node and fix all the dependencies that are in the tuple cycle.
		err := wg.calculateNodeWeightAndFixDependencies(nodeID, tupleCycleDependencies)
		if err != nil {
			return tupleCycles, err
		}
		tupleCycles = wg.removeNodeFromTupleCycles(nodeID, tupleCycles)
		return tupleCycles, nil
	}

	// there is a cycle but the node is not responsible for the cycle and it is not a logical operator node
	// a logical operator node is an operator node, a logical ttu and a logical userset node
	if !wg.isLogicalOperator(node) {
		// even when there is a cycle, if the relation is not recursive then we calculate the weight using the max strategy
		err = wg.calculateNodeWeightWithMaxStrategy(nodeID)
		if err != nil {
			return tupleCycles, err
		}
		return tupleCycles, nil
	}

	// if the node is a logical union operator,
	// a logical union operator is a union node, a logical userset or a logical ttu.
	if wg.isLogicalUnionOperator(node) {
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

// calculateNodeWeightWithMixedStrategy is a mixed weight strategy used for exclusion node (A but not B).
// For all A edges, we take all the types for all the edges and get the max value
// if more than one edge have the same type in their weights.
func (wg *WeightedAuthorizationModelGraph) calculateNodeWeightWithMixedStrategy(nodeID string) error {
	node := wg.nodes[nodeID]
	edges := wg.edges[nodeID]

	if len(edges) == 0 && node.nodeType != SpecificType && node.nodeType != SpecificTypeWildcard {
		return fmt.Errorf("%w: %s node does not have any terminal type to reach to", ErrInvalidModel, node.uniqueLabel)
	}

	if node.nodeType != OperatorNode && node.label != ExclusionOperator {
		return fmt.Errorf("%w: node %s cannot apply mixed strategy, only accepted exclusion nodes", ErrInvalidModel, nodeID)
	}

	// with the logical ttu and userset, an exclusion operation only has two edges
	// the first edge is the one that defines the userset, the second edge is the one that excludes it
	if len(edges) != 2 {
		return fmt.Errorf("%w: invalid number of edges for exclusion node %s", ErrInvalidModel, nodeID)
	}
	nodeWeights := make(map[string]int, len(edges))
	for k, v := range edges[0].weights {
		nodeWeights[k] = v
	}

	for key, value := range edges[1].weights {
		if w, ok := nodeWeights[key]; ok {
			nodeWeights[key] = int(math.Max(float64(w), float64(value)))
		}
	}

	node.weights = nodeWeights
	return nil
}

// This strategy is used in AND operations and enforces that all the edges return the same type
// if an edge does not return the same type, the key is removed from the weights.
// While doing that process we does not have any type to get in the weight it means that for this operation
// not all paths return the same type and the model is not valid.
func (wg *WeightedAuthorizationModelGraph) calculateNodeWeightWithEnforceTypeStrategy(nodeID string) error {
	node := wg.nodes[nodeID]
	edges := wg.edges[nodeID]

	if len(edges) == 0 && node.nodeType != SpecificType && node.nodeType != SpecificTypeWildcard {
		return fmt.Errorf("%w: %s node does not have any terminal type to reach to", ErrInvalidModel, node.uniqueLabel)
	}

	rewriteWeights := make(map[string]int, len(edges))
	for _, edge := range edges {

		if len(rewriteWeights) == 0 {
			for key, weight := range edge.weights {
				rewriteWeights[key] = weight
			}
		} else {
			for key, rewriteWeight := range rewriteWeights {
				if _, existsAlready := edge.GetWeights()[key]; existsAlready {
					rewriteWeights[key] = int(math.Max(float64(edge.weights[key]), float64(rewriteWeight)))
				} else {
					delete(rewriteWeights, key)
				}
			}
		}
	}

	node.weights = rewriteWeights

	if len(node.weights) == 0 {
		return fmt.Errorf("%w: not all paths return the same type for the node %s", ErrInvalidModel, nodeID)
	}

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

	if node.nodeType != SpecificTypeAndRelation && !wg.isLogicalUnionOperator(node) {
		return fmt.Errorf("%w: invalid node, reference node is not a union operator or a relation or a logical userset or logical TTU: %s", ErrTupleCycle, nodeID)
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

func assignRecursiveCycleMetadata(edgesInCycle []*WeightedAuthorizationModelEdge, recursiveRelation string) {
	// add recursive relation metadata to the edge from node and to the edge
	for _, edge := range edgesInCycle {
		edge.from.recursiveRelation = recursiveRelation
		edge.to.recursiveRelation = recursiveRelation
		edge.recursiveRelation = recursiveRelation
	}
}

func assignTupleCycleMetadata(edgesInCycle []*WeightedAuthorizationModelEdge) {
	// add recursive relation metadata to the edge from node and to the edge
	for _, edge := range edgesInCycle {
		edge.from.tupleCycle = true
		edge.to.tupleCycle = true
		edge.tupleCycle = true
	}
}

// This function will fix the weight of the edges that are dependent on the reference node
// We can take this max weight strategy to remove the dependecies in the edges
// because AND or a BUT NOT are not allowed in a tuple cycle.
func (wg *WeightedAuthorizationModelGraph) fixDependantEdgesWeight(nodeCycle string, referenceNodeID string, references []string, tupleCycleDependencies map[string][]*WeightedAuthorizationModelEdge) {
	node := wg.nodes[nodeCycle]
	tupleCycle := node.tupleCycle

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
		if tupleCycle {
			edge.tupleCycle = true
		}
	}
}

// This function will fix the weight of the nodes that are dependent on the reference node
// We can take this max weight strategy to remove the dependecies in the nodes
// because AND or a BUT NOT are not allowed in a tuple cycle.
func (wg *WeightedAuthorizationModelGraph) fixDependantNodesWeight(nodeCycle string, referenceNodeID string, tupleCycleDependencies map[string][]*WeightedAuthorizationModelEdge) {
	node := wg.nodes[nodeCycle]
	tupleCycle := node.tupleCycle
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
		if tupleCycle {
			fromNode.tupleCycle = true
		}
		wg.addReferentialWildcardsToNode(edge.from.uniqueLabel, nodeCycle)
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
