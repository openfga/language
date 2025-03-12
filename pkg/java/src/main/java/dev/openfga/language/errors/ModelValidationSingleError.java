package dev.openfga.language.errors;

public class ModelValidationSingleError extends ParsingError {

    private ValidationMetadata metadata;

    // Needed for Jackson deserialization
    public ModelValidationSingleError() {}

    public ModelValidationSingleError(ErrorProperties properties, ValidationMetadata metadata) {
        super("validation", properties);
        this.metadata = metadata;
    }

    public ValidationMetadata getMetadata() {
        return metadata;
    }

    public void setMetadata(ValidationMetadata metadata) {
        this.metadata = metadata;
    }
}
