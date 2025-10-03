using System.Text.Json.Serialization;
using OpenFga.Language.Errors;

namespace OpenFga.Language.ModFile;

public class ModFileStringProperty
{
    [JsonPropertyName("value")]
    public string Value { get; set; } = string.Empty;

    [JsonPropertyName("line")]
    public StartEnd? Line { get; set; }

    [JsonPropertyName("column")]
    public StartEnd? Column { get; set; }
}
