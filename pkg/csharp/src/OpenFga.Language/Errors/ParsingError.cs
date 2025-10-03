using System.Text.Json.Serialization;
using YamlDotNet.Serialization;

namespace OpenFga.Language.Errors;

public abstract class ParsingError : SimpleError
{
    [JsonPropertyName("line")]
    [YamlMember(Alias = "line")]
    public StartEnd? Line { get; set; }

    [JsonPropertyName("column")]
    [YamlMember(Alias = "column")]
    public StartEnd? Column { get; set; }

    [JsonPropertyName("fullMessage")]
    [YamlMember(Alias = "fullMessage")]
    public string FullMessage { get; set; } = string.Empty;

    protected ParsingError() { }

    protected ParsingError(string type, ErrorProperties properties) : base(properties.Message)
    {
        Line = properties.Line;
        Column = properties.Column;
        FullMessage = properties.GetFullMessage(type);
    }

    public StartEnd? GetLine(int offset = 0)
    {
        if (Line == null)
        {
            return null;
        }
        return Line.WithOffset(offset);
    }

    public StartEnd? GetColumn(int offset = 0)
    {
        if (Column == null)
        {
            return null;
        }
        return Column.WithOffset(offset);
    }

    public override string ToString()
    {
        return string.IsNullOrEmpty(FullMessage) ? Message : FullMessage;
    }
}
