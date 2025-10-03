using Antlr4.Runtime;
using OpenFga.Language.Errors;

namespace OpenFga.Language;

public class OpenFgaDslErrorListener : IAntlrErrorListener<int>, IAntlrErrorListener<IToken>
{
    private readonly List<SyntaxError> _errors = new();

    public IReadOnlyList<SyntaxError> Errors => _errors.AsReadOnly();

    public void SyntaxError(TextWriter output, IRecognizer recognizer, int offendingSymbol, int line, int column, string msg, RecognitionException e)
    {
        // For lexer errors
        var properties = new ErrorProperties(
            new StartEnd(line - 1, line - 1), 
            new StartEnd(column, column), 
            msg);
        
        _errors.Add(new SyntaxError(properties, null, e));
    }

    public void SyntaxError(TextWriter output, IRecognizer recognizer, IToken offendingSymbol, int line, int column, string msg, RecognitionException e)
    {
        // For parser errors
        // line is one based, i.e. the first line will be 1
        // column is zero based, i.e. the first column will 0
        Metadata? metadata = null;
        var columnOffset = 0;

        if (offendingSymbol != null)
        {
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
