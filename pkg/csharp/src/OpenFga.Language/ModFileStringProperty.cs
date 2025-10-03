using OpenFga.Language.Errors;
using System.Text.Json.Serialization;

namespace OpenFga.Language;

public class ModFileStringProperty {
    [JsonPropertyName("value")]
    public string Value { get; set; } = string.Empty;

    [JsonPropertyName("line")]
    public StartEnd? Line { get; set; }

    [JsonPropertyName("column")]
    public StartEnd? Column { get; set; }
}