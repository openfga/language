package graph

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/oklog/ulid/v2"
	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/multi"
)

type AuthorizationModelGraphBuilder struct {
	graph.DirectedMultigraphBuilder

	ids map[string]int64 // nodes: unique labels to ids. Used to find nodes by label.
}

// NewAuthorizationModelGraph builds an authorization model in graph form.
// For example, types such as `group`, usersets such as `group#member` and wildcards `group:*` are encoded as nodes.
//
// The edges are defined by the assignments, e.g.
// `define viewer: [group]` defines an edge from group to document#viewer.
// TODO expand when more use cases are added.
func NewAuthorizationModelGraph(model *openfgav1.AuthorizationModel) (*AuthorizationModelGraph, error) {
	res, err := parseModel(model)
	if err != nil {
		return nil, err
	}

	return &AuthorizationModelGraph{res}, nil
}

func parseModel(model *openfgav1.AuthorizationModel) (*multi.DirectedGraph, error) {
	graphBuilder := &AuthorizationModelGraphBuilder{
		multi.NewDirectedGraph(), map[string]int64{},
	}

	// sort types by name to guarantee stable output
	sortedTypeDefs := make([]*openfgav1.TypeDefinition, len(model.GetTypeDefinitions()))
	copy(sortedTypeDefs, model.GetTypeDefinitions())

	slices.SortFunc(sortedTypeDefs, func(a, b *openfgav1.TypeDefinition) int {
		return cmp.Compare(a.GetType(), b.GetType())
	})

	for _, typeDef := range sortedTypeDefs {
		graphBuilder.GetOrAddNode(typeDef.GetType(), typeDef.GetType(), SpecificType)

		// sort relations by name to guarantee stable output
		sortedRelations := make([]string, 0, len(typeDef.GetRelations()))
		for relationName := range typeDef.GetRelations() {
			sortedRelations = append(sortedRelations, relationName)
		}

		slices.Sort(sortedRelations)

		for _, relation := range sortedRelations {
			_ = parseRelation(model, typeDef, relation, graphBuilder)
		}
	}

	multigraph, ok := graphBuilder.DirectedMultigraphBuilder.(*multi.DirectedGraph)
	if ok {
		return multigraph, nil
	}

	return nil, fmt.Errorf("%w: could not cast to directed graph", ErrBuildingGraph)
}

func parseRelation(model *openfgav1.AuthorizationModel, typeDef *openfgav1.TypeDefinition, relation string, graphBuilder *AuthorizationModelGraphBuilder) graph.Node {
	uniqueLabel := fmt.Sprintf("%s#%s", typeDef.GetType(), relation)
	relationNode := graphBuilder.GetOrAddNode(uniqueLabel, uniqueLabel, SpecificTypeAndRelation)
	rewrite := typeDef.GetRelations()[relation]

	switch rewrite.GetUserset().(type) {
	case *openfgav1.Userset_This:
		directlyRelated := make([]*openfgav1.RelationReference, 0)
		if relationMetadata, ok := typeDef.GetMetadata().GetRelations()[relation]; ok {
			directlyRelated = relationMetadata.GetDirectlyRelatedUserTypes()
		}

		_ = parseThis(directlyRelated, graphBuilder, relationNode)
	case *openfgav1.Userset_ComputedUserset:
		return parseComputed(rewrite, typeDef.GetType(), graphBuilder, relationNode)
	case *openfgav1.Userset_TupleToUserset:
		return parseTupleToUserset(model, rewrite, typeDef, graphBuilder, relationNode)
	case *openfgav1.Userset_Intersection:
		children := rewrite.GetIntersection().GetChild()

		return parseOperator("intersection", relationNode, children, typeDef, relation, graphBuilder)
	case *openfgav1.Userset_Union:
		children := rewrite.GetUnion().GetChild()

		return parseOperator("union", relationNode, children, typeDef, relation, graphBuilder)
	default:
		panic("TODO handle difference")
	}

	return nil
}

