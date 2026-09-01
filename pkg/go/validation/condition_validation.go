package validation

import (
	"maps"
	"slices"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// conditionUse is one place a condition is referenced from: a type restriction
// on a relation.
type conditionUse struct {
	typeName     string
	relationName string
}

// validateConditions runs the three condition checks in the reference's order:
// every referenced condition is defined, every condition's nested name matches
// its key, and every defined condition is referenced.
func validateConditions(model *openfgav1.AuthorizationModel, src source) Findings {
	uses := conditionUses(model)

	fs := undefinedConditions(model, src, uses)

	// A condition whose nested name property differs from its map key. It
	// carries no position, matching the reference.
	conditions := model.GetConditions()
	for _, conditionKey := range slices.Sorted(maps.Keys(conditions)) {
		condition := conditions[conditionKey]
		if condition != nil && condition.GetName() != conditionKey {
			fs = append(fs, differentNestedConditionName(conditionKey, condition.GetName()))
		}
	}

	// A condition defined but never referenced.
	for _, conditionName := range slices.Sorted(maps.Keys(conditions)) {
		if len(uses[conditionName]) > 0 {
			continue
		}

		condition := conditions[conditionName]
		file := condition.GetMetadata().GetSourceInfo().GetFile()
		module := condition.GetMetadata().GetModule()

		fs = append(fs, unusedCondition(conditionName).at(src, src.conditionLine(conditionName)).in(file, module))
	}

	return fs
}

// conditionUses collects where each condition is referenced, in the order the
// references appear walking types in model order and relations in name order —
// the order the reference implementation reports them in.
func conditionUses(model *openfgav1.AuthorizationModel) map[string][]conditionUse {
	uses := make(map[string][]conditionUse)

	for _, typeDef := range model.GetTypeDefinitions() {
		relationsMetadata := typeDef.GetMetadata().GetRelations()
		for _, relationName := range slices.Sorted(maps.Keys(relationsMetadata)) {
			for _, restriction := range relationsMetadata[relationName].GetDirectlyRelatedUserTypes() {
				if condition := restriction.GetCondition(); condition != "" {
					uses[condition] = append(uses[condition], conditionUse{
						typeName:     typeDef.GetType(),
						relationName: relationName,
					})
				}
			}
		}
	}

	return uses
}

// undefinedConditions reports, for every reference to a condition the model
// does not define, one finding per referencing relation.
func undefinedConditions(model *openfgav1.AuthorizationModel, src source,
	uses map[string][]conditionUse) Findings {
	var fs Findings

	defined := model.GetConditions()

	for _, conditionName := range slices.Sorted(maps.Keys(uses)) {
		if _, ok := defined[conditionName]; ok {
			continue
		}

		for _, use := range uses[conditionName] {
			// Anchor the relation line lookup to the referencing type's
			// declaration so the correct `define` is found when several types
			// share a relation name, matching the reference.
			line := src.relationLine(use.relationName, src.typeLine(use.typeName))

			var file, module string
			for _, typeDef := range model.GetTypeDefinitions() {
				if typeDef.GetType() == use.typeName {
					file, module = typeMeta(typeDef)

					break
				}
			}

			fs = append(fs, conditionNotDefined(conditionName, use.typeName, use.relationName).
				at(src, line).in(file, module))
		}
	}

	return fs
}
