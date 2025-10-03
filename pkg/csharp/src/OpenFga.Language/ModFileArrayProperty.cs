using OpenFga.Language.Errors;
using System.Text.Json.Serialization;

namespace OpenFga.Language;

public class ModFileArrayProperty {
    [JsonPropertyName("value")]
    public List<ModFileStringProperty> Value { get; set; } = [];

    [JsonPropertyName("line")]
    public StartEnd? Line { get; set; }

    [JsonPropertyName("column")]
    public StartEnd? Column { get; set; }
}