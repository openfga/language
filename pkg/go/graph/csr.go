package graph

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"iter"
	"math"
	"unsafe"
)

type (
	Node int32
	Edge int32
)

// A CSR is a compressed sparse row representation of the weighted graph.
// It is immutable for its lifetime, has a deterministic structure for a
// given model, and can be encoded to binary and decoded.
//
// To construct a CSR, use the functions [NewCSR] or [CSR.UnmarshalBinary].
type CSR struct {
	nodes int32
	edges int32

	nodeTypes              []uint8
	nodeRecursiveRelations []Node
	nodeCycles             []uint64

	nodeLabelPointers []int32
	nodeLabels        []byte

	nodeIDPointers []int32
	nodeIDs        []byte

	nodeReachabilityPointers []int32
	nodeReachability         []uint64

	nodeWildcardPointers []int32
	nodeWildcards        []Node

	idToNode map[string]int32

	edgePointers            []int32
	edgeTypes               []uint8
	edgeEndpoints           []uint64
	edgeRecursiveRelations  []Node
	edgeRelationDefinitions []Node
	edgeCycles              []uint64

	edgeReachabilityPointers []int32
	edgeReachability         []uint64

	edgeWildcardPointers []int32
	edgeWildcards        []Node

	edgeConditionPointers []int32
	edgeConditions        []byte
}

// Nodes returns an iterator over all nodes in the CSR.
func (csr *CSR) Nodes() iter.Seq2[int, Node] {
	return func(yield func(int, Node) bool) {
		for i := range csr.nodes {
			if !yield(int(i), Node(i)) {
				break
			}
		}
	}
}

// NodeLabel returns the label for a given node.
func (csr *CSR) NodeLabel(node Node) string {
	start := csr.nodeLabelPointers[node]
	end := csr.nodeLabelPointers[node+1]

	if start >= end {
		return ""
	}
	b := csr.nodeLabels[start:end]
	return unsafe.String(unsafe.SliceData(b), len(b))
}

// NodeID returns the unique label for a given node.
func (csr *CSR) NodeID(node Node) string {
	start := csr.nodeIDPointers[node]
	end := csr.nodeIDPointers[node+1]

	if start >= end {
		return ""
	}
	b := csr.nodeIDs[start:end]
	return unsafe.String(unsafe.SliceData(b), len(b))
}

// NodeFromID returns a node for a unique label.
// NodeFromID returns false when a node does not exist
// for the given value.
func (csr *CSR) NodeFromID(id string) (Node, bool) {
	ndx, ok := csr.idToNode[id]
	if !ok {
		return -1, false
	}
	return Node(ndx), true
}

// MustNodeFromID behaves exactly like [CSR.NodeFromID]
// but panics when a node does not exist for the given
// value.
func (csr *CSR) MustNodeFromID(id string) Node {
	if node, ok := csr.NodeFromID(id); ok {
		return node
	}
	panic(fmt.Errorf("node does not exist for id '%s'", id))
}

// NodeType returns a node type for a given node.
func (csr *CSR) NodeType(node Node) NodeType {
	return NodeType(csr.nodeTypes[node])
}

// NodeRecursiveRelation returns the recursive relation
// for a given node. NodeRecursiveRelation returns false
// when the node does not have a recursive relation.
func (csr *CSR) NodeRecursiveRelation(node Node) (Node, bool) {
	recursiveNode := csr.nodeRecursiveRelations[node]
	return recursiveNode, recursiveNode >= 0
}

// NodeWeight returns the maximum weight to a terminal node
// across all of the source node's outgoing edges. NodeWeight
// returns false when the terminal node is not reachable from
// the source node.
func (csr *CSR) NodeWeight(sourceNode Node, terminalNode Node) (int, bool) {
	start := csr.nodeReachabilityPointers[sourceNode]
	end := csr.nodeReachabilityPointers[sourceNode+1]

	for i := start; i < end; i++ {
		n := Node(csr.nodeReachability[i] >> 32)
		w := int32(csr.nodeReachability[i] & 0xFFFFFFFF)

		if n == terminalNode {
			return int(w), w > 0
		}
	}
	return 0, false
}

