package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateTypeName(t *testing.T) {
	t.Parallel()

	t.Run("valid names yield nothing", func(t *testing.T) {
		t.Parallel()

		for _, name := range []string{"document", "user_group", "group2"} {
			assert.NoError(t, validateTypeName(name), name)
		}
	})

	t.Run("reserved keywords", func(t *testing.T) {
		t.Parallel()

		for _, reserved := range []string{"self", "this"} {
			var finding *Finding
			require.ErrorAs(t, validateTypeName(reserved), &finding)
			assert.Equal(t, ReservedTypeKeywords, finding.Metadata.Kind)
			assert.Equal(t, reserved, finding.Metadata.Symbol)
			assert.Equal(t, reserved, finding.Metadata.Type)
		}
	})

	t.Run("names breaking the rule are reported as invalid-name", func(t *testing.T) {
		t.Parallel()

		for _, name := range []string{"doc:ument", "document name", "document#tag", "doc@ument", "doc*ument", ""} {
			var finding *Finding
			require.ErrorAs(t, validateTypeName(name), &finding, name)
			assert.Equal(t, InvalidName, finding.Metadata.Kind, name)
		}
	})
}

func TestValidateRelationName(t *testing.T) {
	t.Parallel()

	t.Run("valid names yield nothing", func(t *testing.T) {
		t.Parallel()

		for _, name := range []string{"viewer", "can_view", "viewer2"} {
			assert.NoError(t, validateRelationName(name, "document"), name)
		}
	})

	t.Run("reserved keywords", func(t *testing.T) {
		t.Parallel()

		for _, reserved := range []string{"self", "this"} {
			var finding *Finding
			require.ErrorAs(t, validateRelationName(reserved, "document"), &finding)
			assert.Equal(t, ReservedRelationKeywords, finding.Metadata.Kind)
			assert.Equal(t, "document", finding.Metadata.Type)
			assert.Equal(t, reserved, finding.Metadata.Relation)
		}
	})

	t.Run("names breaking the rule are reported as invalid-name", func(t *testing.T) {
		t.Parallel()

		for _, name := range []string{"view#er", "viewer:invalid", "can view", "view@er", "view*er", ""} {
			var finding *Finding
			require.ErrorAs(t, validateRelationName(name, "document"), &finding, name)
			assert.Equal(t, InvalidName, finding.Metadata.Kind, name)
		}
	})
}

func TestValidateConditionName(t *testing.T) {
	t.Parallel()

	for _, name := range []string{"is_valid", "is_owner", "condition123"} {
		require.NoError(t, validateConditionName(name), name)
	}

	for _, name := range []string{"condition*", ""} {
		var invalid *Finding
		require.ErrorAs(t, validateConditionName(name), &invalid, name)
		assert.Equal(t, InvalidName, invalid.Metadata.Kind, name)
	}

	var finding *Finding
	require.ErrorAs(t, validateConditionName("has space"), &finding)
	assert.Equal(t, InvalidName, finding.Metadata.Kind)
	assert.Equal(t, "condition 'has space' does not match naming rule: '^[^\\*\\s]{1,50}$'.", finding.Message)
	assert.Equal(t, "has space", finding.Metadata.Condition)
}

func TestValidateNames(t *testing.T) {
	t.Parallel()

	t.Run("positions resolve to the declaring lines", func(t *testing.T) {
		t.Parallel()

		dsl := "model\n  schema 1.1\ntype self\n  relations\n    define this: [self]"

		findings := ExtractAllAs[*Finding](validateNames(mustParse(t, dsl), newSource(dsl)))

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

		dsl := "model\n  schema 1.1\ntype self\n  relations\n    define viewer: [self]"

		findings := ExtractAllAs[*Finding](validateNames(mustParse(t, dsl), source{}))

		require.Len(t, findings, 1)
		assert.Nil(t, findings[0].Line)
		assert.Nil(t, findings[0].Column)
	})
}
