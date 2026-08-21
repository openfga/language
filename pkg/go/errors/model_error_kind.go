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

// wireName returns the name a category serialises as, and an empty string for a
// value with no name. It is the only place the mapping lives, so String, IsValid
// and MarshalText cannot disagree about which values have one.
func (m ModelErrorKind) wireName() string {
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
	default:
		// ModelErrorKindUnspecified lands here too: the zero value has no name by
		// design, so it marshals as a failure rather than as a category.
		return ""
	}
}

// modelErrorKindFromName is the reverse of wireName. The two are written out
// separately rather than derived from one another, so TestModelErrorKindWireNames
// round trips every declared category to keep them in step.
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

// String returns the wire name, or a diagnostic form for a value with none.
func (m ModelErrorKind) String() string {
	if name := m.wireName(); name != "" {
		return name
	}

	if m == ModelErrorKindUnspecified {
		return ""
	}

	return fmt.Sprintf("ModelErrorKind(%d)", int(m))
}

// IsValid reports whether m is a declared category with a wire name.
func (m ModelErrorKind) IsValid() bool {
	return m.wireName() != ""
}

// MarshalText emits the wire name, so the JSON carries "object-type".
//
// It does not go through String, because String has a diagnostic form for an
// undeclared number and this has to have none: a category that cannot be named
// must fail to marshal rather than ship as ModelErrorKind(99).
func (m ModelErrorKind) MarshalText() ([]byte, error) {
	name := m.wireName()
	if name == "" {
		return nil, fmt.Errorf("%w: %d", ErrUnknownModelErrorKind, int(m))
	}

	return []byte(name), nil
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
