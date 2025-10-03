namespace OpenFga.Language.Errors;

public class ModFileValidationSingleError : ParsingError {
    // Needed for JSON deserialization
    public ModFileValidationSingleError() { }

    public ModFileValidationSingleError(ErrorProperties properties) : base("validation", properties) {
    }
}