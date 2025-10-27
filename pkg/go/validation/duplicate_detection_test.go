package validation

import (
	"testing"

	fgaSdk "github.com/openfga/go-sdk"
	"github.com/stretchr/testify/assert"
)

func TestNewDuplicateTypeTracker(t *testing.T) {
	tracker := NewDuplicateTypeTracker()

	assert.NotNil(t, tracker)
	assert.NotNil(t, tracker.typeNames)
	assert.Empty(t, tracker.typeNames)
}

func TestDuplicateTypeTracker_CheckAndAddType(t *testing.T) {
	tests := []struct {
		name               string
		typeNames          []string
		expectedErrorCount int
		expectedDuplicate  string
	}{
		{
			name:               "no duplicates",
			typeNames:          []string{"document", "user", "group"},
			expectedErrorCount: 0,
		},
		{
			name:               "single duplicate",
			typeNames:          []string{"document", "user", "document"},
			expectedErrorCount: 1,
			expectedDuplicate:  "document",
		},
		{
			name:               "multiple duplicates",
			typeNames:          []string{"document", "user", "document", "user", "group"},
			expectedErrorCount: 2,
		},
		{
			name:               "empty type name",
			typeNames:          []string{"", "document", ""},
			expectedErrorCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tracker := NewDuplicateTypeTracker()
			collector := NewErrorCollector(nil)
			meta := &Meta{File: "test.fga", Module: "test"}

			for _, typeName := range tt.typeNames {
				tracker.CheckAndAddType(typeName, collector, meta, nil)
			}

			errors := collector.GetErrors()
			assert.Len(t, errors, tt.expectedErrorCount)

			if tt.expectedErrorCount > 0 && tt.expectedDuplicate != "" {
				found := false
				for _, err := range errors {
					if err.Metadata.Symbol == tt.expectedDuplicate {
						assert.Equal(t, DuplicatedError, err.Metadata.ErrorType)
						assert.Contains(t, err.Message, "defined more than once")
						found = true
						break
					}
				}
				assert.True(t, found, "Expected duplicate error for %s", tt.expectedDuplicate)
			}
		})
	}
}

func TestCheckForDuplicateTypeNamesInRelation(t *testing.T) {
	tests := []struct {
		name               string
		relationMetadata   *fgaSdk.RelationMetadata
		relationName       string
		typeName           string
		expectedErrorCount int
	}{
		{
			name:               "nil relation metadata",
			relationMetadata:   nil,
			expectedErrorCount: 0,
		},
		{
			name: "no duplicates",
			relationMetadata: &fgaSdk.RelationMetadata{
				DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
					{Type: "user"},
					{Type: "group"},
				},
			},
			relationName:       "viewer",
			typeName:           "document",
			expectedErrorCount: 0,
		},
		{
			name: "duplicate type restrictions",
			relationMetadata: &fgaSdk.RelationMetadata{
				DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
					{Type: "user"},
					{Type: "user"},
				},
			},
			relationName:       "viewer",
			typeName:           "document",
			expectedErrorCount: 1,
		},
		{
			name: "duplicate with wildcards",
			relationMetadata: &fgaSdk.RelationMetadata{
				DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
					{Type: "user", Wildcard: &map[string]interface{}{"type": "wildcard"}},
					{Type: "user", Wildcard: &map[string]interface{}{"type": "wildcard"}},
				},
			},
			relationName:       "viewer",
			typeName:           "document",
			expectedErrorCount: 1,
		},
		{
			name: "duplicate with relations",
			relationMetadata: &fgaSdk.RelationMetadata{
				DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
					{Type: "group", Relation: fgaSdk.PtrString("member")},
					{Type: "group", Relation: fgaSdk.PtrString("member")},
				},
			},
			relationName:       "viewer",
			typeName:           "document",
			expectedErrorCount: 1,
		},
		{
			name: "duplicate with conditions",
			relationMetadata: &fgaSdk.RelationMetadata{
				DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
					{Type: "user", Condition: fgaSdk.PtrString("is_owner")},
					{Type: "user", Condition: fgaSdk.PtrString("is_owner")},
				},
			},
			relationName:       "viewer",
			typeName:           "document",
			expectedErrorCount: 1,
		},
		{
			name: "no duplicates with different combinations",
			relationMetadata: &fgaSdk.RelationMetadata{
				DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
					{Type: "user"},
					{Type: "user", Wildcard: &map[string]interface{}{"type": "wildcard"}},
					{Type: "user", Relation: fgaSdk.PtrString("member")},
					{Type: "user", Condition: fgaSdk.PtrString("is_owner")},
				},
			},
			relationName:       "viewer",
			typeName:           "document",
			expectedErrorCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := NewErrorCollector(nil)
			meta := &Meta{File: "test.fga", Module: "test"}

			CheckForDuplicateTypeNamesInRelation(collector, tt.relationMetadata, tt.relationName, tt.typeName, meta, nil)

			errors := collector.GetErrors()
			assert.Len(t, errors, tt.expectedErrorCount)

			if tt.expectedErrorCount > 0 {
				assert.Equal(t, DuplicatedError, errors[0].Metadata.ErrorType)
				assert.Contains(t, errors[0].Message, "defined more than once")
			}
		})
	}
}

