package validation

import (
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

// scope is what a raise site knows that the collector cannot work out: which part of
// the model is at fault, and the enclosing type for the metadata.
type scope struct {
	// part names the part of the model at fault. The raise site builds it, because
	// the code alone does not say which part: duplicated-error is raised about a type
	// from one place and a relation from another, and invalid-name about all three.
	// The sentinel is filled in from the code's table entry, so at this point it
	// wraps nothing.
	part fgaerrors.ModelError

	// offendingType is the enclosing type a finding about another type was written
	// in, matching JS's wire field of the same name. Metadata only: none of the
	// scope types has a slot for it.
	offendingType string
}

// addError adds a finding that names no part of the model, so it is about the model as
// a whole. A raise site that names one, carries file or module metadata, or resolves
// its own column goes through addScopedError instead; none of the codes raised through
// here does any of those.
func (c *ErrorCollector) addError(message string, errorType ValidationErrorType, symbol string,
	lineIndex *int) {
	c.addScopedError(message, errorType, symbol, lineIndex, nil, nil, scope{
		part: &fgaerrors.ErrModel{},
	})
}

// addScopedError resolves where a finding points and records it. The code decides its
// severity and the sentinel it wraps; the raise site's scope decides which part of the
// model it names, and both the category and the metadata are read back off that.
func (c *ErrorCollector) addScopedError(message string, errorType ValidationErrorType, symbol string,
	lineIndex *int, meta *Meta, customResolver ErrorCustomResolver, errorScope scope) {
	line, column := c.position(symbol, lineIndex, customResolver)

	part := errorScope.part
	if part == nil {
		// A raise site that named nothing. Treat it as being about the model as a
		// whole, which is what a code with no scope means.
		part = &fgaerrors.ErrModel{}
	}

	entry := lookupErrorInfo(errorType)
	partScope := part.Scope()

	metadata := &ErrorMetadata{
		Symbol:        symbol,
		ErrorType:     errorType,
		OffendingType: errorScope.offendingType,
		Type:          partScope.ObjectType,
		Relation:      partScope.Relation,
		Condition:     partScope.Condition,
	}

	if meta != nil {
		// Module goes in the metadata, file on the error itself, matching the
		// JS implementation.
		metadata.Module = meta.Module
	}

	validationErr := &ValidationError{
		Message:  message,
		Severity: entry.Severity,
		Category: part.Kind(),
		Line:     line,
		Column:   column,
		Metadata: metadata,

		// A code missing from the table has no sentinel, so there is nothing for
		// errors.Is to match and this is nil. The category and metadata above still
		// report what the raise site named.
		Cause: fgaerrors.WithSentinel(part, entry.Cause),
	}

	if meta != nil {
		validationErr.File = meta.File
	}

	c.errors = append(c.errors, validationErr)
}

// position resolves the line and column a finding points at, both nil when the raise
// site gave no line or the line is outside the source.
func (c *ErrorCollector) position(symbol string, lineIndex *int,
	customResolver ErrorCustomResolver) (line, column *Range) {
	if lineIndex == nil || *lineIndex < 0 || *lineIndex >= len(c.lines) {
		return nil, nil
	}

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

	return line, column
}

// RaiseInvalidName raises an invalid name error.
func (c *ErrorCollector) RaiseInvalidName(symbol, clause string, typeName *string, lineIndex *int, meta *Meta) {
	var message string
	// A nil typeName means the offending name is a type rather than a relation on
	// one, which changes both the message and the scope of the finding.
	errorScope := scope{part: &fgaerrors.ErrObjectType{ObjectType: symbol}}

	if typeName != nil {
		message = fmt.Sprintf("relation '%s' of type '%s' does not match naming rule: '%s'.", symbol, *typeName, clause)
		errorScope = scope{part: &fgaerrors.ErrRelation{ObjectType: *typeName, Relation: symbol}}
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
		part: &fgaerrors.ErrCondition{Condition: symbol},
	})
}

