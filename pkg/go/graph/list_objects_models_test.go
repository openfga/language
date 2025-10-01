package graph

import (
	"fmt"
	language "github.com/openfga/language/pkg/go/transformer"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

func TestThingy(t *testing.T) {
	// This model is just to be sure the logic is correct, we KNOW this model is affected
	// it's from Adrian's test case
	modelFromPr := `
		model
		  schema 1.1
		
		type user
		type user_group
		  relations
			define x: [user]
		type team
		  relations
			define manager: assigned and x from mygroup
			define assigned: [user]
			define mygroup: [user_group]
		type org
		  relations
			define teams: [team]
			define member: manager from teams
	`

	authorizationModel := language.MustTransformDSLToProto(modelFromPr)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)
	require.NotNil(t, graph)
	checkModels()
}

// go through all edges
// if edge is ttu type and the relation definition e.g. `org#member` does not match the .to() relation
// e.g. team#manager
// AND the edge[to()] is pointing TO an intersection or exclusion, flag it and log the relation names for this model
func checkModels() {
	entries, err := os.ReadDir("./models")
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		modelBytes, err := os.ReadFile("./models/" + entry.Name())
		if err != nil {
			panic("line 57 " + err.Error())
		}
		if len(modelBytes) == 0 {
			fmt.Printf("Empty model file %s\n", entry.Name())
			continue
		}
		authorizationModel := language.MustTransformDSLToProto(string(modelBytes))
		wgb := NewWeightedAuthorizationModelGraphBuilder()
		graph, err := wgb.Build(authorizationModel)
		if err != nil {
			fmt.Printf("Error reading model file %s, err: %s\n", entry.Name(), err.Error())
			continue
		}
		checkIfAffectedByListObjectsIssue(graph, entry.Name())
	}
}
func getRelationFromTypeRel(typeRel string) string {
	relation := strings.Split(typeRel, "#")[1]
	return relation
}

// checkIfAffectedByListObjectsIssue checks if this graph has relations which match the following criteria
// 1. have a TTU edge where that TTU relation name does not match the original relation name
// 2. AND that TTU edge leads directly to an intersection or exclusion
// this TTU would be considered affected:
//
//	type user_group
//	  	relations
//			define x: [user]
//	type team
//		relations
//			define manager: assigned and x from mygroup <-- `x from mygroup` not affected, does not lead to intersection
//			define mygroup: [user_group]
//		type org
//		  relations
//			define teams: [team]
//			define member: manager from teams <-- affected, TTU leading direct to AND
func checkIfAffectedByListObjectsIssue(graph *WeightedAuthorizationModelGraph, storeID string) {
	for nodeID, edgeList := range graph.GetEdges() {
		for _, edge := range edgeList {
			if edge.GetEdgeType() == TTUEdge {
				toNode := edge.GetTo()
				if getRelationFromTypeRel(edge.GetRelationDefinition()) != getRelationFromTypeRel(toNode.uniqueLabel) {
					// The relation names mismatch, now we need to see if there's an
					// intersection or exclusion at the next step in the graph
					edgesFromNode, _ := graph.GetEdgesFromNode(toNode)

					// Intersection / Exclusion are always the ONLY edge, and then they can relay to > 1
					if len(edgesFromNode) != 1 {
						continue
					}

					toLabel := edgesFromNode[0].GetTo().GetLabel()
					if toLabel == "intersection" || toLabel == "exclusion" {
						// If we make it here, that means we had a TTU edge whose relation names were different
						// and also led directly to an intersection / exclusion
						fmt.Printf("Possibly affected: store_id: %s, relation: %s\n", storeID, nodeID)
					}
				}
			}
		}
	}
}
