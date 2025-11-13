package validation

import (
	"testing"

	fgaSdk "github.com/openfga/go-sdk"
	"github.com/stretchr/testify/assert"
)

func TestIsValidSchemaVersion(t *testing.T) {
	tests := []struct {
		name     string
		version  string
		expected bool
	}{
		{
			name:     "version 1.1 is valid",
			version:  "1.1",
			expected: true,
		},
		{
			name:     "version 1.2 is valid",
			version:  "1.2",
			expected: true,
		},
		{
			name:     "version 1.0 is invalid",
			version:  "1.0",
			expected: false,
		},
		{
			name:     "version 2.0 is invalid",
			version:  "2.0",
			expected: false,
		},
		{
			name:     "empty string is invalid",
			version:  "",
			expected: false,
		},
		{
			name:     "random string is invalid",
			version:  "invalid",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidSchemaVersion(tt.version)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetSchemaLineNumber(t *testing.T) {
	tests := []struct {
		name          string
		schemaVersion string
		lines         []string
		expected      *int
	}{
		{
			name:          "finds schema version on line 1",
			schemaVersion: "1.1",
			lines: []string{
				"model",
				"  schema 1.1",
				"type document",
			},
			expected: fgaSdk.PtrInt(1),
		},
		{
			name:          "finds schema version with extra whitespace",
			schemaVersion: "1.2",
			lines: []string{
				"model",
				"   schema   1.2   ",
				"type document",
			},
			expected: fgaSdk.PtrInt(1),
		},
		{
			name:          "schema version not found",
			schemaVersion: "1.1",
			lines: []string{
				"model",
				"type document",
			},
			expected: nil,
		},
		{
			name:          "empty lines",
			schemaVersion: "1.1",
			lines:         []string{},
			expected:      nil,
		},
		{
			name:          "nil lines",
			schemaVersion: "1.1",
			lines:         nil,
			expected:      nil,
		},
		{
			name:          "finds first occurrence when multiple matches",
			schemaVersion: "1.1",
			lines: []string{
				"model",
				"  schema 1.1",
				"# comment about schema 1.1",
				"type document",
			},
			expected: fgaSdk.PtrInt(1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetSchemaLineNumber(tt.schemaVersion, tt.lines)
			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tt.expected, *result)
			}
		})
	}
}

func TestValidateSchemaVersion(t *testing.T) {
	tests := []struct {
		name                string
		model               *fgaSdk.AuthorizationModel
		lines               []string
		expectedErrorCount  int
		expectedErrorType   ValidationErrorType
		expectedErrorSymbol string
	}{
		{
			name:               "nil model",
			model:              nil,
			expectedErrorCount: 0,
		},
		{
			name:                "missing schema version",
			model:               &fgaSdk.AuthorizationModel{},
			expectedErrorCount:  1,
			expectedErrorType:   SchemaVersionRequired,
			expectedErrorSymbol: "",
		},
		{
			name: "empty schema version",
			model: &fgaSdk.AuthorizationModel{
				SchemaVersion: "",
			},
			expectedErrorCount:  1,
			expectedErrorType:   SchemaVersionRequired,
			expectedErrorSymbol: "",
		},
		{
			name: "valid schema version 1.1",
			model: &fgaSdk.AuthorizationModel{
				SchemaVersion: "1.1",
			},
			expectedErrorCount: 0,
		},
		{
			name: "valid schema version 1.2",
			model: &fgaSdk.AuthorizationModel{
				SchemaVersion: "1.2",
			},
			expectedErrorCount: 0,
		},
		{
			name: "invalid schema version",
			model: &fgaSdk.AuthorizationModel{
				SchemaVersion: "2.0",
			},
			lines: []string{
				"model",
				"  schema 2.0",
				"type document",
			},
			expectedErrorCount:  1,
			expectedErrorType:   SchemaVersionUnsupported,
			expectedErrorSymbol: "2.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := NewErrorCollector(tt.lines)

			ValidateSchemaVersion(collector, tt.model, tt.lines)

			errors := collector.GetErrors()
			assert.Len(t, errors, tt.expectedErrorCount)

			if tt.expectedErrorCount > 0 {
				assert.Equal(t, tt.expectedErrorType, errors[0].Metadata.ErrorType)
				assert.Equal(t, tt.expectedErrorSymbol, errors[0].Metadata.Symbol)
			}
		})
	}
}

