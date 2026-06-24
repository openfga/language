package validation

import (
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
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
						assert.Contains(t, err.Message, "is a duplicate")
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
		relationMetadata   *openfgav1.RelationMetadata
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
			relationMetadata: &openfgav1.RelationMetadata{
				DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
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
			relationMetadata: &openfgav1.RelationMetadata{
				DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
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
			relationMetadata: &openfgav1.RelationMetadata{
				DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
					{Type: "user", RelationOrWildcard: &openfgav1.RelationReference_Wildcard{Wildcard: &openfgav1.Wildcard{}}},
					{Type: "user", RelationOrWildcard: &openfgav1.RelationReference_Wildcard{Wildcard: &openfgav1.Wildcard{}}},
				},
			},
			relationName:       "viewer",
			typeName:           "document",
			expectedErrorCount: 1,
		},
		{
			name: "duplicate with relations",
			relationMetadata: &openfgav1.RelationMetadata{
				DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
					{Type: "group", RelationOrWildcard: &openfgav1.RelationReference_Relation{Relation: "member"}},
					{Type: "group", RelationOrWildcard: &openfgav1.RelationReference_Relation{Relation: "member"}},
				},
			},
			relationName:       "viewer",
			typeName:           "document",
			expectedErrorCount: 1,
		},
		{
			name: "duplicate with conditions",
			relationMetadata: &openfgav1.RelationMetadata{
				DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
					{Type: "user", Condition: "is_owner"},
					{Type: "user", Condition: "is_owner"},
				},
			},
			relationName:       "viewer",
			typeName:           "document",
			expectedErrorCount: 1,
		},
		{
			name: "no duplicates with different combinations",
			relationMetadata: &openfgav1.RelationMetadata{
				DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
					{Type: "user"},
					{Type: "user", RelationOrWildcard: &openfgav1.RelationReference_Wildcard{Wildcard: &openfgav1.Wildcard{}}},
					{Type: "user", RelationOrWildcard: &openfgav1.RelationReference_Relation{Relation: "member"}},
					{Type: "user", Condition: "is_owner"},
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
				assert.Contains(t, errors[0].Message, "is a duplicate")
			}
		})
	}
}

