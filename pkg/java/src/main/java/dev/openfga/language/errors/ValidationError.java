package dev.openfga.language.errors;

import com.fasterxml.jackson.annotation.JsonValue;
import lombok.AllArgsConstructor;

@AllArgsConstructor
public enum ValidationError {
    AllowedTypesNotValidOnSchema1_0("allowed-type-not-valid-on-schema-1_0"),
    AssignableRelationsMustHaveType("assignable-relation-must-have-type"),
    ConditionNotDefined("condition-not-defined"),
    ConditionNotUsed("condition-not-used"),
    DuplicatedError("duplicated-error"),
    InvalidName("invalid-name"),
    InvalidRelationType("invalid-relation-type"),
    InvalidSchema("invalid-schema"),
    InvalidSyntax("invalid-syntax"),
    InvalidType("invalid-type"),
    MissingDefinition("missing-definition"),
    RelationNoEntrypoint("relation-no-entry-point"),
    RequireSchema1_0("allowed-type-schema-10"),
    ReservedRelationKeywords("reserved-relation-keywords"),
    ReservedTypeKeywords("reserved-type-keywords"),
    SchemaVersionRequired("schema-version-required"),
    SchemaVersionUnsupported("schema-version-unsupported"),
    SelfError("self-error"),
    TuplesetNotDirect("tupleuset-not-direct"),
    TypeRestrictionCannotHaveWildcardAndRelation("type-wildcard-relation");

    private final String value;

    @JsonValue
    public String getValue() {
        return value;
    }
}