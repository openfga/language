
using System.Runtime.Serialization;
using System.Text.Json.Serialization;
using SharpYaml.Serialization;

namespace OpenFga.Language.Errors;

public enum ValidationError
{
    [JsonPropertyName("allowed-type-not-valid-on-schema-1_0")]
    [YamlRemap("allowed-type-not-valid-on-schema-1_0")]
    AllowedTypesNotValidOnSchema1_0,
    
    [JsonPropertyName("assignable-relation-must-have-type")]
    [YamlRemap("assignable-relation-must-have-type")]
    AssignableRelationsMustHaveType,
    
    [JsonPropertyName("condition-not-defined")]
    [YamlRemap("condition-not-defined")]
    ConditionNotDefined,
    
    [JsonPropertyName("condition-not-used")]
    [YamlRemap("condition-not-used")]
    ConditionNotUsed,
    
    [JsonPropertyName("duplicated-error")]
    [YamlRemap("duplicated-error")]
    DuplicatedError,
    
    [JsonPropertyName("invalid-name")]
    [YamlRemap("invalid-name")]
    InvalidName,
    
    [JsonPropertyName("invalid-relation-type")]
    [YamlRemap("invalid-relation-type")]
    InvalidRelationType,
    
    [JsonPropertyName("invalid-schema")]
    [YamlRemap("invalid-schema")]
    InvalidSchema,
    
    [JsonPropertyName("invalid-syntax")]
    [YamlRemap("invalid-syntax")]
    InvalidSyntax,
    
    [JsonPropertyName("invalid-type")]
    [YamlRemap("invalid-type")]
    InvalidType,
    
    [JsonPropertyName("missing-definition")]
    [YamlRemap("missing-definition")]
    MissingDefinition,
    
    [JsonPropertyName("relation-no-entry-point")]
    [YamlRemap("relation-no-entry-point")]
    RelationNoEntrypoint,
    
    [JsonPropertyName("allowed-type-schema-10")]
    RequireSchema1_0,
    
    [JsonPropertyName("reserved-relation-keywords")]
    [YamlRemap("reserved-relation-keywords")]
    ReservedRelationKeywords,
    
    [JsonPropertyName("reserved-type-keywords")]
    [YamlRemap("reserved-type-keywords")]
    ReservedTypeKeywords,
    
    [JsonPropertyName("schema-version-required")]
    [YamlRemap("schema-version-required")]
    SchemaVersionRequired,
    
    [JsonPropertyName("schema-version-unsupported")]
    [YamlRemap("schema-version-unsupported")]
    SchemaVersionUnsupported,
    
    [JsonPropertyName("self-error")]
    [YamlRemap("self-error")]
    SelfError,
    
    [JsonPropertyName("tupleuserset-not-direct")]
    [YamlRemap("tupleuserset-not-direct")]
    TuplesetNotDirect,
    
    [JsonPropertyName("type-wildcard-relation")]
    [YamlRemap("type-wildcard-relation")]
    TypeRestrictionCannotHaveWildcardAndRelation,
    
    [JsonPropertyName("invalid-relation-on-tupleset")]
    [YamlRemap("invalid-relation-on-tupleset")]
    InvalidRelationOnTupleset,
    
    [JsonPropertyName("different-nested-condition-name")]
    [YamlRemap("different-nested-condition-name")]
    DifferentNestedConditionName,
    
    [JsonPropertyName("multiple-modules-in-file")]
    [YamlRemap("multiple-modules-in-file")]
    MultipleModulesInFile
}
