package validation

import (
	"maps"
	"path/filepath"
	"slices"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// MultiFileValidator handles validation across multiple files and modules.
type MultiFileValidator struct {
	model              *openfgav1.AuthorizationModel
	fileToModules      *orderedGroups
	moduleToFiles      *orderedGroups
	typeModuleMap      map[string]string
	conditionModuleMap map[string]string
}

// orderedGroups records a one-to-many mapping, keeping both the keys and each key's
// values in the order they were first added.
//
// A file's module names are joined into the message it reports, and the reference
// lists them in the order they appear in the model. Collecting them into a map would
// report the same model differently from one run to the next, and sorting them would
// report it differently from the other SDKs, so the shared corpus fails either way.
//
// Held by pointer: copying the struct copies keys but shares values, so an add
// through the copy leaves the original holding a value under a key it never
// recorded, and iterating keys then drops it.
type orderedGroups struct {
	keys   []string
	values map[string][]string
}

func newOrderedGroups() *orderedGroups {
	return &orderedGroups{values: make(map[string][]string)}
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

// get returns a non-nil copy, so a caller cannot reorder the record it reads. Not
// slices.Clone, which returns nil for an absent key and so would have the accessors
// hand back a slice that marshals as null rather than [].
func (g *orderedGroups) get(key string) []string {
	values := make([]string, 0, len(g.values[key]))

	return append(values, g.values[key]...)
}

// ModuleInfo represents information about a module.
type ModuleInfo struct {
	Name  string
	Files []string
	Types []string
}

// FileInfo represents information about a file.
type FileInfo struct {
	Path    string
	Modules []string
}

func NewMultiFileValidator(model *openfgav1.AuthorizationModel) *MultiFileValidator {
	validator := &MultiFileValidator{
		model:              model,
		fileToModules:      newOrderedGroups(),
		moduleToFiles:      newOrderedGroups(),
		typeModuleMap:      make(map[string]string),
		conditionModuleMap: make(map[string]string),
	}
	validator.buildFileMappings()

	return validator
}

// buildFileMappings walks the model in the order the reference walks it: every type,
// then every relation, then every condition. The passes are separate because a
// relation's module is reported after the module of every type, not after its own
// type's.
func (mfv *MultiFileValidator) buildFileMappings() {
	if mfv.model == nil {
		return
	}

	for _, typeDef := range mfv.model.GetTypeDefinitions() {
		file := typeDef.GetMetadata().GetSourceInfo().GetFile()
		module := typeDef.GetMetadata().GetModule()

		if file != "" && module != "" {
			mfv.addFileModuleMapping(file, module)
			mfv.typeModuleMap[typeDef.GetType()] = module
		}
	}

	for _, typeDef := range mfv.model.GetTypeDefinitions() {
		// Relation names arrive in a proto map, which has no order of its own, so
		// they are walked in name order.
		for _, relation := range slices.Sorted(maps.Keys(typeDef.GetRelations())) {
			relationMetadata := typeDef.GetMetadata().GetRelations()[relation]

			// A relation may name its own file and module, and falls back to its
			// type's for whichever of the two it leaves unset.
			file := relationMetadata.GetSourceInfo().GetFile()
			if file == "" {
				file = typeDef.GetMetadata().GetSourceInfo().GetFile()
			}

			module := relationMetadata.GetModule()
			if module == "" {
				module = typeDef.GetMetadata().GetModule()
			}

			if file != "" && module != "" {
				mfv.addFileModuleMapping(file, module)
			}
		}
	}

	for _, conditionName := range slices.Sorted(maps.Keys(mfv.model.GetConditions())) {
		condition := mfv.model.GetConditions()[conditionName]
		file := condition.GetMetadata().GetSourceInfo().GetFile()
		module := condition.GetMetadata().GetModule()

		if file != "" && module != "" {
			mfv.addFileModuleMapping(file, module)
			mfv.conditionModuleMap[conditionName] = module
		}
	}
}

func (mfv *MultiFileValidator) addFileModuleMapping(file, module string) {
	file = filepath.Clean(file)
	mfv.fileToModules.add(file, module)
	mfv.moduleToFiles.add(module, file)
}

// ValidateMultiFileConsistency validates consistency across multiple files.
func ValidateMultiFileConsistency(collector *ErrorCollector, model *openfgav1.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}
	// The rule itself lives in ValidateMultipleModulesInFile, which takes the files
	// this validator collected; the two must not drift.
	ValidateMultipleModulesInFile(collector, NewMultiFileValidator(model).GetFileInfo())
}

func (mfv *MultiFileValidator) GetModuleInfo() []ModuleInfo {
	modules := make([]ModuleInfo, 0, len(mfv.moduleToFiles.keys))

	for _, moduleName := range mfv.moduleToFiles.keys {
		info := ModuleInfo{Name: moduleName, Files: mfv.moduleToFiles.get(moduleName), Types: make([]string, 0)}

		// Declaration order: ranging typeModuleMap would list the types in whatever
		// order the runtime handed back. A name the model declares twice is reached
		// once per declaration, and typeModuleMap resolves it to a single module, so
		// each name is listed once rather than once per declaration.
		for _, typeDef := range mfv.model.GetTypeDefinitions() {
			typeName := typeDef.GetType()

			if mfv.typeModuleMap[typeName] != moduleName || slices.Contains(info.Types, typeName) {
				continue
			}

			info.Types = append(info.Types, typeName)
		}

		modules = append(modules, info)
	}

	return modules
}

func (mfv *MultiFileValidator) GetFileInfo() []FileInfo {
	files := make([]FileInfo, 0, len(mfv.fileToModules.keys))
	for _, filePath := range mfv.fileToModules.keys {
		files = append(files, FileInfo{Path: filePath, Modules: mfv.fileToModules.get(filePath)})
	}

	return files
}

func (mfv *MultiFileValidator) IsMultiModuleProject() bool { return len(mfv.moduleToFiles.keys) > 1 }
func (mfv *MultiFileValidator) IsMultiFileProject() bool   { return len(mfv.fileToModules.keys) > 1 }
func (mfv *MultiFileValidator) GetModuleForType(typeName string) string {
	return mfv.typeModuleMap[typeName]
}

func (mfv *MultiFileValidator) GetModuleForCondition(conditionName string) string {
	return mfv.conditionModuleMap[conditionName]
}

func (mfv *MultiFileValidator) GetFilesForModule(moduleName string) []string {
	return mfv.moduleToFiles.get(moduleName)
}

func (mfv *MultiFileValidator) GetModulesForFile(filePath string) []string {
	return mfv.fileToModules.get(filePath)
}
