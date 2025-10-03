using System.Text.Json.Serialization;

namespace OpenFga.Language;

/// <summary>
/// Represents a parsed FGA module file (.fga.mod) containing schema version and content references.
/// This class holds the structured representation of a module definition file.
/// </summary>
public class FgaModFile {
    /// <summary>
    /// Gets or sets the schema version property of the module file.
    /// </summary>
    [JsonPropertyName("schema")]
    public ModFileStringProperty? Schema { get; set; }

    /// <summary>
    /// Gets or sets the contents array property containing references to FGA files.
    /// </summary>
    [JsonPropertyName("contents")]
    public ModFileArrayProperty? Contents { get; set; }
}