func parseThis(directlyRelated []*openfgav1.RelationReference, graphBuilder *AuthorizationModelGraphBuilder, thisNode graph.Node) []graph.Node {
	newNodes := make([]graph.Node, 0, len(directlyRelated))

	for _, directlyRelatedDef := range directlyRelated {
		if directlyRelatedDef.GetRelationOrWildcard() == nil {
			// direct assignment to concrete type
			assignableType := directlyRelatedDef.GetType()
			newNodes = append(newNodes, graphBuilder.GetOrAddNode(assignableType, assignableType, SpecificType))
		}

		if directlyRelatedDef.GetWildcard() != nil {
			// direct assignment to wildcard
			assignableWildcard := directlyRelatedDef.GetType() + ":*"
			newNodes = append(newNodes, graphBuilder.GetOrAddNode(assignableWildcard, assignableWildcard, SpecificType))
		}

		if directlyRelatedDef.GetRelation() != "" {
			// direct assignment to userset
			assignableUserset := directlyRelatedDef.GetType() + "#" + directlyRelatedDef.GetRelation()
			newNodes = append(newNodes, graphBuilder.GetOrAddNode(assignableUserset, assignableUserset, SpecificTypeAndRelation))
		}
	}

	for _, newNode := range newNodes {
		graphBuilder.AddEdge(newNode, thisNode, DirectEdge, "")
	}

	return newNodes
}

func parseComputed(rewrite *openfgav1.Userset, typeName string, graphBuilder *AuthorizationModelGraphBuilder, relationNode graph.Node) graph.Node {
	// e.g. define x: y. Here y is the rewritten relation
	rewrittenRelation := rewrite.GetComputedUserset().GetRelation()
	rewrittenNodeName := fmt.Sprintf("%s#%s", typeName, rewrittenRelation)
	newNode := graphBuilder.GetOrAddNode(rewrittenNodeName, rewrittenNodeName, SpecificTypeAndRelation)
	// new edge from y to x
	graphBuilder.AddEdge(newNode, relationNode, RewriteEdge, "")

	return newNode
}

// parseOperator creates as many intermediate nodes as necessary for each composition of operators.
func parseOperator(operator string, operationNode graph.Node, childrenOfOperator []*openfgav1.Userset, typeDefOfOperator *openfgav1.TypeDefinition, relationNameOfOperator string, graphBuilder *AuthorizationModelGraphBuilder) graph.Node {
	uniqueIdentifierOfIntermediateNode := fmt.Sprintf("%s:%s", operator, ulid.Make().String())

	intermediateOperatorNode := graphBuilder.GetOrAddNode(uniqueIdentifierOfIntermediateNode, operator, OperatorNode)

	// add one edge "operator" -> "relation that defined the operator"
	// Note: if this is a composition of operators, operationNode will be nil and this edge won't be added.
	graphBuilder.AddEdge(intermediateOperatorNode, operationNode, RewriteEdge, "")

	// add one edge per children: "child" -> "operator"
	for _, child := range childrenOfOperator {
		switch child.GetUserset().(type) {
		case *openfgav1.Userset_This:
			directlyRelated := make([]*openfgav1.RelationReference, 0)
			if relationMetadata, ok := typeDefOfOperator.GetMetadata().GetRelations()[relationNameOfOperator]; ok {
				directlyRelated = relationMetadata.GetDirectlyRelatedUserTypes()
			}
			// e.g. `define: [user] and viewer` will add edges `user -> intersection`, and `document#viewer -> intersection`:
			for _, childNode := range parseThis(directlyRelated, graphBuilder, nil) {
				graphBuilder.AddEdge(childNode, intermediateOperatorNode, DirectEdge, "")
			}
		case *openfgav1.Userset_ComputedUserset:
			childNode := parseComputed(child, typeDefOfOperator.GetType(), graphBuilder, nil)
			graphBuilder.AddEdge(childNode, intermediateOperatorNode, RewriteEdge, "")
		case *openfgav1.Userset_Intersection:
			childrenOfIntersection := child.GetIntersection().GetChild()
			childNode := parseOperator("intersection", nil, childrenOfIntersection, typeDefOfOperator, relationNameOfOperator, graphBuilder)
			graphBuilder.AddEdge(childNode, intermediateOperatorNode, RewriteEdge, "")
		case *openfgav1.Userset_Union:
			childrenOfUnion := child.GetUnion().GetChild()
			childNode := parseOperator("union", nil, childrenOfUnion, typeDefOfOperator, relationNameOfOperator, graphBuilder)
			graphBuilder.AddEdge(childNode, intermediateOperatorNode, RewriteEdge, "")
		default:
			panic("TODO handle TTU and difference")
		}
	}

	return intermediateOperatorNode
}

