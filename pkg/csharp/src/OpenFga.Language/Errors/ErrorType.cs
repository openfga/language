using System.Text.Json.Serialization;

namespace OpenFga.Language.Errors;

public enum ErrorType {
    [JsonPropertyName("syntax")]
    Syntax,

    [JsonPropertyName("validation")]
    Validation
}