package validation

import (
	"testing"

	fgaSdk "github.com/openfga/go-sdk"
	"github.com/stretchr/testify/assert"
)

func TestValidationRegexRules(t *testing.T) {
	// Test that regex rules match the JS implementation
	assert.Equal(t, "[^:#@\\*\\s]{1,254}", ValidationRegexRules.Type)
	assert.Equal(t, "[^:#@\\*\\s]{1,50}", ValidationRegexRules.Relation)
	assert.Equal(t, "[^\\*\\s]{1,50}", ValidationRegexRules.Condition)
	assert.Equal(t, "[^#:\\s*][a-zA-Z0-9_|*@.+]*", ValidationRegexRules.ID)
	assert.Equal(t, "[^\\s]{2,256}", ValidationRegexRules.Object)
}

func TestValidateFieldValue(t *testing.T) {
	tests := []struct {
		name     string
		rule     string
		value    string
		expected bool
	}{
		{
			name:     "valid type name",
			rule:     "^[a-zA-Z]+$",
			value:    "document",
			expected: true,
		},
		{
			name:     "invalid type name with numbers",
			rule:     "^[a-zA-Z]+$",
			value:    "document123",
			expected: false,
		},
		{
			name:     "valid relation name",
			rule:     "^[a-zA-Z_]+$",
			value:    "can_view",
			expected: true,
		},
		{
			name:     "empty string with appropriate rule",
			rule:     "^$",
			value:    "",
			expected: true,
		},
		{
			name:     "invalid regex pattern",
			rule:     "[unclosed",
			value:    "test",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validateFieldValue(tt.rule, tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestValidateTypeName(t *testing.T) {
	tests := []struct {
		name               string
		typeName           string
		expectedValid      bool
		expectedErrorType  ValidationErrorType
		expectedErrorCount int
	}{
		{
			name:               "valid type name",
			typeName:           "document",
			expectedValid:      true,
			expectedErrorCount: 0,
		},
		{
			name:               "reserved keyword self",
			typeName:           "self",
			expectedValid:      false,
			expectedErrorType:  ReservedTypeKeywords,
			expectedErrorCount: 1,
		},
		{
			name:               "reserved keyword this",
			typeName:           "this",
			expectedValid:      false,
			expectedErrorType:  ReservedTypeKeywords,
			expectedErrorCount: 1,
		},
		{
			name:               "valid type name with underscore",
			typeName:           "user_group",
			expectedValid:      true,
			expectedErrorCount: 0,
		},
		{
			name:               "type name with invalid characters",
			typeName:           "document:invalid",
			expectedValid:      false,
			expectedErrorType:  InvalidName,
			expectedErrorCount: 1,
		},
		{
			name:               "type name with space",
			typeName:           "document name",
			expectedValid:      false,
			expectedErrorType:  InvalidName,
			expectedErrorCount: 1,
		},
		{
			name:               "type name with hash",
			typeName:           "document#tag",
			expectedValid:      false,
			expectedErrorType:  InvalidName,
			expectedErrorCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := NewErrorCollector(nil)
			lineIndex := 5
			meta := &Meta{File: "test.fga", Module: "test"}

			result := ValidateTypeName(tt.typeName, collector, &lineIndex, meta)

			assert.Equal(t, tt.expectedValid, result)

			errors := collector.GetErrors()
			assert.Len(t, errors, tt.expectedErrorCount)

			if tt.expectedErrorCount > 0 {
				assert.Equal(t, tt.expectedErrorType, errors[0].Metadata.ErrorType)
				assert.Equal(t, tt.typeName, errors[0].Metadata.Symbol)
			}
		})
	}
}

func TestValidateRelationName(t *testing.T) {
	tests := []struct {
		name               string
		relationName       string
		typeName           string
		expectedValid      bool
		expectedErrorType  ValidationErrorType
		expectedErrorCount int
	}{
		{
			name:               "valid relation name",
			relationName:       "viewer",
			typeName:           "document",
			expectedValid:      true,
			expectedErrorCount: 0,
		},
		{
			name:               "reserved keyword self",
			relationName:       "self",
			typeName:           "document",
			expectedValid:      false,
			expectedErrorType:  ReservedRelationKeywords,
			expectedErrorCount: 1,
		},
		{
			name:               "reserved keyword this",
			relationName:       "this",
			typeName:           "document",
			expectedValid:      false,
			expectedErrorType:  ReservedRelationKeywords,
			expectedErrorCount: 1,
		},
		{
			name:               "valid relation name with underscore",
			relationName:       "can_view",
			typeName:           "document",
			expectedValid:      true,
			expectedErrorCount: 0,
		},
		{
			name:               "relation name with invalid characters",
			relationName:       "viewer:invalid",
			typeName:           "document",
			expectedValid:      false,
			expectedErrorType:  InvalidName,
			expectedErrorCount: 1,
		},
		{
			name:               "relation name with space",
			relationName:       "can view",
			typeName:           "document",
			expectedValid:      false,
			expectedErrorType:  InvalidName,
			expectedErrorCount: 1,
		},
		{
			name:               "relation name with hash",
			relationName:       "view#tag",
			typeName:           "document",
			expectedValid:      false,
			expectedErrorType:  InvalidName,
			expectedErrorCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := NewErrorCollector(nil)
			lineIndex := 8
			meta := &Meta{File: "test.fga", Module: "test"}

			result := ValidateRelationName(tt.relationName, tt.typeName, collector, &lineIndex, meta)

			assert.Equal(t, tt.expectedValid, result)

			errors := collector.GetErrors()
			assert.Len(t, errors, tt.expectedErrorCount)

			if tt.expectedErrorCount > 0 {
				assert.Equal(t, tt.expectedErrorType, errors[0].Metadata.ErrorType)
				assert.Equal(t, tt.relationName, errors[0].Metadata.Symbol)

				// Check that error message includes type name for relation errors
				if tt.expectedErrorType == InvalidName {
					assert.Contains(t, errors[0].Message, tt.typeName)
				}
			}
		})
	}
}

func TestValidateConditionName(t *testing.T) {
	tests := []struct {
		name               string
		conditionName      string
		expectedValid      bool
		expectedErrorCount int
	}{
		{
			name:               "valid condition name",
			conditionName:      "is_owner",
			expectedValid:      true,
			expectedErrorCount: 0,
		},
		{
			name:               "valid condition name with numbers",
			conditionName:      "condition123",
			expectedValid:      true,
			expectedErrorCount: 0,
		},
		{
			name:               "condition name with space",
			conditionName:      "is owner",
			expectedValid:      false,
			expectedErrorCount: 1,
		},
		{
			name:               "condition name with asterisk",
			conditionName:      "condition*",
			expectedValid:      false,
			expectedErrorCount: 1,
		},
		{
			name:               "empty condition name",
			conditionName:      "",
			expectedValid:      false,
			expectedErrorCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := NewErrorCollector(nil)
			lineIndex := 10
			meta := &Meta{File: "test.fga", Module: "test"}

			result := ValidateConditionName(tt.conditionName, collector, &lineIndex, meta)

			assert.Equal(t, tt.expectedValid, result)

			errors := collector.GetErrors()
			assert.Len(t, errors, tt.expectedErrorCount)

			if tt.expectedErrorCount > 0 {
				assert.Equal(t, InvalidName, errors[0].Metadata.ErrorType)
				assert.Equal(t, tt.conditionName, errors[0].Metadata.Symbol)
			}
		})
	}
}

func TestGetTypeLineNumber(t *testing.T) {
	tests := []struct {
		name      string
		typeName  string
		lines     []string
		skipIndex *int
		expected  *int
	}{
		{
			name:     "finds type on line 0",
			typeName: "document",
			lines: []string{
				"type document",
				"  relations",
				"    define viewer: [user]",
			},
			expected: fgaSdk.PtrInt(0),
		},
		{
			name:     "finds type on line 2",
			typeName: "user",
			lines: []string{
				"model",
				"  schema 1.1",
				"type user",
				"type document",
			},
			expected: fgaSdk.PtrInt(2),
		},
		{
			name:     "type not found",
			typeName: "nonexistent",
			lines: []string{
				"type document",
				"type user",
			},
			expected: nil,
		},
		{
			name:     "skips specified index",
			typeName: "document",
			lines: []string{
				"type document",
				"  relations",
				"type document",
			},
			skipIndex: fgaSdk.PtrInt(0),
			expected:  fgaSdk.PtrInt(2),
		},
		{
			name:     "empty lines",
			typeName: "document",
			lines:    []string{},
			expected: nil,
		},
		{
			name:     "nil lines",
			typeName: "document",
			lines:    nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetTypeLineNumber(tt.typeName, tt.lines, tt.skipIndex)
			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tt.expected, *result)
			}
		})
	}
}

