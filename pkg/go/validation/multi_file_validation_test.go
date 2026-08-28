package validation

import (
	"slices"
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// multiModuleModel names one file, core.fga, from all three places a model can name a
// module: a type, a relation, and a condition. It is the shape the shared corpus uses
// for this rule, with a second relation and a second condition so map iteration has
// something to reorder.
func multiModuleModel() *openfgav1.AuthorizationModel {
	return &openfgav1.AuthorizationModel{
		SchemaVersion: "1.2",
		TypeDefinitions: []*openfgav1.TypeDefinition{
			{
				Type: "user",
				Relations: map[string]*openfgav1.Userset{
					"granted": {Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}}},
				},
				Metadata: &openfgav1.Metadata{
					Module:     "core",
					SourceInfo: &openfgav1.SourceInfo{File: "core.fga"},
					Relations: map[string]*openfgav1.RelationMetadata{
						"granted": {
							Module:     "usermodule",
							SourceInfo: &openfgav1.SourceInfo{File: "core.fga"},
						},
					},
				},
			},
			{
				Type: "org",
				Relations: map[string]*openfgav1.Userset{
					"member": {Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}}},
					"owner":  {Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}}},
				},
				Metadata: &openfgav1.Metadata{
					Module:     "other",
					SourceInfo: &openfgav1.SourceInfo{File: "core.fga"},
					Relations: map[string]*openfgav1.RelationMetadata{
						"member": {
							Module:     "relationmodule",
							SourceInfo: &openfgav1.SourceInfo{File: "core.fga"},
						},
						"owner": {
							Module:     "ownermodule",
							SourceInfo: &openfgav1.SourceInfo{File: "core.fga"},
						},
					},
				},
			},
		},
		Conditions: map[string]*openfgav1.Condition{
			"zeta": {
				Name: "zeta",
				Metadata: &openfgav1.ConditionMetadata{
					Module:     "zetamodule",
					SourceInfo: &openfgav1.SourceInfo{File: "core.fga"},
				},
			},
			"alpha": {
				Name: "alpha",
				Metadata: &openfgav1.ConditionMetadata{
					Module:     "alphamodule",
					SourceInfo: &openfgav1.SourceInfo{File: "core.fga"},
				},
			},
		},
	}
}

// TestMultiFileCollectionFollowsTheModelOrder pins the order the modules of one file
// are collected in: every type, then every relation, then every condition. That is the
// order the reference collects them in, and a file's modules are joined into the
// message it reports, so the shared corpus fails on any other order.
func TestMultiFileCollectionFollowsTheModelOrder(t *testing.T) {
	t.Parallel()

	model := multiModuleModel()

	// The passes are separate, so the first type's relation is collected after the
	// second type, not after its own type. Deliberately not in name order either:
	// sorting the modules would be deterministic and would still report the model
	// differently from the other SDKs.
	want := []string{
		"core", "other",
		"usermodule", "relationmodule", "ownermodule",
		"alphamodule", "zetamodule",
	}

	// Relations and conditions arrive in proto maps, which have no order of their own,
	// so one passing run proves nothing: seven modules admit 5040 orders, which 100
	// runs would not agree on by chance.
	for i := 0; i < 100; i++ {
		files := NewMultiFileValidator(model).GetFileInfo()

		require.Len(t, files, 1, "every definition in this model names core.fga")
		require.Equal(t, "core.fga", files[0].Path)
		require.Equalf(t, want, files[0].Modules, "run %d collected the modules in a different order", i)
	}
}

// TestValidateMultiFileConsistencyReportsEveryModuleInTheFile is the same rule from the
// entry point the engine calls, including the modules a relation and a condition name,
// which are the two the type-level walk alone would miss.
func TestValidateMultiFileConsistencyReportsEveryModuleInTheFile(t *testing.T) {
	t.Parallel()

	collector := NewValidationErrors(nil)
	ValidateMultiFileConsistency(collector, multiModuleModel(), nil)

	findings := collector.AllFindings()
	require.Len(t, findings, 1)
	assert.Equal(t,
		"file core.fga would contain multiple module definitions "+
			"(core, other, usermodule, relationmodule, ownermodule, alphamodule, zetamodule) "+
			"when transforming to DSL. Only one module can be defined per file.",
		findings[0].Message)
	assert.Equal(t, MultipleModulesInFile, findings[0].Metadata.ErrorType)
	assert.Equal(t, "core.fga", findings[0].Metadata.Symbol)
}

// TestRelationInheritsItsTypesFileAndModule covers the fallback field by field: a
// relation that names only a file belongs to its type's module, and a relation that
// names neither belongs to both of its type's.
func TestRelationInheritsItsTypesFileAndModule(t *testing.T) {
	t.Parallel()

	model := &openfgav1.AuthorizationModel{
		SchemaVersion: "1.2",
		TypeDefinitions: []*openfgav1.TypeDefinition{
			{
				Type: "org",
				Relations: map[string]*openfgav1.Userset{
					"member": {Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}}},
					"owner":  {Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}}},
				},
				Metadata: &openfgav1.Metadata{
					Module:     "core",
					SourceInfo: &openfgav1.SourceInfo{File: "core.fga"},
					Relations: map[string]*openfgav1.RelationMetadata{
						"member": {SourceInfo: &openfgav1.SourceInfo{File: "extra.fga"}},
						"owner":  {},
					},
				},
			},
		},
	}

	validator := NewMultiFileValidator(model)

	assert.Equal(t, []FileInfo{
		{Path: "core.fga", Modules: []string{"core"}},
		{Path: "extra.fga", Modules: []string{"core"}},
	}, validator.GetFileInfo())

	assert.True(t, validator.IsMultiFileProject())
	assert.False(t, validator.IsMultiModuleProject(), "both files hold the same module")
	assert.Equal(t, []string{"core.fga", "extra.fga"}, validator.GetFilesForModule("core"))

	collector := NewValidationErrors(nil)
	ValidateMultiFileConsistency(collector, model, nil)
	assert.Empty(t, collector.AllFindings(), "neither file holds more than one module")
}

