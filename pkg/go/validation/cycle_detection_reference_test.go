package validation

import (
	"maps"
	"math/rand/v2"
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// pathWalk is hasEntryPointOrLoop from pkg/js's validate-dsl.ts: it walks every
// path out of a relation, skipping relations already on the path, and takes
// exponential time on some models. It is the reference resolveEntryPoints must
// match while answering without walking paths.
func pathWalk(idx *index, typeName, relationName string, rewrite *openfgav1.Userset,
	visited map[string]map[string]bool) entryPointResult {
	if relationName == "" || rewrite == nil {
		return entryPointResult{}
	}

	if visited[typeName] == nil {
		visited[typeName] = map[string]bool{}
	}

	visited[typeName][relationName] = true

	if !idx.relationDefined(typeName, relationName) {
		return entryPointResult{}
	}

	switch rewrite.GetUserset().(type) {
	case *openfgav1.Userset_This:
		for _, restriction := range idx.directTypeRestrictions(typeName, relationName) {
			restrictedType, restrictedRelation := restriction.GetType(), restriction.GetRelation()
			if restrictedRelation == "" || restriction.GetWildcard() != nil {
				return entryPointResult{hasEntry: true}
			}

			assignable := idx.userset(restrictedType, restrictedRelation)
			if assignable == nil {
				return entryPointResult{}
			}

			if visited[restrictedType][restrictedRelation] {
				continue
			}

			if pathWalk(idx, restrictedType, restrictedRelation, assignable, copyPath(visited)).hasEntry {
				return entryPointResult{hasEntry: true}
			}
		}

		return entryPointResult{}

	case *openfgav1.Userset_ComputedUserset:
		computed := rewrite.GetComputedUserset().GetRelation()
		if computed == "" || !idx.relationDefined(typeName, computed) {
			return entryPointResult{}
		}

		if visited[typeName][computed] {
			return entryPointResult{loop: true}
		}

		return pathWalk(idx, typeName, computed, idx.userset(typeName, computed), visited)

	case *openfgav1.Userset_TupleToUserset:
		tupleset := rewrite.GetTupleToUserset().GetTupleset().GetRelation()
		computed := rewrite.GetTupleToUserset().GetComputedUserset().GetRelation()

		if tupleset == "" || computed == "" || !idx.relationDefined(typeName, tupleset) {
			return entryPointResult{}
		}

		for _, restriction := range idx.directTypeRestrictions(typeName, tupleset) {
			assignable := idx.userset(restriction.GetType(), computed)
			if assignable == nil || visited[restriction.GetType()][computed] {
				continue
			}

			if pathWalk(idx, restriction.GetType(), computed, assignable, copyPath(visited)).hasEntry {
				return entryPointResult{hasEntry: true}
			}
		}

		return entryPointResult{}

	case *openfgav1.Userset_Union:
		hasLoop := false

		for _, child := range rewrite.GetUnion().GetChild() {
			result := pathWalk(idx, typeName, relationName, child, copyPath(visited))
			if result.hasEntry {
				return entryPointResult{hasEntry: true}
			}

			hasLoop = hasLoop || result.loop
		}

		return entryPointResult{loop: hasLoop}

	case *openfgav1.Userset_Intersection:
		for _, child := range rewrite.GetIntersection().GetChild() {
			result := pathWalk(idx, typeName, relationName, child, copyPath(visited))
			if !result.hasEntry {
				return entryPointResult{loop: result.loop}
			}
		}

		return entryPointResult{hasEntry: true}

	case *openfgav1.Userset_Difference:
		for _, operand := range []*openfgav1.Userset{rewrite.GetDifference().GetBase(), rewrite.GetDifference().GetSubtract()} {
			result := pathWalk(idx, typeName, relationName, operand, copyPath(visited))
			if !result.hasEntry {
				return entryPointResult{loop: result.loop}
			}
		}

		return entryPointResult{hasEntry: true}
	}

	return entryPointResult{}
}

func copyPath(visited map[string]map[string]bool) map[string]map[string]bool {
	copied := make(map[string]map[string]bool, len(visited))
	for typeName, relations := range visited {
		copied[typeName] = maps.Clone(relations)
	}

	return copied
}

// TestEntryPointsMatchPathWalk checks resolveEntryPoints against pathWalk on
// random models mixing every kind of rewrite with loops, shared dependencies
// and references to missing relations.
func TestEntryPointsMatchPathWalk(t *testing.T) {
	random := rand.New(rand.NewPCG(1, 2))

	for range 5000 {
		model := randomModel(random)
		idx := newIndex(model)
		entryPoints := resolveEntryPoints(idx)

		for _, typeDef := range model.GetTypeDefinitions() {
			for relationName, rewrite := range typeDef.GetRelations() {
				want := pathWalk(idx, typeDef.GetType(), relationName, rewrite, map[string]map[string]bool{})
				if got := entryPoints.of(typeDef.GetType(), relationName); got != want {
					t.Fatalf("%s#%s: got %+v, want %+v, in model %v",
						typeDef.GetType(), relationName, got, want, model)
				}
			}
		}
	}
}

var randomRelations = []string{"r0", "r1", "r2", "r3", "r4", "r5"}

// randomModel builds up to three types of up to six relations, each with a
// random rewrite and random type restrictions.
func randomModel(random *rand.Rand) *openfgav1.AuthorizationModel {
	types := []string{"t0", "t1", "t2"}[:1+random.IntN(3)]
	model := &openfgav1.AuthorizationModel{
		SchemaVersion:   "1.1",
		TypeDefinitions: []*openfgav1.TypeDefinition{{Type: "user"}},
	}

	for _, typeName := range types {
		typeDef := &openfgav1.TypeDefinition{
			Type:      typeName,
			Relations: map[string]*openfgav1.Userset{},
			Metadata:  &openfgav1.Metadata{Relations: map[string]*openfgav1.RelationMetadata{}},
		}

		defined := randomRelations[:1+random.IntN(len(randomRelations))]
		for _, relationName := range defined {
			typeDef.Relations[relationName] = randomRewrite(random, 1+random.IntN(3), defined)
			typeDef.Metadata.Relations[relationName] = &openfgav1.RelationMetadata{
				DirectlyRelatedUserTypes: randomRestrictions(random, types),
			}
		}

		model.TypeDefinitions = append(model.TypeDefinitions, typeDef)
	}

	return model
}

func randomRewrite(random *rand.Rand, depth int, defined []string) *openfgav1.Userset {
	pick := func(names []string) string { return names[random.IntN(len(names))] }

	operands := func() []*openfgav1.Userset {
		children := make([]*openfgav1.Userset, 2+random.IntN(2))
		for i := range children {
			children[i] = randomRewrite(random, depth-1, defined)
		}

		return children
	}

	if depth == 0 || random.IntN(3) == 0 {
		switch random.IntN(4) {
		case 0:
			return &openfgav1.Userset{Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}}}
		case 1:
			return &openfgav1.Userset{Userset: &openfgav1.Userset_ComputedUserset{
				ComputedUserset: &openfgav1.ObjectRelation{Relation: pick(defined)},
			}}
		case 2:
			return &openfgav1.Userset{Userset: &openfgav1.Userset_ComputedUserset{
				ComputedUserset: &openfgav1.ObjectRelation{Relation: pick(randomRelations)},
			}}
		default:
			return &openfgav1.Userset{Userset: &openfgav1.Userset_TupleToUserset{
				TupleToUserset: &openfgav1.TupleToUserset{
					Tupleset:        &openfgav1.ObjectRelation{Relation: pick(defined)},
					ComputedUserset: &openfgav1.ObjectRelation{Relation: pick(randomRelations)},
				},
			}}
		}
	}

	switch random.IntN(3) {
	case 0:
		return &openfgav1.Userset{Userset: &openfgav1.Userset_Union{Union: &openfgav1.Usersets{Child: operands()}}}
	case 1:
		return &openfgav1.Userset{Userset: &openfgav1.Userset_Intersection{
			Intersection: &openfgav1.Usersets{Child: operands()},
		}}
	default:
		return &openfgav1.Userset{Userset: &openfgav1.Userset_Difference{Difference: &openfgav1.Difference{
			Base:     randomRewrite(random, depth-1, defined),
			Subtract: randomRewrite(random, depth-1, defined),
		}}}
	}
}

func randomRestrictions(random *rand.Rand, types []string) []*openfgav1.RelationReference {
	restrictions := make([]*openfgav1.RelationReference, random.IntN(3))
	for i := range restrictions {
		restrictedType := types[random.IntN(len(types))]

		switch random.IntN(6) {
		case 0:
			restrictions[i] = &openfgav1.RelationReference{Type: "user"}
		case 1:
			restrictions[i] = &openfgav1.RelationReference{
				Type:               restrictedType,
				RelationOrWildcard: &openfgav1.RelationReference_Wildcard{Wildcard: &openfgav1.Wildcard{}},
			}
		default:
			restrictions[i] = &openfgav1.RelationReference{
				Type: restrictedType,
				RelationOrWildcard: &openfgav1.RelationReference_Relation{
					Relation: randomRelations[random.IntN(len(randomRelations))],
				},
			}
		}
	}

	return restrictions
}
