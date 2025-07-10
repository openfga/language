package validation

import (
	"path/filepath"

	fgaSdk "github.com/openfga/go-sdk"
)

// MultiFileValidator handles validation across multiple files and modules
// This includes module consistency checking and file-to-module mapping validation
type MultiFileValidator struct {
	model              *fgaSdk.AuthorizationModel
	fileToModuleMap    map[string]map[string]bool // file -> modules defined in that file
	moduleToFileMap    map[string]map[string]bool // module -> files where module is defined
	typeModuleMap      map[string]string          // type -> module where type is defined
	conditionModuleMap map[string]string          // condition -> module where condition is defined
}

// ModuleInfo represents information about a module
type ModuleInfo struct {
	Name  string
	Files []string
	Types []string
}

// FileInfo represents information about a file
type FileInfo struct {
	Path    string
	Modules []string
}

// NewMultiFileValidator creates a new multi-file validator
func NewMultiFileValidator(model *fgaSdk.AuthorizationModel) *MultiFileValidator {
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

// buildFileMappings constructs internal mappings between files, modules, types, and conditions
func (mfv *MultiFileValidator) buildFileMappings() {
	if mfv.model == nil {
		return
	}

	// Build mappings from type definitions
	for _, typeDef := range mfv.model.TypeDefinitions {
		if typeDef.Metadata != nil {
			var file, module string

			// Extract file information
			if typeDef.Metadata.HasSourceInfo() {
				sourceInfo := typeDef.Metadata.GetSourceInfo()
				if sourceInfo.HasFile() {
					file = sourceInfo.GetFile()
				}
			}

			// Extract module information
			if typeDef.Metadata.HasModule() {
				module = typeDef.Metadata.GetModule()
			}

			// Update mappings if we have both file and module
			if file != "" && module != "" {
				mfv.addFileModuleMapping(file, module)
				mfv.typeModuleMap[typeDef.Type] = module
			}
		}
	}

	// Build mappings from conditions
	if mfv.model.Conditions != nil {
		for conditionName, condition := range *mfv.model.Conditions {
			if condition.Metadata != nil {
				var file, module string

				// Extract file information
				if condition.Metadata.HasSourceInfo() {
					sourceInfo := condition.Metadata.GetSourceInfo()
					if sourceInfo.HasFile() {
						file = sourceInfo.GetFile()
					}
				}

				// Extract module information
				if condition.Metadata.HasModule() {
					module = condition.Metadata.GetModule()
				}

				// Update mappings if we have both file and module
				if file != "" && module != "" {
					mfv.addFileModuleMapping(file, module)
					mfv.conditionModuleMap[conditionName] = module
				}
			}
		}
	}
}

// addFileModuleMapping adds a mapping between a file and module
func (mfv *MultiFileValidator) addFileModuleMapping(file, module string) {
	// Normalize file path
	file = filepath.Clean(file)

	// Add to file->module mapping
	if mfv.fileToModuleMap[file] == nil {
		mfv.fileToModuleMap[file] = make(map[string]bool)
	}
	mfv.fileToModuleMap[file][module] = true

	// Add to module->file mapping
	if mfv.moduleToFileMap[module] == nil {
		mfv.moduleToFileMap[module] = make(map[string]bool)
	}
	mfv.moduleToFileMap[module][file] = true
}

// ValidateMultiFileConsistency validates consistency across multiple files
func ValidateMultiFileConsistency(collector *ErrorCollector, model *fgaSdk.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}

	validator := NewMultiFileValidator(model)

	// Check for multiple modules in single file
	validator.validateMultipleModulesInFile(collector)
}

// validateMultipleModulesInFile checks for files that define multiple modules
func (mfv *MultiFileValidator) validateMultipleModulesInFile(collector *ErrorCollector) {
	for file, modules := range mfv.fileToModuleMap {
		if len(modules) > 1 {
			// Convert map keys to slice
			moduleNames := make([]string, 0, len(modules))
			for module := range modules {
				moduleNames = append(moduleNames, module)
			}

			collector.RaiseMultipleModulesInSingleFile(file, moduleNames)
		}
	}
}

// GetModuleInfo returns information about all modules
func (mfv *MultiFileValidator) GetModuleInfo() []ModuleInfo {
	modules := make([]ModuleInfo, 0, len(mfv.moduleToFileMap))

	for moduleName, files := range mfv.moduleToFileMap {
		moduleInfo := ModuleInfo{
			Name:  moduleName,
			Files: make([]string, 0, len(files)),
			Types: make([]string, 0),
		}

		// Add files
		for file := range files {
			moduleInfo.Files = append(moduleInfo.Files, file)
		}

		// Add types
		for typeName, typeModule := range mfv.typeModuleMap {
			if typeModule == moduleName {
				moduleInfo.Types = append(moduleInfo.Types, typeName)
			}
		}

		modules = append(modules, moduleInfo)
	}

	return modules
}

// GetFileInfo returns information about all files
func (mfv *MultiFileValidator) GetFileInfo() []FileInfo {
	files := make([]FileInfo, 0, len(mfv.fileToModuleMap))

	for filePath, modules := range mfv.fileToModuleMap {
		fileInfo := FileInfo{
			Path:    filePath,
			Modules: make([]string, 0, len(modules)),
		}

		// Add modules
		for module := range modules {
			fileInfo.Modules = append(fileInfo.Modules, module)
		}

		files = append(files, fileInfo)
	}

	return files
}

// IsMultiModuleProject returns true if the project contains multiple modules
func (mfv *MultiFileValidator) IsMultiModuleProject() bool {
	return len(mfv.moduleToFileMap) > 1
}

// IsMultiFileProject returns true if the project spans multiple files
func (mfv *MultiFileValidator) IsMultiFileProject() bool {
	return len(mfv.fileToModuleMap) > 1
}

// GetModuleForType returns the module name for a given type
func (mfv *MultiFileValidator) GetModuleForType(typeName string) string {
	return mfv.typeModuleMap[typeName]
}

// GetModuleForCondition returns the module name for a given condition
func (mfv *MultiFileValidator) GetModuleForCondition(conditionName string) string {
	return mfv.conditionModuleMap[conditionName]
}

// GetFilesForModule returns all files that define a given module
func (mfv *MultiFileValidator) GetFilesForModule(moduleName string) []string {
	files := make([]string, 0)
	if moduleFiles, exists := mfv.moduleToFileMap[moduleName]; exists {
		for file := range moduleFiles {
			files = append(files, file)
		}
	}
	return files
}

// GetModulesForFile returns all modules defined in a given file
func (mfv *MultiFileValidator) GetModulesForFile(filePath string) []string {
	modules := make([]string, 0)
	if fileModules, exists := mfv.fileToModuleMap[filePath]; exists {
		for module := range fileModules {
			modules = append(modules, module)
		}
	}
	return modules
}
