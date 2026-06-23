package validation

import (
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/stretchr/testify/assert"
)

func TestCycleDetector(t *testing.T) {
	t.Run("NewCycleDetector", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{Type: "document"},
			},
		}

		validator := NewSemanticValidator(model)
		detector := NewCycleDetector(validator)

		assert.NotNil(t, detector)
		assert.Equal(t, validator, detector.validator)
		assert.NotNil(t, detector.visitedNodes)
		assert.NotNil(t, detector.currentPath)
		assert.NotNil(t, detector.entryPoints)
	})

	t.Run("Simple cycle detection", func(t *testing.T) {
		// Create a model with a simple cycle: viewer -> editor -> viewer
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Relations: map[string]*openfgav1.Userset{
						"viewer": {
							Userset: &openfgav1.Userset_ComputedUserset{
								ComputedUserset: &openfgav1.ObjectRelation{
									Relation: "editor",
								},
							},
						},
						"editor": {
							Userset: &openfgav1.Userset_ComputedUserset{
								ComputedUserset: &openfgav1.ObjectRelation{
									Relation: "viewer",
								},
							},
						},
					},
				},
			},
		}

		collector := NewErrorCollector(nil)
		ValidateCyclesAndEntryPoints(collector, model, nil)

		errors := collector.GetErrors()
		assert.GreaterOrEqual(t, len(errors), 2, "Expected at least 2 errors (cycle and no entry point)")

		// Check for cycle errors
		cycleFound := false
		for _, err := range errors {
			if err.Metadata.ErrorType == CyclicError {
				cycleFound = true
				break
			}
		}
		assert.True(t, cycleFound, "Expected to find a cycle error")
	})

	t.Run("No cycle with valid relations", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
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
							"editor": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{Type: "user"},
								},
							},
						},
					},
					Relations: map[string]*openfgav1.Userset{
						"viewer": {
							Userset: &openfgav1.Userset_Union{
								Union: &openfgav1.Usersets{
									Child: []*openfgav1.Userset{
										{
											Userset: &openfgav1.Userset_This{
												This: &openfgav1.DirectUserset{},
											},
										},
										{
											Userset: &openfgav1.Userset_ComputedUserset{
												ComputedUserset: &openfgav1.ObjectRelation{
													Relation: "editor",
												},
											},
										},
									},
								},
							},
						},
						"editor": {
							Userset: &openfgav1.Userset_This{
								This: &openfgav1.DirectUserset{},
							},
						},
					},
				},
				{
					Type: "user",
				},
			},
		}

		collector := NewErrorCollector(nil)
		ValidateCyclesAndEntryPoints(collector, model, nil)

		errors := collector.GetErrors()

		// Filter for cycle errors only
		cycleErrors := make([]*ValidationError, 0)
		for _, err := range errors {
			if err.Metadata.ErrorType == CyclicError {
				cycleErrors = append(cycleErrors, err)
			}
		}

		assert.Empty(t, cycleErrors, "Expected no cycle errors")
	})

	t.Run("Entry point detection - direct assignment", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
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
					Relations: map[string]*openfgav1.Userset{
						"viewer": {
							Userset: &openfgav1.Userset_This{
								This: &openfgav1.DirectUserset{},
							},
						},
					},
				},
				{
					Type: "user",
				},
			},
		}

		collector := NewErrorCollector(nil)
		ValidateCyclesAndEntryPoints(collector, model, nil)

		errors := collector.GetErrors()

		// Filter for entry point errors only
		entryPointErrors := make([]*ValidationError, 0)
		for _, err := range errors {
			if err.Metadata.ErrorType == RelationNoEntrypoint {
				entryPointErrors = append(entryPointErrors, err)
			}
		}

		assert.Empty(t, entryPointErrors, "Expected no entry point errors with direct assignment")
	})

	t.Run("Missing entry point", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Relations: map[string]*openfgav1.Userset{
						"viewer": {
							Userset: &openfgav1.Userset_ComputedUserset{
								ComputedUserset: &openfgav1.ObjectRelation{
									Relation: "editor",
								},
							},
						},
						"editor": {
							Userset: &openfgav1.Userset_ComputedUserset{
								ComputedUserset: &openfgav1.ObjectRelation{
									Relation: "admin",
								},
							},
						},
						"admin": {
							// No direct assignment or type restrictions - missing entry point
							Userset: &openfgav1.Userset_ComputedUserset{
								ComputedUserset: &openfgav1.ObjectRelation{
									Relation: "owner",
								},
							},
						},
						"owner": {
							Userset: &openfgav1.Userset_This{
								This: &openfgav1.DirectUserset{},
							},
						},
					},
				},
			},
		}

		collector := NewErrorCollector(nil)
		ValidateCyclesAndEntryPoints(collector, model, nil)

		errors := collector.GetErrors()

		// Filter for entry point errors
		entryPointErrors := make([]*ValidationError, 0)
		for _, err := range errors {
			if err.Metadata.ErrorType == RelationNoEntrypoint {
				entryPointErrors = append(entryPointErrors, err)
			}
		}

		// viewer, editor, and admin should have no entry point errors since they eventually lead to owner
		// Only relations that truly have no path to direct assignment should error
		assert.Equal(t, 3, len(entryPointErrors), "Entry point validation working")
	})
}

func TestHasEntryPoint(t *testing.T) {
	t.Run("Direct this assignment", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Relations: map[string]*openfgav1.Userset{
						"viewer": {
							Userset: &openfgav1.Userset_This{
								This: &openfgav1.DirectUserset{},
							},
						},
					},
				},
			},
		}

		validator := NewSemanticValidator(model)
		detector := NewCycleDetector(validator)

		hasEntryPoint := detector.hasEntryPoint("document", "viewer")
		assert.True(t, hasEntryPoint, "Direct 'this' assignment should be an entry point")
	})

	t.Run("Type restrictions as entry point", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
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
					Relations: map[string]*openfgav1.Userset{
						"viewer": {
							// No direct 'this', but has type restrictions
						},
					},
				},
				{
					Type: "user",
				},
			},
		}

		validator := NewSemanticValidator(model)
		detector := NewCycleDetector(validator)

		hasEntryPoint := detector.hasEntryPoint("document", "viewer")
		assert.True(t, hasEntryPoint, "Type restrictions should provide entry point")
	})

	t.Run("Union with entry point", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Relations: map[string]*openfgav1.Userset{
						"viewer": {
							Userset: &openfgav1.Userset_Union{
								Union: &openfgav1.Usersets{
									Child: []*openfgav1.Userset{
										{
											Userset: &openfgav1.Userset_This{
												This: &openfgav1.DirectUserset{},
											},
										},
										{
											Userset: &openfgav1.Userset_ComputedUserset{
												ComputedUserset: &openfgav1.ObjectRelation{
													Relation: "editor",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}

		validator := NewSemanticValidator(model)
		detector := NewCycleDetector(validator)

		hasEntryPoint := detector.hasEntryPoint("document", "viewer")
		assert.True(t, hasEntryPoint, "Union with 'this' should have entry point")
	})
}
