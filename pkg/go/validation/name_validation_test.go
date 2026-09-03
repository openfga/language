package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateTypeName(t *testing.T) {
	t.Parallel()

	t.Run("valid name yields nothing", func(t *testing.T) {
		t.Parallel()
		assert.Nil(t, validateTypeName("document"))
	})

	t.Run("reserved keywords", func(t *testing.T) {
		t.Parallel()

		for _, reserved := range []string{"self", "this"} {
			finding := validateTypeName(reserved)

			require.NotNil(t, finding)
			assert.Equal(t, ReservedTypeKeywords, finding.Metadata.Kind)
			assert.Equal(t, "a type cannot be named 'self' or 'this'.", finding.Message)
			assert.Equal(t, reserved, finding.Metadata.Symbol)
			assert.Equal(t, reserved, finding.Metadata.Type)
		}
	})

	t.Run("rule violation quotes the anchored rule", func(t *testing.T) {
		t.Parallel()

		finding := validateTypeName("doc:ument")

		require.NotNil(t, finding)
		assert.Equal(t, InvalidName, finding.Metadata.Kind)
		assert.Equal(t, "type 'doc:ument' does not match naming rule: '^[^:#@\\*\\s]{1,254}$'.", finding.Message)
	})
}

func TestValidateRelationName(t *testing.T) {
	t.Parallel()

	t.Run("valid name yields nothing", func(t *testing.T) {
		t.Parallel()
		assert.Nil(t, validateRelationName("viewer", "document"))
	})

	t.Run("reserved keyword", func(t *testing.T) {
		t.Parallel()

		finding := validateRelationName("self", "document")

		require.NotNil(t, finding)
		assert.Equal(t, ReservedRelationKeywords, finding.Metadata.Kind)
		assert.Equal(t, "a relation cannot be named 'self' or 'this'.", finding.Message)
		assert.Equal(t, "document", finding.Metadata.Type)
		assert.Equal(t, "self", finding.Metadata.Relation)
	})

	t.Run("rule violation names the relation and its type", func(t *testing.T) {
		t.Parallel()

		finding := validateRelationName("view#er", "document")

		require.NotNil(t, finding)
		assert.Equal(t, InvalidName, finding.Metadata.Kind)
		assert.Equal(t,
			"relation 'view#er' of type 'document' does not match naming rule: '^[^:#@\\*\\s]{1,50}$'.",
			finding.Message)
	})
}

func TestValidateConditionName(t *testing.T) {
	t.Parallel()

	assert.Nil(t, validateConditionName("is_valid"))

	finding := validateConditionName("has space")

	require.NotNil(t, finding)
	assert.Equal(t, InvalidName, finding.Metadata.Kind)
	assert.Equal(t, "condition 'has space' does not match naming rule: '^[^\\*\\s]{1,50}$'.", finding.Message)
	assert.Equal(t, "has space", finding.Metadata.Condition)
}

func TestValidateNames(t *testing.T) {
	t.Parallel()

	t.Run("positions resolve to the declaring lines", func(t *testing.T) {
		t.Parallel()

		// The parser would reject these names; build the model directly, as the
		// phase sees it.
		dsl := "model\n  schema 1.1\ntype self\n  relations\n    define this: [self]"
		model := modelWithRelations(t, "self", "this")

		findings := ExtractAllAs[*Finding](validateNames(model, newSource(dsl)))

		require.Len(t, findings, 2)

		assert.Equal(t, ReservedTypeKeywords, findings[0].Metadata.Kind)
		assert.Equal(t, &Range{Start: 2, End: 2}, findings[0].Line)
		assert.Equal(t, &Range{Start: 5, End: 9}, findings[0].Column)

		assert.Equal(t, ReservedRelationKeywords, findings[1].Metadata.Kind)
		assert.Equal(t, &Range{Start: 4, End: 4}, findings[1].Line)
		assert.Equal(t, &Range{Start: 11, End: 15}, findings[1].Column)
	})

	t.Run("no source text means no positions", func(t *testing.T) {
		t.Parallel()

		findings := ExtractAllAs[*Finding](validateNames(modelWithRelations(t, "self", "viewer"), source{}))

		require.Len(t, findings, 1)
		assert.Nil(t, findings[0].Line)
		assert.Nil(t, findings[0].Column)
	})
}