func TestGetRelationLineNumber(t *testing.T) {
	tests := []struct {
		name         string
		relationName string
		lines        []string
		skipIndex    *int
		expected     *int
	}{
		{
			name:         "finds relation on line 2",
			relationName: "viewer",
			lines: []string{
				"type document",
				"  relations",
				"    define viewer: [user]",
				"    define admin: [user]",
			},
			expected: fgaSdk.PtrInt(2),
		},
		{
			name:         "finds relation with complex definition",
			relationName: "can_view",
			lines: []string{
				"type document",
				"  relations",
				"    define viewer: [user]",
				"    define can_view: viewer or admin",
			},
			expected: fgaSdk.PtrInt(3),
		},
		{
			name:         "relation not found",
			relationName: "nonexistent",
			lines: []string{
				"type document",
				"  relations",
				"    define viewer: [user]",
			},
			expected: nil,
		},
		{
			name:         "skips specified index",
			relationName: "viewer",
			lines: []string{
				"    define viewer: [user]",
				"  relations",
				"    define viewer: [group]",
			},
			skipIndex: fgaSdk.PtrInt(0),
			expected:  fgaSdk.PtrInt(2),
		},
		{
			name:         "empty lines",
			relationName: "viewer",
			lines:        []string{},
			expected:     nil,
		},
		{
			name:         "nil lines",
			relationName: "viewer",
			lines:        nil,
			expected:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetRelationLineNumber(tt.relationName, tt.lines, tt.skipIndex)
			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tt.expected, *result)
			}
		})
	}
}

