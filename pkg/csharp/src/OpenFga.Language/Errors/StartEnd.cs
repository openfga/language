using SharpYaml.Serialization;
using System.Text.Json.Serialization;

namespace OpenFga.Language.Errors;

/// <summary>
/// Represents a range with start and end positions, typically used for line and column ranges in error reporting.
/// Provides methods for offset calculations and equality comparisons.
/// </summary>
public sealed class StartEnd {
    /// <summary>
    /// Gets or sets the start position of the range.
    /// </summary>
    [JsonPropertyName("start")]
    [YamlMember("start")]
    public int Start { get; set; }

    /// <summary>
    /// Gets or sets the end position of the range.
    /// </summary>
    [JsonPropertyName("end")]
    [YamlMember("end")]
    public int End { get; set; }

    /// <summary>
    /// Initializes a new instance of the StartEnd class.
    /// Needed for JSON deserialization.
    /// </summary>
    public StartEnd() { }

    /// <summary>
    /// Initializes a new instance of the StartEnd class with start and end positions.
    /// </summary>
    /// <param name="start">The start position</param>
    /// <param name="end">The end position</param>
    public StartEnd(int start, int end) {
        Start = start;
        End = end;
    }

    /// <summary>
    /// Creates a new StartEnd instance with the specified offset applied to both start and end positions.
    /// </summary>
    /// <param name="offset">The offset to apply</param>
    /// <returns>A new StartEnd instance with the offset applied</returns>
    public StartEnd WithOffset(int offset) {
        return new StartEnd(Start + offset, End + offset);
    }

    /// <summary>
    /// Determines whether the specified object is equal to the current StartEnd instance.
    /// </summary>
    /// <param name="obj">The object to compare</param>
    /// <returns>True if the objects are equal, false otherwise</returns>
    public override bool Equals(object? obj) {
        if (obj is StartEnd other) {
            return Start == other.Start && End == other.End;
        }
        return false;
    }

    /// <summary>
    /// Returns the hash code for the current StartEnd instance.
    /// </summary>
    /// <returns>A hash code for the current instance</returns>
    public override int GetHashCode() {
        return HashCode.Combine(Start, End);
    }
}