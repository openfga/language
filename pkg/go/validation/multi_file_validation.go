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
func validateMultiFile(model *openfgav1.AuthorizationModel) error {
	var fs []*Finding

	files, modules := modulesByFile(model)
	for _, file := range files {
		if len(modules[file]) > 1 {
			fs = append(fs, multipleModulesInSingleFile(file, modules[file]))
		}
	}

	return joinFindings(fs...)
}

// modulesByFile returns the files named in the model's metadata, in the order
// first seen, and each file's modules, in the same order. A file's modules are
// joined into its message, and pkg/js lists them in the order it meets them,
// so a map's order or a sorted one would fail the shared corpus.
//
// The model is walked as pkg/js walks it: every type, then every relation,
// then every condition. The passes are separate because a relation's module is
// recorded after the module of every type, not after its own type's.
func modulesByFile(model *openfgav1.AuthorizationModel) ([]string, map[string][]string) {
	var files []string
	modules := make(map[string][]string)

	record := func(file, module string) {
		if file == "" || module == "" {
			return
		}

		file = filepath.Clean(file)
		if _, seen := modules[file]; !seen {
			files = append(files, file)
		}

		if !slices.Contains(modules[file], module) {
			modules[file] = append(modules[file], module)
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

	return files, modules
}
