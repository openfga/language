package validation

import (
	"regexp"
	"strings"
)

// source is the DSL text findings are located in. The zero value is a model
// that arrived as JSON: there is no text, so every lookup reports absent and
// every stamp is a no-op, and findings carry no position.
type source struct {
	lines []string
}

func newSource(dsl string) source {
	return source{lines: strings.Split(dsl, "\n")}
}

// foldInline collapses every run of inline whitespace into one space, so
// `define\towner:` is matched like `define owner:`. Only space, tab and form
// feed fold — exactly the lexer's WHITESPACE alphabet (OpenFGALexer.g4);
// anything else fails to lex and can never reach a line lookup. Lines already
// normal are returned unchanged. Matching only: columns are still resolved
// against the raw line, so folding never shifts a reported position.
//
// TODO(SoulPancake): replace with utils.NormalizeWhitespace once #652 merges,
// so the repo has one whitespace-folder.
func foldInline(line string) string {
	folded := false

	for i := 0; i < len(line); i++ {
		if b := line[i]; b == '\t' || b == '\f' || (b == ' ' && i > 0 && line[i-1] == ' ') {
			folded = true

			break
		}
	}

	if !folded {
		return line
	}

	var b strings.Builder
	b.Grow(len(line))

	inRun := false

	for i := 0; i < len(line); i++ {
		if c := line[i]; c == ' ' || c == '\t' || c == '\f' {
			if !inRun {
				b.WriteByte(' ')
			}

			inRun = true
		} else {
			b.WriteByte(c)
			inRun = false
		}
	}

	return b.String()
}

// typeLine returns the line index a type is declared on, or -1 when the source
// does not declare it.
func (s source) typeLine(typeName string) int {
	for i, line := range s.lines {
		trimmed := foldInline(strings.TrimSpace(line))
		if !strings.HasPrefix(trimmed, "type ") {
			continue
		}

		if fields := strings.Fields(trimmed); len(fields) >= 2 && fields[1] == typeName {
			return i
		}
	}

	return -1
}

// relationLine returns the line index a relation is defined on, searching from
// the given line so the right `define` is found when several types declare a
// relation of the same name. A negative from searches the whole source.
func (s source) relationLine(relationName string, from int) int {
	if from < 0 {
		from = 0
	}

	for i := from; i < len(s.lines); i++ {
		trimmed := foldInline(strings.TrimSpace(s.lines[i]))
		if !strings.HasPrefix(trimmed, "define ") {
			continue
		}

		definePart := strings.TrimPrefix(trimmed, "define ")

		colon := strings.Index(definePart, ":")
		if colon > 0 && strings.TrimSpace(definePart[:colon]) == relationName {
			return i
		}
	}

	return -1
}

// conditionLine returns the line index a condition is declared on, or -1. The
// parameter list's `(` must follow the name, so a condition whose name is a
// prefix of another (e.g. `less` vs `less_than`) cannot match the wrong line.
func (s source) conditionLine(conditionName string) int {
	prefix := "condition " + conditionName

	for i, line := range s.lines {
		trimmed := foldInline(strings.TrimSpace(line))
		if !strings.HasPrefix(trimmed, prefix) {
			continue
		}

		if strings.HasPrefix(strings.TrimLeft(trimmed[len(prefix):], " \t"), "(") {
			return i
		}
	}

	return -1
}

// multiSpaceRegex collapses runs of whitespace when normalizing a DSL line for
// schema-version matching; foldInline is not used here because the pattern's
// own \s+ already admits any single separator and this predates it. Hoisted so
// it is compiled once, not per line.
var multiSpaceRegex = regexp.MustCompile(`\s{2,}`)

// schemaLine returns the line index the schema version is declared on, or -1.
//
// A trailing comment may follow the version, as in `schema 1.1 # note`. The `#`
// has to be preceded by whitespace, so one written against the version is part
// of the version and does not match here.
func (s source) schemaLine(schemaVersion string) int {
	if len(s.lines) == 0 {
		return -1
	}

	pattern := regexp.MustCompile(`^\s*schema\s+` + regexp.QuoteMeta(schemaVersion) + `(\s+#.*)?\s*$`)

	for i, line := range s.lines {
		normalized := multiSpaceRegex.ReplaceAllString(strings.TrimSpace(line), " ")
		if pattern.MatchString(normalized) {
			return i
		}
	}

	return -1
}

// at stamps the position a finding points at: the given line, and the column
// its symbol sits at on that line. No other code computes positions.
// A no-op for a nil finding and for a line the source does not have, which
// covers both a failed line search (-1) and a model with no source text.
// Chainable, so a raise site reads `invalidType(name).at(src, line).in(file, module)`.
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

// wordIndex returns the index of symbol in rawLine matched on word boundaries,
// mirroring the reference's `\bsymbol\b` lookup. This avoids matching a symbol
// as a substring of another word (e.g. finding `t` inside `type`). Returns 0
// when the symbol is not found, matching the reference's fallback.
//
// The boundary check is done directly rather than via a per-call compiled
// regexp: `\b` only requires that the characters flanking the match are not
// word characters, which is cheap to test in place and avoids recompiling a
// pattern for every finding.
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