// EdgesFromNode returns an iterator over all of the given
// node's outgoing edges.
func (csr *CSR) EdgesFromNode(sourceNode Node) iter.Seq2[int, Edge] {
	start := csr.edgePointers[sourceNode]
	end := csr.edgePointers[sourceNode+1]

	return func(yield func(int, Edge) bool) {
		for i := start; i < end; i++ {
			if !yield(int(i), Edge(i)) {
				break
			}
		}
	}
}

// NodeReachesWildcard returns true when the source node can
// reach a wildcard for given type name.
func (csr *CSR) NodeReachesWildcard(sourceNode Node, name string) bool {
	wildcardNode, ok := csr.NodeFromID(name + ":*")
	if !ok {
		return false
	}

	start := csr.nodeWildcardPointers[sourceNode]
	end := csr.nodeWildcardPointers[sourceNode+1]

	for i := start; i < end; i++ {
		if csr.nodeWildcards[i] == wildcardNode {
			return true
		}
	}
	return false
}

// NodeIsCyclic returns true when the given node is part of a
// tuple cycle.
func (csr *CSR) NodeIsCyclic(node Node) bool {
	ndx := node / 64
	bit := node % 64
	return ((csr.nodeCycles[ndx] >> bit) & 1) != 0
}

// EdgeType returns the edge type for the given edge.
func (csr *CSR) EdgeType(edge Edge) EdgeType {
	return EdgeType(csr.edgeTypes[edge])
}

func (csr *CSR) EdgeSource(edge Edge) Node {
	return Node(csr.edgeEndpoints[edge] >> 32)
}

func (csr *CSR) EdgeTarget(edge Edge) Node {
	return Node(csr.edgeEndpoints[edge] & 0xFFFFFFFF)
}

func (csr *CSR) EdgeRecursiveRelation(edge Edge) Node {
	return csr.edgeRecursiveRelations[edge]
}

func (csr *CSR) EdgeRelationDefinition(edge Edge) Node {
	return csr.edgeRelationDefinitions[edge]
}

func (csr *CSR) EdgeIsCyclic(edge Edge) bool {
	ndx := edge / 64
	bit := edge % 64
	return ((csr.edgeCycles[ndx] >> bit) & 1) != 0
}

func (csr *CSR) EdgeWeight(sourceEdge Edge, terminalNode Node) (int, bool) {
	start := csr.edgeReachabilityPointers[sourceEdge]
	end := csr.edgeReachabilityPointers[sourceEdge+1]

	for i := start; i < end; i++ {
		n := Node(csr.edgeReachability[i] >> 32)
		w := int32(csr.edgeReachability[i] & 0xFFFFFFFF)

		if n == terminalNode {
			return int(w), w > 0
		}
	}
	return 0, false
}

func (csr *CSR) EdgeReachesWildcard(sourceEdge Edge, name string) bool {
	wildcardNode, ok := csr.NodeFromID(name + ":*")
	if !ok {
		return false
	}

	start := csr.edgeWildcardPointers[sourceEdge]
	end := csr.edgeWildcardPointers[sourceEdge+1]

	for i := start; i < end; i++ {
		if csr.edgeWildcards[i] == wildcardNode {
			return true
		}
	}
	return false
}

func (csr *CSR) EdgeConditions(edge Edge) iter.Seq2[int, string] {
	start := csr.edgeConditionPointers[edge]
	end := csr.edgeConditionPointers[edge+1]

	conditions := csr.edgeConditions[start:end]

	return func(yield func(int, string) bool) {
		var ctr int
		for len(conditions) > 0 {
			var condition []byte

			i := bytes.IndexByte(conditions, 0x0)
			if 0 > i {
				condition = conditions
				conditions = conditions[:0]
			} else {
				condition = conditions[:i]
				conditions = conditions[i+1:]
			}

			value := unsafe.String(
				unsafe.SliceData(condition),
				len(condition),
			)

			if !yield(ctr, value) {
				break
			}
			ctr++
		}
	}
}

