package graph

import (
	"cmp"
	"fmt"
	"reflect"
	"slices"

	"github.com/oklog/ulid/v2"
	openfgav1 "github.com/openfga/api/proto/openfga/v1"

	"github.com/openfga/language/pkg/go/errors"
)

type WeightedAuthorizationModelGraphBuilder struct {
}

func NewWeightedAuthorizationModelGraphBuilder() *WeightedAuthorizationModelGraphBuilder {
	return &WeightedAuthorizationModelGraphBuilder{}
}

func (wgb *WeightedAuthorizationModelGraphBuilder) ValidateAndBuild(model *openfgav1.AuthorizationModel) (*WeightedAuthorizationModelGraph, error) {
	wb := NewWeightedAuthorizationModelGraph()

	typeDefinitions := model.GetTypeDefinitions()
	// Sort types by name to guarantee stable output - do this once
	sortedTypeDefs := make([]*openfgav1.TypeDefinition, len(typeDefinitions))
	copy(sortedTypeDefs, typeDefinitions)
	slices.SortFunc(sortedTypeDefs, func(a, b *openfgav1.TypeDefinition) int {
		return cmp.Compare(a.GetType(), b.GetType())
	})

	// Prepare sorted relations for each type - do this once
	type typeWithRelations struct {
		typeDef         *openfgav1.TypeDefinition
		sortedRelations []string
	}

	typesWithRelations := make([]typeWithRelations, len(sortedTypeDefs))
	for i, typeDef := range sortedTypeDefs {
		// validate type definition
		err := wgb.validateTypeDefinition(typeDef, wb)
		if err != nil {
			return nil, err
		}

		typeRelations := typeDef.GetRelations()
		// Sort relations by name to guarantee stable output
		sortedRelations := make([]string, 0, len(typeRelations))
		for relationName := range typeRelations {
			sortedRelations = append(sortedRelations, relationName)
		}

		slices.Sort(sortedRelations)
		typesWithRelations[i] = typeWithRelations{
			typeDef:         typeDef,
			sortedRelations: sortedRelations,
		}

		for _, relation := range sortedRelations {

			err := wgb.validateRelation(typeDef.GetType(), relation, wb)
			if err != nil {
				return nil, err
			}
		}
	}

	// Extract condition names as a map for efficient lookup
	conditionUsage := make(map[string]bool, len(model.GetConditions()))
	for conditionName := range model.GetConditions() {
		conditionUsage[conditionName] = false
	}

	// Second pass: Process rewrites for each relation
	for _, typeWithRel := range typesWithRelations {
		typeDef := typeWithRel.typeDef

		for _, relation := range typeWithRel.sortedRelations {
			uniqueLabel := typeDef.GetType() + "#" + relation

			// Get the already created node
			parentNode, found := wb.GetNodeByID(uniqueLabel)
			if !found {

				return nil, errors.RelationObjectTypeError(relation, typeDef.GetType(), errors.ErrRelationUndefined)
			}

			// Parse the rewrite definition
			rewrite := typeDef.GetRelations()[relation]

			err := wgb.parseRewrite(wb, parentNode, model, rewrite, typeDef, relation, conditionUsage, true)
			if err != nil {
				return nil, err
			}
		}
	}

	// iterate over all conditions and mark their usage
	for conditionName := range model.GetConditions() {
		if !conditionUsage[conditionName] {
			return nil, errors.ConditionError(conditionName, errors.ErrConditionUnReferenced)
		}
	}

	// Assign weights after all nodes and edges are created
	err := wb.AssignWeights()
	if err != nil {
		return nil, err
	}

	return wb, nil
}

func (wgb *WeightedAuthorizationModelGraphBuilder) Build(model *openfgav1.AuthorizationModel) (*WeightedAuthorizationModelGraph, error) {
	wb := NewWeightedAuthorizationModelGraph()
	// sort types by name to guarantee stable output
	sortedTypeDefs := make([]*openfgav1.TypeDefinition, len(model.GetTypeDefinitions()))
	copy(sortedTypeDefs, model.GetTypeDefinitions())

	slices.SortFunc(sortedTypeDefs, func(a, b *openfgav1.TypeDefinition) int {
		return cmp.Compare(a.GetType(), b.GetType())
	})

	for _, typeDef := range sortedTypeDefs {
		wb.GetOrAddNode(typeDef.GetType(), typeDef.GetType(), SpecificType, "")

		// sort relations by name to guarantee stable output
		sortedRelations := make([]string, 0, len(typeDef.GetRelations()))
		for relationName := range typeDef.GetRelations() {
			sortedRelations = append(sortedRelations, relationName)
		}

		slices.Sort(sortedRelations)

		for _, relation := range sortedRelations {
			uniqueLabel := typeDef.GetType() + "#" + relation
			parentNode := wb.GetOrAddNode(uniqueLabel, uniqueLabel, SpecificTypeAndRelation, uniqueLabel)

			rewrite := typeDef.GetRelations()[relation]
			err := wgb.parseRewrite(wb, parentNode, model, rewrite, typeDef, relation, nil, false)
			if err != nil {
				return nil, err
			}
		}
	}

	err := wb.AssignWeights()
	if err != nil {
		return nil, err
	}

	return wb, nil
}

