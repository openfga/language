package validation

import (
	fgaerrors "github.com/openfga/language/pkg/go/errors"
)

// errorInfo is what a code implies beyond its message: its severity, the part of
// the model it belongs to, and the sentinel a caller matches with errors.Is.
type errorInfo struct {
	Severity fgaerrors.Severity
	Category fgaerrors.ModelErrorKind
	Cause    error

	// Critical marks a finding that invalidates the model as a whole rather than
	// one part of it, so a consumer may stop at the first one. Critical implies
	// blocking; TestCriticalImpliesBlocking enforces it.
	Critical bool
}

// errorInfoByType maps every code the validator emits to its severity, category,
// cause and criticality. It is the only place those are decided, so a code cannot
// mean one thing in the collector and another in a report.
//
// Every emitted code must appear here; TestErrorInfoCoversEveryEmittedErrorType
// enforces it, and declared-but-unemitted codes go in unemittedErrorTypes instead.
//
// Category is a default. DuplicatedError covers both a duplicate type and a
// duplicate type restriction, so a raise site overrides it through scope.category.
var errorInfoByType = map[ValidationErrorType]errorInfo{
	// Schema.
	InvalidSchema: {
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindInvalidModel,
		Cause:    fgaerrors.ErrInvalidSchemaVersion,
		Critical: true,
	},
	SchemaVersionUnsupported: {
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindInvalidModel,
		Cause:    fgaerrors.ErrSchemaVersionUnsupported,
	},
	SchemaVersionRequired: {
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindInvalidModel,
		Cause:    fgaerrors.ErrSchemaVersionRequired,
	},

	// Naming.
	InvalidName: {
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindRelation,
		Cause:    fgaerrors.ErrInvalidName,
	},
	ReservedTypeKeywords: {
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindObjectType,
		Cause:    fgaerrors.ErrReservedKeywords,
	},
	ReservedRelationKeywords: {
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindRelation,
		Cause:    fgaerrors.ErrReservedKeywords,
	},

	// Duplicates.
	DuplicatedError: {
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindRelation,
		Cause:    fgaerrors.ErrDuplicateDefinition,
		Critical: true,
	},

	// Undefined references.
	UndefinedType: {
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindObjectType,
		Cause:    fgaerrors.ErrObjectTypeUndefined,
		Critical: true,
	},
	UndefinedRelation: {
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindRelation,
		Cause:    fgaerrors.ErrRelationUndefined,
		Critical: true,
	},
	MissingDefinition: {
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindRelation,
		Cause:    fgaerrors.ErrRelationUndefined,
	},

	// Types and type restrictions.
	InvalidType: {
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindObjectType,
		Cause:    fgaerrors.ErrInvalidType,
	},
	InvalidRelationType: {
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindRelation,
		Cause:    fgaerrors.ErrInvalidRelationType,
		Critical: true,
	},
	AssignableRelationsMustHaveType: {
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindRelation,
		Cause:    fgaerrors.ErrDirectlyAssignableRelation,
	},

	// Tuplesets.
	InvalidRelationOnTupleset: {
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindRelation,
		Cause:    fgaerrors.ErrInvalidRelationOnTupleset,
	},
	TuplesetNotDirect: {
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindRelation,
		Cause:    fgaerrors.ErrInvalidRelationOnTuplesetNotDirect,
	},

	// Entrypoints.
	RelationNoEntrypoint: {
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindRelation,
		Cause:    fgaerrors.ErrNoEntrypoints,
		Critical: true,
	},

	// Wildcards.
	InvalidWildcardError: {
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindRelation,
		Cause:    fgaerrors.ErrInvalidWildcard,
	},
	TypeRestrictionCannotHaveWildcardAndRelation: {
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindRelation,
		Cause:    fgaerrors.ErrInvalidWildcard,
	},

	// Conditions.
	ConditionNotDefined: {
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindRelationCondition,
		Cause:    fgaerrors.ErrConditionUndefined,
	},
	ConditionNotUsed: {
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindCondition,
		Cause:    fgaerrors.ErrConditionUnReferenced,
	},
	DifferentNestedConditionName: {
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindCondition,
		Cause:    fgaerrors.ErrConditionNameMismatch,
	},

	// Modules.
	MultipleModulesInFile: {
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindInvalidModel,
		Cause:    fgaerrors.ErrMultipleModulesInFile,
		Critical: true,
	},

	// Weighted graph. The raise site chains the error the builder returned underneath
	// this sentinel, so a caller can match on the build being refused without knowing
	// which of the graph's own sentinels fired. Critical, because a model the graph
	// refuses has no graph, so nothing further can be reported about it.
	GraphModelUnbuildable: {
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindInvalidModel,
		Cause:    fgaerrors.ErrModelNotBuildable,
		Critical: true,
	},
}

// unemittedErrorTypes are declared ValidationErrorType values that no validation
// produces, established by inspecting every errorType argument reaching addError.
//
// They are kept rather than deleted because each has a published documentation
// page, and because SelfError and InvalidSyntax are equally unemitted in
// pkg/js/errors.ts. A cycle with no entrypoint surfaces as RelationNoEntrypoint,
// leaving CyclicError and CyclicRelation nothing to report. InvalidSchemaVersion is
// unreachable because RaiseInvalidSchemaVersion emits InvalidSchema, which is what
// the shared corpus expects.
//
// None get an errorInfoByType entry, so lookupErrorInfo treats them as blocking
// with no cause. Anything that starts emitting one must add it to the table in the
// same change.
var unemittedErrorTypes = map[ValidationErrorType]struct{}{
	SelfError:            {},
	InvalidSyntax:        {},
	CyclicError:          {},
	CyclicRelation:       {},
	InvalidSchemaVersion: {},
}

// allErrorTypes lists every declared ValidationErrorType. A Go const block of a
// string type cannot be enumerated at runtime, so exhaustiveness checks need it
// written out.
//
// Keep in sync with the const block in errors.go.
var allErrorTypes = []ValidationErrorType{
	SchemaVersionRequired,
	SchemaVersionUnsupported,
	ReservedTypeKeywords,
	ReservedRelationKeywords,
	SelfError,
	InvalidName,
	MissingDefinition,
	InvalidRelationType,
	InvalidRelationOnTupleset,
	InvalidType,
	RelationNoEntrypoint,
	TuplesetNotDirect,
	DuplicatedError,
	UndefinedType,
	UndefinedRelation,
	CyclicError,
	InvalidWildcardError,
	AssignableRelationsMustHaveType,
	InvalidSchema,
	InvalidSyntax,
	TypeRestrictionCannotHaveWildcardAndRelation,
	ConditionNotDefined,
	ConditionNotUsed,
	DifferentNestedConditionName,
	MultipleModulesInFile,
	CyclicRelation,
	InvalidSchemaVersion,
	GraphModelUnbuildable,
}

// isCriticalErrorType reports whether a code invalidates the model as a whole.
// Criticality is a field on the errorInfo entry, so a code cannot be critical and
// non-blocking at once. Unknown codes are blocking but not critical.
func isCriticalErrorType(errorType ValidationErrorType) bool {
	return lookupErrorInfo(errorType).Critical
}

// lookupErrorInfo returns the entry for a code.
//
// Unknown and unemitted codes fall back to a blocking error with no cause, so a
// code missing from the table cannot downgrade a finding to non-blocking.
func lookupErrorInfo(errorType ValidationErrorType) errorInfo {
	if entry, ok := errorInfoByType[errorType]; ok {
		return entry
	}
	return errorInfo{
		Severity: fgaerrors.SeverityError,
		Category: fgaerrors.ErrorKindInvalidModel,
	}
}
