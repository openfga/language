using System.Text.Json.Serialization;
using YamlDotNet.Serialization;

namespace OpenFga.Language.Errors;

public class ValidationMetadata
{
    [JsonPropertyName("symbol")]
    [YamlMember(Alias = "symbol")]
    public string Symbol { get; set; } = string.Empty;

    [JsonPropertyName("errorType")]
    [YamlMember(Alias = "errorType")]
    [SharpYaml.Serialization.YamlMember("errorType")]
    public ValidationError ErrorType { get; set; }

    [JsonPropertyName("relation")]
    [YamlMember(Alias = "relation")]
    public string? Relation { get; set; }

    [JsonPropertyName("typeName")]
    [YamlMember(Alias = "typeName")]
    public string? TypeName { get; set; }

    [JsonPropertyName("condition")]
    [YamlMember(Alias = "condition")]
    public string? Condition { get; set; }

    public ValidationMetadata() { }

    public ValidationMetadata(string symbol, ValidationError error)
    {
        Symbol = symbol;
        ErrorType = error;
    }

    public ValidationMetadata(string symbol, ValidationError error, string? relation, string? typeName, string? condition)
    {
        Symbol = symbol;
        ErrorType = error;
        Relation = relation;
        TypeName = typeName;
        Condition = condition;
    }
}
