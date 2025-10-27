using SharpYaml.Serialization;
using System.Text.Json.Serialization;

namespace OpenFga.Language.Errors;

/// <summary>
/// Contains metadata information for validation errors, including symbol details and error context.
/// Used to provide additional context about validation errors such as the offending symbol and related entities.
/// </summary>
public class ValidationMetadata {
    /// <summary>
    /// Gets or sets the symbol that caused the validation error.
    /// </summary>
    [JsonPropertyName("symbol")]
    [YamlMember("symbol")]
    public string Symbol { get; set; } = string.Empty;

    /// <summary>
    /// Gets or sets the type of validation error that occurred.
    /// </summary>
    [JsonPropertyName("errorType")]
    [YamlMember("errorType")]
    public ValidationError ErrorType { get; set; }

    /// <summary>
    /// Gets or sets the relation name associated with the error (if applicable).
    /// </summary>
    [JsonPropertyName("relation")]
    [YamlMember("relation")]
    public string? Relation { get; set; }

    /// <summary>
    /// Gets or sets the type name associated with the error (if applicable).
    /// </summary>
    [JsonPropertyName("typeName")]
    [YamlMember("typeName")]
    public string? TypeName { get; set; }

    /// <summary>
    /// Gets or sets the condition name associated with the error (if applicable).
    /// </summary>
    [JsonPropertyName("condition")]
    [YamlMember("condition")]
    public string? Condition { get; set; }

    /// <summary>
    /// Initializes a new instance of the ValidationMetadata class.
    /// </summary>
    public ValidationMetadata() { }

    /// <summary>
    /// Initializes a new instance of the ValidationMetadata class with symbol and error type.
    /// </summary>
    /// <param name="symbol">The symbol that caused the error</param>
    /// <param name="error">The type of validation error</param>
    public ValidationMetadata(string symbol, ValidationError error) {
        Symbol = symbol;
        ErrorType = error;
    }

    /// <summary>
    /// Initializes a new instance of the ValidationMetadata class with full context.
    /// </summary>
    /// <param name="symbol">The symbol that caused the error</param>
    /// <param name="error">The type of validation error</param>
    /// <param name="relation">The relation name (if applicable)</param>
    /// <param name="typeName">The type name (if applicable)</param>
    /// <param name="condition">The condition name (if applicable)</param>
    public ValidationMetadata(string symbol, ValidationError error, string? relation, string? typeName, string? condition) {
        Symbol = symbol;
        ErrorType = error;
        Relation = relation;
        TypeName = typeName;
        Condition = condition;
    }
}