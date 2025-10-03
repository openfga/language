namespace OpenFga.Language.Errors;

/// <summary>
/// Contains properties that describe an error including location information and error message.
/// Used as a base class for various error types to provide consistent error reporting.
/// </summary>
public class ErrorProperties {
    /// <summary>
    /// Gets or sets the line range where the error occurred.
    /// </summary>
    public StartEnd? Line { get; set; }

    /// <summary>
    /// Gets or sets the column range where the error occurred.
    /// </summary>
    public StartEnd? Column { get; set; }

    /// <summary>
    /// Gets or sets the error message describing what went wrong.
    /// </summary>
    public string Message { get; set; } = string.Empty;

    /// <summary>
    /// Initializes a new instance of the ErrorProperties class.
    /// </summary>
    public ErrorProperties() { }

    /// <summary>
    /// Initializes a new instance of the ErrorProperties class with location and message.
    /// </summary>
    /// <param name="line">The line range where the error occurred</param>
    /// <param name="column">The column range where the error occurred</param>
    /// <param name="message">The error message</param>
    public ErrorProperties(StartEnd? line, StartEnd? column, string message) {
        Line = line;
        Column = column;
        Message = message;
    }

    /// <summary>
    /// Gets a formatted error message including the error type and location information.
    /// </summary>
    /// <param name="type">The type of error (e.g., "syntax", "validation")</param>
    /// <returns>A formatted error message</returns>
    internal string GetFullMessage(string type) {
        if (Line != null && Column != null) {
            return $"{type} error at line={Line.Start}, column={Column.Start}: {Message}";
        }
        else {
            return $"{type} error: {Message}";
        }
    }
}