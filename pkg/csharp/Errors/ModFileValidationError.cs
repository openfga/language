namespace OpenFgaLanguage.Errors;

public class ModFileValidationError : Exception
{
    public List<ModFileValidationSingleError> Errors { get; }

    public ModFileValidationError(List<ModFileValidationSingleError> errors) 
        : base(MessagesFromErrors(errors))
    {
        Errors = errors;
    }

    private static string MessagesFromErrors(IEnumerable<object> errors)
    {
        var delimiter = "\n\t* ";
        var errorsCount = errors.Count();
        var errorsPlural = errorsCount > 1 ? "s" : "";
        var prefix = $"{errorsCount} error{errorsPlural} occurred:{delimiter}";
        var suffix = "\n\n";
        
        return prefix + string.Join("\n\t* ", errors.Select(e => e.ToString())) + suffix;
    }
}
