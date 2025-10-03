using System.Text.Json.Serialization;
using OpenFga.Language.Errors;

namespace OpenFga.Language.ModFile;

public class ModFileArrayProperty
{
    public const string JsonPropertyValue = "value";
    public const string JsonPropertyLine = "line";
    public const string JsonPropertyColumn = "column";

    [JsonPropertyName(JsonPropertyValue)]
    public List<ModFileStringProperty> Value { get; set; } = new();

    [JsonPropertyName(JsonPropertyLine)]
    public StartEnd? Line { get; set; }

    [JsonPropertyName(JsonPropertyColumn)]
    public StartEnd? Column { get; set; }

    public ModFileArrayProperty() { }

    public ModFileArrayProperty SetValue(List<ModFileStringProperty> value)
    {
        Value = value;
        return this;
    }

    public ModFileArrayProperty SetLine(StartEnd? line)
    {
        Line = line;
        return this;
    }

    public ModFileArrayProperty SetColumn(StartEnd? column)
    {
        Column = column;
        return this;
    }
}
