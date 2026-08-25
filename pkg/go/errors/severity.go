package errors

import "fmt"

// Severity states whether a finding makes a model invalid, or only reports
// something about a model that stays valid.
//
// It serialises as its wire name, so the names below are API and the numbers are
// not.
type Severity int

// SeverityUnspecified is the zero value, so a finding that never set a severity
// cannot pass for one that did. It has no wire name, omitempty drops it, and
// Blocks treats it as blocking.
const SeverityUnspecified Severity = 0

const (
	// SeverityError means the model is invalid. Validation fails.
	SeverityError Severity = iota + 1

	// SeverityWarning means the model is valid today but relies on something a
	// future version may not accept. Validation does not fail.
	SeverityWarning

	// SeverityAdvisory means the model is valid, but a request against it may not
	// behave the way the author expects, depending on the tuples written and the
	// checks issued. Validation does not fail.
	SeverityAdvisory
)

// severityFromName maps a wire name back to its severity.
func severityFromName(name string) (Severity, bool) {
	switch name {
	case "error":
		return SeverityError, true
	case "warning":
		return SeverityWarning, true
	case "advisory":
		return SeverityAdvisory, true
	default:
		return SeverityUnspecified, false
	}
}

// String returns the wire name of a declared severity, an empty string for the
// zero value, and a diagnostic form for any other number.
func (s Severity) String() string {
	switch s {
	case SeverityError:
		return "error"
	case SeverityWarning:
		return "warning"
	case SeverityAdvisory:
		return "advisory"
	case SeverityUnspecified:
		return ""
	default:
		return fmt.Sprintf("Severity(%d)", int(s))
	}
}

// IsValid reports whether s is a declared severity.
func (s Severity) IsValid() bool {
	switch s {
	case SeverityError, SeverityWarning, SeverityAdvisory:
		return true
	default:
		return false
	}
}

// Blocks reports whether a finding of this severity makes validation fail.
//
// Only the severities declared as non-blocking answer false, so an unset or
// undeclared value blocks: a severity that cannot be recognised must not let an
// invalid model pass.
func (s Severity) Blocks() bool {
	return s != SeverityWarning && s != SeverityAdvisory
}

// MarshalText emits the wire name, so the JSON carries "warning". An undeclared
// value must fail to marshal rather than ship String's diagnostic form.
func (s Severity) MarshalText() ([]byte, error) {
	if !s.IsValid() {
		return nil, fmt.Errorf("%w: %d", ErrUnknownSeverity, int(s))
	}

	return []byte(s.String()), nil
}

// UnmarshalText resolves a wire name back to its severity, rejecting any name
// this package does not declare.
func (s *Severity) UnmarshalText(text []byte) error {
	severity, ok := severityFromName(string(text))
	if !ok {
		return fmt.Errorf("%w: %q", ErrUnknownSeverity, text)
	}

	*s = severity

	return nil
}
