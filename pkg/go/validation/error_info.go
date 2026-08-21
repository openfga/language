package validation

import (
	fgaerrors "github.com/openfga/language/pkg/go/errors"
)

// errorInfo is what a code implies beyond its message: its severity and the
// sentinel a caller matches with errors.Is.
type errorInfo struct {
	Severity fgaerrors.Severity
	Cause    error

	// Critical marks a finding that invalidates the model as a whole rather than
	// one part of it, so a consumer may stop at the first one. Critical implies
	// blocking; TestCriticalImpliesBlocking enforces it.
	Critical bool
}

// errorInfoByType maps every code the validator emits to its severity, cause and
// criticality. It is the only place those are decided, so a code cannot mean one
// thing in the collector and another in a report.
//
// Which part of the model a finding is about is not here, because it does not follow
// from the code. DuplicatedError covers a duplicate type and a duplicate type
// restriction; InvalidName covers a type, a relation and a condition. The raise site
// states it by building the cause it passes in scope.
//
// Every emitted code must appear here; TestErrorInfoCoversEveryEmittedErrorType
// enforces it, and declared-but-unemitted codes go in unemittedErrorTypes in
// error_info_test.go instead.
var errorInfoByType = map[ValidationErrorType]errorInfo{
	// Schema.
	InvalidSchema: {
		Severity: fgaerrors.SeverityError,
		Cause:    fgaerrors.ErrInvalidSchemaVersion,
		Critical: true,
	},
	SchemaVersionUnsupported: {
		Severity: fgaerrors.SeverityError,
		Cause:    fgaerrors.ErrSchemaVersionUnsupported,
	},
	SchemaVersionRequired: {
		Severity: fgaerrors.SeverityError,
		Cause:    fgaerrors.ErrSchemaVersionRequired,
	},

	// Naming.
	InvalidName: {
		Severity: fgaerrors.SeverityError,
		Cause:    fgaerrors.ErrInvalidName,
	},
	ReservedTypeKeywords: {
		Severity: fgaerrors.SeverityError,
		Cause:    fgaerrors.ErrReservedKeywords,
	},
	ReservedRelationKeywords: {
		Severity: fgaerrors.SeverityError,
		Cause:    fgaerrors.ErrReservedKeywords,
	},

	// Duplicates.
	DuplicatedError: {
		Severity: fgaerrors.SeverityError,
		Cause:    fgaerrors.ErrDuplicateDefinition,
		Critical: true,
	},

	// Undefined references.
	UndefinedType: {
		Severity: fgaerrors.SeverityError,
		Cause:    fgaerrors.ErrObjectTypeUndefined,
		Critical: true,
	},
	UndefinedRelation: {
		Severity: fgaerrors.SeverityError,
		Cause:    fgaerrors.ErrRelationUndefined,
		Critical: true,
	},
	MissingDefinition: {
		Severity: fgaerrors.SeverityError,
		Cause:    fgaerrors.ErrRelationUndefined,
	},

	// Types and type restrictions.
	InvalidType: {
		Severity: fgaerrors.SeverityError,
		Cause:    fgaerrors.ErrInvalidType,
	},
	InvalidRelationType: {
		Severity: fgaerrors.SeverityError,
		Cause:    fgaerrors.ErrInvalidRelationType,
		Critical: true,
	},
	AssignableRelationsMustHaveType: {
		Severity: fgaerrors.SeverityError,
		Cause:    fgaerrors.ErrDirectlyAssignableRelation,
	},

	// Tuplesets.
	InvalidRelationOnTupleset: {
		Severity: fgaerrors.SeverityError,
		Cause:    fgaerrors.ErrInvalidRelationOnTupleset,
	},
	TuplesetNotDirect: {
		Severity: fgaerrors.SeverityError,
		Cause:    fgaerrors.ErrInvalidRelationOnTuplesetNotDirect,
	},

	// Entrypoints.
	RelationNoEntrypoint: {
		Severity: fgaerrors.SeverityError,
		Cause:    fgaerrors.ErrNoEntrypoints,
		Critical: true,
	},

	// Wildcards.
	InvalidWildcardError: {
		Severity: fgaerrors.SeverityError,
		Cause:    fgaerrors.ErrInvalidWildcard,
	},
	TypeRestrictionCannotHaveWildcardAndRelation: {
		Severity: fgaerrors.SeverityError,
		Cause:    fgaerrors.ErrInvalidWildcard,
	},

	// Conditions.
	ConditionNotDefined: {
		Severity: fgaerrors.SeverityError,
		Cause:    fgaerrors.ErrConditionUndefined,
	},
	ConditionNotUsed: {
		Severity: fgaerrors.SeverityError,
		Cause:    fgaerrors.ErrConditionUnReferenced,
	},
	DifferentNestedConditionName: {
		Severity: fgaerrors.SeverityError,
		Cause:    fgaerrors.ErrConditionNameMismatch,
	},

	// Modules.
	MultipleModulesInFile: {
		Severity: fgaerrors.SeverityError,
		Cause:    fgaerrors.ErrMultipleModulesInFile,
		Critical: true,
	},
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
	}
}
