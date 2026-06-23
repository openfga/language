package validation

import (
	"path/filepath"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// MultiFileValidator handles validation across multiple files and modules.
type MultiFileValidator struct {
	model              *openfgav1.AuthorizationModel
	fileToModuleMap    map[string]map[string]bool
	moduleToFileMap    map[string]map[string]bool
	typeModuleMap      map[string]string
	conditionModuleMap map[string]string
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
		fileToModuleMap:    make(map[string]map[string]bool),
		moduleToFileMap:    make(map[string]map[string]bool),
		typeModuleMap:      make(map[string]string),
		conditionModuleMap: make(map[string]string),
	}
	validator.buildFileMappings()
	return validator
}

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
	for conditionName, condition := range mfv.model.GetConditions() {
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
	if mfv.fileToModuleMap[file] == nil {
		mfv.fileToModuleMap[file] = make(map[string]bool)
	}
	mfv.fileToModuleMap[file][module] = true
	if mfv.moduleToFileMap[module] == nil {
		mfv.moduleToFileMap[module] = make(map[string]bool)
	}
	mfv.moduleToFileMap[module][file] = true
}

// ValidateMultiFileConsistency validates consistency across multiple files.
func ValidateMultiFileConsistency(collector *ErrorCollector, model *openfgav1.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}
	validator := NewMultiFileValidator(model)
	validator.validateMultipleModulesInFile(collector)
}

func (mfv *MultiFileValidator) validateMultipleModulesInFile(collector *ErrorCollector) {
	for file, modules := range mfv.fileToModuleMap {
		if len(modules) > 1 {
			moduleNames := make([]string, 0, len(modules))
			for module := range modules {
				moduleNames = append(moduleNames, module)
			}
			collector.RaiseMultipleModulesInSingleFile(file, moduleNames)
		}
	}
}

func (mfv *MultiFileValidator) GetModuleInfo() []ModuleInfo {
	modules := make([]ModuleInfo, 0, len(mfv.moduleToFileMap))
	for moduleName, files := range mfv.moduleToFileMap {
		mi := ModuleInfo{Name: moduleName, Files: make([]string, 0, len(files)), Types: make([]string, 0)}
		for file := range files {
			mi.Files = append(mi.Files, file)
		}
		for typeName, typeModule := range mfv.typeModuleMap {
			if typeModule == moduleName {
				mi.Types = append(mi.Types, typeName)
			}
		}
		modules = append(modules, mi)
	}
	return modules
}

func (mfv *MultiFileValidator) GetFileInfo() []FileInfo {
	files := make([]FileInfo, 0, len(mfv.fileToModuleMap))
	for filePath, modules := range mfv.fileToModuleMap {
		fi := FileInfo{Path: filePath, Modules: make([]string, 0, len(modules))}
		for module := range modules {
			fi.Modules = append(fi.Modules, module)
		}
		files = append(files, fi)
	}
	return files
}

func (mfv *MultiFileValidator) IsMultiModuleProject() bool { return len(mfv.moduleToFileMap) > 1 }
func (mfv *MultiFileValidator) IsMultiFileProject() bool   { return len(mfv.fileToModuleMap) > 1 }
func (mfv *MultiFileValidator) GetModuleForType(typeName string) string {
	return mfv.typeModuleMap[typeName]
}
func (mfv *MultiFileValidator) GetModuleForCondition(conditionName string) string {
	return mfv.conditionModuleMap[conditionName]
}
func (mfv *MultiFileValidator) GetFilesForModule(moduleName string) []string {
	files := make([]string, 0)
	if moduleFiles, exists := mfv.moduleToFileMap[moduleName]; exists {
		for file := range moduleFiles {
			files = append(files, file)
		}
	}
	return files
}
func (mfv *MultiFileValidator) GetModulesForFile(filePath string) []string {
	modules := make([]string, 0)
	if fileModules, exists := mfv.fileToModuleMap[filePath]; exists {
		for module := range fileModules {
			modules = append(modules, module)
		}
	}
	return modules
}
