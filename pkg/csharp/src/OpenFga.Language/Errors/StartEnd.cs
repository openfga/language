using System.Text.Json.Serialization;
using SharpYaml.Serialization;

namespace OpenFga.Language.Errors;

public sealed class StartEnd
{
    [JsonPropertyName("start")]
    [YamlMember("start")]
    public int Start { get; set; }

    [JsonPropertyName("end")]
    [YamlMember("end")]
    public int End { get; set; }

    // Needed for JSON deserialization
    public StartEnd() { }

    public StartEnd(int start, int end)
    {
        Start = start;
        End = end;
    }

    public StartEnd WithOffset(int offset)
    {
        return new StartEnd(Start + offset, End + offset);
    }

    public override bool Equals(object? obj)
    {
        if (obj is StartEnd other)
        {
            return Start == other.Start && End == other.End;
        }
        return false;
    }

    public override int GetHashCode()
    {
        return HashCode.Combine(Start, End);
    }
}
