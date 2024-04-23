package dev.openfga.language.errors;

public class ModuleTransformationSingleError extends ParsingError {
    private ValidationMetadata metadata;

    public ModuleTransformationSingleError() {}

    public ModuleTransformationSingleError(ErrorProperties properties, ValidationMetadata metadata) {
        super("transformation-error", properties);
        this.metadata = metadata;
    }

    public ValidationMetadata getMetadata() {
        return metadata;
    }

    public void setMetadata(ValidationMetadata metadata) {
        this.metadata = metadata;
    }
}