func TestGetRelationDefName(t *testing.T) {
	tests := []struct {
		name     string
		userset  *openfgav1.Userset
		expected string
	}{
		{
			name: "computed userset",
			userset: &openfgav1.Userset{
				Userset: &openfgav1.Userset_ComputedUserset{
					ComputedUserset: &openfgav1.ObjectRelation{
						Relation: "viewer",
					},
				},
			},
			expected: "viewer",
		},
		{
			name: "tuple to userset with target and from",
			userset: &openfgav1.Userset{
				Userset: &openfgav1.Userset_TupleToUserset{
					TupleToUserset: &openfgav1.TupleToUserset{
						ComputedUserset: &openfgav1.ObjectRelation{
							Relation: "viewer",
						},
						Tupleset: &openfgav1.ObjectRelation{
							Relation: "parent",
						},
					},
				},
			},
			expected: "viewer from parent",
		},
		{
			name: "tuple to userset with target only",
			userset: &openfgav1.Userset{
				Userset: &openfgav1.Userset_TupleToUserset{
					TupleToUserset: &openfgav1.TupleToUserset{
						ComputedUserset: &openfgav1.ObjectRelation{
							Relation: "viewer",
						},
						Tupleset: &openfgav1.ObjectRelation{},
					},
				},
			},
			expected: "viewer",
		},
		{
			name: "tuple to userset with from only",
			userset: &openfgav1.Userset{
				Userset: &openfgav1.Userset_TupleToUserset{
					TupleToUserset: &openfgav1.TupleToUserset{
						ComputedUserset: &openfgav1.ObjectRelation{},
						Tupleset: &openfgav1.ObjectRelation{
							Relation: "parent",
						},
					},
				},
			},
			expected: "",
		},
		{
			name:     "empty userset",
			userset:  &openfgav1.Userset{},
			expected: "",
		},
		{
			name: "computed userset with nil relation",
			userset: &openfgav1.Userset{
				Userset: &openfgav1.Userset_ComputedUserset{
					ComputedUserset: &openfgav1.ObjectRelation{
						Relation: "",
					},
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
		typeDef            *openfgav1.TypeDefinition
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
			typeDef: &openfgav1.TypeDefinition{
				Type:      "document",
				Relations: nil,
			},
			relationName:       "viewer",
			expectedErrorCount: 0,
		},
		{
			name: "simple relation with no duplicates",
			typeDef: &openfgav1.TypeDefinition{
				Type: "document",
				Relations: map[string]*openfgav1.Userset{
					"viewer": {
						Userset: &openfgav1.Userset_This{
							This: &openfgav1.DirectUserset{},
						},
					},
				},
			},
			relationName:       "viewer",
			expectedErrorCount: 0,
		},
		{
			name: "union with duplicates",
			typeDef: &openfgav1.TypeDefinition{
				Type: "document",
				Relations: map[string]*openfgav1.Userset{
					"viewer": {
						Userset: &openfgav1.Userset_Union{
							Union: &openfgav1.Usersets{
								Child: []*openfgav1.Userset{
									{
										Userset: &openfgav1.Userset_ComputedUserset{
											ComputedUserset: &openfgav1.ObjectRelation{
												Relation: "admin",
											},
										},
									},
									{
										Userset: &openfgav1.Userset_ComputedUserset{
											ComputedUserset: &openfgav1.ObjectRelation{
												Relation: "admin",
											},
										},
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
			typeDef: &openfgav1.TypeDefinition{
				Type: "document",
				Relations: map[string]*openfgav1.Userset{
					"can_edit": {
						Userset: &openfgav1.Userset_Intersection{
							Intersection: &openfgav1.Usersets{
								Child: []*openfgav1.Userset{
									{
										Userset: &openfgav1.Userset_ComputedUserset{
											ComputedUserset: &openfgav1.ObjectRelation{
												Relation: "admin",
											},
										},
									},
									{
										Userset: &openfgav1.Userset_ComputedUserset{
											ComputedUserset: &openfgav1.ObjectRelation{
												Relation: "admin",
											},
										},
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
		model              *openfgav1.AuthorizationModel
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
			model: &openfgav1.AuthorizationModel{
				TypeDefinitions: nil,
			},
			expectedErrorCount: 0,
		},
		{
			name: "no duplicates",
			model: &openfgav1.AuthorizationModel{
				TypeDefinitions: []*openfgav1.TypeDefinition{
					{
						Type: "document",
						Metadata: &openfgav1.Metadata{
							Relations: map[string]*openfgav1.RelationMetadata{
								"viewer": {
									DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
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
			model: &openfgav1.AuthorizationModel{
				TypeDefinitions: []*openfgav1.TypeDefinition{
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
			model: &openfgav1.AuthorizationModel{
				TypeDefinitions: []*openfgav1.TypeDefinition{
					{
						Type: "document",
						Metadata: &openfgav1.Metadata{
							Relations: map[string]*openfgav1.RelationMetadata{
								"viewer": {
									DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
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
			model: &openfgav1.AuthorizationModel{
				TypeDefinitions: []*openfgav1.TypeDefinition{
					{
						Type: "document",
						Metadata: &openfgav1.Metadata{
							Relations: map[string]*openfgav1.RelationMetadata{
								"viewer": {
									DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
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
		union              *openfgav1.Usersets
		expectedErrorCount int
	}{
		{
			name:               "nil union",
			union:              nil,
			expectedErrorCount: 0,
		},
		{
			name: "union with nil child",
			union: &openfgav1.Usersets{
				Child: nil,
			},
			expectedErrorCount: 0,
		},
		{
			name: "union with no duplicates",
			union: &openfgav1.Usersets{
				Child: []*openfgav1.Userset{
					{
						Userset: &openfgav1.Userset_ComputedUserset{
							ComputedUserset: &openfgav1.ObjectRelation{
								Relation: "admin",
							},
						},
					},
					{
						Userset: &openfgav1.Userset_ComputedUserset{
							ComputedUserset: &openfgav1.ObjectRelation{
								Relation: "viewer",
							},
						},
					},
				},
			},
			expectedErrorCount: 0,
		},
		{
			name: "union with duplicates",
			union: &openfgav1.Usersets{
				Child: []*openfgav1.Userset{
					{
						Userset: &openfgav1.Userset_ComputedUserset{
							ComputedUserset: &openfgav1.ObjectRelation{
								Relation: "admin",
							},
						},
					},
					{
						Userset: &openfgav1.Userset_ComputedUserset{
							ComputedUserset: &openfgav1.ObjectRelation{
								Relation: "admin",
							},
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
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{Type: "document"},
				{Type: "document"}, // Duplicate
			},
		}

		ValidateDuplicates(collector, model, nil)
		errors := collector.GetErrors()
		assert.Len(t, errors, 1)
		assert.Equal(t, DuplicatedError, errors[0].Metadata.ErrorType)
		assert.Contains(t, errors[0].Message, "is a duplicate")
	})

	t.Run("Duplicate Type Restriction Detection", func(t *testing.T) {
		collector := NewErrorCollector(nil)

		// Model with duplicate type restrictions in relation
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
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
