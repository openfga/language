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

	col, width := symbolColumn(src.lines[line], f.Metadata.Symbol)
	f.Column = &Range{Start: col, End: col + width}

	return f
}

// atFromClause is at, with the column moved to the symbol as the operand of a
// `from` clause, so it marks the offending tupleset relation rather than an
// earlier occurrence of the same name on the line.
func (f *Finding) atFromClause(src source, line int) *Finding {
	if f == nil {
		return nil
	}

	f.at(src, line)

	if f.Column == nil {
		return f
	}

	if col := fromClauseIndex(src.lines[line], f.Metadata.Symbol); col >= 0 {
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

// symbolColumn returns where symbol starts on rawLine and how many characters
// it spans there. A symbol of several words, such as `owner from parent`, is
// written with single spaces but may be separated by any run of spaces, tabs or
// form feeds in the source, so its words are matched across such runs and it
// spans them as written. A symbol not found that way is located by wordIndex.
func symbolColumn(rawLine, symbol string) (start, width int) {
	if words := strings.Fields(symbol); len(words) > 1 {
		if start, end, ok := wordsSpan(rawLine, words); ok {
			return start, end - start
		}
	}

	return wordIndex(rawLine, symbol), len(symbol)
}

// wordsSpan finds words on rawLine in order, each separated from the next by
// at least one space, tab or form feed, and returns where the first starts and
// the last ends. As in wordIndex, the match has to start and end on a word
// boundary where the first word starts, or the last word ends, with a word
// character.
func wordsSpan(rawLine string, words []string) (start, end int, ok bool) {
	first, last := words[0], words[len(words)-1]

	for off := 0; ; {
		idx := strings.Index(rawLine[off:], first)
		if idx < 0 {
			return 0, 0, false
		}

		start = off + idx
		off = start + 1

		if isWordChar(first[0]) && start > 0 && isWordChar(rawLine[start-1]) {
			continue
		}

		end, ok = start+len(first), true

		for _, word := range words[1:] {
			gap := end
			for gap < len(rawLine) && isInlineSpace(rawLine[gap]) {
				gap++
			}

			if gap == end || !strings.HasPrefix(rawLine[gap:], word) {
				ok = false

				break
			}

			end = gap + len(word)
		}

		if !ok || (isWordChar(last[len(last)-1]) && end < len(rawLine) && isWordChar(rawLine[end])) {
			continue
		}

		return start, end, true
	}
}

// fromClauseIndex returns where symbol sits as the operand of a `from` clause
// on rawLine, or -1 when no clause has it. The `from` keyword stands on its
// own, after the start of the line or a space, tab or form feed, so a `from`
// inside an identifier such as `from_parent` is not taken for it, and the
// operand is the whole identifier after it.
func fromClauseIndex(rawLine, symbol string) int {
	const keyword = "from"

	for off := 0; ; {
		idx := strings.Index(rawLine[off:], keyword)
		if idx < 0 {
			return -1
		}

		pos := off + idx
		off = pos + 1

		if pos > 0 && !isInlineSpace(rawLine[pos-1]) {
			continue
		}

		operand := pos + len(keyword)
		for operand < len(rawLine) && isInlineSpace(rawLine[operand]) {
			operand++
		}

		if operand == pos+len(keyword) || !strings.HasPrefix(rawLine[operand:], symbol) {
			continue
		}

		if after := operand + len(symbol); after < len(rawLine) && isIdentifierChar(rawLine[after]) {
			continue
		}

		return operand
	}
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

// isIdentifierChar reports whether b can continue a DSL identifier: a word
// character or a `-`.
func isIdentifierChar(b byte) bool {
	return isWordChar(b) || b == '-'
}

// isInlineSpace reports whether b separates tokens within a DSL line: a space,
// tab or form feed, as the lexer's WHITESPACE allows.
func isInlineSpace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\f'
}
