package graph

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"strings"

	"github.com/oklog/ulid/v2"
)

// LogicEdge represents a simplified, canonical edge used for pattern mining.
// It abstracts away the specific details of different edge types (Compute, TTU)
// into a consistent format for hashing and comparison.
type LogicEdge struct {
	// Type indicates the original edge type (e.g., "Compute", "TTU").
	Type string
	// ID is a canonical identifier for the edge's target.
	// - For a Compute edge, it's the target node's unique label (e.g., "document#viewer").
	// - For a TTU edge, it's a composite key like "folder#parent@viewer".
	ID string
	// Edges stores pointers to the original edges in the graph that this LogicEdge represents.
	// A TTU LogicEdge can represent multiple original edges.
	Edges []*WeightedAuthorizationModelEdge
}

// OperatorIndex stores information about a specific operation node in the graph.
// It captures the context of an operation, including its inputs (LogicEdges)
// and whether it represents a complete, self-contained expression.
type OperatorIndex struct {
	// NodeID is the unique identifier of the operation node
	NodeID string
	// ObjectType is the type definition (e.g., "document") this operation belongs to.
	ObjectType string
	// Operation, exclusion, interception, union
	Operation string
	// Edges are the canonicalized input edges to this operation.
	Edges []*LogicEdge
	// CompleteIndex is true if all inputs to the operation are simple relations (Compute/TTU)
	// and not other nested operations. This makes the pattern a prime candidate for canonical replacement.
	CompleteIndex bool
}

// PatternInfo holds metadata about a frequent combinatorial pattern found in the model.
type PatternInfo struct {
	// Frequency is the number of times this pattern occurs.
	Frequency int
	// Size is the number of operands (LogicEdges) in the pattern.
	Size int
	// Pattern is the canonical hash key representing the expression.
	Pattern string
}

// CanonicalNode links a fully replaceable operation (a "canonical" pattern)
// back to the original relation that defines it (e.g., "document#viewer").
type CanonicalNode struct {
	// RelationNode is the top-level relation node (e.g., "document#viewer").
	RelationNode *WeightedAuthorizationModelNode
	// OperationNodeID is the ID of the operation node this relation points to.
	OperationNodeID string
}

// ModelOptimizer encapsulates the state and indexes needed for the optimization process.
type ModelOptimizer struct {
	graph            *WeightedAuthorizationModelGraph
	operatorIndexes  map[string][]*OperatorIndex
	canonicalNodes   map[string][]*CanonicalNode
	frequentPatterns map[string]*PatternInfo
	// incomingEdges is a reverse index mapping a node's ID to all edges that point to it.
	// This is crucial for O(1) lookups when refactoring the graph.
	incomingEdges map[string][]*WeightedAuthorizationModelEdge
}

// OptimizedGraphOptimization is the main entry point for the model optimization algorithm.
// It iteratively finds and replaces common sub-expressions to improve model performance.
func (g *WeightedAuthorizationModelGraph) OptimizeAuthorizationModelGraph() *WeightedAuthorizationModelGraph {
	optimizer := &ModelOptimizer{
		graph: g,
	}

	// The optimization is iterative. After one pass of replacements, new optimization
	// opportunities may have been created. We loop until no more patterns are found
	// or a max iteration count is reached to prevent infinite loops in complex cases.
	const maxIterations = 5
	for i := 0; i < maxIterations; i++ {
		// Step 0: Clean up state from the previous iteration.
		optimizer.cleanUp()

		// Step 1: Build all necessary indexes for this pass.
		optimizer.buildIndexes()

		// Step 2: Filter and sort the identified frequent patterns.
		// We process the most beneficial patterns first (higher frequency and size).
		patterns := optimizer.filterAndSortFrequencyPatterns()
		if len(patterns) == 0 {
			break // No more optimizable patterns found, exit the loop.
		}

		// Step 3: Replace the found patterns in the graph.
		optimizer.replacePatterns(patterns)
	}

	return optimizer.graph
}

