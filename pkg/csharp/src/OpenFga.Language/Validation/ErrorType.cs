using System.Text.Json.Serialization;

namespace OpenFga.Language.Validation;

public enum ErrorType
{
    [JsonPropertyName("syntax")]
    Syntax,
    
    [JsonPropertyName("validation")]
    Validation
}