// RaiseReservedTypeName raises a reserved type name error.
func (c *ErrorCollector) RaiseReservedTypeName(symbol string, lineIndex *int, meta *Meta) {
	message := "a type cannot be named 'self' or 'this'."
	c.addScopedError(message, ReservedTypeKeywords, symbol, lineIndex, meta, nil, scope{
		part: &fgaerrors.ErrObjectType{ObjectType: symbol},
	})
}

// RaiseReservedRelationName raises a reserved relation name error.
func (c *ErrorCollector) RaiseReservedRelationName(symbol, typeName string, lineIndex *int, meta *Meta) {
	message := "a relation cannot be named 'self' or 'this'."
	c.addScopedError(message, ReservedRelationKeywords, symbol, lineIndex, meta, nil, scope{
		part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: symbol},
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
		part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: relation},
	})
}

// RaiseDuplicateTypeName raises a duplicate type name error.
func (c *ErrorCollector) RaiseDuplicateTypeName(symbol string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("the type `%s` is a duplicate.", symbol)
	// A duplicate type is about the type, not a relation on it, so this overrides
	// DuplicatedError's relation-scoped default.
	c.addScopedError(message, DuplicatedError, symbol, lineIndex, meta, nil, scope{
		part: &fgaerrors.ErrObjectType{ObjectType: symbol},
	})
}

// RaiseDuplicateTypeRestriction raises a duplicate type restriction error.
func (c *ErrorCollector) RaiseDuplicateTypeRestriction(symbol, relationName, typeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("the type restriction `%s` is a duplicate in the relation `%s`.", symbol, relationName)
	c.addScopedError(message, DuplicatedError, symbol, lineIndex, meta, nil, scope{
		part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: relationName},
	})
}

// RaiseUndefinedType raises an error for undefined type references.
func (c *ErrorCollector) RaiseUndefinedType(typeName, relationName, parentTypeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("Type '%s' is not defined (referenced in relation '%s' of type '%s')", typeName, relationName, parentTypeName)
	// The undefined type is the subject; parentTypeName is only where it was
	// referenced from, so the scope names the type that does not exist.
	c.addScopedError(message, UndefinedType, typeName, lineIndex, meta, nil, scope{
		part: &fgaerrors.ErrObjectType{ObjectType: typeName},
	})
}

// RaiseUndefinedRelation raises an error for undefined relation references.
func (c *ErrorCollector) RaiseUndefinedRelation(relationName, typeName, parentRelation, parentTypeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("Relation '%s' is not defined on type '%s' (referenced in relation '%s' of type '%s')", relationName, typeName, parentRelation, parentTypeName)
	c.addScopedError(message, UndefinedRelation, relationName, lineIndex, meta, nil, scope{
		part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: relationName},
	})
}

// RaiseDuplicateType raises a duplicate type error in relation.
func (c *ErrorCollector) RaiseDuplicateType(symbol, relationName, typeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("the partial relation definition `%s` is a duplicate in the relation `%s`.",
		symbol, relationName)
	c.addScopedError(message, DuplicatedError, symbol, lineIndex, meta, nil, scope{
		part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: relationName},
	})
}

// RaiseDuplicateRelationshipDefinition raises a duplicate relationship definition error.
func (c *ErrorCollector) RaiseDuplicateRelationshipDefinition(symbol string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("the relation '%s' is defined more than once.", symbol)
	c.addScopedError(message, DuplicatedError, symbol, lineIndex, meta, nil, scope{
		part: &fgaerrors.ErrRelation{Relation: symbol},
	})
}

// RaiseNoEntryPointLoop raises an error for impossible relation with potential loop.
func (c *ErrorCollector) RaiseNoEntryPointLoop(symbol, typeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("`%s` is an impossible relation for `%s` (potential loop).", symbol, typeName)
	c.addScopedError(message, RelationNoEntrypoint, symbol, lineIndex, meta, nil, scope{
		part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: symbol},
	})
}

// RaiseNoEntryPoint raises an error for impossible relation without entry point.
func (c *ErrorCollector) RaiseNoEntryPoint(symbol, typeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("`%s` is an impossible relation for `%s` (no entrypoint).", symbol, typeName)
	c.addScopedError(message, RelationNoEntrypoint, symbol, lineIndex, meta, nil, scope{
		part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: symbol},
	})
}