// buildIndexes constructs the necessary data structures for finding patterns.
// It traverses the graph and populates the optimizer's indexes.
func (o *ModelOptimizer) buildIndexes() {
	visitedNodes := make(map[string]bool)
	for _, node := range o.graph.GetNodes() {
		// We only start the indexing process from relation nodes.
		if node.GetNodeType() == SpecificTypeAndRelation {
			o.buildIndexForRelationNode(node, visitedNodes)
		}
	}
}

// buildIndexForRelationNode is the entry point for indexing a specific relation.
func (o *ModelOptimizer) buildIndexForRelationNode(node *WeightedAuthorizationModelNode, visitedNodes map[string]bool) {
	if visitedNodes[node.GetUniqueLabel()] || node.GetNodeType() != SpecificTypeAndRelation {
		return
	}
	visitedNodes[node.GetUniqueLabel()] = true

	// if a node is a recursive relation, skip the optimization for now
	// the algorithm will need to handle the case and identify that we cannot extract viewer from parent or admin
	// because viewer from parent is a recursive TTU
	// type document
	//      relations
	//        admin: [user]
	//        parent: [document]
	//        member: viewer from parent or admin
	//        viewer: [user] or viewer from parent or admin
	if len(node.recursiveRelation) > 0 {
		return
	}

	edges, ok := o.graph.GetEdges()[node.GetUniqueLabel()]
	// A relation should have exactly one edge defining its logic for this step, looking into relations with only operation link
	if !ok || len(edges) != 1 {
		return
	}

	edge := edges[0]
	// We are only interested in relations defined by an operation (OR, AND, BUT NOT).
	if edge.GetEdgeType() != RewriteEdge || edge.GetTo().GetNodeType() != OperatorNode {
		return
	}

	// name of node relations are in the format objectType#relation
	objectType := strings.Split(node.GetUniqueLabel(), "#")[0]

	// Build the reverse index of incoming edges.
	if _, ok := o.incomingEdges[edge.GetTo().uniqueLabel]; !ok {
		o.incomingEdges[edge.GetTo().uniqueLabel] = []*WeightedAuthorizationModelEdge{}
	}
	o.incomingEdges[edge.GetTo().uniqueLabel] = append(o.incomingEdges[edge.GetTo().uniqueLabel], edge)

	// Recursively build indexes for the operation node.
	indexID := o.buildIndexForOperationNode(edge.GetTo(), objectType)

	// If the operation was a "complete" expression, it's a canonical pattern.
	if indexID != "" {
		if index, ok := o.operatorIndexes[indexID]; ok && len(index) > 0 && index[0].CompleteIndex {
			if _, ok := o.canonicalNodes[indexID]; !ok {
				o.canonicalNodes[indexID] = []*CanonicalNode{}
			}
			o.canonicalNodes[indexID] = append(o.canonicalNodes[indexID], &CanonicalNode{
				RelationNode:    node,
				OperationNodeID: edge.GetTo().uniqueLabel,
			})
		}
	}
}

