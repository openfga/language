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

// severityNames maps each severity to its wire name. A severity missing from here
// fails to marshal, so a constant added without a name is caught.
var severityNames = map[Severity]string{
	SeverityError:    "error",
	SeverityWarning:  "warning",
	SeverityAdvisory: "advisory",
}

// severityValues is the reverse of severityNames, built from it so the two
// cannot disagree.
var severityValues = func() map[string]Severity {
	values := make(map[string]Severity, len(severityNames))
	for severity, name := range severityNames {
		values[name] = severity
	}

	return values
}()

// String returns the wire name, or a diagnostic form for a value with none.
func (s Severity) String() string {
	if name, ok := severityNames[s]; ok {
		return name
	}

	if s == SeverityUnspecified {
		return ""
	}

	return fmt.Sprintf("Severity(%d)", int(s))
}

// IsValid reports whether s is a declared severity with a wire name.
func (s Severity) IsValid() bool {
	_, ok := severityNames[s]

	return ok
}

// Blocks reports whether a finding of this severity makes validation fail.
//
// Only the severities declared as non-blocking answer false, so an unset or
// undeclared value blocks: a severity that cannot be recognised must not let an
// invalid model pass.
func (s Severity) Blocks() bool {
	return s != SeverityWarning && s != SeverityAdvisory
}

// MarshalText emits the wire name, so the JSON carries "warning".
func (s Severity) MarshalText() ([]byte, error) {
	name, ok := severityNames[s]
	if !ok {
		return nil, fmt.Errorf("%w: %d", ErrUnknownSeverity, int(s))
	}

	return []byte(name), nil
}

// UnmarshalText resolves a wire name back to its severity, rejecting any name
// this package does not declare.
func (s *Severity) UnmarshalText(text []byte) error {
	severity, ok := severityValues[string(text)]
	if !ok {
		return fmt.Errorf("%w: %q", ErrUnknownSeverity, text)
	}

	*s = severity

	return nil
}
