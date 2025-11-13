package validation

import (
	"testing"

	fgaSdk "github.com/openfga/go-sdk"
	"github.com/stretchr/testify/assert"
)

func TestNewValidationContext(t *testing.T) {
	lines := []string{"line 1", "line 2", "line 3"}
	ctx := NewValidationContext(lines)

	assert.NotNil(t, ctx)
	assert.NotNil(t, ctx.TypeMap)
	assert.NotNil(t, ctx.VisitedRelations)
	assert.NotNil(t, ctx.UsedConditionNames)
	assert.NotNil(t, ctx.FileToModuleMap)
	assert.NotNil(t, ctx.Conditions)
	assert.Equal(t, lines, ctx.Lines)

	// Test that maps are initialized
	assert.Empty(t, ctx.TypeMap)
	assert.Empty(t, ctx.VisitedRelations)
	assert.Empty(t, ctx.UsedConditionNames)
	assert.Empty(t, ctx.FileToModuleMap)
	assert.Empty(t, ctx.Conditions)
}

func TestValidationContext_AddType(t *testing.T) {
	ctx := NewValidationContext(nil)

	typeDef := &fgaSdk.TypeDefinition{
		Type: "document",
	}

	ctx.AddType("document", typeDef)

	assert.Len(t, ctx.TypeMap, 1)
	assert.Equal(t, typeDef, ctx.TypeMap["document"])
}

func TestValidationContext_GetType(t *testing.T) {
	ctx := NewValidationContext(nil)

	// Test getting non-existent type
	typeDef, exists := ctx.GetType("document")
	assert.Nil(t, typeDef)
	assert.False(t, exists)

	// Add type and test getting it
	expectedTypeDef := &fgaSdk.TypeDefinition{
		Type: "document",
	}
	ctx.AddType("document", expectedTypeDef)

	typeDef, exists = ctx.GetType("document")
	assert.Equal(t, expectedTypeDef, typeDef)
	assert.True(t, exists)
}

func TestValidationContext_MarkRelationVisited(t *testing.T) {
	ctx := NewValidationContext(nil)

	// Initially no relations are visited
	assert.False(t, ctx.IsRelationVisited("document", "viewer"))

	// Mark relation as visited
	ctx.MarkRelationVisited("document", "viewer")

	// Check that it's now visited
	assert.True(t, ctx.IsRelationVisited("document", "viewer"))

	// Check that other relations are not visited
	assert.False(t, ctx.IsRelationVisited("document", "admin"))
	assert.False(t, ctx.IsRelationVisited("user", "viewer"))
}

func TestValidationContext_IsRelationVisited(t *testing.T) {
	ctx := NewValidationContext(nil)

	// Test with non-existent type
	assert.False(t, ctx.IsRelationVisited("nonexistent", "relation"))

	// Test with existing type but non-existent relation
	ctx.MarkRelationVisited("document", "viewer")
	assert.False(t, ctx.IsRelationVisited("document", "nonexistent"))

	// Test with existing type and relation
	assert.True(t, ctx.IsRelationVisited("document", "viewer"))
}

func TestValidationContext_MultipleRelationsPerType(t *testing.T) {
	ctx := NewValidationContext(nil)

	// Mark multiple relations for the same type
	ctx.MarkRelationVisited("document", "viewer")
	ctx.MarkRelationVisited("document", "admin")
	ctx.MarkRelationVisited("document", "owner")

	// All should be visited
	assert.True(t, ctx.IsRelationVisited("document", "viewer"))
	assert.True(t, ctx.IsRelationVisited("document", "admin"))
	assert.True(t, ctx.IsRelationVisited("document", "owner"))

	// Other types should not be affected
	assert.False(t, ctx.IsRelationVisited("user", "viewer"))
}