// buildIndexForOperationNode recursively builds indexes for an operation and its children.
func (o *ModelOptimizer) buildIndexForOperationNode(node *WeightedAuthorizationModelNode, objectType string) string {
	if node.GetNodeType() != OperatorNode {
		return ""
	}

	// Exclusion (BUT NOT) operators have a different structure and logic.
	if node.GetLabel() == ExclusionOperator {
		return o.buildIndexForExclusion(node, objectType)
	}

	edges, ok := o.graph.GetEdges()[node.GetUniqueLabel()]
	if !ok {
		return ""
	}

	completeIndex := true
	logicEdges := make(map[string]*LogicEdge)

	for _, edge := range edges {
		// If an input is a direct relation, it's not part of a combinatorial pattern.
		if edge.GetEdgeType() == DirectEdge {
			completeIndex = false
			continue
		}

		if (edge.GetEdgeType() == ComputedEdge || edge.GetEdgeType() == RewriteEdge) && edge.GetTo().GetNodeType() == SpecificTypeAndRelation {
			if !o.addComputeLogicEdge(logicEdges, edge) {
				// Duplicate compute edge (e.g., "viewer or viewer"), mark for deletion.
				o.graph.deleteEdge(edge.GetFrom().GetUniqueLabel(), edge)
			}
		} else if edge.GetEdgeType() == TTUEdge {
			o.addTTULogicEdge(logicEdges, edge)
		} else if edge.GetTo().nodeType == OperatorNode {
			// If an input is another operation, this one is not "complete".
			completeIndex = false
			toNodeId := edge.GetTo().uniqueLabel
			// Build the reverse index for the nested operation.
			if _, ok := o.incomingEdges[toNodeId]; !ok {
				o.incomingEdges[toNodeId] = []*WeightedAuthorizationModelEdge{}
			}
			o.incomingEdges[toNodeId] = append(o.incomingEdges[toNodeId], edge)
			// Recurse into the nested operation.
			o.buildIndexForOperationNode(edge.GetTo(), objectType)
		}
	}

	// We only care about operations with 2 or more operands.
	if len(logicEdges) < 2 {
		return ""
	}

	// Convert map to slice for sorting and combination generation.
	logicEdgesSlice := make([]*LogicEdge, 0, len(logicEdges))
	for _, le := range logicEdges {
		logicEdgesSlice = append(logicEdgesSlice, le)
	}
	// Sort to ensure canonical representation.
	sort.Slice(logicEdgesSlice, func(i, j int) bool {
		return logicEdgesSlice[i].ID < logicEdgesSlice[j].ID
	})

	// Create the index for the full pattern.
	operatorHash := o.createIndex(node, logicEdgesSlice, objectType, completeIndex)

	// Mine for smaller, frequent sub-patterns.
	o.mineFrequentPattern(node, logicEdgesSlice, objectType, 2)
	o.mineFrequentPattern(node, logicEdgesSlice, objectType, 3)

	return operatorHash
}

// buildIndexForExclusion handles the specific logic for "BUT NOT" operators.
func (o *ModelOptimizer) buildIndexForExclusion(node *WeightedAuthorizationModelNode, objectType string) string {
	edges, ok := o.graph.GetEdges()[node.GetUniqueLabel()]
	if !ok {
		return ""
	}
	createIndex := true

	// in the case of exclusion if starts with a direct edge the exclusion cannot be indexed
	//however we still need to traverse the subgraph if there are other
	//combinations of operations inside of the right side of the exclusion
	if edges[0].GetEdgeType() == DirectEdge {
		createIndex = false
	}

	logicEdges := make(map[string]*LogicEdge)
	for _, edge := range edges {
		if edge.to.nodeType == OperatorNode {
			createIndex = false
			o.buildIndexForOperationNode(edge.to, objectType)
		} else if (edge.GetEdgeType() == ComputedEdge || edge.GetEdgeType() == RewriteEdge) && edge.GetTo().GetNodeType() == SpecificTypeAndRelation {
			if !o.addComputeLogicEdge(logicEdges, edge) {
				// Duplicate compute edge (e.g., "viewer or viewer"), mark for deletion.
				o.graph.deleteEdge(edge.GetFrom().GetUniqueLabel(), edge)
			}
		} else if edge.edgeType == TTUEdge {
			o.addTTULogicEdge(logicEdges, edge)
		}
	}

	if !createIndex || len(logicEdges) < 2 {
		return ""
	}

	// Convert map to slice for sorting and combination generation.
	logicEdgesSlice := make([]*LogicEdge, 0, len(logicEdges))
	for _, le := range logicEdges {
		logicEdgesSlice = append(logicEdgesSlice, le)
	}
	// Sort to ensure canonical representation.
	sort.Slice(logicEdgesSlice, func(i, j int) bool {
		return logicEdgesSlice[i].ID < logicEdgesSlice[j].ID
	})

	// The operation node for the index is the subtract node, not the top-level exclusion node.
	return o.createIndex(node, logicEdgesSlice, objectType, true)
}

