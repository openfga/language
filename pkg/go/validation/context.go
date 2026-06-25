package validation

import (
	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// ValidationContext holds the state during model validation
type ValidationContext struct {
	TypeMap            map[string]*openfgav1.TypeDefinition
	VisitedRelations   map[string]map[string]bool
	UsedConditionNames map[string]bool
	FileToModuleMap    map[string]map[string]bool
	Conditions         map[string]*openfgav1.Condition
	Lines              []string
}

func NewValidationContext(lines []string) *ValidationContext {
	return &ValidationContext{
		TypeMap:            make(map[string]*openfgav1.TypeDefinition),
		VisitedRelations:   make(map[string]map[string]bool),
		UsedConditionNames: make(map[string]bool),
		FileToModuleMap:    make(map[string]map[string]bool),
		Conditions:         make(map[string]*openfgav1.Condition),
		Lines:              lines,
	}
}

func (ctx *ValidationContext) AddType(typeName string, typeDef *openfgav1.TypeDefinition) {
	ctx.TypeMap[typeName] = typeDef
}

func (ctx *ValidationContext) GetType(typeName string) (*openfgav1.TypeDefinition, bool) {
	typeDef, exists := ctx.TypeMap[typeName]
	return typeDef, exists
}

func (ctx *ValidationContext) MarkRelationVisited(typeName, relationName string) {
	if ctx.VisitedRelations[typeName] == nil {
		ctx.VisitedRelations[typeName] = make(map[string]bool)
	}
	ctx.VisitedRelations[typeName][relationName] = true
}

func (ctx *ValidationContext) IsRelationVisited(typeName, relationName string) bool {
	if ctx.VisitedRelations[typeName] == nil {
		return false
	}
	return ctx.VisitedRelations[typeName][relationName]
}

func (ctx *ValidationContext) MarkConditionUsed(conditionName string) {
	ctx.UsedConditionNames[conditionName] = true
}

func (ctx *ValidationContext) IsConditionUsed(conditionName string) bool {
	return ctx.UsedConditionNames[conditionName]
}

func (ctx *ValidationContext) AddModuleToFile(filename, module string) {
	if ctx.FileToModuleMap[filename] == nil {
		ctx.FileToModuleMap[filename] = make(map[string]bool)
	}
	ctx.FileToModuleMap[filename][module] = true
}

func (ctx *ValidationContext) GetModulesForFile(filename string) []string {
	modules := make([]string, 0, len(ctx.FileToModuleMap[filename]))
	if moduleMap := ctx.FileToModuleMap[filename]; moduleMap != nil {
		for module := range moduleMap {
			modules = append(modules, module)
		}
	}
	return modules
}

func (ctx *ValidationContext) HasMultipleModulesInFile(filename string) bool {
	return len(ctx.GetModulesForFile(filename)) > 1
}

func (ctx *ValidationContext) DeepCopyVisitedRelations() map[string]map[string]bool {
	cp := make(map[string]map[string]bool)
	for typeName, relations := range ctx.VisitedRelations {
		cp[typeName] = make(map[string]bool)
		for relationName, visited := range relations {
			cp[typeName][relationName] = visited
		}
	}
	return cp
}

// RelationTargetParserResult represents the result of parsing a relation target.
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

// EntryPointResult represents the result of entry point analysis.
type EntryPointResult struct {
	HasEntry bool `json:"hasEntry"`
	Loop     bool `json:"loop"`
}

// DestructedAssignableType represents a parsed assignable type.
type DestructedAssignableType struct {
	DecodedType      string `json:"decodedType"`
	DecodedRelation  string `json:"decodedRelation,omitempty"`
	IsWildcard       bool   `json:"isWildcard"`
	DecodedCondition string `json:"decodedConditionName,omitempty"`
}

// ValidationRegex represents a validation rule with regex pattern.
type ValidationRegex struct {
	Rule  string `json:"rule"`
	Regex string `json:"regex"`
}

// ValidationOptions represents options for validation.
type ValidationOptions struct {
	TypeValidation     string `json:"typeValidation,omitempty"`
	RelationValidation string `json:"relationValidation,omitempty"`
}
