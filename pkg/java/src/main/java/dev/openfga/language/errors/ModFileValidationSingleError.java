package dev.openfga.language.errors;

public class ModFileValidationSingleError extends ParsingError {
    // Needed for Jackson deserialization
    public ModFileValidationSingleError() {}

    public ModFileValidationSingleError(ErrorProperties properties) {
        super("validation", properties);
    }
}
