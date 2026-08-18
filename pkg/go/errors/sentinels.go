package errors

import "errors"

// Sentinel errors for the conditions model validation reports.
//
// Callers branch on what went wrong with errors.Is rather than on message text.
// Every validation finding wraps exactly one of these, and only conditions the
// validator actually reports get one.
var (
	// ErrInvalidSchemaVersion is reported for a schema version that was never a
	// valid one, as against one that is no longer supported.
	ErrInvalidSchemaVersion = errors.New("invalid schema version")

	// ErrSchemaVersionUnsupported is reported for a schema version that was
	// once valid but is no longer supported.
	ErrSchemaVersionUnsupported = errors.New("schema version no longer supported")

	// ErrSchemaVersionRequired is reported when a model declares no schema
	// version.
	ErrSchemaVersionRequired = errors.New("schema version required")

	// ErrReservedKeywords is reported when a type or relation is named with a
	// reserved word.
	ErrReservedKeywords = errors.New("self and this are reserved keywords")

	// ErrInvalidName is reported when a type or relation name does not match
	// the naming rules.
	ErrInvalidName = errors.New("invalid name")

	// ErrDuplicateDefinition is reported when a type, relation or type
	// restriction is defined more than once.
	ErrDuplicateDefinition = errors.New("duplicate definition")

	// ErrObjectTypeUndefined is reported when a model references an object type
	// that is not defined.
	ErrObjectTypeUndefined = errors.New("undefined object type")

	// ErrRelationUndefined is reported when a model references a relation that
	// is not defined on the type it is used with.
	ErrRelationUndefined = errors.New("undefined relation")

	// ErrInvalidType is reported when a type restriction names something that
	// is not a valid type.
	ErrInvalidType = errors.New("invalid type")

	// ErrInvalidRelationType is reported when a relation is not valid for the
	// type it is referenced against.
	ErrInvalidRelationType = errors.New("invalid relation for type")

	// ErrInvalidRelationOnTupleset is reported when a tupleset relation
	// references a relation that does not exist on the related type.
	ErrInvalidRelationOnTupleset = errors.New("invalid relation on tupleset")

	// ErrInvalidRelationOnTuplesetNotDirect is reported when a relation used
	// inside a `from` clause is not a direct relation.
	ErrInvalidRelationOnTuplesetNotDirect = errors.New(
		"relations that are referenced in a tupleset must be defined with a direct relation")

	// ErrNoEntrypoints is reported when a relation can never be satisfied,
	// either because nothing can enter it or because it only refers back to
	// itself.
	ErrNoEntrypoints = errors.New("no entrypoints defined")

	// ErrDirectlyAssignableRelation is reported when an assignable relation
	// declares no assignable types.
	ErrDirectlyAssignableRelation = errors.New("a direct assignment must contain at least one object type or userset")

	// ErrInvalidWildcard is reported when a wildcard is used somewhere it is
	// not permitted, including alongside a relation in the same type
	// restriction.
	ErrInvalidWildcard = errors.New("invalid wildcard usage")

	// ErrConditionUndefined is reported when a relation references a condition
	// that the model does not define.
	ErrConditionUndefined = errors.New("condition is not defined")

	// ErrConditionUnReferenced is reported when a condition is defined but
	// never used.
	ErrConditionUnReferenced = errors.New("condition is defined but not referenced")

	// ErrConditionNameMismatch is reported when a condition's key differs from
	// the name declared inside it.
	ErrConditionNameMismatch = errors.New("condition name does not match its nested name")

	// ErrMultipleModulesInFile is reported when one file declares more than one
	// module.
	ErrMultipleModulesInFile = errors.New("file contains multiple modules")

	// ErrUnknownModelErrorKind is returned when a ModelErrorKind has no wire
	// name, either marshalling a value this package does not declare or reading
	// a name it does not recognise. It is not a validation finding.
	ErrUnknownModelErrorKind = errors.New("unknown model error type")

	// ErrUnknownSeverity is returned when a Severity has no wire name, either
	// marshalling a value this package does not declare or reading a name it
	// does not recognise. It is not a validation finding.
	ErrUnknownSeverity = errors.New("unknown severity")
)
