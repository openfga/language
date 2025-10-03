using System.Text.Json.Serialization;

namespace OpenFga.Language.Validation;

public enum ValidationError
{
    [JsonPropertyName("allowed-type-not-valid-on-schema-1_0")]
    AllowedTypesNotValidOnSchema1_0,
    
    [JsonPropertyName("assignable-relation-must-have-type")]
    AssignableRelationsMustHaveType,
    
    [JsonPropertyName("condition-not-defined")]
    ConditionNotDefined,
    
    [JsonPropertyName("condition-not-used")]
    ConditionNotUsed,
    
    [JsonPropertyName("duplicated-error")]
    DuplicatedError,
    
    [JsonPropertyName("invalid-name")]
    InvalidName,
    
    [JsonPropertyName("invalid-relation-type")]
    InvalidRelationType,
    
    [JsonPropertyName("invalid-schema")]
    InvalidSchema,
    
    [JsonPropertyName("invalid-syntax")]
    InvalidSyntax,
    
    [JsonPropertyName("invalid-type")]
    InvalidType,
    
    [JsonPropertyName("missing-definition")]
    MissingDefinition,
    
    [JsonPropertyName("relation-no-entry-point")]
    RelationNoEntrypoint,
    
    [JsonPropertyName("allowed-type-schema-10")]
    RequireSchema1_0,
    
    [JsonPropertyName("reserved-relation-keywords")]
    ReservedRelationKeywords,
    
    [JsonPropertyName("reserved-type-keywords")]
    ReservedTypeKeywords,
    
    [JsonPropertyName("schema-version-required")]
    SchemaVersionRequired,
    
    [JsonPropertyName("schema-version-unsupported")]
    SchemaVersionUnsupported,
    
    [JsonPropertyName("self-error")]
    SelfError,
    
    [JsonPropertyName("tupleuserset-not-direct")]
    TuplesetNotDirect,
    
    [JsonPropertyName("type-wildcard-relation")]
    TypeRestrictionCannotHaveWildcardAndRelation,
    
    [JsonPropertyName("invalid-relation-on-tupleset")]
    InvalidRelationOnTupleset,
    
    [JsonPropertyName("different-nested-condition-name")]
    DifferentNestedConditionName,
    
    [JsonPropertyName("multiple-modules-in-file")]
    MultipleModulesInFile
}
