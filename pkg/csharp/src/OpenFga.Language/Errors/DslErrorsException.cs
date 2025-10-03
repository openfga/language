namespace OpenFga.Language.Errors;

/// <summary>
/// Exception thrown when DSL parsing encounters one or more syntax errors.
/// Contains detailed information about all parsing errors encountered.
/// </summary>
/// <param name="errors">Collection of parsing errors encountered during DSL parsing</param>
public class DslErrorsException(IEnumerable<ParsingError> errors)
    : Exception(MessagesFromErrors(errors)) {

    /// <summary>
    /// Gets the read-only list of parsing errors that caused this exception.
    /// </summary>
    public IReadOnlyList<ParsingError> Errors { get; } = errors.ToList().AsReadOnly();

    private static string MessagesFromErrors(IEnumerable<ParsingError> errors) {
        var delimiter = "\n\t* ";
        var errorsCount = errors.Count();
        var errorsPlural = errorsCount > 1 ? "s" : "";
        var prefix = $"{errorsCount} error{errorsPlural} occurred:{delimiter}";
        var suffix = "\n\n";

        return prefix + string.Join("\n\t* ", errors.Select(e => e.ToString())) + suffix;
    }
}