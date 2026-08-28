package validation

import (
	"strings"

	fgaerrors "github.com/openfga/language/pkg/go/errors"
)

// wordIndex returns the index of symbol in rawLine matched on word boundaries,
// mirroring the reference's `\bsymbol\b` lookup. This avoids matching a symbol
// as a substring of another word (e.g. finding `t` inside `type`). Returns 0
// when the symbol is not found, matching the reference's fallback.
//
// The boundary check is done directly rather than via a per-call compiled
// regexp: `\b` only requires that the characters flanking the match are not word
// characters, which is cheap to test in place and avoids recompiling a pattern
// for every error.
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

// scope is what a raise site knows that a finding's code cannot work out on its own:
// which part of the model is at fault, and the enclosing type for the metadata.
type scope struct {
	// part names the part of the model at fault. The raise site builds it, because
	// the code alone does not say which part: duplicated-error is raised about a type
	// from one place and a relation from another, and invalid-name about all three.
	// The sentinel is filled in from the code's table entry, so at this point it
	// wraps nothing.
	part fgaerrors.ModelError

	// offendingType is the enclosing type a finding about another type was written
	// in, matching JS's wire field of the same name. Metadata only: none of the
	// scope types has a slot for it.
	offendingType string
}

// newValidationError builds the finding a raise site describes: the code decides the
// severity and the sentinel it wraps, the scope decides which part of the model it
// names, and both the category and the metadata are read back off that. The line and
// column arguments are the already-resolved position, nil when the raise site gave none.
func newValidationError(message string, errorType ValidationErrorType, symbol string,
	line, column *Range, errorScope scope, meta *Meta) *ValidationError {
	part := errorScope.part
	if part == nil {
		// A raise site that named nothing. Treat it as being about the model as a
		// whole, which is what a code with no scope means.
		part = &fgaerrors.ErrModel{}
	}

	entry := lookupErrorInfo(errorType)
	partScope := part.Scope()

	metadata := &ErrorMetadata{
		Symbol:        symbol,
		ErrorType:     errorType,
		OffendingType: errorScope.offendingType,
		Type:          partScope.ObjectType,
		Relation:      partScope.Relation,
		Condition:     partScope.Condition,
	}

	if meta != nil {
		// Module goes in the metadata, file on the error itself, matching the
		// JS implementation.
		metadata.Module = meta.Module
	}

	validationErr := &ValidationError{
		Message:  message,
		Severity: entry.Severity,
		Category: part.Kind(),
		Line:     line,
		Column:   column,
		Metadata: metadata,

		// A code missing from the table has no sentinel, so there is nothing for
		// errors.Is to match and this is nil. The category and metadata above still
		// report what the raise site named.
		Cause: fgaerrors.WithSentinel(part, entry.Cause),
	}

	if meta != nil {
		validationErr.File = meta.File
	}

	return validationErr
}

// resolvePosition resolves the line and column a finding points at, both nil when the
// raise site gave no line or the line is outside the source.
func resolvePosition(lines []string, symbol string, lineIndex *int,
	customResolver ErrorCustomResolver) (line, column *Range) {
	if lineIndex == nil || *lineIndex < 0 || *lineIndex >= len(lines) {
		return nil, nil
	}

	line = &Range{Start: *lineIndex, End: *lineIndex}

	// Find symbol position in line for column calculation, matching on word
	// boundaries as the reference does.
	rawLine := lines[*lineIndex]
	symbolPos := wordIndex(rawLine, symbol)

	if customResolver != nil {
		symbolPos = customResolver(symbolPos, rawLine, symbol)
	}

	if symbolPos >= 0 {
		column = &Range{
			Start: symbolPos,
			End:   symbolPos + len(symbol),
		}
	}

	return line, column
}
