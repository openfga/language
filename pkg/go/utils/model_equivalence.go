package utils

import (
	"bytes"
	"slices"
	"strings"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"google.golang.org/protobuf/proto"
)

// AreModelsEquivalent returns true if the two authorization models define the same types, relations and
// conditions, so that writing one of them when the other is the latest model of a store would not change any
// authorization decision.
//
// The comparison ignores:
// - the model IDs
// - the order of type definitions, of directly related user types and of the children of a union or intersection
// - how unions and intersections are nested, e.g. `a or (b or c)` is equivalent to `a or b or c`
// - module and source file information, so moving a type to another module or file does not make a model different
// - leading and trailing whitespace in condition expressions
//
// The order of the base and subtracted sets of an exclusion (`but not`) is kept, since swapping them changes the
// meaning of the relation.
func AreModelsEquivalent(model1, model2 *openfgav1.AuthorizationModel) bool {
	if model1 == nil || model2 == nil {
		return model1 == model2
	}

	return proto.Equal(normalizeModel(model1), normalizeModel(model2))
}

// normalizeModel returns a copy of the model with everything that AreModelsEquivalent ignores removed or put in a
// canonical order.
func normalizeModel(model *openfgav1.AuthorizationModel) *openfgav1.AuthorizationModel {
	normalized, _ := proto.Clone(model).(*openfgav1.AuthorizationModel)
	normalized.Id = ""

	for _, typeDef := range normalized.GetTypeDefinitions() {
		for _, rewrite := range typeDef.GetRelations() {
			normalizeUserset(rewrite)
		}

		typeDef.Metadata = normalizeTypeMetadata(typeDef.GetMetadata())
	}

	slices.SortFunc(normalized.GetTypeDefinitions(), func(a, b *openfgav1.TypeDefinition) int {
		return strings.Compare(a.GetType(), b.GetType())
	})

	for _, condition := range normalized.GetConditions() {
		condition.Expression = strings.TrimSpace(condition.GetExpression())
		condition.Metadata = nil
	}

	return normalized
}

func normalizeTypeMetadata(metadata *openfgav1.Metadata) *openfgav1.Metadata {
	if metadata == nil {
		return nil
	}

	relations := make(map[string]*openfgav1.RelationMetadata, len(metadata.GetRelations()))

	for relation, relationMetadata := range metadata.GetRelations() {
		userTypes := relationMetadata.GetDirectlyRelatedUserTypes()
		if len(userTypes) == 0 {
			continue
		}

		slices.SortFunc(userTypes, func(a, b *openfgav1.RelationReference) int {
			return compareProtos(a, b)
		})
		userTypes = slices.CompactFunc(userTypes, func(a, b *openfgav1.RelationReference) bool {
			return proto.Equal(a, b)
		})

		relations[relation] = &openfgav1.RelationMetadata{DirectlyRelatedUserTypes: userTypes}
	}

	if len(relations) == 0 {
		return nil
	}

	return &openfgav1.Metadata{Relations: relations}
}

// normalizeUserset flattens nested unions and intersections and sorts their children in place.
func normalizeUserset(userset *openfgav1.Userset) {
	switch rewrite := userset.GetUserset().(type) {
	case *openfgav1.Userset_Union:
		rewrite.Union.Child = normalizeChildren(rewrite.Union.GetChild(), func(child *openfgav1.Userset) []*openfgav1.Userset {
			return child.GetUnion().GetChild()
		})
	case *openfgav1.Userset_Intersection:
		rewrite.Intersection.Child = normalizeChildren(rewrite.Intersection.GetChild(),
			func(child *openfgav1.Userset) []*openfgav1.Userset {
				return child.GetIntersection().GetChild()
			})
	case *openfgav1.Userset_Difference:
		normalizeUserset(rewrite.Difference.GetBase())
		normalizeUserset(rewrite.Difference.GetSubtract())
	}
}

// normalizeChildren normalizes the children of a union or an intersection. Children that are themselves of the same
// kind (as returned by nested) are merged into the parent, then the children are sorted and duplicates are removed.
func normalizeChildren(
	children []*openfgav1.Userset,
	nested func(*openfgav1.Userset) []*openfgav1.Userset,
) []*openfgav1.Userset {
	flattened := make([]*openfgav1.Userset, 0, len(children))

	for _, child := range children {
		normalizeUserset(child)

		if grandChildren := nested(child); grandChildren != nil {
			flattened = append(flattened, grandChildren...)

			continue
		}

		flattened = append(flattened, child)
	}

	slices.SortFunc(flattened, func(a, b *openfgav1.Userset) int {
		return compareProtos(a, b)
	})

	return slices.CompactFunc(flattened, func(a, b *openfgav1.Userset) bool {
		return proto.Equal(a, b)
	})
}

// compareProtos orders two messages by their deterministic binary encoding, which gives a stable order for sorting.
func compareProtos(a, b proto.Message) int {
	options := proto.MarshalOptions{Deterministic: true}
	aBytes, _ := options.Marshal(a)
	bBytes, _ := options.Marshal(b)

	return bytes.Compare(aBytes, bBytes)
}
