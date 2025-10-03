using SharpYaml.Serialization;
using System.Text.Json.Serialization;

namespace OpenFga.Language.Errors;

public class ValidationMetadata {
    [JsonPropertyName("symbol")]
    [YamlMember("symbol")]
    public string Symbol { get; set; } = string.Empty;

    [JsonPropertyName("errorType")]
    [YamlMember("errorType")]
    public ValidationError ErrorType { get; set; }

    [JsonPropertyName("relation")]
    [YamlMember("relation")]
    public string? Relation { get; set; }

    [JsonPropertyName("typeName")]
    [YamlMember("typeName")]
    public string? TypeName { get; set; }

    [JsonPropertyName("condition")]
    [YamlMember("condition")]
    public string? Condition { get; set; }

    public ValidationMetadata() { }

    public ValidationMetadata(string symbol, ValidationError error) {
        Symbol = symbol;
        ErrorType = error;
    }

    public ValidationMetadata(string symbol, ValidationError error, string? relation, string? typeName, string? condition) {
        Symbol = symbol;
        ErrorType = error;
        Relation = relation;
        TypeName = typeName;
        Condition = condition;
    }
}