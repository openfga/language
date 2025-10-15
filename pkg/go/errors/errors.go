package errors

import (
	"errors"
	"fmt"
)

var (
	// ErrDuplicateTypes is returned when an authorization model contains duplicate types.
	ErrDuplicateTypes = errors.New("an authorization model cannot contain duplicate types")

	// ErrDuplicateRelationsType is returned when an authorization model contains duplicate relations for the same type.
	ErrDuplicateRelationsType = errors.New("an authorization model cannot contain duplicate relations for the same type")

	// ErrInvalidSchemaVersion is returned for an invalid schema version in the authorization model.
	ErrInvalidSchemaVersion = errors.New("invalid schema version")

	// ErrInvalidModel is returned when encountering an invalid authorization model.
	ErrInvalidModel = errors.New("invalid authorization model encountered")

	// ErrRelationUndefined is returned when encountering an undefined relation in the authorization model.
	ErrRelationUndefined = errors.New("undefined relation")

	// ErrObjectTypeUndefined is returned when encountering an undefined object type in the authorization model.
	ErrObjectTypeUndefined = errors.New("undefined object type")

	// ErrInvalidUsersetRewrite is returned for an invalid userset rewrite definition.
	ErrInvalidUsersetRewrite = errors.New("invalid userset rewrite definition")

	// ErrCycle is returned when a cycle is detected in an authorization model.
	// This occurs if an objectType and relation in the model define a rewrite
	// rule that is self-referencing through computed relationships.
	ErrCycle = errors.New("an authorization model cannot contain a model cycle")

	// ErrConstraintTupleCycle is returned when a constraint tuple is part of a model cycle.
	ErrConstraintTupleCycle = errors.New("operands AND or BUT NOT cannot be part of a model cycle")

	// ErrNoEntrypoints is returned when a particular objectType and relation in an authorization
	// model are not accessible via a direct edge, for example from another objectType.
	ErrNoEntrypoints = errors.New("no entrypoints defined")

	// ErrNoEntryPointsLoop is returned when an authorization model contains a cycle
	// because at least one objectType and relation returned ErrNoEntrypoints.
	ErrNoEntryPointsLoop = errors.New("potential loop")

	// ErrCondition is returned when no condition is defined for a relation in the authorization model.
	ErrCondition = errors.New("condition is not defined")

	// ErrConditionUnReferenced is returned when a condition is defined but not referenced in the authorization model.
	ErrConditionUnReferenced = errors.New("condition is defined but not referenced")

	// ErrInvalidType is returned when the type name of a type definition is invalid.
	ErrInvalidType = errors.New("the type name of a type definition cannot be an empty string")

	// ErrReservedKeywords is returned when using reserved keywords "self" and "this".
	ErrReservedKeywords = errors.New("self and this are reserved keywords")

	// ErrInvalidRelation is returned when the relation name of a type is invalid.
	ErrInvalidRelation = errors.New("the relation name of a type cannot be an empty string")

	// ErrInvalidRelationOnTupleset is returned when a relation is referenced in a tupleset but is not defined as a direct relation.
	ErrInvalidRelationOnTupleset = errors.New("relations that are referenced in a tupleset must be defined with a direct relation")

	// ErrDirectlyAssignableRelation is returned when a direct assignment is invalid.
	ErrDirectlyAssignableRelation = errors.New("the a direct assignation must contain at least one object type or userset")

	// ErrInvalidWildcard is returned when the wildcard usage is invalid.
	ErrInvalidWildcard = errors.New("invalid wildcard usage")
)

// ModelErrorType indicates the type of model error
type ModelErrorType int

const (
	// ErrorTypeObjectType indicates an error related only to an object type
	ErrorTypeObjectType ModelErrorType = iota
	// ErrorTypeRelation indicates an error related to a relation and object type
	ErrorTypeRelation
	// ErrorTypeCondition indicates an error related to a condition, relation, and object type
	ErrorRelationCondition
	// ErrorTypeInvalidModel indicates an invalid authorization model error
	ErrorTypeInvalidModel
	// ErrorCondition indicates an error related to a condition
	ErrorCondition
)

func UnsupportedDSLNestingError(typeName string, relationName string) error {
	return fmt.Errorf( //nolint:goerr113
		"the '%s' relation definition under the '%s' type is not supported by the OpenFGA DSL syntax yet",
		relationName,
		typeName,
	)
}

func ConditionNameDoesntMatchError(conditionName string, conditionNestedName string) error {
	return fmt.Errorf( //nolint:goerr113
		"the '%s' condition has a different nested condition name ('%s')",
		conditionName,
		conditionNestedName,
	)
}

type AuthorizationModelError struct {
	Cause      error
	ObjectType string
	Relation   string
	Condition  string
	Type       ModelErrorType
}

func (e *AuthorizationModelError) Error() string {
	switch e.Type {
	case ErrorTypeObjectType:
		return fmt.Sprintf("error in the definition of the object type '%s': %s", e.ObjectType, e.Cause)
	case ErrorTypeRelation:
		return fmt.Sprintf("error in the definition of relation '%s' of object type '%s': %s", e.Relation, e.ObjectType, e.Cause)
	case ErrorRelationCondition:
		return fmt.Sprintf("error in the definition of condition '%s' of relation '%s' in object type '%s': %s",
			e.Condition, e.Relation, e.ObjectType, e.Cause)
	case ErrorCondition:
		return fmt.Sprintf("error in the definition of condition '%s': %s", e.Condition, e.Cause)
	default:
		return fmt.Sprintf("error in authorization model: %s", e.Cause)
	}
}

// Unwrap returns the wrapped error
func (e *AuthorizationModelError) Unwrap() error {
	return e.Cause
}

// ObjectTypeError creates a ModelError specific to an object type
func ObjectTypeError(objectType string, cause error) error {
	return &AuthorizationModelError{
		Type:       ErrorTypeObjectType,
		ObjectType: objectType,
		Cause:      cause,
	}
}

// RelationObjectTypeError creates a ModelError specific to a relation and object type
func RelationObjectTypeError(relation, objectType string, cause error) error {
	return &AuthorizationModelError{
		Type:       ErrorTypeRelation,
		Relation:   relation,
		ObjectType: objectType,
		Cause:      cause,
	}
}

// ConditionRelationObjectTypeError creates a ModelError specific to a condition, relation, and object type
func ConditionRelationObjectTypeError(condition, relation, objectType string, cause error) error {
	return &AuthorizationModelError{
		Type:       ErrorRelationCondition,
		Condition:  condition,
		Relation:   relation,
		ObjectType: objectType,
		Cause:      cause,
	}
}

// ConditionRelationObjectTypeError creates a ModelError specific to a condition, relation, and object type
func ConditionError(condition string, cause error) error {
	return &AuthorizationModelError{
		Type:      ErrorCondition,
		Condition: condition,
		Cause:     cause,
	}
}

// ConditionRelationObjectTypeError creates a ModelError specific to a condition, relation, and object type
func InvalidAuthorizationModelError(cause error) error {
	return &AuthorizationModelError{
		Type:  ErrorTypeInvalidModel,
		Cause: cause,
	}
}
