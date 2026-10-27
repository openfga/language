package validation

import (
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateConditions(t *testing.T) {
	t.Parallel()

	t.Run("defined and used condition yields nothing", func(t *testing.T) {
		t.Parallel()

		dsl := `model
  schema 1.1
type user
type document
  relations
    define viewer: [user with is_valid]
condition is_valid(x: int) {
  x > 0
}`

		require.NoError(t, validateConditions(mustParse(t, dsl), newSource(dsl)))
	})

	t.Run("unused condition is reported at its declaration", func(t *testing.T) {
		t.Parallel()

		dsl := `model
  schema 1.1
type user
type document
  relations
    define viewer: [user]
condition unused(x: int) {
  x > 0
}`

		findings := ExtractAllAs[*Finding](validateConditions(mustParse(t, dsl), newSource(dsl)))

		require.Len(t, findings, 1)
		assert.Equal(t, "`unused` condition is not used in the model.", findings[0].Message)
		assert.Equal(t, ConditionNotUsed, findings[0].Metadata.Kind)
		assert.Equal(t, "unused", findings[0].Metadata.Condition)
		assert.Equal(t, &Range{Start: 6, End: 6}, findings[0].Line)
	})

	t.Run("undefined condition is reported per referencing relation", func(t *testing.T) {
		t.Parallel()

		dsl := `model
  schema 1.1
type user
type document
  relations
    define viewer: [user with missing]
    define editor: [user with missing]`

		findings := ExtractAllAs[*Finding](validateConditions(mustParse(t, dsl), newSource(dsl)))

		require.Len(t, findings, 2)

		for _, finding := range findings {
			assert.Equal(t, "`missing` is not a defined condition in the model.", finding.Message)
			assert.Equal(t, ConditionNotDefined, finding.Metadata.Kind)
			assert.Equal(t, "document", finding.Metadata.Type)
		}

		// Relations are walked in name order.
		assert.Equal(t, "editor", findings[0].Metadata.Relation)
		assert.Equal(t, "viewer", findings[1].Metadata.Relation)
	})

	t.Run("undefined conditions are reported before unused ones", func(t *testing.T) {
		t.Parallel()

		dsl := `model
  schema 1.1
type user
type document
  relations
    define viewer: [user with missing]
condition marked(x: int) {
  x > 0
}`

		findings := ExtractAllAs[*Finding](validateConditions(mustParse(t, dsl), newSource(dsl)))

		require.Len(t, findings, 2)
		assert.Equal(t, ConditionNotDefined, findings[0].Metadata.Kind)
		assert.Equal(t, ConditionNotUsed, findings[1].Metadata.Kind)
	})

	t.Run("a nested name differing from its key carries no position", func(t *testing.T) {
		t.Parallel()

		// The transformer always writes matching names, so build the mismatch
		// directly, as a hand-written JSON model could carry it.
		model := mustParse(t, `model
  schema 1.1
type user
type document
  relations
    define viewer: [user with is_valid]
condition is_valid(x: int) {
  x > 0
}`)
		model.Conditions["is_valid"].Name = "other_name"

		findings := ExtractAllAs[*Finding](validateConditions(model, source{}))

		require.Len(t, findings, 1)
		assert.Equal(t, "condition key is `is_valid` but nested name property is other_name", findings[0].Message)
		assert.Equal(t, DifferentNestedConditionName, findings[0].Metadata.Kind)
		assert.Equal(t, "other_name", findings[0].Metadata.Symbol)
		assert.Equal(t, "is_valid", findings[0].Metadata.Condition)
		assert.Nil(t, findings[0].Line)
	})

	t.Run("the reserved inline expression needs no definition", func(t *testing.T) {
		t.Parallel()

		dsl := `model
  schema 1.1
type user
type document
  relations
    define viewer: [user with $expression]`

		require.NoError(t, validateConditions(mustParse(t, dsl), newSource(dsl)))
	})

	t.Run("findings carry the condition's file and module", func(t *testing.T) {
		t.Parallel()

		model := mustParse(t, `model
  schema 1.1
type user
condition unused(x: int) {
  x > 0
}`)
		model.Conditions["unused"].Metadata = &openfgav1.ConditionMetadata{
			Module:     "core",
			SourceInfo: &openfgav1.SourceInfo{File: "core.fga"},
		}

		findings := ExtractAllAs[*Finding](validateConditions(model, source{}))

		require.Len(t, findings, 1)
		assert.Equal(t, "core.fga", findings[0].File)
		assert.Equal(t, "core", findings[0].Metadata.Module)
	})
}

func TestConditionUses(t *testing.T) {
	t.Parallel()

	dsl := `model
  schema 1.1
type user
type folder
  relations
    define owner: [user with is_valid]
type document
  relations
    define viewer: [user with is_valid, user:* with is_valid]`

	uses := conditionUses(mustParse(t, dsl))

	// One entry per referencing type restriction, types in model order.
	require.Len(t, uses, 1)
	assert.Equal(t, []conditionUse{
		{typeName: "folder", relationName: "owner"},
		{typeName: "document", relationName: "viewer"},
		{typeName: "document", relationName: "viewer"},
	}, uses["is_valid"])
}