func TestValidationContext_MarkConditionUsed(t *testing.T) {
	ctx := NewValidationContext(nil)

	// Initially no conditions are used
	assert.False(t, ctx.IsConditionUsed("condition1"))

	// Mark condition as used
	ctx.MarkConditionUsed("condition1")

	// Check that it's now used
	assert.True(t, ctx.IsConditionUsed("condition1"))

	// Check that other conditions are not used
	assert.False(t, ctx.IsConditionUsed("condition2"))
}

func TestValidationContext_IsConditionUsed(t *testing.T) {
	ctx := NewValidationContext(nil)

	// Test with non-existent condition
	assert.False(t, ctx.IsConditionUsed("nonexistent"))

	// Mark condition and test
	ctx.MarkConditionUsed("test_condition")
	assert.True(t, ctx.IsConditionUsed("test_condition"))
	assert.False(t, ctx.IsConditionUsed("other_condition"))
}

func TestValidationContext_AddModuleToFile(t *testing.T) {
	ctx := NewValidationContext(nil)

	// Initially no modules
	modules := ctx.GetModulesForFile("test.fga")
	assert.Empty(t, modules)

	// Add module to file
	ctx.AddModuleToFile("test.fga", "module1")

	// Check that module is added
	modules = ctx.GetModulesForFile("test.fga")
	assert.Len(t, modules, 1)
	assert.Contains(t, modules, "module1")
}

func TestValidationContext_GetModulesForFile(t *testing.T) {
	ctx := NewValidationContext(nil)

	// Test with non-existent file
	modules := ctx.GetModulesForFile("nonexistent.fga")
	assert.Empty(t, modules)

	// Add multiple modules to same file
	ctx.AddModuleToFile("test.fga", "module1")
	ctx.AddModuleToFile("test.fga", "module2")
	ctx.AddModuleToFile("test.fga", "module3")

	modules = ctx.GetModulesForFile("test.fga")
	assert.Len(t, modules, 3)
	assert.Contains(t, modules, "module1")
	assert.Contains(t, modules, "module2")
	assert.Contains(t, modules, "module3")

	// Test that other files are not affected
	otherModules := ctx.GetModulesForFile("other.fga")
	assert.Empty(t, otherModules)
}

func TestValidationContext_HasMultipleModulesInFile(t *testing.T) {
	ctx := NewValidationContext(nil)

	// Initially no modules
	assert.False(t, ctx.HasMultipleModulesInFile("test.fga"))

	// Add single module
	ctx.AddModuleToFile("test.fga", "module1")
	assert.False(t, ctx.HasMultipleModulesInFile("test.fga"))

	// Add second module
	ctx.AddModuleToFile("test.fga", "module2")
	assert.True(t, ctx.HasMultipleModulesInFile("test.fga"))

	// Test with non-existent file
	assert.False(t, ctx.HasMultipleModulesInFile("nonexistent.fga"))
}

func TestValidationContext_DeepCopyVisitedRelations(t *testing.T) {
	ctx := NewValidationContext(nil)

	// Add some visited relations
	ctx.MarkRelationVisited("document", "viewer")
	ctx.MarkRelationVisited("document", "admin")
	ctx.MarkRelationVisited("user", "member")

	// Create deep copy
	copy := ctx.DeepCopyVisitedRelations()

	// Verify copy has same content
	assert.True(t, copy["document"]["viewer"])
	assert.True(t, copy["document"]["admin"])
	assert.True(t, copy["user"]["member"])

	// Modify original
	ctx.MarkRelationVisited("document", "owner")

	// Verify copy is not affected
	assert.False(t, copy["document"]["owner"])
	assert.True(t, ctx.IsRelationVisited("document", "owner"))

	// Modify copy
	copy["user"]["admin"] = true

	// Verify original is not affected
	assert.False(t, ctx.IsRelationVisited("user", "admin"))
}