func TestGetRelationDefName(t *testing.T) {
	tests := []struct {
		name     string
		userset  fgaSdk.Userset
		expected string
	}{
		{
			name: "computed userset",
			userset: fgaSdk.Userset{
				ComputedUserset: &fgaSdk.ObjectRelation{
					Relation: fgaSdk.PtrString("viewer"),
				},
			},
			expected: "viewer",
		},
		{
			name: "tuple to userset with target and from",
			userset: fgaSdk.Userset{
				TupleToUserset: &fgaSdk.TupleToUserset{
					ComputedUserset: fgaSdk.ObjectRelation{
						Relation: fgaSdk.PtrString("viewer"),
					},
					Tupleset: fgaSdk.ObjectRelation{
						Relation: fgaSdk.PtrString("parent"),
					},
				},
			},
			expected: "viewer from parent",
		},
		{
			name: "tuple to userset with target only",
			userset: fgaSdk.Userset{
				TupleToUserset: &fgaSdk.TupleToUserset{
					ComputedUserset: fgaSdk.ObjectRelation{
						Relation: fgaSdk.PtrString("viewer"),
					},
					Tupleset: fgaSdk.ObjectRelation{},
				},
			},
			expected: "viewer",
		},
		{
			name: "tuple to userset with from only",
			userset: fgaSdk.Userset{
				TupleToUserset: &fgaSdk.TupleToUserset{
					ComputedUserset: fgaSdk.ObjectRelation{},
					Tupleset: fgaSdk.ObjectRelation{
						Relation: fgaSdk.PtrString("parent"),
					},
				},
			},
			expected: "",
		},
		{
			name:     "empty userset",
			userset:  fgaSdk.Userset{},
			expected: "",
		},
		{
			name: "computed userset with nil relation",
			userset: fgaSdk.Userset{
				ComputedUserset: &fgaSdk.ObjectRelation{
					Relation: nil,
				},
			},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getRelationDefName(tt.userset)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCheckForDuplicatesInRelation(t *testing.T) {
	tests := []struct {
		name               string
		typeDef            *fgaSdk.TypeDefinition
		relationName       string
		expectedErrorCount int
	}{
		{
			name:               "nil type definition",
			typeDef:            nil,
			relationName:       "viewer",
			expectedErrorCount: 0,
		},
		{
			name: "nil relations",
			typeDef: &fgaSdk.TypeDefinition{
				Type:      "document",
				Relations: nil,
			},
			relationName:       "viewer",
			expectedErrorCount: 0,
		},
		{
			name: "simple relation with no duplicates",
			typeDef: &fgaSdk.TypeDefinition{
				Type: "document",
				Relations: &map[string]fgaSdk.Userset{
					"viewer": {
						This: &map[string]interface{}{},
					},
				},
			},
			relationName:       "viewer",
			expectedErrorCount: 0,
		},
		{
			name: "union with duplicates",
			typeDef: &fgaSdk.TypeDefinition{
				Type: "document",
				Relations: &map[string]fgaSdk.Userset{
					"viewer": {
						Union: &fgaSdk.Usersets{
							Child: []fgaSdk.Userset{
								{
									ComputedUserset: &fgaSdk.ObjectRelation{
										Relation: fgaSdk.PtrString("admin"),
									},
								},
								{
									ComputedUserset: &fgaSdk.ObjectRelation{
										Relation: fgaSdk.PtrString("admin"),
									},
								},
							},
						},
					},
				},
			},
			relationName:       "viewer",
			expectedErrorCount: 1,
		},
		{
			name: "intersection with duplicates",
			typeDef: &fgaSdk.TypeDefinition{
				Type: "document",
				Relations: &map[string]fgaSdk.Userset{
					"can_edit": {
						Intersection: &fgaSdk.Usersets{
							Child: []fgaSdk.Userset{
								{
									ComputedUserset: &fgaSdk.ObjectRelation{
										Relation: fgaSdk.PtrString("admin"),
									},
								},
								{
									ComputedUserset: &fgaSdk.ObjectRelation{
										Relation: fgaSdk.PtrString("admin"),
									},
								},
							},
						},
					},
				},
			},
			relationName:       "can_edit",
			expectedErrorCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := NewErrorCollector(nil)

			CheckForDuplicatesInRelation(collector, tt.typeDef, tt.relationName, nil)

			errors := collector.GetErrors()
			assert.Len(t, errors, tt.expectedErrorCount)

			if tt.expectedErrorCount > 0 {
				assert.Equal(t, DuplicatedError, errors[0].Metadata.ErrorType)
			}
		})
	}
}

func TestValidateDuplicates(t *testing.T) {
	tests := []struct {
		name               string
		model              *fgaSdk.AuthorizationModel
		expectedErrorCount int
		expectedErrorTypes []ValidationErrorType
	}{
		{
			name:               "nil model",
			model:              nil,
			expectedErrorCount: 0,
		},
		{
			name: "nil type definitions",
			model: &fgaSdk.AuthorizationModel{
				TypeDefinitions: nil,
			},
			expectedErrorCount: 0,
		},
		{
			name: "no duplicates",
			model: &fgaSdk.AuthorizationModel{
				TypeDefinitions: []fgaSdk.TypeDefinition{
					{
						Type: "document",
						Metadata: &fgaSdk.Metadata{
							Relations: &map[string]fgaSdk.RelationMetadata{
								"viewer": {
									DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
										{Type: "user"},
									},
								},
							},
						},
					},
					{
						Type: "user",
					},
				},
			},
			expectedErrorCount: 0,
		},
		{
			name: "duplicate type names",
			model: &fgaSdk.AuthorizationModel{
				TypeDefinitions: []fgaSdk.TypeDefinition{
					{
						Type: "document",
					},
					{
						Type: "document",
					},
				},
			},
			expectedErrorCount: 1,
			expectedErrorTypes: []ValidationErrorType{DuplicatedError},
		},
		{
			name: "duplicate type restrictions in relation",
			model: &fgaSdk.AuthorizationModel{
				TypeDefinitions: []fgaSdk.TypeDefinition{
					{
						Type: "document",
						Metadata: &fgaSdk.Metadata{
							Relations: &map[string]fgaSdk.RelationMetadata{
								"viewer": {
									DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
										{Type: "user"},
										{Type: "user"},
									},
								},
							},
						},
					},
				},
			},
			expectedErrorCount: 1,
			expectedErrorTypes: []ValidationErrorType{DuplicatedError},
		},
		{
			name: "multiple types of duplicates",
			model: &fgaSdk.AuthorizationModel{
				TypeDefinitions: []fgaSdk.TypeDefinition{
					{
						Type: "document",
						Metadata: &fgaSdk.Metadata{
							Relations: &map[string]fgaSdk.RelationMetadata{
								"viewer": {
									DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
										{Type: "user"},
										{Type: "user"},
									},
								},
							},
						},
					},
					{
						Type: "document", // Duplicate type name
					},
				},
			},
			expectedErrorCount: 2,
			expectedErrorTypes: []ValidationErrorType{DuplicatedError, DuplicatedError},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := NewErrorCollector(nil)

			ValidateDuplicates(collector, tt.model, nil)

			errors := collector.GetErrors()
			assert.Len(t, errors, tt.expectedErrorCount)

			for i, expectedType := range tt.expectedErrorTypes {
				if i < len(errors) {
					assert.Equal(t, expectedType, errors[i].Metadata.ErrorType)
				}
			}
		})
	}
}