func pushNode(csr *CSR, node *WeightedAuthorizationModelNode) Node {
	ndx := csr.nodes
	csr.nodes++

	csr.nodeTypes[ndx] = uint8(node.nodeType)

	var nodeCycle uint64
	if node.tupleCycle {
		nodeCycle = 1
	}
	n := ndx / 64
	b := ndx % 64
	csr.nodeCycles[n] |= nodeCycle << b

	csr.nodeLabels = append(csr.nodeLabels, node.label...)
	if len(csr.nodeLabels) > math.MaxInt32 {
		panic("node labels are oversized")
	}
	csr.nodeLabelPointers[csr.nodes] = int32(len(csr.nodeLabels))

	csr.nodeIDs = append(csr.nodeIDs, node.uniqueLabel...)
	if len(csr.nodeIDs) > math.MaxInt32 {
		panic("node ids are oversized")
	}
	csr.nodeIDPointers[csr.nodes] = int32(len(csr.nodeIDs))

	csr.idToNode[node.uniqueLabel] = ndx

	return Node(ndx)
}

func pushEdge(csr *CSR, edge *WeightedAuthorizationModelEdge) Edge {
	ndx := csr.edges
	csr.edges++

	csr.edgeTypes[ndx] = uint8(edge.edgeType)

	var edgeCycle uint64
	if edge.tupleCycle {
		edgeCycle = 1
	}
	n := ndx / 64
	b := ndx % 64
	csr.edgeCycles[n] |= edgeCycle << b

	sourceNode := csr.MustNodeFromID(edge.from.uniqueLabel)
	targetNode := csr.MustNodeFromID(edge.to.uniqueLabel)

	csr.edgeEndpoints[ndx] |= uint64(sourceNode) << 32
	csr.edgeEndpoints[ndx] |= uint64(targetNode)

	for nodeID, weight := range edge.weights {
		leafNode := csr.MustNodeFromID(nodeID)

		var reachability uint64
		reachability |= uint64(leafNode) << 32
		reachability |= uint64(weight)

		csr.edgeReachability = append(
			csr.edgeReachability,
			reachability,
		)

		if len(csr.edgeReachability) > math.MaxInt32 {
			panic("edge reachability is oversized")
		}
	}
	csr.edgeReachabilityPointers[ndx+1] = int32(len(csr.edgeReachability))

	for _, wildcardType := range edge.wildcards {
		targetNode := csr.MustNodeFromID(wildcardType + ":*")
		csr.edgeWildcards = append(csr.edgeWildcards, targetNode)
		if len(csr.edgeWildcards) > math.MaxInt32 {
			panic("edge wildcards is oversized")
		}
	}
	csr.edgeWildcardPointers[ndx+1] = int32(len(csr.edgeWildcards))

	relation := edge.recursiveRelation
	if relation == "" {
		csr.edgeRecursiveRelations[ndx] = -1
	} else {
		csr.edgeRecursiveRelations[ndx] = csr.MustNodeFromID(relation)
	}

	relation = edge.relationDefinition
	csr.edgeRelationDefinitions[ndx] = csr.MustNodeFromID(relation)

	for _, condition := range edge.GetConditions() {
		csr.edgeConditions = append(
			csr.edgeConditions,
			[]byte(condition)...,
		)
		csr.edgeConditions = append(csr.edgeConditions, 0x0)
	}
	csr.edgeConditionPointers[ndx+1] = int32(len(csr.edgeConditions))

	return Edge(ndx)
}

func applyNodeWildcards(csr *CSR, node *WeightedAuthorizationModelNode) {
	for _, wildcardType := range node.wildcards {
		targetNode := csr.MustNodeFromID(wildcardType + ":*")
		csr.nodeWildcards = append(csr.nodeWildcards, targetNode)
		if len(csr.nodeWildcards) > math.MaxInt32 {
			panic("node wildcards is oversized")
		}
	}
}

func applyNodeWeights(csr *CSR, node *WeightedAuthorizationModelNode) {
	for nodeID, weight := range node.weights {
		terminalNode := csr.MustNodeFromID(nodeID)

		var reachability uint64
		reachability |= uint64(terminalNode) << 32
		reachability |= uint64(weight)

		csr.nodeReachability = append(
			csr.nodeReachability,
			reachability,
		)
		if len(csr.nodeReachability) > math.MaxInt32 {
			panic("node reachability is oversized")
		}
	}
}