// createIndex generates a canonical key for an operation and adds it to the indexes.
func (o *ModelOptimizer) createIndex(node *WeightedAuthorizationModelNode, logicEdges []*LogicEdge, objectType string, complete bool) string {
	operatorKey := createCanonicalKey(objectType, node, logicEdges)

	if _, ok := o.operatorIndexes[operatorKey]; !ok {
		o.operatorIndexes[operatorKey] = []*OperatorIndex{}
	}
	o.operatorIndexes[operatorKey] = append(o.operatorIndexes[operatorKey], &OperatorIndex{
		NodeID:        node.GetUniqueLabel(),
		Operation:     node.GetLabel(),
		ObjectType:    objectType,
		Edges:         logicEdges,
		CompleteIndex: complete,
	})

	if info, ok := o.frequentPatterns[operatorKey]; ok {
		info.Frequency++
	} else {
		o.frequentPatterns[operatorKey] = &PatternInfo{
			Frequency: 1,
			Size:      len(logicEdges),
			Pattern:   operatorKey,
		}
	}
	return operatorKey
}

// createCanonicalKey creates a stable, unique string representation of an operation.
// OPTIMIZED: This function now uses strings.Builder for efficient string construction,
// reducing memory allocations compared to fmt.Sprintf or simple concatenation.
func createCanonicalKey(objectType string, node *WeightedAuthorizationModelNode, logicEdges []*LogicEdge) string {
	var sb strings.Builder
	// Pre-calculating the required size can be a further optimization, but for
	// clarity and typical operand counts, this is already a huge improvement.
	sb.WriteString(objectType)
	sb.WriteByte('@')
	sb.WriteString(node.GetLabel())
	sb.WriteByte(':')

	for i, ledge := range logicEdges {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(ledge.ID)
	}
	content := sb.String()

	// Use a hash for very long keys to keep map performance high.
	if len(content) > 100 {
		hash := sha256.Sum256([]byte(content))
		return fmt.Sprintf("hash_%x", hash[:8]) // Use first 8 bytes for a short but unique key.
	}
	return content
}

// addComputeLogicEdge adds a 'compute' edge to the logic edge map. Returns false if it's a duplicate.
func (o *ModelOptimizer) addComputeLogicEdge(logicEdges map[string]*LogicEdge, edge *WeightedAuthorizationModelEdge) bool {
	key := edge.GetTo().uniqueLabel
	if _, ok := logicEdges[key]; ok {
		return false // Duplicate detected.
	}
	logicEdges[key] = &LogicEdge{
		Type:  "ComputeEdge",
		ID:    key,
		Edges: []*WeightedAuthorizationModelEdge{edge},
	}
	return true
}

// addTTULogicEdge adds a 'ttu' edge to the logic edge map.
func (o *ModelOptimizer) addTTULogicEdge(logicEdges map[string]*LogicEdge, edge *WeightedAuthorizationModelEdge) {
	// For TTU, the canonical ID includes the tupleset relation.
	// e.g., "parent@viewer" for "viewer from parent".
	// in a ttu the to Node is always a relation, relation are in teh format of objectType#relation
	// object type can differ, we could have multiple in the case parent references to multiple terminal types
	// but relation is always is going to be the same for the same ttu
	rel := strings.Split(edge.GetTo().uniqueLabel, "#")[1]
	key := fmt.Sprintf("%s@%s", edge.GetTuplesetRelation(), rel)

	if le, ok := logicEdges[key]; ok {
		le.Edges = append(le.Edges, edge)
	} else {
		logicEdges[key] = &LogicEdge{
			Type:  "TTUEdge",
			ID:    key,
			Edges: []*WeightedAuthorizationModelEdge{edge},
		}
	}
}

