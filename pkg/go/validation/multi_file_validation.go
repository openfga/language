package validation

import (
	"maps"
	"path/filepath"
	"slices"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// validateMultiFile reports every file that would contain more than one module
// when transformed back to DSL. Findings carry no position: they are about
// files, not lines.
func validateMultiFile(model *openfgav1.AuthorizationModel) Findings {
	var fs Findings

	files := modulesByFile(model)
	for _, file := range files.keys {
		if modules := files.values[file]; len(modules) > 1 {
			fs = append(fs, multipleModulesInSingleFile(file, modules))
		}
	}

	return fs
}

// orderedGroups records a one-to-many mapping, keeping both the keys and each
// key's values in the order they were first added.
//
// A file's module names are joined into the message it reports, and the
// reference lists them in the order they appear in the model. Collecting them
// into a map would report the same model differently from one run to the next,
// and sorting them would report it differently from the other SDKs, so the
// shared corpus fails either way.
type orderedGroups struct {
	keys   []string
	values map[string][]string
}

func (g *orderedGroups) add(key, value string) {
	existing, seen := g.values[key]
	if !seen {
		g.keys = append(g.keys, key)
	}

	if slices.Contains(existing, value) {
		return
	}

	g.values[key] = append(existing, value)
}

// modulesByFile walks the model in the order the reference walks it: every
// type, then every relation, then every condition. The passes are separate
// because a relation's module is recorded after the module of every type, not
// after its own type's.
func modulesByFile(model *openfgav1.AuthorizationModel) *orderedGroups {
	files := &orderedGroups{values: make(map[string][]string)}

	record := func(file, module string) {
		if file != "" && module != "" {
			files.add(filepath.Clean(file), module)
		}
	}

	for _, typeDef := range model.GetTypeDefinitions() {
		record(typeMeta(typeDef))
	}

	for _, typeDef := range model.GetTypeDefinitions() {
		// Relation names arrive in a proto map, which has no order of its own,
		// so they are walked in name order.
		for _, relationName := range slices.Sorted(maps.Keys(typeDef.GetRelations())) {
			record(relationMeta(typeDef, relationName))
		}
	}

	conditions := model.GetConditions()
	for _, conditionName := range slices.Sorted(maps.Keys(conditions)) {
		condition := conditions[conditionName]
		record(condition.GetMetadata().GetSourceInfo().GetFile(), condition.GetMetadata().GetModule())
	}

	return files
}
