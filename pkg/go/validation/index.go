package validation

import (
	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// index is the model indexed for lookup: types and relations by name. It is
// built once per validation run and shared by every phase that resolves
// references, so each phase does not rebuild its own maps.
//
// When a type is declared more than once the maps are built last-wins, matching
// the reference's typeMap; the duplicate itself is reported by the duplicates
// phase.
type index struct {
	model     *openfgav1.AuthorizationModel
	types     map[string]*openfgav1.TypeDefinition
	relations map[string]map[string]*openfgav1.Userset
}

func newIndex(model *openfgav1.AuthorizationModel) *index {
	idx := &index{
		model:     model,
		types:     make(map[string]*openfgav1.TypeDefinition, len(model.GetTypeDefinitions())),
		relations: make(map[string]map[string]*openfgav1.Userset, len(model.GetTypeDefinitions())),
	}

	for _, typeDef := range model.GetTypeDefinitions() {
		idx.types[typeDef.GetType()] = typeDef

		if relations := typeDef.GetRelations(); len(relations) > 0 {
			byName := make(map[string]*openfgav1.Userset, len(relations))
			for relationName, userset := range relations {
				byName[relationName] = userset
			}

			idx.relations[typeDef.GetType()] = byName
		}
	}

	return idx
}

func (idx *index) typeDefined(typeName string) bool {
	_, ok := idx.types[typeName]

	return ok
}

func (idx *index) relationDefined(typeName, relationName string) bool {
	_, ok := idx.relations[typeName][relationName]

	return ok
}

func (idx *index) typeDef(typeName string) *openfgav1.TypeDefinition {
	return idx.types[typeName]
}

func (idx *index) userset(typeName, relationName string) *openfgav1.Userset {
	return idx.relations[typeName][relationName]
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
// relation is such a single direct assignment. This mirrors the reference
// implementation's allowableTypes helper used for tuple-to-userset validation.
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
