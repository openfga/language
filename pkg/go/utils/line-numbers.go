package utils

import (
	"slices"
	"strings"
)

// intraLineWhitespace is the set of characters the grammar's WHITESPACE token
// admits within a line: `WHITESPACE: ( '\t' | ' ' | '\u000C')+;` (OpenFGALexer.g4).
// `\f` (U+000C) belongs here because the lexer prefers WHITESPACE over NEWLINE for
// a bare `\f`, so `define owner\f: [user]` is a single valid line. Trimming a
// narrower set would make such a declaration unfindable and collapse its reported
// location to 0:0.
const intraLineWhitespace = " \t\f"

// declarationIndex returns the index of the first line that, once trimmed, begins
// with prefix and whose remainder satisfies rest. Requiring the caller to validate
// the remainder is what keeps a declaration whose name is a prefix of another
// (e.g. `type document` vs `type documentation`) from matching the wrong line.
func declarationIndex(lines []string, prefix string, rest func(string) bool) int {
	return slices.IndexFunc(lines, func(line string) bool {
		trimmed := strings.TrimSpace(line)
		if !strings.HasPrefix(trimmed, prefix) {
			return false
		}

		return rest(trimmed[len(prefix):])
	})
}

// endOrComment reports whether the remainder of a declaration line is empty or is
// only a trailing comment, mirroring the reference's `^(\s+#.*)?$`. The `#` must be
// preceded by whitespace so a `#` glued to the name isn't treated as a comment —
// `type doc#x` is not a declaration of `doc`.
func endOrComment(rest string) bool {
	if rest == "" {
		return true
	}

	trimmed := strings.TrimLeft(rest, intraLineWhitespace)
	if trimmed == rest {
		return false
	}

	return strings.HasPrefix(trimmed, "#")
}

// startsWith reports whether rest begins with sep, ignoring leading whitespace.
func startsWith(rest, sep string) bool {
	return strings.HasPrefix(strings.TrimLeft(rest, intraLineWhitespace), sep)
}

// normalizeSpaces collapses runs of spaces into one, mirroring the reference's
// ` {2,}` normalization, so `define  owner:` is matched the same as `define owner:`.
func normalizeSpaces(line string) string {
	for strings.Contains(line, "  ") {
		line = strings.ReplaceAll(line, "  ", " ")
	}

	return line
}

// GetConditionLineNumber returns the index of the line declaring conditionName, or
// -1. The parameter list's `(` must follow the name, so `less` does not match a
// declaration of `less_than`.
func GetConditionLineNumber(conditionName string, lines []string) int {
	return declarationIndex(lines, "condition "+conditionName, func(rest string) bool {
		return startsWith(rest, "(")
	})
}

// GetTypeLineNumber returns the index of the line declaring typeName, or -1. Only a
// trailing comment may follow the name, so `document` does not match a declaration
// of `documentation`.
func GetTypeLineNumber(typeName string, lines []string) int {
	return declarationIndex(lines, "type "+typeName, endOrComment)
}

// GetExtendedTypeLineNumber returns the index of the line extending typeName, or -1.
// The name must be followed only by a trailing comment, as in GetTypeLineNumber.
func GetExtendedTypeLineNumber(typeName string, lines []string) int {
	return declarationIndex(lines, "extend type "+typeName, endOrComment)
}

// GetRelationLineNumber returns the index of the line defining relation, or -1. The
// `:` must follow the name, so `owner` does not match a definition of `owner_group`.
func GetRelationLineNumber(relation string, lines []string) int {
	prefix := "define " + relation

	return slices.IndexFunc(lines, func(line string) bool {
		normalized := normalizeSpaces(strings.TrimSpace(line))
		if !strings.HasPrefix(normalized, prefix) {
			return false
		}

		return startsWith(normalized[len(prefix):], ":")
	})
}

type StartEnd struct {
	Start int
	End   int
}

func ConstructLineAndColumnData(lines []string, lineIndex int, symbol string) (StartEnd, StartEnd) {
	if len(lines) == 0 || lineIndex == -1 {
		return StartEnd{
				Start: 0,
				End:   0,
			},
			StartEnd{
				Start: 0,
				End:   0,
			}
	}

	rawLine := lines[lineIndex]

	wordIdx := strings.Index(rawLine, symbol)

	if wordIdx == -1 {
		wordIdx = 0
	}

	return StartEnd{
			Start: lineIndex,
			End:   lineIndex,
		},
		StartEnd{
			Start: wordIdx,
			End:   wordIdx + len(symbol),
		}
}
