package validation

import (
	"regexp"
	"strings"

	"github.com/openfga/language/pkg/go/utils"
)

// source is the DSL text findings are located in. The zero value is a model
// that arrived as JSON: there is no text, so every lookup reports absent and
// every stamp is a no-op, and findings carry no position.
//
// The line lookups match against lines folded by utils.NormalizeWhitespace, so
// `define\towner:` is found like `define owner:`. Columns are still resolved
// against the raw line, so folding never shifts a reported position.
type source struct {
	lines []string
}

func newSource(dsl string) source {
	return source{lines: strings.Split(dsl, "\n")}
}

// typeLine returns the line index a type is declared on, or -1 when the source
// does not declare it.
func (s source) typeLine(typeName string) int {
	return utils.GetTypeLineNumber(typeName, s.lines)
}

// relationLine returns the line index a relation is defined on, searching from
// the given line so the right `define` is found when several types declare a
// relation of the same name. A negative from searches the whole source.
func (s source) relationLine(relationName string, from int) int {
	from = min(max(from, 0), len(s.lines))

	line := utils.GetRelationLineNumber(relationName, s.lines[from:])
	if line < 0 {
		return -1
	}

	return from + line
}

// conditionLine returns the line index a condition is declared on, or -1.
func (s source) conditionLine(conditionName string) int {
	return utils.GetConditionLineNumber(conditionName, s.lines)
}

// schemaLine returns the line index the schema version is declared on, or -1.
//
// A trailing comment may follow the version, as in `schema 1.1 # note`. The `#`
// has to be preceded by whitespace, so one written against the version is part
// of the version and does not match here.
func (s source) schemaLine(schemaVersion string) int {
	if len(s.lines) == 0 {
		return -1
	}

	pattern := regexp.MustCompile(`^\s*schema\s+` + regexp.QuoteMeta(schemaVersion) + `(\s+#.*)?$`)

	for i, line := range s.lines {
		normalized := utils.NormalizeWhitespace(strings.TrimSpace(line))
		if pattern.MatchString(normalized) {
			return i
		}
	}

	return -1
}

// at stamps the position a finding points at: the given line, and the column
// its symbol sits at on that line. It does nothing for a nil finding or for a
// line the source does not have, which covers both a failed line search (-1)
// and a model with no source text. It returns f, so a raise site reads
// `unusedCondition(name).at(src, line)`.
func (f *Finding) at(src source, line int) *Finding {
	if f == nil || line < 0 || line >= len(src.lines) {
		return f
	}

	f.Line = &Range{Start: line, End: line}

	col := wordIndex(src.lines[line], f.Metadata.Symbol)
	f.Column = &Range{Start: col, End: col + len(f.Metadata.Symbol)}

	return f
}

// atFromClause is at, with the column searched after the `from` keyword so it
// marks the offending tupleset relation rather than an earlier occurrence of
// the same name on the line.
func (f *Finding) atFromClause(src source, line int) *Finding {
	if f == nil {
		return nil
	}

	f.at(src, line)

	if f.Column == nil {
		return f
	}

	rawLine := src.lines[line]
	if clause := strings.Index(rawLine, "from"); clause >= 0 {
		col := clause + len("from") + strings.Index(rawLine[clause+len("from"):], f.Metadata.Symbol)
		f.Column = &Range{Start: col, End: col + len(f.Metadata.Symbol)}
	}

	return f
}

// atRestriction is at, with the column searched on the value side of the `:` so
// it marks the type restriction rather than a relation key sharing its name.
func (f *Finding) atRestriction(src source, line int) *Finding {
	if f == nil {
		return nil
	}

	f.at(src, line)

	if f.Column == nil {
		return f
	}

	rawLine := src.lines[line]
	if colon := strings.Index(rawLine, ":"); colon >= 0 {
		col := colon + 1 + wordIndex(rawLine[colon+1:], f.Metadata.Symbol)
		f.Column = &Range{Start: col, End: col + len(f.Metadata.Symbol)}
	}

	return f
}

// wordIndex returns the index of symbol in rawLine, preferring a match on word
// boundaries, as pkg/js's `\bsymbol\b` lookup does, so `t` is found on its own
// rather than inside `type`. With no such match it falls back to the first
// occurrence anywhere, where pkg/js answers 0, and it returns 0 when symbol is
// absent, as pkg/js does. The boundary test is done in place to avoid compiling
// a regexp per finding.
func wordIndex(rawLine, symbol string) int {
	if symbol == "" {
		return 0
	}
	// Only attempt a word-boundary match when the symbol begins and ends with a
	// word character; symbols containing non-word characters (e.g. `user:*`)
	// can't match `\bsymbol\b` and fall through to the substring search.
	if isWordChar(symbol[0]) && isWordChar(symbol[len(symbol)-1]) {
		for off := 0; ; {
			idx := strings.Index(rawLine[off:], symbol)
			if idx < 0 {
				break
			}

			pos := off + idx
			beforeOK := pos == 0 || !isWordChar(rawLine[pos-1])
			afterPos := pos + len(symbol)
			afterOK := afterPos == len(rawLine) || !isWordChar(rawLine[afterPos])

			if beforeOK && afterOK {
				return pos
			}

			off = pos + 1
		}
	}

	if idx := strings.Index(rawLine, symbol); idx >= 0 {
		return idx
	}

	return 0
}

// isWordChar reports whether b is a regexp `\w` character ([0-9A-Za-z_]).
func isWordChar(b byte) bool {
	return b == '_' ||
		(b >= '0' && b <= '9') ||
		(b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z')
}