func NewCSR(g *WeightedAuthorizationModelGraph) *CSR {
	var csr CSR
	csr.idToNode = make(map[string]int32)

	// ---------- Node Field Array Initialization ----------
	nodeCount := len(g.nodes)
	csr.nodeTypes = make([]uint8, nodeCount)
	csr.nodeCycles = make([]uint64, (nodeCount/64)+1)
	csr.nodeRecursiveRelations = make([]Node, nodeCount)

	// ---------- Node Pointer Array Initialization ----------
	pointerCount := nodeCount + 1
	csr.nodeLabelPointers = make([]int32, pointerCount)
	csr.nodeIDPointers = make([]int32, pointerCount)
	csr.nodeReachabilityPointers = make([]int32, pointerCount)
	csr.nodeWildcardPointers = make([]int32, pointerCount)
	csr.edgePointers = make([]int32, pointerCount)

	// collect nodes into an array to stabalize iteration order
	// set order is determined by map enumeration
	nodes := make([]*WeightedAuthorizationModelNode, 0, nodeCount)

	var nodeLabelSize int
	var nodeIDSize int
	var nodeWildcardSize int
	var nodeReachabilitySize int
	var conditionsSize int

	for _, node := range g.nodes {
		nodes = append(nodes, node)
		nodeLabelSize += len(node.label)
		nodeIDSize += len(node.uniqueLabel)
		nodeWildcardSize += len(node.wildcards)
		nodeReachabilitySize += len(node.weights)

		for _, edge := range g.edges[node.uniqueLabel] {
			conditionsSize += len(edge.GetConditions())
		}
	}

	csr.nodeLabels = make([]byte, 0, nodeLabelSize)
	csr.nodeIDs = make([]byte, 0, nodeIDSize)
	csr.nodeWildcards = make([]Node, 0, nodeWildcardSize)
	csr.nodeReachability = make([]uint64, 0, nodeReachabilitySize)
	csr.edgeConditions = make([]byte, 0, conditionsSize)

	for _, node := range nodes {
		pushNode(&csr, node)
	}

	var edgeCount int

	for ndx, node := range nodes {
		relation := node.recursiveRelation
		recursiveNode, ok := csr.NodeFromID(relation)
		csr.nodeRecursiveRelations[ndx] = recursiveNode
		if !ok {
			csr.nodeRecursiveRelations[ndx] = -1
		}

		applyNodeWeights(&csr, node)
		reachPointer := int32(len(csr.nodeReachability))
		csr.nodeReachabilityPointers[ndx+1] = reachPointer

		applyNodeWildcards(&csr, node)
		wildcardPointer := int32(len(csr.nodeWildcards))
		csr.nodeWildcardPointers[ndx+1] = wildcardPointer

		edgeCount += len(g.edges[node.uniqueLabel])
	}

	// ---------- Edge Field Array Initizlization ----------
	csr.edgeTypes = make([]uint8, edgeCount)
	csr.edgeCycles = make([]uint64, (edgeCount/64)+1)
	csr.edgeEndpoints = make([]uint64, edgeCount)
	csr.edgeRecursiveRelations = make([]Node, edgeCount)
	csr.edgeRelationDefinitions = make([]Node, edgeCount)

	// ---------- Edge Pointer Array Initialization ----------
	pointerCount = edgeCount + 1
	csr.edgeReachabilityPointers = make([]int32, pointerCount)
	csr.edgeWildcardPointers = make([]int32, pointerCount)
	csr.edgeConditionPointers = make([]int32, pointerCount)

	for ndx, node := range nodes {
		for _, edge := range g.edges[node.uniqueLabel] {
			pushEdge(&csr, edge)
		}
		csr.edgePointers[ndx+1] = csr.edges
	}

	return &csr
}

