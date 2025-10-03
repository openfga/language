namespace OpenFga.Language.Errors;

/// <summary>
/// Represents a single validation error that occurred during FGA module file validation.
/// Used for validation errors specific to .fga.mod file parsing and validation.
/// </summary>
public class ModFileValidationSingleError : ParsingError {
    /// <summary>
    /// Initializes a new instance of the ModFileValidationSingleError class.
    /// Needed for JSON deserialization.
    /// </summary>
    public ModFileValidationSingleError() { }

    /// <summary>
    /// Initializes a new instance of the ModFileValidationSingleError class with error properties.
    /// </summary>
    /// <param name="properties">The error properties including location and message</param>
    public ModFileValidationSingleError(ErrorProperties properties) : base("validation", properties) {
    }
}