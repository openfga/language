using Antlr4.Runtime;
using OpenFga.Language.Errors;

namespace OpenFga.Language;

/// <summary>
/// Error listener for OpenFGA DSL parsing that collects syntax errors during lexer and parser operations.
/// Implements ANTLR error listener interfaces to capture and store parsing errors.
/// </summary>
public class OpenFgaDslErrorListener : IAntlrErrorListener<int>, IAntlrErrorListener<IToken> {
    private readonly List<SyntaxError> _errors = new();

    /// <summary>
    /// Gets the read-only list of syntax errors collected during parsing.
    /// </summary>
    public IReadOnlyList<SyntaxError> Errors => _errors.AsReadOnly();

    /// <summary>
    /// Handles syntax errors from the lexer (token recognition errors).
    /// </summary>
    /// <param name="output">The output writer</param>
    /// <param name="recognizer">The recognizer that detected the error</param>
    /// <param name="offendingSymbol">The offending symbol</param>
    /// <param name="line">The line number where the error occurred</param>
    /// <param name="column">The column number where the error occurred</param>
    /// <param name="msg">The error message</param>
    /// <param name="e">The recognition exception</param>
    public void SyntaxError(TextWriter output, IRecognizer recognizer, int offendingSymbol, int line, int column, string msg, RecognitionException e) {
        // For lexer errors
        var properties = new ErrorProperties(
            new StartEnd(line - 1, line - 1),
            new StartEnd(column, column),
            msg);

        _errors.Add(new SyntaxError(properties, null, e));
    }

    /// <summary>
    /// Handles syntax errors from the parser (grammar recognition errors).
    /// </summary>
    /// <param name="output">The output writer</param>
    /// <param name="recognizer">The recognizer that detected the error</param>
    /// <param name="offendingSymbol">The offending token</param>
    /// <param name="line">The line number where the error occurred</param>
    /// <param name="column">The column number where the error occurred</param>
    /// <param name="msg">The error message</param>
    /// <param name="e">The recognition exception</param>
    public void SyntaxError(TextWriter output, IRecognizer recognizer, IToken offendingSymbol, int line, int column, string msg, RecognitionException e) {
        // For parser errors
        // line is one based, i.e. the first line will be 1
        // column is zero based, i.e. the first column will 0
        Metadata? metadata = null;
        var columnOffset = 0;

        if (offendingSymbol != null) {
            metadata = new Metadata(offendingSymbol.Text);
            columnOffset = metadata.Symbol.Length;
        }

        var properties = new ErrorProperties(
            new StartEnd(line - 1, line - 1),
            new StartEnd(column, column + columnOffset),
            msg);

        _errors.Add(new SyntaxError(properties, metadata, e));
    }
}