using System.Text.Json.Serialization;

namespace OpenFga.Language;

public class FgaModFile {
    [JsonPropertyName("schema")]
    public ModFileStringProperty? Schema { get; set; }

    [JsonPropertyName("contents")]
    public ModFileArrayProperty? Contents { get; set; }
}