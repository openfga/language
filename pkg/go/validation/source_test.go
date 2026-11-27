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
	assert.Equal(t, -1, source{}.relationLine("viewer", -1), "no source, no line")
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

	assert.Equal(t, 3, src.conditionLine("less"))
	assert.Equal(t, 6, src.conditionLine("less_than"))
	assert.Equal(t, -1, src.conditionLine("missing"))
	assert.Equal(t, -1, source{}.conditionLine("less"), "no source, no line")
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
		{"first boundary match wins", "a or a", "a", 0},
		{"falls back to a substring with no boundary match", "type usergroup", "user", 5},
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

		assert.Equal(t, &Range{Start: 4, End: 4}, finding.Line)
		assert.Equal(t, &Range{Start: 11, End: 17}, finding.Column, "column is half-open over the symbol")
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

	tests := []struct {
		name    string
		rawLine string
		symbol  string
		want    *Range
	}{
		{
			// `owner` is both the computed relation and the tupleset after `from`.
			name:    "marks the operand after from",
			rawLine: "    define viewer: owner from owner",
			symbol:  "owner",
			want:    &Range{Start: 30, End: 35},
		},
		{
			name:    "skips a from inside the relation name",
			rawLine: "    define from_parent: member from parent",
			symbol:  "parent",
			want:    &Range{Start: 36, End: 42},
		},
		{
			name:    "skips a from clause whose operand only starts with the symbol",
			rawLine: "    define viewer: member from organization or owner from org",
			symbol:  "org",
			want:    &Range{Start: 58, End: 61},
		},
		{
			name:    "follows tabs around from",
			rawLine: "    define viewer: owner\tfrom\tparent",
			symbol:  "parent",
			want:    &Range{Start: 30, End: 36},
		},
		{
			name:    "keeps the symbol's column when no from clause has it",
			rawLine: "    define viewer: [parent]",
			symbol:  "parent",
			want:    &Range{Start: 20, End: 26},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			finding := (&Finding{Metadata: Metadata{Symbol: test.symbol}}).atFromClause(newSource(test.rawLine), 0)

			assert.Equal(t, test.want, finding.Column)
		})
	}
}

func TestSymbolColumn(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		rawLine   string
		symbol    string
		wantStart int
		wantWidth int
	}{
		{"one word", "    define viewer: [user]", "viewer", 11, 6},
		{"words separated by single spaces", "    define viewer: owner from parent", "owner from parent", 19, 17},
		{"words separated by tabs", "    define viewer: owner\tfrom\tparent", "owner from parent", 19, 17},
		{"words separated by several spaces span them", "    define viewer: owner  from   parent", "owner from parent", 19, 20},
		{"a type restriction with a condition", "    define viewer: [user:*  with cond]", "user:* with cond", 20, 17},
		{"the first word must start a word", "    define viewer: co_owner from parent or owner from parent", "owner from parent", 43, 17},
		{"the last word must end a word", "    define viewer: owner from parents or owner from parent", "owner from parent", 41, 17},
		{"words not found fall back to wordIndex", "    define viewer: owner or parent", "owner from parent", 0, 17},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			start, width := symbolColumn(test.rawLine, test.symbol)

			assert.Equal(t, test.wantStart, start)
			assert.Equal(t, test.wantWidth, width)
		})
	}
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