func TestGetConditionLineNumber(t *testing.T) {
	tests := []struct {
		name          string
		conditionName string
		lines         []string
		skipIndex     *int
		expected      *int
	}{
		{
			name:          "finds condition in line",
			conditionName: "is_owner",
			lines: []string{
				"type document",
				"  relations",
				"    define viewer: [user with is_owner]",
			},
			expected: fgaSdk.PtrInt(2),
		},
		{
			name:          "condition not found",
			conditionName: "nonexistent",
			lines: []string{
				"type document",
				"  relations",
				"    define viewer: [user]",
			},
			expected: nil,
		},
		{
			name:          "skips specified index",
			conditionName: "is_owner",
			lines: []string{
				"    define viewer: [user with is_owner]",
				"  relations",
				"    condition is_owner: some_condition",
			},
			skipIndex: fgaSdk.PtrInt(0),
			expected:  fgaSdk.PtrInt(2),
		},
		{
			name:          "empty lines",
			conditionName: "is_owner",
			lines:         []string{},
			expected:      nil,
		},
		{
			name:          "nil lines",
			conditionName: "is_owner",
			lines:         nil,
			expected:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetConditionLineNumber(tt.conditionName, tt.lines, tt.skipIndex)
			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tt.expected, *result)
			}
		})
	}
}

func TestValidateNameRules(t *testing.T) {
	tests := []struct {
		name               string
		typeName           string
		relationNames      []string
		lines              []string
		expectedErrorCount int
	}{
		{
			name:          "valid type and relations",
			typeName:      "document",
			relationNames: []string{"viewer", "admin", "owner"},
			lines: []string{
				"type document",
				"  relations",
				"    define viewer: [user]",
				"    define admin: [user] ",
				"    define owner: [user]",
			},
			expectedErrorCount: 0,
		},
		{
			name:          "reserved type name",
			typeName:      "self",
			relationNames: []string{"viewer"},
			lines: []string{
				"type self",
				"  relations",
				"    define viewer: [user]",
			},
			expectedErrorCount: 1,
		},
		{
			name:          "reserved relation name",
			typeName:      "document",
			relationNames: []string{"this", "viewer"},
			lines: []string{
				"type document",
				"  relations",
				"    define this: [user]",
				"    define viewer: [user]",
			},
			expectedErrorCount: 1,
		},
		{
			name:          "multiple validation errors",
			typeName:      "self",
			relationNames: []string{"this", "viewer"},
			lines: []string{
				"type self",
				"  relations",
				"    define this: [user]",
				"    define viewer: [user]",
			},
			expectedErrorCount: 2,
		},
		{
			name:          "invalid type name characters",
			typeName:      "document:invalid",
			relationNames: []string{"viewer"},
			lines: []string{
				"type document:invalid",
				"  relations",
				"    define viewer: [user]",
			},
			expectedErrorCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := NewErrorCollector(tt.lines)
			typeLineIndex := GetTypeLineNumber(tt.typeName, tt.lines, nil)
			meta := &Meta{File: "test.fga", Module: "test"}

			ValidateNameRules(collector, tt.typeName, tt.relationNames, typeLineIndex, meta, tt.lines)

			errors := collector.GetErrors()
			assert.Len(t, errors, tt.expectedErrorCount)
		})
	}
}

// TestNameValidationIntegration tests name validation with SDK types.
func TestNameValidationIntegration(t *testing.T) {
	t.Run("Valid Names", func(t *testing.T) {
		validNames := []string{"document", "user", "group", "viewer", "editor", "admin"}

		for _, name := range validNames {
			collector := NewErrorCollector(nil)
			typeValid := ValidateTypeName(name, collector, nil, nil)
			assert.True(t, typeValid, "Expected %s to be valid type name", name)

			collector = NewErrorCollector(nil)
			relationValid := ValidateRelationName(name, "parent_type", collector, nil, nil)
			assert.True(t, relationValid, "Expected %s to be valid relation name", name)
		}
	})

	t.Run("Reserved Keywords", func(t *testing.T) {
		reservedKeywords := []string{"this", "self"}

		for _, keyword := range reservedKeywords {
			collector := NewErrorCollector(nil)
			typeValid := ValidateTypeName(keyword, collector, nil, nil)
			assert.False(t, typeValid, "Expected %s to be invalid type name", keyword)

			collector = NewErrorCollector(nil)
			relationValid := ValidateRelationName(keyword, "parent_type", collector, nil, nil)
			assert.False(t, relationValid, "Expected %s to be invalid relation name", keyword)
		}
	})
}