func TestRewriteType(t *testing.T) {
	// Test that rewrite type constants are defined correctly
	assert.Equal(t, "direct", string(RewriteDirect))
	assert.Equal(t, "computed_userset", string(RewriteComputedUserset))
	assert.Equal(t, "tuple_to_userset", string(RewriteTupleToUserset))
}

func TestRelationTargetParserResult(t *testing.T) {
	result := RelationTargetParserResult{
		Target:  "viewer",
		From:    "parent",
		Rewrite: RewriteTupleToUserset,
	}

	assert.Equal(t, "viewer", result.Target)
	assert.Equal(t, "parent", result.From)
	assert.Equal(t, RewriteTupleToUserset, result.Rewrite)
}

func TestEntryPointResult(t *testing.T) {
	result := EntryPointResult{
		HasEntry: true,
		Loop:     false,
	}

	assert.True(t, result.HasEntry)
	assert.False(t, result.Loop)
}

func TestDestructedAssignableType(t *testing.T) {
	assignable := DestructedAssignableType{
		DecodedType:      "user",
		DecodedRelation:  "member",
		IsWildcard:       false,
		DecodedCondition: "condition1",
	}

	assert.Equal(t, "user", assignable.DecodedType)
	assert.Equal(t, "member", assignable.DecodedRelation)
	assert.False(t, assignable.IsWildcard)
	assert.Equal(t, "condition1", assignable.DecodedCondition)
}

func TestValidationRegex(t *testing.T) {
	regex := ValidationRegex{
		Rule:  "[a-zA-Z]+",
		Regex: "^[a-zA-Z]+$",
	}

	assert.Equal(t, "[a-zA-Z]+", regex.Rule)
	assert.Equal(t, "^[a-zA-Z]+$", regex.Regex)
}

func TestValidationOptions(t *testing.T) {
	options := ValidationOptions{
		TypeValidation:     "strict",
		RelationValidation: "loose",
	}

	assert.Equal(t, "strict", options.TypeValidation)
	assert.Equal(t, "loose", options.RelationValidation)
}

func TestValidationContext_Integration(t *testing.T) {
	// Test a more complex integration scenario
	lines := []string{
		"model",
		"  schema 1.1",
		"type document",
		"  relations",
		"    define viewer: [user]",
		"    define admin: [user]",
		"type user",
	}

	ctx := NewValidationContext(lines)

	// Add types
	docType := &fgaSdk.TypeDefinition{Type: "document"}
	userType := &fgaSdk.TypeDefinition{Type: "user"}

	ctx.AddType("document", docType)
	ctx.AddType("user", userType)

	// Mark relations as visited during validation
	ctx.MarkRelationVisited("document", "viewer")
	ctx.MarkRelationVisited("document", "admin")

	// Mark conditions as used
	ctx.MarkConditionUsed("is_owner")

	// Add modules to files
	ctx.AddModuleToFile("model.fga", "main")
	ctx.AddModuleToFile("permissions.fga", "permissions")
	ctx.AddModuleToFile("permissions.fga", "conditions") // Multiple modules in one file

	// Verify state
	assert.Len(t, ctx.TypeMap, 2)
	assert.True(t, ctx.IsRelationVisited("document", "viewer"))
	assert.True(t, ctx.IsRelationVisited("document", "admin"))
	assert.False(t, ctx.IsRelationVisited("user", "member"))
	assert.True(t, ctx.IsConditionUsed("is_owner"))
	assert.False(t, ctx.IsConditionUsed("is_admin"))
	assert.False(t, ctx.HasMultipleModulesInFile("model.fga"))
	assert.True(t, ctx.HasMultipleModulesInFile("permissions.fga"))

	// Test deep copy doesn't affect original
	copy := ctx.DeepCopyVisitedRelations()
	// Initialize user map if it doesn't exist
	if copy["user"] == nil {
		copy["user"] = make(map[string]bool)
	}
	copy["user"]["member"] = true
	assert.False(t, ctx.IsRelationVisited("user", "member"))
}
