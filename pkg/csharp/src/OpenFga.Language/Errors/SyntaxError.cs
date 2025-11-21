using Antlr4.Runtime;

namespace OpenFga.Language.Errors;

/// <summary>
/// Represents a syntax error encountered during DSL parsing.
/// Contains metadata about the error location and the underlying recognition exception.
/// </summary>
public class SyntaxError : ParsingError {
    /// <summary>
    /// Gets or sets the metadata associated with this syntax error.
    /// </summary>
    public Metadata? Metadata { get; set; }

    /// <summary>
    /// Gets or sets the underlying recognition exception that caused this syntax error.
    /// </summary>
    public RecognitionException? Cause { get; set; }

    /// <summary>
    /// Initializes a new instance of the SyntaxError class.
    /// </summary>
    public SyntaxError() { }

    /// <summary>
    /// Initializes a new instance of the SyntaxError class with specified properties.
    /// </summary>
    /// <param name="properties">Error properties including location and message</param>
    /// <param name="metadata">Metadata associated with the error</param>
    /// <param name="cause">The underlying recognition exception</param>
    public SyntaxError(ErrorProperties properties, Metadata? metadata, RecognitionException? cause)
        : base("syntax", properties) {
        Metadata = metadata;
        Cause = cause;
    }
}