func (csr *CSR) MarshalBinary() ([]byte, error) {
	size := 8 // size properties
	size += len(csr.nodeTypes)
	size += len(csr.nodeRecursiveRelations) * 4
	size += len(csr.nodeCycles) * 8
	size += len(csr.nodeLabelPointers) * 4
	size += len(csr.nodeLabels)
	size += len(csr.nodeIDPointers) * 4
	size += len(csr.nodeIDs)
	size += len(csr.nodeReachabilityPointers) * 4
	size += len(csr.nodeReachability) * 8
	size += len(csr.nodeWildcardPointers) * 4
	size += len(csr.nodeWildcards) * 4
	size += len(csr.edgePointers) * 4
	size += len(csr.edgeTypes)
	size += len(csr.edgeEndpoints) * 8
	size += len(csr.edgeRecursiveRelations) * 4
	size += len(csr.edgeRelationDefinitions) * 4
	size += len(csr.edgeCycles) * 8
	size += len(csr.edgeReachabilityPointers) * 4
	size += len(csr.edgeReachability) * 8
	size += len(csr.edgeWildcardPointers) * 4
	size += len(csr.edgeWildcards) * 4
	size += len(csr.edgeConditionPointers) * 4
	size += len(csr.edgeConditions)
	size += 23 * 4 // length encoding

	buf := make([]byte, 0, size)

	buf = binary.BigEndian.AppendUint32(buf, uint32(csr.nodes))
	buf = binary.BigEndian.AppendUint32(buf, uint32(csr.edges))

	buf = binary.BigEndian.AppendUint32(buf, uint32(len(csr.nodeTypes)))
	buf = append(buf, csr.nodeTypes...)

	buf = binary.BigEndian.AppendUint32(
		buf,
		uint32(len(csr.nodeRecursiveRelations)),
	)
	for _, v := range csr.nodeRecursiveRelations {
		buf = binary.BigEndian.AppendUint32(buf, uint32(v))
	}

	buf = binary.BigEndian.AppendUint32(buf, uint32(len(csr.nodeCycles)))
	for _, v := range csr.nodeCycles {
		buf = binary.BigEndian.AppendUint64(buf, v)
	}

	buf = binary.BigEndian.AppendUint32(
		buf,
		uint32(len(csr.nodeLabelPointers)),
	)
	for _, v := range csr.nodeLabelPointers {
		buf = binary.BigEndian.AppendUint32(buf, uint32(v))
	}

	buf = binary.BigEndian.AppendUint32(buf, uint32(len(csr.nodeLabels)))
	buf = append(buf, csr.nodeLabels...)

	buf = binary.BigEndian.AppendUint32(buf, uint32(len(csr.nodeIDPointers)))
	for _, v := range csr.nodeIDPointers {
		buf = binary.BigEndian.AppendUint32(buf, uint32(v))
	}

	buf = binary.BigEndian.AppendUint32(buf, uint32(len(csr.nodeIDs)))
	buf = append(buf, csr.nodeIDs...)

	buf = binary.BigEndian.AppendUint32(buf, uint32(len(csr.nodeReachabilityPointers)))
	for _, v := range csr.nodeReachabilityPointers {
		buf = binary.BigEndian.AppendUint32(buf, uint32(v))
	}

	buf = binary.BigEndian.AppendUint32(buf, uint32(len(csr.nodeReachability)))
	for _, v := range csr.nodeReachability {
		buf = binary.BigEndian.AppendUint64(buf, v)
	}

	buf = binary.BigEndian.AppendUint32(buf, uint32(len(csr.nodeWildcardPointers)))
	for _, v := range csr.nodeWildcardPointers {
		buf = binary.BigEndian.AppendUint32(buf, uint32(v))
	}

	buf = binary.BigEndian.AppendUint32(buf, uint32(len(csr.nodeWildcards)))
	for _, v := range csr.nodeWildcards {
		buf = binary.BigEndian.AppendUint32(buf, uint32(v))
	}

	buf = binary.BigEndian.AppendUint32(buf, uint32(len(csr.edgePointers)))
	for _, v := range csr.edgePointers {
		buf = binary.BigEndian.AppendUint32(buf, uint32(v))
	}

	buf = binary.BigEndian.AppendUint32(buf, uint32(len(csr.edgeTypes)))
	buf = append(buf, csr.edgeTypes...)

	buf = binary.BigEndian.AppendUint32(buf, uint32(len(csr.edgeEndpoints)))
	for _, v := range csr.edgeEndpoints {
		buf = binary.BigEndian.AppendUint64(buf, v)
	}

	buf = binary.BigEndian.AppendUint32(buf, uint32(len(csr.edgeRecursiveRelations)))
	for _, v := range csr.edgeRecursiveRelations {
		buf = binary.BigEndian.AppendUint32(buf, uint32(v))
	}

	buf = binary.BigEndian.AppendUint32(buf, uint32(len(csr.edgeRelationDefinitions)))
	for _, v := range csr.edgeRelationDefinitions {
		buf = binary.BigEndian.AppendUint32(buf, uint32(v))
	}

	buf = binary.BigEndian.AppendUint32(buf, uint32(len(csr.edgeCycles)))
	for _, v := range csr.edgeCycles {
		buf = binary.BigEndian.AppendUint64(buf, v)
	}

	buf = binary.BigEndian.AppendUint32(buf, uint32(len(csr.edgeReachabilityPointers)))
	for _, v := range csr.edgeReachabilityPointers {
		buf = binary.BigEndian.AppendUint32(buf, uint32(v))
	}

	buf = binary.BigEndian.AppendUint32(buf, uint32(len(csr.edgeReachability)))
	for _, v := range csr.edgeReachability {
		buf = binary.BigEndian.AppendUint64(buf, v)
	}

	buf = binary.BigEndian.AppendUint32(buf, uint32(len(csr.edgeWildcardPointers)))
	for _, v := range csr.edgeWildcardPointers {
		buf = binary.BigEndian.AppendUint32(buf, uint32(v))
	}

	buf = binary.BigEndian.AppendUint32(buf, uint32(len(csr.edgeWildcards)))
	for _, v := range csr.edgeWildcards {
		buf = binary.BigEndian.AppendUint32(buf, uint32(v))
	}

	buf = binary.BigEndian.AppendUint32(buf, uint32(len(csr.edgeConditionPointers)))
	for _, v := range csr.edgeConditionPointers {
		buf = binary.BigEndian.AppendUint32(buf, uint32(v))
	}

	buf = binary.BigEndian.AppendUint32(buf, uint32(len(csr.edgeConditions)))
	buf = append(buf, csr.edgeConditions...)

	return buf, nil
}

