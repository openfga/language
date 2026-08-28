package validation

import (
	"fmt"
	"strings"

	fgaerrors "github.com/openfga/language/pkg/go/errors"
)

// newInvalidNameError reports a name that breaks a naming rule. A nil typeName means the
// offending name is a type rather than a relation on one, which changes the message and
// the scope.
func newInvalidNameError(lines []string, symbol, clause string, typeName *string, lineIndex *int, meta *Meta) *ValidationError {
	var message string
	errorScope := scope{part: &fgaerrors.ErrObjectType{ObjectType: symbol}}

	if typeName != nil {
		message = fmt.Sprintf("relation '%s' of type '%s' does not match naming rule: '%s'.", symbol, *typeName, clause)
		errorScope = scope{part: &fgaerrors.ErrRelation{ObjectType: *typeName, Relation: symbol}}
	} else {
		message = fmt.Sprintf("type '%s' does not match naming rule: '%s'.", symbol, clause)
	}

	line, column := resolvePosition(lines, symbol, lineIndex, nil)
	return newValidationError(message, InvalidName, symbol, line, column, errorScope, meta)
}

// newInvalidConditionNameError reports a condition name that breaks a naming rule,
// scoped to the condition rather than a type or relation.
func newInvalidConditionNameError(lines []string, symbol, clause string, lineIndex *int, meta *Meta) *ValidationError {
	message := fmt.Sprintf("condition '%s' does not match naming rule: '%s'.", symbol, clause)
	line, column := resolvePosition(lines, symbol, lineIndex, nil)
	return newValidationError(message, InvalidName, symbol, line, column, scope{part: &fgaerrors.ErrCondition{Condition: symbol}}, meta)
}

// newReservedTypeNameError reports a type named with a reserved keyword.
func newReservedTypeNameError(lines []string, symbol string, lineIndex *int, meta *Meta) *ValidationError {
	line, column := resolvePosition(lines, symbol, lineIndex, nil)
	return newValidationError("a type cannot be named 'self' or 'this'.", ReservedTypeKeywords, symbol, line, column, scope{part: &fgaerrors.ErrObjectType{ObjectType: symbol}}, meta)
}

// newReservedRelationNameError reports a relation named with a reserved keyword.
func newReservedRelationNameError(lines []string, symbol, typeName string, lineIndex *int, meta *Meta) *ValidationError {
	line, column := resolvePosition(lines, symbol, lineIndex, nil)
	return newValidationError("a relation cannot be named 'self' or 'this'.", ReservedRelationKeywords, symbol, line, column, scope{part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: symbol}}, meta)
}

// newTupleUsersetRequiresDirectError reports a tuple-to-userset that is not direct. Its
// column is resolved past the `from` keyword so it marks the offending relation.
func newTupleUsersetRequiresDirectError(lines []string, symbol, typeName, relation string, meta *Meta, lineIndex *int) *ValidationError {
	message := fmt.Sprintf("`%s` relation used inside from allows only direct relation.", symbol)

	customResolver := func(wordIdx int, rawLine, value string) int {
		clauseStartsAt := strings.Index(rawLine, "from") + len("from")
		if clauseStartsAt >= len("from") {
			wordIdx = clauseStartsAt + strings.Index(rawLine[clauseStartsAt:], value)
		}
		return wordIdx
	}

	line, column := resolvePosition(lines, symbol, lineIndex, customResolver)
	return newValidationError(message, TuplesetNotDirect, symbol, line, column, scope{part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: relation}}, meta)
}

// newDuplicateTypeNameError reports a duplicated type. It is about the type, not a
// relation on it, so it overrides DuplicatedError's relation-scoped default.
func newDuplicateTypeNameError(lines []string, symbol string, meta *Meta, lineIndex *int) *ValidationError {
	message := fmt.Sprintf("the type `%s` is a duplicate.", symbol)
	line, column := resolvePosition(lines, symbol, lineIndex, nil)
	return newValidationError(message, DuplicatedError, symbol, line, column, scope{part: &fgaerrors.ErrObjectType{ObjectType: symbol}}, meta)
}

// newDuplicateTypeRestrictionError reports a duplicated type restriction on a relation.
func newDuplicateTypeRestrictionError(lines []string, symbol, relationName, typeName string, meta *Meta, lineIndex *int) *ValidationError {
	message := fmt.Sprintf("the type restriction `%s` is a duplicate in the relation `%s`.", symbol, relationName)
	line, column := resolvePosition(lines, symbol, lineIndex, nil)
	return newValidationError(message, DuplicatedError, symbol, line, column, scope{part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: relationName}}, meta)
}

