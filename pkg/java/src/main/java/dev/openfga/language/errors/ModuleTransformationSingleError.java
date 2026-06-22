package dev.openfga.language.errors;

public class ModuleTransformationSingleError extends ParsingError {
    // Needed for Jackson deserialization
    public ModuleTransformationSingleError() {}

    public ModuleTransformationSingleError(ErrorProperties properties) {
        super("transformation", properties);
    }

    public ModuleTransformationSingleError(ErrorProperties properties, String file) {
        super("transformation", properties);
        setFile(file);
    }
}
