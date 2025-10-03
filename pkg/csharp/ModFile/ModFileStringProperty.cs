using System.Text.Json.Serialization;
using OpenFgaLanguage.Errors;

namespace OpenFgaLanguage.ModFile;

public class ModFileStringProperty
{
    public const string JsonPropertyValue = "value";
    public const string JsonPropertyLine = "line";
    public const string JsonPropertyColumn = "column";

    [JsonPropertyName(JsonPropertyValue)]
    public string Value { get; set; } = string.Empty;

    [JsonPropertyName(JsonPropertyLine)]
    public StartEnd? Line { get; set; }

    [JsonPropertyName(JsonPropertyColumn)]
    public StartEnd? Column { get; set; }

    public ModFileStringProperty() { }

    public ModFileStringProperty SetValue(string value)
    {
        Value = value;
        return this;
    }

    public ModFileStringProperty SetLine(StartEnd? line)
    {
        Line = line;
        return this;
    }

    public ModFileStringProperty SetColumn(StartEnd? column)
    {
        Column = column;
        return this;
    }
}