func (csr *CSR) UnmarshalBinary(b []byte) error {
	csr.nodes = int32(binary.BigEndian.Uint32(b))
	b = b[4:]
	csr.edges = int32(binary.BigEndian.Uint32(b))
	b = b[4:]

	length := int(binary.BigEndian.Uint32(b))
	b = b[4:]
	csr.nodeTypes = make([]uint8, length)
	for i := range length {
		csr.nodeTypes[i] = b[0]
		b = b[1:]
	}

	length = int(binary.BigEndian.Uint32(b))
	b = b[4:]
	csr.nodeRecursiveRelations = make([]Node, length)
	for i := range length {
		csr.nodeRecursiveRelations[i] = Node(binary.BigEndian.Uint32(b))
		b = b[4:]
	}

	length = int(binary.BigEndian.Uint32(b))
	b = b[4:]
	csr.nodeCycles = make([]uint64, length)
	for i := range length {
		csr.nodeCycles[i] = binary.BigEndian.Uint64(b)
		b = b[8:]
	}

	length = int(binary.BigEndian.Uint32(b))
	b = b[4:]
	csr.nodeLabelPointers = make([]int32, length)
	for i := range length {
		csr.nodeLabelPointers[i] = int32(binary.BigEndian.Uint32(b))
		b = b[4:]
	}

	length = int(binary.BigEndian.Uint32(b))
	b = b[4:]
	csr.nodeLabels = make([]byte, length)
	copy(csr.nodeLabels, b[:length])
	b = b[length:]

	length = int(binary.BigEndian.Uint32(b))
	b = b[4:]
	csr.nodeIDPointers = make([]int32, length)
	for i := range length {
		csr.nodeIDPointers[i] = int32(binary.BigEndian.Uint32(b))
		b = b[4:]
	}

	length = int(binary.BigEndian.Uint32(b))
	b = b[4:]
	csr.nodeIDs = make([]byte, length)
	copy(csr.nodeIDs, b[:length])
	b = b[length:]

	csr.idToNode = make(map[string]int32, csr.nodes)
	for _, node := range csr.Nodes() {
		csr.idToNode[csr.NodeID(node)] = int32(node)
	}

	length = int(binary.BigEndian.Uint32(b))
	b = b[4:]
	csr.nodeReachabilityPointers = make([]int32, length)
	for i := range length {
		csr.nodeReachabilityPointers[i] = int32(binary.BigEndian.Uint32(b))
		b = b[4:]
	}

	length = int(binary.BigEndian.Uint32(b))
	b = b[4:]
	csr.nodeReachability = make([]uint64, length)
	for i := range length {
		r := binary.BigEndian.Uint64(b)
		b = b[8:]
		csr.nodeReachability[i] = r
	}

	length = int(binary.BigEndian.Uint32(b))
	b = b[4:]
	csr.nodeWildcardPointers = make([]int32, length)
	for i := range length {
		csr.nodeWildcardPointers[i] = int32(binary.BigEndian.Uint32(b))
		b = b[4:]
	}

	length = int(binary.BigEndian.Uint32(b[:4]))
	b = b[4:]
	csr.nodeWildcards = make([]Node, length)
	for i := range length {
		csr.nodeWildcards[i] = Node(binary.BigEndian.Uint32(b))
		b = b[4:]
	}

	length = int(binary.BigEndian.Uint32(b))
	b = b[4:]
	csr.edgePointers = make([]int32, length)
	for i := range length {
		csr.edgePointers[i] = int32(binary.BigEndian.Uint32(b))
		b = b[4:]
	}

	length = int(binary.BigEndian.Uint32(b))
	b = b[4:]
	csr.edgeTypes = make([]uint8, length)
	for i := range length {
		csr.edgeTypes[i] = b[0]
		b = b[1:]
	}

	length = int(binary.BigEndian.Uint32(b))
	b = b[4:]
	csr.edgeEndpoints = make([]uint64, length)
	for i := range length {
		csr.edgeEndpoints[i] = binary.BigEndian.Uint64(b)
		b = b[8:]
	}

	length = int(binary.BigEndian.Uint32(b))
	b = b[4:]
	csr.edgeRecursiveRelations = make([]Node, length)
	for i := range length {
		csr.edgeRecursiveRelations[i] = Node(binary.BigEndian.Uint32(b))
		b = b[4:]
	}

	length = int(binary.BigEndian.Uint32(b))
	b = b[4:]
	csr.edgeRelationDefinitions = make([]Node, length)
	for i := range length {
		csr.edgeRelationDefinitions[i] = Node(binary.BigEndian.Uint32(b))
		b = b[4:]
	}

	length = int(binary.BigEndian.Uint32(b))
	b = b[4:]
	csr.edgeCycles = make([]uint64, length)
	for i := range length {
		csr.edgeCycles[i] = binary.BigEndian.Uint64(b)
		b = b[8:]
	}

	length = int(binary.BigEndian.Uint32(b))
	b = b[4:]
	csr.edgeReachabilityPointers = make([]int32, length)
	for i := range length {
		csr.edgeReachabilityPointers[i] = int32(binary.BigEndian.Uint32(b))
		b = b[4:]
	}

	length = int(binary.BigEndian.Uint32(b))
	b = b[4:]
	csr.edgeReachability = make([]uint64, length)
	for i := range length {
		r := binary.BigEndian.Uint64(b)
		b = b[8:]
		csr.edgeReachability[i] = r
	}

	length = int(binary.BigEndian.Uint32(b))
	b = b[4:]
	csr.edgeWildcardPointers = make([]int32, length)
	for i := range length {
		csr.edgeWildcardPointers[i] = int32(binary.BigEndian.Uint32(b))
		b = b[4:]
	}

	length = int(binary.BigEndian.Uint32(b))
	b = b[4:]
	csr.edgeWildcards = make([]Node, length)
	for i := range length {
		csr.edgeWildcards[i] = Node(binary.BigEndian.Uint32(b))
		b = b[4:]
	}

	length = int(binary.BigEndian.Uint32(b))
	b = b[4:]
	csr.edgeConditionPointers = make([]int32, length)
	for i := range length {
		csr.edgeConditionPointers[i] = int32(binary.BigEndian.Uint32(b))
		b = b[4:]
	}

	length = int(binary.BigEndian.Uint32(b))
	b = b[4:]
	csr.edgeConditions = make([]byte, length)
	copy(csr.edgeConditions, b[:length])

	return nil
}