// newUndefinedTypeError reports a reference to a type that does not exist. The scope
// names the type that is missing, not the relation it was referenced from.
func newUndefinedTypeError(lines []string, typeName, relationName, parentTypeName string, meta *Meta, lineIndex *int) *ValidationError {
	message := fmt.Sprintf("Type '%s' is not defined (referenced in relation '%s' of type '%s')", typeName, relationName, parentTypeName)
	line, column := resolvePosition(lines, typeName, lineIndex, nil)
	return newValidationError(message, UndefinedType, typeName, line, column, scope{part: &fgaerrors.ErrObjectType{ObjectType: typeName}}, meta)
}

// newUndefinedRelationError reports a reference to a relation that does not exist on its
// type.
func newUndefinedRelationError(lines []string, relationName, typeName, parentRelation, parentTypeName string, meta *Meta, lineIndex *int) *ValidationError {
	message := fmt.Sprintf("Relation '%s' is not defined on type '%s' (referenced in relation '%s' of type '%s')", relationName, typeName, parentRelation, parentTypeName)
	line, column := resolvePosition(lines, relationName, lineIndex, nil)
	return newValidationError(message, UndefinedRelation, relationName, line, column, scope{part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: relationName}}, meta)
}

// newDuplicateTypeError reports a duplicated partial relation definition.
func newDuplicateTypeError(lines []string, symbol, relationName, typeName string, meta *Meta, lineIndex *int) *ValidationError {
	message := fmt.Sprintf("the partial relation definition `%s` is a duplicate in the relation `%s`.",
		symbol, relationName)
	line, column := resolvePosition(lines, symbol, lineIndex, nil)
	return newValidationError(message, DuplicatedError, symbol, line, column, scope{part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: relationName}}, meta)
}

// newDuplicateRelationshipDefinitionError reports a relation defined more than once.
func newDuplicateRelationshipDefinitionError(lines []string, symbol string, meta *Meta, lineIndex *int) *ValidationError {
	message := fmt.Sprintf("the relation '%s' is defined more than once.", symbol)
	line, column := resolvePosition(lines, symbol, lineIndex, nil)
	return newValidationError(message, DuplicatedError, symbol, line, column, scope{part: &fgaerrors.ErrRelation{Relation: symbol}}, meta)
}

// newNoEntryPointLoopError reports an impossible relation with a potential loop.
func newNoEntryPointLoopError(lines []string, symbol, typeName string, meta *Meta, lineIndex *int) *ValidationError {
	message := fmt.Sprintf("`%s` is an impossible relation for `%s` (potential loop).", symbol, typeName)
	line, column := resolvePosition(lines, symbol, lineIndex, nil)
	return newValidationError(message, RelationNoEntrypoint, symbol, line, column, scope{part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: symbol}}, meta)
}

// newNoEntryPointError reports an impossible relation with no entry point.
func newNoEntryPointError(lines []string, symbol, typeName string, meta *Meta, lineIndex *int) *ValidationError {
	message := fmt.Sprintf("`%s` is an impossible relation for `%s` (no entrypoint).", symbol, typeName)
	line, column := resolvePosition(lines, symbol, lineIndex, nil)
	return newValidationError(message, RelationNoEntrypoint, symbol, line, column, scope{part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: symbol}}, meta)
}

// newInvalidRelationOnTuplesetError reports a tupleset relation whose target does not
// exist on the referenced type.
func newInvalidRelationOnTuplesetError(lines []string, symbol, typeName, typeDef, relationName,
	offendingRelation, parent string, lineIndex *int, meta *Meta) *ValidationError {
	message := fmt.Sprintf("the `%s` relation definition on type `%s` is not valid: `%s` does not exist on `%s`, which is of type `%s`.",
		offendingRelation, typeDef, offendingRelation, parent, typeName)
	line, column := resolvePosition(lines, symbol, lineIndex, nil)
	return newValidationError(message, InvalidRelationOnTupleset, symbol, line, column, scope{part: &fgaerrors.ErrRelation{ObjectType: typeDef, Relation: relationName}}, meta)
}

// newInvalidTypeRelationError reports a relation reference that is not valid for a type.
// Its offendingType argument is the enclosing type the reference was written in, kept as
// metadata.
func newInvalidTypeRelationError(lines []string, symbol, typeName, relationName, offendingRelation,
	offendingType string, lineIndex *int, meta *Meta) *ValidationError {
	message := fmt.Sprintf("`%s` is not a valid relation for `%s`.", offendingRelation, typeName)
	line, column := resolvePosition(lines, symbol, lineIndex, nil)
	return newValidationError(message, InvalidRelationType, symbol, line, column, scope{
		part:          &fgaerrors.ErrRelation{ObjectType: typeName, Relation: relationName},
		offendingType: offendingType,
	}, meta)
}

