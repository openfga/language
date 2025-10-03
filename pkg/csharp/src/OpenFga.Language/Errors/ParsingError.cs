using SharpYaml.Serialization;
using System.Text.Json.Serialization;

namespace OpenFga.Language.Errors;

/// <summary>
/// Abstract base class for parsing errors that provides location information and formatted error messages.
/// Extends SimpleError with line/column information and full message formatting capabilities.
/// </summary>
public abstract class ParsingError : SimpleError {
    /// <summary>
    /// Gets or sets the line range where the parsing error occurred.
    /// </summary>
    [JsonPropertyName("line")]
    [YamlMember("line")]
    public StartEnd? Line { get; set; }

    /// <summary>
    /// Gets or sets the column range where the parsing error occurred.
    /// </summary>
    [JsonPropertyName("column")]
    [YamlMember("column")]
    public StartEnd? Column { get; set; }

    /// <summary>
    /// Gets or sets the full formatted error message including location information.
    /// </summary>
    [JsonPropertyName("fullMessage")]
    [YamlMember("fullMessage")]
    public string FullMessage { get; set; } = string.Empty;

    /// <summary>
    /// Initializes a new instance of the ParsingError class.
    /// </summary>
    protected ParsingError() { }

    /// <summary>
    /// Initializes a new instance of the ParsingError class with error type and properties.
    /// </summary>
    /// <param name="type">The type of parsing error</param>
    /// <param name="properties">The error properties including location and message</param>
    protected ParsingError(string type, ErrorProperties properties) : base(properties.Message) {
        Line = properties.Line;
        Column = properties.Column;
        FullMessage = properties.GetFullMessage(type);
    }

    /// <summary>
    /// Gets the line range with an optional offset applied.
    /// </summary>
    /// <param name="offset">The offset to apply to the line range</param>
    /// <returns>The line range with offset applied, or null if no line information is available</returns>
    public StartEnd? GetLine(int offset = 0) {
        if (Line == null) {
            return null;
        }
        return Line.WithOffset(offset);
    }

    /// <summary>
    /// Gets the column range with an optional offset applied.
    /// </summary>
    /// <param name="offset">The offset to apply to the column range</param>
    /// <returns>The column range with offset applied, or null if no column information is available</returns>
    public StartEnd? GetColumn(int offset = 0) {
        if (Column == null) {
            return null;
        }
        return Column.WithOffset(offset);
    }

    /// <summary>
    /// Returns a string representation of the parsing error.
    /// </summary>
    /// <returns>The full message if available, otherwise the base message</returns>
    public override string ToString() {
        return string.IsNullOrEmpty(FullMessage) ? Message : FullMessage;
    }
}