func (wgb *WeightedAuthorizationModelGraphBuilder) validateTypeDefinition(typeDef *openfgav1.TypeDefinition, wb *WeightedAuthorizationModelGraph) error {
	objectType := typeDef.GetType()
	if objectType == "" {
		return errors.InvalidAuthorizationModelError(errors.ErrInvalidType)
	}

	if objectType == "self" || objectType == "this" {
		return errors.InvalidAuthorizationModelError(errors.ErrReservedKeywords) // fmt.Errorf("the definition of type '%s' is invalid", objectType)
	}

	// If a type was already defined it will fail
	_, err := wb.AddNode(objectType, objectType, SpecificType, "")
	if err != nil {
		return errors.ObjectTypeError(objectType, errors.ErrDuplicateTypes)
	}

	return nil
}

func (wgb *WeightedAuthorizationModelGraphBuilder) validateRelation(objectType string, relation string, wb *WeightedAuthorizationModelGraph) error {
	if relation == "" {
		return errors.ObjectTypeError(objectType, errors.ErrInvalidRelation)
	}

	if relation == "self" || relation == "this" {
		return errors.ObjectTypeError(objectType, errors.ErrReservedKeywords)
	}

	uniqueLabel := objectType + "#" + relation
	// If a relation was already defined for the type it will fail
	_, err := wb.AddNode(uniqueLabel, uniqueLabel, SpecificTypeAndRelation, uniqueLabel)
	if err != nil {
		return errors.RelationObjectTypeError(relation, objectType, errors.ErrDuplicateRelationsType)
	}

	return nil
}

func (wgb *WeightedAuthorizationModelGraphBuilder) parseRewrite(wg *WeightedAuthorizationModelGraph, parentNode *WeightedAuthorizationModelNode, model *openfgav1.AuthorizationModel, rewrite *openfgav1.Userset, typeDef *openfgav1.TypeDefinition, relation string, conditions map[string]bool, withValidation bool) error {
	var operator string
	relationNodeName := GetRelationNodeName(typeDef.GetType(), relation)

	if withValidation {
		err := wgb.validateRewrite(rewrite, typeDef.GetType(), relation, wg, typeDef, conditions)
		if err != nil {
			return err
		}
	}

	var children []*openfgav1.Userset

	switch rw := rewrite.GetUserset().(type) {
	case *openfgav1.Userset_This:
		return wgb.parseThis(wg, parentNode, typeDef, relation, relationNodeName)
	case *openfgav1.Userset_ComputedUserset:
		wgb.parseComputed(wg, parentNode, typeDef, rw.ComputedUserset.GetRelation(), relationNodeName)
		return nil
	case *openfgav1.Userset_TupleToUserset:
		return wgb.parseTupleToUserset(wg, parentNode, model, typeDef, rw.TupleToUserset, relationNodeName)
	case *openfgav1.Userset_Union:
		operator = UnionOperator
		children = rw.Union.GetChild()

	case *openfgav1.Userset_Intersection:
		operator = IntersectionOperator
		children = rw.Intersection.GetChild()

	case *openfgav1.Userset_Difference:
		operator = ExclusionOperator
		children = []*openfgav1.Userset{
			rw.Difference.GetBase(),
			rw.Difference.GetSubtract(),
		}
	}

	operatorNodeName := operator + ":" + ulid.Make().String()
	operatorNode := wg.GetOrAddNode(operatorNodeName, operator, OperatorNode, relationNodeName)

	// add one edge "relation" -> "operation that defined the operator"
	// Note: if this is a composition of operators, operationNode will be nil and this edge won't be added.
	wg.AddEdge(parentNode.GetUniqueLabel(), operatorNodeName, RewriteEdge, relationNodeName, "", nil)
	for _, child := range children {
		err := wgb.parseRewrite(wg, operatorNode, model, child, typeDef, relation, conditions, withValidation)
		if err != nil {
			return err
		}
	}
	return nil
}

