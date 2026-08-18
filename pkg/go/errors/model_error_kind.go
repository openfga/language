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

// modelErrorKindNames maps each category to its wire name. A category missing
// from here fails to marshal, so a constant added without a name is caught.
var modelErrorKindNames = map[ModelErrorKind]string{
	ErrorKindObjectType:        "object-type",
	ErrorKindRelation:          "relation",
	ErrorKindRelationCondition: "relation-condition",
	ErrorKindCondition:         "condition",
	ErrorKindInvalidModel:      "invalid-model",
}

// modelErrorKindValues is the reverse of modelErrorKindNames, built from it so
// the two cannot disagree.
var modelErrorKindValues = func() map[string]ModelErrorKind {
	values := make(map[string]ModelErrorKind, len(modelErrorKindNames))
	for errorType, name := range modelErrorKindNames {
		values[name] = errorType
	}

	return values
}()

// String returns the wire name, or a diagnostic form for a value with none.
func (m ModelErrorKind) String() string {
	if name, ok := modelErrorKindNames[m]; ok {
		return name
	}

	if m == ModelErrorKindUnspecified {
		return ""
	}

	return fmt.Sprintf("ModelErrorKind(%d)", int(m))
}

// IsValid reports whether m is a declared category with a wire name.
func (m ModelErrorKind) IsValid() bool {
	_, ok := modelErrorKindNames[m]

	return ok
}

// MarshalText emits the wire name, so the JSON carries "object-type".
func (m ModelErrorKind) MarshalText() ([]byte, error) {
	name, ok := modelErrorKindNames[m]
	if !ok {
		return nil, fmt.Errorf("%w: %d", ErrUnknownModelErrorKind, int(m))
	}

	return []byte(name), nil
}

// UnmarshalText resolves a wire name back to its category, rejecting any name
// this package does not declare.
func (m *ModelErrorKind) UnmarshalText(text []byte) error {
	errorType, ok := modelErrorKindValues[string(text)]
	if !ok {
		return fmt.Errorf("%w: %q", ErrUnknownModelErrorKind, text)
	}

	*m = errorType

	return nil
}
