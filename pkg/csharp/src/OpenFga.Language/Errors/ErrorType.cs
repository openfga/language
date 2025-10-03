using System.Text.Json.Serialization;

namespace OpenFga.Language.Errors;

/// <summary>
/// Enumeration of error types that can occur during OpenFGA language processing.
/// Distinguishes between syntax errors and validation errors.
/// </summary>
public enum ErrorType {
    /// <summary>Syntax error type for parsing errors.</summary>
    [JsonPropertyName("syntax")]
    Syntax,

    /// <summary>Validation error type for semantic validation errors.</summary>
    [JsonPropertyName("validation")]
    Validation
}