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
// Conditions are not encoded in the graph,
// and the two edges in an exclusion are not distinguished.
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
			uniqueLabel := fmt.Sprintf("%s#%s", typeDef.GetType(), relation)
			parentNode := graphBuilder.GetOrAddNode(uniqueLabel, uniqueLabel, SpecificTypeAndRelation)
			rewrite := typeDef.GetRelations()[relation]
			checkRewrite(graphBuilder, parentNode, model, rewrite, typeDef, relation)
		}
	}

	multigraph, ok := graphBuilder.DirectedMultigraphBuilder.(*multi.DirectedGraph)
	if ok {
		return multigraph, nil
	}

	return nil, fmt.Errorf("%w: could not cast to directed graph", ErrBuildingGraph)
}

func checkRewrite(graphBuilder *AuthorizationModelGraphBuilder, parentNode *AuthorizationModelNode, model *openfgav1.AuthorizationModel, rewrite *openfgav1.Userset, typeDef *openfgav1.TypeDefinition, relation string) {
	var operator string

	var children []*openfgav1.Userset

	switch rw := rewrite.GetUserset().(type) {
	case *openfgav1.Userset_This:
		parseThis(graphBuilder, parentNode, typeDef, relation)

		return
	case *openfgav1.Userset_ComputedUserset:
		parseComputed(graphBuilder, parentNode, typeDef, rw.ComputedUserset.GetRelation())

		return
	case *openfgav1.Userset_TupleToUserset:
		parseTupleToUserset(graphBuilder, parentNode, model, typeDef, rw.TupleToUserset)

		return
	case *openfgav1.Userset_Union:
		operator = "union"
		children = rw.Union.GetChild()

	case *openfgav1.Userset_Intersection:
		operator = "intersection"
		children = rw.Intersection.GetChild()

	case *openfgav1.Userset_Difference:
		operator = "exclusion"
		children = []*openfgav1.Userset{
			rw.Difference.GetBase(),
			rw.Difference.GetSubtract(),
		}
	}

	operatorNode := fmt.Sprintf("%s:%s", operator, ulid.Make().String())
	operatorNodeParent := graphBuilder.GetOrAddNode(operatorNode, operator, OperatorNode)

	// add one edge "operator" -> "relation that defined the operator"
	// Note: if this is a composition of operators, operationNode will be nil and this edge won't be added.
	graphBuilder.AddEdge(operatorNodeParent, parentNode, RewriteEdge, "")
	for _, child := range children {
		checkRewrite(graphBuilder, operatorNodeParent, model, child, typeDef, relation)
	}
}

func parseThis(graphBuilder *AuthorizationModelGraphBuilder, parentNode graph.Node, typeDef *openfgav1.TypeDefinition, relation string) {
	directlyRelated := make([]*openfgav1.RelationReference, 0)
	var curNode *AuthorizationModelNode

	if relationMetadata, ok := typeDef.GetMetadata().GetRelations()[relation]; ok {
		directlyRelated = relationMetadata.GetDirectlyRelatedUserTypes()
	}

	for _, directlyRelatedDef := range directlyRelated {
		conditionName := directlyRelatedDef.GetCondition()
		assignableType := directlyRelatedDef.GetType()
		if conditionName != "" {
			assignableType = fmt.Sprintf("%s [with %s]", directlyRelatedDef.GetType(), conditionName)
		}
		if directlyRelatedDef.GetRelationOrWildcard() == nil {
			// direct assignment to concrete type
			curNode = graphBuilder.GetOrAddNode(assignableType, assignableType, SpecificType)
		}

		if directlyRelatedDef.GetWildcard() != nil {
			// direct assignment to wildcard
			assignableWildcard := assignableType + ":*"
			curNode = graphBuilder.GetOrAddNode(assignableWildcard, assignableWildcard, SpecificType)
		}

		if directlyRelatedDef.GetRelation() != "" {
			// direct assignment to userset
			assignableUserset := assignableType + "#" + directlyRelatedDef.GetRelation()
			curNode = graphBuilder.GetOrAddNode(assignableUserset, assignableUserset, SpecificTypeAndRelation)
		}

		if graphBuilder.HasEdge(curNode, parentNode, DirectEdge, "") {
			// de-dup types that are conditioned, e.g. if define viewer: [user, user with condX]
			// we only draw one edge from user to x#viewer
			continue
		}

		graphBuilder.AddEdge(curNode, parentNode, DirectEdge, "")
	}
}

func parseComputed(graphBuilder *AuthorizationModelGraphBuilder, parentNode graph.Node, typeDef *openfgav1.TypeDefinition, relation string) {
	// e.g. define x: y. Here y is the rewritten relation
	rewrittenNodeName := fmt.Sprintf("%s#%s", typeDef.GetType(), relation)
	newNode := graphBuilder.GetOrAddNode(rewrittenNodeName, rewrittenNodeName, TTUNode)
	// new edge from y to x
	graphBuilder.AddEdge(newNode, parentNode, RewriteEdge, "")
}

func parseTupleToUserset(graphBuilder *AuthorizationModelGraphBuilder, parentNode graph.Node, model *openfgav1.AuthorizationModel, typeDef *openfgav1.TypeDefinition, rewrite *openfgav1.TupleToUserset) {
	// e.g. define viewer: admin from parent
	// "parent" is the tupleset
	tuplesetRelation := rewrite.GetTupleset().GetRelation()
	// "admin" is the computed relation
	computedRelation := rewrite.GetComputedUserset().GetRelation()

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

		typeee := tuplesetType
		if relatedType.GetCondition() != "" {
			typeee = fmt.Sprintf("%s [with %s]", tuplesetType, relatedType.GetCondition())
		}
		rewrittenNodeName := fmt.Sprintf("%s#%s", typeee, computedRelation)
		nodeSource := graphBuilder.GetOrAddNode(rewrittenNodeName, rewrittenNodeName, TTUNode)
		conditionedOnNodeName := fmt.Sprintf("(if %s is %s of %s)", typeee, tuplesetRelation, typeDef.GetType())

		// new edge from "xxx#admin" to "yyy#viewer" conditioned on "yyy#parent"
		graphBuilder.AddEdge(nodeSource, parentNode, TTUEdge, conditionedOnNodeName)
	}
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

	l := g.NewLine(to, from)
	newLine := &AuthorizationModelEdge{Line: l, edgeType: edgeType, conditionedOn: conditionedOn}
	g.SetLine(newLine)

	return newLine
}

func (g *AuthorizationModelGraphBuilder) HasEdge(from, to graph.Node, edgeType EdgeType, conditionedOn string) bool {
	if from == nil || to == nil {
		return false
	}

	iter := g.Lines(from.ID(), to.ID())
	for {
		if !iter.Next() {
			return false
		}
		l := iter.Line()
		edge, ok := l.(*AuthorizationModelEdge)
		if !ok {
			return false
		}
		if edge.edgeType == edgeType && edge.conditionedOn == conditionedOn {
			return true
		}
	}
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
