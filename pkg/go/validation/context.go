package validation

import (
	fgaSdk "github.com/openfga/go-sdk"
)

// ValidationContext holds the state during model validation
// This is equivalent to the context maintained in the JS implementation
type ValidationContext struct {
	// TypeMap maps type names to their definitions for quick lookup
	TypeMap map[string]*fgaSdk.TypeDefinition

	// VisitedRelations tracks visited type/relation pairs for cycle detection
	// Format: TypeName -> RelationName -> bool
	VisitedRelations map[string]map[string]bool

	// UsedConditionNames tracks which conditions are actually used in the model
	UsedConditionNames map[string]bool

	// FileToModuleMap tracks which modules are defined in each file
	// Used for detecting multiple modules in single file
	FileToModuleMap map[string]map[string]bool

	// Conditions from the authorization model for validation
	Conditions map[string]*fgaSdk.Condition

	// Lines from the DSL for error reporting with line numbers
	Lines []string
}

// NewValidationContext creates a new validation context.
func NewValidationContext(lines []string) *ValidationContext {
	return &ValidationContext{
		TypeMap:            make(map[string]*fgaSdk.TypeDefinition),
		VisitedRelations:   make(map[string]map[string]bool),
		UsedConditionNames: make(map[string]bool),
		FileToModuleMap:    make(map[string]map[string]bool),
		Conditions:         make(map[string]*fgaSdk.Condition),
		Lines:              lines,
	}
}

// AddType adds a type definition to the type map.
func (ctx *ValidationContext) AddType(typeName string, typeDef *fgaSdk.TypeDefinition) {
	ctx.TypeMap[typeName] = typeDef
}

// GetType retrieves a type definition by name.
func (ctx *ValidationContext) GetType(typeName string) (*fgaSdk.TypeDefinition, bool) {
	typeDef, exists := ctx.TypeMap[typeName]
	return typeDef, exists
}

// MarkRelationVisited marks a type/relation pair as visited for cycle detection.
func (ctx *ValidationContext) MarkRelationVisited(typeName, relationName string) {
	if ctx.VisitedRelations[typeName] == nil {
		ctx.VisitedRelations[typeName] = make(map[string]bool)
	}
	ctx.VisitedRelations[typeName][relationName] = true
}

// IsRelationVisited checks if a type/relation pair has been visited.
func (ctx *ValidationContext) IsRelationVisited(typeName, relationName string) bool {
	if ctx.VisitedRelations[typeName] == nil {
		return false
	}
	return ctx.VisitedRelations[typeName][relationName]
}

// MarkConditionUsed marks a condition as used.
func (ctx *ValidationContext) MarkConditionUsed(conditionName string) {
	ctx.UsedConditionNames[conditionName] = true
}

// IsConditionUsed checks if a condition has been marked as used.
func (ctx *ValidationContext) IsConditionUsed(conditionName string) bool {
	return ctx.UsedConditionNames[conditionName]
}

// AddModuleToFile adds a module to a file mapping.
func (ctx *ValidationContext) AddModuleToFile(filename, module string) {
	if ctx.FileToModuleMap[filename] == nil {
		ctx.FileToModuleMap[filename] = make(map[string]bool)
	}
	ctx.FileToModuleMap[filename][module] = true
}

// GetModulesForFile returns all modules defined in a file.
func (ctx *ValidationContext) GetModulesForFile(filename string) []string {
	modules := make([]string, 0)
	if moduleMap := ctx.FileToModuleMap[filename]; moduleMap != nil {
		for module := range moduleMap {
			modules = append(modules, module)
		}
	}
	return modules
}

// HasMultipleModulesInFile checks if a file has multiple modules.
func (ctx *ValidationContext) HasMultipleModulesInFile(filename string) bool {
	return len(ctx.GetModulesForFile(filename)) > 1
}

// DeepCopyVisitedRelations creates a deep copy of visited relations for recursive validation
// This is equivalent to the deepCopy function in the JS implementation
func (ctx *ValidationContext) DeepCopyVisitedRelations() map[string]map[string]bool {
	copy := make(map[string]map[string]bool)
	for typeName, relations := range ctx.VisitedRelations {
		copy[typeName] = make(map[string]bool)
		for relationName, visited := range relations {
			copy[typeName][relationName] = visited
		}
	}
	return copy
}

// RelationTargetParserResult represents the result of parsing a relation target
// This is equivalent to the RelationTargetParserResult interface in JS
type RelationTargetParserResult struct {
	Target  string      `json:"target,omitempty"`
	From    string      `json:"from,omitempty"`
	Rewrite RewriteType `json:"rewrite"`
}

// RewriteType represents the type of rewrite operation.
type RewriteType string

const (
	RewriteDirect          RewriteType = "direct"
	RewriteComputedUserset RewriteType = "computed_userset"
	RewriteTupleToUserset  RewriteType = "tuple_to_userset"
)

// EntryPointResult represents the result of entry point analysis
// This is equivalent to the return type of hasEntryPointOrLoop in JS
type EntryPointResult struct {
	HasEntry bool `json:"hasEntry"`
	Loop     bool `json:"loop"`
}

// DestructedAssignableType represents a parsed assignable type
// This is equivalent to the DestructedAssignableType interface in JS
type DestructedAssignableType struct {
	DecodedType      string `json:"decodedType"`
	DecodedRelation  string `json:"decodedRelation,omitempty"`
	IsWildcard       bool   `json:"isWildcard"`
	DecodedCondition string `json:"decodedConditionName,omitempty"`
}

// ValidationRegex represents a validation rule with regex pattern
// This is equivalent to the ValidationRegex interface in JS
type ValidationRegex struct {
	Rule  string `json:"rule"`
	Regex string `json:"regex"`
}

// ValidationOptions represents options for validation
// This is equivalent to the ValidationOptions interface in JS
type ValidationOptions struct {
	TypeValidation     string `json:"typeValidation,omitempty"`
	RelationValidation string `json:"relationValidation,omitempty"`
}