func parseTupleToUserset(model *openfgav1.AuthorizationModel, rewrite *openfgav1.Userset, typeDef *openfgav1.TypeDefinition, graphBuilder *AuthorizationModelGraphBuilder, nodeTarget graph.Node) graph.Node {
	// e.g. define viewer: admin from parent
	// "parent" is the tupleset
	tuplesetRelation := rewrite.GetTupleToUserset().GetTupleset().GetRelation()
	// "admin" is the computed relation
	computedRelation := rewrite.GetTupleToUserset().GetComputedUserset().GetRelation()

	// find all the directly related types to the tupleset
	directlyRelated := make([]*openfgav1.RelationReference, 0)
	if relationMetadata, ok := typeDef.GetMetadata().GetRelations()[tuplesetRelation]; ok {
		directlyRelated = relationMetadata.GetDirectlyRelatedUserTypes()
	}

	for _, relatedType := range directlyRelated {
		tuplesetType := relatedType.GetType()

		if !typeAndRelationExists(model, tuplesetType, computedRelation) {
			continue
		}

		rewrittenNodeName := fmt.Sprintf("%s#%s", tuplesetType, computedRelation)
		nodeSource := graphBuilder.GetOrAddNode(rewrittenNodeName, rewrittenNodeName, SpecificTypeAndRelation)

		conditionedOnNodeName := fmt.Sprintf("(%s#%s)", typeDef.GetType(), tuplesetRelation)

		// new edge from "xxx#admin" to "yyy#viewer" conditioned on "yyy#parent"
		graphBuilder.AddEdge(nodeSource, nodeTarget, TTUEdge, conditionedOnNodeName)
	}

	return nodeTarget
}

func (g *AuthorizationModelGraphBuilder) GetOrAddNode(uniqueLabel, label string, nodeType NodeType) *AuthorizationModelNode {
	if existingNode := g.GetNodeFor(uniqueLabel); existingNode != nil {
		return existingNode
	}

	node := g.NewNode()
	nodeid := node.ID()
	newNode := &AuthorizationModelNode{
		Node:        node,
		label:       label,
		nodeType:    nodeType,
		uniqueLabel: uniqueLabel,
	}
	g.AddNode(newNode)
	g.ids[uniqueLabel] = nodeid

	return newNode
}

func (g *AuthorizationModelGraphBuilder) GetNodeFor(uniqueLabel string) *AuthorizationModelNode {
	id, ok := g.ids[uniqueLabel]
	if !ok {
		return nil
	}

	authModelNode, ok := g.Node(id).(*AuthorizationModelNode)
	if !ok {
		return nil
	}

	return authModelNode
}

func (g *AuthorizationModelGraphBuilder) AddEdge(from, to graph.Node, edgeType EdgeType, conditionedOn string) *AuthorizationModelEdge {
	if from == nil || to == nil {
		return nil
	}

	l := g.NewLine(from, to)
	newLine := &AuthorizationModelEdge{Line: l, edgeType: edgeType, conditionedOn: conditionedOn}
	g.SetLine(newLine)

	return newLine
}

func typeAndRelationExists(model *openfgav1.AuthorizationModel, typeName, relation string) bool {
	typeDefs := model.GetTypeDefinitions()
	// TODO this should be made faster, ideally typeDefs is a map
	for _, typeDef := range typeDefs {
		if typeDef.GetType() == typeName {
			relations := typeDef.GetRelations()
			if _, ok := relations[relation]; ok {
				return true
			}
		}
	}

	return false
}
