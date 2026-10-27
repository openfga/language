package validation

import (
	"fmt"
	"strings"
)

// The functions here build one finding each: the message the shared test cases expect,
// the code, and which part of the model is at fault. The raise site stamps
// where it is: the position through at or one of its variants, and the file and
// module from the model's metadata. A finding built from a JSON model has no
// source text, so it goes without a position.

// invalidTypeName reports a type name that breaks the naming rule.
func invalidTypeName(name string) *Finding {
	return &Finding{
		Message:  fmt.Sprintf("type '%s' does not match naming rule: '%s'.", name, typeNameRule),
		Metadata: Metadata{Kind: InvalidName, Symbol: name, Type: name},
	}
}

// invalidRelationName reports a relation name that breaks the naming rule.
func invalidRelationName(name, typeName string) *Finding {
	return &Finding{
		Message:  fmt.Sprintf("relation '%s' of type '%s' does not match naming rule: '%s'.", name, typeName, relationNameRule),
		Metadata: Metadata{Kind: InvalidName, Symbol: name, Type: typeName, Relation: name},
	}
}

// invalidConditionName reports a condition name that breaks the naming rule.
func invalidConditionName(name string) *Finding {
	return &Finding{
		Message:  fmt.Sprintf("condition '%s' does not match naming rule: '%s'.", name, conditionNameRule),
		Metadata: Metadata{Kind: InvalidName, Symbol: name, Condition: name},
	}
}

// reservedTypeName reports a type named with a reserved keyword.
func reservedTypeName(name string) *Finding {
	return &Finding{
		Message:  "a type cannot be named 'self' or 'this'.",
		Metadata: Metadata{Kind: ReservedTypeKeywords, Symbol: name, Type: name},
	}
}

// reservedRelationName reports a relation named with a reserved keyword.
func reservedRelationName(name, typeName string) *Finding {
	return &Finding{
		Message:  "a relation cannot be named 'self' or 'this'.",
		Metadata: Metadata{Kind: ReservedRelationKeywords, Symbol: name, Type: typeName, Relation: name},
	}
}

// tupleUsersetRequiresDirect reports a tuple-to-userset whose tupleset relation
// is not a direct assignment. Stamped with atFromClause so the column marks the
// relation after the `from` keyword.
func tupleUsersetRequiresDirect(fromRelation, typeName, relation string) *Finding {
	return &Finding{
		Message:  fmt.Sprintf("`%s` relation used inside from allows only direct relation.", fromRelation),
		Metadata: Metadata{Kind: TuplesetNotDirect, Symbol: fromRelation, Type: typeName, Relation: relation},
	}
}

// duplicateTypeName reports a duplicated type.
func duplicateTypeName(name string) *Finding {
	return &Finding{
		Message:  fmt.Sprintf("the type `%s` is a duplicate.", name),
		Metadata: Metadata{Kind: DuplicatedError, Symbol: name, Type: name},
	}
}

// duplicateTypeRestriction reports a duplicated type restriction on a relation.
func duplicateTypeRestriction(restriction, relationName, typeName string) *Finding {
	return &Finding{
		Message:  fmt.Sprintf("the type restriction `%s` is a duplicate in the relation `%s`.", restriction, relationName),
		Metadata: Metadata{Kind: DuplicatedError, Symbol: restriction, Type: typeName, Relation: relationName},
	}
}

// duplicatePartialRelation reports a duplicated partial relation definition.
func duplicatePartialRelation(operand, relationName, typeName string) *Finding {
	return &Finding{
		Message:  fmt.Sprintf("the partial relation definition `%s` is a duplicate in the relation `%s`.", operand, relationName),
		Metadata: Metadata{Kind: DuplicatedError, Symbol: operand, Type: typeName, Relation: relationName},
	}
}

// noEntryPointLoop reports an impossible relation with a potential loop.
func noEntryPointLoop(relation, typeName string) *Finding {
	return &Finding{
		Message:  fmt.Sprintf("`%s` is an impossible relation for `%s` (potential loop).", relation, typeName),
		Metadata: Metadata{Kind: RelationNoEntrypoint, Symbol: relation, Type: typeName, Relation: relation},
	}
}

// noEntryPoint reports an impossible relation with no entry point.
func noEntryPoint(relation, typeName string) *Finding {
	return &Finding{
		Message:  fmt.Sprintf("`%s` is an impossible relation for `%s` (no entrypoint).", relation, typeName),
		Metadata: Metadata{Kind: RelationNoEntrypoint, Symbol: relation, Type: typeName, Relation: relation},
	}
}