func (wgb *WeightedAuthorizationModelGraphBuilder) validateComputedRewrite(rewrite *openfgav1.Userset_ComputedUserset, objectType string, relation string, wg *WeightedAuthorizationModelGraph) error {
	computedUserset := rewrite.ComputedUserset.GetRelation()
	if computedUserset == relation {
		return errors.RelationObjectTypeError(relation, objectType, errors.ErrInvalidUsersetRewrite)
	}

	if err := wgb.ValidateRelationExists(objectType, computedUserset, wg); err != nil {
		return err
	}

	return nil
}

func (wgb *WeightedAuthorizationModelGraphBuilder) validateTupleToUsersetRewrite(rewrite *openfgav1.Userset_TupleToUserset, objectType string, relation string, wg *WeightedAuthorizationModelGraph, typeDef *openfgav1.TypeDefinition) error {
	tupleset := rewrite.TupleToUserset.GetTupleset().GetRelation()
	if err := wgb.ValidateRelationExists(objectType, tupleset, wg); err != nil {
		return err
	}

	tuplesetRewrite := typeDef.Relations[tupleset]
	// Tupleset relations must only be direct relationships, no rewrites are allowed on them.
	if reflect.TypeOf(tuplesetRewrite.GetUserset()) != reflect.TypeOf(&openfgav1.Userset_This{}) {
		return errors.RelationObjectTypeError(tupleset, objectType, errors.ErrInvalidRelationOnTupleset)
	}

	relationMetadata, ok := typeDef.GetMetadata().GetRelations()[tupleset]
	if !ok {
		return errors.RelationObjectTypeError(tupleset, objectType, errors.ErrRelationUndefined)
	}

	directRelatedUserTypes := relationMetadata.GetDirectlyRelatedUserTypes()
	if len(directRelatedUserTypes) == 0 {
		return errors.RelationObjectTypeError(tupleset, objectType, errors.ErrInvalidRelationOnTupleset)
	}

	for _, directlyRelatedDef := range directRelatedUserTypes {
		if directlyRelatedDef.GetRelationOrWildcard() != nil {
			// Tupleset relations must only be direct relationships, no rewrites are allowed on them.
			return errors.RelationObjectTypeError(tupleset, objectType, errors.ErrInvalidRelationOnTupleset)
		}
	}

	computedRelation := rewrite.TupleToUserset.GetComputedUserset().GetRelation()
	for _, relatedType := range directRelatedUserTypes {
		tuplesetType := relatedType.GetType()
		if err := wgb.ValidateRelationExists(tuplesetType, computedRelation, wg); err != nil {
			return err
		}
	}
	return nil
}

func (wgb *WeightedAuthorizationModelGraphBuilder) validateThisRewrite(objectType string, relation string, wg *WeightedAuthorizationModelGraph, typeDef *openfgav1.TypeDefinition, conditions map[string]bool) error {
	relationMetadata, ok := typeDef.GetMetadata().GetRelations()[relation]
	if !ok {
		return errors.RelationObjectTypeError(relation, objectType, errors.ErrRelationUndefined)
	}
	directRelatedUserTypes := relationMetadata.GetDirectlyRelatedUserTypes()
	if len(directRelatedUserTypes) == 0 {
		return errors.RelationObjectTypeError(relation, objectType, errors.ErrDirectlyAssignableRelation)
	}

	for _, directlyRelatedDef := range directRelatedUserTypes {

		if directlyRelatedDef.GetRelationOrWildcard() == nil {
			// direct assignment to concrete type
			assignableType := directlyRelatedDef.GetType()
			_, ok = wg.GetNodeByID(assignableType)
			if !ok {
				return errors.ObjectTypeError(assignableType, errors.ErrObjectTypeUndefined)
			}
		} else {
			if directlyRelatedDef.GetWildcard() != nil && len(directlyRelatedDef.GetRelation()) > 0 {
				return errors.RelationObjectTypeError(relation, objectType, errors.ErrInvalidWildcard)
			}
			if directlyRelatedDef.GetWildcard() != nil {
				assignableType := directlyRelatedDef.GetType()
				_, ok = wg.GetNodeByID(assignableType)
				if !ok {
					return errors.ObjectTypeError(assignableType, errors.ErrObjectTypeUndefined)
				}
			} else {
				assignableUserset := directlyRelatedDef.GetType() + "#" + directlyRelatedDef.GetRelation()
				_, ok = wg.GetNodeByID(assignableUserset)
				if !ok {
					return errors.RelationObjectTypeError(relation, objectType, errors.ErrInvalidUsersetRewrite)
				}
			}
		}
		if directlyRelatedDef.GetCondition() != "" {
			// validate the condition that directlyRelatedDef.GetCondition() is contained in conditions
			conditionName := directlyRelatedDef.GetCondition()
			if !conditions[conditionName] {
				return errors.ConditionRelationObjectTypeError(conditionName, relation, objectType, errors.ErrCondition)
			}
			// mark condition as referenced
			conditions[conditionName] = true
		}
	}
	return nil
}

