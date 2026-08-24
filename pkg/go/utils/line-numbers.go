package utils

import (
	"slices"
	"strings"
)

// nameBytes marks every byte that can appear inside a declaration name, keyed by
// the bytes themselves so each entry is visibly correct. IDENTIFIER admits letters,
// digits, `_` and MINUS, and EXTENDED_IDENTIFIER adds SLASH and DOT
// (OpenFGALexer.g4), so `core.doc` is one name rather than `core` and a terminator.
var nameBytes = [256]bool{
	'a': true, 'b': true, 'c': true, 'd': true, 'e': true, 'f': true, 'g': true,
	'h': true, 'i': true, 'j': true, 'k': true, 'l': true, 'm': true, 'n': true,
	'o': true, 'p': true, 'q': true, 'r': true, 's': true, 't': true, 'u': true,
	'v': true, 'w': true, 'x': true, 'y': true, 'z': true,
	'A': true, 'B': true, 'C': true, 'D': true, 'E': true, 'F': true, 'G': true,
	'H': true, 'I': true, 'J': true, 'K': true, 'L': true, 'M': true, 'N': true,
	'O': true, 'P': true, 'Q': true, 'R': true, 'S': true, 'T': true, 'U': true,
	'V': true, 'W': true, 'X': true, 'Y': true, 'Z': true,
	'0': true, '1': true, '2': true, '3': true, '4': true,
	'5': true, '6': true, '7': true, '8': true, '9': true,
	'_': true, '-': true, '/': true, '.': true,
}

// IsNameByte reports whether b can appear inside a declaration name.
func IsNameByte(b byte) bool {
	return nameBytes[b]
}

// name ends there, or -1. Prefix always spans whole words (`type <name>`, `define
// <name>`, `condition <name>`), so requiring the next byte to end the name is what
// stops `document` matching a declaration of `documentation`.
func declarationIndex(lines []string, prefix string) int {
	return slices.IndexFunc(lines, func(line string) bool {
		trimmed := NormalizeWhitespace(strings.TrimSpace(line))
		if !strings.HasPrefix(trimmed, prefix) {
			return false
		}

		rest := trimmed[len(prefix):]

		return rest == "" || !IsNameByte(rest[0])
	})
}

// isInlineWhitespace reports whether b is a byte the lexer's WHITESPACE rule
// accepts between tokens on a line: space, tab or form feed (OpenFGALexer.g4).
func isInlineWhitespace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\f'
}

// NormalizeWhitespace collapses every run of inline whitespace into one space, so
// `define\towner:` and `define  owner:` are matched the same as `define owner:`.
// Only space, tab and form feed are folded — exactly the lexer's WHITESPACE
// alphabet; anything else (nbsp, vertical tab) fails to lex and can never reach a
// line lookup. Lines that are already normal are returned unchanged.
func NormalizeWhitespace(line string) string {
	for i := range len(line) {
		if b := line[i]; b == '\t' || b == '\f' || (b == ' ' && i > 0 && line[i-1] == ' ') {
			return foldInlineWhitespace(line)
		}
	}

	return line
}

func foldInlineWhitespace(line string) string {
	var sb strings.Builder

	sb.Grow(len(line))

	inRun := false

	for i := range len(line) {
		if isInlineWhitespace(line[i]) {
			inRun = true

			continue
		}

		if inRun {
			sb.WriteByte(' ')

			inRun = false
		}

		sb.WriteByte(line[i])
	}

	if inRun {
		sb.WriteByte(' ')
	}

	return sb.String()
}

// GetConditionLineNumber returns the index of the line declaring conditionName, or
// -1. `less` does not match a declaration of `less_than`.
func GetConditionLineNumber(conditionName string, lines []string) int {
	return declarationIndex(lines, "condition "+conditionName)
}

// GetTypeLineNumber returns the index of the line declaring typeName, or -1.
// `document` does not match a declaration of `documentation`.
func GetTypeLineNumber(typeName string, lines []string) int {
	return declarationIndex(lines, "type "+typeName)
}

// GetExtendedTypeLineNumber returns the index of the line extending typeName, or -1.
func GetExtendedTypeLineNumber(typeName string, lines []string) int {
	return declarationIndex(lines, "extend type "+typeName)
}

// GetRelationLineNumber returns the index of the line defining relation, or -1.
// `owner` does not match a definition of `owner_group`.
func GetRelationLineNumber(relation string, lines []string) int {
	return declarationIndex(lines, "define "+relation)
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
