package graph

import (
	"cmp"
	"fmt"
	"slices"

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
			uniqueLabel := fmt.Sprintf("%s#%s", typeDef.GetType(), relation)
			relationNode := graphBuilder.GetOrAddNode(uniqueLabel, uniqueLabel, SpecificTypeAndRelation)

			rewrite := typeDef.GetRelations()[relation]
			if _, ok := rewrite.GetUserset().(*openfgav1.Userset_This); ok {
				directlyRelated := make([]*openfgav1.RelationReference, 0)
				if metadata, ok := typeDef.GetMetadata().GetRelations()[relation]; ok {
					directlyRelated = metadata.GetDirectlyRelatedUserTypes()
				}

				for _, directlyRelatedDef := range directlyRelated {
					assignableType := directlyRelatedDef.GetType()

					newNode := graphBuilder.GetOrAddNode(assignableType, assignableType, SpecificType)
					graphBuilder.AddEdge(newNode, relationNode, DirectEdge)
				}
			}
		}
	}

	multigraph, ok := graphBuilder.DirectedMultigraphBuilder.(*multi.DirectedGraph)
	if ok {
		return multigraph, nil
	}

	return nil, fmt.Errorf("%w: could not cast to directed graph", ErrBuildingGraph)
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

func (g *AuthorizationModelGraphBuilder) AddEdge(from, to graph.Node, edgeType EdgeType) *AuthorizationModelEdge {
	l := g.NewLine(from, to)
	newLine := &AuthorizationModelEdge{Line: l, edgeType: edgeType}
	g.SetLine(newLine)

	return newLine
}
