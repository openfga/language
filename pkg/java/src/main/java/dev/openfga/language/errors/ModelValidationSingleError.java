package dev.openfga.language.errors;

import lombok.*;

@Getter
@Setter
@NoArgsConstructor
public class ModelValidationSingleError extends ParsingError {

    private ValidationMetadata metadata;

    public ModelValidationSingleError(ErrorProperties properties, ValidationMetadata metadata) {
        super(getErrorType(metadata), properties);
        this.metadata = metadata;
    }

    private static String getErrorType(ValidationMetadata metadata) {
        if (metadata.getErrorType() != null) {
            return metadata.getErrorType().getValue();
        }
        return ErrorType.VALIDATION.getValue();
    }
}