// mineFrequentPattern finds and indexes combinatorial sub-patterns of a specific size.
func (o *ModelOptimizer) mineFrequentPattern(node *WeightedAuthorizationModelNode, logicEdges []*LogicEdge, objectType string, subsetSize int) {
	if len(logicEdges) <= subsetSize {
		return
	}

	// Non-recursive combination generator for fixed small sizes (2 and 3).
	// This is more performant than a general recursive solution for this specific use case.
	if subsetSize == 2 {
		for i := 0; i < len(logicEdges)-1; i++ {
			for j := i + 1; j < len(logicEdges); j++ {
				combination := []*LogicEdge{logicEdges[i], logicEdges[j]}
				o.createIndex(node, combination, objectType, false)
			}
		}
	} else if subsetSize == 3 {
		for i := 0; i < len(logicEdges)-2; i++ {
			for j := i + 1; j < len(logicEdges)-1; j++ {
				for k := j + 1; k < len(logicEdges); k++ {
					combination := []*LogicEdge{logicEdges[i], logicEdges[j], logicEdges[k]}
					o.createIndex(node, combination, objectType, false)
				}
			}
		}
	}
}

// filterAndSortFrequencyPatterns prepares the list of patterns to be replaced.
func (o *ModelOptimizer) filterAndSortFrequencyPatterns() []*PatternInfo {
	patterns := make([]*PatternInfo, 0, len(o.frequentPatterns))
	for _, patternInfo := range o.frequentPatterns {
		if patternInfo.Frequency > 1 {
			patterns = append(patterns, patternInfo)
		}
	}

	sort.Slice(patterns, func(i, j int) bool {
		// Primary criterion: sort by frequency, descending.
		if patterns[i].Frequency != patterns[j].Frequency {
			return patterns[i].Frequency > patterns[j].Frequency
		}
		// Secondary criterion (tie-breaker): sort by size, descending.
		return patterns[i].Size > patterns[j].Size
	})
	return patterns
}