// invalidRelationOnTupleset reports a tuple-to-userset whose computed relation
// does not exist on the type the tupleset relation is assignable to.
func invalidRelationOnTupleset(symbol, missingRelation, typeName, tuplesetRelation, targetType, relationName string) *Finding {
	return &Finding{
		Message: fmt.Sprintf("the `%s` relation definition on type `%s` is not valid: `%s` does not exist on `%s`, which is of type `%s`.",
			missingRelation, typeName, missingRelation, tuplesetRelation, targetType),
		Metadata: Metadata{Kind: InvalidRelationOnTupleset, Symbol: symbol, Type: typeName, Relation: relationName},
	}
}

// invalidTypeRelation reports a relation reference that is not valid for a
// type. The offendingType is the enclosing type the reference was written in.
func invalidTypeRelation(symbol, typeName, relationName, offendingRelation, offendingType string) *Finding {
	return &Finding{
		Message:  fmt.Sprintf("`%s` is not a valid relation for `%s`.", offendingRelation, typeName),
		Metadata: Metadata{Kind: InvalidRelationType, Symbol: symbol, Type: typeName, Relation: relationName, OffendingType: offendingType},
	}
}

// invalidType reports an invalid type in an assignable-types list. Stamped with
// atRestriction so the column marks the type, not a relation key sharing its
// name.
func invalidType(typeName string) *Finding {
	return &Finding{
		Message:  fmt.Sprintf("`%s` is not a valid type.", typeName),
		Metadata: Metadata{Kind: InvalidType, Symbol: typeName, Type: typeName},
	}
}

// missingRelation reports a rewrite that names a relation the type does not
// define.
func missingRelation(name, typeName, relation string) *Finding {
	return &Finding{
		Message:  fmt.Sprintf("the relation `%s` does not exist.", name),
		Metadata: Metadata{Kind: MissingDefinition, Symbol: name, Type: typeName, Relation: relation},
	}
}

// invalidSchemaVersion reports a schema version that was never valid (e.g.
// "0.9", "2.0"), as distinct from one that is recognized but no longer
// supported. Its code is InvalidSchema, as in pkg/js; InvalidSchemaVersion is
// never emitted.
func invalidSchemaVersion(version string) *Finding {
	return &Finding{
		Message:  "invalid schema " + version,
		Metadata: Metadata{Kind: InvalidSchema, Symbol: version},
	}
}

// schemaVersionUnsupported reports a recognized but retired schema version
// (e.g. "1.0").
func schemaVersionUnsupported(version string) *Finding {
	return &Finding{
		Message:  "schema version no longer supported",
		Metadata: Metadata{Kind: SchemaVersionUnsupported, Symbol: version},
	}
}

// schemaVersionRequired reports a model with no schema version.
func schemaVersionRequired() *Finding {
	return &Finding{
		Message:  "schema version required",
		Metadata: Metadata{Kind: SchemaVersionRequired},
	}
}

// conditionNotDefined reports a reference to a condition the model does not
// define.
func conditionNotDefined(condition, typeName, relationName string) *Finding {
	return &Finding{
		Message:  fmt.Sprintf("`%s` is not a defined condition in the model.", condition),
		Metadata: Metadata{Kind: ConditionNotDefined, Symbol: condition, Type: typeName, Relation: relationName, Condition: condition},
	}
}

// unusedCondition reports a condition defined but never referenced.
func unusedCondition(condition string) *Finding {
	return &Finding{
		Message:  fmt.Sprintf("`%s` condition is not used in the model.", condition),
		Metadata: Metadata{Kind: ConditionNotUsed, Symbol: condition, Condition: condition},
	}
}

// differentNestedConditionName reports a condition whose nested name property
// differs from its map key. It carries no position, matching the reference.
func differentNestedConditionName(conditionKey, nestedName string) *Finding {
	return &Finding{
		Message:  fmt.Sprintf("condition key is `%s` but nested name property is %s", conditionKey, nestedName),
		Metadata: Metadata{Kind: DifferentNestedConditionName, Symbol: nestedName, Condition: conditionKey},
	}
}

// multipleModulesInSingleFile reports a file that would contain more than one
// module. It carries no position: it is about the file, not a line in it.
func multipleModulesInSingleFile(file string, modules []string) *Finding {
	return &Finding{
		Message: fmt.Sprintf("file %s would contain multiple module definitions (%s) when transforming to DSL. "+
			"Only one module can be defined per file.", file, strings.Join(modules, ", ")),
		Metadata: Metadata{Kind: MultipleModulesInFile, Symbol: file},
	}
}
