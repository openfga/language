
using System.Runtime.Serialization;
using System.Text.Json.Serialization;
using YamlDotNet.Serialization;

namespace OpenFga.Language.Errors;

public enum ValidationError
{
    [JsonPropertyName("allowed-type-not-valid-on-schema-1_0")]
    [SharpYaml.Serialization.YamlRemap("allowed-type-not-valid-on-schema-1_0")]
    AllowedTypesNotValidOnSchema1_0,
    
    [JsonPropertyName("assignable-relation-must-have-type")]
    [SharpYaml.Serialization.YamlRemap("assignable-relation-must-have-type")]
    AssignableRelationsMustHaveType,
    
    [JsonPropertyName("condition-not-defined")]
    [SharpYaml.Serialization.YamlRemap("condition-not-defined")]
    ConditionNotDefined,
    
    [JsonPropertyName("condition-not-used")]
    [SharpYaml.Serialization.YamlRemap("condition-not-used")]
    ConditionNotUsed,
    
    [JsonPropertyName("duplicated-error")]
    [SharpYaml.Serialization.YamlRemap("duplicated-error")]
    DuplicatedError,
    
    [JsonPropertyName("invalid-name")]
    [SharpYaml.Serialization.YamlRemap("invalid-name")]
    InvalidName,
    
    [JsonPropertyName("invalid-relation-type")]
    [SharpYaml.Serialization.YamlRemap("invalid-relation-type")]
    InvalidRelationType,
    
    [JsonPropertyName("invalid-schema")]
    [SharpYaml.Serialization.YamlRemap("invalid-schema")]
    InvalidSchema,
    
    [JsonPropertyName("invalid-syntax")]
    [SharpYaml.Serialization.YamlRemap("invalid-syntax")]
    InvalidSyntax,
    
    [JsonPropertyName("invalid-type")]
    [SharpYaml.Serialization.YamlRemap("invalid-type")]
    InvalidType,
    
    [JsonPropertyName("missing-definition")]
    [SharpYaml.Serialization.YamlRemap("missing-definition")]
    MissingDefinition,
    
    [JsonPropertyName("relation-no-entry-point")]
    [SharpYaml.Serialization.YamlRemap("relation-no-entry-point")]
    RelationNoEntrypoint,
    
    [JsonPropertyName("allowed-type-schema-10")]
    RequireSchema1_0,
    
    [JsonPropertyName("reserved-relation-keywords")]
    [SharpYaml.Serialization.YamlRemap("reserved-relation-keywords")]
    ReservedRelationKeywords,
    
    [JsonPropertyName("reserved-type-keywords")]
    [SharpYaml.Serialization.YamlRemap("reserved-type-keywords")]
    ReservedTypeKeywords,
    
    [JsonPropertyName("schema-version-required")]
    [SharpYaml.Serialization.YamlRemap("schema-version-required")]
    SchemaVersionRequired,
    
    [JsonPropertyName("schema-version-unsupported")]
    [SharpYaml.Serialization.YamlRemap("schema-version-unsupported")]
    SchemaVersionUnsupported,
    
    [JsonPropertyName("self-error")]
    [SharpYaml.Serialization.YamlRemap("self-error")]
    SelfError,
    
    [JsonPropertyName("tupleuserset-not-direct")]
    [SharpYaml.Serialization.YamlRemap("tupleuserset-not-direct")]
    TuplesetNotDirect,
    
    [JsonPropertyName("type-wildcard-relation")]
    [SharpYaml.Serialization.YamlRemap("type-wildcard-relation")]
    TypeRestrictionCannotHaveWildcardAndRelation,
    
    [JsonPropertyName("invalid-relation-on-tupleset")]
    [SharpYaml.Serialization.YamlRemap("invalid-relation-on-tupleset")]
    InvalidRelationOnTupleset,
    
    [JsonPropertyName("different-nested-condition-name")]
    [SharpYaml.Serialization.YamlRemap("different-nested-condition-name")]
    DifferentNestedConditionName,
    
    [JsonPropertyName("multiple-modules-in-file")]
    [SharpYaml.Serialization.YamlRemap("multiple-modules-in-file")]
    MultipleModulesInFile
}
