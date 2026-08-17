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

// declarationIndex returns the index of the first line beginning with prefix whose
// name ends there, or -1. prefix always spans whole words (`type <name>`, `define
// <name>`, `condition <name>`), so requiring the next byte to end the name is what
// stops `document` matching a declaration of `documentation`.
func declarationIndex(lines []string, prefix string, normalize func(string) string) int {
	return slices.IndexFunc(lines, func(line string) bool {
		trimmed := normalize(strings.TrimSpace(line))
		if !strings.HasPrefix(trimmed, prefix) {
			return false
		}

		rest := trimmed[len(prefix):]

		return rest == "" || !IsNameByte(rest[0])
	})
}

// keepSpaces leaves a line as-is. Only the relation helper collapses repeated
// spaces, matching the reference, where ` {2,}` normalization is applied in
// getRelationLineNumber alone.
func keepSpaces(line string) string { return line }

// normalizeSpaces collapses runs of spaces into one, mirroring the reference's
// ` {2,}` normalization, so `define  owner:` is matched the same as `define owner:`.
func normalizeSpaces(line string) string {
	for strings.Contains(line, "  ") {
		line = strings.ReplaceAll(line, "  ", " ")
	}

	return line
}

// GetConditionLineNumber returns the index of the line declaring conditionName, or
// -1. `less` does not match a declaration of `less_than`.
func GetConditionLineNumber(conditionName string, lines []string) int {
	return declarationIndex(lines, "condition "+conditionName, keepSpaces)
}

// GetTypeLineNumber returns the index of the line declaring typeName, or -1.
// `document` does not match a declaration of `documentation`.
func GetTypeLineNumber(typeName string, lines []string) int {
	return declarationIndex(lines, "type "+typeName, keepSpaces)
}

// GetExtendedTypeLineNumber returns the index of the line extending typeName, or -1.
func GetExtendedTypeLineNumber(typeName string, lines []string) int {
	return declarationIndex(lines, "extend type "+typeName, keepSpaces)
}

// GetRelationLineNumber returns the index of the line defining relation, or -1.
// `owner` does not match a definition of `owner_group`.
func GetRelationLineNumber(relation string, lines []string) int {
	return declarationIndex(lines, "define "+relation, normalizeSpaces)
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