func TestCheckDuplicatesInUnion(t *testing.T) {
	tests := []struct {
		name               string
		union              *fgaSdk.Usersets
		expectedErrorCount int
	}{
		{
			name:               "nil union",
			union:              nil,
			expectedErrorCount: 0,
		},
		{
			name: "union with nil child",
			union: &fgaSdk.Usersets{
				Child: nil,
			},
			expectedErrorCount: 0,
		},
		{
			name: "union with no duplicates",
			union: &fgaSdk.Usersets{
				Child: []fgaSdk.Userset{
					{
						ComputedUserset: &fgaSdk.ObjectRelation{
							Relation: fgaSdk.PtrString("admin"),
						},
					},
					{
						ComputedUserset: &fgaSdk.ObjectRelation{
							Relation: fgaSdk.PtrString("viewer"),
						},
					},
				},
			},
			expectedErrorCount: 0,
		},
		{
			name: "union with duplicates",
			union: &fgaSdk.Usersets{
				Child: []fgaSdk.Userset{
					{
						ComputedUserset: &fgaSdk.ObjectRelation{
							Relation: fgaSdk.PtrString("admin"),
						},
					},
					{
						ComputedUserset: &fgaSdk.ObjectRelation{
							Relation: fgaSdk.PtrString("admin"),
						},
					},
				},
			},
			expectedErrorCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := NewErrorCollector(nil)
			meta := &Meta{File: "test.fga", Module: "test"}

			checkDuplicatesInUnion(collector, tt.union, "test_relation", "test_type", meta, nil)

			errors := collector.GetErrors()
			assert.Len(t, errors, tt.expectedErrorCount)

			if tt.expectedErrorCount > 0 {
				assert.Equal(t, DuplicatedError, errors[0].Metadata.ErrorType)
			}
		})
	}
}