func TestValidateMultipleModulesInFile(t *testing.T) {
	tests := []struct {
		name               string
		fileToModuleMap    map[string]map[string]bool
		expectedErrorCount int
		expectedFile       string
		expectedModules    []string
	}{
		{
			name:               "no files",
			fileToModuleMap:    map[string]map[string]bool{},
			expectedErrorCount: 0,
		},
		{
			name: "single module per file",
			fileToModuleMap: map[string]map[string]bool{
				"file1.fga": {"module1": true},
				"file2.fga": {"module2": true},
			},
			expectedErrorCount: 0,
		},
		{
			name: "multiple modules in single file",
			fileToModuleMap: map[string]map[string]bool{
				"file1.fga": {
					"module1": true,
					"module2": true,
					"module3": true,
				},
			},
			expectedErrorCount: 1,
			expectedFile:       "file1.fga",
			expectedModules:    []string{"module1", "module2", "module3"},
		},
		{
			name: "mixed: some files with single, some with multiple modules",
			fileToModuleMap: map[string]map[string]bool{
				"file1.fga": {"module1": true},
				"file2.fga": {
					"module2": true,
					"module3": true,
				},
				"file3.fga": {"module4": true},
			},
			expectedErrorCount: 1,
			expectedFile:       "file2.fga",
			expectedModules:    []string{"module2", "module3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := NewErrorCollector(nil)

			ValidateMultipleModulesInFile(collector, tt.fileToModuleMap)

			errors := collector.GetErrors()
			assert.Len(t, errors, tt.expectedErrorCount)

			if tt.expectedErrorCount > 0 {
				assert.Equal(t, MultipleModulesInFile, errors[0].Metadata.ErrorType)
				assert.Equal(t, tt.expectedFile, errors[0].Metadata.Symbol)

				// Check that all expected modules are mentioned in the error message
				for _, module := range tt.expectedModules {
					assert.Contains(t, errors[0].Message, module)
				}
			}
		})
	}
}

func TestValidateBasicModelStructure(t *testing.T) {
	tests := []struct {
		name               string
		model              *fgaSdk.AuthorizationModel
		fileToModuleMap    map[string]map[string]bool
		lines              []string
		expectedErrorCount int
	}{
		{
			name: "valid model structure",
			model: &fgaSdk.AuthorizationModel{
				SchemaVersion: "1.1",
			},
			fileToModuleMap: map[string]map[string]bool{
				"file1.fga": {"module1": true},
			},
			expectedErrorCount: 0,
		},
		{
			name:               "missing schema version",
			model:              &fgaSdk.AuthorizationModel{},
			fileToModuleMap:    map[string]map[string]bool{},
			expectedErrorCount: 1,
		},
		{
			name: "invalid schema version",
			model: &fgaSdk.AuthorizationModel{
				SchemaVersion: "2.0",
			},
			fileToModuleMap:    map[string]map[string]bool{},
			expectedErrorCount: 1,
		},
		{
			name: "multiple modules in file",
			model: &fgaSdk.AuthorizationModel{
				SchemaVersion: "1.1",
			},
			fileToModuleMap: map[string]map[string]bool{
				"file1.fga": {
					"module1": true,
					"module2": true,
				},
			},
			expectedErrorCount: 1,
		},
		{
			name:  "multiple errors",
			model: &fgaSdk.AuthorizationModel{},
			fileToModuleMap: map[string]map[string]bool{
				"file1.fga": {
					"module1": true,
					"module2": true,
				},
			},
			expectedErrorCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := NewErrorCollector(tt.lines)

			ValidateBasicModelStructure(collector, tt.model, tt.fileToModuleMap, tt.lines)

			errors := collector.GetErrors()
			assert.Len(t, errors, tt.expectedErrorCount)
		})
	}
}

func TestSupportedSchemaVersions(t *testing.T) {
	// Test that the supported schema versions map is properly initialized
	assert.NotNil(t, SupportedSchemaVersions)
	assert.True(t, SupportedSchemaVersions[SchemaVersion11])
	assert.True(t, SupportedSchemaVersions[SchemaVersion12])
	assert.False(t, SupportedSchemaVersions["1.0"])
	assert.False(t, SupportedSchemaVersions["2.0"])
}

func TestSchemaVersionConstants(t *testing.T) {
	// Test that schema version constants are defined correctly
	assert.Equal(t, "1.1", SchemaVersion11)
	assert.Equal(t, "1.2", SchemaVersion12)
}

func TestSchemaVersionValidation(t *testing.T) {
	collector := NewErrorCollector(nil)

	// Test valid schema version
	validModel := &fgaSdk.AuthorizationModel{
		SchemaVersion: "1.1",
	}
	ValidateSchemaVersion(collector, validModel, nil)
	assert.Empty(t, collector.GetErrors())

	// Test invalid schema version
	collector = NewErrorCollector(nil)
	invalidModel := &fgaSdk.AuthorizationModel{
		SchemaVersion: "2.0",
	}
	ValidateSchemaVersion(collector, invalidModel, nil)
	errors := collector.GetErrors()
	assert.Len(t, errors, 1)
	assert.Equal(t, SchemaVersionUnsupported, errors[0].Metadata.ErrorType)
}
