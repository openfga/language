using OpenFga.Language.Errors;
using System.Text.Json.Serialization;

namespace OpenFga.Language;

/// <summary>
/// Represents a string property in a module file with associated line and column information.
/// Used for tracking the location of string values within the source file.
/// </summary>
public class ModFileStringProperty {
    /// <summary>
    /// Gets or sets the string value of the property.
    /// </summary>
    [JsonPropertyName("value")]
    public string Value { get; set; } = string.Empty;

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