func TestValidateDuplicates_Integration(t *testing.T) {
	t.Run("Duplicate Type Detection", func(t *testing.T) {
		collector := NewErrorCollector(nil)

		// Model with duplicate type names
		model := &fgaSdk.AuthorizationModel{
			TypeDefinitions: []fgaSdk.TypeDefinition{
				{Type: "document"},
				{Type: "document"}, // Duplicate
			},
		}

		ValidateDuplicates(collector, model, nil)
		errors := collector.GetErrors()
		assert.Len(t, errors, 1)
		assert.Equal(t, DuplicatedError, errors[0].Metadata.ErrorType)
		assert.Contains(t, errors[0].Message, "defined more than once")
	})

	t.Run("Duplicate Type Restriction Detection", func(t *testing.T) {
		collector := NewErrorCollector(nil)

		// Model with duplicate type restrictions in relation
		model := &fgaSdk.AuthorizationModel{
			TypeDefinitions: []fgaSdk.TypeDefinition{
				{
					Type: "document",
					Metadata: &fgaSdk.Metadata{
						Relations: &map[string]fgaSdk.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
									{Type: "user"},
									{Type: "user"}, // Duplicate
								},
							},
						},
					},
				},
			},
		}

		ValidateDuplicates(collector, model, nil)
		errors := collector.GetErrors()
		assert.Len(t, errors, 1)
		assert.Equal(t, DuplicatedError, errors[0].Metadata.ErrorType)
	})
}