// RaiseInvalidRelationOnTupleset raises an error for invalid relation on tupleset.
func (c *ErrorCollector) RaiseInvalidRelationOnTupleset(symbol, typeName, typeDef, relationName,
	offendingRelation, parent string, lineIndex *int, meta *Meta) {
	message := fmt.Sprintf("the `%s` relation definition on type `%s` is not valid: `%s` does not exist on `%s`, which is of type `%s`.",
		offendingRelation, typeDef, offendingRelation, parent, typeName)
	c.addScopedError(message, InvalidRelationOnTupleset, symbol, lineIndex, meta, nil, scope{
		part: &fgaerrors.ErrRelation{ObjectType: typeDef, Relation: relationName},
	})
}

// RaiseInvalidTypeRelation raises an error for invalid type relation.
func (c *ErrorCollector) RaiseInvalidTypeRelation(symbol, typeName, relationName, offendingRelation,
	offendingType string, lineIndex *int, meta *Meta) {
	message := fmt.Sprintf("`%s` is not a valid relation for `%s`.", offendingRelation, typeName)
	c.addScopedError(message, InvalidRelationType, symbol, lineIndex, meta, nil, scope{
		part:          &fgaerrors.ErrRelation{ObjectType: typeName, Relation: relationName},
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
		part: &fgaerrors.ErrObjectType{ObjectType: symbol},
	})
}

// RaiseAssignableRelationMustHaveTypes raises an error for assignable relations without types.
func (c *ErrorCollector) RaiseAssignableRelationMustHaveTypes(symbol string, lineIndex *int) {
	message := fmt.Sprintf("the assignable relation '%s' must have at least one assignable type.", symbol)
	c.addScopedError(message, AssignableRelationsMustHaveType, symbol, lineIndex, nil, nil, scope{
		part: &fgaerrors.ErrRelation{Relation: symbol},
	})
}

// RaiseAssignableTypeWildcardRelation raises an error for wildcard with relation.
func (c *ErrorCollector) RaiseAssignableTypeWildcardRelation(symbol, typeName, relation string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("the type restriction '%s' on relation '%s' of type '%s' is not allowed to have both a wildcard and a relation.",
		symbol, relation, typeName)
	c.addScopedError(message, TypeRestrictionCannotHaveWildcardAndRelation, symbol, lineIndex, meta, nil, scope{
		part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: relation},
	})
}

// RaiseInvalidRelationError reports a rewrite that names a relation the type does
// not define. The message names the missing relation only, as the reference's does;
// it does not list the relations that do exist.
func (c *ErrorCollector) RaiseInvalidRelationError(symbol, typeName, relation string,
	lineIndex *int, meta *Meta) {
	message := fmt.Sprintf("the relation `%s` does not exist.", symbol)
	c.addScopedError(message, MissingDefinition, symbol, lineIndex, meta, nil, scope{
		part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: relation},
	})
}

// RaiseInvalidSchemaVersion raises an error for a schema version that was never
// valid (e.g. "0.9", "2.0"). This is distinct from a version that is recognized
// but no longer supported (see RaiseSchemaVersionUnsupported).
func (c *ErrorCollector) RaiseInvalidSchemaVersion(symbol string, lineIndex *int) {
	message := fmt.Sprintf("invalid schema %s", symbol)
	c.addError(message, InvalidSchema, symbol, lineIndex)
}

// RaiseSchemaVersionUnsupported raises an error for a recognized but retired
// schema version (e.g. "1.0").
func (c *ErrorCollector) RaiseSchemaVersionUnsupported(symbol string, lineIndex *int) {
	message := "schema version no longer supported"
	c.addError(message, SchemaVersionUnsupported, symbol, lineIndex)
}

// RaiseSchemaVersionRequired raises an error for missing schema version.
func (c *ErrorCollector) RaiseSchemaVersionRequired(symbol string, lineIndex *int) {
	message := "schema version required"
	c.addError(message, SchemaVersionRequired, symbol, lineIndex)
}

