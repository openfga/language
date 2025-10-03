namespace OpenFga.Language.Errors;

public class DslErrorsException(IEnumerable<ParsingError> errors) 
    : Exception(MessagesFromErrors(errors)) {

    public IReadOnlyList<ParsingError> Errors { get; } = errors.ToList().AsReadOnly();

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
