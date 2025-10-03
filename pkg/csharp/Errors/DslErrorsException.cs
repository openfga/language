namespace OpenFgaLanguage.Errors;

public class DslErrorsException : Exception
{
    public IReadOnlyList<ParsingError> Errors { get; }

    public DslErrorsException(IEnumerable<ParsingError> errors) 
        : base(MessagesFromErrors(errors))
    {
        Errors = errors.ToList().AsReadOnly();
    }

    private static string MessagesFromErrors(IEnumerable<ParsingError> errors)
    {
        var delimiter = "\n\t* ";
        var errorsCount = errors.Count();
        var errorsPlural = errorsCount > 1 ? "s" : "";
        var prefix = $"{errorsCount} error{errorsPlural} occurred:{delimiter}";
        var suffix = "\n\n";
        
        return prefix + string.Join("\n\t* ", errors.Select(e => e.ToString())) + suffix;
    }
}
