package validation

import (
	"fmt"
	"strings"
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

// ErrorCollector collects validation errors during model validation
// This is equivalent to the JS ExceptionCollector class
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

// GetErrors returns all collected errors.
func (c *ErrorCollector) GetErrors() []*ValidationError {
	return c.errors
}

// HasErrors returns true if any errors have been collected.
func (c *ErrorCollector) HasErrors() bool {
	return len(c.errors) > 0
}

// Count returns the number of errors collected.
func (c *ErrorCollector) Count() int {
	return len(c.errors)
}

// addError is a helper to add an error to the collection.
func (c *ErrorCollector) addError(message string, errorType ValidationErrorType, symbol string,
	lineIndex *int, meta *Meta, customResolver ErrorCustomResolver) {
	var line *LineRange
	var column *ColumnRange

	// Calculate line and column positions if lineIndex is provided
	if lineIndex != nil && *lineIndex >= 0 && *lineIndex < len(c.lines) {
		line = &LineRange{Start: *lineIndex, End: *lineIndex}

		// Find symbol position in line for column calculation, matching on word
		// boundaries as the reference does.
		rawLine := c.lines[*lineIndex]
		symbolPos := wordIndex(rawLine, symbol)

		if customResolver != nil {
			symbolPos = customResolver(symbolPos, rawLine, symbol)
		}

		if symbolPos >= 0 {
			column = &ColumnRange{
				Start: symbolPos,
				End:   symbolPos + len(symbol),
			}
		}
	}

	metadata := &ErrorMetadata{
		Symbol:    symbol,
		ErrorType: errorType,
	}

	if meta != nil {
		metadata.Module = meta.Module
		// Set file in both metadata and error for consistency with JS implementation
	}

	validationErr := &ValidationError{
		Message:  message,
		Line:     line,
		Column:   column,
		Metadata: metadata,
	}

	if meta != nil {
		validationErr.File = meta.File
	}

	c.errors = append(c.errors, validationErr)
}

// RaiseInvalidName raises an invalid name error.
func (c *ErrorCollector) RaiseInvalidName(symbol, clause string, typeName *string, lineIndex *int, meta *Meta) {
	var message string
	if typeName != nil {
		message = fmt.Sprintf("relation '%s' of type '%s' does not match naming rule: '%s'.", symbol, *typeName, clause)
	} else {
		message = fmt.Sprintf("type '%s' does not match naming rule: '%s'.", symbol, clause)
	}
	c.addError(message, InvalidName, symbol, lineIndex, meta, nil)
}

// RaiseReservedTypeName raises a reserved type name error.
func (c *ErrorCollector) RaiseReservedTypeName(symbol string, lineIndex *int, meta *Meta) {
	message := "a type cannot be named 'self' or 'this'."
	c.addError(message, ReservedTypeKeywords, symbol, lineIndex, meta, nil)
}

// RaiseReservedRelationName raises a reserved relation name error.
func (c *ErrorCollector) RaiseReservedRelationName(symbol string, lineIndex *int, meta *Meta) {
	message := "a relation cannot be named 'self' or 'this'."
	c.addError(message, ReservedRelationKeywords, symbol, lineIndex, meta, nil)
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

	c.addError(message, TuplesetNotDirect, symbol, lineIndex, meta, customResolver)
}

// RaiseDuplicateTypeName raises a duplicate type name error.
func (c *ErrorCollector) RaiseDuplicateTypeName(symbol string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("the type `%s` is a duplicate.", symbol)
	c.addError(message, DuplicatedError, symbol, lineIndex, meta, nil)
}

// RaiseDuplicateTypeRestriction raises a duplicate type restriction error.
func (c *ErrorCollector) RaiseDuplicateTypeRestriction(symbol, relationName, typeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("the type restriction `%s` is a duplicate in the relation `%s`.", symbol, relationName)
	c.addError(message, DuplicatedError, symbol, lineIndex, meta, nil)
}

// RaiseUndefinedType raises an error for undefined type references.
func (c *ErrorCollector) RaiseUndefinedType(typeName, relationName, parentTypeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("Type '%s' is not defined (referenced in relation '%s' of type '%s')", typeName, relationName, parentTypeName)
	c.addError(message, UndefinedType, typeName, lineIndex, meta, nil)
}

// RaiseUndefinedRelation raises an error for undefined relation references.
func (c *ErrorCollector) RaiseUndefinedRelation(relationName, typeName, parentRelation, parentTypeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("Relation '%s' is not defined on type '%s' (referenced in relation '%s' of type '%s')", relationName, typeName, parentRelation, parentTypeName)
	c.addError(message, UndefinedRelation, relationName, lineIndex, meta, nil)
}

// RaiseDuplicateType raises a duplicate type error in relation.
func (c *ErrorCollector) RaiseDuplicateType(symbol, relationName, typeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("the partial relation definition `%s` is a duplicate in the relation `%s`.",
		symbol, relationName)
	c.addError(message, DuplicatedError, symbol, lineIndex, meta, nil)
}

// RaiseDuplicateRelationshipDefinition raises a duplicate relationship definition error.
func (c *ErrorCollector) RaiseDuplicateRelationshipDefinition(symbol string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("the relation '%s' is defined more than once.", symbol)
	c.addError(message, DuplicatedError, symbol, lineIndex, meta, nil)
}

// RaiseNoEntryPointLoop raises an error for impossible relation with potential loop.
func (c *ErrorCollector) RaiseNoEntryPointLoop(symbol, typeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("`%s` is an impossible relation for `%s` (potential loop).", symbol, typeName)
	c.addError(message, RelationNoEntrypoint, symbol, lineIndex, meta, nil)
}

// RaiseNoEntryPoint raises an error for impossible relation without entry point.
func (c *ErrorCollector) RaiseNoEntryPoint(symbol, typeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("`%s` is an impossible relation for `%s` (no entrypoint).", symbol, typeName)
	c.addError(message, RelationNoEntrypoint, symbol, lineIndex, meta, nil)
}

// RaiseInvalidRelationOnTupleset raises an error for invalid relation on tupleset.
func (c *ErrorCollector) RaiseInvalidRelationOnTupleset(symbol, typeName, typeDef, relationName,
	offendingRelation, parent string, lineIndex *int, meta *Meta) {
	message := fmt.Sprintf("the `%s` relation definition on type `%s` is not valid: `%s` does not exist on `%s`, which is of type `%s`.",
		offendingRelation, typeDef, offendingRelation, parent, typeName)
	c.addError(message, InvalidRelationOnTupleset, symbol, lineIndex, meta, nil)
}

// RaiseInvalidTypeRelation raises an error for invalid type relation.
func (c *ErrorCollector) RaiseInvalidTypeRelation(symbol, typeName, relationName, offendingRelation,
	offendingType string, lineIndex *int, meta *Meta) {
	message := fmt.Sprintf("`%s` is not a valid relation for `%s`.", offendingRelation, typeName)
	c.addError(message, InvalidRelationType, symbol, lineIndex, meta, nil)
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
	c.addError(message, InvalidType, symbol, lineIndex, meta, resolver)
}

// RaiseAssignableRelationMustHaveTypes raises an error for assignable relations without types.
func (c *ErrorCollector) RaiseAssignableRelationMustHaveTypes(symbol string, lineIndex *int) {
	message := fmt.Sprintf("the assignable relation '%s' must have at least one assignable type.", symbol)
	c.addError(message, AssignableRelationsMustHaveType, symbol, lineIndex, nil, nil)
}

// RaiseAssignableTypeWildcardRelation raises an error for wildcard with relation.
func (c *ErrorCollector) RaiseAssignableTypeWildcardRelation(symbol, typeName, relation string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("the type restriction '%s' on relation '%s' of type '%s' is not allowed to have both a wildcard and a relation.",
		symbol, relation, typeName)
	c.addError(message, TypeRestrictionCannotHaveWildcardAndRelation, symbol, lineIndex, meta, nil)
}

// RaiseInvalidRelationError raises an error for invalid relation reference.
func (c *ErrorCollector) RaiseInvalidRelationError(symbol, typeName, relation string, validRelations []string,
	lineIndex *int, meta *Meta) {
	message := fmt.Sprintf("the relation `%s` does not exist.", symbol)
	c.addError(message, MissingDefinition, symbol, lineIndex, meta, nil)
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
	c.addError(message, DuplicatedError, symbol, lineIndex, nil, nil)
}

// RaiseInvalidConditionNameInParameter raises an error for invalid condition names.
func (c *ErrorCollector) RaiseInvalidConditionNameInParameter(symbol, typeName, relationName, conditionName string,
	meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("`%s` is not a defined condition in the model.", conditionName)
	c.addError(message, ConditionNotDefined, symbol, lineIndex, meta, nil)
}

// RaiseUnusedCondition raises an error for unused conditions.
func (c *ErrorCollector) RaiseUnusedCondition(symbol string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("`%s` condition is not used in the model.", symbol)
	c.addError(message, ConditionNotUsed, symbol, lineIndex, meta, nil)
}

// RaiseDifferentNestedConditionName raises an error for mismatched condition names.
func (c *ErrorCollector) RaiseDifferentNestedConditionName(condition, nestedConditionName string) {
	message := fmt.Sprintf("the '%s' condition has a different nested condition name ('%s').", condition, nestedConditionName)
	c.addError(message, DifferentNestedConditionName, condition, nil, nil, nil)
}

// RaiseMultipleModulesInSingleFile raises an error for multiple modules in single file.
func (c *ErrorCollector) RaiseMultipleModulesInSingleFile(file string, modules []string) {
	moduleList := strings.Join(modules, ", ")
	message := fmt.Sprintf("file '%s' contains multiple modules: %s.", file, moduleList)
	c.addError(message, MultipleModulesInFile, file, nil, nil, nil)
}

// Complex operation validation error methods

// RaiseRedundantUnionMember raises an error for redundant members in union operations.
func (c *ErrorCollector) RaiseRedundantUnionMember(operation, relationName, typeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("Redundant operation '%s' found in union for relation '%s' of type '%s'", operation, relationName, typeName)
	c.addError(message, DuplicatedError, operation, lineIndex, meta, nil)
}

// RaiseImpossibleIntersection raises an error for intersection operations that cannot succeed.
func (c *ErrorCollector) RaiseImpossibleIntersection(relationName, typeName string, conflictingTypes []string, meta *Meta, lineIndex *int) {
	typeList := strings.Join(conflictingTypes, ", ")
	message := fmt.Sprintf("Impossible intersection in relation '%s' of type '%s': conflicting types [%s]", relationName, typeName, typeList)
	c.addError(message, InvalidRelationType, relationName, lineIndex, meta, nil)
}

// RaiseEmptyDifference raises an error for difference operations that result in empty sets.
func (c *ErrorCollector) RaiseEmptyDifference(relationName, typeName, operation string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("Empty difference operation in relation '%s' of type '%s': subtracting '%s' from itself", relationName, typeName, operation)
	c.addError(message, RelationNoEntrypoint, relationName, lineIndex, meta, nil)
}
