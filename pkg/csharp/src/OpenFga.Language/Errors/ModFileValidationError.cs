namespace OpenFga.Language.Errors;

/// <summary>
/// Exception thrown when FGA module file validation encounters one or more validation errors.
/// Contains detailed information about all validation errors encountered during module file parsing.
/// </summary>
/// <param name="errors">Collection of validation errors encountered during module file parsing</param>
public class ModFileValidationError(List<ModFileValidationSingleError> errors)
    : Exception(MessagesFromErrors(errors)) {

    /// <summary>
    /// Gets the list of validation errors that caused this exception.
    /// </summary>
    public List<ModFileValidationSingleError> Errors { get; } = errors;

    private static string MessagesFromErrors(IEnumerable<object> errors) {
        var delimiter = "\n\t* ";
        var errorsCount = errors.Count();
        var errorsPlural = errorsCount > 1 ? "s" : "";
        var prefix = $"{errorsCount} error{errorsPlural} occurred:{delimiter}";
        var suffix = "\n\n";

        return prefix + string.Join("\n\t* ", errors.Select(e => e.ToString())) + suffix;
    }
}