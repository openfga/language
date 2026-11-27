package validation

import (
	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// index is the model indexed for lookup: types by name, and through them their
// relations. It is built once per validation run and shared by every phase
// that resolves references, so each phase does not rebuild its own map.
//
// When a type is declared more than once the map is built last-wins, matching
// pkg/js's typeMap; the duplicate itself is reported by the duplicates phase.
type index struct {
	model *openfgav1.AuthorizationModel
	types map[string]*openfgav1.TypeDefinition
}

func newIndex(model *openfgav1.AuthorizationModel) *index {
	idx := &index{
		model: model,
		types: make(map[string]*openfgav1.TypeDefinition, len(model.GetTypeDefinitions())),
	}

	for _, typeDef := range model.GetTypeDefinitions() {
		idx.types[typeDef.GetType()] = typeDef
	}

	return idx
}

func (idx *index) typeDefined(typeName string) bool {
	_, ok := idx.types[typeName]

	return ok
}

// relationDefined reports whether typeName defines relationName. An unknown
// type has no relations: the proto getters accept a nil type definition.
func (idx *index) relationDefined(typeName, relationName string) bool {
	_, ok := idx.types[typeName].GetRelations()[relationName]

	return ok
}

func (idx *index) typeDef(typeName string) *openfgav1.TypeDefinition {
	return idx.types[typeName]
}

func (idx *index) userset(typeName, relationName string) *openfgav1.Userset {
	return idx.types[typeName].GetRelations()[relationName]
}

// directTypeRestrictions returns the directly-related user types declared for a
// relation in its metadata.
func (idx *index) directTypeRestrictions(typeName, relationName string) []*openfgav1.RelationReference {
	typeDef := idx.types[typeName]
	if typeDef == nil {
		return nil
	}

	relationMetadata, ok := typeDef.GetMetadata().GetRelations()[relationName]
	if !ok {
		return nil
	}

	return relationMetadata.GetDirectlyRelatedUserTypes()
}

// directlyAssignableTypes returns the type restrictions a relation is directly
// assignable to, but only when that relation is a single direct assignment
// (i.e. `define r: [a, b]` rather than a rewrite). The bool reports whether the
// relation is such a single direct assignment. This mirrors pkg/js's
// allowableTypes helper used for tuple-to-userset validation.
func (idx *index) directlyAssignableTypes(typeName, relationName string) ([]*openfgav1.RelationReference, bool) {
	userset := idx.userset(typeName, relationName)
	if userset == nil {
		return nil, false
	}

	if _, ok := userset.GetUserset().(*openfgav1.Userset_This); !ok {
		return nil, false
	}

	return idx.directTypeRestrictions(typeName, relationName), true
}

// typeMeta resolves the file and module a type was declared in.
func typeMeta(typeDef *openfgav1.TypeDefinition) (file, module string) {
	return typeDef.GetMetadata().GetSourceInfo().GetFile(), typeDef.GetMetadata().GetModule()
}

// relationMeta resolves the file and module for a relation, falling back to its
// type's for whichever of the two the relation leaves unset.
func relationMeta(typeDef *openfgav1.TypeDefinition, relationName string) (file, module string) {
	relationMetadata, ok := typeDef.GetMetadata().GetRelations()[relationName]
	if !ok {
		return typeMeta(typeDef)
	}

	file = relationMetadata.GetSourceInfo().GetFile()
	module = relationMetadata.GetModule()

	if file == "" {
		file = typeDef.GetMetadata().GetSourceInfo().GetFile()
	}

	if module == "" {
		module = typeDef.GetMetadata().GetModule()
	}

	return file, module
}
