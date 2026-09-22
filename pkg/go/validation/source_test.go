package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSourceTypeLine(t *testing.T) {
	t.Parallel()

	src := newSource("model\n  schema 1.1\ntype user\ntype document\n  relations\n    define viewer: [user]")

	assert.Equal(t, 2, src.typeLine("user"))
	assert.Equal(t, 3, src.typeLine("document"))
	assert.Equal(t, -1, src.typeLine("missing"))
	assert.Equal(t, -1, source{}.typeLine("user"), "no source, no line")
}

func TestSourceRelationLine(t *testing.T) {
	t.Parallel()

	src := newSource(`model
  schema 1.1
type user
type folder
  relations
    define viewer: [user]
type document
  relations
    define viewer: [user]`)

	assert.Equal(t, 5, src.relationLine("viewer", -1), "negative from searches the whole source")
	assert.Equal(t, 8, src.relationLine("viewer", src.typeLine("document")),
		"anchoring to the type finds its own define")
	assert.Equal(t, -1, src.relationLine("missing", -1))
}

func TestSourceConditionLine(t *testing.T) {
	t.Parallel()

	src := newSource(`model
  schema 1.1
type user
condition less(x: int) {
  x < 5
}
condition less_than(x: int) {
  x < 10
}`)

	assert.Equal(t, 3, src.conditionLine("less"), "a name that prefixes another must not match its line")
	assert.Equal(t, 6, src.conditionLine("less_than"))
	assert.Equal(t, -1, src.conditionLine("missing"))
}

func TestSourceSchemaLine(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 1, newSource("model\n  schema 1.0\ntype user").schemaLine("1.0"))
	assert.Equal(t, 1, newSource("model\n  schema   1.0  # retired\ntype user").schemaLine("1.0"),
		"a trailing comment and repeated spaces still match")
	assert.Equal(t, -1, newSource("model\n  schema 1.1\ntype user").schemaLine("1.0"))
	assert.Equal(t, -1, source{}.schemaLine("1.0"))
}

func TestWordIndex(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		rawLine string
		symbol  string
		want    int
	}{
		{"word boundary skips substring", "    define type_a: [t]", "t", 20},
		{"plain match", "    define viewer: [user]", "viewer", 11},
		{"symbol with non-word characters", "    define viewer: [user:*]", "user:*", 20},
		{"missing symbol falls back to zero", "    define viewer: [user]", "absent", 0},
		{"empty symbol", "anything", "", 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.want, wordIndex(test.rawLine, test.symbol))
		})
	}
}

func TestFindingAt(t *testing.T) {
	t.Parallel()

	src := newSource("model\n  schema 1.1\ntype document\n  relations\n    define viewer: [user]")

	t.Run("stamps line and column", func(t *testing.T) {
		t.Parallel()

		finding := (&Finding{Metadata: Metadata{Symbol: "viewer"}}).at(src, 4)

		require.NotNil(t, finding.Line)
		require.NotNil(t, finding.Column)
		assert.Equal(t, Range{Start: 4, End: 4}, *finding.Line)
		assert.Equal(t, Range{Start: 11, End: 17}, *finding.Column, "column is half-open over the symbol")
	})

	t.Run("no source means no position", func(t *testing.T) {
		t.Parallel()

		finding := (&Finding{Metadata: Metadata{Symbol: "viewer"}}).at(source{}, 0)

		assert.Nil(t, finding.Line)
		assert.Nil(t, finding.Column)
	})

	t.Run("failed line search means no position", func(t *testing.T) {
		t.Parallel()

		finding := (&Finding{Metadata: Metadata{Symbol: "viewer"}}).at(src, -1)

		assert.Nil(t, finding.Line)
		assert.Nil(t, finding.Column)
	})

	t.Run("nil finding stays nil", func(t *testing.T) {
		t.Parallel()

		var finding *Finding
		assert.Nil(t, finding.at(src, 0))
	})
}

func TestFindingAtFromClause(t *testing.T) {
	t.Parallel()

	// `owner` appears both as the define target side and after `from`; the
	// from-clause stamp must mark the occurrence after `from`.
	src := newSource("    define viewer: owner from owner")

	finding := (&Finding{Metadata: Metadata{Symbol: "owner"}}).atFromClause(src, 0)

	require.NotNil(t, finding.Column)
	assert.Equal(t, 30, finding.Column.Start)
	assert.Equal(t, 35, finding.Column.End)
}

func TestFindingAtRestriction(t *testing.T) {
	t.Parallel()

	// `group` is both the relation name and the restricted type; the
	// restriction stamp must mark the occurrence after the colon.
	src := newSource("    define group: [group]")

	finding := (&Finding{Metadata: Metadata{Symbol: "group"}}).atRestriction(src, 0)

	require.NotNil(t, finding.Column)
	assert.Equal(t, 19, finding.Column.Start)
	assert.Equal(t, 24, finding.Column.End)
}
