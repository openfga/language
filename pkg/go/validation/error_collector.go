package validation

import (
	"fmt"
	"strings"
)

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

		// Find symbol position in line for column calculation
		rawLine := c.lines[*lineIndex]
		symbolPos := strings.Index(rawLine, symbol)

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

	error := &ValidationError{
		Message:  message,
		Line:     line,
		Column:   column,
		Metadata: metadata,
	}

	if meta != nil {
		error.File = meta.File
	}

	c.errors = append(c.errors, error)
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
	message := fmt.Sprintf("the type definition '%s' is defined more than once.", symbol)
	c.addError(message, DuplicatedError, symbol, lineIndex, meta, nil)
}

// RaiseDuplicateTypeRestriction raises a duplicate type restriction error.
func (c *ErrorCollector) RaiseDuplicateTypeRestriction(symbol, relationName, typeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("the type restriction '%s' in relation '%s' of type '%s' is defined more than once.",
		symbol, relationName, typeName)
	c.addError(message, DuplicatedError, symbol, lineIndex, meta, nil)
}

// RaiseUndefinedType raises an error for undefined type references.
func (ec *ErrorCollector) RaiseUndefinedType(typeName, relationName, parentTypeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("Type '%s' is not defined (referenced in relation '%s' of type '%s')", typeName, relationName, parentTypeName)
	ec.addError(message, UndefinedType, typeName, lineIndex, meta, nil)
}

// RaiseUndefinedRelation raises an error for undefined relation references.
func (ec *ErrorCollector) RaiseUndefinedRelation(relationName, typeName, parentRelation, parentTypeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("Relation '%s' is not defined on type '%s' (referenced in relation '%s' of type '%s')", relationName, typeName, parentRelation, parentTypeName)
	ec.addError(message, UndefinedRelation, relationName, lineIndex, meta, nil)
}

// RaiseDuplicateType raises a duplicate type error in relation.
func (c *ErrorCollector) RaiseDuplicateType(symbol, relationName, typeName string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("the type '%s' is defined more than once in relation '%s' of type '%s'.",
		symbol, relationName, typeName)
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
	message := fmt.Sprintf("`%s` is an impossible relation for `%s`.", symbol, typeName)
	c.addError(message, RelationNoEntrypoint, symbol, lineIndex, meta, nil)
}

// RaiseInvalidRelationOnTupleset raises an error for invalid relation on tupleset.
func (c *ErrorCollector) RaiseInvalidRelationOnTupleset(symbol, typeName, typeDef, relationName,
	offendingRelation, parent string, lineIndex *int, meta *Meta) {
	message := fmt.Sprintf("the relation '%s' is not valid on type '%s'.", offendingRelation, typeDef)
	c.addError(message, InvalidRelationOnTupleset, symbol, lineIndex, meta, nil)
}

// RaiseInvalidTypeRelation raises an error for invalid type relation.
func (c *ErrorCollector) RaiseInvalidTypeRelation(symbol, typeName, relationName, offendingRelation,
	offendingType string, lineIndex *int, meta *Meta) {
	message := fmt.Sprintf("the relation '%s' does not exist on type '%s'.", offendingRelation, offendingType)
	c.addError(message, InvalidRelationType, symbol, lineIndex, meta, nil)
}

// RaiseInvalidType raises an error for invalid type.
func (c *ErrorCollector) RaiseInvalidType(symbol, typeName, relation string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("type '%s' is not defined.", symbol)
	c.addError(message, InvalidType, symbol, lineIndex, meta, nil)
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

// RaiseInvalidSchemaVersion raises an error for invalid schema version.
func (c *ErrorCollector) RaiseInvalidSchemaVersion(symbol string, lineIndex *int) {
	message := fmt.Sprintf("the schema version '%s' is not supported.", symbol)
	c.addError(message, SchemaVersionUnsupported, symbol, lineIndex, nil, nil)
}

// RaiseSchemaVersionRequired raises an error for missing schema version.
func (c *ErrorCollector) RaiseSchemaVersionRequired(symbol string, lineIndex *int) {
	message := "a schema version is required in the model."
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
	message := fmt.Sprintf("condition parameter name '%s' is invalid.", conditionName)
	c.addError(message, ConditionNotDefined, symbol, lineIndex, meta, nil)
}

// RaiseUnusedCondition raises an error for unused conditions.
func (c *ErrorCollector) RaiseUnusedCondition(symbol string, meta *Meta, lineIndex *int) {
	message := fmt.Sprintf("condition '%s' is defined but not used.", symbol)
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
