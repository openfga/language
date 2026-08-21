package validation

import (
	"errors"
	"fmt"
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

// ErrorCollector collects validation errors during model validation.
// This is equivalent to the JS ExceptionCollector class.
type ErrorCollector struct {
	errors []*ValidationError
	lines  []string // DSL lines for line number resolution
}

// NewErrorCollector creates a new error collector.
func NewErrorCollector(lines []string) *ErrorCollector {
	return &ErrorCollector{
		errors: make([]*ValidationError, 0),
		lines:  lines,
	}
}

// AllFindings returns every collected finding, blocking or not. The collector is the
// raw record; ValidationErrors is where findings are filtered by severity, which is
// why this is not called GetErrors: on ValidationErrors that name means the blocking
// ones only.
func (c *ErrorCollector) AllFindings() []*ValidationError {
	return c.errors
}

// HasErrors reports whether any collected finding makes the model invalid.
//
// The cascade in RunAllValidations gates on this, so it counts blocking findings
// only: one advisory must not skip every phase that runs after it.
func (c *ErrorCollector) HasErrors() bool {
	for _, err := range c.errors {
		if err.Blocks() {
			return true
		}
	}
	return false
}

// Count returns the number of collected findings that make the model invalid.
func (c *ErrorCollector) Count() int {
	count := 0
	for _, err := range c.errors {
		if err.Blocks() {
			count++
		}
	}
	return count
}

// CountAll returns the total number of collected findings, blocking or not.
func (c *ErrorCollector) CountAll() int {
	return len(c.errors)
}

// scope names the model entity a finding is about, so addScopedError can build the
// cause and derive the metadata from one description. A zero scope means the raise
// site has nothing to add beyond the symbol, and the table's category stands alone.
type scope struct {
	objectType string
	relation   string
	condition  string

	// offendingType is the enclosing type a finding about another type was written
	// in, matching JS's wire field of the same name. Metadata only: no scoped error
	// type has a slot for it.
	offendingType string

	// category overrides the table default when set, for codes raised from places
	// with different scopes: a duplicate type and a duplicate type restriction share
	// one code without being the same kind of finding.
	category fgaerrors.ModelErrorKind

	// cause overrides the table's sentinel when set, for a finding that carries an
	// error it did not raise itself. It is wrapped in the scoped type as any sentinel
	// would be, so errors.Is still reaches whatever the original error wrapped.
	cause error
}

// addError is a helper to add an error to the collection.
func (c *ErrorCollector) addError(message string, errorType ValidationErrorType, symbol string,
	lineIndex *int, meta *Meta, customResolver ErrorCustomResolver) {
	c.addScopedError(message, errorType, symbol, lineIndex, meta, customResolver, scope{})
}

// addScopedError adds an error that knows which type, relation or condition it
// concerns. Callers with nothing to add beyond the symbol use addError instead.
func (c *ErrorCollector) addScopedError(message string, errorType ValidationErrorType, symbol string,
	lineIndex *int, meta *Meta, customResolver ErrorCustomResolver, errorScope scope) {
	var line *Range
	var column *Range

	// Calculate line and column positions if lineIndex is provided
	if lineIndex != nil && *lineIndex >= 0 && *lineIndex < len(c.lines) {
		line = &Range{Start: *lineIndex, End: *lineIndex}

		// Find symbol position in line for column calculation, matching on word
		// boundaries as the reference does.
		rawLine := c.lines[*lineIndex]
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
	}

	entry := lookupErrorInfo(errorType)

	category := entry.Category
	if errorScope.category != fgaerrors.ModelErrorKindUnspecified {
		category = errorScope.category
	}

	sentinel := entry.Cause
	if errorScope.cause != nil {
		sentinel = errorScope.cause
	}

	// The cause carries the scope and the metadata is derived from it, so the JSON
	// and the errors.As payload cannot disagree. offendingType is metadata only, so
	// it comes straight off the scope.
	cause := newScopedCause(category, errorScope, sentinel)
	objectType, relation, condition := causeScope(cause)

	metadata := &ErrorMetadata{
		Symbol:        symbol,
		ErrorType:     errorType,
		OffendingType: errorScope.offendingType,
		Type:          objectType,
		Relation:      relation,
		Condition:     condition,
	}

	if meta != nil {
		// Module goes in the metadata, file on the error itself, matching the
		// JS implementation.
		metadata.Module = meta.Module
	}

	validationErr := &ValidationError{
		Message:  message,
		Severity: entry.Severity,
		Category: category,
		Line:     line,
		Column:   column,
		Metadata: metadata,
		Cause:    cause,
	}

	if meta != nil {
		validationErr.File = meta.File
	}

	c.errors = append(c.errors, validationErr)
}

// newScopedCause wraps sentinel in the error type matching category, carrying
// whichever scope fields that type declares. Returns nil when there is no sentinel,
// which is the case for codes absent from the table.
func newScopedCause(category fgaerrors.ModelErrorKind, errorScope scope, sentinel error) error {
	if sentinel == nil {
		return nil
	}

	switch category {
	case fgaerrors.ErrorKindObjectType:
		return &fgaerrors.ErrObjectType{
			ObjectType: errorScope.objectType,
			Cause:      sentinel,
		}
	case fgaerrors.ErrorKindRelation:
		return &fgaerrors.ErrRelation{
			ObjectType: errorScope.objectType,
			Relation:   errorScope.relation,
			Cause:      sentinel,
		}
	case fgaerrors.ErrorKindRelationCondition:
		return &fgaerrors.ErrRelationCondition{
			ObjectType: errorScope.objectType,
			Relation:   errorScope.relation,
			Condition:  errorScope.condition,
			Cause:      sentinel,
		}
	case fgaerrors.ErrorKindCondition:
		return &fgaerrors.ErrCondition{
			Condition: errorScope.condition,
			Cause:     sentinel,
		}
	default:
		// ErrorKindInvalidModel, and anything unrecognised: a finding no part of the
		// model owns.
		return &fgaerrors.ErrModel{Cause: sentinel}
	}
}

// causeScope reads the scope off whichever error type cause is, so the metadata
// carries exactly the fields that type declares. A cause with no scope to report,
// *ErrModel or nil, yields three empty strings, which omitempty drops.
func causeScope(cause error) (objectType, relation, condition string) {
	var (
		objectTypeErr        *fgaerrors.ErrObjectType
		relationErr          *fgaerrors.ErrRelation
		relationConditionErr *fgaerrors.ErrRelationCondition
		conditionErr         *fgaerrors.ErrCondition
	)

	switch {
	case errors.As(cause, &objectTypeErr):
		return objectTypeErr.ObjectType, "", ""
	case errors.As(cause, &relationErr):
		return relationErr.ObjectType, relationErr.Relation, ""
	case errors.As(cause, &relationConditionErr):
		return relationConditionErr.ObjectType, relationConditionErr.Relation, relationConditionErr.Condition
	case errors.As(cause, &conditionErr):
		return "", "", conditionErr.Condition
	default:
		return "", "", ""
	}
}

// RaiseInvalidName raises an invalid name error.
func (c *ErrorCollector) RaiseInvalidName(symbol, clause string, typeName *string, lineIndex *int, meta *Meta) {
	var message string
	// A nil typeName means the offending name is a type rather than a relation on
	// one, which changes both the message and the scope of the finding.
	errorScope := scope{objectType: symbol, category: fgaerrors.ErrorKindObjectType}

	if typeName != nil {
		message = fmt.Sprintf("relation '%s' of type '%s' does not match naming rule: '%s'.", symbol, *typeName, clause)
		errorScope = scope{objectType: *typeName, relation: symbol}
	} else {
		message = fmt.Sprintf("type '%s' does not match naming rule: '%s'.", symbol, clause)
	}

	c.addScopedError(message, InvalidName, symbol, lineIndex, meta, nil, errorScope)
}

// RaiseInvalidConditionName raises an invalid name error for a condition, scoped
// to the condition rather than RaiseInvalidName's type or relation.
func (c *ErrorCollector) RaiseInvalidConditionName(symbol, clause string, lineIndex *int, meta *Meta) {
	message := fmt.Sprintf("condition '%s' does not match naming rule: '%s'.", symbol, clause)
	c.addScopedError(message, InvalidName, symbol, lineIndex, meta, nil, scope{
		condition: symbol,
		category:  fgaerrors.ErrorKindCondition,
	})
}

// RaiseReservedTypeName raises a reserved type name error.
func (c *ErrorCollector) RaiseReservedTypeName(symbol string, lineIndex *int, meta *Meta) {
	message := "a type cannot be named 'self' or 'this'."
	c.addScopedError(message, ReservedTypeKeywords, symbol, lineIndex, meta, nil, scope{
		objectType: symbol,
	})
}

// RaiseReservedRelationName raises a reserved relation name error.
func (c *ErrorCollector) RaiseReservedRelationName(symbol, typeName string, lineIndex *int, meta *Meta) {
	message := "a relation cannot be named 'self' or 'this'."
	c.addScopedError(message, ReservedRelationKeywords, symbol, lineIndex, meta, nil, scope{
		objectType: typeName,
		relation:   symbol,
	})
}

// RaiseTupleUsersetRequiresDirect raises an error for tuple-to-userset not being direct.
func (c *ErrorCollector) RaiseTupleUsersetRequiresDirect(symbol, typeName, relation string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("`%s` relation used inside from allows only direct relation.", symbol)

	// Custom resolver for "from" clause positioning
	customResolver := func(wordIdx int, rawLine, value string) int {
		clauseStartsAt := strings.Index(rawLine, "from") + len("from")
		if clauseStartsAt >= len("from") {
			wordIdx = clauseStartsAt + strings.Index(rawLine[clauseStartsAt:], value)
		}
		return wordIdx
	}

	c.addScopedError(message, TuplesetNotDirect, symbol, lineIndex, meta, customResolver, scope{
		objectType: typeName,
		relation:   relation,
	})
}

// RaiseDuplicateTypeName raises a duplicate type name error.
func (c *ErrorCollector) RaiseDuplicateTypeName(symbol string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("the type `%s` is a duplicate.", symbol)
	// A duplicate type is about the type, not a relation on it, so this overrides
	// DuplicatedError's relation-scoped default.
	c.addScopedError(message, DuplicatedError, symbol, lineIndex, meta, nil, scope{
		objectType: symbol,
		category:   fgaerrors.ErrorKindObjectType,
	})
}

// RaiseDuplicateTypeRestriction raises a duplicate type restriction error.
func (c *ErrorCollector) RaiseDuplicateTypeRestriction(symbol, relationName, typeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("the type restriction `%s` is a duplicate in the relation `%s`.", symbol, relationName)
	c.addScopedError(message, DuplicatedError, symbol, lineIndex, meta, nil, scope{
		objectType: typeName,
		relation:   relationName,
	})
}

// RaiseUndefinedType raises an error for undefined type references.
func (c *ErrorCollector) RaiseUndefinedType(typeName, relationName, parentTypeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("Type '%s' is not defined (referenced in relation '%s' of type '%s')", typeName, relationName, parentTypeName)
	// The undefined type is the subject; parentTypeName is only where it was
	// referenced from, so the scope names the type that does not exist.
	c.addScopedError(message, UndefinedType, typeName, lineIndex, meta, nil, scope{
		objectType: typeName,
	})
}

// RaiseUndefinedRelation raises an error for undefined relation references.
func (c *ErrorCollector) RaiseUndefinedRelation(relationName, typeName, parentRelation, parentTypeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("Relation '%s' is not defined on type '%s' (referenced in relation '%s' of type '%s')", relationName, typeName, parentRelation, parentTypeName)
	c.addScopedError(message, UndefinedRelation, relationName, lineIndex, meta, nil, scope{
		objectType: typeName,
		relation:   relationName,
	})
}

// RaiseDuplicateType raises a duplicate type error in relation.
func (c *ErrorCollector) RaiseDuplicateType(symbol, relationName, typeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("the partial relation definition `%s` is a duplicate in the relation `%s`.",
		symbol, relationName)
	c.addScopedError(message, DuplicatedError, symbol, lineIndex, meta, nil, scope{
		objectType: typeName,
		relation:   relationName,
	})
}

// RaiseDuplicateRelationshipDefinition raises a duplicate relationship definition error.
func (c *ErrorCollector) RaiseDuplicateRelationshipDefinition(symbol string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("the relation '%s' is defined more than once.", symbol)
	c.addScopedError(message, DuplicatedError, symbol, lineIndex, meta, nil, scope{
		relation: symbol,
	})
}

// RaiseNoEntryPointLoop raises an error for impossible relation with potential loop.
func (c *ErrorCollector) RaiseNoEntryPointLoop(symbol, typeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("`%s` is an impossible relation for `%s` (potential loop).", symbol, typeName)
	c.addScopedError(message, RelationNoEntrypoint, symbol, lineIndex, meta, nil, scope{
		objectType: typeName,
		relation:   symbol,
	})
}

// RaiseNoEntryPoint raises an error for impossible relation without entry point.
func (c *ErrorCollector) RaiseNoEntryPoint(symbol, typeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("`%s` is an impossible relation for `%s` (no entrypoint).", symbol, typeName)
	c.addScopedError(message, RelationNoEntrypoint, symbol, lineIndex, meta, nil, scope{
		objectType: typeName,
		relation:   symbol,
	})
}

// RaiseInvalidRelationOnTupleset raises an error for invalid relation on tupleset.
func (c *ErrorCollector) RaiseInvalidRelationOnTupleset(symbol, typeName, typeDef, relationName,
	offendingRelation, parent string, lineIndex *int, meta *Meta) {
	message := fmt.Sprintf("the `%s` relation definition on type `%s` is not valid: `%s` does not exist on `%s`, which is of type `%s`.",
		offendingRelation, typeDef, offendingRelation, parent, typeName)
	c.addScopedError(message, InvalidRelationOnTupleset, symbol, lineIndex, meta, nil, scope{
		objectType: typeDef,
		relation:   relationName,
	})
}

// RaiseInvalidTypeRelation raises an error for invalid type relation.
func (c *ErrorCollector) RaiseInvalidTypeRelation(symbol, typeName, relationName, offendingRelation,
	offendingType string, lineIndex *int, meta *Meta) {
	message := fmt.Sprintf("`%s` is not a valid relation for `%s`.", offendingRelation, typeName)
	c.addScopedError(message, InvalidRelationType, symbol, lineIndex, meta, nil, scope{
		objectType:    typeName,
		relation:      relationName,
		offendingType: offendingType,
	})
}

// RaiseInvalidType raises an error for invalid type.
func (c *ErrorCollector) RaiseInvalidType(symbol, typeName, relation string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("`%s` is not a valid type.", symbol)
	// The invalid type appears in the assignable-types list (after the colon),
	// which may share a name with the relation key before the colon. Resolve the
	// column to the value side so it marks the type, not the relation name —
	// mirroring the reference's customResolver.
	resolver := func(_ int, rawLine, sym string) int {
		colon := strings.Index(rawLine, ":")
		if colon < 0 {
			return wordIndex(rawLine, sym)
		}
		value := rawLine[colon+1:]
		idx := wordIndex(value, sym)
		return colon + 1 + idx
	}
	c.addScopedError(message, InvalidType, symbol, lineIndex, meta, resolver, scope{
		objectType: symbol,
	})
}

// RaiseAssignableRelationMustHaveTypes raises an error for assignable relations without types.
func (c *ErrorCollector) RaiseAssignableRelationMustHaveTypes(symbol string, lineIndex *int) {
	message := fmt.Sprintf("the assignable relation '%s' must have at least one assignable type.", symbol)
	c.addScopedError(message, AssignableRelationsMustHaveType, symbol, lineIndex, nil, nil, scope{
		relation: symbol,
	})
}

// RaiseAssignableTypeWildcardRelation raises an error for wildcard with relation.
func (c *ErrorCollector) RaiseAssignableTypeWildcardRelation(symbol, typeName, relation string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("the type restriction '%s' on relation '%s' of type '%s' is not allowed to have both a wildcard and a relation.",
		symbol, relation, typeName)
	c.addScopedError(message, TypeRestrictionCannotHaveWildcardAndRelation, symbol, lineIndex, meta, nil, scope{
		objectType: typeName,
		relation:   relation,
	})
}

// RaiseInvalidRelationError reports a rewrite that names a relation the type does
// not define. The message names the missing relation only, as the reference's does;
// it does not list the relations that do exist.
func (c *ErrorCollector) RaiseInvalidRelationError(symbol, typeName, relation string,
	lineIndex *int, meta *Meta) {
	message := fmt.Sprintf("the relation `%s` does not exist.", symbol)
	c.addScopedError(message, MissingDefinition, symbol, lineIndex, meta, nil, scope{
		objectType: typeName,
		relation:   relation,
	})
}

// RaiseInvalidSchemaVersion raises an error for a schema version that was never
// valid (e.g. "0.9", "2.0"). This is distinct from a version that is recognized
// but no longer supported (see RaiseSchemaVersionUnsupported).
func (c *ErrorCollector) RaiseInvalidSchemaVersion(symbol string, lineIndex *int) {
	message := fmt.Sprintf("invalid schema %s", symbol)
	c.addError(message, InvalidSchema, symbol, lineIndex, nil, nil)
}

// RaiseSchemaVersionUnsupported raises an error for a recognized but retired
// schema version (e.g. "1.0").
func (c *ErrorCollector) RaiseSchemaVersionUnsupported(symbol string, lineIndex *int) {
	message := "schema version no longer supported"
	c.addError(message, SchemaVersionUnsupported, symbol, lineIndex, nil, nil)
}

// RaiseSchemaVersionRequired raises an error for missing schema version.
func (c *ErrorCollector) RaiseSchemaVersionRequired(symbol string, lineIndex *int) {
	message := "schema version required"
	c.addError(message, SchemaVersionRequired, symbol, lineIndex, nil, nil)
}

// RaiseMaximumOneDirectRelationship raises an error for multiple direct relationships.
func (c *ErrorCollector) RaiseMaximumOneDirectRelationship(symbol string, lineIndex *int) {
	message := fmt.Sprintf("the relation '%s' can have at most one direct relationship.", symbol)
	c.addScopedError(message, DuplicatedError, symbol, lineIndex, nil, nil, scope{
		relation: symbol,
	})
}

// RaiseInvalidConditionNameInParameter raises an error for invalid condition names.
func (c *ErrorCollector) RaiseInvalidConditionNameInParameter(symbol, typeName, relationName, conditionName string,
	meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("`%s` is not a defined condition in the model.", conditionName)
	// Scoped to the relation the condition is applied to, not the condition's own
	// definition: the condition does not exist to have a definition.
	c.addScopedError(message, ConditionNotDefined, symbol, lineIndex, meta, nil, scope{
		objectType: typeName,
		relation:   relationName,
		condition:  conditionName,
	})
}

// RaiseUnusedCondition raises an error for unused conditions.
func (c *ErrorCollector) RaiseUnusedCondition(symbol string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("`%s` condition is not used in the model.", symbol)
	c.addScopedError(message, ConditionNotUsed, symbol, lineIndex, meta, nil, scope{
		condition: symbol,
	})
}

// RaiseDifferentNestedConditionName raises an error for a condition whose nested
// name property differs from its map key. The message mirrors the reference.
func (c *ErrorCollector) RaiseDifferentNestedConditionName(condition, nestedConditionName string) {
	message := fmt.Sprintf("condition key is `%s` but nested name property is %s", condition, nestedConditionName)
	c.addScopedError(message, DifferentNestedConditionName, nestedConditionName, nil, nil, nil, scope{
		condition: condition,
	})
}

// RaiseMultipleModulesInSingleFile raises an error for multiple modules in single
// file. The modules are listed in the order the model declares them, and the message
// mirrors the reference.
func (c *ErrorCollector) RaiseMultipleModulesInSingleFile(file string, modules []string) {
	moduleList := strings.Join(modules, ", ")
	message := fmt.Sprintf("file %s would contain multiple module definitions (%s) when transforming to DSL. "+
		"Only one module can be defined per file.", file, moduleList)
	c.addError(message, MultipleModulesInFile, file, nil, nil, nil)
}

// Complex operation validation error methods

// RaiseRedundantUnionMember raises an error for redundant members in union operations.
func (c *ErrorCollector) RaiseRedundantUnionMember(operation, relationName, typeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("Redundant operation '%s' found in union for relation '%s' of type '%s'", operation, relationName, typeName)
	c.addScopedError(message, DuplicatedError, operation, lineIndex, meta, nil, scope{
		objectType: typeName,
		relation:   relationName,
	})
}

// RaiseImpossibleIntersection raises an error for intersection operations that cannot succeed.
func (c *ErrorCollector) RaiseImpossibleIntersection(relationName, typeName string, conflictingTypes []string, meta *Meta, lineIndex *int) {
	typeList := strings.Join(conflictingTypes, ", ")
	message := fmt.Sprintf("Impossible intersection in relation '%s' of type '%s': conflicting types [%s]", relationName, typeName, typeList)
	c.addScopedError(message, InvalidRelationType, relationName, lineIndex, meta, nil, scope{
		objectType: typeName,
		relation:   relationName,
	})
}

// RaiseEmptyDifference raises an error for difference operations that result in empty sets.
func (c *ErrorCollector) RaiseEmptyDifference(relationName, typeName, operation string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("Empty difference operation in relation '%s' of type '%s': subtracting '%s' from itself", relationName, typeName, operation)
	c.addScopedError(message, RelationNoEntrypoint, relationName, lineIndex, meta, nil, scope{
		objectType: typeName,
		relation:   relationName,
	})
}

// Graph validation error methods

// RaiseModelUnbuildable raises an error for a model the weighted graph refuses to build.
//
// The finding carries no position, and one refused model raises one finding however many
// problems it has. Both follow from the builder returning on the first problem it meets
// and returning no graph with it: there is nothing left to walk for the rest, and the
// error it returns names a count rather than the relations responsible.
//
// The builder's error is chained under ErrModelNotBuildable, so errors.Is matches the
// build being refused and the specific reason alike.
func (c *ErrorCollector) RaiseModelUnbuildable(cause error) {
	message := fmt.Sprintf("the model cannot be built into a weighted graph: %s", cause)
	chained := fmt.Errorf("%w: %w", fgaerrors.ErrModelNotBuildable, cause)
	c.addScopedError(message, GraphModelUnbuildable, "", nil, nil, nil, scope{cause: chained})
}
