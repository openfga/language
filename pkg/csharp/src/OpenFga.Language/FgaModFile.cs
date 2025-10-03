using System.Text.Json.Serialization;

namespace OpenFga.Language.ModFile;

public class FgaModFile
{
    [JsonPropertyName("schema")]
    public ModFileStringProperty? Schema { get; set; }

    [JsonPropertyName("contents")]
    public ModFileArrayProperty? Contents { get; set; }
}