// newInvalidTypeError reports an invalid type in an assignable-types list. Its column is
// resolved to the value side of the colon so it marks the type, not a relation key that
// shares its name.
func newInvalidTypeError(lines []string, symbol string, meta *Meta, lineIndex *int) *ValidationError {
	message := fmt.Sprintf("`%s` is not a valid type.", symbol)
	resolver := func(_ int, rawLine, sym string) int {
		colon := strings.Index(rawLine, ":")
		if colon < 0 {
			return wordIndex(rawLine, sym)
		}
		value := rawLine[colon+1:]
		idx := wordIndex(value, sym)
		return colon + 1 + idx
	}
	line, column := resolvePosition(lines, symbol, lineIndex, resolver)
	return newValidationError(message, InvalidType, symbol, line, column, scope{part: &fgaerrors.ErrObjectType{ObjectType: symbol}}, meta)
}

// newAssignableRelationMustHaveTypesError reports an assignable relation with no
// assignable type.
func newAssignableRelationMustHaveTypesError(lines []string, symbol string, lineIndex *int) *ValidationError {
	message := fmt.Sprintf("the assignable relation '%s' must have at least one assignable type.", symbol)
	line, column := resolvePosition(lines, symbol, lineIndex, nil)
	return newValidationError(message, AssignableRelationsMustHaveType, symbol, line, column, scope{part: &fgaerrors.ErrRelation{Relation: symbol}}, nil)
}

// newAssignableTypeWildcardRelationError reports a type restriction that carries both a
// wildcard and a relation.
func newAssignableTypeWildcardRelationError(lines []string, symbol, typeName, relation string, meta *Meta, lineIndex *int) *ValidationError {
	message := fmt.Sprintf("the type restriction '%s' on relation '%s' of type '%s' is not allowed to have both a wildcard and a relation.",
		symbol, relation, typeName)
	line, column := resolvePosition(lines, symbol, lineIndex, nil)
	return newValidationError(message, TypeRestrictionCannotHaveWildcardAndRelation, symbol, line, column, scope{part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: relation}}, meta)
}

// newInvalidRelationError reports a rewrite that names a relation the type does not
// define. The message names the missing relation only, as the reference's does.
func newInvalidRelationError(lines []string, symbol, typeName, relation string,
	lineIndex *int, meta *Meta) *ValidationError {
	message := fmt.Sprintf("the relation `%s` does not exist.", symbol)
	line, column := resolvePosition(lines, symbol, lineIndex, nil)
	return newValidationError(message, MissingDefinition, symbol, line, column, scope{part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: relation}}, meta)
}

// newInvalidSchemaVersionError reports a schema version that was never valid (e.g.
// "0.9", "2.0"), as distinct from one that is recognized but no longer supported.
func newInvalidSchemaVersionError(lines []string, symbol string, lineIndex *int) *ValidationError {
	line, column := resolvePosition(lines, symbol, lineIndex, nil)
	return newValidationError(fmt.Sprintf("invalid schema %s", symbol), InvalidSchema, symbol, line, column, scope{}, nil)
}

// newSchemaVersionUnsupportedError reports a recognized but retired schema version
// (e.g. "1.0").
func newSchemaVersionUnsupportedError(lines []string, symbol string, lineIndex *int) *ValidationError {
	line, column := resolvePosition(lines, symbol, lineIndex, nil)
	return newValidationError("schema version no longer supported", SchemaVersionUnsupported, symbol, line, column, scope{}, nil)
}

// newSchemaVersionRequiredError reports a model with no schema version. It names no part
// of the model, so it is about the model as a whole.
func newSchemaVersionRequiredError(lines []string, lineIndex *int) *ValidationError {
	line, column := resolvePosition(lines, "", lineIndex, nil)
	return newValidationError("schema version required", SchemaVersionRequired, "", line, column, scope{}, nil)
}

// newMaximumOneDirectRelationshipError reports a relation with more than one direct
// relationship.
func newMaximumOneDirectRelationshipError(lines []string, symbol string, lineIndex *int) *ValidationError {
	message := fmt.Sprintf("the relation '%s' can have at most one direct relationship.", symbol)
	line, column := resolvePosition(lines, symbol, lineIndex, nil)
	return newValidationError(message, DuplicatedError, symbol, line, column, scope{part: &fgaerrors.ErrRelation{Relation: symbol}}, nil)
}

// newInvalidConditionNameInParameterError reports a reference to a condition that is not
// defined. It is scoped to the relation the condition is applied to, since the condition
// has no definition to point at.
func newInvalidConditionNameInParameterError(lines []string, symbol, typeName, relationName, conditionName string,
	meta *Meta, lineIndex *int) *ValidationError {
	message := fmt.Sprintf("`%s` is not a defined condition in the model.", conditionName)
	line, column := resolvePosition(lines, symbol, lineIndex, nil)
	return newValidationError(message, ConditionNotDefined, symbol, line, column, scope{part: &fgaerrors.ErrRelationCondition{ObjectType: typeName, Relation: relationName, Condition: conditionName}}, meta)
}