// replacePatterns iterates through the sorted patterns and refactors the graph.
func (o *ModelOptimizer) replacePatterns(patterns []*PatternInfo) error {
	modifiedNodes := make(map[string]bool)
	for _, pattern := range patterns {
		// Get only the operator indexes that haven't been modified in this pass.
		operatorIndexes := o.getValidOperatorIndexes(pattern.Pattern, modifiedNodes)
		if len(operatorIndexes) > 1 {
			err := o.replacePattern(pattern.Pattern, operatorIndexes, modifiedNodes)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// getValidOperatorIndex filters out operator indexes that belong to already modified nodes.
// OPTIMIZED: Pre-allocates the slice with a capacity to avoid reallocations.
func (o *ModelOptimizer) getValidOperatorIndexes(key string, modifiedNodes map[string]bool) []*OperatorIndex {
	originalIndexes := o.operatorIndexes[key]
	// Pre-allocate with the maximum possible size to avoid append-related reallocations.
	validIndexes := make([]*OperatorIndex, 0, len(originalIndexes))
	for _, index := range originalIndexes {
		if !modifiedNodes[index.NodeID] {
			validIndexes = append(validIndexes, index)
		}
	}
	return validIndexes
}

// replacePattern performs the graph surgery for a single frequent pattern.
func (o *ModelOptimizer) replacePattern(key string, operatorIndexes []*OperatorIndex, modifiedNodes map[string]bool) error {
	var relationNode *WeightedAuthorizationModelNode
	var operatorNodeID string

	// If a canonical definition exists, reuse it.
	if canonical, ok := o.canonicalNodes[key]; ok && len(canonical) > 0 {
		relationNode = canonical[0].RelationNode
		operatorNodeID = canonical[0].OperationNodeID
	} else {
		// Otherwise, create a new optimized relation (e.g., "opt_...").
		firstIndex := operatorIndexes[0]
		operationNode, err := o.addOperationNode(firstIndex.Operation, firstIndex.Edges)
		if err != nil {
			return err
		}
		relationNode = o.addRelationNode(firstIndex.ObjectType, operationNode)
	}

	o.replacePatternForRelationNode(relationNode, operatorNodeID, operatorIndexes, modifiedNodes)
	return nil
}

// replacePatternForRelationNode connects the new opt relation back into the original expressions.
func (o *ModelOptimizer) replacePatternForRelationNode(relationNode *WeightedAuthorizationModelNode, operatorNodeID string, operatorIndexes []*OperatorIndex, modifiedNodes map[string]bool) {
	for _, opIndex := range operatorIndexes {
		// Don't modify the canonical node itself.
		if operatorNodeID != "" && operatorNodeID == opIndex.NodeID {
			continue
		}

		// Delete the old edges that are now part of the optimized relation.
		for _, ledge := range opIndex.Edges {
			for _, edge := range ledge.Edges {
				o.graph.deleteEdge(opIndex.NodeID, edge)
			}
		}

		// If the entire operation is replaced (CompleteIndex), we can rewire and delete the old op node.
		if opIndex.CompleteIndex {
			incomingEdges := o.incomingEdges[opIndex.NodeID]
			for _, edge := range incomingEdges {
				// Point the parent relation directly to the new/reused opt relation.
				edge.to = relationNode
				edge.edgeType = ComputedEdge // It was Referential, now it's a direct Compute.
			}
			o.graph.deleteNode(opIndex.NodeID)
			modifiedNodes[opIndex.NodeID] = true
		} else {
			// Otherwise, just add a new edge from the original operation to the new opt relation.
			o.graph.AddEdge(opIndex.NodeID, relationNode.GetUniqueLabel(), ComputedEdge, "", nil)
			modifiedNodes[opIndex.NodeID] = true
		}
	}
}

// addRelationNode creates a new relation node (e.g., "document#opt_123").
func (o *ModelOptimizer) addRelationNode(objectType string, operationNode *WeightedAuthorizationModelNode) *WeightedAuthorizationModelNode {
	// Generate a unique name for the new relation using strings.Builder for performance.
	var sb strings.Builder
	sb.WriteString("opt_")
	sb.WriteString(strings.ReplaceAll(ulid.Make().String(), "-", ""))
	relationName := sb.String()
	uniqueLabel := fmt.Sprintf("%s#%s", objectType, relationName)
	// Add the node to the graph.
	relationNode := o.graph.GetOrAddNode(uniqueLabel, uniqueLabel, SpecificTypeAndRelation)
	o.graph.AddEdge(uniqueLabel, operationNode.uniqueLabel, RewriteEdge, "", nil)
	rewriteEdge := o.graph.edges[uniqueLabel][0]
	rewriteEdge.weights = operationNode.weights
	rewriteEdge.wildcards = operationNode.wildcards
	relationNode.weights = rewriteEdge.weights
	relationNode.wildcards = rewriteEdge.wildcards
	return relationNode
}

// addOperationNode creates a new relation node (e.g., "document#opt_123").
func (o *ModelOptimizer) addOperationNode(operation string, logicEdges []*LogicEdge) (*WeightedAuthorizationModelNode, error) {
	operatorNodeName := operation + ":" + ulid.Make().String()
	// Add the node to the graph.
	operatorNode := o.graph.GetOrAddNode(operatorNodeName, operation, OperatorNode)
	// Copy the edges from the pattern to this new  node.
	for _, ledge := range logicEdges {
		for _, edge := range ledge.Edges {
			o.graph.AddEdge(operatorNodeName, edge.to.uniqueLabel, edge.edgeType, edge.tuplesetRelation, edge.conditions)
		}
	}

	// Recalculate graph properties for the new node.
	err := o.graph.calculateNodeWeightBasedOnPrecomputedEdges(operatorNode)
	if err != nil {
		return operatorNode, err
	}
	err = o.graph.calculateWildcardNodeBasedOnEdges(operatorNodeName)
	if err != nil {
		return operatorNode, err
	}

	return operatorNode, nil
}

// cleanUp resets the optimizer's state for a new iteration.
func (o *ModelOptimizer) cleanUp() {
	o.operatorIndexes = make(map[string][]*OperatorIndex)
	o.canonicalNodes = make(map[string][]*CanonicalNode)
	o.frequentPatterns = make(map[string]*PatternInfo)
	o.incomingEdges = make(map[string][]*WeightedAuthorizationModelEdge)
}
