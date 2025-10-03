using System.Text.Json.Serialization;
using OpenFga.Language.Errors;

namespace OpenFga.Language.ModFile;

public class ModFileArrayProperty
{
    [JsonPropertyName("value")]
    public List<ModFileStringProperty> Value { get; set; } = [];

    [JsonPropertyName("line")]
    public StartEnd? Line { get; set; }

    [JsonPropertyName("column")]
    public StartEnd? Column { get; set; }
}
