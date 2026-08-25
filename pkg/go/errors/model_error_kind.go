package errors

import "fmt"

// ModelErrorKind is the part of a model a finding is attached to: which kind of
// thing is wrong, as against the specific problem (the error code) and the
// identity of the thing (the finding's metadata).
//
// It serialises as its wire name, so the names below are API and the numbers are
// not.
type ModelErrorKind int

// ModelErrorKindUnspecified is the zero value, reserved so that a finding which
// never set a category cannot pass for one that did. It has no wire name and
// omitempty drops it.
const ModelErrorKindUnspecified ModelErrorKind = 0

const (
	// ErrorKindObjectType is a finding about an object type as a whole.
	ErrorKindObjectType ModelErrorKind = iota + 1

	// ErrorKindRelation is a finding about a relation on an object type.
	ErrorKindRelation

	// ErrorKindRelationCondition is a finding about a condition applied to a
	// relation on an object type.
	ErrorKindRelationCondition

	// ErrorKindCondition is a finding about a condition definition itself,
	// independent of where it is applied.
	ErrorKindCondition

	// ErrorKindInvalidModel is a finding about the model as a whole, which
	// cannot be attributed to a single type, relation or condition.
	ErrorKindInvalidModel
)

// modelErrorKindFromName maps a wire name back to its category.
func modelErrorKindFromName(name string) (ModelErrorKind, bool) {
	switch name {
	case "object-type":
		return ErrorKindObjectType, true
	case "relation":
		return ErrorKindRelation, true
	case "relation-condition":
		return ErrorKindRelationCondition, true
	case "condition":
		return ErrorKindCondition, true
	case "invalid-model":
		return ErrorKindInvalidModel, true
	default:
		return ModelErrorKindUnspecified, false
	}
}

// String returns the wire name of a declared category, an empty string for the
// zero value, and a diagnostic form for any other number.
func (m ModelErrorKind) String() string {
	switch m {
	case ErrorKindObjectType:
		return "object-type"
	case ErrorKindRelation:
		return "relation"
	case ErrorKindRelationCondition:
		return "relation-condition"
	case ErrorKindCondition:
		return "condition"
	case ErrorKindInvalidModel:
		return "invalid-model"
	case ModelErrorKindUnspecified:
		return ""
	default:
		return fmt.Sprintf("ModelErrorKind(%d)", int(m))
	}
}

// IsValid reports whether m is a declared category.
func (m ModelErrorKind) IsValid() bool {
	switch m {
	case ErrorKindObjectType,
		ErrorKindRelation,
		ErrorKindRelationCondition,
		ErrorKindCondition,
		ErrorKindInvalidModel:
		return true
	default:
		return false
	}
}

// MarshalText emits the wire name, so the JSON carries "object-type". An
// undeclared value must fail to marshal rather than ship String's diagnostic form.
func (m ModelErrorKind) MarshalText() ([]byte, error) {
	if !m.IsValid() {
		return nil, fmt.Errorf("%w: %d", ErrUnknownModelErrorKind, int(m))
	}

	return []byte(m.String()), nil
}

// UnmarshalText resolves a wire name back to its category, rejecting any name
// this package does not declare.
func (m *ModelErrorKind) UnmarshalText(text []byte) error {
	errorType, ok := modelErrorKindFromName(string(text))
	if !ok {
		return fmt.Errorf("%w: %q", ErrUnknownModelErrorKind, text)
	}

	*m = errorType

	return nil
}
