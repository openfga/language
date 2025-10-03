using System.Text.Json.Serialization;

namespace OpenFga.Language.Validation;

public class ValidationMetadata
{
    [JsonPropertyName("symbol")]
    public string Symbol { get; set; } = string.Empty;

    [JsonPropertyName("error")]
    public ValidationError Error { get; set; }

    [JsonPropertyName("relation")]
    public string? Relation { get; set; }

    [JsonPropertyName("type")]
    public string? Type { get; set; }

    [JsonPropertyName("condition")]
    public string? Condition { get; set; }

    public ValidationMetadata() { }

    public ValidationMetadata(string symbol, ValidationError error)
    {
        Symbol = symbol;
        Error = error;
    }

    public ValidationMetadata(string symbol, ValidationError error, string? relation, string? type, string? condition)
    {
        Symbol = symbol;
        Error = error;
        Relation = relation;
        Type = type;
        Condition = condition;
    }
}