// RaiseMaximumOneDirectRelationship raises an error for multiple direct relationships.
func (c *ErrorCollector) RaiseMaximumOneDirectRelationship(symbol string, lineIndex *int) {
	message := fmt.Sprintf("the relation '%s' can have at most one direct relationship.", symbol)
	c.addScopedError(message, DuplicatedError, symbol, lineIndex, nil, nil, scope{
		part: &fgaerrors.ErrRelation{Relation: symbol},
	})
}

// RaiseInvalidConditionNameInParameter raises an error for invalid condition names.
func (c *ErrorCollector) RaiseInvalidConditionNameInParameter(symbol, typeName, relationName, conditionName string,
	meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("`%s` is not a defined condition in the model.", conditionName)
	// Scoped to the relation the condition is applied to, not the condition's own
	// definition: the condition does not exist to have a definition.
	c.addScopedError(message, ConditionNotDefined, symbol, lineIndex, meta, nil, scope{
		part: &fgaerrors.ErrRelationCondition{ObjectType: typeName, Relation: relationName, Condition: conditionName},
	})
}

// RaiseUnusedCondition raises an error for unused conditions.
func (c *ErrorCollector) RaiseUnusedCondition(symbol string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("`%s` condition is not used in the model.", symbol)
	c.addScopedError(message, ConditionNotUsed, symbol, lineIndex, meta, nil, scope{
		part: &fgaerrors.ErrCondition{Condition: symbol},
	})
}

// RaiseDifferentNestedConditionName raises an error for a condition whose nested
// name property differs from its map key. The message mirrors the reference.
func (c *ErrorCollector) RaiseDifferentNestedConditionName(condition, nestedConditionName string) {
	message := fmt.Sprintf("condition key is `%s` but nested name property is %s", condition, nestedConditionName)
	c.addScopedError(message, DifferentNestedConditionName, nestedConditionName, nil, nil, nil, scope{
		part: &fgaerrors.ErrCondition{Condition: condition},
	})
}

// RaiseMultipleModulesInSingleFile raises an error for multiple modules in single
// file. The modules are listed in the order the model declares them, and the message
// mirrors the reference.
func (c *ErrorCollector) RaiseMultipleModulesInSingleFile(file string, modules []string) {
	moduleList := strings.Join(modules, ", ")
	message := fmt.Sprintf("file %s would contain multiple module definitions (%s) when transforming to DSL. "+
		"Only one module can be defined per file.", file, moduleList)
	c.addError(message, MultipleModulesInFile, file, nil)
}

// Complex operation validation error methods

// RaiseRedundantUnionMember raises an error for redundant members in union operations.
func (c *ErrorCollector) RaiseRedundantUnionMember(operation, relationName, typeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("Redundant operation '%s' found in union for relation '%s' of type '%s'", operation, relationName, typeName)
	c.addScopedError(message, DuplicatedError, operation, lineIndex, meta, nil, scope{
		part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: relationName},
	})
}

// RaiseImpossibleIntersection raises an error for intersection operations that cannot succeed.
func (c *ErrorCollector) RaiseImpossibleIntersection(relationName, typeName string, conflictingTypes []string, meta *Meta, lineIndex *int) {
	typeList := strings.Join(conflictingTypes, ", ")
	message := fmt.Sprintf("Impossible intersection in relation '%s' of type '%s': conflicting types [%s]", relationName, typeName, typeList)
	c.addScopedError(message, InvalidRelationType, relationName, lineIndex, meta, nil, scope{
		part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: relationName},
	})
}

// RaiseEmptyDifference raises an error for difference operations that result in empty sets.
func (c *ErrorCollector) RaiseEmptyDifference(relationName, typeName, operation string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("Empty difference operation in relation '%s' of type '%s': subtracting '%s' from itself", relationName, typeName, operation)
	c.addScopedError(message, RelationNoEntrypoint, relationName, lineIndex, meta, nil, scope{
		part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: relationName},
	})
}