func TestMultiFileValidatorReads(t *testing.T) {
	t.Parallel()

	validator := NewMultiFileValidator(multiModuleModel())

	assert.Equal(t, "core", validator.GetModuleForType("user"))
	assert.Equal(t, "other", validator.GetModuleForType("org"))
	assert.Empty(t, validator.GetModuleForType("absent"))

	assert.Equal(t, "alphamodule", validator.GetModuleForCondition("alpha"))
	assert.Equal(t, "zetamodule", validator.GetModuleForCondition("zeta"))
	assert.Empty(t, validator.GetModuleForCondition("absent"))

	assert.Equal(t, []string{
		"core", "other",
		"usermodule", "relationmodule", "ownermodule",
		"alphamodule", "zetamodule",
	}, validator.GetModulesForFile("core.fga"))
	// Empty rather than nil, so an accessor's result marshals as [] whichever key it
	// was asked for. assert.Equal tells the two apart where assert.Empty does not.
	assert.Equal(t, []string{}, validator.GetModulesForFile("absent.fga"))
	assert.Equal(t, []string{}, validator.GetFilesForModule("absent"))

	// A module reaches GetModuleInfo whether a type, a relation or a condition named
	// it, and only a type's module carries types.
	assert.Equal(t, []ModuleInfo{
		{Name: "core", Files: []string{"core.fga"}, Types: []string{"user"}},
		{Name: "other", Files: []string{"core.fga"}, Types: []string{"org"}},
		{Name: "usermodule", Files: []string{"core.fga"}, Types: []string{}},
		{Name: "relationmodule", Files: []string{"core.fga"}, Types: []string{}},
		{Name: "ownermodule", Files: []string{"core.fga"}, Types: []string{}},
		{Name: "alphamodule", Files: []string{"core.fga"}, Types: []string{}},
		{Name: "zetamodule", Files: []string{"core.fga"}, Types: []string{}},
	}, validator.GetModuleInfo())

	assert.True(t, validator.IsMultiModuleProject())
	assert.False(t, validator.IsMultiFileProject(), "one file, however many modules it holds")

	// The reads return copies: reordering what a caller was handed must not reorder
	// the record it came from.
	modules := validator.GetModulesForFile("core.fga")
	modules[0] = "clobbered"
	assert.Equal(t, "core", validator.GetModulesForFile("core.fga")[0])
}

// TestModuleInfoListsADuplicatedTypeOnce covers reading a model that declares one type
// name twice, which is itself a duplicated-error but is exactly when a caller reaches
// for GetModuleInfo. Walking the declarations reaches such a name once per declaration,
// and typeModuleMap holds one module for it, so without the guard the module that name
// resolves to lists it as many times as the model declares it.
func TestModuleInfoListsADuplicatedTypeOnce(t *testing.T) {
	t.Parallel()

	declaredIn := func(module string) *openfgav1.Metadata {
		return &openfgav1.Metadata{
			Module:     module,
			SourceInfo: &openfgav1.SourceInfo{File: "core.fga"},
		}
	}

	model := &openfgav1.AuthorizationModel{
		SchemaVersion: "1.2",
		TypeDefinitions: []*openfgav1.TypeDefinition{
			{Type: "document", Metadata: declaredIn("first")},
			{Type: "folder", Metadata: declaredIn("first")},
			{Type: "document", Metadata: declaredIn("second")},
		},
	}

	modules := NewMultiFileValidator(model).GetModuleInfo()

	// document resolves to the module that declared it last, and appears there once.
	// The module it no longer resolves to does not list it at all.
	assert.Equal(t, []ModuleInfo{
		{Name: "first", Files: []string{"core.fga"}, Types: []string{"folder"}},
		{Name: "second", Files: []string{"core.fga"}, Types: []string{"document"}},
	}, modules)

	for _, module := range modules {
		assert.Len(t, slices.Compact(slices.Clone(module.Types)), len(module.Types),
			"module %q lists a type name more than once", module.Name)
	}
}

// TestMultiFileValidatorWithoutModules covers a model that names no file or module at
// all, which is every model written as a single DSL file.
func TestMultiFileValidatorWithoutModules(t *testing.T) {
	t.Parallel()

	model := &openfgav1.AuthorizationModel{
		SchemaVersion: "1.1",
		TypeDefinitions: []*openfgav1.TypeDefinition{
			{Type: "user"},
			{
				Type:      "document",
				Relations: map[string]*openfgav1.Userset{"viewer": {Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}}}},
			},
		},
	}

	validator := NewMultiFileValidator(model)

	assert.Empty(t, validator.GetFileInfo())
	assert.Empty(t, validator.GetModuleInfo())
	assert.False(t, validator.IsMultiFileProject())
	assert.False(t, validator.IsMultiModuleProject())

	collector := NewValidationErrors(nil)
	ValidateMultiFileConsistency(collector, model, nil)
	assert.Empty(t, collector.AllFindings())

	// A nil model reaches the same entry point through the engine, and reports nothing
	// rather than panicking.
	nilCollector := NewValidationErrors(nil)
	ValidateMultiFileConsistency(nilCollector, nil, nil)
	assert.Empty(t, nilCollector.AllFindings())
	assert.Empty(t, NewMultiFileValidator(nil).GetFileInfo())
}