// newUnusedConditionError reports a condition defined but never referenced.
func newUnusedConditionError(lines []string, symbol string, meta *Meta, lineIndex *int) *ValidationError {
	message := fmt.Sprintf("`%s` condition is not used in the model.", symbol)
	line, column := resolvePosition(lines, symbol, lineIndex, nil)
	return newValidationError(message, ConditionNotUsed, symbol, line, column, scope{part: &fgaerrors.ErrCondition{Condition: symbol}}, meta)
}

// newDifferentNestedConditionNameError reports a condition whose nested name property
// differs from its map key. It carries no position, matching the reference.
func newDifferentNestedConditionNameError(condition, nestedConditionName string) *ValidationError {
	message := fmt.Sprintf("condition key is `%s` but nested name property is %s", condition, nestedConditionName)
	return newValidationError(message, DifferentNestedConditionName, nestedConditionName, nil, nil, scope{part: &fgaerrors.ErrCondition{Condition: condition}}, nil)
}

// newMultipleModulesInSingleFileError reports a file that would contain more than one
// module. It names no part of the model, so it is about the model as a whole.
func newMultipleModulesInSingleFileError(file string, modules []string) *ValidationError {
	moduleList := strings.Join(modules, ", ")
	message := fmt.Sprintf("file %s would contain multiple module definitions (%s) when transforming to DSL. "+
		"Only one module can be defined per file.", file, moduleList)
	return newValidationError(message, MultipleModulesInFile, file, nil, nil, scope{}, nil)
}

// newRedundantUnionMemberError reports a redundant member in a union operation.
func newRedundantUnionMemberError(lines []string, operation, relationName, typeName string, meta *Meta, lineIndex *int) *ValidationError {
	message := fmt.Sprintf("Redundant operation '%s' found in union for relation '%s' of type '%s'", operation, relationName, typeName)
	line, column := resolvePosition(lines, operation, lineIndex, nil)
	return newValidationError(message, DuplicatedError, operation, line, column, scope{part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: relationName}}, meta)
}

// newImpossibleIntersectionError reports an intersection operation that cannot succeed.
func newImpossibleIntersectionError(lines []string, relationName, typeName string, conflictingTypes []string, meta *Meta, lineIndex *int) *ValidationError {
	typeList := strings.Join(conflictingTypes, ", ")
	message := fmt.Sprintf("Impossible intersection in relation '%s' of type '%s': conflicting types [%s]", relationName, typeName, typeList)
	line, column := resolvePosition(lines, relationName, lineIndex, nil)
	return newValidationError(message, InvalidRelationType, relationName, line, column, scope{part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: relationName}}, meta)
}

// newEmptyDifferenceError reports a difference operation that results in an empty set.
func newEmptyDifferenceError(lines []string, relationName, typeName, operation string, meta *Meta, lineIndex *int) *ValidationError {
	message := fmt.Sprintf("Empty difference operation in relation '%s' of type '%s': subtracting '%s' from itself", relationName, typeName, operation)
	line, column := resolvePosition(lines, relationName, lineIndex, nil)
	return newValidationError(message, RelationNoEntrypoint, relationName, line, column, scope{part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: relationName}}, meta)
}

// newInvalidWildcardUsageError reports a wildcard used where it is not allowed. The
// wildcard is written in a relation of parentTypeName; typeName is the restriction it
// appears in, which the symbol already records.
func newInvalidWildcardUsageError(lines []string, typeName, relationName, parentTypeName, reason string, meta *Meta, lineIndex *int) *ValidationError {
	message := fmt.Sprintf("Invalid wildcard usage for type '%s' in relation '%s' of type '%s': %s",
		typeName, relationName, parentTypeName, reason)
	line, column := resolvePosition(lines, typeName, lineIndex, nil)
	return newValidationError(message, InvalidWildcardError, typeName, line, column, scope{part: &fgaerrors.ErrRelation{ObjectType: parentTypeName, Relation: relationName}}, meta)
}

// newTuplesetNotDirectError reports a tupleset relation that does not allow direct
// assignment.
func newTuplesetNotDirectError(lines []string, tuplesetRelation, typeName, parentRelation string, meta *Meta, lineIndex *int) *ValidationError {
	message := fmt.Sprintf("Tupleset relation '%s' on type '%s' must allow direct assignment (used in relation '%s')",
		tuplesetRelation, typeName, parentRelation)
	line, column := resolvePosition(lines, tuplesetRelation, lineIndex, nil)
	return newValidationError(message, TuplesetNotDirect, tuplesetRelation, line, column, scope{part: &fgaerrors.ErrRelation{ObjectType: typeName, Relation: tuplesetRelation}}, meta)
}