func (wgb *WeightedAuthorizationModelGraphBuilder) validateRewrite(rewrite *openfgav1.Userset, objectType string, relation string, wg *WeightedAuthorizationModelGraph, typeDef *openfgav1.TypeDefinition, conditions map[string]bool) error {
	if rewrite.GetUserset() == nil {
		return errors.RelationObjectTypeError(relation, objectType, errors.ErrRelationUndefined)
	}

	switch r := rewrite.GetUserset().(type) {
	case *openfgav1.Userset_This:
		return wgb.validateThisRewrite(objectType, relation, wg, typeDef, conditions)
	case *openfgav1.Userset_ComputedUserset:
		return wgb.validateComputedRewrite(r, objectType, relation, wg)
	case *openfgav1.Userset_TupleToUserset:
		return wgb.validateTupleToUsersetRewrite(r, objectType, relation, wg, typeDef)
	case *openfgav1.Userset_Union:
		for _, child := range r.Union.GetChild() {
			err := wgb.validateRewrite(child, objectType, relation, wg, typeDef, conditions)
			if err != nil {
				return err
			}
		}
	case *openfgav1.Userset_Intersection:
		for _, child := range r.Intersection.GetChild() {
			err := wgb.validateRewrite(child, objectType, relation, wg, typeDef, conditions)
			if err != nil {
				return err
			}
		}
	case *openfgav1.Userset_Difference:
		err := wgb.validateRewrite(r.Difference.GetBase(), objectType, relation, wg, typeDef, conditions)
		if err != nil {
			return err
		}

		err = wgb.validateRewrite(r.Difference.GetSubtract(), objectType, relation, wg, typeDef, conditions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (wgb *WeightedAuthorizationModelGraphBuilder) parseTupleToUserset(wg *WeightedAuthorizationModelGraph, parentNode *WeightedAuthorizationModelNode, model *openfgav1.AuthorizationModel, typeDef *openfgav1.TypeDefinition, rewrite *openfgav1.TupleToUserset, parentRelationName string) error {
	// e.g. define viewer: admin from parent
	// "parent" is the tupleset
	tuplesetRelation := rewrite.GetTupleset().GetRelation()
	// "admin" is the computed relation
	computedRelation := rewrite.GetComputedUserset().GetRelation()

	// find all the directly related types to the tupleset
	relationMetadata, ok := typeDef.GetMetadata().GetRelations()[tuplesetRelation]
	if !ok {
		return errors.RelationObjectTypeError(tuplesetRelation, typeDef.GetType(), errors.ErrRelationUndefined)
	}
	directlyRelated := relationMetadata.GetDirectlyRelatedUserTypes()
	if len(directlyRelated) == 0 {
		return errors.RelationObjectTypeError(tuplesetRelation, typeDef.GetType(), errors.ErrInvalidRelationOnTupleset)
	}

	typeTuplesetRelation := typeDef.GetType() + "#" + tuplesetRelation
	node := parentNode
	if parentNode.nodeType != SpecificTypeAndRelation && len(directlyRelated) > 1 {
		uniqueLabel := typeDef.GetType() + "#ttu:" + tuplesetRelation + "#" + computedRelation
		// add a logical ttu node for grouping of TTU that are part of the same tuplesetrelation and computed relation
		logicalNode := wg.GetOrAddNode(uniqueLabel, uniqueLabel, LogicalTTUGrouping, parentRelationName)
		wg.AddEdge(parentNode.uniqueLabel, logicalNode.uniqueLabel, TTULogicalEdge, parentRelationName, typeTuplesetRelation, nil)
		node = logicalNode
	}

	for _, relatedType := range directlyRelated {
		tuplesetType := relatedType.GetType()

		if !typeAndRelationExists(model, tuplesetType, computedRelation) {
			return errors.RelationObjectTypeError(computedRelation, tuplesetType, errors.ErrRelationUndefined)
		}

		rewrittenNodeName := GetRelationNodeName(tuplesetType, computedRelation)
		nodeSource := wg.GetOrAddNode(rewrittenNodeName, rewrittenNodeName, SpecificTypeAndRelation, rewrittenNodeName)

		if wg.HasEdge(node, nodeSource, TTUEdge, typeTuplesetRelation) {
			// we don't need to do any condition update, only de-dup the edge. In case of TTU
			// the direct relation will have the conditions
			// for example, in the case of
			// type group
			//   relations
			// 		define rel1: [user] or rel1 from parent
			//		define parent: [group, group with condX]
			// In the graph we only have one TTU edge from the OR node to the group#rel1 node, but there are no conditions associated to it
			// the conditions are associated to the edge from group#parent node to the group node. This direct edge has two conditions: none and condX
			continue
		}

		// new edge from "xxx#admin" to "yyy#viewer" tuplesetRelation on "yyy#parent"
		wg.UpsertEdge(node, nodeSource, TTUEdge, parentRelationName, typeTuplesetRelation, relatedType.GetCondition())
	}
	return nil
}

func (wgb *WeightedAuthorizationModelGraphBuilder) parseComputed(wg *WeightedAuthorizationModelGraph, parentNode *WeightedAuthorizationModelNode, typeDef *openfgav1.TypeDefinition, relation string, parentRelationName string) {
	nodeType := RewriteEdge
	// e.g. define x: y. Here y is the rewritten relation
	rewrittenNodeName := typeDef.GetType() + "#" + relation
	newNode := wg.GetOrAddNode(rewrittenNodeName, rewrittenNodeName, SpecificTypeAndRelation, rewrittenNodeName)
	// new edge from x to y
	if parentNode.nodeType == SpecificTypeAndRelation && newNode.nodeType == SpecificTypeAndRelation {
		nodeType = ComputedEdge
	}
	wg.AddEdge(parentNode.uniqueLabel, newNode.uniqueLabel, nodeType, parentRelationName, "", nil)
}

func (wgb *WeightedAuthorizationModelGraphBuilder) parseThis(wg *WeightedAuthorizationModelGraph, parentNode *WeightedAuthorizationModelNode, typeDef *openfgav1.TypeDefinition, relation string, parentRelationName string) error {
	var directlyRelated []*openfgav1.RelationReference
	var curNode *WeightedAuthorizationModelNode

	if relationMetadata, ok := typeDef.GetMetadata().GetRelations()[relation]; ok {
		directlyRelated = relationMetadata.GetDirectlyRelatedUserTypes()
	}
	node := parentNode
	// add a logical userset node for grouping of direct usersets that are defined in the same relation
	if parentNode.nodeType != SpecificTypeAndRelation && len(directlyRelated) > 1 {
		uniqueLabel := typeDef.GetType() + "#direct:" + relation
		logicalNode := wg.GetOrAddNode(uniqueLabel, uniqueLabel, LogicalDirectGrouping, parentRelationName)
		wg.AddEdge(parentNode.uniqueLabel, logicalNode.uniqueLabel, DirectLogicalEdge, parentRelationName, "", nil)
		node = logicalNode
	}

	for _, directlyRelatedDef := range directlyRelated {
		switch {
		case directlyRelatedDef.GetRelationOrWildcard() == nil:
			// direct assignment to concrete type
			assignableType := directlyRelatedDef.GetType()
			curNode = wg.GetOrAddNode(assignableType, assignableType, SpecificType, "")
		case directlyRelatedDef.GetWildcard() != nil:
			// direct assignment to wildcard
			assignableWildcard := directlyRelatedDef.GetType() + ":*"
			curNode = wg.GetOrAddNode(assignableWildcard, assignableWildcard, SpecificTypeWildcard, "")
		default:
			// direct assignment to userset
			assignableUserset := directlyRelatedDef.GetType() + "#" + directlyRelatedDef.GetRelation()
			curNode = wg.GetOrAddNode(assignableUserset, assignableUserset, SpecificTypeAndRelation, assignableUserset)
		}

		// de-dup types that are conditioned, e.g. if define viewer: [user, user with condX]
		// we only draw one edge from user to x#viewer, but with two conditions: none and condX
		err := wg.UpsertEdge(node, curNode, DirectEdge, parentRelationName, "", directlyRelatedDef.GetCondition())
		if err != nil {
			return err
		}
	}
	return nil
}

func GetRelationNodeName(objectType string, relation string) string {
	return fmt.Sprintf("%s#%s", objectType, relation)
}

func (wgb *WeightedAuthorizationModelGraphBuilder) ValidateRelationExists(objectType string, relation string, wg *WeightedAuthorizationModelGraph) error {
	nodeName := GetRelationNodeName(objectType, relation)
	_, exists := wg.GetNodeByID(nodeName)
	if !exists {
		return errors.RelationObjectTypeError(relation, objectType, errors.ErrRelationUndefined)
	}
	return nil
}
