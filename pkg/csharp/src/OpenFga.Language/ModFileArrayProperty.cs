using OpenFga.Language.Errors;
using System.Text.Json.Serialization;

namespace OpenFga.Language;

/// <summary>
/// Represents an array property in a module file with associated line and column information.
/// Used for tracking the location of array values within the source file.
/// </summary>
public class ModFileArrayProperty {
    /// <summary>
    /// Gets or sets the list of string properties that make up the array value.
    /// </summary>
    [JsonPropertyName("value")]
    public List<ModFileStringProperty> Value { get; set; } = [];

    /// <summary>
    /// Gets or sets the line range information for this property in the source file.
    /// </summary>
    [JsonPropertyName("line")]
    public StartEnd? Line { get; set; }

    /// <summary>
    /// Gets or sets the column range information for this property in the source file.
    /// </summary>
    [JsonPropertyName("column")]
    public StartEnd? Column { get; set; }
}