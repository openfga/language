using Antlr4.Runtime;

namespace OpenFga.Language.Errors;

public class SyntaxError : ParsingError
{
    public Metadata? Metadata { get; set; }
    public RecognitionException? Cause { get; set; }

    public SyntaxError() { }

    public SyntaxError(ErrorProperties properties, Metadata? metadata, RecognitionException? cause) 
        : base("syntax", properties)
    {
        Metadata = metadata;
        Cause = cause;
    }
}
